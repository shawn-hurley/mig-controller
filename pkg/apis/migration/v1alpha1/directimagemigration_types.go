/*
Copyright 2020 Red Hat Inc.

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
	"strings"

	imagev1 "github.com/openshift/api/image/v1"
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// DirectImageMigrationSpec defines the desired state of DirectImageMigration
type DirectImageMigrationSpec struct {
	SrcMigClusterRef  *kapi.ObjectReference `json:"srcMigClusterRef,omitempty"`
	DestMigClusterRef *kapi.ObjectReference `json:"destMigClusterRef,omitempty"`
	Namespaces        []string              `json:"namespaces,omitempty"`
}

// DirectImageMigrationStatus defines the observed state of DirectImageMigration
type DirectImageMigrationStatus struct {
	ImageStreams []ImageStreamListItem `json:"imageStreams,omitempty"`
	Conditions
	ObservedDigest string       `json:"observedDigest,omitempty"`
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`
	Phase          string       `json:"phase,omitempty"`
	Itinerary      string       `json:"itinerary,omitempty"`
	Errors         []string     `json:"errors,omitempty"`
}

type ImageStreamListItem struct {
	Namespace     string `json:"namespace,omitempty"`
	Name          string `json:"name,omitempty"`
	DestNamespace string `json:"destNamespace,omitempty"`
	NotFound      bool   `json:"notFound,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DirectImageMigration is the Schema for the directimagemigrations API
// +kubebuilder:resource:path=directimagemigrations,shortName=dim
// +k8s:openapi-gen=true
type DirectImageMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectImageMigrationSpec   `json:"spec,omitempty"`
	Status DirectImageMigrationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DirectImageMigrationList contains a list of DirectImageMigration
type DirectImageMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectImageMigration `json:"items"`
}

func (r *DirectImageMigration) GetSourceCluster(client k8sclient.Client) (*MigCluster, error) {
	return GetCluster(client, r.Spec.SrcMigClusterRef)
}

func (r *DirectImageMigration) GetDestinationCluster(client k8sclient.Client) (*MigCluster, error) {
	return GetCluster(client, r.Spec.DestMigClusterRef)
}

// GetSourceNamespaces get source namespaces without mapping
func (r *DirectImageMigration) GetSourceNamespaces() []string {
	includedNamespaces := []string{}
	for _, namespace := range r.Spec.Namespaces {
		namespace = strings.Split(namespace, ":")[0]
		includedNamespaces = append(includedNamespaces, namespace)
	}

	return includedNamespaces
}

// GetDestinationNamespaces get destination namespaces without mapping
func (r *DirectImageMigration) GetDestinationNamespaces() []string {
	includedNamespaces := []string{}
	for _, namespace := range r.Spec.Namespaces {
		namespaces := strings.Split(namespace, ":")
		if len(namespaces) > 1 {
			includedNamespaces = append(includedNamespaces, namespaces[1])
		} else {
			includedNamespaces = append(includedNamespaces, namespaces[0])
		}
	}

	return includedNamespaces
}

// GetNamespaceMapping gets a map of src to dest namespaces
func (r *DirectImageMigration) GetNamespaceMapping() map[string]string {
	nsMapping := make(map[string]string)
	for _, namespace := range r.Spec.Namespaces {
		namespaces := strings.Split(namespace, ":")
		if len(namespaces) > 1 {
			nsMapping[namespaces[0]] = namespaces[1]
		} else {
			nsMapping[namespaces[0]] = namespaces[0]
		}
	}

	return nsMapping
}

// Add (de-duplicated) errors.
func (r *DirectImageMigration) AddErrors(errors []string) {
	m := map[string]bool{}
	for _, e := range r.Status.Errors {
		m[e] = true
	}
	for _, error := range errors {
		_, found := m[error]
		if !found {
			r.Status.Errors = append(r.Status.Errors, error)
		}
	}
}

func (r *DirectImageMigration) DirectImageStreamMigrationLabels(is imagev1.ImageStream) map[string]string {
	labels := r.GetCorrelationLabels()
	labels[labelKey(is)] = string(is.UID)
	return labels
}

// HasErrors will notify about error presence on the DirectImageMigration resource
func (r *DirectImageMigration) HasErrors() bool {
	return len(r.Status.Errors) > 0
}

// HasCompleted gets whether a DirectImageMigration has completed and a list of errors, if any
func (r *DirectImageMigration) HasCompleted() (bool, []string) {
	completed := r.Status.Phase == "Completed"
	reasons := r.Status.Errors
	return completed, reasons
}

func init() {
	SchemeBuilder.Register(&DirectImageMigration{}, &DirectImageMigrationList{})
}