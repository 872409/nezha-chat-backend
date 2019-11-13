// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Post(ctx context.Context, in *PostReq, opts ...client.CallOption) (*PostResp, error)
	CheckPassword(ctx context.Context, in *CheckPasswordReq, opts ...client.CallOption) (*CheckPasswordResp, error)
	GetList(ctx context.Context, in *GetListReq, opts ...client.CallOption) (*GetListResp, error)
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
		name = "nezha.chat.user.srv.service"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Post(ctx context.Context, in *PostReq, opts ...client.CallOption) (*PostResp, error) {
	req := c.c.NewRequest(c.name, "User.Post", in)
	out := new(PostResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CheckPassword(ctx context.Context, in *CheckPasswordReq, opts ...client.CallOption) (*CheckPasswordResp, error) {
	req := c.c.NewRequest(c.name, "User.CheckPassword", in)
	out := new(CheckPasswordResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetList(ctx context.Context, in *GetListReq, opts ...client.CallOption) (*GetListResp, error) {
	req := c.c.NewRequest(c.name, "User.GetList", in)
	out := new(GetListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Post(context.Context, *PostReq, *PostResp) error
	CheckPassword(context.Context, *CheckPasswordReq, *CheckPasswordResp) error
	GetList(context.Context, *GetListReq, *GetListResp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Post(ctx context.Context, in *PostReq, out *PostResp) error
		CheckPassword(ctx context.Context, in *CheckPasswordReq, out *CheckPasswordResp) error
		GetList(ctx context.Context, in *GetListReq, out *GetListResp) error
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

func (h *userHandler) Post(ctx context.Context, in *PostReq, out *PostResp) error {
	return h.UserHandler.Post(ctx, in, out)
}

func (h *userHandler) CheckPassword(ctx context.Context, in *CheckPasswordReq, out *CheckPasswordResp) error {
	return h.UserHandler.CheckPassword(ctx, in, out)
}

func (h *userHandler) GetList(ctx context.Context, in *GetListReq, out *GetListResp) error {
	return h.UserHandler.GetList(ctx, in, out)
}
