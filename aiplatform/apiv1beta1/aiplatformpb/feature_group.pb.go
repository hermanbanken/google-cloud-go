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
// source: google/cloud/aiplatform/v1beta1/feature_group.proto

package aiplatformpb

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

// Vertex AI Feature Group.
type FeatureGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*FeatureGroup_BigQuery_
	Source isFeatureGroup_Source `protobuf_oneof:"source"`
	// Identifier. Name of the FeatureGroup. Format:
	// `projects/{project}/locations/{location}/featureGroups/{featureGroup}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when this FeatureGroup was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this FeatureGroup was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag string `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// FeatureGroup.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// No more than 64 user labels can be associated with one
	// FeatureGroup(System labels are excluded)." System reserved label keys
	// are prefixed with "aiplatform.googleapis.com/" and are immutable.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Description of the FeatureGroup.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FeatureGroup) Reset() {
	*x = FeatureGroup{}
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureGroup) ProtoMessage() {}

func (x *FeatureGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureGroup.ProtoReflect.Descriptor instead.
func (*FeatureGroup) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescGZIP(), []int{0}
}

func (m *FeatureGroup) GetSource() isFeatureGroup_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *FeatureGroup) GetBigQuery() *FeatureGroup_BigQuery {
	if x, ok := x.GetSource().(*FeatureGroup_BigQuery_); ok {
		return x.BigQuery
	}
	return nil
}

func (x *FeatureGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureGroup) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FeatureGroup) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *FeatureGroup) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *FeatureGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *FeatureGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type isFeatureGroup_Source interface {
	isFeatureGroup_Source()
}

type FeatureGroup_BigQuery_ struct {
	// Indicates that features for this group come from BigQuery Table/View.
	// By default treats the source as a sparse time series source. The BigQuery
	// source table or view must have at least one entity ID column and a column
	// named `feature_timestamp`.
	BigQuery *FeatureGroup_BigQuery `protobuf:"bytes,7,opt,name=big_query,json=bigQuery,proto3,oneof"`
}

func (*FeatureGroup_BigQuery_) isFeatureGroup_Source() {}

// Input source type for BigQuery Tables and Views.
type FeatureGroup_BigQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The BigQuery source URI that points to either a
	// BigQuery Table or View.
	BigQuerySource *BigQuerySource `protobuf:"bytes,1,opt,name=big_query_source,json=bigQuerySource,proto3" json:"big_query_source,omitempty"`
	// Optional. Columns to construct entity_id / row keys.
	// If not provided defaults to `entity_id`.
	EntityIdColumns []string `protobuf:"bytes,2,rep,name=entity_id_columns,json=entityIdColumns,proto3" json:"entity_id_columns,omitempty"`
	// Optional. Set if the data source is not a time-series.
	StaticDataSource bool `protobuf:"varint,3,opt,name=static_data_source,json=staticDataSource,proto3" json:"static_data_source,omitempty"`
	// Optional. If the source is a time-series source, this can be set to
	// control how downstream sources (ex:
	// [FeatureView][google.cloud.aiplatform.v1beta1.FeatureView] ) will treat
	// time-series sources. If not set, will treat the source as a time-series
	// source with `feature_timestamp` as timestamp column and no scan boundary.
	TimeSeries *FeatureGroup_BigQuery_TimeSeries `protobuf:"bytes,4,opt,name=time_series,json=timeSeries,proto3" json:"time_series,omitempty"`
	// Optional. If set, all feature values will be fetched
	// from a single row per unique entityId including nulls.
	// If not set, will collapse all rows for each unique entityId into a singe
	// row with any non-null values if present, if no non-null values are
	// present will sync null.
	// ex: If source has schema
	// `(entity_id, feature_timestamp, f0, f1)` and the following rows:
	// `(e1, 2020-01-01T10:00:00.123Z, 10, 15)`
	// `(e1, 2020-02-01T10:00:00.123Z, 20, null)`
	// If dense is set, `(e1, 20, null)` is synced to online stores. If dense is
	// not set, `(e1, 20, 15)` is synced to online stores.
	Dense bool `protobuf:"varint,5,opt,name=dense,proto3" json:"dense,omitempty"`
}

func (x *FeatureGroup_BigQuery) Reset() {
	*x = FeatureGroup_BigQuery{}
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureGroup_BigQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureGroup_BigQuery) ProtoMessage() {}

