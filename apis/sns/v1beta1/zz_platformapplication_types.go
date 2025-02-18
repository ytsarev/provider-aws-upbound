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

type PlatformApplicationObservation struct {

	// The ARN of the SNS platform application
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ARN of the SNS platform application
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PlatformApplicationParameters struct {

	// The bundle identifier that's assigned to your iOS app. May only include alphanumeric characters, hyphens (-), and periods (.).
	// +kubebuilder:validation:Optional
	ApplePlatformBundleID *string `json:"applePlatformBundleId,omitempty" tf:"apple_platform_bundle_id,omitempty"`

	// The identifier that's assigned to your Apple developer account team. Must be 10 alphanumeric characters.
	// +kubebuilder:validation:Optional
	ApplePlatformTeamID *string `json:"applePlatformTeamId,omitempty" tf:"apple_platform_team_id,omitempty"`

	// The ARN of the SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	// +kubebuilder:validation:Optional
	EventDeliveryFailureTopicArn *string `json:"eventDeliveryFailureTopicArn,omitempty" tf:"event_delivery_failure_topic_arn,omitempty"`

	// The ARN of the SNS Topic triggered when a new platform endpoint is added to your platform application.
	// +kubebuilder:validation:Optional
	EventEndpointCreatedTopicArn *string `json:"eventEndpointCreatedTopicArn,omitempty" tf:"event_endpoint_created_topic_arn,omitempty"`

	// The ARN of the SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	// +kubebuilder:validation:Optional
	EventEndpointDeletedTopicArn *string `json:"eventEndpointDeletedTopicArn,omitempty" tf:"event_endpoint_deleted_topic_arn,omitempty"`

	// The ARN of the SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	// +kubebuilder:validation:Optional
	EventEndpointUpdatedTopicArn *string `json:"eventEndpointUpdatedTopicArn,omitempty" tf:"event_endpoint_updated_topic_arn,omitempty"`

	// The IAM role ARN permitted to receive failure feedback for this application and give SNS write access to use CloudWatch logs on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	FailureFeedbackRoleArn *string `json:"failureFeedbackRoleArn,omitempty" tf:"failure_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate failureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FailureFeedbackRoleArnRef *v1.Reference `json:"failureFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate failureFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	FailureFeedbackRoleArnSelector *v1.Selector `json:"failureFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The platform that the app is registered with. See Platform for supported platforms.
	// +kubebuilder:validation:Required
	Platform *string `json:"platform" tf:"platform,omitempty"`

	// Application Platform credential. See Credential for type of credential required for platform.
	// +kubebuilder:validation:Required
	PlatformCredentialSecretRef v1.SecretKeySelector `json:"platformCredentialSecretRef" tf:"-"`

	// Application Platform principal. See Principal for type of principal required for platform.
	// +kubebuilder:validation:Optional
	PlatformPrincipalSecretRef *v1.SecretKeySelector `json:"platformPrincipalSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The IAM role ARN permitted to receive success feedback for this application and give SNS write access to use CloudWatch logs on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SuccessFeedbackRoleArn *string `json:"successFeedbackRoleArn,omitempty" tf:"success_feedback_role_arn,omitempty"`

	// Reference to a Role in iam to populate successFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SuccessFeedbackRoleArnRef *v1.Reference `json:"successFeedbackRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate successFeedbackRoleArn.
	// +kubebuilder:validation:Optional
	SuccessFeedbackRoleArnSelector *v1.Selector `json:"successFeedbackRoleArnSelector,omitempty" tf:"-"`

	// The sample rate percentage (0-100) of successfully delivered messages.
	// +kubebuilder:validation:Optional
	SuccessFeedbackSampleRate *string `json:"successFeedbackSampleRate,omitempty" tf:"success_feedback_sample_rate,omitempty"`
}

// PlatformApplicationSpec defines the desired state of PlatformApplication
type PlatformApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlatformApplicationParameters `json:"forProvider"`
}

// PlatformApplicationStatus defines the observed state of PlatformApplication.
type PlatformApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlatformApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PlatformApplication is the Schema for the PlatformApplications API. Provides an SNS platform application resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PlatformApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformApplicationSpec   `json:"spec"`
	Status            PlatformApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlatformApplicationList contains a list of PlatformApplications
type PlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlatformApplication `json:"items"`
}

// Repository type metadata.
var (
	PlatformApplication_Kind             = "PlatformApplication"
	PlatformApplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlatformApplication_Kind}.String()
	PlatformApplication_KindAPIVersion   = PlatformApplication_Kind + "." + CRDGroupVersion.String()
	PlatformApplication_GroupVersionKind = CRDGroupVersion.WithKind(PlatformApplication_Kind)
)

func init() {
	SchemeBuilder.Register(&PlatformApplication{}, &PlatformApplicationList{})
}
