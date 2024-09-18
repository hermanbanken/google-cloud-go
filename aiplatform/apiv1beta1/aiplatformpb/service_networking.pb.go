// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/aiplatform/v1beta1/service_networking.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PSC config that is used to automatically create forwarding rule via
// ServiceConnectionMap.
type PSCAutomationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Project id used to create forwarding rule.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Required. The full name of the Google Compute Engine
	// [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks).
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// `projects/{project}/global/networks/{network}`.
	// Where {project} is a project number, as in '12345', and {network} is
	// network name.
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *PSCAutomationConfig) Reset() {
	*x = PSCAutomationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSCAutomationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSCAutomationConfig) ProtoMessage() {}

func (x *PSCAutomationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSCAutomationConfig.ProtoReflect.Descriptor instead.
func (*PSCAutomationConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescGZIP(), []int{0}
}

func (x *PSCAutomationConfig) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PSCAutomationConfig) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

// Represents configuration for private service connect.
type PrivateServiceConnectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. If true, expose the IndexEndpoint via private service connect.
	EnablePrivateServiceConnect bool `protobuf:"varint,1,opt,name=enable_private_service_connect,json=enablePrivateServiceConnect,proto3" json:"enable_private_service_connect,omitempty"`
	// A list of Projects from which the forwarding rule will target the service
	// attachment.
	ProjectAllowlist []string `protobuf:"bytes,2,rep,name=project_allowlist,json=projectAllowlist,proto3" json:"project_allowlist,omitempty"`
	// Output only. The name of the generated service attachment resource.
	// This is only populated if the endpoint is deployed with
	// PrivateServiceConnect.
	ServiceAttachment string `protobuf:"bytes,5,opt,name=service_attachment,json=serviceAttachment,proto3" json:"service_attachment,omitempty"`
}

func (x *PrivateServiceConnectConfig) Reset() {
	*x = PrivateServiceConnectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateServiceConnectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateServiceConnectConfig) ProtoMessage() {}

func (x *PrivateServiceConnectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateServiceConnectConfig.ProtoReflect.Descriptor instead.
func (*PrivateServiceConnectConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateServiceConnectConfig) GetEnablePrivateServiceConnect() bool {
	if x != nil {
		return x.EnablePrivateServiceConnect
	}
	return false
}

func (x *PrivateServiceConnectConfig) GetProjectAllowlist() []string {
	if x != nil {
		return x.ProjectAllowlist
	}
	return nil
}

func (x *PrivateServiceConnectConfig) GetServiceAttachment() string {
	if x != nil {
		return x.ServiceAttachment
	}
	return ""
}

// PscAutomatedEndpoints defines the output of the forwarding rule
// automatically created by each PscAutomationConfig.
type PscAutomatedEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Corresponding project_id in pscAutomationConfigs
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Corresponding network in pscAutomationConfigs.
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// Ip Address created by the automated forwarding rule.
	MatchAddress string `protobuf:"bytes,3,opt,name=match_address,json=matchAddress,proto3" json:"match_address,omitempty"`
}

func (x *PscAutomatedEndpoints) Reset() {
	*x = PscAutomatedEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PscAutomatedEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PscAutomatedEndpoints) ProtoMessage() {}

func (x *PscAutomatedEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PscAutomatedEndpoints.ProtoReflect.Descriptor instead.
func (*PscAutomatedEndpoints) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescGZIP(), []int{2}
}

func (x *PscAutomatedEndpoints) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PscAutomatedEndpoints) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *PscAutomatedEndpoints) GetMatchAddress() string {
	if x != nil {
		return x.MatchAddress
	}
	return ""
}

// Configuration for PSC-I.
type PscInterfaceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The full name of the Compute Engine
	// [network
	// attachment](https://cloud.google.com/vpc/docs/about-network-attachments) to
	// attach to the resource.
	// For example, `projects/12345/regions/us-central1/networkAttachments/myNA`.
	// is of the form
	// `projects/{project}/regions/{region}/networkAttachments/{networkAttachment}`.
	// Where {project} is a project number, as in `12345`, and {networkAttachment}
	// is a network attachment name.
	// To specify this field, you must have already [created a network attachment]
	// (https://cloud.google.com/vpc/docs/create-manage-network-attachments#create-network-attachments).
	// This field is only used for resources using PSC-I.
	NetworkAttachment string `protobuf:"bytes,1,opt,name=network_attachment,json=networkAttachment,proto3" json:"network_attachment,omitempty"`
}

func (x *PscInterfaceConfig) Reset() {
	*x = PscInterfaceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PscInterfaceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PscInterfaceConfig) ProtoMessage() {}

func (x *PscInterfaceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PscInterfaceConfig.ProtoReflect.Descriptor instead.
func (*PscInterfaceConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescGZIP(), []int{3}
}

func (x *PscInterfaceConfig) GetNetworkAttachment() string {
	if x != nil {
		return x.NetworkAttachment
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_service_networking_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x13, 0x50, 0x53, 0x43, 0x41, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x22, 0xc8, 0x01, 0x0a, 0x1b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x48, 0x0a, 0x1e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x1b,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x75, 0x0a, 0x15,
	0x50, 0x73, 0x63, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x75, 0x0a, 0x12, 0x50, 0x73, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5f, 0x0a, 0x12, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x01, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xe6, 0x02, 0xea, 0x41, 0x76,
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x16, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_service_networking_proto_goTypes = []any{
	(*PSCAutomationConfig)(nil),         // 0: google.cloud.aiplatform.v1beta1.PSCAutomationConfig
	(*PrivateServiceConnectConfig)(nil), // 1: google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig
	(*PscAutomatedEndpoints)(nil),       // 2: google.cloud.aiplatform.v1beta1.PscAutomatedEndpoints
	(*PscInterfaceConfig)(nil),          // 3: google.cloud.aiplatform.v1beta1.PscInterfaceConfig
}
var file_google_cloud_aiplatform_v1beta1_service_networking_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_service_networking_proto_init() }
func file_google_cloud_aiplatform_v1beta1_service_networking_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_service_networking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PSCAutomationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateServiceConnectConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PscAutomatedEndpoints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PscInterfaceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_service_networking_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_service_networking_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_service_networking_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_service_networking_proto = out.File
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_service_networking_proto_depIdxs = nil
}
