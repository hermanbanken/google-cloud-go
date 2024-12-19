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
// source: google/cloud/securityposture/v1/sha_custom_config.proto

package securityposturepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	expr "google.golang.org/genproto/googleapis/type/expr"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines the valid value options for the severity of a finding.
type CustomConfig_Severity int32

const (
	// Unspecified severity.
	CustomConfig_SEVERITY_UNSPECIFIED CustomConfig_Severity = 0
	// Critical severity.
	CustomConfig_CRITICAL CustomConfig_Severity = 1
	// High severity.
	CustomConfig_HIGH CustomConfig_Severity = 2
	// Medium severity.
	CustomConfig_MEDIUM CustomConfig_Severity = 3
	// Low severity.
	CustomConfig_LOW CustomConfig_Severity = 4
)

// Enum value maps for CustomConfig_Severity.
var (
	CustomConfig_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "CRITICAL",
		2: "HIGH",
		3: "MEDIUM",
		4: "LOW",
	}
	CustomConfig_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"CRITICAL":             1,
		"HIGH":                 2,
		"MEDIUM":               3,
		"LOW":                  4,
	}
)

func (x CustomConfig_Severity) Enum() *CustomConfig_Severity {
	p := new(CustomConfig_Severity)
	*p = x
	return p
}

func (x CustomConfig_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomConfig_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_enumTypes[0].Descriptor()
}

func (CustomConfig_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_securityposture_v1_sha_custom_config_proto_enumTypes[0]
}

func (x CustomConfig_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomConfig_Severity.Descriptor instead.
func (CustomConfig_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP(), []int{0, 0}
}

// Defines the properties in a custom module configuration for Security
// Health Analytics. Use the custom module configuration to create custom
// detectors that generate custom findings for resources that you specify.
type CustomConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The CEL expression to evaluate to produce findings. When the
	// expression evaluates to true against a resource, a finding is generated.
	Predicate *expr.Expr `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	// Optional. Custom output properties.
	CustomOutput *CustomConfig_CustomOutputSpec `protobuf:"bytes,2,opt,name=custom_output,json=customOutput,proto3" json:"custom_output,omitempty"`
	// Required. The resource types that the custom module operates on. Each
	// custom module can specify up to 5 resource types.
	ResourceSelector *CustomConfig_ResourceSelector `protobuf:"bytes,3,opt,name=resource_selector,json=resourceSelector,proto3" json:"resource_selector,omitempty"`
	// Required. The severity to assign to findings generated by the module.
	Severity CustomConfig_Severity `protobuf:"varint,4,opt,name=severity,proto3,enum=google.cloud.securityposture.v1.CustomConfig_Severity" json:"severity,omitempty"`
	// Optional. Text that describes the vulnerability or misconfiguration that
	// the custom module detects. This explanation is returned with each finding
	// instance to help investigators understand the detected issue. The text must
	// be enclosed in quotation marks.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. An explanation of the recommended steps that security teams can
	// take to resolve the detected issue. This explanation is returned with each
	// finding generated by this module in the `nextSteps` property of the finding
	// JSON.
	Recommendation string `protobuf:"bytes,6,opt,name=recommendation,proto3" json:"recommendation,omitempty"`
}

func (x *CustomConfig) Reset() {
	*x = CustomConfig{}
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConfig) ProtoMessage() {}

func (x *CustomConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConfig.ProtoReflect.Descriptor instead.
func (*CustomConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP(), []int{0}
}

func (x *CustomConfig) GetPredicate() *expr.Expr {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *CustomConfig) GetCustomOutput() *CustomConfig_CustomOutputSpec {
	if x != nil {
		return x.CustomOutput
	}
	return nil
}

func (x *CustomConfig) GetResourceSelector() *CustomConfig_ResourceSelector {
	if x != nil {
		return x.ResourceSelector
	}
	return nil
}

func (x *CustomConfig) GetSeverity() CustomConfig_Severity {
	if x != nil {
		return x.Severity
	}
	return CustomConfig_SEVERITY_UNSPECIFIED
}

func (x *CustomConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomConfig) GetRecommendation() string {
	if x != nil {
		return x.Recommendation
	}
	return ""
}

// A set of optional name-value pairs that define custom source properties to
// return with each finding that is generated by the custom module. The custom
// source properties that are defined here are included in the finding JSON
// under `sourceProperties`.
type CustomConfig_CustomOutputSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A list of custom output properties to add to the finding.
	Properties []*CustomConfig_CustomOutputSpec_Property `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *CustomConfig_CustomOutputSpec) Reset() {
	*x = CustomConfig_CustomOutputSpec{}
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomConfig_CustomOutputSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConfig_CustomOutputSpec) ProtoMessage() {}

