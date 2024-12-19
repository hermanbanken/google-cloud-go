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
// source: google/cloud/aiplatform/v1beta1/featurestore.proto

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

// Possible states a featurestore can have.
type Featurestore_State int32

const (
	// Default value. This value is unused.
	Featurestore_STATE_UNSPECIFIED Featurestore_State = 0
	// State when the featurestore configuration is not being updated and the
	// fields reflect the current configuration of the featurestore. The
	// featurestore is usable in this state.
	Featurestore_STABLE Featurestore_State = 1
	// The state of the featurestore configuration when it is being updated.
	// During an update, the fields reflect either the original configuration
	// or the updated configuration of the featurestore. For example,
	// `online_serving_config.fixed_node_count` can take minutes to update.
	// While the update is in progress, the featurestore is in the UPDATING
	// state, and the value of `fixed_node_count` can be the original value or
	// the updated value, depending on the progress of the operation. Until the
	// update completes, the actual number of nodes can still be the original
	// value of `fixed_node_count`. The featurestore is still usable in this
	// state.
	Featurestore_UPDATING Featurestore_State = 2
)

// Enum value maps for Featurestore_State.
var (
	Featurestore_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STABLE",
		2: "UPDATING",
	}
	Featurestore_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STABLE":            1,
		"UPDATING":          2,
	}
)

func (x Featurestore_State) Enum() *Featurestore_State {
	p := new(Featurestore_State)
	*p = x
	return p
}

func (x Featurestore_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Featurestore_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_enumTypes[0].Descriptor()
}

func (Featurestore_State) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_featurestore_proto_enumTypes[0]
}

func (x Featurestore_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Featurestore_State.Descriptor instead.
func (Featurestore_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescGZIP(), []int{0, 0}
}

// Vertex AI Feature Store provides a centralized repository for organizing,
// storing, and serving ML features. The Featurestore is a top-level container
// for your features and their values.
type Featurestore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the Featurestore. Format:
	// `projects/{project}/locations/{location}/featurestores/{featurestore}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when this Featurestore was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Featurestore was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag string `protobuf:"bytes,5,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// Featurestore.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// No more than 64 user labels can be associated with one Featurestore(System
	// labels are excluded)."
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Config for online storage resources. The field should not
	// co-exist with the field of `OnlineStoreReplicationConfig`. If both of it
	// and OnlineStoreReplicationConfig are unset, the feature store will not have
	// an online store and cannot be used for online serving.
	OnlineServingConfig *Featurestore_OnlineServingConfig `protobuf:"bytes,7,opt,name=online_serving_config,json=onlineServingConfig,proto3" json:"online_serving_config,omitempty"`
	// Output only. State of the featurestore.
	State Featurestore_State `protobuf:"varint,8,opt,name=state,proto3,enum=google.cloud.aiplatform.v1beta1.Featurestore_State" json:"state,omitempty"`
	// Optional. TTL in days for feature values that will be stored in online
	// serving storage. The Feature Store online storage periodically removes
	// obsolete feature values older than `online_storage_ttl_days` since the
	// feature generation time. Note that `online_storage_ttl_days` should be less
	// than or equal to `offline_storage_ttl_days` for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays int32 `protobuf:"varint,13,opt,name=online_storage_ttl_days,json=onlineStorageTtlDays,proto3" json:"online_storage_ttl_days,omitempty"`
	// Optional. Customer-managed encryption key spec for data storage. If set,
	// both of the online and offline data storage will be secured by this key.
	EncryptionSpec *EncryptionSpec `protobuf:"bytes,10,opt,name=encryption_spec,json=encryptionSpec,proto3" json:"encryption_spec,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzs bool `protobuf:"varint,14,opt,name=satisfies_pzs,json=satisfiesPzs,proto3" json:"satisfies_pzs,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzi bool `protobuf:"varint,15,opt,name=satisfies_pzi,json=satisfiesPzi,proto3" json:"satisfies_pzi,omitempty"`
}

func (x *Featurestore) Reset() {
	*x = Featurestore{}
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Featurestore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Featurestore) ProtoMessage() {}

