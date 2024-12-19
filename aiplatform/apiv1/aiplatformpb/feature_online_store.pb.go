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
// source: google/cloud/aiplatform/v1/feature_online_store.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible states a featureOnlineStore can have.
type FeatureOnlineStore_State int32

const (
	// Default value. This value is unused.
	FeatureOnlineStore_STATE_UNSPECIFIED FeatureOnlineStore_State = 0
	// State when the featureOnlineStore configuration is not being updated and
	// the fields reflect the current configuration of the featureOnlineStore.
	// The featureOnlineStore is usable in this state.
	FeatureOnlineStore_STABLE FeatureOnlineStore_State = 1
	// The state of the featureOnlineStore configuration when it is being
	// updated. During an update, the fields reflect either the original
	// configuration or the updated configuration of the featureOnlineStore. The
	// featureOnlineStore is still usable in this state.
	FeatureOnlineStore_UPDATING FeatureOnlineStore_State = 2
)

// Enum value maps for FeatureOnlineStore_State.
var (
	FeatureOnlineStore_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STABLE",
		2: "UPDATING",
	}
	FeatureOnlineStore_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STABLE":            1,
		"UPDATING":          2,
	}
)

func (x FeatureOnlineStore_State) Enum() *FeatureOnlineStore_State {
	p := new(FeatureOnlineStore_State)
	*p = x
	return p
}

func (x FeatureOnlineStore_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureOnlineStore_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_enumTypes[0].Descriptor()
}

func (FeatureOnlineStore_State) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_feature_online_store_proto_enumTypes[0]
}

func (x FeatureOnlineStore_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureOnlineStore_State.Descriptor instead.
func (FeatureOnlineStore_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0, 0}
}

// Vertex AI Feature Online Store provides a centralized repository for serving
// ML features and embedding indexes at low latency. The Feature Online Store is
// a top-level container.
type FeatureOnlineStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to StorageType:
	//
	//	*FeatureOnlineStore_Bigtable_
	//	*FeatureOnlineStore_Optimized_
	StorageType isFeatureOnlineStore_StorageType `protobuf_oneof:"storage_type"`
	// Identifier. Name of the FeatureOnlineStore. Format:
	// `projects/{project}/locations/{location}/featureOnlineStores/{featureOnlineStore}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when this FeatureOnlineStore was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this FeatureOnlineStore was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag string `protobuf:"bytes,5,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// FeatureOnlineStore.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// No more than 64 user labels can be associated with one
	// FeatureOnlineStore(System labels are excluded)." System reserved label keys
	// are prefixed with "aiplatform.googleapis.com/" and are immutable.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. State of the featureOnlineStore.
	State FeatureOnlineStore_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.cloud.aiplatform.v1.FeatureOnlineStore_State" json:"state,omitempty"`
	// Optional. The dedicated serving endpoint for this FeatureOnlineStore, which
	// is different from common Vertex service endpoint.
	DedicatedServingEndpoint *FeatureOnlineStore_DedicatedServingEndpoint `protobuf:"bytes,10,opt,name=dedicated_serving_endpoint,json=dedicatedServingEndpoint,proto3" json:"dedicated_serving_endpoint,omitempty"`
	// Optional. Customer-managed encryption key spec for data storage. If set,
	// online store will be secured by this key.
	EncryptionSpec *EncryptionSpec `protobuf:"bytes,13,opt,name=encryption_spec,json=encryptionSpec,proto3" json:"encryption_spec,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzs bool `protobuf:"varint,15,opt,name=satisfies_pzs,json=satisfiesPzs,proto3" json:"satisfies_pzs,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzi bool `protobuf:"varint,16,opt,name=satisfies_pzi,json=satisfiesPzi,proto3" json:"satisfies_pzi,omitempty"`
}

func (x *FeatureOnlineStore) Reset() {
	*x = FeatureOnlineStore{}
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureOnlineStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureOnlineStore) ProtoMessage() {}

func (x *FeatureOnlineStore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureOnlineStore.ProtoReflect.Descriptor instead.
func (*FeatureOnlineStore) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0}
}

func (m *FeatureOnlineStore) GetStorageType() isFeatureOnlineStore_StorageType {
	if m != nil {
		return m.StorageType
	}
	return nil
}

