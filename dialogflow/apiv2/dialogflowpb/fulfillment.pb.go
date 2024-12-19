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
// source: google/cloud/dialogflow/v2/fulfillment.proto

package dialogflowpb

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
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of the feature.
type Fulfillment_Feature_Type int32

const (
	// Feature type not specified.
	Fulfillment_Feature_TYPE_UNSPECIFIED Fulfillment_Feature_Type = 0
	// Fulfillment is enabled for SmallTalk.
	Fulfillment_Feature_SMALLTALK Fulfillment_Feature_Type = 1
)

// Enum value maps for Fulfillment_Feature_Type.
var (
	Fulfillment_Feature_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "SMALLTALK",
	}
	Fulfillment_Feature_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"SMALLTALK":        1,
	}
)

func (x Fulfillment_Feature_Type) Enum() *Fulfillment_Feature_Type {
	p := new(Fulfillment_Feature_Type)
	*p = x
	return p
}

func (x Fulfillment_Feature_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Fulfillment_Feature_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_enumTypes[0].Descriptor()
}

func (Fulfillment_Feature_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_v2_fulfillment_proto_enumTypes[0]
}

func (x Fulfillment_Feature_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Fulfillment_Feature_Type.Descriptor instead.
func (Fulfillment_Feature_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{0, 1, 0}
}

// By default, your agent responds to a matched intent with a static response.
// As an alternative, you can provide a more dynamic response by using
// fulfillment. When you enable fulfillment for an intent, Dialogflow responds
// to that intent by calling a service that you define. For example, if an
// end-user wants to schedule a haircut on Friday, your service can check your
// database and respond to the end-user with availability information for
// Friday.
//
// For more information, see the [fulfillment
// guide](https://cloud.google.com/dialogflow/docs/fulfillment-overview).
type Fulfillment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique identifier of the fulfillment.
	// Supported formats:
	//
	// - `projects/<Project ID>/agent/fulfillment`
	// - `projects/<Project ID>/locations/<Location ID>/agent/fulfillment`
	//
	// This field is not used for Fulfillment in an Environment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The human-readable name of the fulfillment, unique within the
	// agent.
	//
	// This field is not used for Fulfillment in an Environment.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The fulfillment configuration.
	//
	// Types that are assignable to Fulfillment:
	//
	//	*Fulfillment_GenericWebService_
	Fulfillment isFulfillment_Fulfillment `protobuf_oneof:"fulfillment"`
	// Optional. Whether fulfillment is enabled.
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Optional. The field defines whether the fulfillment is enabled for certain
	// features.
	Features []*Fulfillment_Feature `protobuf:"bytes,5,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *Fulfillment) Reset() {
	*x = Fulfillment{}
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fulfillment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment) ProtoMessage() {}

func (x *Fulfillment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment.ProtoReflect.Descriptor instead.
func (*Fulfillment) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{0}
}

func (x *Fulfillment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fulfillment) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (m *Fulfillment) GetFulfillment() isFulfillment_Fulfillment {
	if m != nil {
		return m.Fulfillment
	}
	return nil
}

func (x *Fulfillment) GetGenericWebService() *Fulfillment_GenericWebService {
	if x, ok := x.GetFulfillment().(*Fulfillment_GenericWebService_); ok {
		return x.GenericWebService
	}
	return nil
}

func (x *Fulfillment) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Fulfillment) GetFeatures() []*Fulfillment_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

type isFulfillment_Fulfillment interface {
	isFulfillment_Fulfillment()
}

type Fulfillment_GenericWebService_ struct {
	// Configuration for a generic web service.
	GenericWebService *Fulfillment_GenericWebService `protobuf:"bytes,3,opt,name=generic_web_service,json=genericWebService,proto3,oneof"`
}

func (*Fulfillment_GenericWebService_) isFulfillment_Fulfillment() {}

// The request message for
// [Fulfillments.GetFulfillment][google.cloud.dialogflow.v2.Fulfillments.GetFulfillment].
type GetFulfillmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the fulfillment.
	// Format: `projects/<Project ID>/agent/fulfillment`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFulfillmentRequest) Reset() {
	*x = GetFulfillmentRequest{}
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFulfillmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFulfillmentRequest) ProtoMessage() {}

func (x *GetFulfillmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFulfillmentRequest.ProtoReflect.Descriptor instead.
func (*GetFulfillmentRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{1}
}

func (x *GetFulfillmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message for
// [Fulfillments.UpdateFulfillment][google.cloud.dialogflow.v2.Fulfillments.UpdateFulfillment].
type UpdateFulfillmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fulfillment to update.
	Fulfillment *Fulfillment `protobuf:"bytes,1,opt,name=fulfillment,proto3" json:"fulfillment,omitempty"`
	// Required. The mask to control which fields get updated. If the mask is not
	// present, all fields will be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateFulfillmentRequest) Reset() {
	*x = UpdateFulfillmentRequest{}
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFulfillmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFulfillmentRequest) ProtoMessage() {}

func (x *UpdateFulfillmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFulfillmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateFulfillmentRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFulfillmentRequest) GetFulfillment() *Fulfillment {
	if x != nil {
		return x.Fulfillment
	}
	return nil
}

func (x *UpdateFulfillmentRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Represents configuration for a generic web service.
// Dialogflow supports two mechanisms for authentications:
//
// - Basic authentication with username and password.
// - Authentication with additional authentication headers.
//
// More information could be found at:
// https://cloud.google.com/dialogflow/docs/fulfillment-configure.
type Fulfillment_GenericWebService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fulfillment URI for receiving POST requests.
	// It must use https protocol.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Optional. The user name for HTTP Basic authentication.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Optional. The password for HTTP Basic authentication.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Optional. The HTTP request headers to send together with fulfillment
	// requests.
	RequestHeaders map[string]string `protobuf:"bytes,4,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Indicates if generic web service is created through Cloud
	// Functions integration. Defaults to false.
	//
	// is_cloud_function is deprecated. Cloud functions can be configured by
	// its uri as a regular web service now.
	//
	// Deprecated: Marked as deprecated in google/cloud/dialogflow/v2/fulfillment.proto.
	IsCloudFunction bool `protobuf:"varint,5,opt,name=is_cloud_function,json=isCloudFunction,proto3" json:"is_cloud_function,omitempty"`
}

