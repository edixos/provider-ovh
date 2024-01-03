// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IpxeScriptInitParameters struct {

	// For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Content of your IPXE script
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type IpxeScriptObservation struct {

	// For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Content of your IPXE script
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type IpxeScriptParameters struct {

	// For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Content of your IPXE script
	// +kubebuilder:validation:Optional
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

// IpxeScriptSpec defines the desired state of IpxeScript
type IpxeScriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpxeScriptParameters `json:"forProvider"`
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
	InitProvider IpxeScriptInitParameters `json:"initProvider,omitempty"`
}

// IpxeScriptStatus defines the observed state of IpxeScript.
type IpxeScriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpxeScriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpxeScript is the Schema for the IpxeScripts API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type IpxeScript struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.script) || (has(self.initProvider) && has(self.initProvider.script))",message="spec.forProvider.script is a required parameter"
	Spec   IpxeScriptSpec   `json:"spec"`
	Status IpxeScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpxeScriptList contains a list of IpxeScripts
type IpxeScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpxeScript `json:"items"`
}

// Repository type metadata.
var (
	IpxeScript_Kind             = "IpxeScript"
	IpxeScript_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpxeScript_Kind}.String()
	IpxeScript_KindAPIVersion   = IpxeScript_Kind + "." + CRDGroupVersion.String()
	IpxeScript_GroupVersionKind = CRDGroupVersion.WithKind(IpxeScript_Kind)
)

func init() {
	SchemeBuilder.Register(&IpxeScript{}, &IpxeScriptList{})
}