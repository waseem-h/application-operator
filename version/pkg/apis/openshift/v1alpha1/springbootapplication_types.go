package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpringBootApplicationSpec defines the desired state of SpringBootApplication
// +k8s:openapi-gen=true
type SpringBootApplicationSpec struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:ExclusiveMinimum=false
	Replicas        int32         `json:"replicas,omitempty"`
	ImagePullPolicy string        `json:"imagePullPolicy,omitempty"`
	Image           string        `json:"image"`
	ReloadOnChange  bool          `json:"reloadOnChange,omitempty"`
	RollingParams   RollingParams `json:"rollingParams,omitempty"`
	Service         ServiceSpec   `json:"service,omitempty"`
}

// RollingParams defines the rolling update strategy for the application
type RollingParams struct {
	MaxSurge       string `json:"maxSurge,omitempty"`
	MaxUnavailable string `json:"maxUnavailable,omitempty"`
}

// ServiceSpec defines how the service will be created for the application
type ServiceSpec struct {
	Type     string `json:"type,omitempty"`
	BasePath string `json:"basePath,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Port     string `json:"port,omitempty"`
}

// SpringBootApplicationStatus defines the observed state of SpringBootApplication
// +k8s:openapi-gen=true
type SpringBootApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpringBootApplication is the Schema for the springbootapplications API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=springbootapplications,scope=Namespaced
type SpringBootApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpringBootApplicationSpec   `json:"spec,omitempty"`
	Status SpringBootApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpringBootApplicationList contains a list of SpringBootApplication
type SpringBootApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringBootApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpringBootApplication{}, &SpringBootApplicationList{})
}
