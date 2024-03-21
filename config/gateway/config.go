package gateway

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "gateway"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectGateway"
		r.References["network_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/network/v1alpha1.Subnet",
		}
	})
}
