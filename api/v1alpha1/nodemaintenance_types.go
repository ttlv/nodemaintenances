/*
Copyright 2020 The HarmonyCloud authors.

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

// NodeMaintenanceSpec defines the desired state of NodeMaintenance
type NodeMaintenanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NodeName   string         `json:"nodeName"`
	Proxy      Proxy          `json:"proxy"`
	Services   []SpecServices `json:"services"`
	MacAddress string         `json:"macAddress"`
	HostName   string         `json:"hostName"`
}

// NodeMaintenanceStatus defines the observed state of NodeMaintenance
type NodeMaintenanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Services   []StatusServices `json:"services"`
	Conditions []Conditions     `json:"conditions"`
	BindStatus BindStatus       `json:"bindStatus"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=nm

// NodeMaintenance is the Schema for the NodeMaintenances API
type NodeMaintenance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeMaintenanceSpec   `json:"spec,omitempty"`
	Status NodeMaintenanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeMaintenanceList contains a list of NodeMaintenance
type NodeMaintenanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeMaintenance `json:"items"`
}

type Proxy struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

type SpecServices struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	ProxyPort          string `json:"proxyPort"`
	FrpServerIpAddress string `json:"frpServerIpAddress"`
	UniqueID           string `json:"uniqueID"`
}

type StatusServices struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	LastUpdate string `json:"lastUpdate"`
}

type Conditions struct {
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	TimeStamp string `json:"timeStamp"`
}

type BindStatus struct {
	Phase                   string `json:"phase"`
	NodeDeploymentReference string `json:"nodeDeploymentReference"`
	TimeStamp               string `json:"timeStamp"`
}

func init() {
	SchemeBuilder.Register(&NodeMaintenance{}, &NodeMaintenanceList{})
}
