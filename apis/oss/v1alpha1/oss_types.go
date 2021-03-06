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
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true

// BucketList contains a list of Bucket
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// +kubebuilder:object:root=true

// Bucket is a managed resource that represents an Bucket instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="WARNING",type="string",JSONPath=".status.atProvider.message"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibaba}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	runtimev1.ResourceSpec `json:",inline"`
	BucketParameter        `json:",inline"`

	// Profile is used to extend store business information
	// +kubebuilder:pruning:PreserveUnknownFields
	Profile *runtime.RawExtension `json:"profile,omitempty"`
}

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	runtimev1.ResourceStatus `json:",inline"`
	AtProvider               BucketObservation `json:"atProvider,omitempty"`
}

// BucketParameter is the isolated place to store files
type BucketParameter struct {
	ACL                string `json:"acl,omitempty"`
	StorageClass       string `json:"storageClass,omitempty"`
	DataRedundancyType string `json:"dataRedundancyType,omitempty"`
}

// BucketObservation is the representation of the current state that is observed.
type BucketObservation struct {
	ExtranetEndpoint string `json:"extranetEndpoint,omitempty"`
	IntranetEndpoint string `json:"intranetEndpoint,omitempty"`
	Message          string `json:"message,omitempty"`
}
