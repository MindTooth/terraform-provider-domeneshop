package domeneshop

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/MindTooth/go-domeneshop"
)

func resourceDomeneshopRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomeneshopRecordCreate,
		Read:   resourceDomeneshopRecordRead,
		Update: resourceDomeneshopRecordUpdate,
		Delete: resourceDomeneshopRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDomeneshopRecordImport,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"A",
					"AAAA",
					"CNAME",
					"MX",
					"SRV",
					"TXT",
				}, false),
			},
			"domain_id": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
				De
			},
			"data": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(60, 604800),
			},
		},
	}
}

func resourceDomeneshopRecordCreate(d *schema.ResourceData, meta interface{}) error {


	client := m.(*Client).getClient()
 	ctx := m.(*Client).getContext()

 	domain_id := int32(d.Get("domain_id").(int))
 	name := d.Get("host").(string)

 	record, resp, err := client.DnsApi.DomainsDomainIdDnsPost(ctx, domain_id,

	d.SetId(fmt.Sprint(record_id))

	return nil
}

func resourceDomeneshopRecordRead(d *schema.ResourceData, meta interface{}) error   {}
func resourceDomeneshopRecordUpdate(d *schema.ResourceData, meta interface{}) error {}
func resourceDomeneshopRecordDelete(d *schema.ResourceData, meta interface{}) error {}
func resourceDomeneshopRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
}
