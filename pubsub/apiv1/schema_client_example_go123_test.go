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

package pubsub_test

import (
	"context"

	pubsub "cloud.google.com/go/pubsub/apiv1"
	pubsubpb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
)

func ExampleSchemaClient_ListSchemaRevisions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &pubsubpb.ListSchemaRevisionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/pubsub/apiv1/pubsubpb#ListSchemaRevisionsRequest.
	}
	for resp, err := range c.ListSchemaRevisions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleSchemaClient_ListSchemas_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &pubsubpb.ListSchemasRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/pubsub/apiv1/pubsubpb#ListSchemasRequest.
	}
	for resp, err := range c.ListSchemas(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
