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

package appengine

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newInstancesClientHook clientHook

// InstancesCallOptions contains the retry settings for each method of InstancesClient.
type InstancesCallOptions struct {
	ListInstances  []gax.CallOption
	GetInstance    []gax.CallOption
	DeleteInstance []gax.CallOption
	DebugInstance  []gax.CallOption
}

func defaultInstancesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("appengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultInstancesCallOptions() *InstancesCallOptions {
	return &InstancesCallOptions{
		ListInstances: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DebugInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultInstancesRESTCallOptions() *InstancesCallOptions {
	return &InstancesCallOptions{
		ListInstances: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DebugInstance: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalInstancesClient is an interface that defines the methods available from App Engine Admin API.
type internalInstancesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListInstances(context.Context, *appenginepb.ListInstancesRequest, ...gax.CallOption) *InstanceIterator
	GetInstance(context.Context, *appenginepb.GetInstanceRequest, ...gax.CallOption) (*appenginepb.Instance, error)
	DeleteInstance(context.Context, *appenginepb.DeleteInstanceRequest, ...gax.CallOption) (*DeleteInstanceOperation, error)
	DeleteInstanceOperation(name string) *DeleteInstanceOperation
	DebugInstance(context.Context, *appenginepb.DebugInstanceRequest, ...gax.CallOption) (*DebugInstanceOperation, error)
	DebugInstanceOperation(name string) *DebugInstanceOperation
}

// InstancesClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages instances of a version.
type InstancesClient struct {
	// The internal transport-dependent client.
	internalClient internalInstancesClient

	// The call options for this service.
	CallOptions *InstancesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InstancesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InstancesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InstancesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListInstances lists the instances of a version.
//
// Tip: To aggregate details about instances over time, see the
// Stackdriver Monitoring API (at https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
func (c *InstancesClient) ListInstances(ctx context.Context, req *appenginepb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	return c.internalClient.ListInstances(ctx, req, opts...)
}

// GetInstance gets instance information.
func (c *InstancesClient) GetInstance(ctx context.Context, req *appenginepb.GetInstanceRequest, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	return c.internalClient.GetInstance(ctx, req, opts...)
}

// DeleteInstance stops a running instance.
//
// The instance might be automatically recreated based on the scaling settings
// of the version. For more information, see “How Instances are Managed”
// (standard environment (at https://cloud.google.com/appengine/docs/standard/python/how-instances-are-managed) |
// flexible environment (at https://cloud.google.com/appengine/docs/flexible/python/how-instances-are-managed)).
//
// To ensure that instances are not re-created and avoid getting billed, you
// can stop all instances within the target version by changing the serving
// status of the version to STOPPED with the
// apps.services.versions.patch (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions/patch)
// method.
func (c *InstancesClient) DeleteInstance(ctx context.Context, req *appenginepb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	return c.internalClient.DeleteInstance(ctx, req, opts...)
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *InstancesClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	return c.internalClient.DeleteInstanceOperation(name)
}

// DebugInstance enables debugging on a VM instance. This allows you to use the SSH
// command to connect to the virtual machine where the instance lives.
// While in “debug mode”, the instance continues to serve live traffic.
// You should delete the instance when you are done debugging and then
// allow the system to take over and determine if another instance
// should be started.
//
// Only applicable for instances in App Engine flexible environment.
func (c *InstancesClient) DebugInstance(ctx context.Context, req *appenginepb.DebugInstanceRequest, opts ...gax.CallOption) (*DebugInstanceOperation, error) {
	return c.internalClient.DebugInstance(ctx, req, opts...)
}

// DebugInstanceOperation returns a new DebugInstanceOperation from a given name.
// The name must be that of a previously created DebugInstanceOperation, possibly from a different process.
func (c *InstancesClient) DebugInstanceOperation(name string) *DebugInstanceOperation {
	return c.internalClient.DebugInstanceOperation(name)
}

// instancesGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type instancesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing InstancesClient
	CallOptions **InstancesCallOptions

	// The gRPC API client.
	instancesClient appenginepb.InstancesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewInstancesClient creates a new instances client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages instances of a version.
func NewInstancesClient(ctx context.Context, opts ...option.ClientOption) (*InstancesClient, error) {
	clientOpts := defaultInstancesGRPCClientOptions()
	if newInstancesClientHook != nil {
		hookOpts, err := newInstancesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InstancesClient{CallOptions: defaultInstancesCallOptions()}

	c := &instancesGRPCClient{
		connPool:        connPool,
		instancesClient: appenginepb.NewInstancesClient(connPool),
		CallOptions:     &client.CallOptions,
		logger:          internaloption.GetLogger(opts),
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
func (c *instancesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *instancesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *instancesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type instancesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing InstancesClient
	CallOptions **InstancesCallOptions

	logger *slog.Logger
}

// NewInstancesRESTClient creates a new instances rest client.
//
// Manages instances of a version.
func NewInstancesRESTClient(ctx context.Context, opts ...option.ClientOption) (*InstancesClient, error) {
	clientOpts := append(defaultInstancesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultInstancesRESTCallOptions()
	c := &instancesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
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

	return &InstancesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultInstancesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://appengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://appengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://appengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *instancesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *instancesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *instancesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *instancesGRPCClient) ListInstances(ctx context.Context, req *appenginepb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListInstances[0:len((*c.CallOptions).ListInstances):len((*c.CallOptions).ListInstances)], opts...)
	it := &InstanceIterator{}
	req = proto.Clone(req).(*appenginepb.ListInstancesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Instance, string, error) {
		resp := &appenginepb.ListInstancesResponse{}
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
			resp, err = executeRPC(ctx, c.instancesClient.ListInstances, req, settings.GRPC, c.logger, "ListInstances")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetInstances(), resp.GetNextPageToken(), nil
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

func (c *instancesGRPCClient) GetInstance(ctx context.Context, req *appenginepb.GetInstanceRequest, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetInstance[0:len((*c.CallOptions).GetInstance):len((*c.CallOptions).GetInstance)], opts...)
	var resp *appenginepb.Instance
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.instancesClient.GetInstance, req, settings.GRPC, c.logger, "GetInstance")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instancesGRPCClient) DeleteInstance(ctx context.Context, req *appenginepb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteInstance[0:len((*c.CallOptions).DeleteInstance):len((*c.CallOptions).DeleteInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.instancesClient.DeleteInstance, req, settings.GRPC, c.logger, "DeleteInstance")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *instancesGRPCClient) DebugInstance(ctx context.Context, req *appenginepb.DebugInstanceRequest, opts ...gax.CallOption) (*DebugInstanceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DebugInstance[0:len((*c.CallOptions).DebugInstance):len((*c.CallOptions).DebugInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.instancesClient.DebugInstance, req, settings.GRPC, c.logger, "DebugInstance")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DebugInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// ListInstances lists the instances of a version.
//
// Tip: To aggregate details about instances over time, see the
// Stackdriver Monitoring API (at https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
func (c *instancesRESTClient) ListInstances(ctx context.Context, req *appenginepb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	it := &InstanceIterator{}
	req = proto.Clone(req).(*appenginepb.ListInstancesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Instance, string, error) {
		resp := &appenginepb.ListInstancesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/instances", req.GetParent())

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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListInstances")
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
		return resp.GetInstances(), resp.GetNextPageToken(), nil
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

// GetInstance gets instance information.
func (c *instancesRESTClient) GetInstance(ctx context.Context, req *appenginepb.GetInstanceRequest, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetInstance[0:len((*c.CallOptions).GetInstance):len((*c.CallOptions).GetInstance)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &appenginepb.Instance{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetInstance")
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

// DeleteInstance stops a running instance.
//
// The instance might be automatically recreated based on the scaling settings
// of the version. For more information, see “How Instances are Managed”
// (standard environment (at https://cloud.google.com/appengine/docs/standard/python/how-instances-are-managed) |
// flexible environment (at https://cloud.google.com/appengine/docs/flexible/python/how-instances-are-managed)).
//
// To ensure that instances are not re-created and avoid getting billed, you
// can stop all instances within the target version by changing the serving
// status of the version to STOPPED with the
// apps.services.versions.patch (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions/patch)
// method.
func (c *instancesRESTClient) DeleteInstance(ctx context.Context, req *appenginepb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteInstance")
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DeleteInstanceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DebugInstance enables debugging on a VM instance. This allows you to use the SSH
// command to connect to the virtual machine where the instance lives.
// While in “debug mode”, the instance continues to serve live traffic.
// You should delete the instance when you are done debugging and then
// allow the system to take over and determine if another instance
// should be started.
//
// Only applicable for instances in App Engine flexible environment.
func (c *instancesRESTClient) DebugInstance(ctx context.Context, req *appenginepb.DebugInstanceRequest, opts ...gax.CallOption) (*DebugInstanceOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:debug", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "DebugInstance")
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DebugInstanceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DebugInstanceOperation returns a new DebugInstanceOperation from a given name.
// The name must be that of a previously created DebugInstanceOperation, possibly from a different process.
func (c *instancesGRPCClient) DebugInstanceOperation(name string) *DebugInstanceOperation {
	return &DebugInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DebugInstanceOperation returns a new DebugInstanceOperation from a given name.
// The name must be that of a previously created DebugInstanceOperation, possibly from a different process.
func (c *instancesRESTClient) DebugInstanceOperation(name string) *DebugInstanceOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &DebugInstanceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *instancesGRPCClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *instancesRESTClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &DeleteInstanceOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
