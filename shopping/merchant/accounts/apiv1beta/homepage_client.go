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
	"time"

	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newHomepageClientHook clientHook

// HomepageCallOptions contains the retry settings for each method of HomepageClient.
type HomepageCallOptions struct {
	GetHomepage     []gax.CallOption
	UpdateHomepage  []gax.CallOption
	ClaimHomepage   []gax.CallOption
	UnclaimHomepage []gax.CallOption
}

func defaultHomepageGRPCClientOptions() []option.ClientOption {
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

func defaultHomepageCallOptions() *HomepageCallOptions {
	return &HomepageCallOptions{
		GetHomepage: []gax.CallOption{
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
		UpdateHomepage: []gax.CallOption{
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
		ClaimHomepage: []gax.CallOption{
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
		UnclaimHomepage: []gax.CallOption{
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

func defaultHomepageRESTCallOptions() *HomepageCallOptions {
	return &HomepageCallOptions{
		GetHomepage: []gax.CallOption{
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
		UpdateHomepage: []gax.CallOption{
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
		ClaimHomepage: []gax.CallOption{
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
		UnclaimHomepage: []gax.CallOption{
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

// internalHomepageClient is an interface that defines the methods available from Merchant API.
type internalHomepageClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetHomepage(context.Context, *accountspb.GetHomepageRequest, ...gax.CallOption) (*accountspb.Homepage, error)
	UpdateHomepage(context.Context, *accountspb.UpdateHomepageRequest, ...gax.CallOption) (*accountspb.Homepage, error)
	ClaimHomepage(context.Context, *accountspb.ClaimHomepageRequest, ...gax.CallOption) (*accountspb.Homepage, error)
	UnclaimHomepage(context.Context, *accountspb.UnclaimHomepageRequest, ...gax.CallOption) (*accountspb.Homepage, error)
}

// HomepageClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to support an API for a store’s homepage.
type HomepageClient struct {
	// The internal transport-dependent client.
	internalClient internalHomepageClient

	// The call options for this service.
	CallOptions *HomepageCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *HomepageClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *HomepageClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *HomepageClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetHomepage retrieves a store’s homepage.
func (c *HomepageClient) GetHomepage(ctx context.Context, req *accountspb.GetHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	return c.internalClient.GetHomepage(ctx, req, opts...)
}

// UpdateHomepage updates a store’s homepage. Executing this method requires admin access.
func (c *HomepageClient) UpdateHomepage(ctx context.Context, req *accountspb.UpdateHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	return c.internalClient.UpdateHomepage(ctx, req, opts...)
}

// ClaimHomepage claims a store’s homepage. Executing this method requires admin access.
//
// If the homepage is already claimed, this will recheck the
// verification (unless the merchant is exempted from claiming, which also
// exempts from verification) and return a successful response. If ownership
// can no longer be verified, it will return an error, but it won’t clear the
// claim. In case of failure, a canonical error message will be returned:
// * PERMISSION_DENIED: user doesn’t have the necessary permissions on this
// MC account;
// * FAILED_PRECONDITION:
// - The account is not a Merchant Center account;
// - MC account doesn’t have a homepage;
// - claiming failed (in this case the error message will contain more
// details).
func (c *HomepageClient) ClaimHomepage(ctx context.Context, req *accountspb.ClaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	return c.internalClient.ClaimHomepage(ctx, req, opts...)
}

// UnclaimHomepage unclaims a store’s homepage. Executing this method requires admin access.
func (c *HomepageClient) UnclaimHomepage(ctx context.Context, req *accountspb.UnclaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	return c.internalClient.UnclaimHomepage(ctx, req, opts...)
}

// homepageGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type homepageGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing HomepageClient
	CallOptions **HomepageCallOptions

	// The gRPC API client.
	homepageClient accountspb.HomepageServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewHomepageClient creates a new homepage service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to support an API for a store’s homepage.
func NewHomepageClient(ctx context.Context, opts ...option.ClientOption) (*HomepageClient, error) {
	clientOpts := defaultHomepageGRPCClientOptions()
	if newHomepageClientHook != nil {
		hookOpts, err := newHomepageClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := HomepageClient{CallOptions: defaultHomepageCallOptions()}

	c := &homepageGRPCClient{
		connPool:       connPool,
		homepageClient: accountspb.NewHomepageServiceClient(connPool),
		CallOptions:    &client.CallOptions,
		logger:         internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *homepageGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *homepageGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *homepageGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type homepageRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing HomepageClient
	CallOptions **HomepageCallOptions

	logger *slog.Logger
}

// NewHomepageRESTClient creates a new homepage service rest client.
//
// Service to support an API for a store’s homepage.
func NewHomepageRESTClient(ctx context.Context, opts ...option.ClientOption) (*HomepageClient, error) {
	clientOpts := append(defaultHomepageRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultHomepageRESTCallOptions()
	c := &homepageRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &HomepageClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultHomepageRESTClientOptions() []option.ClientOption {
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
func (c *homepageRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *homepageRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *homepageRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *homepageGRPCClient) GetHomepage(ctx context.Context, req *accountspb.GetHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetHomepage[0:len((*c.CallOptions).GetHomepage):len((*c.CallOptions).GetHomepage)], opts...)
	var resp *accountspb.Homepage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.homepageClient.GetHomepage, req, settings.GRPC, c.logger, "GetHomepage")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *homepageGRPCClient) UpdateHomepage(ctx context.Context, req *accountspb.UpdateHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "homepage.name", url.QueryEscape(req.GetHomepage().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateHomepage[0:len((*c.CallOptions).UpdateHomepage):len((*c.CallOptions).UpdateHomepage)], opts...)
	var resp *accountspb.Homepage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.homepageClient.UpdateHomepage, req, settings.GRPC, c.logger, "UpdateHomepage")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *homepageGRPCClient) ClaimHomepage(ctx context.Context, req *accountspb.ClaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ClaimHomepage[0:len((*c.CallOptions).ClaimHomepage):len((*c.CallOptions).ClaimHomepage)], opts...)
	var resp *accountspb.Homepage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.homepageClient.ClaimHomepage, req, settings.GRPC, c.logger, "ClaimHomepage")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *homepageGRPCClient) UnclaimHomepage(ctx context.Context, req *accountspb.UnclaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UnclaimHomepage[0:len((*c.CallOptions).UnclaimHomepage):len((*c.CallOptions).UnclaimHomepage)], opts...)
	var resp *accountspb.Homepage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.homepageClient.UnclaimHomepage, req, settings.GRPC, c.logger, "UnclaimHomepage")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetHomepage retrieves a store’s homepage.
func (c *homepageRESTClient) GetHomepage(ctx context.Context, req *accountspb.GetHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
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
	opts = append((*c.CallOptions).GetHomepage[0:len((*c.CallOptions).GetHomepage):len((*c.CallOptions).GetHomepage)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.Homepage{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetHomepage")
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

// UpdateHomepage updates a store’s homepage. Executing this method requires admin access.
func (c *homepageRESTClient) UpdateHomepage(ctx context.Context, req *accountspb.UpdateHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetHomepage()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetHomepage().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "homepage.name", url.QueryEscape(req.GetHomepage().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateHomepage[0:len((*c.CallOptions).UpdateHomepage):len((*c.CallOptions).UpdateHomepage)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.Homepage{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateHomepage")
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

// ClaimHomepage claims a store’s homepage. Executing this method requires admin access.
//
// If the homepage is already claimed, this will recheck the
// verification (unless the merchant is exempted from claiming, which also
// exempts from verification) and return a successful response. If ownership
// can no longer be verified, it will return an error, but it won’t clear the
// claim. In case of failure, a canonical error message will be returned:
// * PERMISSION_DENIED: user doesn’t have the necessary permissions on this
// MC account;
// * FAILED_PRECONDITION:
// - The account is not a Merchant Center account;
// - MC account doesn’t have a homepage;
// - claiming failed (in this case the error message will contain more
// details).
func (c *homepageRESTClient) ClaimHomepage(ctx context.Context, req *accountspb.ClaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v:claim", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ClaimHomepage[0:len((*c.CallOptions).ClaimHomepage):len((*c.CallOptions).ClaimHomepage)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.Homepage{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ClaimHomepage")
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

// UnclaimHomepage unclaims a store’s homepage. Executing this method requires admin access.
func (c *homepageRESTClient) UnclaimHomepage(ctx context.Context, req *accountspb.UnclaimHomepageRequest, opts ...gax.CallOption) (*accountspb.Homepage, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v:unclaim", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UnclaimHomepage[0:len((*c.CallOptions).UnclaimHomepage):len((*c.CallOptions).UnclaimHomepage)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.Homepage{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UnclaimHomepage")
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
