package domeneshop

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
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
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:  d.Get("token").(string),
		Secret: d.Get("secret").(string),
	}

	return config.Client()
}
