package schema

import (
	"testing"

	"github.com/eugenetaranov/terraform/config"
	"github.com/eugenetaranov/terraform/terraform"
)

// TestResourceDataRaw creates a ResourceData from a raw configuration map.
func TestResourceDataRaw(
	t *testing.T, schema map[string]*Schema, raw map[string]interface{}) *ResourceData {
	c, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	sm := schemaMap(schema)
	diff, err := sm.Diff(nil, terraform.NewResourceConfig(c))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	result, err := sm.Data(nil, diff)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return result
}