func (x *Fulfillment_GenericWebService) Reset() {
	*x = Fulfillment_GenericWebService{}
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fulfillment_GenericWebService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_GenericWebService) ProtoMessage() {}

func (x *Fulfillment_GenericWebService) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_GenericWebService.ProtoReflect.Descriptor instead.
func (*Fulfillment_GenericWebService) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Fulfillment_GenericWebService) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Fulfillment_GenericWebService) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Fulfillment_GenericWebService) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Fulfillment_GenericWebService) GetRequestHeaders() map[string]string {
	if x != nil {
		return x.RequestHeaders
	}
	return nil
}

// Deprecated: Marked as deprecated in google/cloud/dialogflow/v2/fulfillment.proto.
func (x *Fulfillment_GenericWebService) GetIsCloudFunction() bool {
	if x != nil {
		return x.IsCloudFunction
	}
	return false
}

// Whether fulfillment is enabled for the specific feature.
type Fulfillment_Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the feature that enabled for fulfillment.
	Type Fulfillment_Feature_Type `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.dialogflow.v2.Fulfillment_Feature_Type" json:"type,omitempty"`
}

func (x *Fulfillment_Feature) Reset() {
	*x = Fulfillment_Feature{}
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fulfillment_Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_Feature) ProtoMessage() {}

func (x *Fulfillment_Feature) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_Feature.ProtoReflect.Descriptor instead.
func (*Fulfillment_Feature) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Fulfillment_Feature) GetType() Fulfillment_Feature_Type {
	if x != nil {
		return x.Type
	}
	return Fulfillment_Feature_TYPE_UNSPECIFIED
}