func (x *Featurestore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Featurestore.ProtoReflect.Descriptor instead.
func (*Featurestore) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescGZIP(), []int{0}
}

func (x *Featurestore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Featurestore) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Featurestore) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Featurestore) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Featurestore) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Featurestore) GetOnlineServingConfig() *Featurestore_OnlineServingConfig {
	if x != nil {
		return x.OnlineServingConfig
	}
	return nil
}

func (x *Featurestore) GetState() Featurestore_State {
	if x != nil {
		return x.State
	}
	return Featurestore_STATE_UNSPECIFIED
}

func (x *Featurestore) GetOnlineStorageTtlDays() int32 {
	if x != nil {
		return x.OnlineStorageTtlDays
	}
	return 0
}

func (x *Featurestore) GetEncryptionSpec() *EncryptionSpec {
	if x != nil {
		return x.EncryptionSpec
	}
	return nil
}

func (x *Featurestore) GetSatisfiesPzs() bool {
	if x != nil {
		return x.SatisfiesPzs
	}
	return false
}

func (x *Featurestore) GetSatisfiesPzi() bool {
	if x != nil {
		return x.SatisfiesPzi
	}
	return false
}

// OnlineServingConfig specifies the details for provisioning online serving
// resources.
type Featurestore_OnlineServingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of nodes for the online store. The number of nodes doesn't
	// scale automatically, but you can manually update the number of
	// nodes. If set to 0, the featurestore will not have an
	// online store and cannot be used for online serving.
	FixedNodeCount int32 `protobuf:"varint,2,opt,name=fixed_node_count,json=fixedNodeCount,proto3" json:"fixed_node_count,omitempty"`
	// Online serving scaling configuration.
	// Only one of `fixed_node_count` and `scaling` can be set. Setting one will
	// reset the other.
	Scaling *Featurestore_OnlineServingConfig_Scaling `protobuf:"bytes,4,opt,name=scaling,proto3" json:"scaling,omitempty"`
}

func (x *Featurestore_OnlineServingConfig) Reset() {
	*x = Featurestore_OnlineServingConfig{}
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Featurestore_OnlineServingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Featurestore_OnlineServingConfig) ProtoMessage() {}

func (x *Featurestore_OnlineServingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Featurestore_OnlineServingConfig.ProtoReflect.Descriptor instead.
func (*Featurestore_OnlineServingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Featurestore_OnlineServingConfig) GetFixedNodeCount() int32 {
	if x != nil {
		return x.FixedNodeCount
	}
	return 0
}

func (x *Featurestore_OnlineServingConfig) GetScaling() *Featurestore_OnlineServingConfig_Scaling {
	if x != nil {
		return x.Scaling
	}
	return nil
}

// Online serving scaling configuration. If min_node_count and
// max_node_count are set to the same value, the cluster will be configured
// with the fixed number of node (no auto-scaling).
type Featurestore_OnlineServingConfig_Scaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The minimum number of nodes to scale down to. Must be greater
	// than or equal to 1.
	MinNodeCount int32 `protobuf:"varint,1,opt,name=min_node_count,json=minNodeCount,proto3" json:"min_node_count,omitempty"`
	// The maximum number of nodes to scale up to. Must be greater than
	// min_node_count, and less than or equal to 10 times of 'min_node_count'.
	MaxNodeCount int32 `protobuf:"varint,2,opt,name=max_node_count,json=maxNodeCount,proto3" json:"max_node_count,omitempty"`
	// Optional. The cpu utilization that the Autoscaler should be trying to
	// achieve. This number is on a scale from 0 (no utilization) to 100
	// (total utilization), and is limited between 10 and 80. When a cluster's
	// CPU utilization exceeds the target that you have set, Bigtable
	// immediately adds nodes to the cluster. When CPU utilization is
	// substantially lower than the target, Bigtable removes nodes. If not set
	// or set to 0, default to 50.
	CpuUtilizationTarget int32 `protobuf:"varint,3,opt,name=cpu_utilization_target,json=cpuUtilizationTarget,proto3" json:"cpu_utilization_target,omitempty"`
}

