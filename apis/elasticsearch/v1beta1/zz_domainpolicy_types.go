/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainPolicyParameters struct {

	// IAM policy document specifying the access policies for the domain
	// +kubebuilder:validation:Required
	AccessPolicies *string `json:"accessPolicies" tf:"access_policies,omitempty"`

	// Name of the domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticsearch/v1beta1.Domain
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in elasticsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in elasticsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainPolicySpec defines the desired state of DomainPolicy
type DomainPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainPolicyParameters `json:"forProvider"`
}

// DomainPolicyStatus defines the observed state of DomainPolicy.
type DomainPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicy is the Schema for the DomainPolicys API. Provides an Elasticsearch Domain Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainPolicySpec   `json:"spec"`
	Status            DomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicyList contains a list of DomainPolicys
type DomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	DomainPolicy_Kind             = "DomainPolicy"
	DomainPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainPolicy_Kind}.String()
	DomainPolicy_KindAPIVersion   = DomainPolicy_Kind + "." + CRDGroupVersion.String()
	DomainPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DomainPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainPolicy{}, &DomainPolicyList{})
}
