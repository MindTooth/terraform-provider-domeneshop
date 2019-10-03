package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/MindTooth/terraform-provider-domeneshop/domeneshop"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: domeneshop.Provider})
}
