package object_storage

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "storage"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_region_storage_presign", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_storage_efs", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EFS"
	})
	p.AddResourceConfigurator("ovh_cloud_storage_block_volume", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BlockVolume"
	})
	p.AddResourceConfigurator("ovh_cloud_storage_block_volume_backup", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BlockVolumeBackup"
		r.References["volume_id"] = config.Reference{
			TerraformName: "ovh_cloud_storage_block_volume",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_storage_block_volume_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BlockVolumeSnapshot"
		r.References["volume_id"] = config.Reference{
			TerraformName: "ovh_cloud_storage_block_volume",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_storage_file_share_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FileShareNetwork"
	})
	p.AddResourceConfigurator("ovh_cloud_storage_file_share", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FileShare"
		r.References["share_network_id"] = config.Reference{
			TerraformName: "ovh_cloud_storage_file_share_network",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_storage_file_share_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FileShareSnapshot"
		r.References["share_id"] = config.Reference{
			TerraformName: "ovh_cloud_storage_file_share",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_file_storage_share_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectFileStorageShareNetwork"
	})
	p.AddResourceConfigurator("ovh_cloud_project_file_storage_share", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectFileStorageShare"
		r.References["share_network_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_file_storage_share_network",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_storage_object_bucket_lifecycle_configuration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectStorageLifecycleConfiguration"
	})
	p.AddResourceConfigurator("ovh_cloud_project_storage_replication_job", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectStorageReplicationJob"
	})
}
