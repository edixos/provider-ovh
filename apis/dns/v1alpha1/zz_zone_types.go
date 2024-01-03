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

type ConfigurationInitParameters struct {

	// Identifier of the resource
	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationObservation struct {

	// Identifier of the resource
	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationParameters struct {

	// Identifier of the resource
	// Identifier of the resource
	// +kubebuilder:validation:Optional
	Label *string `json:"label" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DetailsInitParameters struct {
}

type DetailsObservation struct {

	// description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// expiration date
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// order detail id
	OrderDetailID *float64 `json:"orderDetailId,omitempty" tf:"order_detail_id,omitempty"`

	// quantity
	Quantity *string `json:"quantity,omitempty" tf:"quantity,omitempty"`
}

type DetailsParameters struct {
}

type OrderInitParameters struct {
}

type OrderObservation struct {

	// date
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Information about a Bill entry
	Details []DetailsObservation `json:"details,omitempty" tf:"details,omitempty"`

	// expiration date
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// order id
	OrderID *float64 `json:"orderId,omitempty" tf:"order_id,omitempty"`
}

type OrderParameters struct {
}

type PlanInitParameters struct {

	// Catalog name
	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanObservation struct {

	// Catalog name
	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionConfigurationInitParameters struct {

	// Identifier of the resource
	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlanOptionConfigurationObservation struct {

	// Identifier of the resource
	// Identifier of the resource
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlanOptionConfigurationParameters struct {

	// Identifier of the resource
	// Identifier of the resource
	// +kubebuilder:validation:Optional
	Label *string `json:"label" tf:"label,omitempty"`

	// Path to the resource in API.OVH.COM
	// Path to the resource in API.OVH.COM
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type PlanOptionInitParameters struct {

	// Catalog name
	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	Configuration []PlanOptionConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionObservation struct {

	// Catalog name
	// Catalog name
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	Configuration []PlanOptionConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	PlanCode *string `json:"planCode,omitempty" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	PricingMode *string `json:"pricingMode,omitempty" tf:"pricing_mode,omitempty"`
}

type PlanOptionParameters struct {

	// Catalog name
	// Catalog name
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	// +kubebuilder:validation:Optional
	Configuration []PlanOptionConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	// +kubebuilder:validation:Optional
	PlanCode *string `json:"planCode" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	// +kubebuilder:validation:Optional
	PricingMode *string `json:"pricingMode" tf:"pricing_mode,omitempty"`
}

type PlanParameters struct {

	// Catalog name
	// Catalog name
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Representation of a configuration item for personalizing product
	// Representation of a configuration item for personalizing product
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// duration
	// duration
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// Plan code
	// Plan code
	// +kubebuilder:validation:Optional
	PlanCode *string `json:"planCode" tf:"plan_code,omitempty"`

	// Pricing model identifier
	// Pricing model identifier
	// +kubebuilder:validation:Optional
	PricingMode *string `json:"pricingMode" tf:"pricing_mode,omitempty"`
}

type ZoneInitParameters struct {

	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at /1.0/me.json under
	// Ovh Subsidiary
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	// Product Plan to order
	Plan []PlanInitParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	// Product Plan to order
	PlanOption []PlanOptionInitParameters `json:"planOption,omitempty" tf:"plan_option,omitempty"`
}

type ZoneObservation struct {

	// Is DNSSEC supported by this zone
	// Is DNSSEC supported by this zone
	DNSSECSupported *bool `json:"dnssecSupported,omitempty" tf:"dnssec_supported,omitempty"`

	// hasDnsAnycast flag of the DNS zone
	// hasDnsAnycast flag of the DNS zone
	HasDNSAnycast *bool `json:"hasDnsAnycast,omitempty" tf:"has_dns_anycast,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update date of the DNS zone
	// Last update date of the DNS zone
	LastUpdate *string `json:"lastUpdate,omitempty" tf:"last_update,omitempty"`

	// Name servers that host the DNS zone
	// Name servers that host the DNS zone
	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	// Details about an Order
	// Details about an Order
	Order []OrderObservation `json:"order,omitempty" tf:"order,omitempty"`

	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at /1.0/me.json under
	// Ovh Subsidiary
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	// Product Plan to order
	Plan []PlanObservation `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	// Product Plan to order
	PlanOption []PlanOptionObservation `json:"planOption,omitempty" tf:"plan_option,omitempty"`

	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type ZoneParameters struct {

	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at /1.0/me.json under
	// Ovh Subsidiary
	// +kubebuilder:validation:Optional
	OvhSubsidiary *string `json:"ovhSubsidiary,omitempty" tf:"ovh_subsidiary,omitempty"`

	// Ovh payment mode
	// +kubebuilder:validation:Optional
	PaymentMean *string `json:"paymentMean,omitempty" tf:"payment_mean,omitempty"`

	// Product Plan to order
	// Product Plan to order
	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Product Plan to order
	// Product Plan to order
	// +kubebuilder:validation:Optional
	PlanOption []PlanOptionParameters `json:"planOption,omitempty" tf:"plan_option,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
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
	InitProvider ZoneInitParameters `json:"initProvider,omitempty"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API. ovh_domain_zone.html.markdownsubcategory : "Domain names"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ovhSubsidiary) || (has(self.initProvider) && has(self.initProvider.ovhSubsidiary))",message="spec.forProvider.ovhSubsidiary is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
