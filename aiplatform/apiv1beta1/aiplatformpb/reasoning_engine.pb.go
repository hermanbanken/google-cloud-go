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
// source: google/cloud/aiplatform/v1beta1/reasoning_engine.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ReasoningEngine configurations
type ReasoningEngineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User provided package spec of the ReasoningEngine.
	PackageSpec *ReasoningEngineSpec_PackageSpec `protobuf:"bytes,2,opt,name=package_spec,json=packageSpec,proto3" json:"package_spec,omitempty"`
	// Optional. Declarations for object class methods in OpenAPI specification
	// format.
	ClassMethods []*structpb.Struct `protobuf:"bytes,3,rep,name=class_methods,json=classMethods,proto3" json:"class_methods,omitempty"`
}

func (x *ReasoningEngineSpec) Reset() {
	*x = ReasoningEngineSpec{}
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReasoningEngineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReasoningEngineSpec) ProtoMessage() {}

func (x *ReasoningEngineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReasoningEngineSpec.ProtoReflect.Descriptor instead.
func (*ReasoningEngineSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescGZIP(), []int{0}
}

func (x *ReasoningEngineSpec) GetPackageSpec() *ReasoningEngineSpec_PackageSpec {
	if x != nil {
		return x.PackageSpec
	}
	return nil
}

func (x *ReasoningEngineSpec) GetClassMethods() []*structpb.Struct {
	if x != nil {
		return x.ClassMethods
	}
	return nil
}

// ReasoningEngine provides a customizable runtime for models to determine
// which actions to take and in which order.
type ReasoningEngine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the ReasoningEngine.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the ReasoningEngine.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. The description of the ReasoningEngine.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Configurations of the ReasoningEngine
	Spec *ReasoningEngineSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// Output only. Timestamp when this ReasoningEngine was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this ReasoningEngine was most recently updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag string `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *ReasoningEngine) Reset() {
	*x = ReasoningEngine{}
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReasoningEngine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReasoningEngine) ProtoMessage() {}

func (x *ReasoningEngine) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReasoningEngine.ProtoReflect.Descriptor instead.
func (*ReasoningEngine) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescGZIP(), []int{1}
}

func (x *ReasoningEngine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReasoningEngine) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ReasoningEngine) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReasoningEngine) GetSpec() *ReasoningEngineSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ReasoningEngine) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ReasoningEngine) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ReasoningEngine) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// User provided package spec like pickled object and package requirements.
type ReasoningEngineSpec_PackageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The Cloud Storage URI of the pickled python object.
	PickleObjectGcsUri string `protobuf:"bytes,1,opt,name=pickle_object_gcs_uri,json=pickleObjectGcsUri,proto3" json:"pickle_object_gcs_uri,omitempty"`
	// Optional. The Cloud Storage URI of the dependency files in tar.gz format.
	DependencyFilesGcsUri string `protobuf:"bytes,2,opt,name=dependency_files_gcs_uri,json=dependencyFilesGcsUri,proto3" json:"dependency_files_gcs_uri,omitempty"`
	// Optional. The Cloud Storage URI of the `requirements.txt` file
	RequirementsGcsUri string `protobuf:"bytes,3,opt,name=requirements_gcs_uri,json=requirementsGcsUri,proto3" json:"requirements_gcs_uri,omitempty"`
	// Optional. The Python version. Currently support 3.8, 3.9, 3.10, 3.11.
	// If not specified, default value is 3.10.
	PythonVersion string `protobuf:"bytes,4,opt,name=python_version,json=pythonVersion,proto3" json:"python_version,omitempty"`
}

func (x *ReasoningEngineSpec_PackageSpec) Reset() {
	*x = ReasoningEngineSpec_PackageSpec{}
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReasoningEngineSpec_PackageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReasoningEngineSpec_PackageSpec) ProtoMessage() {}

func (x *ReasoningEngineSpec_PackageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReasoningEngineSpec_PackageSpec.ProtoReflect.Descriptor instead.
func (*ReasoningEngineSpec_PackageSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReasoningEngineSpec_PackageSpec) GetPickleObjectGcsUri() string {
	if x != nil {
		return x.PickleObjectGcsUri
	}
	return ""
}

func (x *ReasoningEngineSpec_PackageSpec) GetDependencyFilesGcsUri() string {
	if x != nil {
		return x.DependencyFilesGcsUri
	}
	return ""
}

func (x *ReasoningEngineSpec_PackageSpec) GetRequirementsGcsUri() string {
	if x != nil {
		return x.RequirementsGcsUri
	}
	return ""
}

func (x *ReasoningEngineSpec_PackageSpec) GetPythonVersion() string {
	if x != nil {
		return x.PythonVersion
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_reasoning_engine_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x03, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x68, 0x0a, 0x0c,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0xe6, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x36, 0x0a, 0x15, 0x70, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x70,
	0x69, 0x63, 0x6b, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x63, 0x73, 0x55, 0x72,
	0x69, 0x12, 0x3c, 0x0a, 0x18, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x15, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x47, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12,
	0x35, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x67, 0x63, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x47, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0d, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x87, 0x04, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x3a, 0x9f, 0x01, 0xea, 0x41, 0x9b,
	0x01, 0x0a, 0x29, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x4b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x7d, 0x2a, 0x10, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x32, 0x0f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0xeb, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_goTypes = []any{
	(*ReasoningEngineSpec)(nil),             // 0: google.cloud.aiplatform.v1beta1.ReasoningEngineSpec
	(*ReasoningEngine)(nil),                 // 1: google.cloud.aiplatform.v1beta1.ReasoningEngine
	(*ReasoningEngineSpec_PackageSpec)(nil), // 2: google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec
	(*structpb.Struct)(nil),                 // 3: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),           // 4: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.package_spec:type_name -> google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.PackageSpec
	3, // 1: google.cloud.aiplatform.v1beta1.ReasoningEngineSpec.class_methods:type_name -> google.protobuf.Struct
	0, // 2: google.cloud.aiplatform.v1beta1.ReasoningEngine.spec:type_name -> google.cloud.aiplatform.v1beta1.ReasoningEngineSpec
	4, // 3: google.cloud.aiplatform.v1beta1.ReasoningEngine.create_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.cloud.aiplatform.v1beta1.ReasoningEngine.update_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_init() }
func file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_reasoning_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_reasoning_engine_proto = out.File
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_reasoning_engine_proto_depIdxs = nil
}
