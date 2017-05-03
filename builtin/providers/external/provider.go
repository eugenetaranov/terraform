package external

import (
	"github.com/eugenetaranov/terraform/helper/schema"
	"github.com/eugenetaranov/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"external": dataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{},
	}
}
