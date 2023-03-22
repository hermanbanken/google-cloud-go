// Copyright 2023 Google LLC
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

package appengine

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newServicesClientHook clientHook

// ServicesCallOptions contains the retry settings for each method of ServicesClient.
type ServicesCallOptions struct {
	ListServices  []gax.CallOption
	GetService    []gax.CallOption
	UpdateService []gax.CallOption
	DeleteService []gax.CallOption
}

func defaultServicesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServicesCallOptions() *ServicesCallOptions {
	return &ServicesCallOptions{
		ListServices:  []gax.CallOption{},
		GetService:    []gax.CallOption{},
		UpdateService: []gax.CallOption{},
		DeleteService: []gax.CallOption{},
	}
}

func defaultServicesRESTCallOptions() *ServicesCallOptions {
	return &ServicesCallOptions{
		ListServices:  []gax.CallOption{},
		GetService:    []gax.CallOption{},
		UpdateService: []gax.CallOption{},
		DeleteService: []gax.CallOption{},
	}
}

// internalServicesClient is an interface that defines the methods available from App Engine Admin API.
type internalServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListServices(context.Context, *appenginepb.ListServicesRequest, ...gax.CallOption) *ServiceIterator
	GetService(context.Context, *appenginepb.GetServiceRequest, ...gax.CallOption) (*appenginepb.Service, error)
	UpdateService(context.Context, *appenginepb.UpdateServiceRequest, ...gax.CallOption) (*UpdateServiceOperation, error)
	UpdateServiceOperation(name string) *UpdateServiceOperation
	DeleteService(context.Context, *appenginepb.DeleteServiceRequest, ...gax.CallOption) (*DeleteServiceOperation, error)
	DeleteServiceOperation(name string) *DeleteServiceOperation
}

// ServicesClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages services of an application.
type ServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalServicesClient

	// The call options for this service.
	CallOptions *ServicesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListServices lists all the services in the application.
func (c *ServicesClient) ListServices(ctx context.Context, req *appenginepb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	return c.internalClient.ListServices(ctx, req, opts...)
}

// GetService gets the current configuration of the specified service.
func (c *ServicesClient) GetService(ctx context.Context, req *appenginepb.GetServiceRequest, opts ...gax.CallOption) (*appenginepb.Service, error) {
	return c.internalClient.GetService(ctx, req, opts...)
}

// UpdateService updates the configuration of the specified service.
func (c *ServicesClient) UpdateService(ctx context.Context, req *appenginepb.UpdateServiceRequest, opts ...gax.CallOption) (*UpdateServiceOperation, error) {
	return c.internalClient.UpdateService(ctx, req, opts...)
}

// UpdateServiceOperation returns a new UpdateServiceOperation from a given name.
// The name must be that of a previously created UpdateServiceOperation, possibly from a different process.
func (c *ServicesClient) UpdateServiceOperation(name string) *UpdateServiceOperation {
	return c.internalClient.UpdateServiceOperation(name)
}

// DeleteService deletes the specified service and all enclosed versions.
func (c *ServicesClient) DeleteService(ctx context.Context, req *appenginepb.DeleteServiceRequest, opts ...gax.CallOption) (*DeleteServiceOperation, error) {
	return c.internalClient.DeleteService(ctx, req, opts...)
}

// DeleteServiceOperation returns a new DeleteServiceOperation from a given name.
// The name must be that of a previously created DeleteServiceOperation, possibly from a different process.
func (c *ServicesClient) DeleteServiceOperation(name string) *DeleteServiceOperation {
	return c.internalClient.DeleteServiceOperation(name)
}

// servicesGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type servicesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ServicesClient
	CallOptions **ServicesCallOptions

	// The gRPC API client.
	servicesClient appenginepb.ServicesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServicesClient creates a new services client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages services of an application.
