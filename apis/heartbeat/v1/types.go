package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Heartbeat describes a heartbeat.
type Heartbeat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Revision          int64 `json:",string"`
	UpdateTime        metav1.Time
	Online            bool
	IpsecRingLength   int
	//Spec HeartBeatSpec `json:"spec"`
}

// HeartbeatSpec is the spec for a Heartbeat resource
type HeartbeatSpec struct {
	//NodeName string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HeartbeatList is a list of Heartbeat resources
type HeartbeatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Heartbeat `json:"items"`
}
