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

type ModelObservation struct {

	// ID of the model
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ModelParameters struct {

	// Content type of the model
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Description of the model
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the model
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Schema of the model in a JSON form
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelParameters `json:"forProvider"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Model is the Schema for the Models API. Provides a Model for a REST API Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec"`
	Status            ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	Model_Kind             = "Model"
	Model_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Model_Kind}.String()
	Model_KindAPIVersion   = Model_Kind + "." + CRDGroupVersion.String()
	Model_GroupVersionKind = CRDGroupVersion.WithKind(Model_Kind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}
