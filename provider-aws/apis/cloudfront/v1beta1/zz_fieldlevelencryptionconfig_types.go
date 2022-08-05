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

type ContentTypeProfileConfigObservation struct {
}

type ContentTypeProfileConfigParameters struct {

	// +kubebuilder:validation:Required
	ContentTypeProfiles []ContentTypeProfilesParameters `json:"contentTypeProfiles" tf:"content_type_profiles,omitempty"`

	// +kubebuilder:validation:Required
	ForwardWhenContentTypeIsUnknown *bool `json:"forwardWhenContentTypeIsUnknown" tf:"forward_when_content_type_is_unknown,omitempty"`
}

type ContentTypeProfilesItemsObservation struct {
}

type ContentTypeProfilesItemsParameters struct {

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`
}

type ContentTypeProfilesObservation struct {
}

type ContentTypeProfilesParameters struct {

	// +kubebuilder:validation:Required
	Items []ContentTypeProfilesItemsParameters `json:"items" tf:"items,omitempty"`
}

type FieldLevelEncryptionConfigObservation struct {
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FieldLevelEncryptionConfigParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	ContentTypeProfileConfig []ContentTypeProfileConfigParameters `json:"contentTypeProfileConfig" tf:"content_type_profile_config,omitempty"`

	// +kubebuilder:validation:Required
	QueryArgProfileConfig []QueryArgProfileConfigParameters `json:"queryArgProfileConfig" tf:"query_arg_profile_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type QueryArgProfileConfigObservation struct {
}

type QueryArgProfileConfigParameters struct {

	// +kubebuilder:validation:Required
	ForwardWhenQueryArgProfileIsUnknown *bool `json:"forwardWhenQueryArgProfileIsUnknown" tf:"forward_when_query_arg_profile_is_unknown,omitempty"`

	// +kubebuilder:validation:Optional
	QueryArgProfiles []QueryArgProfilesParameters `json:"queryArgProfiles,omitempty" tf:"query_arg_profiles,omitempty"`
}

type QueryArgProfilesItemsObservation struct {
}

type QueryArgProfilesItemsParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/cloudfront/v1beta1.FieldLevelEncryptionProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProfileIDRef *v1.Reference `json:"profileIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProfileIDSelector *v1.Selector `json:"profileIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	QueryArg *string `json:"queryArg" tf:"query_arg,omitempty"`
}

type QueryArgProfilesObservation struct {
}

type QueryArgProfilesParameters struct {

	// +kubebuilder:validation:Optional
	Items []QueryArgProfilesItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

// FieldLevelEncryptionConfigSpec defines the desired state of FieldLevelEncryptionConfig
type FieldLevelEncryptionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FieldLevelEncryptionConfigParameters `json:"forProvider"`
}

// FieldLevelEncryptionConfigStatus defines the observed state of FieldLevelEncryptionConfig.
type FieldLevelEncryptionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FieldLevelEncryptionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionConfig is the Schema for the FieldLevelEncryptionConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FieldLevelEncryptionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FieldLevelEncryptionConfigSpec   `json:"spec"`
	Status            FieldLevelEncryptionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionConfigList contains a list of FieldLevelEncryptionConfigs
type FieldLevelEncryptionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FieldLevelEncryptionConfig `json:"items"`
}

// Repository type metadata.
var (
	FieldLevelEncryptionConfig_Kind             = "FieldLevelEncryptionConfig"
	FieldLevelEncryptionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FieldLevelEncryptionConfig_Kind}.String()
	FieldLevelEncryptionConfig_KindAPIVersion   = FieldLevelEncryptionConfig_Kind + "." + CRDGroupVersion.String()
	FieldLevelEncryptionConfig_GroupVersionKind = CRDGroupVersion.WithKind(FieldLevelEncryptionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FieldLevelEncryptionConfig{}, &FieldLevelEncryptionConfigList{})
}
