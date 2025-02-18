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

type CommandObservation struct {
}

type CommandParameters struct {

	// –  The name you assign to this job. It must be unique in your account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Python version being used to execute a Python shell job. Allowed values are 2, 3 or 3.9. Version 3 refers to Python 3.6.
	// +kubebuilder:validation:Optional
	PythonVersion *string `json:"pythonVersion,omitempty" tf:"python_version,omitempty"`

	// Specifies the S3 path to a script that executes a job.
	// +kubebuilder:validation:Required
	ScriptLocation *string `json:"scriptLocation" tf:"script_location,omitempty"`
}

type ExecutionPropertyObservation struct {
}

type ExecutionPropertyParameters struct {

	// The maximum number of concurrent runs allowed for a job. The default is 1.
	// +kubebuilder:validation:Optional
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`
}

type JobObservation struct {

	// Amazon Resource Name (ARN) of Glue Job
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Job name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type JobParameters struct {

	// –  The command of the job. Defined below.
	// +kubebuilder:validation:Required
	Command []CommandParameters `json:"command" tf:"command,omitempty"`

	// –  The list of connections used for this job.
	// +kubebuilder:validation:Optional
	Connections []*string `json:"connections,omitempty" tf:"connections,omitempty"`

	// execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the Calling AWS Glue APIs in Python topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the Special Parameters Used by AWS Glue topic in the developer guide.
	// +kubebuilder:validation:Optional
	DefaultArguments map[string]*string `json:"defaultArguments,omitempty" tf:"default_arguments,omitempty"`

	// –  Description of the job.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: FLEX, STANDARD.
	// +kubebuilder:validation:Optional
	ExecutionClass *string `json:"executionClass,omitempty" tf:"execution_class,omitempty"`

	// –  Execution property of the job. Defined below.
	// +kubebuilder:validation:Optional
	ExecutionProperty []ExecutionPropertyParameters `json:"executionProperty,omitempty" tf:"execution_property,omitempty"`

	// The version of glue to use, for example "1.0". For information about available versions, see the AWS Glue Release Notes.
	// +kubebuilder:validation:Optional
	GlueVersion *string `json:"glueVersion,omitempty" tf:"glue_version,omitempty"`

	// –  The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. Required when pythonshell is set, accept either 0.0625 or 1.0. Use number_of_workers and worker_type arguments instead with glue_version 2.0 and above.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// –  The maximum number of times to retry this job if it fails.
	// +kubebuilder:validation:Optional
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// overridable arguments for this job, specified as name-value pairs.
	// +kubebuilder:validation:Optional
	NonOverridableArguments map[string]*string `json:"nonOverridableArguments,omitempty" tf:"non_overridable_arguments,omitempty"`

	// Notification property of the job. Defined below.
	// +kubebuilder:validation:Optional
	NotificationProperty []NotificationPropertyParameters `json:"notificationProperty,omitempty" tf:"notification_property,omitempty"`

	// The number of workers of a defined workerType that are allocated when a job runs.
	// +kubebuilder:validation:Optional
	NumberOfWorkers *float64 `json:"numberOfWorkers,omitempty" tf:"number_of_workers,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  The ARN of the IAM role associated with this job.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The name of the Security Configuration to be associated with the job.
	// +kubebuilder:validation:Optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The job timeout in minutes. The default is 2880 minutes (48 hours) for glueetl and pythonshell jobs, and null (unlimited) for gluestreaming jobs.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	// +kubebuilder:validation:Optional
	WorkerType *string `json:"workerType,omitempty" tf:"worker_type,omitempty"`
}

type NotificationPropertyObservation struct {
}

type NotificationPropertyParameters struct {

	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	// +kubebuilder:validation:Optional
	NotifyDelayAfter *float64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after,omitempty"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API. Provides an Glue Job resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
