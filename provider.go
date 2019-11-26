package sdm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	apiV1 "github.com/strongdm/strongdm-sdk-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "An authenticated JWT token used to access the StrongDM API.",
			},

			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "api.strongdm.com:443",
				Description: "The host and port of the StrongDM API endpoint.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sdm_node": resourceNode(),
			"sdm_role": resourceRole(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			client, err := apiV1.New(d.Get("host").(string), d.Get("api_key").(string))
			if err != nil {
				return nil, fmt.Errorf("cannot dial API server: %w", err)
			}
			return client, nil
		},
	}
}
