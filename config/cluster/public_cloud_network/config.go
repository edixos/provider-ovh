package public_cloud_network

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "network"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_network_private", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNetwork"
	})
	p.AddResourceConfigurator("ovh_cloud_project_network_private_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Subnet"
		r.References["network_id"] = config.Reference{
			Type:      "github.com/edixos/provider-ovh/apis/cluster/network/v1alpha1.PrivateNetwork",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_network_private_subnet_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SubnetV2"
		r.References["network_id"] = config.Reference{
			Type:      "github.com/edixos/provider-ovh/apis/cluster/network/v1alpha1.PrivateNetwork",
			Extractor: "github.com/edixos/provider-ovh/config/common.PrivateNetworkOpenStackIdExtractor()",
		}
	})
}
