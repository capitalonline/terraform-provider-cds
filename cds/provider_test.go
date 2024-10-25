package cds

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var defaultRegionToTest = os.Getenv("CDS_REGION")

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cds": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err:%s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CDS_SECRET_ID"); v == "" {
		t.Fatal("CDS_SECRET_ID must be set for acceptance tests")
	}
	if v := os.Getenv("CDS_SECRET_KEY"); v == "" {
		t.Fatal("CDS_SECRET_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("CDS_REGION"); v == "" {
		t.Fatal("CDS_REGION= must be set for acceptance tests")
	}
}
