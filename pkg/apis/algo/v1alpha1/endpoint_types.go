package v1alpha1

import (
	"algo-runner-go/swagger"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EndpointSpec defines the desired state of Endpoint
// +k8s:openapi-gen=true
type EndpointSpec struct {
	EndpointOwnerUserName string `json:"endpointOwnerUserName,omitempty"`

	EndpointName string `json:"endpointName,omitempty"`

	PipelineOwnerUserName string `json:"pipelineOwnerUserName,omitempty"`

	PipelineName string `json:"pipelineName,omitempty"`

	PipelineVersionTag string `json:"pipelineVersionTag,omitempty"`

	AlgoConfigs []swagger.AlgoConfig `json:"algoConfigs,omitempty"`

	TopicConfigs []swagger.TopicConfigModel `json:"topicConfigs,omitempty"`

	Pipes []swagger.PipeModel `json:"pipes,omitempty"`
}

// EndpointStatus defines the observed state of Endpoint
// +k8s:openapi-gen=true
type EndpointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Endpoint is the Schema for the endpoints API
// +k8s:openapi-gen=true
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EndpointSpec   `json:"spec,omitempty"`
	Status EndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EndpointList contains a list of Endpoint
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
