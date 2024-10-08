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

type MoveInitParameters struct {

	// Custom description on your ip
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Routage information
	RoutedTo []RoutedToInitParameters `json:"routedTo,omitempty" tf:"routed_to,omitempty"`
}

type MoveObservation struct {
	CanBeTerminated *bool `json:"canBeTerminated,omitempty" tf:"can_be_terminated,omitempty"`

	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Custom description on your ip
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	OrganisationID *string `json:"organisationId,omitempty" tf:"organisation_id,omitempty"`

	// Routage information
	RoutedTo []RoutedToObservation `json:"routedTo,omitempty" tf:"routed_to,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStartDate *string `json:"taskStartDate,omitempty" tf:"task_start_date,omitempty"`

	// Status field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStatus *string `json:"taskStatus,omitempty" tf:"task_status,omitempty"`

	// Possible values for ip type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MoveParameters struct {

	// Custom description on your ip
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Routage information
	// +kubebuilder:validation:Optional
	RoutedTo []RoutedToParameters `json:"routedTo,omitempty" tf:"routed_to,omitempty"`
}

type RoutedToInitParameters struct {

	// Service where ip is routed to
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type RoutedToObservation struct {

	// Service where ip is routed to
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type RoutedToParameters struct {

	// Service where ip is routed to
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// MoveSpec defines the desired state of Move
type MoveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MoveParameters `json:"forProvider"`
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
	InitProvider MoveInitParameters `json:"initProvider,omitempty"`
}

// MoveStatus defines the observed state of Move.
type MoveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MoveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Move is the Schema for the Moves API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type Move struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ip) || (has(self.initProvider) && has(self.initProvider.ip))",message="spec.forProvider.ip is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routedTo) || (has(self.initProvider) && has(self.initProvider.routedTo))",message="spec.forProvider.routedTo is a required parameter"
	Spec   MoveSpec   `json:"spec"`
	Status MoveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MoveList contains a list of Moves
type MoveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Move `json:"items"`
}

// Repository type metadata.
var (
	Move_Kind             = "Move"
	Move_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Move_Kind}.String()
	Move_KindAPIVersion   = Move_Kind + "." + CRDGroupVersion.String()
	Move_GroupVersionKind = CRDGroupVersion.WithKind(Move_Kind)
)

func init() {
	SchemeBuilder.Register(&Move{}, &MoveList{})
}