var File_google_cloud_dialogflow_v2_fulfillment_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_v2_fulfillment_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x07, 0x0a, 0x0b, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x6b, 0x0a, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x57, 0x65,
	0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c, 0x66,
	0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0xdf,
	0x02, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1f, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x7b, 0x0a,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x11, 0x69, 0x73,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x05, 0xe0, 0x41, 0x01, 0x18, 0x01, 0x52, 0x0f, 0x69, 0x73,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x41, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x80, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x54, 0x41, 0x4c,
	0x4b, 0x10, 0x01, 0x3a, 0x8c, 0x01, 0xea, 0x41, 0x88, 0x01, 0x0a, 0x25, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x24, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66,
	0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27,
	0x0a, 0x25, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x75, 0x6c, 0x66,
	0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xac, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x66, 0x75,
	0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0x91, 0x05, 0x0a,
	0x0c, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xdb, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x6d, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x60, 0x5a, 0x35, 0x12, 0x33, 0x2f,
	0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x7d, 0x12, 0x27, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x12, 0xa8, 0x02, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0xb3, 0x01, 0xda, 0x41, 0x17, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x92, 0x01, 0x3a, 0x0b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x5a, 0x4e, 0x3a, 0x0b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x32, 0x3f, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x7d, 0x32, 0x33, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x1a, 0x78, 0xca, 0x41, 0x19, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x59, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x42, 0x96, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x32, 0x42, 0x10, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x70, 0x62, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x1a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescData = file_google_cloud_dialogflow_v2_fulfillment_proto_rawDesc
)

func file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_v2_fulfillment_proto_rawDescData
}

