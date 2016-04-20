package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"sync"

	"gopkg.in/rightscale/rsc.v5/cm15"
	"gopkg.in/rightscale/rsc.v5/rsapi"
)

var mutex = &sync.Mutex{}

func resourceRightScaleDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceRightScaleDeploymentCreate,
		Read:   resourceRightScaleDeploymentRead,
		Update: resourceRightScaleDeploymentUpdate,
		Delete: resourceRightScaleDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Deployment name",
			},

			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "Deployment description",
			},
		},
	}
}

func resourceRightScaleDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	mutex.Lock()
	client := meta.(*cm15.API)

	deploymentLocator := client.DeploymentLocator("/api/deployments")
	deployment, err := deploymentLocator.Create(&cm15.DeploymentParam{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})

	if err != nil {
		log.Printf("[RIGHTSCALE] DEPLOYMENT CREATE ERROR: %s", err.Error())
	}

	// Set this resource id to RightScale HREF
	d.SetId(string(deployment.Href))

	mutex.Unlock()
	return resourceRightScaleDeploymentRead(d, meta)
}

func resourceRightScaleDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	client := meta.(*cm15.API)
	deployment, err := client.DeploymentLocator(d.Id()).Show(rsapi.APIParams{})

	d.Set("name", deployment.Name)
	d.Set("description", deployment.Description)

	if err != nil {
		log.Printf("[RIGHTSCALE] DEPLOYMENT READ ERROR %s", err.Error())
	}

	return nil
}

func resourceRightScaleDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	if d.HasChange("description") {
		if err := client.DeploymentLocator(d.Id()).Update(&cm15.DeploymentParam{Description: d.Get("description").(string)}); err != nil {
			return err
		}
	}

	if d.HasChange("name") {
		if err := client.DeploymentLocator(d.Id()).Update(&cm15.DeploymentParam{Name: d.Get("name").(string)}); err != nil {
			return err
		}
	}

	return resourceRightScaleDeploymentRead(d, meta)
}

func resourceRightScaleDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
