package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"

	"gopkg.in/rightscale/rsc.v5/cm15"
	"gopkg.in/rightscale/rsc.v5/rsapi"

	"log"
)

func resourceRightScaleSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceRightScaleSSHKeyCreate,
		Read:   resourceRightScaleSSHKeyRead,
		Delete: resourceRightScaleSSHKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of ssh key to be created",
			},
			"cloud_href": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Href of cloud to create ssh key in",
			},
		},
	}
}

func resourceRightScaleSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)
	locator := client.SshKeyLocator(d.Get("cloud_href").(string))

	sshKey, err := locator.Create(&cm15.SshKeyParam{
		Name: d.Get("name").(string),
	})

	if err != nil {
		log.Printf("[RIGHTSCALE] SSH KEY CREATE ERROR: %s", err.Error())
	}

	d.SetId(string(sshKey.Href))

	return resourceRightScaleSSHKeyRead(d, meta)
}

func resourceRightScaleSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)
	sshKey, err := client.SshKeyLocator(d.Id()).Show(rsapi.APIParams{})

	if err != nil {
		log.Printf("[RIGHTSCALE] SSH KEY READ ERROR: %s", err.Error())
	}

	d.Set("name", sshKey.ResourceUid)

	return nil
}

func resourceRightScaleSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)
	locator := client.SshKeyLocator(d.Id())

	err := locator.Destroy()

	if err != nil {
		log.Printf("[RIGHTSCALE] SSH KEY DELETE ERROR: %s", err.Error())
	}

	return nil
}
