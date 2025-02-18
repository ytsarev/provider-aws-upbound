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

type TagObservation struct {

	// Transfer Family resource identifier and key, separated by a comma (,)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TagParameters struct {

	// Tag name.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the Transfer Family resource to tag.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/transfer/v1beta1.Server
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Server in transfer to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Server in transfer to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// Tag value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tag is the Schema for the Tags API. Manages an individual Transfer Family resource tag
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagSpec   `json:"spec"`
	Status            TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}
