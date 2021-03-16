/*

 Copyright 2021 The Crossplane Authors.

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
	runtimev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// OSSList contains a list of OSS
type OSSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSS `json:"items"`
}

// +kubebuilder:object:root=true

// OSS is a managed resource that represents an OSS instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="WARNING",type="string",JSONPath=".status.atProvider.message"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibaba}
type OSS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSSSpec   `json:"spec,omitempty"`
	Status OSSStatus `json:"status,omitempty"`
}

// OSSSpec defines the desired state of OSS
type OSSSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	ForProvider            OSSParameters `json:"forProvider"`
}

// OSSStatus defines the observed state of OSS
type OSSStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               OSSObservation `json:"atProvider,omitempty"`
}

// Bucket is the isolated place to store files
type Bucket struct {
	Name               string `json:"name"`
	ACL                string `json:"acl,omitempty"`
	StorageClass       string `json:"storageClass,omitempty"`
	DataRedundancyType string `json:"dataRedundancyType,omitempty"`
}

// OSSParameters define the desired state of an OSS
type OSSParameters struct {
	Bucket Bucket `json:"bucket,omitempty"`
}

// OSSObservation is the representation of the current state that is observed.
type OSSObservation struct {
	BucketName       string `json:"bucketName,omitempty"`
	ACL              string `json:"acl,omitempty"`
	StorageClass     string `json:"storageClass,omitempty"`
	RedundancyType   string `json:"redundancyType,omitempty"`
	ExtranetEndpoint string `json:"extranetEndpoint,omitempty"`
	IntranetEndpoint string `json:"intranetEndpoint,omitempty"`
	Message          string `json:"message,omitempty"`
}
