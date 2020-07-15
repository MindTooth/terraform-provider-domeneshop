package domeneshop

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOMENESHOP_TOKEN", nil),
				Description: "Token for accessing service.",
			},
			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOMENESHOP_SECRET", nil),
				Description: "Secret to access service.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"domeneshop_domain": dataSourceDomeneshopDomain(),
			"domeneshop_record": dataSourceDomeneshopRecord(),
		},
		/*
			ResourcesMap: map[string]*schema.Resource{
				"domeneshop_record": resourceDomeneshopRecord(),
			},
		*/
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		Token:            d.Get("token").(string),
		Secret:           d.Get("secret").(string),
		TerraformVersion: terraformVersion,
	}

	return config.Client()
}
