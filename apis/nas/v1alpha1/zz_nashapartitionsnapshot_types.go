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

type NashaPartitionSnapshotInitParameters struct {
	PartitionName *string `json:"partitionName,omitempty" tf:"partition_name,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NashaPartitionSnapshotObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PartitionName *string `json:"partitionName,omitempty" tf:"partition_name,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NashaPartitionSnapshotParameters struct {

	// +kubebuilder:validation:Optional
	PartitionName *string `json:"partitionName,omitempty" tf:"partition_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// NashaPartitionSnapshotSpec defines the desired state of NashaPartitionSnapshot
type NashaPartitionSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NashaPartitionSnapshotParameters `json:"forProvider"`
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
	InitProvider NashaPartitionSnapshotInitParameters `json:"initProvider,omitempty"`
}

// NashaPartitionSnapshotStatus defines the observed state of NashaPartitionSnapshot.
type NashaPartitionSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NashaPartitionSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NashaPartitionSnapshot is the Schema for the NashaPartitionSnapshots API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type NashaPartitionSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionName) || (has(self.initProvider) && has(self.initProvider.partitionName))",message="spec.forProvider.partitionName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   NashaPartitionSnapshotSpec   `json:"spec"`
	Status NashaPartitionSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NashaPartitionSnapshotList contains a list of NashaPartitionSnapshots
type NashaPartitionSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NashaPartitionSnapshot `json:"items"`
}

// Repository type metadata.
var (
	NashaPartitionSnapshot_Kind             = "NashaPartitionSnapshot"
	NashaPartitionSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NashaPartitionSnapshot_Kind}.String()
	NashaPartitionSnapshot_KindAPIVersion   = NashaPartitionSnapshot_Kind + "." + CRDGroupVersion.String()
	NashaPartitionSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(NashaPartitionSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&NashaPartitionSnapshot{}, &NashaPartitionSnapshotList{})
}
