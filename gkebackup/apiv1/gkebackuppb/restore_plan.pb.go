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
// source: google/cloud/gkebackup/v1/restore_plan.proto

package gkebackuppb

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

// State
type RestorePlan_State int32

const (
	// Default first value for Enums.
	RestorePlan_STATE_UNSPECIFIED RestorePlan_State = 0
	// Waiting for cluster state to be RUNNING.
	RestorePlan_CLUSTER_PENDING RestorePlan_State = 1
	// The RestorePlan has successfully been created and is ready for Restores.
	RestorePlan_READY RestorePlan_State = 2
	// RestorePlan creation has failed.
	RestorePlan_FAILED RestorePlan_State = 3
	// The RestorePlan is in the process of being deleted.
	RestorePlan_DELETING RestorePlan_State = 4
)

// Enum value maps for RestorePlan_State.
var (
	RestorePlan_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CLUSTER_PENDING",
		2: "READY",
		3: "FAILED",
		4: "DELETING",
	}
	RestorePlan_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CLUSTER_PENDING":   1,
		"READY":             2,
		"FAILED":            3,
		"DELETING":          4,
	}
)

func (x RestorePlan_State) Enum() *RestorePlan_State {
	p := new(RestorePlan_State)
	*p = x
	return p
}

func (x RestorePlan_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RestorePlan_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_gkebackup_v1_restore_plan_proto_enumTypes[0].Descriptor()
}

func (RestorePlan_State) Type() protoreflect.EnumType {
	return &file_google_cloud_gkebackup_v1_restore_plan_proto_enumTypes[0]
}

