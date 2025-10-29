package registry

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "registry"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContainerRegistry"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContainerRegistryOIDC"
		r.References["registry_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/registry/v1alpha1.ContainerRegistry",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_ip_restrictions_management", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContainerRegistryIPRestrictionsManagement"
		r.References["registry_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/registry/v1alpha1.ContainerRegistry",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_ip_restrictions_registry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContainerRegistryIPRestrictionsRegistry"
		r.References["registry_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/registry/v1alpha1.ContainerRegistry",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContainerRegistryUser"
		r.References["registry_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cluster/registry/v1alpha1.ContainerRegistry",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["login"].(string); ok {
				conn["login"] = []byte(a)
			}
			return conn, nil
		}
	})
}
