/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	firewall "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/firewall"
	firewallrule "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/firewallrule"
	mitigation "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/mitigation"
	move "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/move"
	projectfailoveripattach "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/projectfailoveripattach"
	reverse "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/reverse"
	service "github.com/edixos/provider-ovh/internal/controller/namespaced/additionalip/service"
	project "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/project"
	s3credentials "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/s3credentials"
	s3policy "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/s3policy"
	user "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/user"
	cephacl "github.com/edixos/provider-ovh/internal/controller/namespaced/clouddiskarray/cephacl"
	projectdatabase "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabase"
	projectdatabasedatabase "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasedatabase"
	projectdatabaseintegration "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseintegration"
	projectdatabaseiprestriction "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseiprestriction"
	projectdatabasekafkaacl "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasekafkaacl"
	projectdatabasekafkaschemaregistryacl "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasekafkaschemaregistryacl"
	projectdatabasekafkatopic "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasekafkatopic"
	projectdatabasem3dbnamespace "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasem3dbnamespace"
	projectdatabasem3dbuser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasem3dbuser"
	projectdatabasemongodbuser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasemongodbuser"
	projectdatabaseopensearchpattern "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseopensearchpattern"
	projectdatabaseopensearchuser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseopensearchuser"
	projectdatabasepostgresqlconnectionpool "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasepostgresqlconnectionpool"
	projectdatabasepostgresqluser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabasepostgresqluser"
	projectdatabaseredisuser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseredisuser"
	projectdatabaseuser "github.com/edixos/provider-ovh/internal/controller/namespaced/databases/projectdatabaseuser"
	servernetworking "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/serverupdate"
	zone "github.com/edixos/provider-ovh/internal/controller/namespaced/dns/zone"
	zonednssec "github.com/edixos/provider-ovh/internal/controller/namespaced/dns/zonednssec"
	zonerecord "github.com/edixos/provider-ovh/internal/controller/namespaced/dns/zonerecord"
	zoneredirection "github.com/edixos/provider-ovh/internal/controller/namespaced/dns/zoneredirection"
	projectgateway "github.com/edixos/provider-ovh/internal/controller/namespaced/gateway/projectgateway"
	iampermissionsgroup "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iampermissionsgroup"
	iampolicy "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iampolicy"
	iamresourcegroup "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iamresourcegroup"
	cluster "github.com/edixos/provider-ovh/internal/controller/namespaced/kube/cluster"
	iprestriction "github.com/edixos/provider-ovh/internal/controller/namespaced/kube/iprestriction"
	nodepool "github.com/edixos/provider-ovh/internal/controller/namespaced/kube/nodepool"
	oidcconfiguration "github.com/edixos/provider-ovh/internal/controller/namespaced/kube/oidcconfiguration"
	httpfarm "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfarm"
	httpfarmserver "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfarmserver"
	httpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfrontend"
	httproute "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httproute"
	httprouterule "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httprouterule"
	iploadbalancing "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/iploadbalancing"
	projectregionloadbalancerlogsubscription "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/projectregionloadbalancerlogsubscription"
	refresh "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/refresh"
	ssl "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/ssl"
	tcpfarm "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfarm"
	tcpfarmserver "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfarmserver"
	tcpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfrontend"
	tcproute "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcproute"
	tcprouterule "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcprouterule"
	udpfarm "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/udpfarm"
	udpfarmserver "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/udpfarmserver"
	udpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/udpfrontend"
	vracknetwork "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/vracknetwork"
	logscluster "github.com/edixos/provider-ovh/internal/controller/namespaced/logs/logscluster"
	logsinput "github.com/edixos/provider-ovh/internal/controller/namespaced/logs/logsinput"
	logsoutputopensearchalias "github.com/edixos/provider-ovh/internal/controller/namespaced/logs/logsoutputopensearchalias"
	logsoutputopensearchindex "github.com/edixos/provider-ovh/internal/controller/namespaced/logs/logsoutputopensearchindex"
	logstoken "github.com/edixos/provider-ovh/internal/controller/namespaced/logs/logstoken"
	group "github.com/edixos/provider-ovh/internal/controller/namespaced/me/group"
	oauth2client "github.com/edixos/provider-ovh/internal/controller/namespaced/me/oauth2client"
	userme "github.com/edixos/provider-ovh/internal/controller/namespaced/me/user"
	nashapartition "github.com/edixos/provider-ovh/internal/controller/namespaced/nas/nashapartition"
	nashapartitionaccess "github.com/edixos/provider-ovh/internal/controller/namespaced/nas/nashapartitionaccess"
	nashapartitionsnapshot "github.com/edixos/provider-ovh/internal/controller/namespaced/nas/nashapartitionsnapshot"
	privatenetwork "github.com/edixos/provider-ovh/internal/controller/namespaced/network/privatenetwork"
	subnet "github.com/edixos/provider-ovh/internal/controller/namespaced/network/subnet"
	subnetv2 "github.com/edixos/provider-ovh/internal/controller/namespaced/network/subnetv2"
	privatedatabase "github.com/edixos/provider-ovh/internal/controller/namespaced/privatesql/privatedatabase"
	privatedatabasedatabase "github.com/edixos/provider-ovh/internal/controller/namespaced/privatesql/privatedatabasedatabase"
	privatedatabaseuser "github.com/edixos/provider-ovh/internal/controller/namespaced/privatesql/privatedatabaseuser"
	privatedatabaseusergrant "github.com/edixos/provider-ovh/internal/controller/namespaced/privatesql/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/edixos/provider-ovh/internal/controller/namespaced/privatesql/privatedatabasewhitelist"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/namespaced/providerconfig"
	containerregistry "github.com/edixos/provider-ovh/internal/controller/namespaced/registry/containerregistry"
	containerregistryiprestrictionsmanagement "github.com/edixos/provider-ovh/internal/controller/namespaced/registry/containerregistryiprestrictionsmanagement"
	containerregistryiprestrictionsregistry "github.com/edixos/provider-ovh/internal/controller/namespaced/registry/containerregistryiprestrictionsregistry"
	containerregistryoidc "github.com/edixos/provider-ovh/internal/controller/namespaced/registry/containerregistryoidc"
	containerregistryuser "github.com/edixos/provider-ovh/internal/controller/namespaced/registry/containerregistryuser"
	projectregionstoragepresign "github.com/edixos/provider-ovh/internal/controller/namespaced/storage/projectregionstoragepresign"
	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/namespaced/vminstances/projectworkflowbackup"
	cloudproject "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/cloudproject"
	dedicatedserver "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/dedicatedserver"
	dedicatedserverinterface "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/dedicatedserverinterface"
	ip "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/ip"
	iploadbalancingvrack "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/iploadbalancing"
	vrack "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/vrack"
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.SetupGated,
		firewallrule.SetupGated,
		mitigation.SetupGated,
		move.SetupGated,
		projectfailoveripattach.SetupGated,
		reverse.SetupGated,
		service.SetupGated,
		project.SetupGated,
		s3credentials.SetupGated,
		s3policy.SetupGated,
		user.SetupGated,
		cephacl.SetupGated,
		projectdatabase.SetupGated,
		projectdatabasedatabase.SetupGated,
		projectdatabaseintegration.SetupGated,
		projectdatabaseiprestriction.SetupGated,
		projectdatabasekafkaacl.SetupGated,
		projectdatabasekafkaschemaregistryacl.SetupGated,
		projectdatabasekafkatopic.SetupGated,
		projectdatabasem3dbnamespace.SetupGated,
		projectdatabasem3dbuser.SetupGated,
		projectdatabasemongodbuser.SetupGated,
		projectdatabaseopensearchpattern.SetupGated,
		projectdatabaseopensearchuser.SetupGated,
		projectdatabasepostgresqlconnectionpool.SetupGated,
		projectdatabasepostgresqluser.SetupGated,
		projectdatabaseredisuser.SetupGated,
		projectdatabaseuser.SetupGated,
		servernetworking.SetupGated,
		serverreboottask.SetupGated,
		serverupdate.SetupGated,
		zone.SetupGated,
		zonednssec.SetupGated,
		zonerecord.SetupGated,
		zoneredirection.SetupGated,
		projectgateway.SetupGated,
		iampermissionsgroup.SetupGated,
		iampolicy.SetupGated,
		iamresourcegroup.SetupGated,
		cluster.SetupGated,
		iprestriction.SetupGated,
		nodepool.SetupGated,
		oidcconfiguration.SetupGated,
		httpfarm.SetupGated,
		httpfarmserver.SetupGated,
		httpfrontend.SetupGated,
		httproute.SetupGated,
		httprouterule.SetupGated,
		iploadbalancing.SetupGated,
		projectregionloadbalancerlogsubscription.SetupGated,
		refresh.SetupGated,
		ssl.SetupGated,
		tcpfarm.SetupGated,
		tcpfarmserver.SetupGated,
		tcpfrontend.SetupGated,
		tcproute.SetupGated,
		tcprouterule.SetupGated,
		udpfarm.SetupGated,
		udpfarmserver.SetupGated,
		udpfrontend.SetupGated,
		vracknetwork.SetupGated,
		logscluster.SetupGated,
		logsinput.SetupGated,
		logsoutputopensearchalias.SetupGated,
		logsoutputopensearchindex.SetupGated,
		logstoken.SetupGated,
		group.SetupGated,
		oauth2client.SetupGated,
		userme.SetupGated,
		nashapartition.SetupGated,
		nashapartitionaccess.SetupGated,
		nashapartitionsnapshot.SetupGated,
		privatenetwork.SetupGated,
		subnet.SetupGated,
		subnetv2.SetupGated,
		privatedatabase.SetupGated,
		privatedatabasedatabase.SetupGated,
		privatedatabaseuser.SetupGated,
		privatedatabaseusergrant.SetupGated,
		privatedatabasewhitelist.SetupGated,
		providerconfig.SetupGated,
		containerregistry.SetupGated,
		containerregistryiprestrictionsmanagement.SetupGated,
		containerregistryiprestrictionsregistry.SetupGated,
		containerregistryoidc.SetupGated,
		containerregistryuser.SetupGated,
		projectregionstoragepresign.SetupGated,
		projectworkflowbackup.SetupGated,
		cloudproject.SetupGated,
		dedicatedserver.SetupGated,
		dedicatedserverinterface.SetupGated,
		ip.SetupGated,
		iploadbalancingvrack.SetupGated,
		vrack.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
