package domeneshop

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceDomeneshopForward() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomeneshopForwardRead,
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The domain ID to search with",
				ValidateFunc: validation.NoZeroValues,
			},
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The (sub)hostname to search with, www or @ for root",
				ValidateFunc: validation.NoZeroValues,
			},
			// Computer
			"frame": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether forwarding use iframe",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The returned URL which domain is forwarded to",
			},
		},
	}
}

func dataSourceDomeneshopForwardRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).getClient()
	ctx := m.(*Client).getContext()

	domain_id := int32(d.Get("domain_id").(int))
	name := d.Get("host").(string)

	forward, resp, err := client.ForwardsApi.DomainsDomainIdForwardsHostGet(*ctx, domain_id, name)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("domain not found: %s", err)
		}
		return fmt.Errorf("error getting forwarding: %v", err)
	}

	d.Set("host", forward.Host)
	d.Set("url", forward.Url)
	d.Set("frame", forward.Frame)

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
