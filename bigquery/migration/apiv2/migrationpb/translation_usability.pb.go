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
// source: google/cloud/bigquery/migration/v2/translation_usability.proto

package migrationpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A record in the aggregate CSV report for a migration workflow
type GcsReportLogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Severity of the translation record.
	Severity string `protobuf:"bytes,1,opt,name=severity,proto3" json:"severity,omitempty"`
	// Category of the error/warning. Example: SyntaxError
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// The file path in which the error occurred
	FilePath string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// The file name in which the error occurred
	Filename string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	// Specifies the row from the source text where the error occurred (0 based,
	// -1 for messages without line location). Example: 2
	SourceScriptLine int32 `protobuf:"varint,5,opt,name=source_script_line,json=sourceScriptLine,proto3" json:"source_script_line,omitempty"`
	// Specifies the column from the source texts where the error occurred. (0
	// based, -1 for messages without column location) example: 6
	SourceScriptColumn int32 `protobuf:"varint,6,opt,name=source_script_column,json=sourceScriptColumn,proto3" json:"source_script_column,omitempty"`
	// Detailed message of the record.
	Message string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	// The script context (obfuscated) in which the error occurred
	ScriptContext string `protobuf:"bytes,8,opt,name=script_context,json=scriptContext,proto3" json:"script_context,omitempty"`
	// Category of the error/warning. Example: SyntaxError
	Action string `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty"`
	// Effect of the error/warning. Example: COMPATIBILITY
	Effect string `protobuf:"bytes,10,opt,name=effect,proto3" json:"effect,omitempty"`
	// Name of the affected object in the log message.
	ObjectName string `protobuf:"bytes,11,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
}

func (x *GcsReportLogMessage) Reset() {
	*x = GcsReportLogMessage{}
	mi := &file_google_cloud_bigquery_migration_v2_translation_usability_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcsReportLogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsReportLogMessage) ProtoMessage() {}

func (x *GcsReportLogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2_translation_usability_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsReportLogMessage.ProtoReflect.Descriptor instead.
func (*GcsReportLogMessage) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescGZIP(), []int{0}
}

func (x *GcsReportLogMessage) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *GcsReportLogMessage) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GcsReportLogMessage) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *GcsReportLogMessage) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *GcsReportLogMessage) GetSourceScriptLine() int32 {
	if x != nil {
		return x.SourceScriptLine
	}
	return 0
}

func (x *GcsReportLogMessage) GetSourceScriptColumn() int32 {
	if x != nil {
		return x.SourceScriptColumn
	}
	return 0
}

func (x *GcsReportLogMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GcsReportLogMessage) GetScriptContext() string {
	if x != nil {
		return x.ScriptContext
	}
	return ""
}

func (x *GcsReportLogMessage) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GcsReportLogMessage) GetEffect() string {
	if x != nil {
		return x.Effect
	}
	return ""
}

func (x *GcsReportLogMessage) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

var File_google_cloud_bigquery_migration_v2_translation_usability_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x75, 0x73, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x02, 0x0a, 0x13, 0x47, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xd5, 0x01, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescData = file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2_translation_usability_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_bigquery_migration_v2_translation_usability_proto_goTypes = []any{
	(*GcsReportLogMessage)(nil), // 0: google.cloud.bigquery.migration.v2.GcsReportLogMessage
}
var file_google_cloud_bigquery_migration_v2_translation_usability_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2_translation_usability_proto_init() }
func file_google_cloud_bigquery_migration_v2_translation_usability_proto_init() {
	if File_google_cloud_bigquery_migration_v2_translation_usability_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2_translation_usability_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2_translation_usability_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_migration_v2_translation_usability_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2_translation_usability_proto = out.File
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2_translation_usability_proto_depIdxs = nil
}
