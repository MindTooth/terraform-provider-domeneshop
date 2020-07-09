package domeneshop

import (
	"fmt"
	"log"

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
			"jens": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The id of the domain",
			},
		},
	}
}

func dataSourceDomeneshopDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).getClient()
	ctx := m.(*Client).getContext()

	name := d.Get("name").(string)

	domain, resp, err := client.DomainsApi.GetDomains(*ctx, &domeneshop.GetDomainsOpts{Domain: optional.NewString(name)})

	log.Println(resp)
	fmt.Println(resp)

	log.Println(domain)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("domain not found: %s", err)
		}
		return fmt.Errorf("error getting dns domains: %v", err)
	}

	d.SetId(domain[0].Domain)
	d.Set("jens", domain[0].Id)

	return nil
}
