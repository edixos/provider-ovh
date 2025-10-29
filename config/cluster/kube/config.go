package kube

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "kube"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_kube", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.UseAsync = true
		r.Kind = "Cluster"
		r.References["private_network_id"] = config.Reference{
			Type:      "github.com/edixos/provider-ovh/apis/cluster/network/v1alpha1.PrivateNetwork",
			Extractor: "github.com/edixos/provider-ovh/config/common.PrivateNetworkOpenStackIdExtractor()",
		}
		r.References["nodes_subnet_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/network/v1alpha1.Subnet",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_iprestrictions", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/kube/v1alpha1.Cluster",
		}
		r.Kind = "IpRestriction"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_nodepool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.UseAsync = true
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/kube/v1alpha1.Cluster",
		}
		r.Kind = "NodePool"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/kube/v1alpha1.Cluster",
		}
		r.Kind = "OIDCConfiguration"
	})
}
