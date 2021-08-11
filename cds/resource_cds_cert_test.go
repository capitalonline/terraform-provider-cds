package cds

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_certificate.my_certificate", "certificate_name", "my_certificate"),
				),
			},
		},
	})
}

const testAccCertificateConfig = `
resource cds_certificate my_haproxy {
	certificate_name = "my_certificate"
	certificate = "XXXXXX"
	private_key = "XXXXXX"
}
`
