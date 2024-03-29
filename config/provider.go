/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/ok-amba/provider-datadog/config/authnmapping"
	"github.com/ok-amba/provider-datadog/config/dashboards"
	"github.com/ok-amba/provider-datadog/config/role"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "datadog"
	modulePath     = "github.com/ok-amba/provider-datadog"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dashboards.Configure,
		authnmapping.Configure,
		role.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
