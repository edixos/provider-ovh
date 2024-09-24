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

type IPPoolsInitParameters struct {
}

type IPPoolsObservation struct {
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type IPPoolsParameters struct {
}

type SubnetInitParameters struct {
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type SubnetObservation struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	End *string `json:"end,omitempty" tf:"end,omitempty"`

	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPPools []IPPoolsObservation `json:"ipPools,omitempty" tf:"ip_pools,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type SubnetParameters struct {

	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// +kubebuilder:validation:Optional
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
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
	InitProvider SubnetInitParameters `json:"initProvider,omitempty"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Subnet is the Schema for the Subnets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.end) || (has(self.initProvider) && has(self.initProvider.end))",message="spec.forProvider.end is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.start) || (has(self.initProvider) && has(self.initProvider.start))",message="spec.forProvider.start is a required parameter"
	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
