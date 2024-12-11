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
// source: google/cloud/dataqna/v1alpha/user_feedback.proto

package dataqnapb

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

// Enumeration of feedback ratings.
type UserFeedback_UserFeedbackRating int32

const (
	// No rating was specified.
	UserFeedback_USER_FEEDBACK_RATING_UNSPECIFIED UserFeedback_UserFeedbackRating = 0
	// The user provided positive feedback.
	UserFeedback_POSITIVE UserFeedback_UserFeedbackRating = 1
	// The user provided negative feedback.
	UserFeedback_NEGATIVE UserFeedback_UserFeedbackRating = 2
)

// Enum value maps for UserFeedback_UserFeedbackRating.
var (
	UserFeedback_UserFeedbackRating_name = map[int32]string{
		0: "USER_FEEDBACK_RATING_UNSPECIFIED",
		1: "POSITIVE",
		2: "NEGATIVE",
	}
	UserFeedback_UserFeedbackRating_value = map[string]int32{
		"USER_FEEDBACK_RATING_UNSPECIFIED": 0,
		"POSITIVE":                         1,
		"NEGATIVE":                         2,
	}
)

func (x UserFeedback_UserFeedbackRating) Enum() *UserFeedback_UserFeedbackRating {
	p := new(UserFeedback_UserFeedbackRating)
	*p = x
	return p
}

func (x UserFeedback_UserFeedbackRating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserFeedback_UserFeedbackRating) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataqna_v1alpha_user_feedback_proto_enumTypes[0].Descriptor()
}

func (UserFeedback_UserFeedbackRating) Type() protoreflect.EnumType {
	return &file_google_cloud_dataqna_v1alpha_user_feedback_proto_enumTypes[0]
}

func (x UserFeedback_UserFeedbackRating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserFeedback_UserFeedbackRating.Descriptor instead.
func (UserFeedback_UserFeedbackRating) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescGZIP(), []int{0, 0}
}

// Feedback provided by a user.
type UserFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique identifier for the user feedback.
	// User feedback is a singleton resource on a Question.
	// Example: `projects/foo/locations/bar/questions/1234/userFeedback`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Free form user feedback, such as a text box.
	FreeFormFeedback string `protobuf:"bytes,2,opt,name=free_form_feedback,json=freeFormFeedback,proto3" json:"free_form_feedback,omitempty"`
	// The user feedback rating
	Rating UserFeedback_UserFeedbackRating `protobuf:"varint,3,opt,name=rating,proto3,enum=google.cloud.dataqna.v1alpha.UserFeedback_UserFeedbackRating" json:"rating,omitempty"`
}

func (x *UserFeedback) Reset() {
	*x = UserFeedback{}
	mi := &file_google_cloud_dataqna_v1alpha_user_feedback_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFeedback) ProtoMessage() {}

func (x *UserFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_user_feedback_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFeedback.ProtoReflect.Descriptor instead.
func (*UserFeedback) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *UserFeedback) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserFeedback) GetFreeFormFeedback() string {
	if x != nil {
		return x.FreeFormFeedback
	}
	return ""
}

func (x *UserFeedback) GetRating() UserFeedback_UserFeedbackRating {
	if x != nil {
		return x.Rating
	}
	return UserFeedback_USER_FEEDBACK_RATING_UNSPECIFIED
}

var File_google_cloud_dataqna_v1alpha_user_feedback_proto protoreflect.FileDescriptor

var file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x66, 0x72, 0x65, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x55, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x56, 0x0a, 0x12, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x42, 0x41,
	0x43, 0x4b, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x02, 0x3a, 0x73, 0xea, 0x41, 0x70, 0x0a, 0x23, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e,
	0x61, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x49, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x42, 0xd3, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x71,
	0x6e, 0x61, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x70, 0x62, 0xaa, 0x02,
	0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x51, 0x6e, 0x41, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x51, 0x6e, 0x41, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61,
	0x74, 0x61, 0x51, 0x6e, 0x41, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescOnce sync.Once
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescData = file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDesc
)

func file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescGZIP() []byte {
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescData)
	})
	return file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDescData
}

var file_google_cloud_dataqna_v1alpha_user_feedback_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dataqna_v1alpha_user_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_dataqna_v1alpha_user_feedback_proto_goTypes = []any{
	(UserFeedback_UserFeedbackRating)(0), // 0: google.cloud.dataqna.v1alpha.UserFeedback.UserFeedbackRating
	(*UserFeedback)(nil),                 // 1: google.cloud.dataqna.v1alpha.UserFeedback
}
var file_google_cloud_dataqna_v1alpha_user_feedback_proto_depIdxs = []int32{
	0, // 0: google.cloud.dataqna.v1alpha.UserFeedback.rating:type_name -> google.cloud.dataqna.v1alpha.UserFeedback.UserFeedbackRating
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_dataqna_v1alpha_user_feedback_proto_init() }
func file_google_cloud_dataqna_v1alpha_user_feedback_proto_init() {
	if File_google_cloud_dataqna_v1alpha_user_feedback_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataqna_v1alpha_user_feedback_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataqna_v1alpha_user_feedback_proto_depIdxs,
		EnumInfos:         file_google_cloud_dataqna_v1alpha_user_feedback_proto_enumTypes,
		MessageInfos:      file_google_cloud_dataqna_v1alpha_user_feedback_proto_msgTypes,
	}.Build()
	File_google_cloud_dataqna_v1alpha_user_feedback_proto = out.File
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_rawDesc = nil
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_goTypes = nil
	file_google_cloud_dataqna_v1alpha_user_feedback_proto_depIdxs = nil
}
