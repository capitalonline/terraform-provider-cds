package cds

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/security_group"
	"github.com/capitalonline/cds-gic-sdk-go/security_group_rule"
	u "terraform-provider-cds/cds/utils"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCdsSecurityGroup() *schema.Resource {
	return &schema.Resource{

		Create: resourceCdsSecurityGroupCreate,
		Read:   resourceCdsSecurityGroupRead,
		Update: resourceCdsSecurityGroupUpdate,
		Delete: resourceCdsSecurityGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: u.ValidateStringLengthInRange(1, 128),
				Description:  "Name of the security group to be queried.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: u.ValidateStringLengthInRange(1, 256),
				Description:  "Description of the security group.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: u.ValidateStringLengthInRange(1, 36),
				Description:  "Description of the security group.",
			},
			"rule": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				MaxItems:   15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Id of the rule.",
						},
						"action": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 1),
							Description:  "Name of the security group rule to be queried.",
						},
						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 256),
							Description:  "Description of the security group rule.",
						},
						"targetaddress": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"targetport": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"localport": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"priority": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 12),
							Description:  "Description of the security group rule.",
						},
						"direction": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 16),
							Description:  "Description of the security group rule.",
						},
						"protocol": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 16),
							Description:  "Description of the security group rule.",
						},
						"ruletype": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 36),
							Description:  "Description of the security group rule.",
						},
					},
				},
			},
			"rule_current": {
				Type:       schema.TypeSet,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				MaxItems:   15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Id of the rule.",
						},
						"action": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 1),
							Description:  "Name of the security group rule to be queried.",
						},
						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 256),
							Description:  "Description of the security group rule.",
						},
						"targetaddress": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"targetport": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"localport": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 200),
							Description:  "Description of the security group rule.",
						},
						"priority": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 12),
							Description:  "Description of the security group rule.",
						},
						"direction": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 16),
							Description:  "Description of the security group rule.",
						},
						"protocol": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 16),
							Description:  "Description of the security group rule.",
						},
						"ruletype": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: u.ValidateStringLengthInRange(1, 36),
							Description:  "Description of the security group rule.",
						},
					},
				},
				Set: resourceSecurityRuleHash,
			},
		},
	}
}

func resourceCdsSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_security_group.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}

	groupName := ""
	if name, ok := d.GetOk("name"); ok {
		gName := name.(string)
		if len(gName) > 0 {
			groupName = gName
		}
	}
	log.Println("create security group: ", groupName)
	description := ""
	if desc, ok := d.GetOk("description"); ok {
		gdescription := desc.(string)
		if len(gdescription) > 0 {
			description = gdescription
		}
	}
	groupType := ""
	if securityGroupType, ok := d.GetOk("type"); ok {
		gsecurityGroupType := securityGroupType.(string)
		if len(gsecurityGroupType) > 0 {
			groupType = gsecurityGroupType
		}
	}

	var rules []map[string]interface{}
	if rulea, ok := d.GetOk("rule"); ok {
		ruleArray := rulea.(*schema.Set).List()
		for _, value := range ruleArray {
			i := value.(map[string]interface{})
			rules = append(rules, i)
		}

	}

	id, errRet := securityGroupService.CreateSecurityGroup(ctx, groupName, description, groupType, rules)
	if errRet != nil {
		return errRet
	}
	d.SetId(id)
	return resourceCdsSecurityGroupRead(d, meta)
}

func resourceCdsSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_security_group.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	log.Printf("GroupRead")

	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}

	id := d.Id()
	readRequest := security_group.NewDescribeSecurityGroupRequest()
	readRequest.SecurityGroupId = common.StringPtr(id)

	readResponse, errRet := securityGroupService.DescribeSecurityGroup(ctx, readRequest)
	if errRet != nil {
		return errRet
	}

	currentGroups := readResponse.Data.SecurityGroup
	if len(currentGroups) == 0 {
		d.SetId("")
		return nil
	}

	currentGroup := readResponse.Data.SecurityGroup[0]
	d.Set("description", *currentGroup.Description)
	d.Set("name", *currentGroup.SecurityGroupName)
	d.Set("type", *currentGroup.SecurityGroupType)

	readRuleRequest := security_group.NewDescribeSecurityGroupRuleRequest()
	readRuleRequest.SecurityGroupId = common.StringPtr(id)
	readRuleResponse, errRet := securityGroupService.DescribeSecurityGroupRule(ctx, readRuleRequest)
	if errRet != nil {
		return errRet
	}
	rules := &schema.Set{
		F: resourceSecurityRuleHash,
	}
	for _, value := range readRuleResponse.Data.SecurityGroupRules {
		rule := map[string]interface{}{}
		rule["id"] = *value.RuleId
		rule["action"] = strconv.Itoa(*value.Action)
		rule["description"] = *value.Description
		rule["targetaddress"] = *value.TargetAddress
		rule["targetport"] = *value.TargetPort
		rule["localport"] = *value.LocalPort
		rule["priority"] = strconv.Itoa(*value.Priority)
		rule["direction"] = *value.Direction
		rule["protocol"] = *value.Protocol
		rule["ruletype"] = *value.Type
		rules.Add(rule)
	}
	err := d.Set("rule_current", rules)
	if err != nil {
		return err
	}

	return nil
}

func resourceCdsSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("update security group")
	defer logElapsed("resource.cds_security_group.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}

	id := d.Id()
	d.Partial(true)

	var (
		name        string
		description string
		updateAttr  []string
	)

	old, now := d.GetChange("name")
	if d.HasChange("name") {
		updateAttr = append(updateAttr, "name")

		name = now.(string)
	} else {
		name = old.(string)
	}

	old, now = d.GetChange("description")
	if d.HasChange("description") {
		updateAttr = append(updateAttr, "description")

		description = now.(string)
	} else {
		description = old.(string)
	}

	modifyRequest := security_group.NewModifySecurityGroupRequest()
	modifyRequest.SecurityGroupName = common.StringPtr(name)
	modifyRequest.SecurityGroupId = common.StringPtr(id)
	modifyRequest.Description = common.StringPtr(description)
	_, errRet := securityGroupService.ModifySecurityGroup(ctx, modifyRequest)
	if errRet != nil {
		return errRet
	}

	for _, attr := range updateAttr {
		d.SetPartial(attr)
	}

	err := resourceCdsSecurityGroupUpdateRules(d, meta, id, ctx)
	if err != nil {
		return err
	}

	d.Partial(false)
	return resourceCdsSecurityGroupRead(d, meta)
}

func resourceCdsSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("delete security group")
	defer logElapsed("resource.cds_security_group.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}
	request := security_group.NewDeleteSecurityGroupRequest()
	request.SecurityGroupId = common.StringPtr(d.Id())
	_, errRet := securityGroupService.DeleteSecurityGroup(ctx, request)
	if errRet != nil {
		return errRet
	}
	time.Sleep(10 * time.Second)
	return nil
}

func resourceSecurityRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	//if v, ok := m["action"]; ok {
	//	buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	//}

	if v, ok := m["description"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["targetaddress"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["targetport"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["localport"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["direction"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["ruletype"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	if v, ok := m["protocol"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}

func resourceCdsSecurityGroupUpdateRules(
	d *schema.ResourceData, meta interface{}, id string, ctx context.Context) error {

	if d.HasChange("rule") {
		c := d.Get("rule_current")
		o, n := d.GetChange("rule")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}
		if c == nil {
			c = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)
		cs := c.(*schema.Set)

		// TODO: Currently there's no rule updating logic
		// If any rule were updated, we take it as new rule.
		// since it's complicated to judge between updating and adding.
		// Additionly, not all rule options could be modified by API.
		add := ns.Difference(os).List()
		remove := getRemoveRules(cs.List(), ns.List())

		securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}
		// Remove old rules
		if len(remove) > 0 {
			deleteRuleRequest := security_group_rule.NewDeleteSecurityGroupRuleRequest()
			deleteRuleRequest.SecurityGroupId = common.StringPtr(id)
			for _, rule := range remove {
				//rule := value.(map[string]interface{})
				log.Println("remove rule: ", rule)
				deleteRuleRequest.RuleIds = append(deleteRuleRequest.RuleIds, common.StringPtr(rule["id"].(string)))

			}
			securityGroupService.DeleteSecurityGroupRule(ctx, deleteRuleRequest)
		}

		// Add new rules
		if len(add) > 0 {
			for _, value := range add {
				log.Println("add rule: ", value)
				rule := value.(map[string]interface{})
				securityRule := security_group_rule.NewAddSecurityGroupRuleRequest()
				errRet := u.Mapstructure(rule, securityRule)
				securityRule.SecurityGroupId = common.StringPtr(id)
				if errRet != nil {
					return errRet
				}
				_, errRet = securityGroupService.client.UseSecurityRuleClient().CreateSecurityGroupRule(securityRule)
				time.Sleep(2 * time.Second)
			}
		}

		d.SetPartial("rule")
	}

	return nil
}

func getRemoveRules(oldWithID, newWithoutID []interface{}) []map[string]interface{} {
	var newStrRules []string
	var removeRules []map[string]interface{}

	for _, rule := range newWithoutID {
		r := rule.(map[string]interface{})
		newStrRules = append(newStrRules, sortStrRule(r))
	}

	for _, rule := range oldWithID {
		r := rule.(map[string]interface{})
		sortR := sortStrRule(r)
		if !containsString(newStrRules, sortR) {
			removeRules = append(removeRules, r)
		}
	}

	return removeRules
}

func sortStrRule(rule map[string]interface{}) string {
	var keys []string

	copyRule := make(map[string]interface{})
	for k, v := range rule {
		copyRule[k] = v
	}
	delete(copyRule, "id")

	for k := range copyRule {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var s string
	var l []string
	for _, k := range keys {
		s = fmt.Sprintf("%s=\"%s\"", k, copyRule[k])
		l = append(l, s)
	}
	strRule := strings.Join(l, ",")

	return strRule
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
