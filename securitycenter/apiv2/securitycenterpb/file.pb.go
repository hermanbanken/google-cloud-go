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
// source: google/cloud/securitycenter/v2/file.proto

package securitycenterpb

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

// File information about the related binary/library used by an executable, or
// the script used by a script interpreter
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Absolute path of the file as a JSON encoded string.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Size of the file in bytes.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// SHA256 hash of the first hashed_size bytes of the file encoded as a
	// hex string.  If hashed_size == size, sha256 represents the SHA256 hash
	// of the entire file.
	Sha256 string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// The length in bytes of the file prefix that was hashed.  If
	// hashed_size == size, any hashes reported represent the entire
	// file.
	HashedSize int64 `protobuf:"varint,4,opt,name=hashed_size,json=hashedSize,proto3" json:"hashed_size,omitempty"`
	// True when the hash covers only a prefix of the file.
	PartiallyHashed bool `protobuf:"varint,5,opt,name=partially_hashed,json=partiallyHashed,proto3" json:"partially_hashed,omitempty"`
	// Prefix of the file contents as a JSON-encoded string.
	Contents string `protobuf:"bytes,6,opt,name=contents,proto3" json:"contents,omitempty"`
	// Path of the file in terms of underlying disk/partition identifiers.
	DiskPath *File_DiskPath `protobuf:"bytes,7,opt,name=disk_path,json=diskPath,proto3" json:"disk_path,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_google_cloud_securitycenter_v2_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *File) GetHashedSize() int64 {
	if x != nil {
		return x.HashedSize
	}
	return 0
}

func (x *File) GetPartiallyHashed() bool {
	if x != nil {
		return x.PartiallyHashed
	}
	return false
}

func (x *File) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

func (x *File) GetDiskPath() *File_DiskPath {
	if x != nil {
		return x.DiskPath
	}
	return nil
}

// Path of the file in terms of underlying disk/partition identifiers.
type File_DiskPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the partition (format
	// https://wiki.archlinux.org/title/persistent_block_device_naming#by-uuid)
	PartitionUuid string `protobuf:"bytes,1,opt,name=partition_uuid,json=partitionUuid,proto3" json:"partition_uuid,omitempty"`
	// Relative path of the file in the partition as a JSON encoded string.
	// Example: /home/user1/executable_file.sh
	RelativePath string `protobuf:"bytes,2,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
}

func (x *File_DiskPath) Reset() {
	*x = File_DiskPath{}
	mi := &file_google_cloud_securitycenter_v2_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File_DiskPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File_DiskPath) ProtoMessage() {}

func (x *File_DiskPath) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File_DiskPath.ProtoReflect.Descriptor instead.
func (*File_DiskPath) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_file_proto_rawDescGZIP(), []int{0, 0}
}

func (x *File_DiskPath) GetPartitionUuid() string {
	if x != nil {
		return x.PartitionUuid
	}
	return ""
}

func (x *File_DiskPath) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

var File_google_cloud_securitycenter_v2_file_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_file_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x22, 0xd2, 0x02, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68,
	0x61, 0x32, 0x35, 0x36, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x6c, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x09,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x56, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x42, 0xe3, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_file_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_file_proto_rawDescData = file_google_cloud_securitycenter_v2_file_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_file_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_file_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_file_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_file_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v2_file_proto_goTypes = []any{
	(*File)(nil),          // 0: google.cloud.securitycenter.v2.File
	(*File_DiskPath)(nil), // 1: google.cloud.securitycenter.v2.File.DiskPath
}
var file_google_cloud_securitycenter_v2_file_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v2.File.disk_path:type_name -> google.cloud.securitycenter.v2.File.DiskPath
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_file_proto_init() }
func file_google_cloud_securitycenter_v2_file_proto_init() {
	if File_google_cloud_securitycenter_v2_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_file_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_file_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v2_file_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_file_proto = out.File
	file_google_cloud_securitycenter_v2_file_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_file_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_file_proto_depIdxs = nil
}
