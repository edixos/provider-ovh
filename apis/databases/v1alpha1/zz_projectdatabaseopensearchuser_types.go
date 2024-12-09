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

type AclsInitParameters struct {

	// Pattern of the ACL
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Permission of the ACL
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type AclsObservation struct {

	// Pattern of the ACL
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Permission of the ACL
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type AclsParameters struct {

	// Pattern of the ACL
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern" tf:"pattern,omitempty"`

	// Permission of the ACL
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`
}

type ProjectDatabaseOpensearchUserInitParameters struct {

	// Acls of the user
	Acls []AclsInitParameters `json:"acls,omitempty" tf:"acls,omitempty"`

	// Id of the database cluster
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/databases/v1alpha1.ProjectDatabase
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a ProjectDatabase in databases to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseOpensearchUserObservation struct {

	// Acls of the user
	Acls []AclsObservation `json:"acls,omitempty" tf:"acls,omitempty"`

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Date of the creation of the user
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Current status of the user
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ProjectDatabaseOpensearchUserParameters struct {

	// Acls of the user
	// +kubebuilder:validation:Optional
	Acls []AclsParameters `json:"acls,omitempty" tf:"acls,omitempty"`

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

	// Name of the user
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	// +kubebuilder:validation:Optional
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectDatabaseOpensearchUserSpec defines the desired state of ProjectDatabaseOpensearchUser
type ProjectDatabaseOpensearchUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseOpensearchUserParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseOpensearchUserInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseOpensearchUserStatus defines the observed state of ProjectDatabaseOpensearchUser.
type ProjectDatabaseOpensearchUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseOpensearchUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseOpensearchUser is the Schema for the ProjectDatabaseOpensearchUsers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectDatabaseOpensearchUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseOpensearchUserSpec   `json:"spec"`
	Status ProjectDatabaseOpensearchUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseOpensearchUserList contains a list of ProjectDatabaseOpensearchUsers
type ProjectDatabaseOpensearchUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseOpensearchUser `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseOpensearchUser_Kind             = "ProjectDatabaseOpensearchUser"
	ProjectDatabaseOpensearchUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseOpensearchUser_Kind}.String()
	ProjectDatabaseOpensearchUser_KindAPIVersion   = ProjectDatabaseOpensearchUser_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseOpensearchUser_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseOpensearchUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseOpensearchUser{}, &ProjectDatabaseOpensearchUserList{})
}
