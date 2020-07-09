package main

import (
	"github.com/MindTooth/terraform-provider-domeneshop/domeneshop"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return domeneshop.Provider()
		},
	})
}
