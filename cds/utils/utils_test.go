package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/security_group_rule"
)

func TestUtil_MapToStruct(t *testing.T) {
	v := &security_group_rule.AddSecurityGroupRuleRequest{}
	m := make(map[string]interface{})
	m["securitygroupid"] = "a"
	m["action"] = "a"
	m["description"] = "a"
	m["targetaddress"] = "a"
	m["targetport"] = "a"
	m["localport"] = "a"
	m["direction"] = "a"
	m["priority"] = "a"
	m["protocol"] = "a"
	m["ruletype"] = "a"

	t1 := time.Now() // get current time
	Mapstructure(m, v)
	elapsed := time.Since(t1)
	fmt.Println(v.ToJsonString())
	fmt.Println("App elapsed: ", elapsed)
}

func TestUtil_Merge(t *testing.T) {
	m1 := map[string]interface{}{}
	m2 := map[string]interface{}{}
	m3 := map[string]interface{}{}

	m1["i1"] = "1"
	m1["i2"] = 3
	m1["i3"] = "3"

	m2["i2"] = "1"
	m3["i2"] = 2

	res := merge(m1, m2, m3)
	fmt.Println(res)
}

func merge(ms ...map[string]interface{}) map[string][]interface{} {
	res := map[string][]interface{}{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if v == v2 {
					continue srcMap
				}
			}
			res[k] = append(res[k], v)
		}
	}
	return res
}
