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
// source: google/cloud/security/publicca/v1/resources.proto

package publiccapb

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

// A representation of an ExternalAccountKey used for [external account
// binding](https://tools.ietf.org/html/rfc8555#section-7.3.4) within ACME.
type ExternalAccountKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name.
	// projects/{project}/locations/{location}/externalAccountKeys/{key_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Key ID.
	// It is generated by the PublicCertificateAuthorityService
	// when the ExternalAccountKey is created
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Output only. Base64-URL-encoded HS256 key.
	// It is generated by the PublicCertificateAuthorityService
	// when the ExternalAccountKey is created
	B64MacKey []byte `protobuf:"bytes,3,opt,name=b64_mac_key,json=b64MacKey,proto3" json:"b64_mac_key,omitempty"`
}

func (x *ExternalAccountKey) Reset() {
	*x = ExternalAccountKey{}
	mi := &file_google_cloud_security_publicca_v1_resources_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExternalAccountKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAccountKey) ProtoMessage() {}

func (x *ExternalAccountKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_security_publicca_v1_resources_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAccountKey.ProtoReflect.Descriptor instead.
func (*ExternalAccountKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_security_publicca_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalAccountKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalAccountKey) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *ExternalAccountKey) GetB64MacKey() []byte {
	if x != nil {
		return x.B64MacKey
	}
	return nil
}

var File_google_cloud_security_publicca_v1_resources_proto protoreflect.FileDescriptor

var file_google_cloud_security_publicca_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x63, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0b, 0x62, 0x36, 0x34, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x62, 0x36, 0x34, 0x4d, 0x61, 0x63,
	0x4b, 0x65, 0x79, 0x3a, 0x84, 0x01, 0xea, 0x41, 0x80, 0x01, 0x0a, 0x2a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x63, 0x61, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x52, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0xef, 0x01, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x63,
	0x61, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x63, 0x61, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x63, 0x61, 0x70, 0x62, 0x3b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x63, 0x61, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x41, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43,
	0x41, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x3a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x41, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_security_publicca_v1_resources_proto_rawDescOnce sync.Once
	file_google_cloud_security_publicca_v1_resources_proto_rawDescData = file_google_cloud_security_publicca_v1_resources_proto_rawDesc
)

func file_google_cloud_security_publicca_v1_resources_proto_rawDescGZIP() []byte {
	file_google_cloud_security_publicca_v1_resources_proto_rawDescOnce.Do(func() {
		file_google_cloud_security_publicca_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_security_publicca_v1_resources_proto_rawDescData)
	})
	return file_google_cloud_security_publicca_v1_resources_proto_rawDescData
}

var file_google_cloud_security_publicca_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_security_publicca_v1_resources_proto_goTypes = []any{
	(*ExternalAccountKey)(nil), // 0: google.cloud.security.publicca.v1.ExternalAccountKey
}
var file_google_cloud_security_publicca_v1_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_security_publicca_v1_resources_proto_init() }
func file_google_cloud_security_publicca_v1_resources_proto_init() {
	if File_google_cloud_security_publicca_v1_resources_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_security_publicca_v1_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_security_publicca_v1_resources_proto_goTypes,
		DependencyIndexes: file_google_cloud_security_publicca_v1_resources_proto_depIdxs,
		MessageInfos:      file_google_cloud_security_publicca_v1_resources_proto_msgTypes,
	}.Build()
	File_google_cloud_security_publicca_v1_resources_proto = out.File
	file_google_cloud_security_publicca_v1_resources_proto_rawDesc = nil
	file_google_cloud_security_publicca_v1_resources_proto_goTypes = nil
	file_google_cloud_security_publicca_v1_resources_proto_depIdxs = nil
}
