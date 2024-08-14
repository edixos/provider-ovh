/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/edixos/provider-ovh/config/gateway"
	"github.com/edixos/provider-ovh/config/vps"

	"github.com/edixos/provider-ovh/config/cloud"
	"github.com/edixos/provider-ovh/config/iam"
	"github.com/edixos/provider-ovh/config/me"

	"github.com/edixos/provider-ovh/config/vrack"

	"github.com/edixos/provider-ovh/config/web_cloud_private_sql"

	"github.com/edixos/provider-ovh/config/object_storage"

	"github.com/edixos/provider-ovh/config/nas"

	"github.com/edixos/provider-ovh/config/registry"

	"github.com/edixos/provider-ovh/config/kube"

	"github.com/edixos/provider-ovh/config/databases"

	"github.com/edixos/provider-ovh/config/logs"

	"github.com/edixos/provider-ovh/config/lb"

	"github.com/edixos/provider-ovh/config/dns"

	"github.com/edixos/provider-ovh/config/dedicated_server"

	"github.com/edixos/provider-ovh/config/cloud_disk_array"

	"github.com/edixos/provider-ovh/config/additional_ip"
	"github.com/edixos/provider-ovh/config/vm_instances"

	"github.com/edixos/provider-ovh/config/public_cloud_network"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "ovh"
	modulePath     = "github.com/edixos/provider-ovh"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ovh.edixos.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		public_cloud_network.Configure,
		iam.Configure,
		me.Configure,
		cloud.Configure,
		additional_ip.Configure,
		vm_instances.Configure,
		cloud_disk_array.Configure,
		dedicated_server.Configure,
		dns.Configure,
		lb.Configure,
		logs.Configure,
		databases.Configure,
		kube.Configure,
		registry.Configure,
		nas.Configure,
		object_storage.Configure,
		web_cloud_private_sql.Configure,
		vrack.Configure,
		vps.Configure,
		gateway.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
