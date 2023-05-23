package dashboards

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_dashboard", func(r *config.Resource) {
		r.Kind = "Dashboard"
		r.ShortGroup = "dashboard"
	})

	p.AddResourceConfigurator("datadog_dashboard_json", func(r *config.Resource) {
		r.Kind = "DashboardJSON"
		r.ShortGroup = "dashboard"
	})
}