func (x *CustomConfig_CustomOutputSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConfig_CustomOutputSpec.ProtoReflect.Descriptor instead.
func (*CustomConfig_CustomOutputSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomConfig_CustomOutputSpec) GetProperties() []*CustomConfig_CustomOutputSpec_Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

// Resource for selecting resource type.
type CustomConfig_ResourceSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource types to run the detector on.
	ResourceTypes []string `protobuf:"bytes,1,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
}

func (x *CustomConfig_ResourceSelector) Reset() {
	*x = CustomConfig_ResourceSelector{}
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomConfig_ResourceSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConfig_ResourceSelector) ProtoMessage() {}

func (x *CustomConfig_ResourceSelector) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConfig_ResourceSelector.ProtoReflect.Descriptor instead.
func (*CustomConfig_ResourceSelector) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CustomConfig_ResourceSelector) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

// An individual name-value pair that defines a custom source property.
type CustomConfig_CustomOutputSpec_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the property for the custom output.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The CEL expression for the custom output. A resource property
	// can be specified to return the value of the property or a text string
	// enclosed in quotation marks.
	ValueExpression *expr.Expr `protobuf:"bytes,2,opt,name=value_expression,json=valueExpression,proto3" json:"value_expression,omitempty"`
}

func (x *CustomConfig_CustomOutputSpec_Property) Reset() {
	*x = CustomConfig_CustomOutputSpec_Property{}
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomConfig_CustomOutputSpec_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConfig_CustomOutputSpec_Property) ProtoMessage() {}

func (x *CustomConfig_CustomOutputSpec_Property) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConfig_CustomOutputSpec_Property.ProtoReflect.Descriptor instead.
func (*CustomConfig_CustomOutputSpec_Property) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CustomConfig_CustomOutputSpec_Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomConfig_CustomOutputSpec_Property) GetValueExpression() *expr.Expr {
	if x != nil {
		return x.ValueExpression
	}
	return nil
}

var File_google_cloud_securityposture_v1_sha_custom_config_proto protoreflect.FileDescriptor

var file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x61, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x06, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x70, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x57, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0xe8, 0x01, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x6c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x66, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3e,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x51,
	0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45,
	0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10,
	0x04, 0x42, 0x8c, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70,
	0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x53, 0x68, 0x61, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70,
	0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x70, 0x62, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescOnce sync.Once
	file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescData = file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDesc
)

func file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescGZIP() []byte {
	file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescData)
	})
	return file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDescData
}

var file_google_cloud_securityposture_v1_sha_custom_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_securityposture_v1_sha_custom_config_proto_goTypes = []any{
	(CustomConfig_Severity)(0),                     // 0: google.cloud.securityposture.v1.CustomConfig.Severity
	(*CustomConfig)(nil),                           // 1: google.cloud.securityposture.v1.CustomConfig
	(*CustomConfig_CustomOutputSpec)(nil),          // 2: google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec
	(*CustomConfig_ResourceSelector)(nil),          // 3: google.cloud.securityposture.v1.CustomConfig.ResourceSelector
	(*CustomConfig_CustomOutputSpec_Property)(nil), // 4: google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec.Property
	(*expr.Expr)(nil),                              // 5: google.type.Expr
}
var file_google_cloud_securityposture_v1_sha_custom_config_proto_depIdxs = []int32{
	5, // 0: google.cloud.securityposture.v1.CustomConfig.predicate:type_name -> google.type.Expr
	2, // 1: google.cloud.securityposture.v1.CustomConfig.custom_output:type_name -> google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec
	3, // 2: google.cloud.securityposture.v1.CustomConfig.resource_selector:type_name -> google.cloud.securityposture.v1.CustomConfig.ResourceSelector
	0, // 3: google.cloud.securityposture.v1.CustomConfig.severity:type_name -> google.cloud.securityposture.v1.CustomConfig.Severity
	4, // 4: google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec.properties:type_name -> google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec.Property
	5, // 5: google.cloud.securityposture.v1.CustomConfig.CustomOutputSpec.Property.value_expression:type_name -> google.type.Expr
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_securityposture_v1_sha_custom_config_proto_init() }
func file_google_cloud_securityposture_v1_sha_custom_config_proto_init() {
	if File_google_cloud_securityposture_v1_sha_custom_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securityposture_v1_sha_custom_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_securityposture_v1_sha_custom_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_securityposture_v1_sha_custom_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_securityposture_v1_sha_custom_config_proto_msgTypes,
	}.Build()
	File_google_cloud_securityposture_v1_sha_custom_config_proto = out.File
	file_google_cloud_securityposture_v1_sha_custom_config_proto_rawDesc = nil
	file_google_cloud_securityposture_v1_sha_custom_config_proto_goTypes = nil
	file_google_cloud_securityposture_v1_sha_custom_config_proto_depIdxs = nil
}