var file_google_cloud_dialogflow_v2_fulfillment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_dialogflow_v2_fulfillment_proto_goTypes = []any{
	(Fulfillment_Feature_Type)(0),         // 0: google.cloud.dialogflow.v2.Fulfillment.Feature.Type
	(*Fulfillment)(nil),                   // 1: google.cloud.dialogflow.v2.Fulfillment
	(*GetFulfillmentRequest)(nil),         // 2: google.cloud.dialogflow.v2.GetFulfillmentRequest
	(*UpdateFulfillmentRequest)(nil),      // 3: google.cloud.dialogflow.v2.UpdateFulfillmentRequest
	(*Fulfillment_GenericWebService)(nil), // 4: google.cloud.dialogflow.v2.Fulfillment.GenericWebService
	(*Fulfillment_Feature)(nil),           // 5: google.cloud.dialogflow.v2.Fulfillment.Feature
	nil,                                   // 6: google.cloud.dialogflow.v2.Fulfillment.GenericWebService.RequestHeadersEntry
	(*fieldmaskpb.FieldMask)(nil),         // 7: google.protobuf.FieldMask
}
var file_google_cloud_dialogflow_v2_fulfillment_proto_depIdxs = []int32{
	4, // 0: google.cloud.dialogflow.v2.Fulfillment.generic_web_service:type_name -> google.cloud.dialogflow.v2.Fulfillment.GenericWebService
	5, // 1: google.cloud.dialogflow.v2.Fulfillment.features:type_name -> google.cloud.dialogflow.v2.Fulfillment.Feature
	1, // 2: google.cloud.dialogflow.v2.UpdateFulfillmentRequest.fulfillment:type_name -> google.cloud.dialogflow.v2.Fulfillment
	7, // 3: google.cloud.dialogflow.v2.UpdateFulfillmentRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 4: google.cloud.dialogflow.v2.Fulfillment.GenericWebService.request_headers:type_name -> google.cloud.dialogflow.v2.Fulfillment.GenericWebService.RequestHeadersEntry
	0, // 5: google.cloud.dialogflow.v2.Fulfillment.Feature.type:type_name -> google.cloud.dialogflow.v2.Fulfillment.Feature.Type
	2, // 6: google.cloud.dialogflow.v2.Fulfillments.GetFulfillment:input_type -> google.cloud.dialogflow.v2.GetFulfillmentRequest
	3, // 7: google.cloud.dialogflow.v2.Fulfillments.UpdateFulfillment:input_type -> google.cloud.dialogflow.v2.UpdateFulfillmentRequest
	1, // 8: google.cloud.dialogflow.v2.Fulfillments.GetFulfillment:output_type -> google.cloud.dialogflow.v2.Fulfillment
	1, // 9: google.cloud.dialogflow.v2.Fulfillments.UpdateFulfillment:output_type -> google.cloud.dialogflow.v2.Fulfillment
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_v2_fulfillment_proto_init() }
func file_google_cloud_dialogflow_v2_fulfillment_proto_init() {
	if File_google_cloud_dialogflow_v2_fulfillment_proto != nil {
		return
	}
	file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes[0].OneofWrappers = []any{
		(*Fulfillment_GenericWebService_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_v2_fulfillment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dialogflow_v2_fulfillment_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_v2_fulfillment_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_v2_fulfillment_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_v2_fulfillment_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_v2_fulfillment_proto = out.File
	file_google_cloud_dialogflow_v2_fulfillment_proto_rawDesc = nil
	file_google_cloud_dialogflow_v2_fulfillment_proto_goTypes = nil
	file_google_cloud_dialogflow_v2_fulfillment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FulfillmentsClient is the client API for Fulfillments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FulfillmentsClient interface {
	// Retrieves the fulfillment.
	GetFulfillment(ctx context.Context, in *GetFulfillmentRequest, opts ...grpc.CallOption) (*Fulfillment, error)
	// Updates the fulfillment.
	UpdateFulfillment(ctx context.Context, in *UpdateFulfillmentRequest, opts ...grpc.CallOption) (*Fulfillment, error)
}

type fulfillmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewFulfillmentsClient(cc grpc.ClientConnInterface) FulfillmentsClient {
	return &fulfillmentsClient{cc}
}

func (c *fulfillmentsClient) GetFulfillment(ctx context.Context, in *GetFulfillmentRequest, opts ...grpc.CallOption) (*Fulfillment, error) {
	out := new(Fulfillment)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.Fulfillments/GetFulfillment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulfillmentsClient) UpdateFulfillment(ctx context.Context, in *UpdateFulfillmentRequest, opts ...grpc.CallOption) (*Fulfillment, error) {
	out := new(Fulfillment)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.Fulfillments/UpdateFulfillment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulfillmentsServer is the server API for Fulfillments service.
type FulfillmentsServer interface {
	// Retrieves the fulfillment.
	GetFulfillment(context.Context, *GetFulfillmentRequest) (*Fulfillment, error)
	// Updates the fulfillment.
	UpdateFulfillment(context.Context, *UpdateFulfillmentRequest) (*Fulfillment, error)
}

// UnimplementedFulfillmentsServer can be embedded to have forward compatible implementations.
type UnimplementedFulfillmentsServer struct {
}

func (*UnimplementedFulfillmentsServer) GetFulfillment(context.Context, *GetFulfillmentRequest) (*Fulfillment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFulfillment not implemented")
}
func (*UnimplementedFulfillmentsServer) UpdateFulfillment(context.Context, *UpdateFulfillmentRequest) (*Fulfillment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFulfillment not implemented")
}

func RegisterFulfillmentsServer(s *grpc.Server, srv FulfillmentsServer) {
	s.RegisterService(&_Fulfillments_serviceDesc, srv)
}

func _Fulfillments_GetFulfillment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFulfillmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulfillmentsServer).GetFulfillment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.Fulfillments/GetFulfillment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulfillmentsServer).GetFulfillment(ctx, req.(*GetFulfillmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulfillments_UpdateFulfillment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFulfillmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulfillmentsServer).UpdateFulfillment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.Fulfillments/UpdateFulfillment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulfillmentsServer).UpdateFulfillment(ctx, req.(*UpdateFulfillmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fulfillments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2.Fulfillments",
	HandlerType: (*FulfillmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFulfillment",
			Handler:    _Fulfillments_GetFulfillment_Handler,
		},
		{
			MethodName: "UpdateFulfillment",
			Handler:    _Fulfillments_UpdateFulfillment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2/fulfillment.proto",
}