func (x *FeatureOnlineStore) GetBigtable() *FeatureOnlineStore_Bigtable {
	if x, ok := x.GetStorageType().(*FeatureOnlineStore_Bigtable_); ok {
		return x.Bigtable
	}
	return nil
}

func (x *FeatureOnlineStore) GetOptimized() *FeatureOnlineStore_Optimized {
	if x, ok := x.GetStorageType().(*FeatureOnlineStore_Optimized_); ok {
		return x.Optimized
	}
	return nil
}

func (x *FeatureOnlineStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureOnlineStore) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FeatureOnlineStore) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *FeatureOnlineStore) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *FeatureOnlineStore) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *FeatureOnlineStore) GetState() FeatureOnlineStore_State {
	if x != nil {
		return x.State
	}
	return FeatureOnlineStore_STATE_UNSPECIFIED
}

func (x *FeatureOnlineStore) GetDedicatedServingEndpoint() *FeatureOnlineStore_DedicatedServingEndpoint {
	if x != nil {
		return x.DedicatedServingEndpoint
	}
	return nil
}

func (x *FeatureOnlineStore) GetEncryptionSpec() *EncryptionSpec {
	if x != nil {
		return x.EncryptionSpec
	}
	return nil
}

func (x *FeatureOnlineStore) GetSatisfiesPzs() bool {
	if x != nil {
		return x.SatisfiesPzs
	}
	return false
}

func (x *FeatureOnlineStore) GetSatisfiesPzi() bool {
	if x != nil {
		return x.SatisfiesPzi
	}
	return false
}

type isFeatureOnlineStore_StorageType interface {
	isFeatureOnlineStore_StorageType()
}

type FeatureOnlineStore_Bigtable_ struct {
	// Contains settings for the Cloud Bigtable instance that will be created
	// to serve featureValues for all FeatureViews under this
	// FeatureOnlineStore.
	Bigtable *FeatureOnlineStore_Bigtable `protobuf:"bytes,8,opt,name=bigtable,proto3,oneof"`
}

type FeatureOnlineStore_Optimized_ struct {
	// Contains settings for the Optimized store that will be created
	// to serve featureValues for all FeatureViews under this
	// FeatureOnlineStore. When choose Optimized storage type, need to set
	// [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	// to use private endpoint. Otherwise will use public endpoint by default.
	Optimized *FeatureOnlineStore_Optimized `protobuf:"bytes,12,opt,name=optimized,proto3,oneof"`
}

func (*FeatureOnlineStore_Bigtable_) isFeatureOnlineStore_StorageType() {}

func (*FeatureOnlineStore_Optimized_) isFeatureOnlineStore_StorageType() {}

type FeatureOnlineStore_Bigtable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Autoscaling config applied to Bigtable Instance.
	AutoScaling *FeatureOnlineStore_Bigtable_AutoScaling `protobuf:"bytes,1,opt,name=auto_scaling,json=autoScaling,proto3" json:"auto_scaling,omitempty"`
}

func (x *FeatureOnlineStore_Bigtable) Reset() {
	*x = FeatureOnlineStore_Bigtable{}
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureOnlineStore_Bigtable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureOnlineStore_Bigtable) ProtoMessage() {}

func (x *FeatureOnlineStore_Bigtable) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureOnlineStore_Bigtable.ProtoReflect.Descriptor instead.
func (*FeatureOnlineStore_Bigtable) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeatureOnlineStore_Bigtable) GetAutoScaling() *FeatureOnlineStore_Bigtable_AutoScaling {
	if x != nil {
		return x.AutoScaling
	}
	return nil
}

// Optimized storage type
type FeatureOnlineStore_Optimized struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeatureOnlineStore_Optimized) Reset() {
	*x = FeatureOnlineStore_Optimized{}
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureOnlineStore_Optimized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureOnlineStore_Optimized) ProtoMessage() {}

func (x *FeatureOnlineStore_Optimized) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureOnlineStore_Optimized.ProtoReflect.Descriptor instead.
func (*FeatureOnlineStore_Optimized) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0, 1}
}

