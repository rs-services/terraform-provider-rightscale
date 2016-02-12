package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"

	//	"gopkg.in/rightscale/rsc.v5/cm15"
	//	"gopkg.in/rightscale/rsc.v5/rsapi"
)

func resourceRightScaleSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceRightScaleSSHKeyCreate,
		Read:   resourceRightScaleSSHKeyRead,
		Delete: resourceRightScaleSSHKeyDelete,

		Schema: map[string]*schema.Schema{
			/*
				"href": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
					ForceNew: true,
				},
			*/

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRightScaleSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceRightScaleSSHKeyRead(d, meta)
}

func resourceRightScaleSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
