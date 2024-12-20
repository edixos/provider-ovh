/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectDatabaseKafkaSchemaregistryaclInitParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Permission to give to this username on this resource
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Resource affected by this acl
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Username affected by this acl
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProjectDatabaseKafkaSchemaregistryaclObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Permission to give to this username on this resource
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Resource affected by this acl
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Username affected by this acl
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProjectDatabaseKafkaSchemaregistryaclParameters struct {

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Permission to give to this username on this resource
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Resource affected by this acl
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Username affected by this acl
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// ProjectDatabaseKafkaSchemaregistryaclSpec defines the desired state of ProjectDatabaseKafkaSchemaregistryacl
type ProjectDatabaseKafkaSchemaregistryaclSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseKafkaSchemaregistryaclParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProjectDatabaseKafkaSchemaregistryaclInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseKafkaSchemaregistryaclStatus defines the observed state of ProjectDatabaseKafkaSchemaregistryacl.
type ProjectDatabaseKafkaSchemaregistryaclStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseKafkaSchemaregistryaclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseKafkaSchemaregistryacl is the Schema for the ProjectDatabaseKafkaSchemaregistryacls API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectDatabaseKafkaSchemaregistryacl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permission) || (has(self.initProvider) && has(self.initProvider.permission))",message="spec.forProvider.permission is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resource) || (has(self.initProvider) && has(self.initProvider.resource))",message="spec.forProvider.resource is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   ProjectDatabaseKafkaSchemaregistryaclSpec   `json:"spec"`
	Status ProjectDatabaseKafkaSchemaregistryaclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseKafkaSchemaregistryaclList contains a list of ProjectDatabaseKafkaSchemaregistryacls
type ProjectDatabaseKafkaSchemaregistryaclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseKafkaSchemaregistryacl `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseKafkaSchemaregistryacl_Kind             = "ProjectDatabaseKafkaSchemaregistryacl"
	ProjectDatabaseKafkaSchemaregistryacl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseKafkaSchemaregistryacl_Kind}.String()
	ProjectDatabaseKafkaSchemaregistryacl_KindAPIVersion   = ProjectDatabaseKafkaSchemaregistryacl_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseKafkaSchemaregistryacl_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseKafkaSchemaregistryacl_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseKafkaSchemaregistryacl{}, &ProjectDatabaseKafkaSchemaregistryaclList{})
}
