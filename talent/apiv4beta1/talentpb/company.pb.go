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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/talent/v4beta1/company.proto

package talentpb

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

// A Company resource represents a company in the service. A company is the
// entity that owns job postings, that is, the hiring entity responsible for
// employing applicants for the job position.
type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required during company update.
	//
	// The resource name for a company. This is generated by the service when a
	// company is created.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/foo/tenants/bar/companies/baz".
	//
	// If tenant id is unspecified, the default tenant is used. For
	// example, "projects/foo/companies/bar".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the company, for example, "Google LLC".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. Client side company identifier, used to uniquely identify the
	// company.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// The employer's company size.
	Size CompanySize `protobuf:"varint,4,opt,name=size,proto3,enum=google.cloud.talent.v4beta1.CompanySize" json:"size,omitempty"`
	// The street address of the company's main headquarters, which may be
	// different from the job location. The service attempts
	// to geolocate the provided address, and populates a more specific
	// location wherever possible in
	// [DerivedInfo.headquarters_location][google.cloud.talent.v4beta1.Company.DerivedInfo.headquarters_location].
	HeadquartersAddress string `protobuf:"bytes,5,opt,name=headquarters_address,json=headquartersAddress,proto3" json:"headquarters_address,omitempty"`
	// Set to true if it is the hiring agency that post jobs for other
	// employers.
	//
	// Defaults to false if not provided.
	HiringAgency bool `protobuf:"varint,6,opt,name=hiring_agency,json=hiringAgency,proto3" json:"hiring_agency,omitempty"`
	// Equal Employment Opportunity legal disclaimer text to be
	// associated with all jobs, and typically to be displayed in all
	// roles.
	//
	// The maximum number of allowed characters is 500.
	EeoText string `protobuf:"bytes,7,opt,name=eeo_text,json=eeoText,proto3" json:"eeo_text,omitempty"`
	// The URI representing the company's primary web site or home page,
	// for example, "https://www.google.com".
	//
	// The maximum number of allowed characters is 255.
	WebsiteUri string `protobuf:"bytes,8,opt,name=website_uri,json=websiteUri,proto3" json:"website_uri,omitempty"`
	// The URI to employer's career site or careers page on the employer's web
	// site, for example, "https://careers.google.com".
	CareerSiteUri string `protobuf:"bytes,9,opt,name=career_site_uri,json=careerSiteUri,proto3" json:"career_site_uri,omitempty"`
	// A URI that hosts the employer's company logo.
	ImageUri string `protobuf:"bytes,10,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	// This field is deprecated. Please set the searchability of the custom
	// attribute in the
	// [Job.custom_attributes][google.cloud.talent.v4beta1.Job.custom_attributes]
	// going forward.
	//
	// A list of keys of filterable
	// [Job.custom_attributes][google.cloud.talent.v4beta1.Job.custom_attributes],
	// whose corresponding `string_values` are used in keyword searches. Jobs with
	// `string_values` under these specified field keys are returned if any
	// of the values match the search keyword. Custom field values with
	// parenthesis, brackets and special symbols are not searchable as-is,
	// and those keyword queries must be surrounded by quotes.
	//
	// Deprecated: Marked as deprecated in google/cloud/talent/v4beta1/company.proto.
	KeywordSearchableJobCustomAttributes []string `protobuf:"bytes,11,rep,name=keyword_searchable_job_custom_attributes,json=keywordSearchableJobCustomAttributes,proto3" json:"keyword_searchable_job_custom_attributes,omitempty"`
	// Output only. Derived details about the company.
	DerivedInfo *Company_DerivedInfo `protobuf:"bytes,12,opt,name=derived_info,json=derivedInfo,proto3" json:"derived_info,omitempty"`
	// Output only. Indicates whether a company is flagged to be suspended from
	// public availability by the service when job content appears suspicious,
	// abusive, or spammy.
	Suspended bool `protobuf:"varint,13,opt,name=suspended,proto3" json:"suspended,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4beta1_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4beta1_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4beta1_company_proto_rawDescGZIP(), []int{0}
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Company) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Company) GetSize() CompanySize {
	if x != nil {
		return x.Size
	}
	return CompanySize_COMPANY_SIZE_UNSPECIFIED
}

func (x *Company) GetHeadquartersAddress() string {
	if x != nil {
		return x.HeadquartersAddress
	}
	return ""
}

func (x *Company) GetHiringAgency() bool {
	if x != nil {
		return x.HiringAgency
	}
	return false
}

func (x *Company) GetEeoText() string {
	if x != nil {
		return x.EeoText
	}
	return ""
}

func (x *Company) GetWebsiteUri() string {
	if x != nil {
		return x.WebsiteUri
	}
	return ""
}

