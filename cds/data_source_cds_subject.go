package cds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/platform"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsSubjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubjectsRead,
		Schema: map[string]*schema.Schema{
			"subjects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subject_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Subject id.",
						},
						"name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Subject name.",
						},
						"balance": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Balance. Available balance, unit yuan/US dollar.",
						},
						"begin_time": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Begin time.",
						},
						"end_time": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "End time.",
						},
						"bill_method": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Bill method.",
						},
						"bill_method_display": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Bill method display.",
						},
						"goods_ids": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Goods ids.",
						},
						"goods_names": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Goods names.",
						},
						"site_ids": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Site ids. Supported created sites, with no limitations on this field representing support for all sites.",
						},
						"site_names": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Site names. Supported created sites, with no limitations on this field representing support for all sites.",
						},
					},
				},
				Description: "Subjects.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Output file path.",
			},
		},
		Description: "Data source subjects \n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data "cds_data_source_subjects" "subjects" {
}
` +
			"\n```",
	}
}

func dataSourceSubjectsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.subject.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	platformService := PlatformService{client: meta.(*CdsClient).apiConn}
	request := platform.NewDescribeSubjectsRequest()
	response, err := platformService.DescribeSubjects(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return errors.New(*response.Msg)
	}
	var subjects = make([]map[string]interface{}, 0, len(response.Data.SubjectList))
	for i := 0; i < len(response.Data.SubjectList); i++ {
		item := response.Data.SubjectList[i]
		subject := make(map[string]interface{})
		subject["subject_id"] = item.Id
		subject["name"] = item.Name
		subject["balance"] = item.Balance
		subject["begin_time"] = item.BeginTime
		subject["end_time"] = item.EndTime
		subject["bill_method"] = item.BillMethod
		subject["bill_method_display"] = item.BillMethodDisplay
		subject["goods_ids"] = item.GoodsIds
		subject["goods_names"] = item.GoodsNames
		subject["site_ids"] = item.SiteIds
		subject["site_names"] = item.SiteNames
		subjects = append(subjects, subject)
	}

	if len(subjects) > 0 {
		err = d.Set("subjects", subjects)
		bytes, _ := json.Marshal(subjects)
		logElapsed(fmt.Sprintf("set subjects:%s", string(bytes)))()
		if err != nil {
			return err
		}
	}
	if path, ok := d.GetOk("result_output_file"); ok {
		if err = writeToFile(path.(string), subjects); err != nil {
			return err
		}
	}
	if d.Id() == "" {
		id := fmt.Sprintf("cds_datasource_subjects")
		d.SetId(id)
	}
	return nil
}
