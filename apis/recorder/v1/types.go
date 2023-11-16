package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Recorder describes a recorder.
type Recorder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec         RecorderSpec     `json:"spec"`
	Cmd          CmdSpec          //主动下发的
	RunningState RunningStateSpec //主动上报的
}

// RecorderSpec is the spec for a Recorder resource
type RecorderSpec struct {
	Url      string
	Vendor   string
	Revision int64
}

type CmdSpec struct {
	Created  metav1.Time
	Updated  metav1.Time
	Val      string //deleting,deleted,creating,created,pausing,paused
	Feedback string //ok,otherwise error string
}

type RunningStateSpec struct {
	Updated metav1.Time
	Val     string //work, pause, rest
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecorderList is a list of Recorder resources
type RecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Recorder `json:"items"`
}
