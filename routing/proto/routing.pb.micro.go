// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/routing.proto

package routing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Routing service

func NewRoutingEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Routing service

type RoutingService interface {
	// Route returns a gps route from origin to destination based on lat/lng
	Route(ctx context.Context, in *RouteRequest, opts ...client.CallOption) (*RouteResponse, error)
	// Eta returns an estimated time of arrival for a route
	Eta(ctx context.Context, in *EtaRequest, opts ...client.CallOption) (*EtaResponse, error)
}

type routingService struct {
	c    client.Client
	name string
}

func NewRoutingService(name string, c client.Client) RoutingService {
	return &routingService{
		c:    c,
		name: name,
	}
}

func (c *routingService) Route(ctx context.Context, in *RouteRequest, opts ...client.CallOption) (*RouteResponse, error) {
	req := c.c.NewRequest(c.name, "Routing.Route", in)
	out := new(RouteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingService) Eta(ctx context.Context, in *EtaRequest, opts ...client.CallOption) (*EtaResponse, error) {
	req := c.c.NewRequest(c.name, "Routing.Eta", in)
	out := new(EtaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Routing service

type RoutingHandler interface {
	// Route returns a gps route from origin to destination based on lat/lng
	Route(context.Context, *RouteRequest, *RouteResponse) error
	// Eta returns an estimated time of arrival for a route
	Eta(context.Context, *EtaRequest, *EtaResponse) error
}

func RegisterRoutingHandler(s server.Server, hdlr RoutingHandler, opts ...server.HandlerOption) error {
	type routing interface {
		Route(ctx context.Context, in *RouteRequest, out *RouteResponse) error
		Eta(ctx context.Context, in *EtaRequest, out *EtaResponse) error
	}
	type Routing struct {
		routing
	}
	h := &routingHandler{hdlr}
	return s.Handle(s.NewHandler(&Routing{h}, opts...))
}

type routingHandler struct {
	RoutingHandler
}

func (h *routingHandler) Route(ctx context.Context, in *RouteRequest, out *RouteResponse) error {
	return h.RoutingHandler.Route(ctx, in, out)
}

func (h *routingHandler) Eta(ctx context.Context, in *EtaRequest, out *EtaResponse) error {
	return h.RoutingHandler.Eta(ctx, in, out)
}
