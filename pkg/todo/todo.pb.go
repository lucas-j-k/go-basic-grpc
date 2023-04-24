// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/todo/todo.proto

package go_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the search name
type TodoRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoRequest) Reset()         { *m = TodoRequest{} }
func (m *TodoRequest) String() string { return proto.CompactTextString(m) }
func (*TodoRequest) ProtoMessage()    {}
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd38d41dde9f46d4, []int{0}
}

func (m *TodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoRequest.Unmarshal(m, b)
}
func (m *TodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoRequest.Marshal(b, m, deterministic)
}
func (m *TodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoRequest.Merge(m, src)
}
func (m *TodoRequest) XXX_Size() int {
	return xxx_messageInfo_TodoRequest.Size(m)
}
func (m *TodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodoRequest proto.InternalMessageInfo

func (m *TodoRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the stringified response
type TodoReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoReply) Reset()         { *m = TodoReply{} }
func (m *TodoReply) String() string { return proto.CompactTextString(m) }
func (*TodoReply) ProtoMessage()    {}
func (*TodoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd38d41dde9f46d4, []int{1}
}

func (m *TodoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoReply.Unmarshal(m, b)
}
func (m *TodoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoReply.Marshal(b, m, deterministic)
}
func (m *TodoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoReply.Merge(m, src)
}
func (m *TodoReply) XXX_Size() int {
	return xxx_messageInfo_TodoReply.Size(m)
}
func (m *TodoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoReply.DiscardUnknown(m)
}

var xxx_messageInfo_TodoReply proto.InternalMessageInfo

func (m *TodoReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TodoReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*TodoRequest)(nil), "todo.TodoRequest")
	proto.RegisterType((*TodoReply)(nil), "todo.TodoReply")
}

func init() { proto.RegisterFile("pkg/todo/todo.proto", fileDescriptor_cd38d41dde9f46d4) }

var fileDescriptor_cd38d41dde9f46d4 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0x4e, 0xd7,
	0x2f, 0xc9, 0x4f, 0xc9, 0x07, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6,
	0x92, 0x2c, 0x17, 0x77, 0x48, 0x7e, 0x4a, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x8a, 0x92,
	0x29, 0x17, 0x27, 0x44, 0xba, 0x20, 0xa7, 0x12, 0x5d, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f,
	0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x35, 0x32, 0xe7,
	0x62, 0x01, 0x69, 0x13, 0xd2, 0xe7, 0x62, 0x77, 0x4f, 0x2d, 0x01, 0x33, 0x05, 0xf5, 0xc0, 0x76,
	0x23, 0x59, 0x26, 0xc5, 0x8f, 0x2c, 0x54, 0x90, 0x53, 0xa9, 0xc4, 0xe0, 0x24, 0x18, 0xc5, 0x9f,
	0x5a, 0x91, 0x98, 0x5b, 0x90, 0x93, 0xaa, 0x9f, 0x9e, 0xaf, 0x9b, 0x5e, 0x54, 0x90, 0x9c, 0xc4,
	0x06, 0x76, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xe1, 0xf2, 0x96, 0xc5, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	// Get Todo URL
	GetTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoReply, error) {
	out := new(TodoReply)
	err := c.cc.Invoke(ctx, "/todo.Todo/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	// Get Todo URL
	GetTodo(context.Context, *TodoRequest) (*TodoReply, error)
}

// UnimplementedTodoServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (*UnimplementedTodoServer) GetTodo(ctx context.Context, req *TodoRequest) (*TodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTodo(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodo",
			Handler:    _Todo_GetTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/todo/todo.proto",
}
