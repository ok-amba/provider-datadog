package role

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_role", func(r *config.Resource) {
		r.Kind = "Role"
		r.ShortGroup = ""
	})
}
