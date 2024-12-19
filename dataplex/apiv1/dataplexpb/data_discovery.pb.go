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
// source: google/cloud/dataplex/v1/data_discovery.proto

package dataplexpb

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

// Determines how discovered tables are published.
type DataDiscoverySpec_BigQueryPublishingConfig_TableType int32

const (
	// Table type unspecified.
	DataDiscoverySpec_BigQueryPublishingConfig_TABLE_TYPE_UNSPECIFIED DataDiscoverySpec_BigQueryPublishingConfig_TableType = 0
	// Default. Discovered tables are published as BigQuery external tables
	// whose data is accessed using the credentials of the user querying the
	// table.
	DataDiscoverySpec_BigQueryPublishingConfig_EXTERNAL DataDiscoverySpec_BigQueryPublishingConfig_TableType = 1
	// Discovered tables are published as BigLake external tables whose data
	// is accessed using the credentials of the associated BigQuery
	// connection.
	DataDiscoverySpec_BigQueryPublishingConfig_BIGLAKE DataDiscoverySpec_BigQueryPublishingConfig_TableType = 2
)

// Enum value maps for DataDiscoverySpec_BigQueryPublishingConfig_TableType.
var (
	DataDiscoverySpec_BigQueryPublishingConfig_TableType_name = map[int32]string{
		0: "TABLE_TYPE_UNSPECIFIED",
		1: "EXTERNAL",
		2: "BIGLAKE",
	}
	DataDiscoverySpec_BigQueryPublishingConfig_TableType_value = map[string]int32{
		"TABLE_TYPE_UNSPECIFIED": 0,
		"EXTERNAL":               1,
		"BIGLAKE":                2,
	}
)

func (x DataDiscoverySpec_BigQueryPublishingConfig_TableType) Enum() *DataDiscoverySpec_BigQueryPublishingConfig_TableType {
	p := new(DataDiscoverySpec_BigQueryPublishingConfig_TableType)
	*p = x
	return p
}

func (x DataDiscoverySpec_BigQueryPublishingConfig_TableType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataDiscoverySpec_BigQueryPublishingConfig_TableType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataplex_v1_data_discovery_proto_enumTypes[0].Descriptor()
}

func (DataDiscoverySpec_BigQueryPublishingConfig_TableType) Type() protoreflect.EnumType {
	return &file_google_cloud_dataplex_v1_data_discovery_proto_enumTypes[0]
}

