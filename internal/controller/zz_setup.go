// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	projectfailoveripattach "github.com/edixos/provider-ovh/internal/controller/additionalip/projectfailoveripattach"
	reverse "github.com/edixos/provider-ovh/internal/controller/additionalip/reverse"
	service "github.com/edixos/provider-ovh/internal/controller/additionalip/service"
	project "github.com/edixos/provider-ovh/internal/controller/cloud/project"
	s3credentials "github.com/edixos/provider-ovh/internal/controller/cloud/s3credentials"
	s3policy "github.com/edixos/provider-ovh/internal/controller/cloud/s3policy"
	user "github.com/edixos/provider-ovh/internal/controller/cloud/user"
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
	iampolicy "github.com/edixos/provider-ovh/internal/controller/iam/iampolicy"
	cluster "github.com/edixos/provider-ovh/internal/controller/kube/cluster"
	iprestriction "github.com/edixos/provider-ovh/internal/controller/kube/iprestriction"
	nodepool "github.com/edixos/provider-ovh/internal/controller/kube/nodepool"
	oidcconfiguration "github.com/edixos/provider-ovh/internal/controller/kube/oidcconfiguration"
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
	group "github.com/edixos/provider-ovh/internal/controller/me/group"
	sshkey "github.com/edixos/provider-ovh/internal/controller/me/sshkey"
	userme "github.com/edixos/provider-ovh/internal/controller/me/user"
	nashapartition "github.com/edixos/provider-ovh/internal/controller/nas/nashapartition"
	nashapartitionaccess "github.com/edixos/provider-ovh/internal/controller/nas/nashapartitionaccess"
	nashapartitionsnapshot "github.com/edixos/provider-ovh/internal/controller/nas/nashapartitionsnapshot"
	privatenetwork "github.com/edixos/provider-ovh/internal/controller/network/privatenetwork"
	subnet "github.com/edixos/provider-ovh/internal/controller/network/subnet"
	privatedatabase "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabase"
	privatedatabasedatabase "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabasedatabase"
	privatedatabaseuser "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabaseuser"
	privatedatabaseusergrant "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabasewhitelist"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/providerconfig"
	containerregistry "github.com/edixos/provider-ovh/internal/controller/registry/containerregistry"
	containerregistryoidc "github.com/edixos/provider-ovh/internal/controller/registry/containerregistryoidc"
	containerregistryuser "github.com/edixos/provider-ovh/internal/controller/registry/containerregistryuser"
	projectregionstoragepresign "github.com/edixos/provider-ovh/internal/controller/storage/projectregionstoragepresign"
	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/vminstances/projectworkflowbackup"
	cloudproject "github.com/edixos/provider-ovh/internal/controller/vrack/cloudproject"
	dedicatedserver "github.com/edixos/provider-ovh/internal/controller/vrack/dedicatedserver"
	dedicatedserverinterface "github.com/edixos/provider-ovh/internal/controller/vrack/dedicatedserverinterface"
	ip "github.com/edixos/provider-ovh/internal/controller/vrack/ip"
	iploadbalancingvrack "github.com/edixos/provider-ovh/internal/controller/vrack/iploadbalancing"
	vrack "github.com/edixos/provider-ovh/internal/controller/vrack/vrack"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectfailoveripattach.Setup,
		reverse.Setup,
		service.Setup,
		project.Setup,
		s3credentials.Setup,
		s3policy.Setup,
		user.Setup,
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
		iampolicy.Setup,
		cluster.Setup,
		iprestriction.Setup,
		nodepool.Setup,
		oidcconfiguration.Setup,
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
		group.Setup,
		sshkey.Setup,
		userme.Setup,
		nashapartition.Setup,
		nashapartitionaccess.Setup,
		nashapartitionsnapshot.Setup,
		privatenetwork.Setup,
		subnet.Setup,
		privatedatabase.Setup,
		privatedatabasedatabase.Setup,
		privatedatabaseuser.Setup,
		privatedatabaseusergrant.Setup,
		privatedatabasewhitelist.Setup,
		providerconfig.Setup,
		containerregistry.Setup,
		containerregistryoidc.Setup,
		containerregistryuser.Setup,
		projectregionstoragepresign.Setup,
		projectworkflowbackup.Setup,
		cloudproject.Setup,
		dedicatedserver.Setup,
		dedicatedserverinterface.Setup,
		ip.Setup,
		iploadbalancingvrack.Setup,
		vrack.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
