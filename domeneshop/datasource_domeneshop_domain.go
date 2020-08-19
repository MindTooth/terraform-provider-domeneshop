package domeneshop

import (
	"fmt"

	"github.com/MindTooth/go-domeneshop"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceDomeneshopDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomeneshopDomainRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The domain to search id for",
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed
			"domain_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The id of the domain",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain active or not",
			},
		},
	}
}

func dataSourceDomeneshopDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).getClient()
	ctx := m.(*Client).getContext()

	name := d.Get("name").(string)

	domain, resp, err := client.DomainsApi.GetDomains(*ctx, &domeneshop.GetDomainsOpts{Domain: optional.NewString(name)})

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("domain not found: %s", err)
		}
		return fmt.Errorf("error getting dns domains: %v", err)
	}

	d.SetId(fmt.Sprint(domain[0].Id))
	d.Set("domain_id", domain[0].Id)
	d.Set("status", domain[0].Status)

	return nil
}
