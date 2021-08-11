package cds

import (
	"context"
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testCertificateResource = `
data cds_data_source_certificate my_certificate {
    certificate_name = "my_certificate"
	certificate = "XXXXXXXX"
	private_key = "XXXXXXXX"
}
`

func TestCertificateSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCertificateResource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_certificate.my_certificate", "certificate_name", "my_certificate"),
				),
			},
		},
	})
}

func testAccCertificateDestroy(s *terraform.State) error {
	defer logElapsed("resource.certificate.read")

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_haproxy" {
			continue
		}

		haproxyService := HaproxyService{client: testAccProvider.Meta().(*CdsClient).apiConn}

		request := haproxy.NewDescribeCACertificateRequest()
		request.CertificateId = common.StringPtr(rs.Primary.Attributes["certificate_id"])

		response, err := haproxyService.DescribeCACertificate(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			return fmt.Errorf("read certificate failed, error: %s", *response.Message)
		}
	}

	return nil
}
