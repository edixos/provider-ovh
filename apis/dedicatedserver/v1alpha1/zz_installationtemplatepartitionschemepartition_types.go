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

type InstallationTemplatePartitionSchemePartitionInitParameters struct {

	// Partition filesystem
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`

	// partition mount point
	Mountpoint *string `json:"mountpoint,omitempty" tf:"mountpoint,omitempty"`

	// step or order. specifies the creation order of the partition on the disk
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// raid partition type
	Raid *string `json:"raid,omitempty" tf:"raid,omitempty"`

	// name of this partitioning scheme
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// size of partition in MB, 0 => rest of the space
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// partition type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The volume name needed for proxmox distribution
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

type InstallationTemplatePartitionSchemePartitionObservation struct {

	// Partition filesystem
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// partition mount point
	Mountpoint *string `json:"mountpoint,omitempty" tf:"mountpoint,omitempty"`

	// step or order. specifies the creation order of the partition on the disk
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// raid partition type
	Raid *string `json:"raid,omitempty" tf:"raid,omitempty"`

	// name of this partitioning scheme
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// size of partition in MB, 0 => rest of the space
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// partition type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The volume name needed for proxmox distribution
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

type InstallationTemplatePartitionSchemePartitionParameters struct {

	// Partition filesystem
	// +kubebuilder:validation:Optional
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`

	// partition mount point
	// +kubebuilder:validation:Optional
	Mountpoint *string `json:"mountpoint,omitempty" tf:"mountpoint,omitempty"`

	// step or order. specifies the creation order of the partition on the disk
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// raid partition type
	// +kubebuilder:validation:Optional
	Raid *string `json:"raid,omitempty" tf:"raid,omitempty"`

	// name of this partitioning scheme
	// +kubebuilder:validation:Optional
	SchemeName *string `json:"schemeName,omitempty" tf:"scheme_name,omitempty"`

	// size of partition in MB, 0 => rest of the space
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Template name
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// partition type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The volume name needed for proxmox distribution
	// +kubebuilder:validation:Optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

// InstallationTemplatePartitionSchemePartitionSpec defines the desired state of InstallationTemplatePartitionSchemePartition
type InstallationTemplatePartitionSchemePartitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstallationTemplatePartitionSchemePartitionParameters `json:"forProvider"`
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
	InitProvider InstallationTemplatePartitionSchemePartitionInitParameters `json:"initProvider,omitempty"`
}

// InstallationTemplatePartitionSchemePartitionStatus defines the observed state of InstallationTemplatePartitionSchemePartition.
type InstallationTemplatePartitionSchemePartitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstallationTemplatePartitionSchemePartitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstallationTemplatePartitionSchemePartition is the Schema for the InstallationTemplatePartitionSchemePartitions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type InstallationTemplatePartitionSchemePartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filesystem) || (has(self.initProvider) && has(self.initProvider.filesystem))",message="spec.forProvider.filesystem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mountpoint) || (has(self.initProvider) && has(self.initProvider.mountpoint))",message="spec.forProvider.mountpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.order) || (has(self.initProvider) && has(self.initProvider.order))",message="spec.forProvider.order is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemeName) || (has(self.initProvider) && has(self.initProvider.schemeName))",message="spec.forProvider.schemeName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateName) || (has(self.initProvider) && has(self.initProvider.templateName))",message="spec.forProvider.templateName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   InstallationTemplatePartitionSchemePartitionSpec   `json:"spec"`
	Status InstallationTemplatePartitionSchemePartitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallationTemplatePartitionSchemePartitionList contains a list of InstallationTemplatePartitionSchemePartitions
type InstallationTemplatePartitionSchemePartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstallationTemplatePartitionSchemePartition `json:"items"`
}

// Repository type metadata.
var (
	InstallationTemplatePartitionSchemePartition_Kind             = "InstallationTemplatePartitionSchemePartition"
	InstallationTemplatePartitionSchemePartition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstallationTemplatePartitionSchemePartition_Kind}.String()
	InstallationTemplatePartitionSchemePartition_KindAPIVersion   = InstallationTemplatePartitionSchemePartition_Kind + "." + CRDGroupVersion.String()
	InstallationTemplatePartitionSchemePartition_GroupVersionKind = CRDGroupVersion.WithKind(InstallationTemplatePartitionSchemePartition_Kind)
)

func init() {
	SchemeBuilder.Register(&InstallationTemplatePartitionSchemePartition{}, &InstallationTemplatePartitionSchemePartitionList{})
}
