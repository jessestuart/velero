/*
Copyright 2018 the Heptio Ark contributors.

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

package v1

import (
	corev1api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodVolumeRestoreSpec is the specification for a PodVolumeRestore.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type PodVolumeRestoreSpec struct {
	// Pod is a reference to the pod containing the volume to be restored.
	Pod corev1api.ObjectReference `json:"pod"`

	// Volume is the name of the volume within the Pod to be restored.
	Volume string `json:"volume"`

	// BackupStorageLocation is the name of the backup storage location
	// where the restic repository is stored.
	BackupStorageLocation string `json:"backupStorageLocation"`

	// RepoIdentifier is the restic repository identifier.
	RepoIdentifier string `json:"repoIdentifier"`

	// SnapshotID is the ID of the volume snapshot to be restored.
	SnapshotID string `json:"snapshotID"`
}

// PodVolumeRestorePhase represents the lifecycle phase of a PodVolumeRestore.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type PodVolumeRestorePhase string

const (
	PodVolumeRestorePhaseNew        PodVolumeRestorePhase = "New"
	PodVolumeRestorePhaseInProgress PodVolumeRestorePhase = "InProgress"
	PodVolumeRestorePhaseCompleted  PodVolumeRestorePhase = "Completed"
	PodVolumeRestorePhaseFailed     PodVolumeRestorePhase = "Failed"
)

// PodVolumeRestoreStatus is the current status of a PodVolumeRestore.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type PodVolumeRestoreStatus struct {
	// Phase is the current state of the PodVolumeRestore.
	Phase PodVolumeRestorePhase `json:"phase"`

	// Message is a message about the pod volume restore's status.
	Message string `json:"message"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type PodVolumeRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   PodVolumeRestoreSpec   `json:"spec"`
	Status PodVolumeRestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodVolumeRestoreList is a list of PodVolumeRestores.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type PodVolumeRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PodVolumeRestore `json:"items"`
}
