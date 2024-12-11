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
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/cloud/gkeconnect/gateway/v1beta1/control.proto

package gatewaypb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Operating systems requiring specialized kubeconfigs.
type GenerateCredentialsRequest_OperatingSystem int32

const (
	// Generates a kubeconfig that works for all operating systems not defined
	// below.
	GenerateCredentialsRequest_OPERATING_SYSTEM_UNSPECIFIED GenerateCredentialsRequest_OperatingSystem = 0
	// Generates a kubeconfig that is specifically designed to work with
	// Windows.
	GenerateCredentialsRequest_OPERATING_SYSTEM_WINDOWS GenerateCredentialsRequest_OperatingSystem = 1
)

// Enum value maps for GenerateCredentialsRequest_OperatingSystem.
var (
	GenerateCredentialsRequest_OperatingSystem_name = map[int32]string{
		0: "OPERATING_SYSTEM_UNSPECIFIED",
		1: "OPERATING_SYSTEM_WINDOWS",
	}
	GenerateCredentialsRequest_OperatingSystem_value = map[string]int32{
		"OPERATING_SYSTEM_UNSPECIFIED": 0,
		"OPERATING_SYSTEM_WINDOWS":     1,
	}
)

func (x GenerateCredentialsRequest_OperatingSystem) Enum() *GenerateCredentialsRequest_OperatingSystem {
	p := new(GenerateCredentialsRequest_OperatingSystem)
	*p = x
	return p
}

func (x GenerateCredentialsRequest_OperatingSystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenerateCredentialsRequest_OperatingSystem) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_enumTypes[0].Descriptor()
}

func (GenerateCredentialsRequest_OperatingSystem) Type() protoreflect.EnumType {
	return &file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_enumTypes[0]
}

func (x GenerateCredentialsRequest_OperatingSystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenerateCredentialsRequest_OperatingSystem.Descriptor instead.
func (GenerateCredentialsRequest_OperatingSystem) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescGZIP(), []int{0, 0}
}

// A request for connection information for a particular membership.
type GenerateCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The Fleet membership resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Whether to force the use of Connect Agent-based transport.
	//
	// This will return a configuration that uses Connect Agent as the underlying
	// transport mechanism for cluster types that would otherwise have used a
	// different transport. Requires that Connect Agent be installed on the
	// cluster. Setting this field to false is equivalent to not setting it.
	ForceUseAgent bool `protobuf:"varint,2,opt,name=force_use_agent,json=forceUseAgent,proto3" json:"force_use_agent,omitempty"`
	// Optional. The Connect Gateway version to be used in the resulting
	// configuration.
	//
	// Leave this field blank to let the server choose the version (recommended).
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Optional. The namespace to use in the kubeconfig context.
	//
	// If this field is specified, the server will set the `namespace` field in
	// kubeconfig context. If not specified, the `namespace` field is omitted.
	KubernetesNamespace string `protobuf:"bytes,4,opt,name=kubernetes_namespace,json=kubernetesNamespace,proto3" json:"kubernetes_namespace,omitempty"`
	// Optional. The operating system where the kubeconfig will be used.
	OperatingSystem GenerateCredentialsRequest_OperatingSystem `protobuf:"varint,5,opt,name=operating_system,json=operatingSystem,proto3,enum=google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest_OperatingSystem" json:"operating_system,omitempty"`
}

func (x *GenerateCredentialsRequest) Reset() {
	*x = GenerateCredentialsRequest{}
	mi := &file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCredentialsRequest) ProtoMessage() {}

