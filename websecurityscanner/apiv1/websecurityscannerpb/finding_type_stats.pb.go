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
// source: google/cloud/websecurityscanner/v1/finding_type_stats.proto

package websecurityscannerpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A FindingTypeStats resource represents stats regarding a specific FindingType
// of Findings under a given ScanRun.
type FindingTypeStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The finding type associated with the stats.
	FindingType string `protobuf:"bytes,1,opt,name=finding_type,json=findingType,proto3" json:"finding_type,omitempty"`
	// Output only. The count of findings belonging to this finding type.
	FindingCount int32 `protobuf:"varint,2,opt,name=finding_count,json=findingCount,proto3" json:"finding_count,omitempty"`
}

func (x *FindingTypeStats) Reset() {
	*x = FindingTypeStats{}
	mi := &file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindingTypeStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindingTypeStats) ProtoMessage() {}

func (x *FindingTypeStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindingTypeStats.ProtoReflect.Descriptor instead.
func (*FindingTypeStats) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescGZIP(), []int{0}
}

func (x *FindingTypeStats) GetFindingType() string {
	if x != nil {
		return x.FindingType
	}
	return ""
}

func (x *FindingTypeStats) GetFindingCount() int32 {
	if x != nil {
		return x.FindingCount
	}
	return 0
}

var File_google_cloud_websecurityscanner_v1_finding_type_stats_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x5a, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x8b, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x70, 0x62, 0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x57, 0x65,
	0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescData = file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_goTypes = []any{
	(*FindingTypeStats)(nil), // 0: google.cloud.websecurityscanner.v1.FindingTypeStats
}
var file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_init() }
func file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_init() {
	if File_google_cloud_websecurityscanner_v1_finding_type_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_depIdxs,
		MessageInfos:      file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1_finding_type_stats_proto = out.File
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1_finding_type_stats_proto_depIdxs = nil
}
