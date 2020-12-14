package command

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/packer/packer"
	plugingetter "github.com/hashicorp/packer/packer/plugin-getter"
)

// CoreWrapper wraps a packer.Core in order to have it's Initialize func return
// a diagnostic.
type CoreWrapper struct {
	*packer.Core
}

func (c *CoreWrapper) Initialize() hcl.Diagnostics {
	err := c.Core.Initialize()
	if err != nil {
		return hcl.Diagnostics{
			&hcl.Diagnostic{
				Detail:   err.Error(),
				Severity: hcl.DiagError,
			},
		}
	}
	return nil
}

func (c *CoreWrapper) PluginRequirements() (plugingetter.List, hcl.Diagnostics) {
	return nil, hcl.Diagnostics{
		&hcl.Diagnostic{
			Summary:  "Packer init is not supported in Packer - JSON",
			Detail:   "Please manually install plugins or use a HCL2 configuration that will do that for you.",
			Severity: hcl.DiagError,
		},
	}
}
