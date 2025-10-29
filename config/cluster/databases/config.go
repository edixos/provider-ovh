package databases

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "databases"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_ip_restriction", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_kafka_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_kafka_schemaregistryacl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_m3db_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_m3db_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_mongodb_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_opensearch_pattern", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_opensearch_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_postgresql_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_postgresql_connection_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_redis_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_database_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
}
