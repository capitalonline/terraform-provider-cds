package cds

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHaproxyCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyCertificateRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"ha_cert_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Ha certificate list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"brand": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Brand. Certificate Brand",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate id.",
						},
						"certificate_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate name.",
						},
						"certificate_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate type.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start time.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End time.",
						},
						"organization": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Organization.",
						},
						"valid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Valid.",
						},
					},
				},
			},
		},
		Description: "Data source haproxy certificate. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#4describeloadbalancers)",
	}
}

func dataSourceHaproxyCertificateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.certificate.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}
	response, err := haproxyService.DescribeCACertificates(ctx, haproxy.NewDescribeCACertificatesRequest())
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("haproxy certificate read failed, error: %s", *response.Message)
	}
	haCertList := make([]map[string]interface{}, 0, len(response.Data))

	for _, v := range response.Data {
		mapping := map[string]interface{}{
			"brand":            *v.Brand,
			"certificate_id":   *v.CertificateId,
			"certificate_name": *v.CertificateName,
			"certificate_type": *v.CertificateType,
			"domain":           *v.Domain,
			"start_time":       *v.StartTime,
			"end_time":         *v.EndTime,
			"organization":     *v.Organization,
			"valid":            *v.Valid,
		}
		haCertList = append(haCertList, mapping)
	}

	d.SetId(strconv.Itoa(int(time.Now().Unix())))
	err = d.Set("ha_cert_list", haCertList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set configuration list fail, reason:%s\n ", logId, err.Error())
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), response.Data); err != nil {
			return err
		}
	}

	return nil
}
