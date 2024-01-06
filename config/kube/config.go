package kube

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "kube"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_kube", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_iprestrictions", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/kube/v1alpha1.ProjectKube",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_nodepool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.UseAsync = true
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/kube/v1alpha1.ProjectKube",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["kube_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/kube/v1alpha1.ProjectKube",
		}
	})
}
