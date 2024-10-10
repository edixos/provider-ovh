package databases

import (
	"errors"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "databases"
)

var ErrTransformEndpointToConnection = errors.New("Error enriching Connection Details")

func addConnectionInfo(attr map[string]any) (map[string][]byte, error) {
	conn := map[string][]byte{}

	if endpoints, ok := attr["endpoints"]; ok {
		if endpointList, ok := endpoints.([]any); ok {
			for idx, item := range endpointList {
				if endpoint, ok := item.(map[string]any); ok {
					for key, value := range endpoint {
						endpointKey := fmt.Sprintf("endpoint_%d_%s", idx, key)

						var val []byte
						switch value.(type) {
						case float64:
							val = []byte(fmt.Sprintf("%.0f", value))
						case bool:
							val = []byte(fmt.Sprintf("%b", value))
						case string:
							val = []byte(fmt.Sprintf("%s", value))
						default:
							continue
						}

						conn[endpointKey] = val
					}
				}
			}
		} else {
			return nil, fmt.Errorf("Could not decode endpoints: %w", ErrTransformEndpointToConnection)
		}
	} else {
		return nil, fmt.Errorf("Could not find endpoints: %w", ErrTransformEndpointToConnection)
	}

	return conn, nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Sensitive.AdditionalConnectionDetailsFn = addConnectionInfo
		r.UseAsync = true
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

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			conn["username"] = []byte(attr["name"].(string))

			return conn, nil
		}
	})

	p.AddResourceConfigurator("ovh_cloud_project_database_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_database",
		}
	})
}
