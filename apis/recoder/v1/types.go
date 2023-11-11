package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Recorder describes a recoder.
type Recorder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecorderSpec `json:"spec"`
	Status StatusSpec
}

// RecorderSpec is the spec for a Recorder resource
type RecorderSpec struct {
	Url    string
	Vendor string
}

type StatusSpec struct {
	Updated string
	Val     string //deleting,deleted,creating,created,pausing,paused
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecorderList is a list of Recorder resources
type RecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Recorder `json:"items"`
}