func (x *FeatureGroup_BigQuery) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureGroup_BigQuery.ProtoReflect.Descriptor instead.
func (*FeatureGroup_BigQuery) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeatureGroup_BigQuery) GetBigQuerySource() *BigQuerySource {
	if x != nil {
		return x.BigQuerySource
	}
	return nil
}

func (x *FeatureGroup_BigQuery) GetEntityIdColumns() []string {
	if x != nil {
		return x.EntityIdColumns
	}
	return nil
}

func (x *FeatureGroup_BigQuery) GetStaticDataSource() bool {
	if x != nil {
		return x.StaticDataSource
	}
	return false
}

func (x *FeatureGroup_BigQuery) GetTimeSeries() *FeatureGroup_BigQuery_TimeSeries {
	if x != nil {
		return x.TimeSeries
	}
	return nil
}

func (x *FeatureGroup_BigQuery) GetDense() bool {
	if x != nil {
		return x.Dense
	}
	return false
}

type FeatureGroup_BigQuery_TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Column hosting timestamp values for a time-series source.
	// Will be used to determine the latest `feature_values` for each entity.
	// Optional. If not provided, column named `feature_timestamp` of
	// type `TIMESTAMP` will be used.
	TimestampColumn string `protobuf:"bytes,1,opt,name=timestamp_column,json=timestampColumn,proto3" json:"timestamp_column,omitempty"`
}

func (x *FeatureGroup_BigQuery_TimeSeries) Reset() {
	*x = FeatureGroup_BigQuery_TimeSeries{}
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureGroup_BigQuery_TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureGroup_BigQuery_TimeSeries) ProtoMessage() {}

func (x *FeatureGroup_BigQuery_TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureGroup_BigQuery_TimeSeries.ProtoReflect.Descriptor instead.
func (*FeatureGroup_BigQuery_TimeSeries) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FeatureGroup_BigQuery_TimeSeries) GetTimestampColumn() string {
	if x != nil {
		return x.TimestampColumn
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_feature_group_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x08,
	0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x55,
	0x0a, 0x09, 0x62, 0x69, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x62, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x56, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x93, 0x03, 0x0a, 0x08, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x61, 0x0a, 0x10, 0x62, 0x69, 0x67, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x62, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x10, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x67,
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x64, 0x65, 0x6e,
	0x73, 0x65, 0x1a, 0x3c, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x90, 0x01, 0xea, 0x41,
	0x8c, 0x01, 0x0a, 0x26, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x45, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2f, 0x7b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x7d, 0x2a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x32, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0xe8, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x11, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_feature_group_proto_goTypes = []any{
	(*FeatureGroup)(nil),                     // 0: google.cloud.aiplatform.v1beta1.FeatureGroup
	(*FeatureGroup_BigQuery)(nil),            // 1: google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery
	nil,                                      // 2: google.cloud.aiplatform.v1beta1.FeatureGroup.LabelsEntry
	(*FeatureGroup_BigQuery_TimeSeries)(nil), // 3: google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery.TimeSeries
	(*timestamppb.Timestamp)(nil),            // 4: google.protobuf.Timestamp
	(*BigQuerySource)(nil),                   // 5: google.cloud.aiplatform.v1beta1.BigQuerySource
}
var file_google_cloud_aiplatform_v1beta1_feature_group_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1beta1.FeatureGroup.big_query:type_name -> google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery
	4, // 1: google.cloud.aiplatform.v1beta1.FeatureGroup.create_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.aiplatform.v1beta1.FeatureGroup.update_time:type_name -> google.protobuf.Timestamp
	2, // 3: google.cloud.aiplatform.v1beta1.FeatureGroup.labels:type_name -> google.cloud.aiplatform.v1beta1.FeatureGroup.LabelsEntry
	5, // 4: google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery.big_query_source:type_name -> google.cloud.aiplatform.v1beta1.BigQuerySource
	3, // 5: google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery.time_series:type_name -> google.cloud.aiplatform.v1beta1.FeatureGroup.BigQuery.TimeSeries
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_feature_group_proto_init() }
func file_google_cloud_aiplatform_v1beta1_feature_group_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_feature_group_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_io_proto_init()
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes[0].OneofWrappers = []any{
		(*FeatureGroup_BigQuery_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_feature_group_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_feature_group_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_feature_group_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_feature_group_proto = out.File
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_feature_group_proto_depIdxs = nil
}
