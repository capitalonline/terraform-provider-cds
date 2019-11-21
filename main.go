package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-cds/cds"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cds.Provider})
}
