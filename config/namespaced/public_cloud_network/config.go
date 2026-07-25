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
			TerraformName: "ovh_cloud_project_network_private",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_network_private_subnet_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SubnetV2"
		r.References["network_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_network_private",
			Extractor:     "github.com/edixos/provider-ovh/config/common.PrivateNetworkOpenStackIdExtractor()",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_region_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_network_private_vrack", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrivateVrackNetwork"
	})
	p.AddResourceConfigurator("ovh_cloud_network_private_vrack_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrivateVrackSubnet"
		r.References["network_id"] = config.Reference{
			TerraformName: "ovh_cloud_network_private_vrack",
		}
	})
}
