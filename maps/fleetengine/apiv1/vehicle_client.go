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

package fleetengine

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"regexp"
	"strings"
	"time"

	fleetenginepb "cloud.google.com/go/maps/fleetengine/apiv1/fleetenginepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newVehicleClientHook clientHook

// VehicleCallOptions contains the retry settings for each method of VehicleClient.
type VehicleCallOptions struct {
	CreateVehicle           []gax.CallOption
	GetVehicle              []gax.CallOption
	UpdateVehicle           []gax.CallOption
	UpdateVehicleAttributes []gax.CallOption
	ListVehicles            []gax.CallOption
	SearchVehicles          []gax.CallOption
}

func defaultVehicleGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("fleetengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("fleetengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("fleetengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://fleetengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultVehicleCallOptions() *VehicleCallOptions {
	return &VehicleCallOptions{
		CreateVehicle: []gax.CallOption{
			gax.WithTimeout(15000 * time.Millisecond),
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
		GetVehicle: []gax.CallOption{
			gax.WithTimeout(15000 * time.Millisecond),
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
		UpdateVehicle: []gax.CallOption{
			gax.WithTimeout(15000 * time.Millisecond),
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
		UpdateVehicleAttributes: []gax.CallOption{
			gax.WithTimeout(15000 * time.Millisecond),
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
		ListVehicles: []gax.CallOption{},
		SearchVehicles: []gax.CallOption{
			gax.WithTimeout(15000 * time.Millisecond),
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

// internalVehicleClient is an interface that defines the methods available from Local Rides and Deliveries API.
type internalVehicleClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateVehicle(context.Context, *fleetenginepb.CreateVehicleRequest, ...gax.CallOption) (*fleetenginepb.Vehicle, error)
	GetVehicle(context.Context, *fleetenginepb.GetVehicleRequest, ...gax.CallOption) (*fleetenginepb.Vehicle, error)
	UpdateVehicle(context.Context, *fleetenginepb.UpdateVehicleRequest, ...gax.CallOption) (*fleetenginepb.Vehicle, error)
	UpdateVehicleAttributes(context.Context, *fleetenginepb.UpdateVehicleAttributesRequest, ...gax.CallOption) (*fleetenginepb.UpdateVehicleAttributesResponse, error)
	ListVehicles(context.Context, *fleetenginepb.ListVehiclesRequest, ...gax.CallOption) *VehicleIterator
	SearchVehicles(context.Context, *fleetenginepb.SearchVehiclesRequest, ...gax.CallOption) (*fleetenginepb.SearchVehiclesResponse, error)
}

// VehicleClient is a client for interacting with Local Rides and Deliveries API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Vehicle management service.
type VehicleClient struct {
	// The internal transport-dependent client.
	internalClient internalVehicleClient

	// The call options for this service.
	CallOptions *VehicleCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *VehicleClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *VehicleClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *VehicleClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateVehicle instantiates a new vehicle associated with an on-demand rideshare or
// deliveries provider. Each Vehicle must have a unique vehicle ID.
//
// The following Vehicle fields are required when creating a Vehicle:
//
//	vehicleState
//
//	supportedTripTypes
//
//	maximumCapacity
//
//	vehicleType
//
// The following Vehicle fields are ignored when creating a Vehicle:
//
//	name
//
//	currentTrips
//
//	availableCapacity
//
//	current_route_segment
//
//	current_route_segment_end_point
//
//	current_route_segment_version
//
//	current_route_segment_traffic
//
//	route
//
//	waypoints
//
//	waypoints_version
//
//	remaining_distance_meters
//
//	remaining_time_seconds
//
//	eta_to_next_waypoint
//
//	navigation_status
//
// All other fields are optional and used if provided.
func (c *VehicleClient) CreateVehicle(ctx context.Context, req *fleetenginepb.CreateVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	return c.internalClient.CreateVehicle(ctx, req, opts...)
}

// GetVehicle returns a vehicle from the Fleet Engine.
func (c *VehicleClient) GetVehicle(ctx context.Context, req *fleetenginepb.GetVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	return c.internalClient.GetVehicle(ctx, req, opts...)
}

// UpdateVehicle writes updated vehicle data to the Fleet Engine.
//
// When updating a Vehicle, the following fields cannot be updated since
// they are managed by the server:
//
//	currentTrips
//
//	availableCapacity
//
//	current_route_segment_version
//
//	waypoints_version
//
// The vehicle name also cannot be updated.
//
// If the attributes field is updated, all the vehicle’s attributes are
// replaced with the attributes provided in the request. If you want to update
// only some attributes, see the UpdateVehicleAttributes method. Likewise,
// the waypoints field can be updated, but must contain all the waypoints
// currently on the vehicle, and no other waypoints.
func (c *VehicleClient) UpdateVehicle(ctx context.Context, req *fleetenginepb.UpdateVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	return c.internalClient.UpdateVehicle(ctx, req, opts...)
}

// UpdateVehicleAttributes partially updates a vehicle’s attributes.
// Only the attributes mentioned in the request will be updated, other
// attributes will NOT be altered. Note: this is different in UpdateVehicle,
// where the whole attributes field will be replaced by the one in
// UpdateVehicleRequest, attributes not in the request would be removed.
func (c *VehicleClient) UpdateVehicleAttributes(ctx context.Context, req *fleetenginepb.UpdateVehicleAttributesRequest, opts ...gax.CallOption) (*fleetenginepb.UpdateVehicleAttributesResponse, error) {
	return c.internalClient.UpdateVehicleAttributes(ctx, req, opts...)
}

// ListVehicles returns a paginated list of vehicles associated with
// a provider that match the request options.
func (c *VehicleClient) ListVehicles(ctx context.Context, req *fleetenginepb.ListVehiclesRequest, opts ...gax.CallOption) *VehicleIterator {
	return c.internalClient.ListVehicles(ctx, req, opts...)
}

// SearchVehicles returns a list of vehicles that match the request options.
func (c *VehicleClient) SearchVehicles(ctx context.Context, req *fleetenginepb.SearchVehiclesRequest, opts ...gax.CallOption) (*fleetenginepb.SearchVehiclesResponse, error) {
	return c.internalClient.SearchVehicles(ctx, req, opts...)
}

// vehicleGRPCClient is a client for interacting with Local Rides and Deliveries API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type vehicleGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing VehicleClient
	CallOptions **VehicleCallOptions

	// The gRPC API client.
	vehicleClient fleetenginepb.VehicleServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewVehicleClient creates a new vehicle service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Vehicle management service.
func NewVehicleClient(ctx context.Context, opts ...option.ClientOption) (*VehicleClient, error) {
	clientOpts := defaultVehicleGRPCClientOptions()
	if newVehicleClientHook != nil {
		hookOpts, err := newVehicleClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := VehicleClient{CallOptions: defaultVehicleCallOptions()}

	c := &vehicleGRPCClient{
		connPool:      connPool,
		vehicleClient: fleetenginepb.NewVehicleServiceClient(connPool),
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
func (c *vehicleGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *vehicleGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *vehicleGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *vehicleGRPCClient) CreateVehicle(ctx context.Context, req *fleetenginepb.CreateVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetParent()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateVehicle[0:len((*c.CallOptions).CreateVehicle):len((*c.CallOptions).CreateVehicle)], opts...)
	var resp *fleetenginepb.Vehicle
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.vehicleClient.CreateVehicle, req, settings.GRPC, c.logger, "CreateVehicle")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vehicleGRPCClient) GetVehicle(ctx context.Context, req *fleetenginepb.GetVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetVehicle[0:len((*c.CallOptions).GetVehicle):len((*c.CallOptions).GetVehicle)], opts...)
	var resp *fleetenginepb.Vehicle
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.vehicleClient.GetVehicle, req, settings.GRPC, c.logger, "GetVehicle")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vehicleGRPCClient) UpdateVehicle(ctx context.Context, req *fleetenginepb.UpdateVehicleRequest, opts ...gax.CallOption) (*fleetenginepb.Vehicle, error) {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateVehicle[0:len((*c.CallOptions).UpdateVehicle):len((*c.CallOptions).UpdateVehicle)], opts...)
	var resp *fleetenginepb.Vehicle
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.vehicleClient.UpdateVehicle, req, settings.GRPC, c.logger, "UpdateVehicle")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vehicleGRPCClient) UpdateVehicleAttributes(ctx context.Context, req *fleetenginepb.UpdateVehicleAttributesRequest, opts ...gax.CallOption) (*fleetenginepb.UpdateVehicleAttributesResponse, error) {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetName()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetName())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateVehicleAttributes[0:len((*c.CallOptions).UpdateVehicleAttributes):len((*c.CallOptions).UpdateVehicleAttributes)], opts...)
	var resp *fleetenginepb.UpdateVehicleAttributesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.vehicleClient.UpdateVehicleAttributes, req, settings.GRPC, c.logger, "UpdateVehicleAttributes")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vehicleGRPCClient) ListVehicles(ctx context.Context, req *fleetenginepb.ListVehiclesRequest, opts ...gax.CallOption) *VehicleIterator {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetParent()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListVehicles[0:len((*c.CallOptions).ListVehicles):len((*c.CallOptions).ListVehicles)], opts...)
	it := &VehicleIterator{}
	req = proto.Clone(req).(*fleetenginepb.ListVehiclesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*fleetenginepb.Vehicle, string, error) {
		resp := &fleetenginepb.ListVehiclesResponse{}
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
			resp, err = executeRPC(ctx, c.vehicleClient.ListVehicles, req, settings.GRPC, c.logger, "ListVehicles")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetVehicles(), resp.GetNextPageToken(), nil
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

func (c *vehicleGRPCClient) SearchVehicles(ctx context.Context, req *fleetenginepb.SearchVehiclesRequest, opts ...gax.CallOption) (*fleetenginepb.SearchVehiclesResponse, error) {
	routingHeaders := ""
	routingHeadersMap := make(map[string]string)
	if reg := regexp.MustCompile("(?P<provider_id>providers/[^/]+)"); reg.MatchString(req.GetParent()) && len(url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])) > 0 {
		routingHeadersMap["provider_id"] = url.QueryEscape(reg.FindStringSubmatch(req.GetParent())[1])
	}
	for headerName, headerValue := range routingHeadersMap {
		routingHeaders = fmt.Sprintf("%s%s=%s&", routingHeaders, headerName, headerValue)
	}
	routingHeaders = strings.TrimSuffix(routingHeaders, "&")
	hds := []string{"x-goog-request-params", routingHeaders}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SearchVehicles[0:len((*c.CallOptions).SearchVehicles):len((*c.CallOptions).SearchVehicles)], opts...)
	var resp *fleetenginepb.SearchVehiclesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.vehicleClient.SearchVehicles, req, settings.GRPC, c.logger, "SearchVehicles")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
