/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MGCClusterSpec defines the desired state of MGCCluster
type MGCClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	APIServerFloatingIP string `json:"apiServerFloatingIP,omitempty"`

	// Foo is an example field of MGCCluster. Edit mgccluster_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// MGCClusterStatus defines the observed state of MGCCluster
type MGCClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ready bool `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MGCCluster is the Schema for the mgcclusters API
type MGCCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MGCClusterSpec   `json:"spec,omitempty"`
	Status MGCClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MGCClusterList contains a list of MGCCluster
type MGCClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MGCCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MGCCluster{}, &MGCClusterList{})
}
