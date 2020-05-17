// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/engine.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CommandRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfae4ba294cbbcbf, []int{0}
}

func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (m *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(m, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "proto.CommandRequest")
}

func init() { proto.RegisterFile("proto/engine.proto", fileDescriptor_dfae4ba294cbbcbf) }

var fileDescriptor_dfae4ba294cbbcbf = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x94,
	0x20, 0x44, 0x2a, 0x29, 0xb1, 0x38, 0x33, 0x19, 0x22, 0xa3, 0xa4, 0xc5, 0xc5, 0xe7, 0x9c, 0x9f,
	0x9b, 0x9b, 0x98, 0x97, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc1, 0xc5, 0x9e,
	0x0c, 0x11, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x8d, 0x4a, 0xb8, 0xd8, 0x5c,
	0xc1, 0xa6, 0x0a, 0x59, 0x70, 0xb1, 0x43, 0x75, 0x09, 0x89, 0x42, 0x0c, 0xd2, 0x43, 0x35, 0x45,
	0x4a, 0x04, 0x2a, 0xec, 0x04, 0xb2, 0x2b, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x89,
	0x41, 0xc8, 0x80, 0x8b, 0x3d, 0xb4, 0x20, 0xbd, 0x28, 0x31, 0x25, 0x55, 0x88, 0x1f, 0xaa, 0xc4,
	0x25, 0xbf, 0x3c, 0x2f, 0x27, 0x3f, 0x31, 0x45, 0x4a, 0x08, 0x2a, 0x10, 0x50, 0x94, 0x9f, 0x5e,
	0x94, 0x98, 0xeb, 0x99, 0x97, 0x96, 0xaf, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x34, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x49, 0x10, 0x9a, 0x90, 0xd8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineClient interface {
	// stop restart
	Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*ProgramInfo, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Engine/Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Engine/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
type EngineServer interface {
	// stop restart
	Command(context.Context, *CommandRequest) (*BasicResponse, error)
	Upgrade(context.Context, *Download) (*ProgramInfo, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Engine/Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Command(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Download)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Engine/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Upgrade(ctx, req.(*Download))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Command",
			Handler:    _Engine_Command_Handler,
		},
		{
			MethodName: "Upgrade",
			Handler:    _Engine_Upgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/engine.proto",
}
