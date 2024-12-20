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

type CustomizationInitParameters struct {

	// Set up the server using the provided hostname instead of the default hostname
	CustomHostname *string `json:"customHostname,omitempty" tf:"custom_hostname,omitempty"`
}

type CustomizationObservation struct {

	// Set up the server using the provided hostname instead of the default hostname
	CustomHostname *string `json:"customHostname,omitempty" tf:"custom_hostname,omitempty"`
}

type CustomizationParameters struct {

	// Set up the server using the provided hostname instead of the default hostname
	// +kubebuilder:validation:Optional
	CustomHostname *string `json:"customHostname,omitempty" tf:"custom_hostname,omitempty"`
}

type InputsInitParameters struct {
}

type InputsObservation struct {
	Default *string `json:"default,omitempty" tf:"default,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Enum []*string `json:"enum,omitempty" tf:"enum,omitempty"`

	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InputsParameters struct {
}

type InstallationTemplateInitParameters struct {

	// OVH template name yours will be based on, choose one among the list given by compatibleTemplates function
	BaseTemplateName *string `json:"baseTemplateName,omitempty" tf:"base_template_name,omitempty"`

	Customization []CustomizationInitParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// Remove default partition schemes at creation
	RemoveDefaultPartitionSchemes *bool `json:"removeDefaultPartitionSchemes,omitempty" tf:"remove_default_partition_schemes,omitempty"`

	// This template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type InstallationTemplateObservation struct {

	// OVH template name yours will be based on, choose one among the list given by compatibleTemplates function
	BaseTemplateName *string `json:"baseTemplateName,omitempty" tf:"base_template_name,omitempty"`

	// This template bit format (32 or 64)
	BitFormat *float64 `json:"bitFormat,omitempty" tf:"bit_format,omitempty"`

	// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation)
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	Customization []CustomizationObservation `json:"customization,omitempty" tf:"customization,omitempty"`

	// information about this template
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the distribution this template is based on
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// after this date, install of this template will not be possible at OVH
	EndOfInstall *string `json:"endOfInstall,omitempty" tf:"end_of_install,omitempty"`

	// this template family type
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Filesystems available
	Filesystems []*string `json:"filesystems,omitempty" tf:"filesystems,omitempty"`

	// This distribution supports hardware raid configuration through the OVH API
	HardRaidConfiguration *bool `json:"hardRaidConfiguration,omitempty" tf:"hard_raid_configuration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Inputs []InputsObservation `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// Whether this distribution supports Logical Volumes (Linux LVM)
	LvmReady *bool `json:"lvmReady,omitempty" tf:"lvm_ready,omitempty"`

	// Partitioning customization is not available for this OS template
	NoPartitioning *bool `json:"noPartitioning,omitempty" tf:"no_partitioning,omitempty"`

	// Remove default partition schemes at creation
	RemoveDefaultPartitionSchemes *bool `json:"removeDefaultPartitionSchemes,omitempty" tf:"remove_default_partition_schemes,omitempty"`

	// Partitioning customization is available but limited to mirroring for this OS template
	SoftRaidOnlyMirroring *bool `json:"softRaidOnlyMirroring,omitempty" tf:"soft_raid_only_mirroring,omitempty"`

	// this template subfamily type
	Subfamily *string `json:"subfamily,omitempty" tf:"subfamily,omitempty"`

	// This template name
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

type InstallationTemplateParameters struct {

	// OVH template name yours will be based on, choose one among the list given by compatibleTemplates function
	// +kubebuilder:validation:Optional
	BaseTemplateName *string `json:"baseTemplateName,omitempty" tf:"base_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Customization []CustomizationParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// Remove default partition schemes at creation
	// +kubebuilder:validation:Optional
	RemoveDefaultPartitionSchemes *bool `json:"removeDefaultPartitionSchemes,omitempty" tf:"remove_default_partition_schemes,omitempty"`

	// This template name
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`
}

// InstallationTemplateSpec defines the desired state of InstallationTemplate
type InstallationTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstallationTemplateParameters `json:"forProvider"`
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
	InitProvider InstallationTemplateInitParameters `json:"initProvider,omitempty"`
}

// InstallationTemplateStatus defines the observed state of InstallationTemplate.
type InstallationTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstallationTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstallationTemplate is the Schema for the InstallationTemplates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type InstallationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseTemplateName) || (has(self.initProvider) && has(self.initProvider.baseTemplateName))",message="spec.forProvider.baseTemplateName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateName) || (has(self.initProvider) && has(self.initProvider.templateName))",message="spec.forProvider.templateName is a required parameter"
	Spec   InstallationTemplateSpec   `json:"spec"`
	Status InstallationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallationTemplateList contains a list of InstallationTemplates
type InstallationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstallationTemplate `json:"items"`
}

// Repository type metadata.
var (
	InstallationTemplate_Kind             = "InstallationTemplate"
	InstallationTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstallationTemplate_Kind}.String()
	InstallationTemplate_KindAPIVersion   = InstallationTemplate_Kind + "." + CRDGroupVersion.String()
	InstallationTemplate_GroupVersionKind = CRDGroupVersion.WithKind(InstallationTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&InstallationTemplate{}, &InstallationTemplateList{})
}
