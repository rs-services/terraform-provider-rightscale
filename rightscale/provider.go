package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_ACCOUNT_ID", nil),
			},

			"refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_REFRESH_TOKEN", nil),
			},

			"api_host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_API_HOST", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rightscale_ssh_key":    resourceRightScaleSSHKey(),
			"rightscale_deployment": resourceRightScaleDeployment(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccountID:    d.Get("account_id").(int),
		RefreshToken: d.Get("refresh_token").(string),
		APIHost:      d.Get("api_host").(string),
	}

	return config.Client()
}
