package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//customPod
type CustomPod struct {
	metav1.TypeMeta  `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec CustomPodSpec `json:"spec,omitempty"`
} 

//customPodSpec
type CustomPodSpec struct {
	RestartPolicy string `json:"restartPolicy"`
	Containers []CustomContainer `json:"containers"`

}

//CustomContainer
type CustomContainer struct {
	Name string `json:"name"`
	Image string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
}



// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pod List
type CustomPodList struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ListMeta		`json:"metadata,omitempty"`
	Items 	[]CustomPod	`json:"items"`
}