// The dedicated serving endpoint for this FeatureOnlineStore. Only need to
// set when you choose Optimized storage type. Public endpoint is provisioned
// by default.
type FeatureOnlineStore_DedicatedServingEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. This field will be populated with the domain name to use for
	// this FeatureOnlineStore
	PublicEndpointDomainName string `protobuf:"bytes,2,opt,name=public_endpoint_domain_name,json=publicEndpointDomainName,proto3" json:"public_endpoint_domain_name,omitempty"`
	// Optional. Private service connect config. The private service connection
	// is available only for Optimized storage type, not for embedding
	// management now. If
	// [PrivateServiceConnectConfig.enable_private_service_connect][google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect]
	// set to true, customers will use private service connection to send
	// request. Otherwise, the connection will set to public endpoint.
	PrivateServiceConnectConfig *PrivateServiceConnectConfig `protobuf:"bytes,3,opt,name=private_service_connect_config,json=privateServiceConnectConfig,proto3" json:"private_service_connect_config,omitempty"`
	// Output only. The name of the service attachment resource. Populated if
	// private service connect is enabled and after FeatureViewSync is created.
	ServiceAttachment string `protobuf:"bytes,4,opt,name=service_attachment,json=serviceAttachment,proto3" json:"service_attachment,omitempty"`
}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) Reset() {
	*x = FeatureOnlineStore_DedicatedServingEndpoint{}
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureOnlineStore_DedicatedServingEndpoint) ProtoMessage() {}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureOnlineStore_DedicatedServingEndpoint.ProtoReflect.Descriptor instead.
func (*FeatureOnlineStore_DedicatedServingEndpoint) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0, 2}
}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) GetPublicEndpointDomainName() string {
	if x != nil {
		return x.PublicEndpointDomainName
	}
	return ""
}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) GetPrivateServiceConnectConfig() *PrivateServiceConnectConfig {
	if x != nil {
		return x.PrivateServiceConnectConfig
	}
	return nil
}

func (x *FeatureOnlineStore_DedicatedServingEndpoint) GetServiceAttachment() string {
	if x != nil {
		return x.ServiceAttachment
	}
	return ""
}

type FeatureOnlineStore_Bigtable_AutoScaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The minimum number of nodes to scale down to. Must be greater
	// than or equal to 1.
	MinNodeCount int32 `protobuf:"varint,1,opt,name=min_node_count,json=minNodeCount,proto3" json:"min_node_count,omitempty"`
	// Required. The maximum number of nodes to scale up to. Must be greater
	// than or equal to min_node_count, and less than or equal to 10 times of
	// 'min_node_count'.
	MaxNodeCount int32 `protobuf:"varint,2,opt,name=max_node_count,json=maxNodeCount,proto3" json:"max_node_count,omitempty"`
	// Optional. A percentage of the cluster's CPU capacity. Can be from 10%
	// to 80%. When a cluster's CPU utilization exceeds the target that you
	// have set, Bigtable immediately adds nodes to the cluster. When CPU
	// utilization is substantially lower than the target, Bigtable removes
	// nodes. If not set will default to 50%.
	CpuUtilizationTarget int32 `protobuf:"varint,3,opt,name=cpu_utilization_target,json=cpuUtilizationTarget,proto3" json:"cpu_utilization_target,omitempty"`
}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) Reset() {
	*x = FeatureOnlineStore_Bigtable_AutoScaling{}
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureOnlineStore_Bigtable_AutoScaling) ProtoMessage() {}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureOnlineStore_Bigtable_AutoScaling.ProtoReflect.Descriptor instead.
func (*FeatureOnlineStore_Bigtable_AutoScaling) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) GetMinNodeCount() int32 {
	if x != nil {
		return x.MinNodeCount
	}
	return 0
}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) GetMaxNodeCount() int32 {
	if x != nil {
		return x.MaxNodeCount
	}
	return 0
}

func (x *FeatureOnlineStore_Bigtable_AutoScaling) GetCpuUtilizationTarget() int32 {
	if x != nil {
		return x.CpuUtilizationTarget
	}
	return 0
}