func NewServicesClient(ctx context.Context, opts ...option.ClientOption) (*ServicesClient, error) {
	clientOpts := defaultServicesGRPCClientOptions()
	if newServicesClientHook != nil {
		hookOpts, err := newServicesClientHook(ctx, clientHookParams{})
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
	client := ServicesClient{CallOptions: defaultServicesCallOptions()}

	c := &servicesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		servicesClient:   appenginepb.NewServicesClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *servicesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *servicesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *servicesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type servicesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ServicesClient
	CallOptions **ServicesCallOptions
}

// NewServicesRESTClient creates a new services rest client.
//
// Manages services of an application.
func NewServicesRESTClient(ctx context.Context, opts ...option.ClientOption) (*ServicesClient, error) {
	clientOpts := append(defaultServicesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultServicesRESTCallOptions()
	c := &servicesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &ServicesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultServicesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://appengine.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://appengine.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *servicesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *servicesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *servicesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *servicesGRPCClient) ListServices(ctx context.Context, req *appenginepb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListServices[0:len((*c.CallOptions).ListServices):len((*c.CallOptions).ListServices)], opts...)
	it := &ServiceIterator{}
	req = proto.Clone(req).(*appenginepb.ListServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Service, string, error) {
		resp := &appenginepb.ListServicesResponse{}
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
			resp, err = c.servicesClient.ListServices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServices(), resp.GetNextPageToken(), nil
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

func (c *servicesGRPCClient) GetService(ctx context.Context, req *appenginepb.GetServiceRequest, opts ...gax.CallOption) (*appenginepb.Service, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetService[0:len((*c.CallOptions).GetService):len((*c.CallOptions).GetService)], opts...)
	var resp *appenginepb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.GetService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servicesGRPCClient) UpdateService(ctx context.Context, req *appenginepb.UpdateServiceRequest, opts ...gax.CallOption) (*UpdateServiceOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateService[0:len((*c.CallOptions).UpdateService):len((*c.CallOptions).UpdateService)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.UpdateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *servicesGRPCClient) DeleteService(ctx context.Context, req *appenginepb.DeleteServiceRequest, opts ...gax.CallOption) (*DeleteServiceOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteService[0:len((*c.CallOptions).DeleteService):len((*c.CallOptions).DeleteService)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servicesClient.DeleteService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// ListServices lists all the services in the application.
func (c *servicesRESTClient) ListServices(ctx context.Context, req *appenginepb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	it := &ServiceIterator{}
	req = proto.Clone(req).(*appenginepb.ListServicesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Service, string, error) {
		resp := &appenginepb.ListServicesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/services", req.GetParent())

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
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetServices(), resp.GetNextPageToken(), nil
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

// GetService gets the current configuration of the specified service.
func (c *servicesRESTClient) GetService(ctx context.Context, req *appenginepb.GetServiceRequest, opts ...gax.CallOption) (*appenginepb.Service, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetService[0:len((*c.CallOptions).GetService):len((*c.CallOptions).GetService)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &appenginepb.Service{}
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// UpdateService updates the configuration of the specified service.
func (c *servicesRESTClient) UpdateService(ctx context.Context, req *appenginepb.UpdateServiceRequest, opts ...gax.CallOption) (*UpdateServiceOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetService()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetMigrateTraffic() {
		params.Add("migrateTraffic", fmt.Sprintf("%v", req.GetMigrateTraffic()))
	}
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &UpdateServiceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DeleteService deletes the specified service and all enclosed versions.
func (c *servicesRESTClient) DeleteService(ctx context.Context, req *appenginepb.DeleteServiceRequest, opts ...gax.CallOption) (*DeleteServiceOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DeleteServiceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DeleteServiceOperation manages a long-running operation from DeleteService.
type DeleteServiceOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// DeleteServiceOperation returns a new DeleteServiceOperation from a given name.
// The name must be that of a previously created DeleteServiceOperation, possibly from a different process.
func (c *servicesGRPCClient) DeleteServiceOperation(name string) *DeleteServiceOperation {
	return &DeleteServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteServiceOperation returns a new DeleteServiceOperation from a given name.
// The name must be that of a previously created DeleteServiceOperation, possibly from a different process.
func (c *servicesRESTClient) DeleteServiceOperation(name string) *DeleteServiceOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &DeleteServiceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteServiceOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteServiceOperation) Name() string {
	return op.lro.Name()
}

// UpdateServiceOperation manages a long-running operation from UpdateService.
type UpdateServiceOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// UpdateServiceOperation returns a new UpdateServiceOperation from a given name.
// The name must be that of a previously created UpdateServiceOperation, possibly from a different process.
func (c *servicesGRPCClient) UpdateServiceOperation(name string) *UpdateServiceOperation {
	return &UpdateServiceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateServiceOperation returns a new UpdateServiceOperation from a given name.
// The name must be that of a previously created UpdateServiceOperation, possibly from a different process.
func (c *servicesRESTClient) UpdateServiceOperation(name string) *UpdateServiceOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &UpdateServiceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp appenginepb.Service
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp appenginepb.Service
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateServiceOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateServiceOperation) Name() string {
	return op.lro.Name()
}

// ServiceIterator manages a stream of *appenginepb.Service.
type ServiceIterator struct {
	items    []*appenginepb.Service
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*appenginepb.Service, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceIterator) Next() (*appenginepb.Service, error) {
	var item *appenginepb.Service
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
