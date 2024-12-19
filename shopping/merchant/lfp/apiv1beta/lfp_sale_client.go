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

package lfp

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	lfppb "cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newLfpSaleClientHook clientHook

// LfpSaleCallOptions contains the retry settings for each method of LfpSaleClient.
type LfpSaleCallOptions struct {
	InsertLfpSale []gax.CallOption
}

func defaultLfpSaleGRPCClientOptions() []option.ClientOption {
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

func defaultLfpSaleCallOptions() *LfpSaleCallOptions {
	return &LfpSaleCallOptions{
		InsertLfpSale: []gax.CallOption{
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

func defaultLfpSaleRESTCallOptions() *LfpSaleCallOptions {
	return &LfpSaleCallOptions{
		InsertLfpSale: []gax.CallOption{
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

// internalLfpSaleClient is an interface that defines the methods available from Merchant API.
type internalLfpSaleClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	InsertLfpSale(context.Context, *lfppb.InsertLfpSaleRequest, ...gax.CallOption) (*lfppb.LfpSale, error)
}

// LfpSaleClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for a LFP
// partner (at https://support.google.com/merchants/answer/7676652) to submit sales
// data for a merchant.
type LfpSaleClient struct {
	// The internal transport-dependent client.
	internalClient internalLfpSaleClient

	// The call options for this service.
	CallOptions *LfpSaleCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *LfpSaleClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *LfpSaleClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *LfpSaleClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// InsertLfpSale inserts a LfpSale for the given merchant.
func (c *LfpSaleClient) InsertLfpSale(ctx context.Context, req *lfppb.InsertLfpSaleRequest, opts ...gax.CallOption) (*lfppb.LfpSale, error) {
	return c.internalClient.InsertLfpSale(ctx, req, opts...)
}

// lfpSaleGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type lfpSaleGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing LfpSaleClient
	CallOptions **LfpSaleCallOptions

	// The gRPC API client.
	lfpSaleClient lfppb.LfpSaleServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewLfpSaleClient creates a new lfp sale service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for a LFP
// partner (at https://support.google.com/merchants/answer/7676652) to submit sales
// data for a merchant.
func NewLfpSaleClient(ctx context.Context, opts ...option.ClientOption) (*LfpSaleClient, error) {
	clientOpts := defaultLfpSaleGRPCClientOptions()
	if newLfpSaleClientHook != nil {
		hookOpts, err := newLfpSaleClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := LfpSaleClient{CallOptions: defaultLfpSaleCallOptions()}

	c := &lfpSaleGRPCClient{
		connPool:      connPool,
		lfpSaleClient: lfppb.NewLfpSaleServiceClient(connPool),
		CallOptions:   &client.CallOptions,
		logger:        internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *lfpSaleGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *lfpSaleGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *lfpSaleGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type lfpSaleRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing LfpSaleClient
	CallOptions **LfpSaleCallOptions

	logger *slog.Logger
}

// NewLfpSaleRESTClient creates a new lfp sale service rest client.
//
// Service for a LFP
// partner (at https://support.google.com/merchants/answer/7676652) to submit sales
// data for a merchant.
func NewLfpSaleRESTClient(ctx context.Context, opts ...option.ClientOption) (*LfpSaleClient, error) {
	clientOpts := append(defaultLfpSaleRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultLfpSaleRESTCallOptions()
	c := &lfpSaleRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &LfpSaleClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultLfpSaleRESTClientOptions() []option.ClientOption {
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
func (c *lfpSaleRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *lfpSaleRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *lfpSaleRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *lfpSaleGRPCClient) InsertLfpSale(ctx context.Context, req *lfppb.InsertLfpSaleRequest, opts ...gax.CallOption) (*lfppb.LfpSale, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertLfpSale[0:len((*c.CallOptions).InsertLfpSale):len((*c.CallOptions).InsertLfpSale)], opts...)
	var resp *lfppb.LfpSale
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.lfpSaleClient.InsertLfpSale, req, settings.GRPC, c.logger, "InsertLfpSale")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// InsertLfpSale inserts a LfpSale for the given merchant.
func (c *lfpSaleRESTClient) InsertLfpSale(ctx context.Context, req *lfppb.InsertLfpSaleRequest, opts ...gax.CallOption) (*lfppb.LfpSale, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetLfpSale()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/lfp/v1beta/%v/lfpSales:insert", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertLfpSale[0:len((*c.CallOptions).InsertLfpSale):len((*c.CallOptions).InsertLfpSale)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &lfppb.LfpSale{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "InsertLfpSale")
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