func (x *Company) GetCareerSiteUri() string {
	if x != nil {
		return x.CareerSiteUri
	}
	return ""
}

func (x *Company) GetImageUri() string {
	if x != nil {
		return x.ImageUri
	}
	return ""
}

// Deprecated: Marked as deprecated in google/cloud/talent/v4beta1/company.proto.
func (x *Company) GetKeywordSearchableJobCustomAttributes() []string {
	if x != nil {
		return x.KeywordSearchableJobCustomAttributes
	}
	return nil
}

func (x *Company) GetDerivedInfo() *Company_DerivedInfo {
	if x != nil {
		return x.DerivedInfo
	}
	return nil
}

func (x *Company) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

// Derived details about the company.
type Company_DerivedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A structured headquarters location of the company, resolved from
	// [Company.headquarters_address][google.cloud.talent.v4beta1.Company.headquarters_address]
	// if provided.
	HeadquartersLocation *Location `protobuf:"bytes,1,opt,name=headquarters_location,json=headquartersLocation,proto3" json:"headquarters_location,omitempty"`
}

func (x *Company_DerivedInfo) Reset() {
	*x = Company_DerivedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4beta1_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company_DerivedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company_DerivedInfo) ProtoMessage() {}

func (x *Company_DerivedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4beta1_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company_DerivedInfo.ProtoReflect.Descriptor instead.
func (*Company_DerivedInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4beta1_company_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Company_DerivedInfo) GetHeadquartersLocation() *Location {
	if x != nil {
		return x.HeadquartersLocation
	}
	return nil
}

var File_google_cloud_talent_v4beta1_company_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4beta1_company_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca,
	0x06, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x68, 0x65,
	0x61, 0x64, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x68, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x65, 0x6f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x65, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x69, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x53,
	0x69, 0x74, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x69, 0x12, 0x5a, 0x0a, 0x28, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x24, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x58, 0x0a, 0x0c, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x44, 0x65, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x64, 0x65,
	0x72, 0x69, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x75, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x1a, 0x69, 0x0a, 0x0b,
	0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x15, 0x68,
	0x65, 0x61, 0x64, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x14, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x73, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x81, 0x01, 0xea, 0x41, 0x7e, 0x0a, 0x1b, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x37, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x7d, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x7d, 0x12, 0x26, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x7d, 0x42, 0x78, 0x0a, 0x1f, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x14,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2,
	0x02, 0x03, 0x43, 0x54, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_talent_v4beta1_company_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4beta1_company_proto_rawDescData = file_google_cloud_talent_v4beta1_company_proto_rawDesc
)

func file_google_cloud_talent_v4beta1_company_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4beta1_company_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4beta1_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4beta1_company_proto_rawDescData)
	})
	return file_google_cloud_talent_v4beta1_company_proto_rawDescData
}

var file_google_cloud_talent_v4beta1_company_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_talent_v4beta1_company_proto_goTypes = []interface{}{
	(*Company)(nil),             // 0: google.cloud.talent.v4beta1.Company
	(*Company_DerivedInfo)(nil), // 1: google.cloud.talent.v4beta1.Company.DerivedInfo
	(CompanySize)(0),            // 2: google.cloud.talent.v4beta1.CompanySize
	(*Location)(nil),            // 3: google.cloud.talent.v4beta1.Location
}
var file_google_cloud_talent_v4beta1_company_proto_depIdxs = []int32{
	2, // 0: google.cloud.talent.v4beta1.Company.size:type_name -> google.cloud.talent.v4beta1.CompanySize
	1, // 1: google.cloud.talent.v4beta1.Company.derived_info:type_name -> google.cloud.talent.v4beta1.Company.DerivedInfo
	3, // 2: google.cloud.talent.v4beta1.Company.DerivedInfo.headquarters_location:type_name -> google.cloud.talent.v4beta1.Location
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4beta1_company_proto_init() }
func file_google_cloud_talent_v4beta1_company_proto_init() {
	if File_google_cloud_talent_v4beta1_company_proto != nil {
		return
	}
	file_google_cloud_talent_v4beta1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_talent_v4beta1_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_google_cloud_talent_v4beta1_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company_DerivedInfo); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_talent_v4beta1_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_talent_v4beta1_company_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4beta1_company_proto_depIdxs,
		MessageInfos:      file_google_cloud_talent_v4beta1_company_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4beta1_company_proto = out.File
	file_google_cloud_talent_v4beta1_company_proto_rawDesc = nil
	file_google_cloud_talent_v4beta1_company_proto_goTypes = nil
	file_google_cloud_talent_v4beta1_company_proto_depIdxs = nil
}
