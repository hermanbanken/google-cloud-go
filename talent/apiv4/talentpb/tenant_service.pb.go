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
// source: google/cloud/talent/v4/tenant_service.proto

package talentpb

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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Request of the CreateTenant method.
type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the project under which the tenant is created.
	//
	// The format is "projects/{project_id}", for example,
	// "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The tenant to be created.
	Tenant *Tenant `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTenantRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateTenantRequest) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

// Request for getting a tenant by name.
type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the tenant to be retrieved.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenants/bar".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for updating a specified tenant.
type UpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The tenant resource to replace the current resource in the
	// system.
	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Strongly recommended for the best service experience.
	//
	// If [update_mask][google.cloud.talent.v4.UpdateTenantRequest.update_mask] is
	// provided, only the specified fields in
	// [tenant][google.cloud.talent.v4.UpdateTenantRequest.tenant] are updated.
	// Otherwise all the fields are updated.
	//
	// A field mask to specify the tenant fields to be updated. Only
	// top level fields of [Tenant][google.cloud.talent.v4.Tenant] are supported.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTenantRequest) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *UpdateTenantRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request to delete a tenant.
type DeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the tenant to be deleted.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenants/bar".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// List tenants for which the client has ACL visibility.
type ListTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the project under which the tenant is created.
	//
	// The format is "projects/{project_id}", for example,
	// "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of tenants to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTenantsRequest) Reset() {
	*x = ListTenantsRequest{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsRequest) ProtoMessage() {}

func (x *ListTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsRequest.ProtoReflect.Descriptor instead.
func (*ListTenantsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListTenantsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTenantsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTenantsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// The List tenants response object.
type ListTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tenants for the current client.
	Tenants []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ListTenantsResponse) Reset() {
	*x = ListTenantsResponse{}
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsResponse) ProtoMessage() {}

func (x *ListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_tenant_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListTenantsResponse) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *ListTenantsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListTenantsResponse) GetMetadata() *ResponseMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_google_cloud_talent_v4_tenant_service_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4_tenant_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x34, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x8f, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73,
	0x6b, 0x22, 0x4d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x9d, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a,
	0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0xbd, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x34, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xf0, 0x06, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x22, 0x3f, 0xda, 0x41, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x22, 0x1f, 0x2f, 0x76, 0x34, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x12, 0x85, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x34, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x2e, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x34, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa8, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x4b, 0xda, 0x41, 0x12, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x30, 0x3a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x32, 0x26, 0x2f, 0x76,
	0x34, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x83, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x76, 0x34, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x98, 0x01, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x30, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x34, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x73, 0x1a, 0x6c, 0xca, 0x41, 0x13, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x53,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x42, 0x6c, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x34, 0x42, 0x12, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x6c,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x54,
	0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_talent_v4_tenant_service_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4_tenant_service_proto_rawDescData = file_google_cloud_talent_v4_tenant_service_proto_rawDesc
)

func file_google_cloud_talent_v4_tenant_service_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4_tenant_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4_tenant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4_tenant_service_proto_rawDescData)
	})
	return file_google_cloud_talent_v4_tenant_service_proto_rawDescData
}

var file_google_cloud_talent_v4_tenant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_talent_v4_tenant_service_proto_goTypes = []any{
	(*CreateTenantRequest)(nil),   // 0: google.cloud.talent.v4.CreateTenantRequest
	(*GetTenantRequest)(nil),      // 1: google.cloud.talent.v4.GetTenantRequest
	(*UpdateTenantRequest)(nil),   // 2: google.cloud.talent.v4.UpdateTenantRequest
	(*DeleteTenantRequest)(nil),   // 3: google.cloud.talent.v4.DeleteTenantRequest
	(*ListTenantsRequest)(nil),    // 4: google.cloud.talent.v4.ListTenantsRequest
	(*ListTenantsResponse)(nil),   // 5: google.cloud.talent.v4.ListTenantsResponse
	(*Tenant)(nil),                // 6: google.cloud.talent.v4.Tenant
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*ResponseMetadata)(nil),      // 8: google.cloud.talent.v4.ResponseMetadata
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_google_cloud_talent_v4_tenant_service_proto_depIdxs = []int32{
	6,  // 0: google.cloud.talent.v4.CreateTenantRequest.tenant:type_name -> google.cloud.talent.v4.Tenant
	6,  // 1: google.cloud.talent.v4.UpdateTenantRequest.tenant:type_name -> google.cloud.talent.v4.Tenant
	7,  // 2: google.cloud.talent.v4.UpdateTenantRequest.update_mask:type_name -> google.protobuf.FieldMask
	6,  // 3: google.cloud.talent.v4.ListTenantsResponse.tenants:type_name -> google.cloud.talent.v4.Tenant
	8,  // 4: google.cloud.talent.v4.ListTenantsResponse.metadata:type_name -> google.cloud.talent.v4.ResponseMetadata
	0,  // 5: google.cloud.talent.v4.TenantService.CreateTenant:input_type -> google.cloud.talent.v4.CreateTenantRequest
	1,  // 6: google.cloud.talent.v4.TenantService.GetTenant:input_type -> google.cloud.talent.v4.GetTenantRequest
	2,  // 7: google.cloud.talent.v4.TenantService.UpdateTenant:input_type -> google.cloud.talent.v4.UpdateTenantRequest
	3,  // 8: google.cloud.talent.v4.TenantService.DeleteTenant:input_type -> google.cloud.talent.v4.DeleteTenantRequest
	4,  // 9: google.cloud.talent.v4.TenantService.ListTenants:input_type -> google.cloud.talent.v4.ListTenantsRequest
	6,  // 10: google.cloud.talent.v4.TenantService.CreateTenant:output_type -> google.cloud.talent.v4.Tenant
	6,  // 11: google.cloud.talent.v4.TenantService.GetTenant:output_type -> google.cloud.talent.v4.Tenant
	6,  // 12: google.cloud.talent.v4.TenantService.UpdateTenant:output_type -> google.cloud.talent.v4.Tenant
	9,  // 13: google.cloud.talent.v4.TenantService.DeleteTenant:output_type -> google.protobuf.Empty
	5,  // 14: google.cloud.talent.v4.TenantService.ListTenants:output_type -> google.cloud.talent.v4.ListTenantsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4_tenant_service_proto_init() }
func file_google_cloud_talent_v4_tenant_service_proto_init() {
	if File_google_cloud_talent_v4_tenant_service_proto != nil {
		return
	}
	file_google_cloud_talent_v4_common_proto_init()
	file_google_cloud_talent_v4_tenant_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_talent_v4_tenant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_talent_v4_tenant_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4_tenant_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_talent_v4_tenant_service_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4_tenant_service_proto = out.File
	file_google_cloud_talent_v4_tenant_service_proto_rawDesc = nil
	file_google_cloud_talent_v4_tenant_service_proto_goTypes = nil
	file_google_cloud_talent_v4_tenant_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	// Creates a new tenant entity.
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Retrieves specified tenant.
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Updates specified tenant.
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Deletes specified tenant.
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists all tenants associated with the project.
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.TenantService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.TenantService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.TenantService/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.TenantService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.TenantService/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	// Creates a new tenant entity.
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	// Retrieves specified tenant.
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	// Updates specified tenant.
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	// Deletes specified tenant.
	DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error)
	// Lists all tenants associated with the project.
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
}

// UnimplementedTenantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (*UnimplementedTenantServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (*UnimplementedTenantServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (*UnimplementedTenantServiceServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (*UnimplementedTenantServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (*UnimplementedTenantServiceServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.TenantService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.TenantService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.TenantService/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.TenantService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.TenantService/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4/tenant_service.proto",
}
