// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	identitygroup "github.com/edixos/provider-ovh/internal/controller/accountmanagement/identitygroup"
	identityuser "github.com/edixos/provider-ovh/internal/controller/accountmanagement/identityuser"
	policy "github.com/edixos/provider-ovh/internal/controller/accountmanagement/policy"
	project "github.com/edixos/provider-ovh/internal/controller/accountmanagement/project"
	projectuser "github.com/edixos/provider-ovh/internal/controller/accountmanagement/projectuser"
	projectusers3credential "github.com/edixos/provider-ovh/internal/controller/accountmanagement/projectusers3credential"
	projectusers3policy "github.com/edixos/provider-ovh/internal/controller/accountmanagement/projectusers3policy"
	sshkey "github.com/edixos/provider-ovh/internal/controller/accountmanagement/sshkey"
	projectfailoveripattach "github.com/edixos/provider-ovh/internal/controller/additionalip/projectfailoveripattach"
	reverse "github.com/edixos/provider-ovh/internal/controller/additionalip/reverse"
	service "github.com/edixos/provider-ovh/internal/controller/additionalip/service"
	cephacl "github.com/edixos/provider-ovh/internal/controller/clouddiskarray/cephacl"
	installationtemplate "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplate"
	installationtemplatepartitionscheme "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionscheme"
	installationtemplatepartitionschemehardwareraid "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionschemehardwareraid"
	installationtemplatepartitionschemepartition "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionschemepartition"
	ipxescript "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/ipxescript"
	serverinstalltask "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverinstalltask"
	servernetworking "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverupdate"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/providerconfig"
	projectnetworkprivate "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivate"
	projectnetworkprivatesubnet "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivatesubnet"
	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/vminstances/projectworkflowbackup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identitygroup.Setup,
		identityuser.Setup,
		policy.Setup,
		project.Setup,
		projectuser.Setup,
		projectusers3credential.Setup,
		projectusers3policy.Setup,
		sshkey.Setup,
		projectfailoveripattach.Setup,
		reverse.Setup,
		service.Setup,
		cephacl.Setup,
		installationtemplate.Setup,
		installationtemplatepartitionscheme.Setup,
		installationtemplatepartitionschemehardwareraid.Setup,
		installationtemplatepartitionschemepartition.Setup,
		ipxescript.Setup,
		serverinstalltask.Setup,
		servernetworking.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
		providerconfig.Setup,
		projectnetworkprivate.Setup,
		projectnetworkprivatesubnet.Setup,
		projectworkflowbackup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