var File_google_cloud_aiplatform_v1_feature_online_store_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x0d, 0x0a, 0x12, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x55,
	0x0a, 0x08, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x57, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4f,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x8a, 0x01, 0x0a, 0x1a, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x18, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x0f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66,
	0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0c, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x50, 0x7a, 0x73,
	0x12, 0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a,
	0x69, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x61,
	0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x50, 0x7a, 0x69, 0x1a, 0x98, 0x02, 0x0a, 0x08, 0x42,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x9e, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x16, 0x63, 0x70,
	0x75, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x14, 0x63, 0x70, 0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x0b, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x1a, 0x96, 0x02, 0x0a, 0x18, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x42, 0x0a, 0x1b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x18, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x1e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x1b, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x3a, 0x86, 0x01, 0xea, 0x41, 0x82, 0x01, 0x0a, 0x2c, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x52, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xd5, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescData = file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_feature_online_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1_feature_online_store_proto_goTypes = []any{
	(FeatureOnlineStore_State)(0),                       // 0: google.cloud.aiplatform.v1.FeatureOnlineStore.State
	(*FeatureOnlineStore)(nil),                          // 1: google.cloud.aiplatform.v1.FeatureOnlineStore
	(*FeatureOnlineStore_Bigtable)(nil),                 // 2: google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable
	(*FeatureOnlineStore_Optimized)(nil),                // 3: google.cloud.aiplatform.v1.FeatureOnlineStore.Optimized
	(*FeatureOnlineStore_DedicatedServingEndpoint)(nil), // 4: google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint
	nil, // 5: google.cloud.aiplatform.v1.FeatureOnlineStore.LabelsEntry
	(*FeatureOnlineStore_Bigtable_AutoScaling)(nil), // 6: google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling
	(*timestamppb.Timestamp)(nil),                   // 7: google.protobuf.Timestamp
	(*EncryptionSpec)(nil),                          // 8: google.cloud.aiplatform.v1.EncryptionSpec
	(*PrivateServiceConnectConfig)(nil),             // 9: google.cloud.aiplatform.v1.PrivateServiceConnectConfig
}
var file_google_cloud_aiplatform_v1_feature_online_store_proto_depIdxs = []int32{
	2,  // 0: google.cloud.aiplatform.v1.FeatureOnlineStore.bigtable:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable
	3,  // 1: google.cloud.aiplatform.v1.FeatureOnlineStore.optimized:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.Optimized
	7,  // 2: google.cloud.aiplatform.v1.FeatureOnlineStore.create_time:type_name -> google.protobuf.Timestamp
	7,  // 3: google.cloud.aiplatform.v1.FeatureOnlineStore.update_time:type_name -> google.protobuf.Timestamp
	5,  // 4: google.cloud.aiplatform.v1.FeatureOnlineStore.labels:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.LabelsEntry
	0,  // 5: google.cloud.aiplatform.v1.FeatureOnlineStore.state:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.State
	4,  // 6: google.cloud.aiplatform.v1.FeatureOnlineStore.dedicated_serving_endpoint:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint
	8,  // 7: google.cloud.aiplatform.v1.FeatureOnlineStore.encryption_spec:type_name -> google.cloud.aiplatform.v1.EncryptionSpec
	6,  // 8: google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.auto_scaling:type_name -> google.cloud.aiplatform.v1.FeatureOnlineStore.Bigtable.AutoScaling
	9,  // 9: google.cloud.aiplatform.v1.FeatureOnlineStore.DedicatedServingEndpoint.private_service_connect_config:type_name -> google.cloud.aiplatform.v1.PrivateServiceConnectConfig
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_feature_online_store_proto_init() }
func file_google_cloud_aiplatform_v1_feature_online_store_proto_init() {
	if File_google_cloud_aiplatform_v1_feature_online_store_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_encryption_spec_proto_init()
	file_google_cloud_aiplatform_v1_service_networking_proto_init()
	file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes[0].OneofWrappers = []any{
		(*FeatureOnlineStore_Bigtable_)(nil),
		(*FeatureOnlineStore_Optimized_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_feature_online_store_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_feature_online_store_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_feature_online_store_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_feature_online_store_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_feature_online_store_proto = out.File
	file_google_cloud_aiplatform_v1_feature_online_store_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_feature_online_store_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_feature_online_store_proto_depIdxs = nil
}
