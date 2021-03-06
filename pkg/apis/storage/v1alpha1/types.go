package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnterpriseRecyclerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EnterpriseRecycler `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnterpriseRecycler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EnterpriseRecyclerSpec   `json:"spec"`
	Status            EnterpriseRecyclerStatus `json:"status,omitempty"`
}

type EnterpriseRecyclerSpec struct {
	// Fill me
}
type EnterpriseRecyclerStatus struct {
	// Fill me
}
