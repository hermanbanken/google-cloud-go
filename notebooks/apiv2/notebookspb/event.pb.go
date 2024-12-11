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
// source: google/cloud/notebooks/v2/event.proto

package notebookspb

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

// The definition of the event types.
type Event_EventType int32

const (
	// Event is not specified.
	Event_EVENT_TYPE_UNSPECIFIED Event_EventType = 0
	// The instance / runtime is idle
	Event_IDLE Event_EventType = 1
	// The instance / runtime is available.
	// This event indicates that instance / runtime underlying compute is
	// operational.
	Event_HEARTBEAT Event_EventType = 2
	// The instance / runtime health is available.
	// This event indicates that instance / runtime health information.
	Event_HEALTH Event_EventType = 3
	// The instance / runtime is available.
	// This event allows instance / runtime to send Host maintenance
	// information to Control Plane.
	// https://cloud.google.com/compute/docs/gpus/gpu-host-maintenance
	Event_MAINTENANCE Event_EventType = 4
	// The instance / runtime is available.
	// This event indicates that the instance had metadata that needs to be
	// modified.
	Event_METADATA_CHANGE Event_EventType = 5
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "IDLE",
		2: "HEARTBEAT",
		3: "HEALTH",
		4: "MAINTENANCE",
		5: "METADATA_CHANGE",
	}
	Event_EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED": 0,
		"IDLE":                   1,
		"HEARTBEAT":              2,
		"HEALTH":                 3,
		"MAINTENANCE":            4,
		"METADATA_CHANGE":        5,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_notebooks_v2_event_proto_enumTypes[0].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_google_cloud_notebooks_v2_event_proto_enumTypes[0]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v2_event_proto_rawDescGZIP(), []int{0, 0}
}

// The definition of an Event for a managed / semi-managed notebook instance.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Event report time.
	ReportTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=report_time,json=reportTime,proto3" json:"report_time,omitempty"`
	// Optional. Event type.
	Type Event_EventType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.notebooks.v2.Event_EventType" json:"type,omitempty"`
	// Optional. Event details. This field is used to pass event information.
	Details map[string]string `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_google_cloud_notebooks_v2_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v2_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v2_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetReportTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReportTime
	}
	return nil
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_google_cloud_notebooks_v2_event_proto protoreflect.FileDescriptor

var file_google_cloud_notebooks_v2_event_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x40,
	0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x72, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x05, 0x42, 0xc1, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x70, 0x62, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62,
	0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x19, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_notebooks_v2_event_proto_rawDescOnce sync.Once
	file_google_cloud_notebooks_v2_event_proto_rawDescData = file_google_cloud_notebooks_v2_event_proto_rawDesc
)

func file_google_cloud_notebooks_v2_event_proto_rawDescGZIP() []byte {
	file_google_cloud_notebooks_v2_event_proto_rawDescOnce.Do(func() {
		file_google_cloud_notebooks_v2_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_notebooks_v2_event_proto_rawDescData)
	})
	return file_google_cloud_notebooks_v2_event_proto_rawDescData
}

var file_google_cloud_notebooks_v2_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_notebooks_v2_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_notebooks_v2_event_proto_goTypes = []any{
	(Event_EventType)(0),          // 0: google.cloud.notebooks.v2.Event.EventType
	(*Event)(nil),                 // 1: google.cloud.notebooks.v2.Event
	nil,                           // 2: google.cloud.notebooks.v2.Event.DetailsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_notebooks_v2_event_proto_depIdxs = []int32{
	3, // 0: google.cloud.notebooks.v2.Event.report_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.cloud.notebooks.v2.Event.type:type_name -> google.cloud.notebooks.v2.Event.EventType
	2, // 2: google.cloud.notebooks.v2.Event.details:type_name -> google.cloud.notebooks.v2.Event.DetailsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_notebooks_v2_event_proto_init() }
func file_google_cloud_notebooks_v2_event_proto_init() {
	if File_google_cloud_notebooks_v2_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_notebooks_v2_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_notebooks_v2_event_proto_goTypes,
		DependencyIndexes: file_google_cloud_notebooks_v2_event_proto_depIdxs,
		EnumInfos:         file_google_cloud_notebooks_v2_event_proto_enumTypes,
		MessageInfos:      file_google_cloud_notebooks_v2_event_proto_msgTypes,
	}.Build()
	File_google_cloud_notebooks_v2_event_proto = out.File
	file_google_cloud_notebooks_v2_event_proto_rawDesc = nil
	file_google_cloud_notebooks_v2_event_proto_goTypes = nil
	file_google_cloud_notebooks_v2_event_proto_depIdxs = nil
}
