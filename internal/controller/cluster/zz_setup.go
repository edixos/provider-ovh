/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	firewall "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/firewall"
	firewallrule "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/firewallrule"
	mitigation "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/mitigation"
	move "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/move"
	projectfailoveripattach "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/projectfailoveripattach"
	reverse "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/reverse"
	service "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/service"
	floatingip "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/floatingip"
	instance "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/instance"
	instancegroup "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/instancegroup"
	plan "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/plan"
	project "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/project"
	projectcontainerregistryiam "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectcontainerregistryiam"
	projectgatewayinterface "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectgatewayinterface"
	projectinstance "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectinstance"
	projectinstancesnapshot "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectinstancesnapshot"
	projectrancher "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectrancher"
	projectregion "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectregion"
	projectsshkey "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectsshkey"
	projectstorage "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectstorage"
	projectvolume "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectvolume"
	projectvolumebackup "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/projectvolumebackup"
	quota "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/quota"
	s3credentials "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/s3credentials"
	s3policy "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/s3policy"
	securitygroup "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/securitygroup"
	sshkey "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/sshkey"
	user "github.com/edixos/provider-ovh/internal/controller/cluster/cloud/user"
	cephacl "github.com/edixos/provider-ovh/internal/controller/cluster/clouddiskarray/cephacl"
	projectdatabase "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabase"
	projectdatabaseclickhouseuser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseclickhouseuser"
	projectdatabasedatabase "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasedatabase"
	projectdatabaseintegration "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseintegration"
	projectdatabasekafkaacl "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasekafkaacl"
	projectdatabasekafkaschemaregistryacl "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasekafkaschemaregistryacl"
	projectdatabasekafkatopic "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasekafkatopic"
	projectdatabaselogsubscription "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaselogsubscription"
	projectdatabasemongodbprometheus "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasemongodbprometheus"
	projectdatabasemongodbuser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasemongodbuser"
	projectdatabaseopensearchpattern "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseopensearchpattern"
	projectdatabaseopensearchuser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseopensearchuser"
	projectdatabasepostgresqlconnectionpool "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasepostgresqlconnectionpool"
	projectdatabasepostgresqluser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasepostgresqluser"
	projectdatabaseprometheus "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseprometheus"
	projectdatabaseuser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabaseuser"
	projectdatabasevalkeyuser "github.com/edixos/provider-ovh/internal/controller/cluster/databases/projectdatabasevalkeyuser"
	server "github.com/edixos/provider-ovh/internal/controller/cluster/dedicated/server"
	serverreinstalltask "github.com/edixos/provider-ovh/internal/controller/cluster/dedicated/serverreinstalltask"
	servernetworking "github.com/edixos/provider-ovh/internal/controller/cluster/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/cluster/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/cluster/dedicatedserver/serverupdate"
	dsrecords "github.com/edixos/provider-ovh/internal/controller/cluster/dns/dsrecords"
	name "github.com/edixos/provider-ovh/internal/controller/cluster/dns/name"
	nameservers "github.com/edixos/provider-ovh/internal/controller/cluster/dns/nameservers"
	zone "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zone"
	zonednssec "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonednssec"
	zonedynhostlogin "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonedynhostlogin"
	zonedynhostrecord "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonedynhostrecord"
	zonerecord "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonerecord"
	zoneredirection "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zoneredirection"
	zoneimport "github.com/edixos/provider-ovh/internal/controller/cluster/domain/zoneimport"
	domainaccount "github.com/edixos/provider-ovh/internal/controller/cluster/email/domainaccount"
	cloudgateway "github.com/edixos/provider-ovh/internal/controller/cluster/gateway/cloudgateway"
	projectgateway "github.com/edixos/provider-ovh/internal/controller/cluster/gateway/projectgateway"
	iampermissionsgroup "github.com/edixos/provider-ovh/internal/controller/cluster/iam/iampermissionsgroup"
	iampolicy "github.com/edixos/provider-ovh/internal/controller/cluster/iam/iampolicy"
	iamresourcegroup "github.com/edixos/provider-ovh/internal/controller/cluster/iam/iamresourcegroup"
	iamresourcetags "github.com/edixos/provider-ovh/internal/controller/cluster/iam/iamresourcetags"
	credential "github.com/edixos/provider-ovh/internal/controller/cluster/kms/credential"
	keymanagercontainer "github.com/edixos/provider-ovh/internal/controller/cluster/kms/keymanagercontainer"
	keymanagercontainerconsumer "github.com/edixos/provider-ovh/internal/controller/cluster/kms/keymanagercontainerconsumer"
	keymanagersecret "github.com/edixos/provider-ovh/internal/controller/cluster/kms/keymanagersecret"
	keymanagersecretconsumer "github.com/edixos/provider-ovh/internal/controller/cluster/kms/keymanagersecretconsumer"
	okms "github.com/edixos/provider-ovh/internal/controller/cluster/kms/okms"
	secret "github.com/edixos/provider-ovh/internal/controller/cluster/kms/secret"
	servicekey "github.com/edixos/provider-ovh/internal/controller/cluster/kms/servicekey"
	cluster "github.com/edixos/provider-ovh/internal/controller/cluster/kube/cluster"
	iprestriction "github.com/edixos/provider-ovh/internal/controller/cluster/kube/iprestriction"
	logsubscription "github.com/edixos/provider-ovh/internal/controller/cluster/kube/logsubscription"
	nodepool "github.com/edixos/provider-ovh/internal/controller/cluster/kube/nodepool"
	oidcconfiguration "github.com/edixos/provider-ovh/internal/controller/cluster/kube/oidcconfiguration"
	httpfarm "github.com/edixos/provider-ovh/internal/controller/cluster/lb/httpfarm"
	httpfarmserver "github.com/edixos/provider-ovh/internal/controller/cluster/lb/httpfarmserver"
	httpfrontend "github.com/edixos/provider-ovh/internal/controller/cluster/lb/httpfrontend"
	httproute "github.com/edixos/provider-ovh/internal/controller/cluster/lb/httproute"
	httprouterule "github.com/edixos/provider-ovh/internal/controller/cluster/lb/httprouterule"
	iploadbalancing "github.com/edixos/provider-ovh/internal/controller/cluster/lb/iploadbalancing"
	projectloadbalancer "github.com/edixos/provider-ovh/internal/controller/cluster/lb/projectloadbalancer"
	projectregionloadbalancerlogsubscription "github.com/edixos/provider-ovh/internal/controller/cluster/lb/projectregionloadbalancerlogsubscription"
	refresh "github.com/edixos/provider-ovh/internal/controller/cluster/lb/refresh"
	ssl "github.com/edixos/provider-ovh/internal/controller/cluster/lb/ssl"
	tcpfarm "github.com/edixos/provider-ovh/internal/controller/cluster/lb/tcpfarm"
	tcpfarmserver "github.com/edixos/provider-ovh/internal/controller/cluster/lb/tcpfarmserver"
	tcpfrontend "github.com/edixos/provider-ovh/internal/controller/cluster/lb/tcpfrontend"
	tcproute "github.com/edixos/provider-ovh/internal/controller/cluster/lb/tcproute"
	tcprouterule "github.com/edixos/provider-ovh/internal/controller/cluster/lb/tcprouterule"
	udpfarm "github.com/edixos/provider-ovh/internal/controller/cluster/lb/udpfarm"
	udpfarmserver "github.com/edixos/provider-ovh/internal/controller/cluster/lb/udpfarmserver"
	udpfrontend "github.com/edixos/provider-ovh/internal/controller/cluster/lb/udpfrontend"
	vracknetwork "github.com/edixos/provider-ovh/internal/controller/cluster/lb/vracknetwork"
	logscluster "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logscluster"
	logsencryptionkey "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsencryptionkey"
	logsinput "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsinput"
	logsoutputgraylogstream "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsoutputgraylogstream"
	logsoutputopensearchalias "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsoutputopensearchalias"
	logsoutputopensearchindex "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsoutputopensearchindex"
	logsrole "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsrole"
	logsrolepermissionstream "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsrolepermissionstream"
	logstoken "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logstoken"
	group "github.com/edixos/provider-ovh/internal/controller/cluster/me/group"
	identityusertoken "github.com/edixos/provider-ovh/internal/controller/cluster/me/identityusertoken"
	oauth2client "github.com/edixos/provider-ovh/internal/controller/cluster/me/oauth2client"
	userme "github.com/edixos/provider-ovh/internal/controller/cluster/me/user"
	nashapartition "github.com/edixos/provider-ovh/internal/controller/cluster/nas/nashapartition"
	nashapartitionaccess "github.com/edixos/provider-ovh/internal/controller/cluster/nas/nashapartitionaccess"
	nashapartitionsnapshot "github.com/edixos/provider-ovh/internal/controller/cluster/nas/nashapartitionsnapshot"
	privatenetwork "github.com/edixos/provider-ovh/internal/controller/cluster/network/privatenetwork"
	privatevracknetwork "github.com/edixos/provider-ovh/internal/controller/cluster/network/privatevracknetwork"
	privatevracksubnet "github.com/edixos/provider-ovh/internal/controller/cluster/network/privatevracksubnet"
	projectregionnetwork "github.com/edixos/provider-ovh/internal/controller/cluster/network/projectregionnetwork"
	subnet "github.com/edixos/provider-ovh/internal/controller/cluster/network/subnet"
	subnetv2 "github.com/edixos/provider-ovh/internal/controller/cluster/network/subnetv2"
	servicekeyjwk "github.com/edixos/provider-ovh/internal/controller/cluster/okms/servicekeyjwk"
	vps "github.com/edixos/provider-ovh/internal/controller/cluster/ovh/vps"
	privatedatabase "github.com/edixos/provider-ovh/internal/controller/cluster/privatesql/privatedatabase"
	privatedatabasedatabase "github.com/edixos/provider-ovh/internal/controller/cluster/privatesql/privatedatabasedatabase"
	privatedatabaseuser "github.com/edixos/provider-ovh/internal/controller/cluster/privatesql/privatedatabaseuser"
	privatedatabaseusergrant "github.com/edixos/provider-ovh/internal/controller/cluster/privatesql/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/edixos/provider-ovh/internal/controller/cluster/privatesql/privatedatabasewhitelist"
	providerconfig "github.com/edixos/provider-ovh/internal/controller/cluster/providerconfig"
	containerregistry "github.com/edixos/provider-ovh/internal/controller/cluster/registry/containerregistry"
	containerregistryiprestrictionsmanagement "github.com/edixos/provider-ovh/internal/controller/cluster/registry/containerregistryiprestrictionsmanagement"
	containerregistryiprestrictionsregistry "github.com/edixos/provider-ovh/internal/controller/cluster/registry/containerregistryiprestrictionsregistry"
	containerregistryoidc "github.com/edixos/provider-ovh/internal/controller/cluster/registry/containerregistryoidc"
	containerregistryuser "github.com/edixos/provider-ovh/internal/controller/cluster/registry/containerregistryuser"
	blockvolume "github.com/edixos/provider-ovh/internal/controller/cluster/storage/blockvolume"
	blockvolumebackup "github.com/edixos/provider-ovh/internal/controller/cluster/storage/blockvolumebackup"
	blockvolumesnapshot "github.com/edixos/provider-ovh/internal/controller/cluster/storage/blockvolumesnapshot"
	efs "github.com/edixos/provider-ovh/internal/controller/cluster/storage/efs"
	efsshare "github.com/edixos/provider-ovh/internal/controller/cluster/storage/efsshare"
	efsshareacl "github.com/edixos/provider-ovh/internal/controller/cluster/storage/efsshareacl"
	efssharesnapshot "github.com/edixos/provider-ovh/internal/controller/cluster/storage/efssharesnapshot"
	fileshare "github.com/edixos/provider-ovh/internal/controller/cluster/storage/fileshare"
	fileshareacl "github.com/edixos/provider-ovh/internal/controller/cluster/storage/fileshareacl"
	filesharenetwork "github.com/edixos/provider-ovh/internal/controller/cluster/storage/filesharenetwork"
	filesharesnapshot "github.com/edixos/provider-ovh/internal/controller/cluster/storage/filesharesnapshot"
	projectfilestorageshare "github.com/edixos/provider-ovh/internal/controller/cluster/storage/projectfilestorageshare"
	projectfilestoragesharenetwork "github.com/edixos/provider-ovh/internal/controller/cluster/storage/projectfilestoragesharenetwork"
	projectregionstoragepresign "github.com/edixos/provider-ovh/internal/controller/cluster/storage/projectregionstoragepresign"
	projectstoragelifecycleconfiguration "github.com/edixos/provider-ovh/internal/controller/cluster/storage/projectstoragelifecycleconfiguration"
	projectstoragereplicationjob "github.com/edixos/provider-ovh/internal/controller/cluster/storage/projectstoragereplicationjob"
	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/cluster/vminstances/projectworkflowbackup"
	cloudproject "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/cloudproject"
	connectpopconfig "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/connectpopconfig"
	connectpopdatacenterconfig "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/connectpopdatacenterconfig"
	connectpopdatacenterextraconfig "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/connectpopdatacenterextraconfig"
	dedicatedcloud "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/dedicatedcloud"
	dedicatedclouddatacenter "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/dedicatedclouddatacenter"
	dedicatedserver "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/dedicatedserver"
	dedicatedserverinterface "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/dedicatedserverinterface"
	ip "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/ip"
	iploadbalancingvrack "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/iploadbalancing"
	ipv6 "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/ipv6"
	ipv6routedsubrange "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/ipv6routedsubrange"
	ovhcloudconnect "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/ovhcloudconnect"
	publicroutingpriority "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/publicroutingpriority"
	vrack "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/vrack"
	vrackservices "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/vrackservices"
	vrackservicesorder "github.com/edixos/provider-ovh/internal/controller/cluster/vrack/vrackservicesorder"
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
		floatingip.Setup,
		instance.Setup,
		instancegroup.Setup,
		plan.Setup,
		project.Setup,
		projectcontainerregistryiam.Setup,
		projectgatewayinterface.Setup,
		projectinstance.Setup,
		projectinstancesnapshot.Setup,
		projectrancher.Setup,
		projectregion.Setup,
		projectsshkey.Setup,
		projectstorage.Setup,
		projectvolume.Setup,
		projectvolumebackup.Setup,
		quota.Setup,
		s3credentials.Setup,
		s3policy.Setup,
		securitygroup.Setup,
		sshkey.Setup,
		user.Setup,
		cephacl.Setup,
		projectdatabase.Setup,
		projectdatabaseclickhouseuser.Setup,
		projectdatabasedatabase.Setup,
		projectdatabaseintegration.Setup,
		projectdatabasekafkaacl.Setup,
		projectdatabasekafkaschemaregistryacl.Setup,
		projectdatabasekafkatopic.Setup,
		projectdatabaselogsubscription.Setup,
		projectdatabasemongodbprometheus.Setup,
		projectdatabasemongodbuser.Setup,
		projectdatabaseopensearchpattern.Setup,
		projectdatabaseopensearchuser.Setup,
		projectdatabasepostgresqlconnectionpool.Setup,
		projectdatabasepostgresqluser.Setup,
		projectdatabaseprometheus.Setup,
		projectdatabaseuser.Setup,
		projectdatabasevalkeyuser.Setup,
		server.Setup,
		serverreinstalltask.Setup,
		servernetworking.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
		dsrecords.Setup,
		name.Setup,
		nameservers.Setup,
		zone.Setup,
		zonednssec.Setup,
		zonedynhostlogin.Setup,
		zonedynhostrecord.Setup,
		zonerecord.Setup,
		zoneredirection.Setup,
		zoneimport.Setup,
		domainaccount.Setup,
		cloudgateway.Setup,
		projectgateway.Setup,
		iampermissionsgroup.Setup,
		iampolicy.Setup,
		iamresourcegroup.Setup,
		iamresourcetags.Setup,
		credential.Setup,
		keymanagercontainer.Setup,
		keymanagercontainerconsumer.Setup,
		keymanagersecret.Setup,
		keymanagersecretconsumer.Setup,
		okms.Setup,
		secret.Setup,
		servicekey.Setup,
		cluster.Setup,
		iprestriction.Setup,
		logsubscription.Setup,
		nodepool.Setup,
		oidcconfiguration.Setup,
		httpfarm.Setup,
		httpfarmserver.Setup,
		httpfrontend.Setup,
		httproute.Setup,
		httprouterule.Setup,
		iploadbalancing.Setup,
		projectloadbalancer.Setup,
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
		logsencryptionkey.Setup,
		logsinput.Setup,
		logsoutputgraylogstream.Setup,
		logsoutputopensearchalias.Setup,
		logsoutputopensearchindex.Setup,
		logsrole.Setup,
		logsrolepermissionstream.Setup,
		logstoken.Setup,
		group.Setup,
		identityusertoken.Setup,
		oauth2client.Setup,
		userme.Setup,
		nashapartition.Setup,
		nashapartitionaccess.Setup,
		nashapartitionsnapshot.Setup,
		privatenetwork.Setup,
		privatevracknetwork.Setup,
		privatevracksubnet.Setup,
		projectregionnetwork.Setup,
		subnet.Setup,
		subnetv2.Setup,
		servicekeyjwk.Setup,
		vps.Setup,
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
		blockvolume.Setup,
		blockvolumebackup.Setup,
		blockvolumesnapshot.Setup,
		efs.Setup,
		efsshare.Setup,
		efsshareacl.Setup,
		efssharesnapshot.Setup,
		fileshare.Setup,
		fileshareacl.Setup,
		filesharenetwork.Setup,
		filesharesnapshot.Setup,
		projectfilestorageshare.Setup,
		projectfilestoragesharenetwork.Setup,
		projectregionstoragepresign.Setup,
		projectstoragelifecycleconfiguration.Setup,
		projectstoragereplicationjob.Setup,
		projectworkflowbackup.Setup,
		cloudproject.Setup,
		connectpopconfig.Setup,
		connectpopdatacenterconfig.Setup,
		connectpopdatacenterextraconfig.Setup,
		dedicatedcloud.Setup,
		dedicatedclouddatacenter.Setup,
		dedicatedserver.Setup,
		dedicatedserverinterface.Setup,
		ip.Setup,
		iploadbalancingvrack.Setup,
		ipv6.Setup,
		ipv6routedsubrange.Setup,
		ovhcloudconnect.Setup,
		publicroutingpriority.Setup,
		vrack.Setup,
		vrackservices.Setup,
		vrackservicesorder.Setup,
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
		floatingip.SetupGated,
		instance.SetupGated,
		instancegroup.SetupGated,
		plan.SetupGated,
		project.SetupGated,
		projectcontainerregistryiam.SetupGated,
		projectgatewayinterface.SetupGated,
		projectinstance.SetupGated,
		projectinstancesnapshot.SetupGated,
		projectrancher.SetupGated,
		projectregion.SetupGated,
		projectsshkey.SetupGated,
		projectstorage.SetupGated,
		projectvolume.SetupGated,
		projectvolumebackup.SetupGated,
		quota.SetupGated,
		s3credentials.SetupGated,
		s3policy.SetupGated,
		securitygroup.SetupGated,
		sshkey.SetupGated,
		user.SetupGated,
		cephacl.SetupGated,
		projectdatabase.SetupGated,
		projectdatabaseclickhouseuser.SetupGated,
		projectdatabasedatabase.SetupGated,
		projectdatabaseintegration.SetupGated,
		projectdatabasekafkaacl.SetupGated,
		projectdatabasekafkaschemaregistryacl.SetupGated,
		projectdatabasekafkatopic.SetupGated,
		projectdatabaselogsubscription.SetupGated,
		projectdatabasemongodbprometheus.SetupGated,
		projectdatabasemongodbuser.SetupGated,
		projectdatabaseopensearchpattern.SetupGated,
		projectdatabaseopensearchuser.SetupGated,
		projectdatabasepostgresqlconnectionpool.SetupGated,
		projectdatabasepostgresqluser.SetupGated,
		projectdatabaseprometheus.SetupGated,
		projectdatabaseuser.SetupGated,
		projectdatabasevalkeyuser.SetupGated,
		server.SetupGated,
		serverreinstalltask.SetupGated,
		servernetworking.SetupGated,
		serverreboottask.SetupGated,
		serverupdate.SetupGated,
		dsrecords.SetupGated,
		name.SetupGated,
		nameservers.SetupGated,
		zone.SetupGated,
		zonednssec.SetupGated,
		zonedynhostlogin.SetupGated,
		zonedynhostrecord.SetupGated,
		zonerecord.SetupGated,
		zoneredirection.SetupGated,
		zoneimport.SetupGated,
		domainaccount.SetupGated,
		cloudgateway.SetupGated,
		projectgateway.SetupGated,
		iampermissionsgroup.SetupGated,
		iampolicy.SetupGated,
		iamresourcegroup.SetupGated,
		iamresourcetags.SetupGated,
		credential.SetupGated,
		keymanagercontainer.SetupGated,
		keymanagercontainerconsumer.SetupGated,
		keymanagersecret.SetupGated,
		keymanagersecretconsumer.SetupGated,
		okms.SetupGated,
		secret.SetupGated,
		servicekey.SetupGated,
		cluster.SetupGated,
		iprestriction.SetupGated,
		logsubscription.SetupGated,
		nodepool.SetupGated,
		oidcconfiguration.SetupGated,
		httpfarm.SetupGated,
		httpfarmserver.SetupGated,
		httpfrontend.SetupGated,
		httproute.SetupGated,
		httprouterule.SetupGated,
		iploadbalancing.SetupGated,
		projectloadbalancer.SetupGated,
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
		logsencryptionkey.SetupGated,
		logsinput.SetupGated,
		logsoutputgraylogstream.SetupGated,
		logsoutputopensearchalias.SetupGated,
		logsoutputopensearchindex.SetupGated,
		logsrole.SetupGated,
		logsrolepermissionstream.SetupGated,
		logstoken.SetupGated,
		group.SetupGated,
		identityusertoken.SetupGated,
		oauth2client.SetupGated,
		userme.SetupGated,
		nashapartition.SetupGated,
		nashapartitionaccess.SetupGated,
		nashapartitionsnapshot.SetupGated,
		privatenetwork.SetupGated,
		privatevracknetwork.SetupGated,
		privatevracksubnet.SetupGated,
		projectregionnetwork.SetupGated,
		subnet.SetupGated,
		subnetv2.SetupGated,
		servicekeyjwk.SetupGated,
		vps.SetupGated,
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
		blockvolume.SetupGated,
		blockvolumebackup.SetupGated,
		blockvolumesnapshot.SetupGated,
		efs.SetupGated,
		efsshare.SetupGated,
		efsshareacl.SetupGated,
		efssharesnapshot.SetupGated,
		fileshare.SetupGated,
		fileshareacl.SetupGated,
		filesharenetwork.SetupGated,
		filesharesnapshot.SetupGated,
		projectfilestorageshare.SetupGated,
		projectfilestoragesharenetwork.SetupGated,
		projectregionstoragepresign.SetupGated,
		projectstoragelifecycleconfiguration.SetupGated,
		projectstoragereplicationjob.SetupGated,
		projectworkflowbackup.SetupGated,
		cloudproject.SetupGated,
		connectpopconfig.SetupGated,
		connectpopdatacenterconfig.SetupGated,
		connectpopdatacenterextraconfig.SetupGated,
		dedicatedcloud.SetupGated,
		dedicatedclouddatacenter.SetupGated,
		dedicatedserver.SetupGated,
		dedicatedserverinterface.SetupGated,
		ip.SetupGated,
		iploadbalancingvrack.SetupGated,
		ipv6.SetupGated,
		ipv6routedsubrange.SetupGated,
		ovhcloudconnect.SetupGated,
		publicroutingpriority.SetupGated,
		vrack.SetupGated,
		vrackservices.SetupGated,
		vrackservicesorder.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
