package domeneshop

import (
	"fmt"

	"github.com/MindTooth/go-domeneshop"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceDomeneshopRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomeneshopRecordRead,
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The domain id to search with",
				ValidateFunc: validation.NoZeroValues,
			},
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Domain active or not",
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed
			"data": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain active or not",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Domain active or not",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain active or not",
			},
		},
	}
}

func dataSourceDomeneshopRecordRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).getClient()
	ctx := m.(*Client).getContext()

	domain_id := int32(d.Get("domain_id").(int))
	name := d.Get("host").(string)

	record, resp, err := client.DnsApi.GetDnsRecords(*ctx, domain_id, &domeneshop.GetDnsRecordsOpts{Host: optional.NewString(name)})

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("domain not found: %s", err)
		}
		return fmt.Errorf("error getting dns domains: %v", err)
	}

	d.SetId(fmt.Sprint(record[0].Id))
	d.Set("host", record[0].Host)
	d.Set("data", record[0].Data)
	d.Set("ttl", record[0].Ttl)
	d.Set("ttl", record[0].Type)

	return nil
}
