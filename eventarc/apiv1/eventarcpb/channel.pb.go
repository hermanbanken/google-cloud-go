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
// source: google/cloud/eventarc/v1/channel.proto

package eventarcpb

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

// State lists all the possible states of a Channel
type Channel_State int32

const (
	// Default value. This value is unused.
	Channel_STATE_UNSPECIFIED Channel_State = 0
	// The PENDING state indicates that a Channel has been created successfully
	// and there is a new activation token available for the subscriber to use
	// to convey the Channel to the provider in order to create a Connection.
	Channel_PENDING Channel_State = 1
	// The ACTIVE state indicates that a Channel has been successfully
	// connected with the event provider.
	// An ACTIVE Channel is ready to receive and route events from the
	// event provider.
	Channel_ACTIVE Channel_State = 2
	// The INACTIVE state indicates that the Channel cannot receive events
	// permanently. There are two possible cases this state can happen:
	//
	//  1. The SaaS provider disconnected from this Channel.
	//  2. The Channel activation token has expired but the SaaS provider
	//     wasn't connected.
	//
	// To re-establish a Connection with a provider, the subscriber
	// should create a new Channel and give it to the provider.
	Channel_INACTIVE Channel_State = 3
)

// Enum value maps for Channel_State.
var (
	Channel_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "ACTIVE",
		3: "INACTIVE",
	}
	Channel_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"ACTIVE":            2,
		"INACTIVE":          3,
	}
)

func (x Channel_State) Enum() *Channel_State {
	p := new(Channel_State)
	*p = x
	return p
}

func (x Channel_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_eventarc_v1_channel_proto_enumTypes[0].Descriptor()
}

func (Channel_State) Type() protoreflect.EnumType {
	return &file_google_cloud_eventarc_v1_channel_proto_enumTypes[0]
}

