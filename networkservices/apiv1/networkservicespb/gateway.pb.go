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
// source: google/cloud/networkservices/v1/gateway.proto

package networkservicespb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of the customer-managed gateway.
// Possible values are:
// * OPEN_MESH
// * SECURE_WEB_GATEWAY
type Gateway_Type int32

const (
	// The type of the customer managed gateway is unspecified.
	Gateway_TYPE_UNSPECIFIED Gateway_Type = 0
	// The type of the customer managed gateway is TrafficDirector Open
	// Mesh.
	Gateway_OPEN_MESH Gateway_Type = 1
	// The type of the customer managed gateway is SecureWebGateway (SWG).
	Gateway_SECURE_WEB_GATEWAY Gateway_Type = 2
)

// Enum value maps for Gateway_Type.
var (
	Gateway_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "OPEN_MESH",
		2: "SECURE_WEB_GATEWAY",
	}
	Gateway_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":   0,
		"OPEN_MESH":          1,
		"SECURE_WEB_GATEWAY": 2,
	}
)

func (x Gateway_Type) Enum() *Gateway_Type {
	p := new(Gateway_Type)
	*p = x
	return p
}

func (x Gateway_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gateway_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_networkservices_v1_gateway_proto_enumTypes[0].Descriptor()
}

func (Gateway_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_networkservices_v1_gateway_proto_enumTypes[0]
}

func (x Gateway_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gateway_Type.Descriptor instead.
func (Gateway_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{0, 0}
}

// Gateway represents the configuration for a proxy, typically a load balancer.
// It captures the ip:port over which the services are exposed by the proxy,
// along with any policy configurations. Routes have reference to to Gateways to
// dictate how requests should be routed by this Gateway.
type Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the Gateway resource. It matches pattern
	// `projects/*/locations/*/gateways/<gateway_name>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Server-defined URL of this resource
	SelfLink string `protobuf:"bytes,13,opt,name=self_link,json=selfLink,proto3" json:"self_link,omitempty"`
	// Output only. The timestamp when the resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when the resource was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Set of label tags associated with the Gateway resource.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. A free-text description of the resource. Max length 1024
	// characters.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Immutable. The type of the customer managed gateway.
	// This field is required. If unspecified, an error is returned.
	Type Gateway_Type `protobuf:"varint,6,opt,name=type,proto3,enum=google.cloud.networkservices.v1.Gateway_Type" json:"type,omitempty"`
	// Required. One or more ports that the Gateway must receive traffic on. The
	// proxy binds to the ports specified. Gateway listen on 0.0.0.0 on the ports
	// specified below.
	Ports []int32 `protobuf:"varint,11,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	// Required. Immutable. Scope determines how configuration across multiple
	// Gateway instances are merged. The configuration for multiple Gateway
	// instances with the same scope will be merged as presented as a single
	// coniguration to the proxy/load balancer.
	//
	// Max length 64 characters.
	// Scope should start with a letter and can only have letters, numbers,
	// hyphens.
	Scope string `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how
	// TLS traffic is terminated. If empty, TLS termination is disabled.
	ServerTlsPolicy string `protobuf:"bytes,9,opt,name=server_tls_policy,json=serverTlsPolicy,proto3" json:"server_tls_policy,omitempty"`
}

func (x *Gateway) Reset() {
	*x = Gateway{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateway) ProtoMessage() {}

func (x *Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateway.ProtoReflect.Descriptor instead.
func (*Gateway) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Gateway) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Gateway) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

func (x *Gateway) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Gateway) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Gateway) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Gateway) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Gateway) GetType() Gateway_Type {
	if x != nil {
		return x.Type
	}
	return Gateway_TYPE_UNSPECIFIED
}

func (x *Gateway) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Gateway) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Gateway) GetServerTlsPolicy() string {
	if x != nil {
		return x.ServerTlsPolicy
	}
	return ""
}

