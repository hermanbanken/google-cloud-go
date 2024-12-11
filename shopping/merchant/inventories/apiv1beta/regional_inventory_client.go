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

package inventories

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	inventoriespb "cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionalInventoryClientHook clientHook

// RegionalInventoryCallOptions contains the retry settings for each method of RegionalInventoryClient.
type RegionalInventoryCallOptions struct {
	ListRegionalInventories []gax.CallOption
	InsertRegionalInventory []gax.CallOption
	DeleteRegionalInventory []gax.CallOption
}

func defaultRegionalInventoryGRPCClientOptions() []option.ClientOption {
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

func defaultRegionalInventoryCallOptions() *RegionalInventoryCallOptions {
	return &RegionalInventoryCallOptions{
		ListRegionalInventories: []gax.CallOption{
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
		InsertRegionalInventory: []gax.CallOption{
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
		DeleteRegionalInventory: []gax.CallOption{
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

func defaultRegionalInventoryRESTCallOptions() *RegionalInventoryCallOptions {
	return &RegionalInventoryCallOptions{
		ListRegionalInventories: []gax.CallOption{
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
		InsertRegionalInventory: []gax.CallOption{
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
		DeleteRegionalInventory: []gax.CallOption{
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

// internalRegionalInventoryClient is an interface that defines the methods available from Merchant API.
type internalRegionalInventoryClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListRegionalInventories(context.Context, *inventoriespb.ListRegionalInventoriesRequest, ...gax.CallOption) *RegionalInventoryIterator
	InsertRegionalInventory(context.Context, *inventoriespb.InsertRegionalInventoryRequest, ...gax.CallOption) (*inventoriespb.RegionalInventory, error)
	DeleteRegionalInventory(context.Context, *inventoriespb.DeleteRegionalInventoryRequest, ...gax.CallOption) error
}

// RegionalInventoryClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage regional inventory for products. There is also separate
// regions resource and API to manage regions definitions.
type RegionalInventoryClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionalInventoryClient

	// The call options for this service.
	CallOptions *RegionalInventoryCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionalInventoryClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionalInventoryClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionalInventoryClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListRegionalInventories lists the RegionalInventory resources for the given product in your
// merchant account. The response might contain fewer items than specified by
// pageSize.  If pageToken was returned in previous request, it can be
// used to obtain additional results.
//
// RegionalInventory resources are listed per product for a given account.
func (c *RegionalInventoryClient) ListRegionalInventories(ctx context.Context, req *inventoriespb.ListRegionalInventoriesRequest, opts ...gax.CallOption) *RegionalInventoryIterator {
	return c.internalClient.ListRegionalInventories(ctx, req, opts...)
}

// InsertRegionalInventory inserts a RegionalInventory to a given product in your
// merchant account.
//
// Replaces the full RegionalInventory resource if an entry with the same
// [region][google.shopping.merchant.inventories.v1beta.RegionalInventory.region]
// already exists for the product.
//
// It might take up to 30 minutes for the new or updated RegionalInventory
// resource to appear in products.
func (c *RegionalInventoryClient) InsertRegionalInventory(ctx context.Context, req *inventoriespb.InsertRegionalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.RegionalInventory, error) {
	return c.internalClient.InsertRegionalInventory(ctx, req, opts...)
}

// DeleteRegionalInventory deletes the specified RegionalInventory resource from the given product
// in your merchant account.  It might take up to an hour for the
// RegionalInventory to be deleted from the specific product.
// Once you have received a successful delete response, wait for that
// period before attempting a delete again.
func (c *RegionalInventoryClient) DeleteRegionalInventory(ctx context.Context, req *inventoriespb.DeleteRegionalInventoryRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteRegionalInventory(ctx, req, opts...)
}

// regionalInventoryGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionalInventoryGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionalInventoryClient
	CallOptions **RegionalInventoryCallOptions

	// The gRPC API client.
	regionalInventoryClient inventoriespb.RegionalInventoryServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewRegionalInventoryClient creates a new regional inventory service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage regional inventory for products. There is also separate
// regions resource and API to manage regions definitions.
func NewRegionalInventoryClient(ctx context.Context, opts ...option.ClientOption) (*RegionalInventoryClient, error) {
	clientOpts := defaultRegionalInventoryGRPCClientOptions()
	if newRegionalInventoryClientHook != nil {
		hookOpts, err := newRegionalInventoryClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionalInventoryClient{CallOptions: defaultRegionalInventoryCallOptions()}

	c := &regionalInventoryGRPCClient{
		connPool:                connPool,
		regionalInventoryClient: inventoriespb.NewRegionalInventoryServiceClient(connPool),
		CallOptions:             &client.CallOptions,
		logger:                  internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *regionalInventoryGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionalInventoryGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionalInventoryGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionalInventoryRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing RegionalInventoryClient
	CallOptions **RegionalInventoryCallOptions

	logger *slog.Logger
}

// NewRegionalInventoryRESTClient creates a new regional inventory service rest client.
//
// Service to manage regional inventory for products. There is also separate
// regions resource and API to manage regions definitions.
func NewRegionalInventoryRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionalInventoryClient, error) {
	clientOpts := append(defaultRegionalInventoryRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRegionalInventoryRESTCallOptions()
	c := &regionalInventoryRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &RegionalInventoryClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRegionalInventoryRESTClientOptions() []option.ClientOption {
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
func (c *regionalInventoryRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionalInventoryRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *regionalInventoryRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *regionalInventoryGRPCClient) ListRegionalInventories(ctx context.Context, req *inventoriespb.ListRegionalInventoriesRequest, opts ...gax.CallOption) *RegionalInventoryIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListRegionalInventories[0:len((*c.CallOptions).ListRegionalInventories):len((*c.CallOptions).ListRegionalInventories)], opts...)
	it := &RegionalInventoryIterator{}
	req = proto.Clone(req).(*inventoriespb.ListRegionalInventoriesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventoriespb.RegionalInventory, string, error) {
		resp := &inventoriespb.ListRegionalInventoriesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.regionalInventoryClient.ListRegionalInventories, req, settings.GRPC, c.logger, "ListRegionalInventories")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRegionalInventories(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *regionalInventoryGRPCClient) InsertRegionalInventory(ctx context.Context, req *inventoriespb.InsertRegionalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.RegionalInventory, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertRegionalInventory[0:len((*c.CallOptions).InsertRegionalInventory):len((*c.CallOptions).InsertRegionalInventory)], opts...)
	var resp *inventoriespb.RegionalInventory
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.regionalInventoryClient.InsertRegionalInventory, req, settings.GRPC, c.logger, "InsertRegionalInventory")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionalInventoryGRPCClient) DeleteRegionalInventory(ctx context.Context, req *inventoriespb.DeleteRegionalInventoryRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteRegionalInventory[0:len((*c.CallOptions).DeleteRegionalInventory):len((*c.CallOptions).DeleteRegionalInventory)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.regionalInventoryClient.DeleteRegionalInventory, req, settings.GRPC, c.logger, "DeleteRegionalInventory")
		return err
	}, opts...)
	return err
}

// ListRegionalInventories lists the RegionalInventory resources for the given product in your
// merchant account. The response might contain fewer items than specified by
// pageSize.  If pageToken was returned in previous request, it can be
// used to obtain additional results.
//
// RegionalInventory resources are listed per product for a given account.
func (c *regionalInventoryRESTClient) ListRegionalInventories(ctx context.Context, req *inventoriespb.ListRegionalInventoriesRequest, opts ...gax.CallOption) *RegionalInventoryIterator {
	it := &RegionalInventoryIterator{}
	req = proto.Clone(req).(*inventoriespb.ListRegionalInventoriesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventoriespb.RegionalInventory, string, error) {
		resp := &inventoriespb.ListRegionalInventoriesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v/regionalInventories", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListRegionalInventories")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetRegionalInventories(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// InsertRegionalInventory inserts a RegionalInventory to a given product in your
// merchant account.
//
// Replaces the full RegionalInventory resource if an entry with the same
// [region][google.shopping.merchant.inventories.v1beta.RegionalInventory.region]
// already exists for the product.
//
// It might take up to 30 minutes for the new or updated RegionalInventory
// resource to appear in products.
func (c *regionalInventoryRESTClient) InsertRegionalInventory(ctx context.Context, req *inventoriespb.InsertRegionalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.RegionalInventory, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetRegionalInventory()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v/regionalInventories:insert", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertRegionalInventory[0:len((*c.CallOptions).InsertRegionalInventory):len((*c.CallOptions).InsertRegionalInventory)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &inventoriespb.RegionalInventory{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "InsertRegionalInventory")
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

// DeleteRegionalInventory deletes the specified RegionalInventory resource from the given product
// in your merchant account.  It might take up to an hour for the
// RegionalInventory to be deleted from the specific product.
// Once you have received a successful delete response, wait for that
// period before attempting a delete again.
func (c *regionalInventoryRESTClient) DeleteRegionalInventory(ctx context.Context, req *inventoriespb.DeleteRegionalInventoryRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteRegionalInventory")
		return err
	}, opts...)
}
