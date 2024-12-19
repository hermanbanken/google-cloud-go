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
// source: google/cloud/websecurityscanner/v1/scan_run_log.proto

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

// A ScanRunLog is an output-only proto used for Stackdriver customer logging.
// It is used for logs covering the start and end of scan pipelines.
// Other than an added summary, this is a subset of the ScanRun.
// Representation in logs is either a proto Struct, or converted to JSON.
// Next id: 9
type ScanRunLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Human friendly message about the event.
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// The resource name of the ScanRun being logged.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The execution state of the ScanRun.
	ExecutionState ScanRun_ExecutionState `protobuf:"varint,3,opt,name=execution_state,json=executionState,proto3,enum=google.cloud.websecurityscanner.v1.ScanRun_ExecutionState" json:"execution_state,omitempty"`
	// The result state of the ScanRun.
	ResultState      ScanRun_ResultState `protobuf:"varint,4,opt,name=result_state,json=resultState,proto3,enum=google.cloud.websecurityscanner.v1.ScanRun_ResultState" json:"result_state,omitempty"`
	UrlsCrawledCount int64               `protobuf:"varint,5,opt,name=urls_crawled_count,json=urlsCrawledCount,proto3" json:"urls_crawled_count,omitempty"`
	UrlsTestedCount  int64               `protobuf:"varint,6,opt,name=urls_tested_count,json=urlsTestedCount,proto3" json:"urls_tested_count,omitempty"`
	HasFindings      bool                `protobuf:"varint,7,opt,name=has_findings,json=hasFindings,proto3" json:"has_findings,omitempty"`
	ErrorTrace       *ScanRunErrorTrace  `protobuf:"bytes,8,opt,name=error_trace,json=errorTrace,proto3" json:"error_trace,omitempty"`
}

func (x *ScanRunLog) Reset() {
	*x = ScanRunLog{}
	mi := &file_google_cloud_websecurityscanner_v1_scan_run_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanRunLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRunLog) ProtoMessage() {}

func (x *ScanRunLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_scan_run_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRunLog.ProtoReflect.Descriptor instead.
func (*ScanRunLog) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRunLog) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ScanRunLog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScanRunLog) GetExecutionState() ScanRun_ExecutionState {
	if x != nil {
		return x.ExecutionState
	}
	return ScanRun_EXECUTION_STATE_UNSPECIFIED
}

func (x *ScanRunLog) GetResultState() ScanRun_ResultState {
	if x != nil {
		return x.ResultState
	}
	return ScanRun_RESULT_STATE_UNSPECIFIED
}

func (x *ScanRunLog) GetUrlsCrawledCount() int64 {
	if x != nil {
		return x.UrlsCrawledCount
	}
	return 0
}

func (x *ScanRunLog) GetUrlsTestedCount() int64 {
	if x != nil {
		return x.UrlsTestedCount
	}
	return 0
}

func (x *ScanRunLog) GetHasFindings() bool {
	if x != nil {
		return x.HasFindings
	}
	return false
}

func (x *ScanRunLog) GetErrorTrace() *ScanRunErrorTrace {
	if x != nil {
		return x.ErrorTrace
	}
	return nil
}

var File_google_cloud_websecurityscanner_v1_scan_run_log_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x31, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x65, 0x62,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x03,
	0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x63, 0x0a, 0x0f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x5a, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52,
	0x75, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x75,
	0x72, 0x6c, 0x73, 0x5f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x72, 0x6c, 0x73, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x72, 0x6c,
	0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x72, 0x6c, 0x73, 0x54, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x66, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73,
	0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x42, 0x85, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53, 0x63, 0x61,
	0x6e, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70,
	0x62, 0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x57, 0x65, 0x62, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescData = file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1_scan_run_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_websecurityscanner_v1_scan_run_log_proto_goTypes = []any{
	(*ScanRunLog)(nil),          // 0: google.cloud.websecurityscanner.v1.ScanRunLog
	(ScanRun_ExecutionState)(0), // 1: google.cloud.websecurityscanner.v1.ScanRun.ExecutionState
	(ScanRun_ResultState)(0),    // 2: google.cloud.websecurityscanner.v1.ScanRun.ResultState
	(*ScanRunErrorTrace)(nil),   // 3: google.cloud.websecurityscanner.v1.ScanRunErrorTrace
}
var file_google_cloud_websecurityscanner_v1_scan_run_log_proto_depIdxs = []int32{
	1, // 0: google.cloud.websecurityscanner.v1.ScanRunLog.execution_state:type_name -> google.cloud.websecurityscanner.v1.ScanRun.ExecutionState
	2, // 1: google.cloud.websecurityscanner.v1.ScanRunLog.result_state:type_name -> google.cloud.websecurityscanner.v1.ScanRun.ResultState
	3, // 2: google.cloud.websecurityscanner.v1.ScanRunLog.error_trace:type_name -> google.cloud.websecurityscanner.v1.ScanRunErrorTrace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1_scan_run_log_proto_init() }
func file_google_cloud_websecurityscanner_v1_scan_run_log_proto_init() {
	if File_google_cloud_websecurityscanner_v1_scan_run_log_proto != nil {
		return
	}
	file_google_cloud_websecurityscanner_v1_scan_run_proto_init()
	file_google_cloud_websecurityscanner_v1_scan_run_error_trace_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1_scan_run_log_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1_scan_run_log_proto_depIdxs,
		MessageInfos:      file_google_cloud_websecurityscanner_v1_scan_run_log_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1_scan_run_log_proto = out.File
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1_scan_run_log_proto_depIdxs = nil
}