func (x *Featurestore_OnlineServingConfig_Scaling) Reset() {
	*x = Featurestore_OnlineServingConfig_Scaling{}
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Featurestore_OnlineServingConfig_Scaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Featurestore_OnlineServingConfig_Scaling) ProtoMessage() {}

func (x *Featurestore_OnlineServingConfig_Scaling) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Featurestore_OnlineServingConfig_Scaling.ProtoReflect.Descriptor instead.
func (*Featurestore_OnlineServingConfig_Scaling) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Featurestore_OnlineServingConfig_Scaling) GetMinNodeCount() int32 {
	if x != nil {
		return x.MinNodeCount
	}
	return 0
}

func (x *Featurestore_OnlineServingConfig_Scaling) GetMaxNodeCount() int32 {
	if x != nil {
		return x.MaxNodeCount
	}
	return 0
}

func (x *Featurestore_OnlineServingConfig_Scaling) GetCpuUtilizationTarget() int32 {
	if x != nil {
		return x.CpuUtilizationTarget
	}
	return 0
}

var File_google_cloud_aiplatform_v1beta1_featurestore_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x09, 0x0a, 0x0c, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x12, 0x56, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x7a, 0x0a, 0x15, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x13, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x17, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x74, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x14, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x74, 0x6c, 0x44, 0x61, 0x79, 0x73,
	0x12, 0x5d, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x61, 0x74,
	0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x50, 0x7a, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74,
	0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a, 0x69, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73,
	0x50, 0x7a, 0x69, 0x1a, 0xbc, 0x02, 0x0a, 0x13, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x69, 0x78, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x95, 0x01, 0x0a, 0x07, 0x53,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4e, 0x6f,
	0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x16, 0x63, 0x70, 0x75, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x14, 0x63, 0x70,
	0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x3a, 0x71, 0xea, 0x41, 0x6e, 0x0a, 0x26, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x44, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x42, 0xe8, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x11, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_featurestore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_featurestore_proto_goTypes = []any{
	(Featurestore_State)(0),                  // 0: google.cloud.aiplatform.v1beta1.Featurestore.State
	(*Featurestore)(nil),                     // 1: google.cloud.aiplatform.v1beta1.Featurestore
	(*Featurestore_OnlineServingConfig)(nil), // 2: google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig
	nil,                                      // 3: google.cloud.aiplatform.v1beta1.Featurestore.LabelsEntry
	(*Featurestore_OnlineServingConfig_Scaling)(nil), // 4: google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling
	(*timestamppb.Timestamp)(nil),                    // 5: google.protobuf.Timestamp
	(*EncryptionSpec)(nil),                           // 6: google.cloud.aiplatform.v1beta1.EncryptionSpec
}
var file_google_cloud_aiplatform_v1beta1_featurestore_proto_depIdxs = []int32{
	5, // 0: google.cloud.aiplatform.v1beta1.Featurestore.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: google.cloud.aiplatform.v1beta1.Featurestore.update_time:type_name -> google.protobuf.Timestamp
	3, // 2: google.cloud.aiplatform.v1beta1.Featurestore.labels:type_name -> google.cloud.aiplatform.v1beta1.Featurestore.LabelsEntry
	2, // 3: google.cloud.aiplatform.v1beta1.Featurestore.online_serving_config:type_name -> google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig
	0, // 4: google.cloud.aiplatform.v1beta1.Featurestore.state:type_name -> google.cloud.aiplatform.v1beta1.Featurestore.State
	6, // 5: google.cloud.aiplatform.v1beta1.Featurestore.encryption_spec:type_name -> google.cloud.aiplatform.v1beta1.EncryptionSpec
	4, // 6: google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.scaling:type_name -> google.cloud.aiplatform.v1beta1.Featurestore.OnlineServingConfig.Scaling
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_featurestore_proto_init() }
func file_google_cloud_aiplatform_v1beta1_featurestore_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_featurestore_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_encryption_spec_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_featurestore_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_featurestore_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_featurestore_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_featurestore_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_featurestore_proto = out.File
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_featurestore_proto_depIdxs = nil
}
