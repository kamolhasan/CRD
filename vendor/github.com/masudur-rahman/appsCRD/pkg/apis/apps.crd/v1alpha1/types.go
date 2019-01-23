package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclien:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Deployment
type CustomDeployment struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec CustomDeploymentSpec	`json:"spec,omitempty"`
}

// Deployment Spec
type CustomDeploymentSpec struct {
	Replicas *int32	`json:"replicas,omitempty"`
	Selector *metav1.LabelSelector	`json:"selector"`
	Template CustomPodTemplateSpec	`json:"template"`
}
// Custom Pod Template Spec
type CustomPodTemplateSpec struct {
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec corev1.PodSpec	`json:"spec,omitempty"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Deployment List
type CustomDeploymentList struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ListMeta		`json:"metadata,omitempty"`
	Items 	[]CustomDeployment	`json:"items"`
}
