// Copyright 2022 Google LLC
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

package vision

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1p1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newImageAnnotatorClientHook clientHook

// ImageAnnotatorCallOptions contains the retry settings for each method of ImageAnnotatorClient.
type ImageAnnotatorCallOptions struct {
	BatchAnnotateImages []gax.CallOption
}

func defaultImageAnnotatorGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("vision.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("vision.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://vision.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultImageAnnotatorCallOptions() *ImageAnnotatorCallOptions {
	return &ImageAnnotatorCallOptions{
		BatchAnnotateImages: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalImageAnnotatorClient is an interface that defines the methods availaible from Cloud Vision API.
type internalImageAnnotatorClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	BatchAnnotateImages(context.Context, *visionpb.BatchAnnotateImagesRequest, ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error)
}

// ImageAnnotatorClient is a client for interacting with Cloud Vision API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
type ImageAnnotatorClient struct {
	// The internal transport-dependent client.
	internalClient internalImageAnnotatorClient

	// The call options for this service.
	CallOptions *ImageAnnotatorCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageAnnotatorClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageAnnotatorClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ImageAnnotatorClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// BatchAnnotateImages run image detection and annotation for a batch of images.
func (c *ImageAnnotatorClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	return c.internalClient.BatchAnnotateImages(ctx, req, opts...)
}

// imageAnnotatorGRPCClient is a client for interacting with Cloud Vision API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageAnnotatorGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ImageAnnotatorClient
	CallOptions **ImageAnnotatorCallOptions

	// The gRPC API client.
	imageAnnotatorClient visionpb.ImageAnnotatorClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewImageAnnotatorClient creates a new image annotator client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
func NewImageAnnotatorClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	clientOpts := defaultImageAnnotatorGRPCClientOptions()
	if newImageAnnotatorClientHook != nil {
		hookOpts, err := newImageAnnotatorClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ImageAnnotatorClient{CallOptions: defaultImageAnnotatorCallOptions()}

	c := &imageAnnotatorGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		imageAnnotatorClient: visionpb.NewImageAnnotatorClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *imageAnnotatorGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageAnnotatorGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageAnnotatorGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *imageAnnotatorGRPCClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).BatchAnnotateImages[0:len((*c.CallOptions).BatchAnnotateImages):len((*c.CallOptions).BatchAnnotateImages)], opts...)
	var resp *visionpb.BatchAnnotateImagesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
