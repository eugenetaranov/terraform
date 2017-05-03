package local

import (
	"testing"

	"github.com/eugenetaranov/terraform/helper/schema"
	"github.com/eugenetaranov/terraform/terraform"
)

var testProviders = map[string]terraform.ResourceProvider{
	"local": Provider(),
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