func (x Channel_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel_State.Descriptor instead.
func (Channel_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_v1_channel_proto_rawDescGZIP(), []int{0, 0}
}

// A representation of the Channel resource.
// A Channel is a resource on which event providers publish their events.
// The published events are delivered through the transport associated with the
// channel. Note that a channel is associated with exactly one event provider.
type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the channel. Must be unique within the
	// location on the project and must be in
	// `projects/{project}/locations/{location}/channels/{channel_id}` format.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Server assigned unique identifier for the channel. The value
	// is a UUID4 string and guaranteed to remain unchanged until the resource is
	// deleted.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. The creation time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last-modified time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The name of the event provider (e.g. Eventarc SaaS partner) associated
	// with the channel. This provider will be granted permissions to publish
	// events to the channel. Format:
	// `projects/{project}/locations/{location}/providers/{provider_id}`.
	Provider string `protobuf:"bytes,7,opt,name=provider,proto3" json:"provider,omitempty"`
	// Types that are assignable to Transport:
	//
	//	*Channel_PubsubTopic
	Transport isChannel_Transport `protobuf_oneof:"transport"`
	// Output only. The state of a Channel.
	State Channel_State `protobuf:"varint,9,opt,name=state,proto3,enum=google.cloud.eventarc.v1.Channel_State" json:"state,omitempty"`
	// Output only. The activation token for the channel. The token must be used
	// by the provider to register the channel for publishing.
	ActivationToken string `protobuf:"bytes,10,opt,name=activation_token,json=activationToken,proto3" json:"activation_token,omitempty"`
	// Resource name of a KMS crypto key (managed by the user) used to
	// encrypt/decrypt their event data.
	//
	// It must match the pattern
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName string `protobuf:"bytes,11,opt,name=crypto_key_name,json=cryptoKeyName,proto3" json:"crypto_key_name,omitempty"`
	// Output only. Whether or not this Channel satisfies the requirements of
	// physical zone separation
	SatisfiesPzs bool `protobuf:"varint,12,opt,name=satisfies_pzs,json=satisfiesPzs,proto3" json:"satisfies_pzs,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	mi := &file_google_cloud_eventarc_v1_channel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_eventarc_v1_channel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_v1_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Channel) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Channel) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Channel) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Channel) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (m *Channel) GetTransport() isChannel_Transport {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (x *Channel) GetPubsubTopic() string {
	if x, ok := x.GetTransport().(*Channel_PubsubTopic); ok {
		return x.PubsubTopic
	}
	return ""
}

func (x *Channel) GetState() Channel_State {
	if x != nil {
		return x.State
	}
	return Channel_STATE_UNSPECIFIED
}

func (x *Channel) GetActivationToken() string {
	if x != nil {
		return x.ActivationToken
	}
	return ""
}

func (x *Channel) GetCryptoKeyName() string {
	if x != nil {
		return x.CryptoKeyName
	}
	return ""
}

func (x *Channel) GetSatisfiesPzs() bool {
	if x != nil {
		return x.SatisfiesPzs
	}
	return false
}

type isChannel_Transport interface {
	isChannel_Transport()
}

type Channel_PubsubTopic struct {
	// Output only. The name of the Pub/Sub topic created and managed by
	// Eventarc system as a transport for the event delivery. Format:
	// `projects/{project}/topics/{topic_id}`.
	PubsubTopic string `protobuf:"bytes,8,opt,name=pubsub_topic,json=pubsubTopic,proto3,oneof"`
}

func (*Channel_PubsubTopic) isChannel_Transport() {}

var File_google_cloud_eventarc_v1_channel_proto protoreflect.FileDescriptor

var file_google_cloud_eventarc_v1_channel_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xba, 0x05, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x10, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4e, 0x0a, 0x0f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74,
	0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73,
	0x50, 0x7a, 0x73, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x3a, 0x73, 0xea, 0x41, 0x70, 0x0a,
	0x1f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x2a, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x32, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42,
	0x0b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x42, 0xbc, 0x01, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x70, 0x62, 0x3b, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x61, 0x72, 0x63, 0x70, 0x62, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_eventarc_v1_channel_proto_rawDescOnce sync.Once
	file_google_cloud_eventarc_v1_channel_proto_rawDescData = file_google_cloud_eventarc_v1_channel_proto_rawDesc
)

func file_google_cloud_eventarc_v1_channel_proto_rawDescGZIP() []byte {
	file_google_cloud_eventarc_v1_channel_proto_rawDescOnce.Do(func() {
		file_google_cloud_eventarc_v1_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_eventarc_v1_channel_proto_rawDescData)
	})
	return file_google_cloud_eventarc_v1_channel_proto_rawDescData
}

var file_google_cloud_eventarc_v1_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_eventarc_v1_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_eventarc_v1_channel_proto_goTypes = []any{
	(Channel_State)(0),            // 0: google.cloud.eventarc.v1.Channel.State
	(*Channel)(nil),               // 1: google.cloud.eventarc.v1.Channel
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_cloud_eventarc_v1_channel_proto_depIdxs = []int32{
	2, // 0: google.cloud.eventarc.v1.Channel.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.eventarc.v1.Channel.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.cloud.eventarc.v1.Channel.state:type_name -> google.cloud.eventarc.v1.Channel.State
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_eventarc_v1_channel_proto_init() }
func file_google_cloud_eventarc_v1_channel_proto_init() {
	if File_google_cloud_eventarc_v1_channel_proto != nil {
		return
	}
	file_google_cloud_eventarc_v1_channel_proto_msgTypes[0].OneofWrappers = []any{
		(*Channel_PubsubTopic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_eventarc_v1_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_eventarc_v1_channel_proto_goTypes,
		DependencyIndexes: file_google_cloud_eventarc_v1_channel_proto_depIdxs,
		EnumInfos:         file_google_cloud_eventarc_v1_channel_proto_enumTypes,
		MessageInfos:      file_google_cloud_eventarc_v1_channel_proto_msgTypes,
	}.Build()
	File_google_cloud_eventarc_v1_channel_proto = out.File
	file_google_cloud_eventarc_v1_channel_proto_rawDesc = nil
	file_google_cloud_eventarc_v1_channel_proto_goTypes = nil
	file_google_cloud_eventarc_v1_channel_proto_depIdxs = nil
}
