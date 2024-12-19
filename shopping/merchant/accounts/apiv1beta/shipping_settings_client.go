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

package accounts

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newShippingSettingsClientHook clientHook

// ShippingSettingsCallOptions contains the retry settings for each method of ShippingSettingsClient.
type ShippingSettingsCallOptions struct {
	GetShippingSettings    []gax.CallOption
	InsertShippingSettings []gax.CallOption
}

func defaultShippingSettingsGRPCClientOptions() []option.ClientOption {
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

func defaultShippingSettingsCallOptions() *ShippingSettingsCallOptions {
	return &ShippingSettingsCallOptions{
		GetShippingSettings:    []gax.CallOption{},
		InsertShippingSettings: []gax.CallOption{},
	}
}

func defaultShippingSettingsRESTCallOptions() *ShippingSettingsCallOptions {
	return &ShippingSettingsCallOptions{
		GetShippingSettings:    []gax.CallOption{},
		InsertShippingSettings: []gax.CallOption{},
	}
}

// internalShippingSettingsClient is an interface that defines the methods available from Merchant API.
type internalShippingSettingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetShippingSettings(context.Context, *accountspb.GetShippingSettingsRequest, ...gax.CallOption) (*accountspb.ShippingSettings, error)
	InsertShippingSettings(context.Context, *accountspb.InsertShippingSettingsRequest, ...gax.CallOption) (*accountspb.ShippingSettings, error)
}

// ShippingSettingsClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to get method call shipping setting information per Merchant API
// method.
type ShippingSettingsClient struct {
	// The internal transport-dependent client.
	internalClient internalShippingSettingsClient

	// The call options for this service.
	CallOptions *ShippingSettingsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ShippingSettingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ShippingSettingsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ShippingSettingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetShippingSettings retrieve shipping setting information.
func (c *ShippingSettingsClient) GetShippingSettings(ctx context.Context, req *accountspb.GetShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	return c.internalClient.GetShippingSettings(ctx, req, opts...)
}

// InsertShippingSettings replace the shipping setting of a merchant with the request shipping
// setting. Executing this method requires admin access.
func (c *ShippingSettingsClient) InsertShippingSettings(ctx context.Context, req *accountspb.InsertShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	return c.internalClient.InsertShippingSettings(ctx, req, opts...)
}

// shippingSettingsGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type shippingSettingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ShippingSettingsClient
	CallOptions **ShippingSettingsCallOptions

	// The gRPC API client.
	shippingSettingsClient accountspb.ShippingSettingsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewShippingSettingsClient creates a new shipping settings service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to get method call shipping setting information per Merchant API
// method.
func NewShippingSettingsClient(ctx context.Context, opts ...option.ClientOption) (*ShippingSettingsClient, error) {
	clientOpts := defaultShippingSettingsGRPCClientOptions()
	if newShippingSettingsClientHook != nil {
		hookOpts, err := newShippingSettingsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ShippingSettingsClient{CallOptions: defaultShippingSettingsCallOptions()}

	c := &shippingSettingsGRPCClient{
		connPool:               connPool,
		shippingSettingsClient: accountspb.NewShippingSettingsServiceClient(connPool),
		CallOptions:            &client.CallOptions,
		logger:                 internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *shippingSettingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *shippingSettingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *shippingSettingsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type shippingSettingsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ShippingSettingsClient
	CallOptions **ShippingSettingsCallOptions

	logger *slog.Logger
}

// NewShippingSettingsRESTClient creates a new shipping settings service rest client.
//
// Service to get method call shipping setting information per Merchant API
// method.
func NewShippingSettingsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ShippingSettingsClient, error) {
	clientOpts := append(defaultShippingSettingsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultShippingSettingsRESTCallOptions()
	c := &shippingSettingsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &ShippingSettingsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultShippingSettingsRESTClientOptions() []option.ClientOption {
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
func (c *shippingSettingsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *shippingSettingsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *shippingSettingsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *shippingSettingsGRPCClient) GetShippingSettings(ctx context.Context, req *accountspb.GetShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetShippingSettings[0:len((*c.CallOptions).GetShippingSettings):len((*c.CallOptions).GetShippingSettings)], opts...)
	var resp *accountspb.ShippingSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.shippingSettingsClient.GetShippingSettings, req, settings.GRPC, c.logger, "GetShippingSettings")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *shippingSettingsGRPCClient) InsertShippingSettings(ctx context.Context, req *accountspb.InsertShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertShippingSettings[0:len((*c.CallOptions).InsertShippingSettings):len((*c.CallOptions).InsertShippingSettings)], opts...)
	var resp *accountspb.ShippingSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.shippingSettingsClient.InsertShippingSettings, req, settings.GRPC, c.logger, "InsertShippingSettings")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetShippingSettings retrieve shipping setting information.
func (c *shippingSettingsRESTClient) GetShippingSettings(ctx context.Context, req *accountspb.GetShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetShippingSettings[0:len((*c.CallOptions).GetShippingSettings):len((*c.CallOptions).GetShippingSettings)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.ShippingSettings{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetShippingSettings")
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

// InsertShippingSettings replace the shipping setting of a merchant with the request shipping
// setting. Executing this method requires admin access.
func (c *shippingSettingsRESTClient) InsertShippingSettings(ctx context.Context, req *accountspb.InsertShippingSettingsRequest, opts ...gax.CallOption) (*accountspb.ShippingSettings, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetShippingSetting()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v/shippingSettings:insert", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertShippingSettings[0:len((*c.CallOptions).InsertShippingSettings):len((*c.CallOptions).InsertShippingSettings)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.ShippingSettings{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "InsertShippingSettings")
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
