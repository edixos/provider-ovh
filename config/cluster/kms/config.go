package kms

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "kms"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_okms", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_okms_service_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_okms_credential", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_okms_secret", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_key_manager_container", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "KeyManagerContainer"
	})
	p.AddResourceConfigurator("ovh_cloud_key_manager_container_consumer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "KeyManagerContainerConsumer"
		r.References["container_id"] = config.Reference{
			TerraformName: "ovh_cloud_key_manager_container",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_key_manager_secret", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "KeyManagerSecret"
	})
	p.AddResourceConfigurator("ovh_cloud_key_manager_secret_consumer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "KeyManagerSecretConsumer"
		r.References["secret_id"] = config.Reference{
			TerraformName: "ovh_cloud_key_manager_secret",
		}
	})
}
