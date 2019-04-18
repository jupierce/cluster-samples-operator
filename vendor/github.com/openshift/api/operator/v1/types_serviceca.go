package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCA provides information to configure an operator to manage the service cert controllers
type ServiceCA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// +required
	//spec holds user settable values for configuration
	Spec ServiceCASpec `json:"spec"`
	// +optional
	// status holds observed values from the cluster. They may not be overridden.
	Status ServiceCAStatus `json:"status"`
}

type ServiceCASpec struct {
	OperatorSpec `json:",inline"`
}

type ServiceCAStatus struct {
	OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCAList is a collection of items
type ServiceCAList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata"`
	// Items contains the items
	Items []ServiceCA `json:"items"`
}
