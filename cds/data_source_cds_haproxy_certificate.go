package cds

import (
	"context"
	"fmt"

	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHaproxyCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyCertificateRead,

		Schema: map[string]*schema.Schema{
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
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

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), response.Data); err != nil {
			return err
		}
	}

	return nil
}
