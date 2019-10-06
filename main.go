package main

import (
	"github.com/MindTooth/terraform-provider-domeneshop/domeneshop"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: domeneshop.Provider})
}
