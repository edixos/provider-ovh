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
	projectdatabase "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabase"
	projectdatabasedatabase "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasedatabase"
	projectdatabaseintegration "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseintegration"
	projectdatabaseiprestriction "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseiprestriction"
	projectdatabasekafkaacl "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasekafkaacl"
	projectdatabasekafkaschemaregistryacl "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasekafkaschemaregistryacl"
	projectdatabasekafkatopic "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasekafkatopic"
	projectdatabasem3dbnamespace "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasem3dbnamespace"
	projectdatabasem3dbuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasem3dbuser"
	projectdatabasemongodbuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasemongodbuser"
	projectdatabaseopensearchpattern "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseopensearchpattern"
	projectdatabaseopensearchuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseopensearchuser"
	projectdatabasepostgresqluser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasepostgresqluser"
	projectdatabaseredisuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseredisuser"
	projectdatabaseuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseuser"
	installationtemplate "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplate"
	installationtemplatepartitionscheme "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionscheme"
	installationtemplatepartitionschemehardwareraid "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionschemehardwareraid"
	installationtemplatepartitionschemepartition "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/installationtemplatepartitionschemepartition"
	ipxescript "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/ipxescript"
	serverinstalltask "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverinstalltask"
	servernetworking "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverupdate"
	zone "github.com/edixos/provider-ovh/internal/controller/dns/zone"
	zonerecord "github.com/edixos/provider-ovh/internal/controller/dns/zonerecord"
	zoneredirection "github.com/edixos/provider-ovh/internal/controller/dns/zoneredirection"
	projectkube "github.com/edixos/provider-ovh/internal/controller/kube/projectkube"
	projectkubeiprestrictions "github.com/edixos/provider-ovh/internal/controller/kube/projectkubeiprestrictions"
	projectkubenodepool "github.com/edixos/provider-ovh/internal/controller/kube/projectkubenodepool"
	projectkubeoidc "github.com/edixos/provider-ovh/internal/controller/kube/projectkubeoidc"
	httpfarm "github.com/edixos/provider-ovh/internal/controller/lb/httpfarm"
	httpfarmserver "github.com/edixos/provider-ovh/internal/controller/lb/httpfarmserver"
	httpfrontend "github.com/edixos/provider-ovh/internal/controller/lb/httpfrontend"
	httproute "github.com/edixos/provider-ovh/internal/controller/lb/httproute"
	httprouterule "github.com/edixos/provider-ovh/internal/controller/lb/httprouterule"
	iploadbalancing "github.com/edixos/provider-ovh/internal/controller/lb/iploadbalancing"
	refresh "github.com/edixos/provider-ovh/internal/controller/lb/refresh"
	tcpfarm "github.com/edixos/provider-ovh/internal/controller/lb/tcpfarm"
	tcpfarmserver "github.com/edixos/provider-ovh/internal/controller/lb/tcpfarmserver"
	tcpfrontend "github.com/edixos/provider-ovh/internal/controller/lb/tcpfrontend"
	tcproute "github.com/edixos/provider-ovh/internal/controller/lb/tcproute"
	tcprouterule "github.com/edixos/provider-ovh/internal/controller/lb/tcprouterule"
	vracknetwork "github.com/edixos/provider-ovh/internal/controller/lb/vracknetwork"
	logscluster "github.com/edixos/provider-ovh/internal/controller/logs/logscluster"
	logsinput "github.com/edixos/provider-ovh/internal/controller/logs/logsinput"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/providerconfig"
	projectnetworkprivate "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivate"
	projectnetworkprivatesubnet "github.com/edixos/provider-ovh/internal/controller/publiccloudnetwork/projectnetworkprivatesubnet"
	projectcontainerregistry "github.com/edixos/provider-ovh/internal/controller/registry/projectcontainerregistry"
	projectcontainerregistryoidc "github.com/edixos/provider-ovh/internal/controller/registry/projectcontainerregistryoidc"
	projectcontainerregistryuser "github.com/edixos/provider-ovh/internal/controller/registry/projectcontainerregistryuser"
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
		projectdatabase.Setup,
		projectdatabasedatabase.Setup,
		projectdatabaseintegration.Setup,
		projectdatabaseiprestriction.Setup,
		projectdatabasekafkaacl.Setup,
		projectdatabasekafkaschemaregistryacl.Setup,
		projectdatabasekafkatopic.Setup,
		projectdatabasem3dbnamespace.Setup,
		projectdatabasem3dbuser.Setup,
		projectdatabasemongodbuser.Setup,
		projectdatabaseopensearchpattern.Setup,
		projectdatabaseopensearchuser.Setup,
		projectdatabasepostgresqluser.Setup,
		projectdatabaseredisuser.Setup,
		projectdatabaseuser.Setup,
		installationtemplate.Setup,
		installationtemplatepartitionscheme.Setup,
		installationtemplatepartitionschemehardwareraid.Setup,
		installationtemplatepartitionschemepartition.Setup,
		ipxescript.Setup,
		serverinstalltask.Setup,
		servernetworking.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
		zone.Setup,
		zonerecord.Setup,
		zoneredirection.Setup,
		projectkube.Setup,
		projectkubeiprestrictions.Setup,
		projectkubenodepool.Setup,
		projectkubeoidc.Setup,
		httpfarm.Setup,
		httpfarmserver.Setup,
		httpfrontend.Setup,
		httproute.Setup,
		httprouterule.Setup,
		iploadbalancing.Setup,
		refresh.Setup,
		tcpfarm.Setup,
		tcpfarmserver.Setup,
		tcpfrontend.Setup,
		tcproute.Setup,
		tcprouterule.Setup,
		vracknetwork.Setup,
		logscluster.Setup,
		logsinput.Setup,
		providerconfig.Setup,
		projectnetworkprivate.Setup,
		projectnetworkprivatesubnet.Setup,
		projectcontainerregistry.Setup,
		projectcontainerregistryoidc.Setup,
		projectcontainerregistryuser.Setup,
		projectworkflowbackup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
