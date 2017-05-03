package aws

import (
	"github.com/eugenetaranov/terraform/helper/schema"
)

func resourceAwsOpsworksStaticWebLayer() *schema.Resource {
	layerType := &opsworksLayerType{
		TypeName:         "web",
		DefaultLayerName: "Static Web Server",

		Attributes: map[string]*opsworksLayerTypeAttribute{},
	}

	return layerType.SchemaResource()
}
