// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Login(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Login(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Login(ctx context.Context, in *go_api.Request, out *go_api.Response) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Login(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.UserHandler.Login(ctx, in, out)
}