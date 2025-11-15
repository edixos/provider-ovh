/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	kms__cluster "github.com/edixos/provider-ovh/config/cluster/kms"
	kms__namespaced "github.com/edixos/provider-ovh/config/namespaced/kms"

	gateway__cluster "github.com/edixos/provider-ovh/config/cluster/gateway"
	vps__cluster "github.com/edixos/provider-ovh/config/cluster/vps"
	gateway__namespaced "github.com/edixos/provider-ovh/config/namespaced/gateway"
	vps__namespaced "github.com/edixos/provider-ovh/config/namespaced/vps"

	cloud__cluster "github.com/edixos/provider-ovh/config/cluster/cloud"
	iam__cluster "github.com/edixos/provider-ovh/config/cluster/iam"
	me__cluster "github.com/edixos/provider-ovh/config/cluster/me"
	cloud__namespaced "github.com/edixos/provider-ovh/config/namespaced/cloud"
	iam__namespaced "github.com/edixos/provider-ovh/config/namespaced/iam"
	me__namespaced "github.com/edixos/provider-ovh/config/namespaced/me"

	vrack__cluster "github.com/edixos/provider-ovh/config/cluster/vrack"
	vrack__namespaced "github.com/edixos/provider-ovh/config/namespaced/vrack"

	web_cloud_private_sql__cluster "github.com/edixos/provider-ovh/config/cluster/web_cloud_private_sql"
	web_cloud_private_sql__namespaced "github.com/edixos/provider-ovh/config/namespaced/web_cloud_private_sql"

	object_storage__cluster "github.com/edixos/provider-ovh/config/cluster/object_storage"
	object_storage__namespaced "github.com/edixos/provider-ovh/config/namespaced/object_storage"

	nas__cluster "github.com/edixos/provider-ovh/config/cluster/nas"
	nas__namespaced "github.com/edixos/provider-ovh/config/namespaced/nas"

	registry__cluster "github.com/edixos/provider-ovh/config/cluster/registry"
	registry__namespaced "github.com/edixos/provider-ovh/config/namespaced/registry"

	kube__cluster "github.com/edixos/provider-ovh/config/cluster/kube"
	kube__namespaced "github.com/edixos/provider-ovh/config/namespaced/kube"

	databases__cluster "github.com/edixos/provider-ovh/config/cluster/databases"
	databases__namespaced "github.com/edixos/provider-ovh/config/namespaced/databases"

	logs__cluster "github.com/edixos/provider-ovh/config/cluster/logs"
	logs__namespaced "github.com/edixos/provider-ovh/config/namespaced/logs"

	lb__cluster "github.com/edixos/provider-ovh/config/cluster/lb"
	lb__namespaced "github.com/edixos/provider-ovh/config/namespaced/lb"

	dns__cluster "github.com/edixos/provider-ovh/config/cluster/dns"
	dns__namespaced "github.com/edixos/provider-ovh/config/namespaced/dns"

	dedicated_server__cluster "github.com/edixos/provider-ovh/config/cluster/dedicated_server"
	dedicated_server__namespaced "github.com/edixos/provider-ovh/config/namespaced/dedicated_server"

	cloud_disk_array__cluster "github.com/edixos/provider-ovh/config/cluster/cloud_disk_array"
	cloud_disk_array__namespaced "github.com/edixos/provider-ovh/config/namespaced/cloud_disk_array"

	additional_ip__cluster "github.com/edixos/provider-ovh/config/cluster/additional_ip"
	vm_instances__cluster "github.com/edixos/provider-ovh/config/cluster/vm_instances"
	additional_ip__namespaced "github.com/edixos/provider-ovh/config/namespaced/additional_ip"
	vm_instances__namespaced "github.com/edixos/provider-ovh/config/namespaced/vm_instances"

	public_cloud_network__cluster "github.com/edixos/provider-ovh/config/cluster/public_cloud_network"
	public_cloud_network__namespaced "github.com/edixos/provider-ovh/config/namespaced/public_cloud_network"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
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
		public_cloud_network__cluster.Configure,
		iam__cluster.Configure,
		me__cluster.Configure,
		cloud__cluster.Configure,
		additional_ip__cluster.Configure,
		vm_instances__cluster.Configure,
		cloud_disk_array__cluster.Configure,
		dedicated_server__cluster.Configure,
		dns__cluster.Configure,
		lb__cluster.Configure,
		logs__cluster.Configure,
		databases__cluster.Configure,
		kube__cluster.Configure,
		registry__cluster.Configure,
		nas__cluster.Configure,
		object_storage__cluster.Configure,
		web_cloud_private_sql__cluster.Configure,
		vrack__cluster.Configure,
		vps__cluster.Configure,
		gateway__cluster.Configure,
		kms__cluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ovh.m.edixos.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		public_cloud_network__namespaced.Configure,
		iam__namespaced.Configure,
		me__namespaced.Configure,
		cloud__namespaced.Configure,
		additional_ip__namespaced.Configure,
		vm_instances__namespaced.Configure,
		cloud_disk_array__namespaced.Configure,
		dedicated_server__namespaced.Configure,
		dns__namespaced.Configure,
		lb__namespaced.Configure,
		logs__namespaced.Configure,
		databases__namespaced.Configure,
		kube__namespaced.Configure,
		registry__namespaced.Configure,
		nas__namespaced.Configure,
		object_storage__namespaced.Configure,
		web_cloud_private_sql__namespaced.Configure,
		vrack__namespaced.Configure,
		vps__namespaced.Configure,
		gateway__namespaced.Configure,
		kms__namespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
