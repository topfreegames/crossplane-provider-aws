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

// Code generated by ack-generate. DO NOT EDIT.

package manualv1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VPCPeeringConnectionParameters defines the desired state of VPCPeeringConnection
type VPCPeeringConnectionParameters struct {
	// Region is which region the VPCPeeringConnection will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The Amazon Web Services account ID of the owner of the accepter VPC.
	//
	// Default: Your Amazon Web Services account ID
	PeerOwnerID *string `json:"peerOwnerID,omitempty"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a
	// Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string `json:"peerRegion,omitempty"`
	// The tags to assign to the peering connection.
	TagSpecifications                    []*TagSpecification `json:"tagSpecifications,omitempty"`
	CustomVPCPeeringConnectionParameters `json:",inline"`
}

// VPCPeeringConnectionSpec defines the desired state of VPCPeeringConnection
type VPCPeeringConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCPeeringConnectionParameters `json:"forProvider"`
}

// VPCPeeringConnectionObservation defines the observed state of VPCPeeringConnection
type VPCPeeringConnectionObservation struct {
	// Information about the accepter VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	AccepterVPCInfo *VPCPeeringConnectionVPCInfo `json:"accepterVPCInfo,omitempty"`
	// The time that an unaccepted VPC peering connection will expire.
	ExpirationTime *metav1.Time `json:"expirationTime,omitempty"`
	// Information about the requester VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	RequesterVPCInfo *VPCPeeringConnectionVPCInfo `json:"requesterVPCInfo,omitempty"`
	// The status of the VPC peering connection.
	Status *VPCPeeringConnectionStateReason `json:"status,omitempty"`
	// Any tags assigned to the resource.
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the VPC peering connection.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionID,omitempty"`
}

// VPCPeeringConnectionStatus defines the observed state of VPCPeeringConnection.
type VPCPeeringConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCPeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// VPCPeeringConnection is the Schema for the VPCPeeringConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionSpec   `json:"spec"`
	Status            VPCPeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionList contains a list of VPCPeeringConnections
type VPCPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnection `json:"items"`
}

// CustomVPCPeeringConnectionParameters are custom parameters for VPCPeeringConnection
type CustomVPCPeeringConnectionParameters struct {
	// The ID of the requester VPC. You must specify this parameter in the request.
	// +crossplane:generate:reference:type=github.com/crossplane/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcID,omitempty"`
	// VPCIDRef is a reference to an API used to set
	// the VPCID.
	// +optional
	VPCIDRef *xpv1.Reference `json:"vpcIDRef,omitempty"`
	// VPCIDSelector selects references to API used
	// to set the VPCID.
	// +optional
	VPCIDSelector *xpv1.Selector `json:"vpcIDSelector,omitempty"`
	// The ID of the VPC with which you are creating the VPC peering connection.
	// You must specify this parameter in the request.
	PeerVPCID *string `json:"peerVPCID,omitempty"`
	// PeerVPCIDRef is a reference to an API used to set
	// the PeerVPCID.
	// +optional
	PeerVPCIDRef *xpv1.Reference `json:"peerVPCIDRef,omitempty"`
	// PeerVPCIDSelector selects references to API used
	// to set the PeerVPCID.
	// +optional
	PeerVPCIDSelector *xpv1.Selector `json:"peerVPCIDSelector,omitempty"`
	// Automatically accepts the peering connection. If this is not set, the peering connection
	// will be created, but will be in pending-acceptance state. This will only lead to an active
	// connection if both VPCs are in the same tenant.
	AcceptRequest bool `json:"acceptRequest,omitempty"`
}

type VPCPeeringConnectionOptionsDescription struct {
	AllowDNSResolutionFromRemoteVPC *bool `json:"allowDNSResolutionFromRemoteVPC,omitempty"`

	AllowEgressFromLocalClassicLinkToRemoteVPC *bool `json:"allowEgressFromLocalClassicLinkToRemoteVPC,omitempty"`

	AllowEgressFromLocalVPCToRemoteClassicLink *bool `json:"allowEgressFromLocalVPCToRemoteClassicLink,omitempty"`
}

type VPCPeeringConnectionStateReason struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`
}

type VPCPeeringConnectionVPCInfo struct {
	CIDRBlock *string `json:"cidrBlock,omitempty"`

	CIDRBlockSet []*CIDRBlock `json:"cidrBlockSet,omitempty"`

	IPv6CIDRBlockSet []*IPv6CIDRBlock `json:"ipv6CIDRBlockSet,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`
	// Describes the VPC peering connection options.
	PeeringOptions *VPCPeeringConnectionOptionsDescription `json:"peeringOptions,omitempty"`

	Region *string `json:"region,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

type VPCPeeringConnection_SDK struct {
	// Describes a VPC in a VPC peering connection.
	AccepterVPCInfo *VPCPeeringConnectionVPCInfo `json:"accepterVPCInfo,omitempty"`

	ExpirationTime *metav1.Time `json:"expirationTime,omitempty"`
	// Describes a VPC in a VPC peering connection.
	RequesterVPCInfo *VPCPeeringConnectionVPCInfo `json:"requesterVPCInfo,omitempty"`
	// Describes the status of a VPC peering connection.
	Status *VPCPeeringConnectionStateReason `json:"status,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionID,omitempty"`
}

type IPv6CIDRBlock struct {
	IPv6CIDRBlock *string `json:"ipv6CIDRBlock,omitempty"`
}

type CIDRBlock struct {
	CIDRBlock *string `json:"cidrBlock,omitempty"`
}

type VPCPeeringConnectionStateReasonCode string

const (
	VPCPeeringConnectionStateReasonCode_initiating_request VPCPeeringConnectionStateReasonCode = "initiating-request"
	VPCPeeringConnectionStateReasonCode_pending_acceptance VPCPeeringConnectionStateReasonCode = "pending-acceptance"
	VPCPeeringConnectionStateReasonCode_active             VPCPeeringConnectionStateReasonCode = "active"
	VPCPeeringConnectionStateReasonCode_deleted            VPCPeeringConnectionStateReasonCode = "deleted"
	VPCPeeringConnectionStateReasonCode_rejected           VPCPeeringConnectionStateReasonCode = "rejected"
	VPCPeeringConnectionStateReasonCode_failed             VPCPeeringConnectionStateReasonCode = "failed"
	VPCPeeringConnectionStateReasonCode_expired            VPCPeeringConnectionStateReasonCode = "expired"
	VPCPeeringConnectionStateReasonCode_provisioning       VPCPeeringConnectionStateReasonCode = "provisioning"
	VPCPeeringConnectionStateReasonCode_deleting           VPCPeeringConnectionStateReasonCode = "deleting"
)