func (x DataDiscoverySpec_BigQueryPublishingConfig_TableType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataDiscoverySpec_BigQueryPublishingConfig_TableType.Descriptor instead.
func (DataDiscoverySpec_BigQueryPublishingConfig_TableType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Spec for a data discovery scan.
type DataDiscoverySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Configuration for metadata publishing.
	BigqueryPublishingConfig *DataDiscoverySpec_BigQueryPublishingConfig `protobuf:"bytes,1,opt,name=bigquery_publishing_config,json=bigqueryPublishingConfig,proto3" json:"bigquery_publishing_config,omitempty"`
	// The configurations of the data discovery scan resource.
	//
	// Types that are assignable to ResourceConfig:
	//
	//	*DataDiscoverySpec_StorageConfig_
	ResourceConfig isDataDiscoverySpec_ResourceConfig `protobuf_oneof:"resource_config"`
}

func (x *DataDiscoverySpec) Reset() {
	*x = DataDiscoverySpec{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoverySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoverySpec) ProtoMessage() {}

func (x *DataDiscoverySpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoverySpec.ProtoReflect.Descriptor instead.
func (*DataDiscoverySpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DataDiscoverySpec) GetBigqueryPublishingConfig() *DataDiscoverySpec_BigQueryPublishingConfig {
	if x != nil {
		return x.BigqueryPublishingConfig
	}
	return nil
}

func (m *DataDiscoverySpec) GetResourceConfig() isDataDiscoverySpec_ResourceConfig {
	if m != nil {
		return m.ResourceConfig
	}
	return nil
}

func (x *DataDiscoverySpec) GetStorageConfig() *DataDiscoverySpec_StorageConfig {
	if x, ok := x.GetResourceConfig().(*DataDiscoverySpec_StorageConfig_); ok {
		return x.StorageConfig
	}
	return nil
}

type isDataDiscoverySpec_ResourceConfig interface {
	isDataDiscoverySpec_ResourceConfig()
}

type DataDiscoverySpec_StorageConfig_ struct {
	// Cloud Storage related configurations.
	StorageConfig *DataDiscoverySpec_StorageConfig `protobuf:"bytes,100,opt,name=storage_config,json=storageConfig,proto3,oneof"`
}

func (*DataDiscoverySpec_StorageConfig_) isDataDiscoverySpec_ResourceConfig() {}

// The output of a data discovery scan.
type DataDiscoveryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Configuration for metadata publishing.
	BigqueryPublishing *DataDiscoveryResult_BigQueryPublishing `protobuf:"bytes,1,opt,name=bigquery_publishing,json=bigqueryPublishing,proto3" json:"bigquery_publishing,omitempty"`
}

func (x *DataDiscoveryResult) Reset() {
	*x = DataDiscoveryResult{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoveryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoveryResult) ProtoMessage() {}

func (x *DataDiscoveryResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoveryResult.ProtoReflect.Descriptor instead.
func (*DataDiscoveryResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DataDiscoveryResult) GetBigqueryPublishing() *DataDiscoveryResult_BigQueryPublishing {
	if x != nil {
		return x.BigqueryPublishing
	}
	return nil
}

// Describes BigQuery publishing configurations.
type DataDiscoverySpec_BigQueryPublishingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Determines whether to  publish discovered tables as BigLake
	// external tables or non-BigLake external tables.
	TableType DataDiscoverySpec_BigQueryPublishingConfig_TableType `protobuf:"varint,2,opt,name=table_type,json=tableType,proto3,enum=google.cloud.dataplex.v1.DataDiscoverySpec_BigQueryPublishingConfig_TableType" json:"table_type,omitempty"`
	// Optional. The BigQuery connection used to create BigLake tables.
	// Must be in the form
	// `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	Connection string `protobuf:"bytes,3,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *DataDiscoverySpec_BigQueryPublishingConfig) Reset() {
	*x = DataDiscoverySpec_BigQueryPublishingConfig{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoverySpec_BigQueryPublishingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoverySpec_BigQueryPublishingConfig) ProtoMessage() {}

func (x *DataDiscoverySpec_BigQueryPublishingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoverySpec_BigQueryPublishingConfig.ProtoReflect.Descriptor instead.
func (*DataDiscoverySpec_BigQueryPublishingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataDiscoverySpec_BigQueryPublishingConfig) GetTableType() DataDiscoverySpec_BigQueryPublishingConfig_TableType {
	if x != nil {
		return x.TableType
	}
	return DataDiscoverySpec_BigQueryPublishingConfig_TABLE_TYPE_UNSPECIFIED
}

func (x *DataDiscoverySpec_BigQueryPublishingConfig) GetConnection() string {
	if x != nil {
		return x.Connection
	}
	return ""
}

// Configurations related to Cloud Storage as the data source.
type DataDiscoverySpec_StorageConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Defines the data to include during discovery when only a subset
	// of the data should be considered. Provide a list of patterns that
	// identify the data to include. For Cloud Storage bucket assets, these
	// patterns are interpreted as glob patterns used to match object names. For
	// BigQuery dataset assets, these patterns are interpreted as patterns to
	// match table names.
	IncludePatterns []string `protobuf:"bytes,1,rep,name=include_patterns,json=includePatterns,proto3" json:"include_patterns,omitempty"`
	// Optional. Defines the data to exclude during discovery. Provide a list of
	// patterns that identify the data to exclude. For Cloud Storage bucket
	// assets, these patterns are interpreted as glob patterns used to match
	// object names. For BigQuery dataset assets, these patterns are interpreted
	// as patterns to match table names.
	ExcludePatterns []string `protobuf:"bytes,2,rep,name=exclude_patterns,json=excludePatterns,proto3" json:"exclude_patterns,omitempty"`
	// Optional. Configuration for CSV data.
	CsvOptions *DataDiscoverySpec_StorageConfig_CsvOptions `protobuf:"bytes,3,opt,name=csv_options,json=csvOptions,proto3" json:"csv_options,omitempty"`
	// Optional. Configuration for JSON data.
	JsonOptions *DataDiscoverySpec_StorageConfig_JsonOptions `protobuf:"bytes,4,opt,name=json_options,json=jsonOptions,proto3" json:"json_options,omitempty"`
}

func (x *DataDiscoverySpec_StorageConfig) Reset() {
	*x = DataDiscoverySpec_StorageConfig{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoverySpec_StorageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoverySpec_StorageConfig) ProtoMessage() {}

func (x *DataDiscoverySpec_StorageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoverySpec_StorageConfig.ProtoReflect.Descriptor instead.
func (*DataDiscoverySpec_StorageConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DataDiscoverySpec_StorageConfig) GetIncludePatterns() []string {
	if x != nil {
		return x.IncludePatterns
	}
	return nil
}

func (x *DataDiscoverySpec_StorageConfig) GetExcludePatterns() []string {
	if x != nil {
		return x.ExcludePatterns
	}
	return nil
}

func (x *DataDiscoverySpec_StorageConfig) GetCsvOptions() *DataDiscoverySpec_StorageConfig_CsvOptions {
	if x != nil {
		return x.CsvOptions
	}
	return nil
}

func (x *DataDiscoverySpec_StorageConfig) GetJsonOptions() *DataDiscoverySpec_StorageConfig_JsonOptions {
	if x != nil {
		return x.JsonOptions
	}
	return nil
}

// Describes CSV and similar semi-structured data formats.
type DataDiscoverySpec_StorageConfig_CsvOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The number of rows to interpret as header rows that should be
	// skipped when reading data rows.
	HeaderRows int32 `protobuf:"varint,1,opt,name=header_rows,json=headerRows,proto3" json:"header_rows,omitempty"`
	// Optional. The delimiter that is used to separate values. The default is
	// `,` (comma).
	Delimiter string `protobuf:"bytes,2,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding string `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
	// Optional. Whether to disable the inference of data types for CSV data.
	// If true, all columns are registered as strings.
	TypeInferenceDisabled bool `protobuf:"varint,4,opt,name=type_inference_disabled,json=typeInferenceDisabled,proto3" json:"type_inference_disabled,omitempty"`
	// Optional. The character used to quote column values. Accepts `"`
	// (double quotation mark) or `'` (single quotation mark). If unspecified,
	// defaults to `"` (double quotation mark).
	Quote string `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) Reset() {
	*x = DataDiscoverySpec_StorageConfig_CsvOptions{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoverySpec_StorageConfig_CsvOptions) ProtoMessage() {}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoverySpec_StorageConfig_CsvOptions.ProtoReflect.Descriptor instead.
func (*DataDiscoverySpec_StorageConfig_CsvOptions) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) GetHeaderRows() int32 {
	if x != nil {
		return x.HeaderRows
	}
	return 0
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) GetDelimiter() string {
	if x != nil {
		return x.Delimiter
	}
	return ""
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) GetTypeInferenceDisabled() bool {
	if x != nil {
		return x.TypeInferenceDisabled
	}
	return false
}

func (x *DataDiscoverySpec_StorageConfig_CsvOptions) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

// Describes JSON data format.
type DataDiscoverySpec_StorageConfig_JsonOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding string `protobuf:"bytes,1,opt,name=encoding,proto3" json:"encoding,omitempty"`
	// Optional. Whether to disable the inference of data types for JSON data.
	// If true, all columns are registered as their primitive types
	// (strings, number, or boolean).
	TypeInferenceDisabled bool `protobuf:"varint,2,opt,name=type_inference_disabled,json=typeInferenceDisabled,proto3" json:"type_inference_disabled,omitempty"`
}

func (x *DataDiscoverySpec_StorageConfig_JsonOptions) Reset() {
	*x = DataDiscoverySpec_StorageConfig_JsonOptions{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoverySpec_StorageConfig_JsonOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoverySpec_StorageConfig_JsonOptions) ProtoMessage() {}

func (x *DataDiscoverySpec_StorageConfig_JsonOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoverySpec_StorageConfig_JsonOptions.ProtoReflect.Descriptor instead.
func (*DataDiscoverySpec_StorageConfig_JsonOptions) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *DataDiscoverySpec_StorageConfig_JsonOptions) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *DataDiscoverySpec_StorageConfig_JsonOptions) GetTypeInferenceDisabled() bool {
	if x != nil {
		return x.TypeInferenceDisabled
	}
	return false
}

// Describes BigQuery publishing configurations.
type DataDiscoveryResult_BigQueryPublishing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The BigQuery dataset to publish to. It takes the form
	// `projects/{project_id}/datasets/{dataset_id}`.
	// If not set, the service creates a default publishing dataset.
	Dataset string `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *DataDiscoveryResult_BigQueryPublishing) Reset() {
	*x = DataDiscoveryResult_BigQueryPublishing{}
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataDiscoveryResult_BigQueryPublishing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDiscoveryResult_BigQueryPublishing) ProtoMessage() {}

func (x *DataDiscoveryResult_BigQueryPublishing) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDiscoveryResult_BigQueryPublishing.ProtoReflect.Descriptor instead.
func (*DataDiscoveryResult_BigQueryPublishing) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DataDiscoveryResult_BigQueryPublishing) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

var File_google_cloud_dataplex_v1_data_discovery_proto protoreflect.FileDescriptor

var file_google_cloud_dataplex_v1_data_discovery_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x09, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x87, 0x01, 0x0a, 0x1a,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x18, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x62, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xa8, 0x02, 0x0a, 0x18, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x72, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34,
	0xe0, 0x41, 0x01, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x42, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x47, 0x4c, 0x41,
	0x4b, 0x45, 0x10, 0x02, 0x1a, 0x88, 0x05, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x6a, 0x0a, 0x0b, 0x63, 0x73, 0x76, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x73, 0x76, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x63, 0x73, 0x76, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x6d, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0xce, 0x01, 0x0a, 0x0a, 0x43, 0x73, 0x76, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x24, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x17, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x15, 0x74, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x1a, 0x6b, 0x0a, 0x0b, 0x4a, 0x73, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x15, 0x74, 0x79, 0x70, 0x65, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42,
	0x11, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x76, 0x0a, 0x13, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x1a, 0x57, 0x0a, 0x12, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xe0, 0x41, 0x03, 0xfa, 0x41,
	0x21, 0x0a, 0x1f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0xac, 0x02, 0xea, 0x41,
	0x48, 0x0a, 0x1f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0xea, 0x41, 0x70, 0x0a, 0x2c, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x44, 0x61, 0x74, 0x61,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x70, 0x62, 0x3b,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_dataplex_v1_data_discovery_proto_rawDescOnce sync.Once
	file_google_cloud_dataplex_v1_data_discovery_proto_rawDescData = file_google_cloud_dataplex_v1_data_discovery_proto_rawDesc
)

func file_google_cloud_dataplex_v1_data_discovery_proto_rawDescGZIP() []byte {
	file_google_cloud_dataplex_v1_data_discovery_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataplex_v1_data_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataplex_v1_data_discovery_proto_rawDescData)
	})
	return file_google_cloud_dataplex_v1_data_discovery_proto_rawDescData
}

