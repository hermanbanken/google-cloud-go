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

package meet

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	meetpb "cloud.google.com/go/apps/meet/apiv2beta/meetpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newSpacesClientHook clientHook

// SpacesCallOptions contains the retry settings for each method of SpacesClient.
type SpacesCallOptions struct {
	CreateSpace         []gax.CallOption
	GetSpace            []gax.CallOption
	UpdateSpace         []gax.CallOption
	EndActiveConference []gax.CallOption
}

func defaultSpacesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("meet.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("meet.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("meet.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://meet.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSpacesCallOptions() *SpacesCallOptions {
	return &SpacesCallOptions{
		CreateSpace: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetSpace: []gax.CallOption{
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
		UpdateSpace: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		EndActiveConference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultSpacesRESTCallOptions() *SpacesCallOptions {
	return &SpacesCallOptions{
		CreateSpace: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetSpace: []gax.CallOption{
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
		UpdateSpace: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		EndActiveConference: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalSpacesClient is an interface that defines the methods available from Google Meet API.
type internalSpacesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSpace(context.Context, *meetpb.CreateSpaceRequest, ...gax.CallOption) (*meetpb.Space, error)
	GetSpace(context.Context, *meetpb.GetSpaceRequest, ...gax.CallOption) (*meetpb.Space, error)
	UpdateSpace(context.Context, *meetpb.UpdateSpaceRequest, ...gax.CallOption) (*meetpb.Space, error)
	EndActiveConference(context.Context, *meetpb.EndActiveConferenceRequest, ...gax.CallOption) error
}

// SpacesClient is a client for interacting with Google Meet API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// REST API for services dealing with spaces.
type SpacesClient struct {
	// The internal transport-dependent client.
	internalClient internalSpacesClient

	// The call options for this service.
	CallOptions *SpacesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SpacesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SpacesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SpacesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Creates a space.
func (c *SpacesClient) CreateSpace(ctx context.Context, req *meetpb.CreateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	return c.internalClient.CreateSpace(ctx, req, opts...)
}

// GetSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Gets a space by space_id or meeting_code.
func (c *SpacesClient) GetSpace(ctx context.Context, req *meetpb.GetSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	return c.internalClient.GetSpace(ctx, req, opts...)
}

// UpdateSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Updates a space.
func (c *SpacesClient) UpdateSpace(ctx context.Context, req *meetpb.UpdateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	return c.internalClient.UpdateSpace(ctx, req, opts...)
}

// EndActiveConference Developer Preview (at https://developers.google.com/workspace/preview).
// Ends an active conference (if there is one).
func (c *SpacesClient) EndActiveConference(ctx context.Context, req *meetpb.EndActiveConferenceRequest, opts ...gax.CallOption) error {
	return c.internalClient.EndActiveConference(ctx, req, opts...)
}

// spacesGRPCClient is a client for interacting with Google Meet API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type spacesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SpacesClient
	CallOptions **SpacesCallOptions

	// The gRPC API client.
	spacesClient meetpb.SpacesServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewSpacesClient creates a new spaces service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// REST API for services dealing with spaces.
func NewSpacesClient(ctx context.Context, opts ...option.ClientOption) (*SpacesClient, error) {
	clientOpts := defaultSpacesGRPCClientOptions()
	if newSpacesClientHook != nil {
		hookOpts, err := newSpacesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SpacesClient{CallOptions: defaultSpacesCallOptions()}

	c := &spacesGRPCClient{
		connPool:     connPool,
		spacesClient: meetpb.NewSpacesServiceClient(connPool),
		CallOptions:  &client.CallOptions,
		logger:       internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *spacesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *spacesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *spacesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type spacesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing SpacesClient
	CallOptions **SpacesCallOptions

	logger *slog.Logger
}

// NewSpacesRESTClient creates a new spaces service rest client.
//
// REST API for services dealing with spaces.
func NewSpacesRESTClient(ctx context.Context, opts ...option.ClientOption) (*SpacesClient, error) {
	clientOpts := append(defaultSpacesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultSpacesRESTCallOptions()
	c := &spacesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &SpacesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultSpacesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://meet.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://meet.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://meet.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://meet.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *spacesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *spacesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *spacesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *spacesGRPCClient) CreateSpace(ctx context.Context, req *meetpb.CreateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateSpace[0:len((*c.CallOptions).CreateSpace):len((*c.CallOptions).CreateSpace)], opts...)
	var resp *meetpb.Space
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.spacesClient.CreateSpace, req, settings.GRPC, c.logger, "CreateSpace")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *spacesGRPCClient) GetSpace(ctx context.Context, req *meetpb.GetSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSpace[0:len((*c.CallOptions).GetSpace):len((*c.CallOptions).GetSpace)], opts...)
	var resp *meetpb.Space
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.spacesClient.GetSpace, req, settings.GRPC, c.logger, "GetSpace")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *spacesGRPCClient) UpdateSpace(ctx context.Context, req *meetpb.UpdateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "space.name", url.QueryEscape(req.GetSpace().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateSpace[0:len((*c.CallOptions).UpdateSpace):len((*c.CallOptions).UpdateSpace)], opts...)
	var resp *meetpb.Space
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.spacesClient.UpdateSpace, req, settings.GRPC, c.logger, "UpdateSpace")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *spacesGRPCClient) EndActiveConference(ctx context.Context, req *meetpb.EndActiveConferenceRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).EndActiveConference[0:len((*c.CallOptions).EndActiveConference):len((*c.CallOptions).EndActiveConference)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.spacesClient.EndActiveConference, req, settings.GRPC, c.logger, "EndActiveConference")
		return err
	}, opts...)
	return err
}

// CreateSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Creates a space.
func (c *spacesRESTClient) CreateSpace(ctx context.Context, req *meetpb.CreateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetSpace()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta/spaces")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateSpace[0:len((*c.CallOptions).CreateSpace):len((*c.CallOptions).CreateSpace)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &meetpb.Space{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateSpace")
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

// GetSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Gets a space by space_id or meeting_code.
func (c *spacesRESTClient) GetSpace(ctx context.Context, req *meetpb.GetSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetSpace[0:len((*c.CallOptions).GetSpace):len((*c.CallOptions).GetSpace)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &meetpb.Space{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetSpace")
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

// UpdateSpace Developer Preview (at https://developers.google.com/workspace/preview).
// Updates a space.
func (c *spacesRESTClient) UpdateSpace(ctx context.Context, req *meetpb.UpdateSpaceRequest, opts ...gax.CallOption) (*meetpb.Space, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetSpace()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta/%v", req.GetSpace().GetName())

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
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "space.name", url.QueryEscape(req.GetSpace().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateSpace[0:len((*c.CallOptions).UpdateSpace):len((*c.CallOptions).UpdateSpace)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &meetpb.Space{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateSpace")
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

// EndActiveConference Developer Preview (at https://developers.google.com/workspace/preview).
// Ends an active conference (if there is one).
func (c *spacesRESTClient) EndActiveConference(ctx context.Context, req *meetpb.EndActiveConferenceRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v2beta/%v:endActiveConference", req.GetName())

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
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "EndActiveConference")
		return err
	}, opts...)
}
