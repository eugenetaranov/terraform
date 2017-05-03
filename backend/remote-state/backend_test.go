package remotestate

import (
	"testing"

	"github.com/eugenetaranov/terraform/backend"
)

func TestBackend_impl(t *testing.T) {
	var _ backend.Backend = new(Backend)
}