func (x RestorePlan_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RestorePlan_State.Descriptor instead.
func (RestorePlan_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescGZIP(), []int{0, 0}
}

// The configuration of a potential series of Restore operations to be performed
// against Backups belong to a particular BackupPlan.
type RestorePlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The full name of the RestorePlan resource.
	// Format: `projects/*/locations/*/restorePlans/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Server generated global unique identifier of
	// [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. The timestamp when this RestorePlan resource was
	// created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when this RestorePlan resource was last
	// updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. User specified descriptive string for this RestorePlan.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Immutable. A reference to the
	// [BackupPlan][google.cloud.gkebackup.v1.BackupPlan] from which Backups may
	// be used as the source for Restores created via this RestorePlan. Format:
	// `projects/*/locations/*/backupPlans/*`.
	BackupPlan string `protobuf:"bytes,6,opt,name=backup_plan,json=backupPlan,proto3" json:"backup_plan,omitempty"`
	// Required. Immutable. The target cluster into which Restores created via
	// this RestorePlan will restore data. NOTE: the cluster's region must be the
	// same as the RestorePlan. Valid formats:
	//
	//   - `projects/*/locations/*/clusters/*`
	//   - `projects/*/zones/*/clusters/*`
	Cluster string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Required. Configuration of Restores created via this RestorePlan.
	RestoreConfig *RestoreConfig `protobuf:"bytes,8,opt,name=restore_config,json=restoreConfig,proto3" json:"restore_config,omitempty"`
	// Optional. A set of custom labels supplied by user.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. `etag` is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a restore from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform restore updates in order to avoid
	// race conditions: An `etag` is returned in the response to `GetRestorePlan`,
	// and systems are expected to put that etag in the request to
	// `UpdateRestorePlan` or `DeleteRestorePlan` to ensure that their change
	// will be applied to the same version of the resource.
	Etag string `protobuf:"bytes,10,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. State of the RestorePlan. This State field reflects the
	// various stages a RestorePlan can be in
	// during the Create operation.
	State RestorePlan_State `protobuf:"varint,11,opt,name=state,proto3,enum=google.cloud.gkebackup.v1.RestorePlan_State" json:"state,omitempty"`
	// Output only. Human-readable description of why RestorePlan is in the
	// current `state`
	StateReason string `protobuf:"bytes,12,opt,name=state_reason,json=stateReason,proto3" json:"state_reason,omitempty"`
}

func (x *RestorePlan) Reset() {
	*x = RestorePlan{}
	mi := &file_google_cloud_gkebackup_v1_restore_plan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestorePlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestorePlan) ProtoMessage() {}

func (x *RestorePlan) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_restore_plan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestorePlan.ProtoReflect.Descriptor instead.
func (*RestorePlan) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescGZIP(), []int{0}
}

func (x *RestorePlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestorePlan) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RestorePlan) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RestorePlan) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RestorePlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RestorePlan) GetBackupPlan() string {
	if x != nil {
		return x.BackupPlan
	}
	return ""
}

func (x *RestorePlan) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *RestorePlan) GetRestoreConfig() *RestoreConfig {
	if x != nil {
		return x.RestoreConfig
	}
	return nil
}

func (x *RestorePlan) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RestorePlan) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *RestorePlan) GetState() RestorePlan_State {
	if x != nil {
		return x.State
	}
	return RestorePlan_STATE_UNSPECIFIED
}

func (x *RestorePlan) GetStateReason() string {
	if x != nil {
		return x.StateReason
	}
	return ""
}

var File_google_cloud_gkebackup_v1_restore_plan_proto protoreflect.FileDescriptor

var file_google_cloud_gkebackup_v1_restore_plan_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb6, 0x07, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0b, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2e, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x67, 0x6b, 0x65, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41,
	0x05, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x54, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x58,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x3a, 0x6e, 0xea, 0x41, 0x6b, 0x0a, 0x24, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x7d, 0x42, 0xc7, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b,
	0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70, 0x62, 0x3b,
	0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70, 0x62, 0xaa, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescOnce sync.Once
	file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescData = file_google_cloud_gkebackup_v1_restore_plan_proto_rawDesc
)

func file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescGZIP() []byte {
	file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescOnce.Do(func() {
		file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescData)
	})
	return file_google_cloud_gkebackup_v1_restore_plan_proto_rawDescData
}

var file_google_cloud_gkebackup_v1_restore_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_gkebackup_v1_restore_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_gkebackup_v1_restore_plan_proto_goTypes = []any{
	(RestorePlan_State)(0),        // 0: google.cloud.gkebackup.v1.RestorePlan.State
	(*RestorePlan)(nil),           // 1: google.cloud.gkebackup.v1.RestorePlan
	nil,                           // 2: google.cloud.gkebackup.v1.RestorePlan.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*RestoreConfig)(nil),         // 4: google.cloud.gkebackup.v1.RestoreConfig
}
var file_google_cloud_gkebackup_v1_restore_plan_proto_depIdxs = []int32{
	3, // 0: google.cloud.gkebackup.v1.RestorePlan.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.gkebackup.v1.RestorePlan.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.gkebackup.v1.RestorePlan.restore_config:type_name -> google.cloud.gkebackup.v1.RestoreConfig
	2, // 3: google.cloud.gkebackup.v1.RestorePlan.labels:type_name -> google.cloud.gkebackup.v1.RestorePlan.LabelsEntry
	0, // 4: google.cloud.gkebackup.v1.RestorePlan.state:type_name -> google.cloud.gkebackup.v1.RestorePlan.State
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_gkebackup_v1_restore_plan_proto_init() }
func file_google_cloud_gkebackup_v1_restore_plan_proto_init() {
	if File_google_cloud_gkebackup_v1_restore_plan_proto != nil {
		return
	}
	file_google_cloud_gkebackup_v1_restore_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gkebackup_v1_restore_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gkebackup_v1_restore_plan_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkebackup_v1_restore_plan_proto_depIdxs,
		EnumInfos:         file_google_cloud_gkebackup_v1_restore_plan_proto_enumTypes,
		MessageInfos:      file_google_cloud_gkebackup_v1_restore_plan_proto_msgTypes,
	}.Build()
	File_google_cloud_gkebackup_v1_restore_plan_proto = out.File
	file_google_cloud_gkebackup_v1_restore_plan_proto_rawDesc = nil
	file_google_cloud_gkebackup_v1_restore_plan_proto_goTypes = nil
	file_google_cloud_gkebackup_v1_restore_plan_proto_depIdxs = nil
}
