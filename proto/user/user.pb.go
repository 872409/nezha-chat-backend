// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type UserItem struct {
	Img                  string   `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserItem) Reset()         { *m = UserItem{} }
func (m *UserItem) String() string { return proto.CompactTextString(m) }
func (*UserItem) ProtoMessage()    {}
func (*UserItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *UserItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserItem.Unmarshal(m, b)
}
func (m *UserItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserItem.Marshal(b, m, deterministic)
}
func (m *UserItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserItem.Merge(m, src)
}
func (m *UserItem) XXX_Size() int {
	return xxx_messageInfo_UserItem.Size(m)
}
func (m *UserItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UserItem.DiscardUnknown(m)
}

var xxx_messageInfo_UserItem proto.InternalMessageInfo

func (m *UserItem) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *UserItem) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PostReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReq) Reset()         { *m = PostReq{} }
func (m *PostReq) String() string { return proto.CompactTextString(m) }
func (*PostReq) ProtoMessage()    {}
func (*PostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *PostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReq.Unmarshal(m, b)
}
func (m *PostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReq.Marshal(b, m, deterministic)
}
func (m *PostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReq.Merge(m, src)
}
func (m *PostReq) XXX_Size() int {
	return xxx_messageInfo_PostReq.Size(m)
}
func (m *PostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReq.DiscardUnknown(m)
}

var xxx_messageInfo_PostReq proto.InternalMessageInfo

func (m *PostReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PostResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResp) Reset()         { *m = PostResp{} }
func (m *PostResp) String() string { return proto.CompactTextString(m) }
func (*PostResp) ProtoMessage()    {}
func (*PostResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *PostResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResp.Unmarshal(m, b)
}
func (m *PostResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResp.Marshal(b, m, deterministic)
}
func (m *PostResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResp.Merge(m, src)
}
func (m *PostResp) XXX_Size() int {
	return xxx_messageInfo_PostResp.Size(m)
}
func (m *PostResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResp.DiscardUnknown(m)
}

var xxx_messageInfo_PostResp proto.InternalMessageInfo

func (m *PostResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *LoginResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetListReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListReq) Reset()         { *m = GetListReq{} }
func (m *GetListReq) String() string { return proto.CompactTextString(m) }
func (*GetListReq) ProtoMessage()    {}
func (*GetListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *GetListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListReq.Unmarshal(m, b)
}
func (m *GetListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListReq.Marshal(b, m, deterministic)
}
func (m *GetListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListReq.Merge(m, src)
}
func (m *GetListReq) XXX_Size() int {
	return xxx_messageInfo_GetListReq.Size(m)
}
func (m *GetListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetListReq proto.InternalMessageInfo

func (m *GetListReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetListReq) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetListResp struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	List                 *UserItem `protobuf:"bytes,2,opt,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetListResp) Reset()         { *m = GetListResp{} }
func (m *GetListResp) String() string { return proto.CompactTextString(m) }
func (*GetListResp) ProtoMessage()    {}
func (*GetListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *GetListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResp.Unmarshal(m, b)
}
func (m *GetListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResp.Marshal(b, m, deterministic)
}
func (m *GetListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResp.Merge(m, src)
}
func (m *GetListResp) XXX_Size() int {
	return xxx_messageInfo_GetListResp.Size(m)
}
func (m *GetListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResp proto.InternalMessageInfo

func (m *GetListResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetListResp) GetList() *UserItem {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*UserItem)(nil), "nezha.chat.user.srv.service.UserItem")
	proto.RegisterType((*Error)(nil), "nezha.chat.user.srv.service.Error")
	proto.RegisterType((*PostReq)(nil), "nezha.chat.user.srv.service.PostReq")
	proto.RegisterType((*PostResp)(nil), "nezha.chat.user.srv.service.PostResp")
	proto.RegisterType((*LoginReq)(nil), "nezha.chat.user.srv.service.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "nezha.chat.user.srv.service.LoginResp")
	proto.RegisterType((*GetListReq)(nil), "nezha.chat.user.srv.service.GetListReq")
	proto.RegisterType((*GetListResp)(nil), "nezha.chat.user.srv.service.GetListResp")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9b, 0xb4, 0xe9, 0xf4, 0x22, 0x8b, 0x42, 0x89, 0x97, 0xb2, 0x58, 0xed, 0xc5,
	0x08, 0xf1, 0x52, 0x8f, 0x16, 0x45, 0x84, 0x1e, 0x24, 0x50, 0x10, 0x05, 0x21, 0x36, 0x4b, 0x1b,
	0x34, 0xd9, 0xb8, 0xb3, 0x56, 0xf0, 0xe8, 0x4b, 0xf8, 0xba, 0xb2, 0x93, 0xa4, 0xda, 0x4b, 0x28,
	0xe4, 0x12, 0x66, 0x26, 0xf3, 0xfd, 0xfb, 0x67, 0x66, 0x03, 0x87, 0xb9, 0x92, 0x5a, 0x9e, 0x7f,
	0xa0, 0x50, 0xf4, 0xf0, 0x29, 0x67, 0x47, 0x99, 0xf8, 0x5a, 0x45, 0xfe, 0x62, 0x15, 0x69, 0x9f,
	0xca, 0xa8, 0xd6, 0x3e, 0x0a, 0xb5, 0x4e, 0x16, 0x82, 0x4f, 0xc0, 0x9d, 0xa3, 0x50, 0x77, 0x5a,
	0xa4, 0x6c, 0x1f, 0xda, 0x49, 0xba, 0x1c, 0x58, 0x43, 0x6b, 0xdc, 0x0b, 0x4d, 0xc8, 0x3c, 0x70,
	0x0d, 0x91, 0x45, 0xa9, 0x18, 0xb4, 0xa8, 0xbc, 0xc9, 0xf9, 0x19, 0x38, 0x37, 0x4a, 0x49, 0xc5,
	0x18, 0xd8, 0x0b, 0x19, 0x0b, 0xe2, 0x9c, 0x90, 0x62, 0x23, 0x95, 0xe2, 0xb2, 0x64, 0x4c, 0xc8,
	0xaf, 0xa0, 0x7b, 0x2f, 0x51, 0x87, 0xe2, 0x7d, 0x4b, 0xd5, 0xda, 0x56, 0x35, 0xef, 0xf2, 0x08,
	0xf1, 0x53, 0xaa, 0xb8, 0x3a, 0xb1, 0xca, 0xf9, 0x35, 0xb8, 0x85, 0x04, 0xe6, 0x6c, 0x02, 0x8e,
	0x30, 0xa7, 0x93, 0x40, 0x3f, 0xe0, 0x7e, 0xcd, 0x47, 0xfa, 0xe4, 0x33, 0x2c, 0x00, 0x3e, 0x05,
	0x77, 0x26, 0x97, 0x49, 0xd6, 0xc4, 0xc9, 0x13, 0xf4, 0x4a, 0x8d, 0x26, 0x56, 0xd8, 0x01, 0x38,
	0x5a, 0xbe, 0x8a, 0xac, 0xd4, 0x2f, 0x12, 0x1e, 0x00, 0xdc, 0x0a, 0x3d, 0x4b, 0x8a, 0x61, 0x31,
	0xb0, 0xff, 0xd9, 0xa3, 0x98, 0x16, 0x15, 0xe3, 0xa0, 0x35, 0x6c, 0xd3, 0xa2, 0x62, 0xe4, 0xdf,
	0x16, 0xf4, 0x37, 0x50, 0x23, 0x4f, 0x97, 0x60, 0xbf, 0x25, 0xa8, 0xc9, 0x52, 0x3f, 0x18, 0xd5,
	0x82, 0xd5, 0xcd, 0x09, 0x09, 0x09, 0x7e, 0x5a, 0x60, 0x9b, 0x12, 0x9b, 0x83, 0x6d, 0x16, 0xc5,
	0x8e, 0x6b, 0xe9, 0xf2, 0x3a, 0x78, 0xa3, 0x1d, 0xba, 0x30, 0xe7, 0x7b, 0xec, 0x01, 0x1c, 0x9a,
	0x3a, 0xab, 0x27, 0xaa, 0xed, 0x7a, 0x27, 0xbb, 0xb4, 0x91, 0xf2, 0x33, 0x74, 0xcb, 0xe9, 0xb1,
	0xd3, 0x5a, 0xe8, 0x6f, 0x31, 0xde, 0x78, 0xb7, 0x46, 0xa3, 0x3f, 0xed, 0x3c, 0xda, 0xa6, 0xe3,
	0xa5, 0x43, 0x7f, 0xe4, 0xc5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xff, 0x76, 0x6a, 0xaa,
	0x03, 0x00, 0x00,
}
