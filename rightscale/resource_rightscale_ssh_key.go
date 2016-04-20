package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"sync"

	"gopkg.in/rightscale/rsc.v5/cm15"
	"gopkg.in/rightscale/rsc.v5/rsapi"

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
	mutex.Lock()

	client := meta.(*cm15.API)

	sshKeyLocator := client.SshKeyLocator
	sshKey, err := sshKeyLocator.Create(&cm15.SshKeyParam{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})

	if err != nil {
		log.Printf("[RIGHTSCALE] SSH KEY CREATE ERROR: %s", err.Error())
	}

	// Set this resource id to RightScale HREF
	d.SetId(string(sshKey.Href))

	mutex.Unlock()

	return resourceRightScaleSSHKeyRead(d, meta)
}

func resourceRightScaleSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
