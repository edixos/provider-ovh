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

type ProjectDatabaseM3DbUserInitParameters struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Group of the user
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectDatabaseM3DbUserObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Date of the creation of the user
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Group of the user
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the user
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Current status of the user
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ProjectDatabaseM3DbUserParameters struct {

	// Id of the database cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Group of the user
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Name of the user
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Arbitrary string to change to trigger a password update
	// +kubebuilder:validation:Optional
	PasswordReset *string `json:"passwordReset,omitempty" tf:"password_reset,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectDatabaseM3DbUserSpec defines the desired state of ProjectDatabaseM3DbUser
type ProjectDatabaseM3DbUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseM3DbUserParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseM3DbUserInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseM3DbUserStatus defines the observed state of ProjectDatabaseM3DbUser.
type ProjectDatabaseM3DbUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseM3DbUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectDatabaseM3DbUser is the Schema for the ProjectDatabaseM3DbUsers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type ProjectDatabaseM3DbUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterId) || (has(self.initProvider) && has(self.initProvider.clusterId))",message="spec.forProvider.clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectDatabaseM3DbUserSpec   `json:"spec"`
	Status ProjectDatabaseM3DbUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseM3DbUserList contains a list of ProjectDatabaseM3DbUsers
type ProjectDatabaseM3DbUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseM3DbUser `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseM3DbUser_Kind             = "ProjectDatabaseM3DbUser"
	ProjectDatabaseM3DbUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseM3DbUser_Kind}.String()
	ProjectDatabaseM3DbUser_KindAPIVersion   = ProjectDatabaseM3DbUser_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseM3DbUser_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseM3DbUser_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseM3DbUser{}, &ProjectDatabaseM3DbUserList{})
}
