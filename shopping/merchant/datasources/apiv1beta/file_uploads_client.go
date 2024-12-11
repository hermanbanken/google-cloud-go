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

package datasources

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	datasourcespb "cloud.google.com/go/shopping/merchant/datasources/apiv1beta/datasourcespb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newFileUploadsClientHook clientHook

// FileUploadsCallOptions contains the retry settings for each method of FileUploadsClient.
type FileUploadsCallOptions struct {
	GetFileUpload []gax.CallOption
}

func defaultFileUploadsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("merchantapi.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFileUploadsCallOptions() *FileUploadsCallOptions {
	return &FileUploadsCallOptions{
		GetFileUpload: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultFileUploadsRESTCallOptions() *FileUploadsCallOptions {
	return &FileUploadsCallOptions{
		GetFileUpload: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalFileUploadsClient is an interface that defines the methods available from Merchant API.
type internalFileUploadsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFileUpload(context.Context, *datasourcespb.GetFileUploadRequest, ...gax.CallOption) (*datasourcespb.FileUpload, error)
}

// FileUploadsClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage data source file uploads.
type FileUploadsClient struct {
	// The internal transport-dependent client.
	internalClient internalFileUploadsClient

	// The call options for this service.
	CallOptions *FileUploadsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FileUploadsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FileUploadsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *FileUploadsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFileUpload gets the latest data source file upload. Only the latest alias is
// accepted for a file upload.
func (c *FileUploadsClient) GetFileUpload(ctx context.Context, req *datasourcespb.GetFileUploadRequest, opts ...gax.CallOption) (*datasourcespb.FileUpload, error) {
	return c.internalClient.GetFileUpload(ctx, req, opts...)
}

// fileUploadsGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type fileUploadsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing FileUploadsClient
	CallOptions **FileUploadsCallOptions

	// The gRPC API client.
	fileUploadsClient datasourcespb.FileUploadsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewFileUploadsClient creates a new file uploads service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage data source file uploads.
func NewFileUploadsClient(ctx context.Context, opts ...option.ClientOption) (*FileUploadsClient, error) {
	clientOpts := defaultFileUploadsGRPCClientOptions()
	if newFileUploadsClientHook != nil {
		hookOpts, err := newFileUploadsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FileUploadsClient{CallOptions: defaultFileUploadsCallOptions()}

	c := &fileUploadsGRPCClient{
		connPool:          connPool,
		fileUploadsClient: datasourcespb.NewFileUploadsServiceClient(connPool),
		CallOptions:       &client.CallOptions,
		logger:            internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *fileUploadsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *fileUploadsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *fileUploadsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type fileUploadsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing FileUploadsClient
	CallOptions **FileUploadsCallOptions

	logger *slog.Logger
}

// NewFileUploadsRESTClient creates a new file uploads service rest client.
//
// Service to manage data source file uploads.
func NewFileUploadsRESTClient(ctx context.Context, opts ...option.ClientOption) (*FileUploadsClient, error) {
	clientOpts := append(defaultFileUploadsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultFileUploadsRESTCallOptions()
	c := &fileUploadsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &FileUploadsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultFileUploadsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://merchantapi.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *fileUploadsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *fileUploadsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *fileUploadsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *fileUploadsGRPCClient) GetFileUpload(ctx context.Context, req *datasourcespb.GetFileUploadRequest, opts ...gax.CallOption) (*datasourcespb.FileUpload, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetFileUpload[0:len((*c.CallOptions).GetFileUpload):len((*c.CallOptions).GetFileUpload)], opts...)
	var resp *datasourcespb.FileUpload
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.fileUploadsClient.GetFileUpload, req, settings.GRPC, c.logger, "GetFileUpload")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFileUpload gets the latest data source file upload. Only the latest alias is
// accepted for a file upload.
func (c *fileUploadsRESTClient) GetFileUpload(ctx context.Context, req *datasourcespb.GetFileUploadRequest, opts ...gax.CallOption) (*datasourcespb.FileUpload, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/datasources/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetFileUpload[0:len((*c.CallOptions).GetFileUpload):len((*c.CallOptions).GetFileUpload)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &datasourcespb.FileUpload{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetFileUpload")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
