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

type ConfigObservation struct {
}

type ConfigParameters struct {

	// A list of strings to filter events by PagerDuty event type. "incident.triggered" is required. The follow event types are also possible:
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// Allows you to filter events by priority. Needs to be an array of PagerDuty priority IDs. Available through pagerduty_priority data source.
	// +kubebuilder:validation:Optional
	Priorities []*string `json:"priorities,omitempty" tf:"priorities,omitempty"`

	// Allows you to filter events by urgency. Either high or low.
	// +kubebuilder:validation:Optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`
}

type ConnectionObservation struct {

	// Name of the Slack channel in Slack connection.
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// The ID of the slack connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the source (team or service) in Slack connection.
	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`
}

type ConnectionParameters struct {

	// The ID of a Slack channel in the workspace.
	// +kubebuilder:validation:Required
	ChannelID *string `json:"channelId" tf:"channel_id,omitempty"`

	// Configuration options for the Slack connection that provide options to filter events.
	// +kubebuilder:validation:Required
	Config []ConfigParameters `json:"config" tf:"config,omitempty"`

	// Type of notification. Either responder or stakeholder.
	// +kubebuilder:validation:Required
	NotificationType *string `json:"notificationType" tf:"notification_type,omitempty"`

	// The ID of the source in PagerDuty. Valid sources are services or teams.
	// +kubebuilder:validation:Required
	SourceID *string `json:"sourceId" tf:"source_id,omitempty"`

	// The type of the source. Either team_reference or service_reference.
	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`

	// The ID of the connected Slack workspace. Can also be defined by the SLACK_CONNECTION_WORKSPACE_ID environment variable.
	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Creates and manages a slack connection in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}