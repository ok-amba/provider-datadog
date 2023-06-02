package authnmapping

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_authn_mapping", func(r *config.Resource) {
		r.Kind = "AuthnMapping"
		r.ShortGroup = "authnmapping"
	})
}
