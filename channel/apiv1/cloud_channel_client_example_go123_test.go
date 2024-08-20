// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package channel_test

import (
	"context"

	channel "cloud.google.com/go/channel/apiv1"
	channelpb "cloud.google.com/go/channel/apiv1/channelpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
)

func ExampleCloudChannelClient_ListChannelPartnerLinks_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListChannelPartnerLinksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListChannelPartnerLinksRequest.
	}
	for resp, err := range c.ListChannelPartnerLinks(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListChannelPartnerRepricingConfigs_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListChannelPartnerRepricingConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListChannelPartnerRepricingConfigsRequest.
	}
	for resp, err := range c.ListChannelPartnerRepricingConfigs(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListCustomerRepricingConfigs_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListCustomerRepricingConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListCustomerRepricingConfigsRequest.
	}
	for resp, err := range c.ListCustomerRepricingConfigs(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListCustomers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListCustomersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListCustomersRequest.
	}
	for resp, err := range c.ListCustomers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListEntitlementChanges_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListEntitlementChangesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListEntitlementChangesRequest.
	}
	for resp, err := range c.ListEntitlementChanges(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListEntitlements_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListEntitlementsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListEntitlementsRequest.
	}
	for resp, err := range c.ListEntitlements(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListOffers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListOffersRequest.
	}
	for resp, err := range c.ListOffers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListProducts_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListProductsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListProductsRequest.
	}
	for resp, err := range c.ListProducts(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListPurchasableOffers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListPurchasableOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListPurchasableOffersRequest.
	}
	for resp, err := range c.ListPurchasableOffers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListPurchasableSkus_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListPurchasableSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListPurchasableSkusRequest.
	}
	for resp, err := range c.ListPurchasableSkus(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListSkuGroupBillableSkus_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSkuGroupBillableSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListSkuGroupBillableSkusRequest.
	}
	for resp, err := range c.ListSkuGroupBillableSkus(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListSkuGroups_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSkuGroupsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListSkuGroupsRequest.
	}
	for resp, err := range c.ListSkuGroups(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListSkus_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListSkusRequest.
	}
	for resp, err := range c.ListSkus(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListSubscribers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSubscribersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListSubscribersRequest.
	}
	for resp, err := range c.ListSubscribers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListTransferableOffers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListTransferableOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListTransferableOffersRequest.
	}
	for resp, err := range c.ListTransferableOffers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListTransferableSkus_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListTransferableSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/channel/apiv1/channelpb#ListTransferableSkusRequest.
	}
	for resp, err := range c.ListTransferableSkus(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudChannelClient_ListOperations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	for resp, err := range c.ListOperations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
