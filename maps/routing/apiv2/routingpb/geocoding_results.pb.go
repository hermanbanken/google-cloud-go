// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/maps/routing/v2/geocoding_results.proto

package routingpb

import (
	reflect "reflect"
	sync "sync"

	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contains [GeocodedWaypoints][google.maps.routing.v2.GeocodedWaypoint] for
// origin, destination and intermediate waypoints. Only populated for address
// waypoints.
type GeocodingResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Origin geocoded waypoint.
	Origin *GeocodedWaypoint `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	// Destination geocoded waypoint.
	Destination *GeocodedWaypoint `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// A list of intermediate geocoded waypoints each containing an index field
	// that corresponds to the zero-based position of the waypoint in the order
	// they were specified in the request.
	Intermediates []*GeocodedWaypoint `protobuf:"bytes,3,rep,name=intermediates,proto3" json:"intermediates,omitempty"`
}

func (x *GeocodingResults) Reset() {
	*x = GeocodingResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_geocoding_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeocodingResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeocodingResults) ProtoMessage() {}

func (x *GeocodingResults) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_geocoding_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeocodingResults.ProtoReflect.Descriptor instead.
func (*GeocodingResults) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_geocoding_results_proto_rawDescGZIP(), []int{0}
}

func (x *GeocodingResults) GetOrigin() *GeocodedWaypoint {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *GeocodingResults) GetDestination() *GeocodedWaypoint {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *GeocodingResults) GetIntermediates() []*GeocodedWaypoint {
	if x != nil {
		return x.Intermediates
	}
	return nil
}

// Details about the locations used as waypoints. Only populated for address
// waypoints. Includes details about the geocoding results for the purposes of
// determining what the address was geocoded to.
type GeocodedWaypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the status code resulting from the geocoding operation.
	GeocoderStatus *status.Status `protobuf:"bytes,1,opt,name=geocoder_status,json=geocoderStatus,proto3" json:"geocoder_status,omitempty"`
	// The index of the corresponding intermediate waypoint in the request.
	// Only populated if the corresponding waypoint is an intermediate
	// waypoint.
	IntermediateWaypointRequestIndex *int32 `protobuf:"varint,2,opt,name=intermediate_waypoint_request_index,json=intermediateWaypointRequestIndex,proto3,oneof" json:"intermediate_waypoint_request_index,omitempty"`
	// The type(s) of the result, in the form of zero or more type tags.
	// Supported types:
	// https://developers.google.com/maps/documentation/geocoding/requests-geocoding#Types
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// Indicates that the geocoder did not return an exact match for the original
	// request, though it was able to match part of the requested address. You may
	// wish to examine the original request for misspellings and/or an incomplete
	// address.
	PartialMatch bool `protobuf:"varint,4,opt,name=partial_match,json=partialMatch,proto3" json:"partial_match,omitempty"`
	// The place ID for this result.
	PlaceId string `protobuf:"bytes,5,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *GeocodedWaypoint) Reset() {
	*x = GeocodedWaypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_geocoding_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeocodedWaypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeocodedWaypoint) ProtoMessage() {}

func (x *GeocodedWaypoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_geocoding_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeocodedWaypoint.ProtoReflect.Descriptor instead.
func (*GeocodedWaypoint) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_geocoding_results_proto_rawDescGZIP(), []int{1}
}

func (x *GeocodedWaypoint) GetGeocoderStatus() *status.Status {
	if x != nil {
		return x.GeocoderStatus
	}
	return nil
}

func (x *GeocodedWaypoint) GetIntermediateWaypointRequestIndex() int32 {
	if x != nil && x.IntermediateWaypointRequestIndex != nil {
		return *x.IntermediateWaypointRequestIndex
	}
	return 0
}

func (x *GeocodedWaypoint) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *GeocodedWaypoint) GetPartialMatch() bool {
	if x != nil {
		return x.PartialMatch
	}
	return false
}

func (x *GeocodedWaypoint) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

var File_google_maps_routing_v2_geocoding_results_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_geocoding_results_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x4a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x57,
	0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x57, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x67, 0x65, 0x6f,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x52, 0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x26,
	0x0a, 0x24, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0xcd, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x15, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_geocoding_results_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_geocoding_results_proto_rawDescData = file_google_maps_routing_v2_geocoding_results_proto_rawDesc
)

func file_google_maps_routing_v2_geocoding_results_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_geocoding_results_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_geocoding_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_geocoding_results_proto_rawDescData)
	})
	return file_google_maps_routing_v2_geocoding_results_proto_rawDescData
}

var file_google_maps_routing_v2_geocoding_results_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_routing_v2_geocoding_results_proto_goTypes = []interface{}{
	(*GeocodingResults)(nil), // 0: google.maps.routing.v2.GeocodingResults
	(*GeocodedWaypoint)(nil), // 1: google.maps.routing.v2.GeocodedWaypoint
	(*status.Status)(nil),    // 2: google.rpc.Status
}
var file_google_maps_routing_v2_geocoding_results_proto_depIdxs = []int32{
	1, // 0: google.maps.routing.v2.GeocodingResults.origin:type_name -> google.maps.routing.v2.GeocodedWaypoint
	1, // 1: google.maps.routing.v2.GeocodingResults.destination:type_name -> google.maps.routing.v2.GeocodedWaypoint
	1, // 2: google.maps.routing.v2.GeocodingResults.intermediates:type_name -> google.maps.routing.v2.GeocodedWaypoint
	2, // 3: google.maps.routing.v2.GeocodedWaypoint.geocoder_status:type_name -> google.rpc.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_geocoding_results_proto_init() }
func file_google_maps_routing_v2_geocoding_results_proto_init() {
	if File_google_maps_routing_v2_geocoding_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routing_v2_geocoding_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeocodingResults); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_maps_routing_v2_geocoding_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeocodedWaypoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_maps_routing_v2_geocoding_results_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_geocoding_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_geocoding_results_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_geocoding_results_proto_depIdxs,
		MessageInfos:      file_google_maps_routing_v2_geocoding_results_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_geocoding_results_proto = out.File
	file_google_maps_routing_v2_geocoding_results_proto_rawDesc = nil
	file_google_maps_routing_v2_geocoding_results_proto_goTypes = nil
	file_google_maps_routing_v2_geocoding_results_proto_depIdxs = nil
}