// Request used with the ListGateways method.
type ListGatewaysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The project and location from which the Gateways should be
	// listed, specified in the format `projects/*/locations/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Maximum number of Gateways to return per call.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value returned by the last `ListGatewaysResponse`
	// Indicates that this is a continuation of a prior `ListGateways` call,
	// and that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListGatewaysRequest) Reset() {
	*x = ListGatewaysRequest{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGatewaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGatewaysRequest) ProtoMessage() {}

func (x *ListGatewaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGatewaysRequest.ProtoReflect.Descriptor instead.
func (*ListGatewaysRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *ListGatewaysRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListGatewaysRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGatewaysRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response returned by the ListGateways method.
type ListGatewaysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Gateway resources.
	Gateways []*Gateway `protobuf:"bytes,1,rep,name=gateways,proto3" json:"gateways,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `next_page_token` is included. To get the next set of results, call this
	// method again using the value of `next_page_token` as `page_token`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListGatewaysResponse) Reset() {
	*x = ListGatewaysResponse{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGatewaysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGatewaysResponse) ProtoMessage() {}

func (x *ListGatewaysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGatewaysResponse.ProtoReflect.Descriptor instead.
func (*ListGatewaysResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *ListGatewaysResponse) GetGateways() []*Gateway {
	if x != nil {
		return x.Gateways
	}
	return nil
}

func (x *ListGatewaysResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request used by the GetGateway method.
type GetGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the Gateway to get. Must be in the format
	// `projects/*/locations/*/gateways/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetGatewayRequest) Reset() {
	*x = GetGatewayRequest{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayRequest) ProtoMessage() {}

func (x *GetGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayRequest.ProtoReflect.Descriptor instead.
func (*GetGatewayRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *GetGatewayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request used by the CreateGateway method.
type CreateGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource of the Gateway. Must be in the
	// format `projects/*/locations/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Short name of the Gateway resource to be created.
	GatewayId string `protobuf:"bytes,2,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// Required. Gateway resource to be created.
	Gateway *Gateway `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *CreateGatewayRequest) Reset() {
	*x = CreateGatewayRequest{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGatewayRequest) ProtoMessage() {}

func (x *CreateGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGatewayRequest.ProtoReflect.Descriptor instead.
func (*CreateGatewayRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGatewayRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateGatewayRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *CreateGatewayRequest) GetGateway() *Gateway {
	if x != nil {
		return x.Gateway
	}
	return nil
}

// Request used by the UpdateGateway method.
type UpdateGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Field mask is used to specify the fields to be overwritten in the
	// Gateway resource by the update.
	// The fields specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the mask. If the
	// user does not provide a mask then all fields will be overwritten.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. Updated Gateway resource.
	Gateway *Gateway `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *UpdateGatewayRequest) Reset() {
	*x = UpdateGatewayRequest{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGatewayRequest) ProtoMessage() {}

func (x *UpdateGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGatewayRequest.ProtoReflect.Descriptor instead.
func (*UpdateGatewayRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateGatewayRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateGatewayRequest) GetGateway() *Gateway {
	if x != nil {
		return x.Gateway
	}
	return nil
}

// Request used by the DeleteGateway method.
type DeleteGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the Gateway to delete. Must be in the format
	// `projects/*/locations/*/gateways/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteGatewayRequest) Reset() {
	*x = DeleteGatewayRequest{}
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGatewayRequest) ProtoMessage() {}

func (x *DeleteGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networkservices_v1_gateway_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGatewayRequest.ProtoReflect.Descriptor instead.
func (*DeleteGatewayRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteGatewayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_networkservices_v1_gateway_proto protoreflect.FileDescriptor

var file_google_cloud_networkservices_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdd, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6c, 0x73,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x43, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x4d, 0x45, 0x53, 0x48, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x47, 0x41, 0x54,
	0x45, 0x57, 0x41, 0x59, 0x10, 0x02, 0x3a, 0x67, 0xea, 0x41, 0x64, 0x0a, 0x26, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x7d, 0x22,
	0x99, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x12,
	0x26, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x52, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x57, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x12, 0x26, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0a,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x12, 0x47, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x12, 0x47, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x5a, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xed, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x62,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_networkservices_v1_gateway_proto_rawDescOnce sync.Once
	file_google_cloud_networkservices_v1_gateway_proto_rawDescData = file_google_cloud_networkservices_v1_gateway_proto_rawDesc
)

func file_google_cloud_networkservices_v1_gateway_proto_rawDescGZIP() []byte {
	file_google_cloud_networkservices_v1_gateway_proto_rawDescOnce.Do(func() {
		file_google_cloud_networkservices_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_networkservices_v1_gateway_proto_rawDescData)
	})
	return file_google_cloud_networkservices_v1_gateway_proto_rawDescData
}

var file_google_cloud_networkservices_v1_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_networkservices_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_networkservices_v1_gateway_proto_goTypes = []any{
	(Gateway_Type)(0),             // 0: google.cloud.networkservices.v1.Gateway.Type
	(*Gateway)(nil),               // 1: google.cloud.networkservices.v1.Gateway
	(*ListGatewaysRequest)(nil),   // 2: google.cloud.networkservices.v1.ListGatewaysRequest
	(*ListGatewaysResponse)(nil),  // 3: google.cloud.networkservices.v1.ListGatewaysResponse
	(*GetGatewayRequest)(nil),     // 4: google.cloud.networkservices.v1.GetGatewayRequest
	(*CreateGatewayRequest)(nil),  // 5: google.cloud.networkservices.v1.CreateGatewayRequest
	(*UpdateGatewayRequest)(nil),  // 6: google.cloud.networkservices.v1.UpdateGatewayRequest
	(*DeleteGatewayRequest)(nil),  // 7: google.cloud.networkservices.v1.DeleteGatewayRequest
	nil,                           // 8: google.cloud.networkservices.v1.Gateway.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 10: google.protobuf.FieldMask
}
var file_google_cloud_networkservices_v1_gateway_proto_depIdxs = []int32{
	9,  // 0: google.cloud.networkservices.v1.Gateway.create_time:type_name -> google.protobuf.Timestamp
	9,  // 1: google.cloud.networkservices.v1.Gateway.update_time:type_name -> google.protobuf.Timestamp
	8,  // 2: google.cloud.networkservices.v1.Gateway.labels:type_name -> google.cloud.networkservices.v1.Gateway.LabelsEntry
	0,  // 3: google.cloud.networkservices.v1.Gateway.type:type_name -> google.cloud.networkservices.v1.Gateway.Type
	1,  // 4: google.cloud.networkservices.v1.ListGatewaysResponse.gateways:type_name -> google.cloud.networkservices.v1.Gateway
	1,  // 5: google.cloud.networkservices.v1.CreateGatewayRequest.gateway:type_name -> google.cloud.networkservices.v1.Gateway
	10, // 6: google.cloud.networkservices.v1.UpdateGatewayRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 7: google.cloud.networkservices.v1.UpdateGatewayRequest.gateway:type_name -> google.cloud.networkservices.v1.Gateway
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_networkservices_v1_gateway_proto_init() }
func file_google_cloud_networkservices_v1_gateway_proto_init() {
	if File_google_cloud_networkservices_v1_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_networkservices_v1_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_networkservices_v1_gateway_proto_goTypes,
		DependencyIndexes: file_google_cloud_networkservices_v1_gateway_proto_depIdxs,
		EnumInfos:         file_google_cloud_networkservices_v1_gateway_proto_enumTypes,
		MessageInfos:      file_google_cloud_networkservices_v1_gateway_proto_msgTypes,
	}.Build()
	File_google_cloud_networkservices_v1_gateway_proto = out.File
	file_google_cloud_networkservices_v1_gateway_proto_rawDesc = nil
	file_google_cloud_networkservices_v1_gateway_proto_goTypes = nil
	file_google_cloud_networkservices_v1_gateway_proto_depIdxs = nil
}
