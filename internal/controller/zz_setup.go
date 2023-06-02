/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	authnmapping "github.com/ok-amba/provider-datadog/internal/controller/authnmapping/authnmapping"
	dashboard "github.com/ok-amba/provider-datadog/internal/controller/dashboard/dashboard"
	dashboardjson "github.com/ok-amba/provider-datadog/internal/controller/dashboard/dashboardjson"
	providerconfig "github.com/ok-amba/provider-datadog/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authnmapping.Setup,
		dashboard.Setup,
		dashboardjson.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