func (x *GenerateCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCredentialsRequest.ProtoReflect.Descriptor instead.
func (*GenerateCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateCredentialsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenerateCredentialsRequest) GetForceUseAgent() bool {
	if x != nil {
		return x.ForceUseAgent
	}
	return false
}

func (x *GenerateCredentialsRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GenerateCredentialsRequest) GetKubernetesNamespace() string {
	if x != nil {
		return x.KubernetesNamespace
	}
	return ""
}

func (x *GenerateCredentialsRequest) GetOperatingSystem() GenerateCredentialsRequest_OperatingSystem {
	if x != nil {
		return x.OperatingSystem
	}
	return GenerateCredentialsRequest_OPERATING_SYSTEM_UNSPECIFIED
}

// Connection information for a particular membership.
type GenerateCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A full YAML kubeconfig in serialized format.
	Kubeconfig []byte `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	// The generated URI of the cluster as accessed through the Connect Gateway
	// API.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *GenerateCredentialsResponse) Reset() {
	*x = GenerateCredentialsResponse{}
	mi := &file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCredentialsResponse) ProtoMessage() {}

func (x *GenerateCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCredentialsResponse.ProtoReflect.Descriptor instead.
func (*GenerateCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateCredentialsResponse) GetKubeconfig() []byte {
	if x != nil {
		return x.Kubeconfig
	}
	return nil
}

func (x *GenerateCredentialsResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_google_cloud_gkeconnect_gateway_v1beta1_control_proto protoreflect.FileDescriptor

var file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x1a, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x14,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x13, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x51, 0x0a, 0x0f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a,
	0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x01, 0x22, 0x59, 0x0a,
	0x1b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0xd8, 0x02, 0x0a, 0x0e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0xf2, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x12, 0x48, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0x51, 0xca, 0x41, 0x1d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x42, 0x86, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x6b, 0x65, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x70, 0x62,
	0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x70, 0x62, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x47, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescOnce sync.Once
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescData = file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDesc
)

func file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescGZIP() []byte {
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescOnce.Do(func() {
		file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescData)
	})
	return file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDescData
}

var file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_goTypes = []any{
	(GenerateCredentialsRequest_OperatingSystem)(0), // 0: google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest.OperatingSystem
	(*GenerateCredentialsRequest)(nil),              // 1: google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest
	(*GenerateCredentialsResponse)(nil),             // 2: google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsResponse
}
var file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_depIdxs = []int32{
	0, // 0: google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest.operating_system:type_name -> google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest.OperatingSystem
	1, // 1: google.cloud.gkeconnect.gateway.v1beta1.GatewayControl.GenerateCredentials:input_type -> google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsRequest
	2, // 2: google.cloud.gkeconnect.gateway.v1beta1.GatewayControl.GenerateCredentials:output_type -> google.cloud.gkeconnect.gateway.v1beta1.GenerateCredentialsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_init() }
func file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_init() {
	if File_google_cloud_gkeconnect_gateway_v1beta1_control_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_depIdxs,
		EnumInfos:         file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_enumTypes,
		MessageInfos:      file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_msgTypes,
	}.Build()
	File_google_cloud_gkeconnect_gateway_v1beta1_control_proto = out.File
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_rawDesc = nil
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_goTypes = nil
	file_google_cloud_gkeconnect_gateway_v1beta1_control_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayControlClient is the client API for GatewayControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayControlClient interface {
	// GenerateCredentials provides connection information that allows a user to
	// access the specified membership using Connect Gateway.
	GenerateCredentials(ctx context.Context, in *GenerateCredentialsRequest, opts ...grpc.CallOption) (*GenerateCredentialsResponse, error)
}

type gatewayControlClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayControlClient(cc grpc.ClientConnInterface) GatewayControlClient {
	return &gatewayControlClient{cc}
}

func (c *gatewayControlClient) GenerateCredentials(ctx context.Context, in *GenerateCredentialsRequest, opts ...grpc.CallOption) (*GenerateCredentialsResponse, error) {
	out := new(GenerateCredentialsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.gkeconnect.gateway.v1beta1.GatewayControl/GenerateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayControlServer is the server API for GatewayControl service.
type GatewayControlServer interface {
	// GenerateCredentials provides connection information that allows a user to
	// access the specified membership using Connect Gateway.
	GenerateCredentials(context.Context, *GenerateCredentialsRequest) (*GenerateCredentialsResponse, error)
}

// UnimplementedGatewayControlServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayControlServer struct {
}

func (*UnimplementedGatewayControlServer) GenerateCredentials(context.Context, *GenerateCredentialsRequest) (*GenerateCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCredentials not implemented")
}

func RegisterGatewayControlServer(s *grpc.Server, srv GatewayControlServer) {
	s.RegisterService(&_GatewayControl_serviceDesc, srv)
}

func _GatewayControl_GenerateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayControlServer).GenerateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gkeconnect.gateway.v1beta1.GatewayControl/GenerateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayControlServer).GenerateCredentials(ctx, req.(*GenerateCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.gkeconnect.gateway.v1beta1.GatewayControl",
	HandlerType: (*GatewayControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCredentials",
			Handler:    _GatewayControl_GenerateCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/gkeconnect/gateway/v1beta1/control.proto",
}
