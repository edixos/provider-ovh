// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/cloud/projectworkflowbackup"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/providerconfig"
	projectnetworkprivate "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivate"
	projectnetworkprivatesubnet "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivatesubnet"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectworkflowbackup.Setup,
		providerconfig.Setup,
		projectnetworkprivate.Setup,
		projectnetworkprivatesubnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
