/*
Copyright 2021.

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

// AgentMachineTemplateSpec defines the desired state of AgentMachineTemplate
type AgentMachineTemplateSpec struct {
	Template AgentMachineTemplateResource `json:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=agentmachinetemplates,scope=Namespaced,categories=cluster-api,shortName=agentmt
//+kubebuilder:deprecatedversion:warning="v1alpha1 is a deprecated version for AgentMachineTemplate"

// AgentMachineTemplate is the Schema for the agentmachinetemplates API
type AgentMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AgentMachineTemplateSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// AgentMachineTemplateList contains a list of AgentMachineTemplate.
type AgentMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentMachineTemplate{}, &AgentMachineTemplateList{})
}
