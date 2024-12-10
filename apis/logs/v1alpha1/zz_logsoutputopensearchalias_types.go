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

type LogsOutputOpensearchAliasInitParameters struct {

	// Alias description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indexes attached to alias
	// +listType=set
	Indexes []*string `json:"indexes,omitempty" tf:"indexes,omitempty"`

	// Number of index
	NbIndex *float64 `json:"nbIndex,omitempty" tf:"nb_index,omitempty"`

	// Number of shard
	NbStream *float64 `json:"nbStream,omitempty" tf:"nb_stream,omitempty"`

	// The service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Streams attached to alias
	// +listType=set
	Streams []*string `json:"streams,omitempty" tf:"streams,omitempty"`

	// Alias suffix
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type LogsOutputOpensearchAliasObservation struct {

	// Alias used
	AliasID *string `json:"aliasId,omitempty" tf:"alias_id,omitempty"`

	// Operation creation
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Current alias size (in bytes)
	CurrentSize *float64 `json:"currentSize,omitempty" tf:"current_size,omitempty"`

	// Alias description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indexes attached to alias
	// +listType=set
	Indexes []*string `json:"indexes,omitempty" tf:"indexes,omitempty"`

	// Indicates if you are allowed to edit entry
	IsEditable *bool `json:"isEditable,omitempty" tf:"is_editable,omitempty"`

	// Alias name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of index
	NbIndex *float64 `json:"nbIndex,omitempty" tf:"nb_index,omitempty"`

	// Number of shard
	NbStream *float64 `json:"nbStream,omitempty" tf:"nb_stream,omitempty"`

	// The service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Streams attached to alias
	// +listType=set
	Streams []*string `json:"streams,omitempty" tf:"streams,omitempty"`

	// Alias suffix
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`

	// Operation last update
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type LogsOutputOpensearchAliasParameters struct {

	// Alias description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indexes attached to alias
	// +kubebuilder:validation:Optional
	// +listType=set
	Indexes []*string `json:"indexes,omitempty" tf:"indexes,omitempty"`

	// Number of index
	// +kubebuilder:validation:Optional
	NbIndex *float64 `json:"nbIndex,omitempty" tf:"nb_index,omitempty"`

	// Number of shard
	// +kubebuilder:validation:Optional
	NbStream *float64 `json:"nbStream,omitempty" tf:"nb_stream,omitempty"`

	// The service name
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Streams attached to alias
	// +kubebuilder:validation:Optional
	// +listType=set
	Streams []*string `json:"streams,omitempty" tf:"streams,omitempty"`

	// Alias suffix
	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

// LogsOutputOpensearchAliasSpec defines the desired state of LogsOutputOpensearchAlias
type LogsOutputOpensearchAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogsOutputOpensearchAliasParameters `json:"forProvider"`
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
	InitProvider LogsOutputOpensearchAliasInitParameters `json:"initProvider,omitempty"`
}

// LogsOutputOpensearchAliasStatus defines the observed state of LogsOutputOpensearchAlias.
type LogsOutputOpensearchAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogsOutputOpensearchAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogsOutputOpensearchAlias is the Schema for the LogsOutputOpensearchAliass API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type LogsOutputOpensearchAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.suffix) || (has(self.initProvider) && has(self.initProvider.suffix))",message="spec.forProvider.suffix is a required parameter"
	Spec   LogsOutputOpensearchAliasSpec   `json:"spec"`
	Status LogsOutputOpensearchAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogsOutputOpensearchAliasList contains a list of LogsOutputOpensearchAliass
type LogsOutputOpensearchAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogsOutputOpensearchAlias `json:"items"`
}

// Repository type metadata.
var (
	LogsOutputOpensearchAlias_Kind             = "LogsOutputOpensearchAlias"
	LogsOutputOpensearchAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogsOutputOpensearchAlias_Kind}.String()
	LogsOutputOpensearchAlias_KindAPIVersion   = LogsOutputOpensearchAlias_Kind + "." + CRDGroupVersion.String()
	LogsOutputOpensearchAlias_GroupVersionKind = CRDGroupVersion.WithKind(LogsOutputOpensearchAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&LogsOutputOpensearchAlias{}, &LogsOutputOpensearchAliasList{})
}