var file_google_cloud_dataplex_v1_data_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_dataplex_v1_data_discovery_proto_goTypes = []any{
	(DataDiscoverySpec_BigQueryPublishingConfig_TableType)(0), // 0: google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig.TableType
	(*DataDiscoverySpec)(nil),                                 // 1: google.cloud.dataplex.v1.DataDiscoverySpec
	(*DataDiscoveryResult)(nil),                               // 2: google.cloud.dataplex.v1.DataDiscoveryResult
	(*DataDiscoverySpec_BigQueryPublishingConfig)(nil),        // 3: google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig
	(*DataDiscoverySpec_StorageConfig)(nil),                   // 4: google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig
	(*DataDiscoverySpec_StorageConfig_CsvOptions)(nil),        // 5: google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions
	(*DataDiscoverySpec_StorageConfig_JsonOptions)(nil),       // 6: google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.JsonOptions
	(*DataDiscoveryResult_BigQueryPublishing)(nil),            // 7: google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing
}
var file_google_cloud_dataplex_v1_data_discovery_proto_depIdxs = []int32{
	3, // 0: google.cloud.dataplex.v1.DataDiscoverySpec.bigquery_publishing_config:type_name -> google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig
	4, // 1: google.cloud.dataplex.v1.DataDiscoverySpec.storage_config:type_name -> google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig
	7, // 2: google.cloud.dataplex.v1.DataDiscoveryResult.bigquery_publishing:type_name -> google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing
	0, // 3: google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig.table_type:type_name -> google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig.TableType
	5, // 4: google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.csv_options:type_name -> google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.CsvOptions
	6, // 5: google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.json_options:type_name -> google.cloud.dataplex.v1.DataDiscoverySpec.StorageConfig.JsonOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_dataplex_v1_data_discovery_proto_init() }
func file_google_cloud_dataplex_v1_data_discovery_proto_init() {
	if File_google_cloud_dataplex_v1_data_discovery_proto != nil {
		return
	}
	file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes[0].OneofWrappers = []any{
		(*DataDiscoverySpec_StorageConfig_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dataplex_v1_data_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataplex_v1_data_discovery_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataplex_v1_data_discovery_proto_depIdxs,
		EnumInfos:         file_google_cloud_dataplex_v1_data_discovery_proto_enumTypes,
		MessageInfos:      file_google_cloud_dataplex_v1_data_discovery_proto_msgTypes,
	}.Build()
	File_google_cloud_dataplex_v1_data_discovery_proto = out.File
	file_google_cloud_dataplex_v1_data_discovery_proto_rawDesc = nil
	file_google_cloud_dataplex_v1_data_discovery_proto_goTypes = nil
	file_google_cloud_dataplex_v1_data_discovery_proto_depIdxs = nil
}
