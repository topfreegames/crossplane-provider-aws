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
// +kubebuilder:object:generate=true
// +groupName=ec2.aws.wildlife.io
// +versionName=v1alpha1

package manualv1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "ec2.aws.wildlife.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// VPCCIDRBlock type metadata.
var (
	VPCCIDRBlockKind             = reflect.TypeOf(VPCCIDRBlock{}).Name()
	VPCCIDRBlockGroupKind        = schema.GroupKind{Group: Group, Kind: VPCCIDRBlockKind}.String()
	VPCCIDRBlockKindAPIVersion   = VPCCIDRBlockKind + "." + SchemeGroupVersion.String()
	VPCCIDRBlockGroupVersionKind = SchemeGroupVersion.WithKind(VPCCIDRBlockKind)
)

// Instance type metadata.
var (
	InstanceKind             = reflect.TypeOf(Instance{}).Name()
	InstanceGroupKind        = schema.GroupKind{Group: Group, Kind: InstanceKind}.String()
	InstanceKindAPIVersion   = InstanceKind + "." + SchemeGroupVersion.String()
	InstanceGroupVersionKind = SchemeGroupVersion.WithKind(InstanceKind)
)

// VPCPeeringConnection type metadata.
var (
	VPCPeeringConnectionKind             = reflect.TypeOf(VPCPeeringConnection{}).Name()
	VPCPeeringConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: VPCPeeringConnectionKind}.String()
	VPCPeeringConnectionKindAPIVersion   = VPCPeeringConnectionKind + "." + SchemeGroupVersion.String()
	VPCPeeringConnectionGroupVersionKind = SchemeGroupVersion.WithKind(VPCPeeringConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VPCCIDRBlock{}, &VPCCIDRBlockList{})
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
	SchemeBuilder.Register(&VPCPeeringConnection{}, &VPCPeeringConnectionList{})
}
