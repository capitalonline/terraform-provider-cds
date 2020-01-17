package main

import (
	"terraform-provider-cds/cds"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cds.Provider})
}
