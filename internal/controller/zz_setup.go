/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	firewall "github.com/edixos/provider-ovh/internal/controller/additionalip/firewall"
	firewallrule "github.com/edixos/provider-ovh/internal/controller/additionalip/firewallrule"
	mitigation "github.com/edixos/provider-ovh/internal/controller/additionalip/mitigation"
	move "github.com/edixos/provider-ovh/internal/controller/additionalip/move"
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
	projectdatabasepostgresqlconnectionpool "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasepostgresqlconnectionpool"
	projectdatabasepostgresqluser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabasepostgresqluser"
	projectdatabaseredisuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseredisuser"
	projectdatabaseuser "github.com/edixos/provider-ovh/internal/controller/databases/projectdatabaseuser"
	servernetworking "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/dedicatedserver/serverupdate"
	zone "github.com/edixos/provider-ovh/internal/controller/dns/zone"
	zonednssec "github.com/edixos/provider-ovh/internal/controller/dns/zonednssec"
	zonerecord "github.com/edixos/provider-ovh/internal/controller/dns/zonerecord"
	zoneredirection "github.com/edixos/provider-ovh/internal/controller/dns/zoneredirection"
	projectgateway "github.com/edixos/provider-ovh/internal/controller/gateway/projectgateway"
	iampermissionsgroup "github.com/edixos/provider-ovh/internal/controller/iam/iampermissionsgroup"
	iampolicy "github.com/edixos/provider-ovh/internal/controller/iam/iampolicy"
	iamresourcegroup "github.com/edixos/provider-ovh/internal/controller/iam/iamresourcegroup"
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
	projectregionloadbalancerlogsubscription "github.com/edixos/provider-ovh/internal/controller/lb/projectregionloadbalancerlogsubscription"
	refresh "github.com/edixos/provider-ovh/internal/controller/lb/refresh"
	ssl "github.com/edixos/provider-ovh/internal/controller/lb/ssl"
	tcpfarm "github.com/edixos/provider-ovh/internal/controller/lb/tcpfarm"
	tcpfarmserver "github.com/edixos/provider-ovh/internal/controller/lb/tcpfarmserver"
	tcpfrontend "github.com/edixos/provider-ovh/internal/controller/lb/tcpfrontend"
	tcproute "github.com/edixos/provider-ovh/internal/controller/lb/tcproute"
	tcprouterule "github.com/edixos/provider-ovh/internal/controller/lb/tcprouterule"
	udpfarm "github.com/edixos/provider-ovh/internal/controller/lb/udpfarm"
	udpfarmserver "github.com/edixos/provider-ovh/internal/controller/lb/udpfarmserver"
	udpfrontend "github.com/edixos/provider-ovh/internal/controller/lb/udpfrontend"
	vracknetwork "github.com/edixos/provider-ovh/internal/controller/lb/vracknetwork"
	logscluster "github.com/edixos/provider-ovh/internal/controller/logs/logscluster"
	logsinput "github.com/edixos/provider-ovh/internal/controller/logs/logsinput"
	logsoutputopensearchalias "github.com/edixos/provider-ovh/internal/controller/logs/logsoutputopensearchalias"
	logsoutputopensearchindex "github.com/edixos/provider-ovh/internal/controller/logs/logsoutputopensearchindex"
	logstoken "github.com/edixos/provider-ovh/internal/controller/logs/logstoken"
	group "github.com/edixos/provider-ovh/internal/controller/me/group"
	oauth2client "github.com/edixos/provider-ovh/internal/controller/me/oauth2client"
	userme "github.com/edixos/provider-ovh/internal/controller/me/user"
	nashapartition "github.com/edixos/provider-ovh/internal/controller/nas/nashapartition"
	nashapartitionaccess "github.com/edixos/provider-ovh/internal/controller/nas/nashapartitionaccess"
	nashapartitionsnapshot "github.com/edixos/provider-ovh/internal/controller/nas/nashapartitionsnapshot"
	privatenetwork "github.com/edixos/provider-ovh/internal/controller/network/privatenetwork"
	subnet "github.com/edixos/provider-ovh/internal/controller/network/subnet"
	subnetv2 "github.com/edixos/provider-ovh/internal/controller/network/subnetv2"
	privatedatabase "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabase"
	privatedatabasedatabase "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabasedatabase"
	privatedatabaseuser "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabaseuser"
	privatedatabaseusergrant "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/edixos/provider-ovh/internal/controller/privatesql/privatedatabasewhitelist"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/providerconfig"
	containerregistry "github.com/edixos/provider-ovh/internal/controller/registry/containerregistry"
	containerregistryiprestrictionsmanagement "github.com/edixos/provider-ovh/internal/controller/registry/containerregistryiprestrictionsmanagement"
	containerregistryiprestrictionsregistry "github.com/edixos/provider-ovh/internal/controller/registry/containerregistryiprestrictionsregistry"
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
		firewall.Setup,
		firewallrule.Setup,
		mitigation.Setup,
		move.Setup,
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
		projectdatabasepostgresqlconnectionpool.Setup,
		projectdatabasepostgresqluser.Setup,
		projectdatabaseredisuser.Setup,
		projectdatabaseuser.Setup,
		servernetworking.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
		zone.Setup,
		zonednssec.Setup,
		zonerecord.Setup,
		zoneredirection.Setup,
		projectgateway.Setup,
		iampermissionsgroup.Setup,
		iampolicy.Setup,
		iamresourcegroup.Setup,
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
		projectregionloadbalancerlogsubscription.Setup,
		refresh.Setup,
		ssl.Setup,
		tcpfarm.Setup,
		tcpfarmserver.Setup,
		tcpfrontend.Setup,
		tcproute.Setup,
		tcprouterule.Setup,
		udpfarm.Setup,
		udpfarmserver.Setup,
		udpfrontend.Setup,
		vracknetwork.Setup,
		logscluster.Setup,
		logsinput.Setup,
		logsoutputopensearchalias.Setup,
		logsoutputopensearchindex.Setup,
		logstoken.Setup,
		group.Setup,
		oauth2client.Setup,
		userme.Setup,
		nashapartition.Setup,
		nashapartitionaccess.Setup,
		nashapartitionsnapshot.Setup,
		privatenetwork.Setup,
		subnet.Setup,
		subnetv2.Setup,
		privatedatabase.Setup,
		privatedatabasedatabase.Setup,
		privatedatabaseuser.Setup,
		privatedatabaseusergrant.Setup,
		privatedatabasewhitelist.Setup,
		providerconfig.Setup,
		containerregistry.Setup,
		containerregistryiprestrictionsmanagement.Setup,
		containerregistryiprestrictionsregistry.Setup,
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
