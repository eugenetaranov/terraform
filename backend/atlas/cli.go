package atlas

import (
	"github.com/eugenetaranov/terraform/backend"
)

// backend.CLI impl.
func (b *Backend) CLIInit(opts *backend.CLIOpts) error {
	b.CLI = opts.CLI
	b.CLIColor = opts.CLIColor
	b.ContextOpts = opts.ContextOpts
	return nil
}
