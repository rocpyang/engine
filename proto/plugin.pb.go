// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/plugin.proto

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

type PluginInfo struct {
	// 插件名唯一
	Name    string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Version []int32 `protobuf:"varint,2,rep,packed,name=Version,proto3" json:"Version,omitempty"`
	MD5     string  `protobuf:"bytes,3,opt,name=MD5,proto3" json:"MD5,omitempty"`
	// json字符串
	Conf                 string   `protobuf:"bytes,4,opt,name=Conf,proto3" json:"Conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfo) Reset()         { *m = PluginInfo{} }
func (m *PluginInfo) String() string { return proto.CompactTextString(m) }
func (*PluginInfo) ProtoMessage()    {}
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4368bc2f41172306, []int{0}
}

func (m *PluginInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfo.Unmarshal(m, b)
}
func (m *PluginInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfo.Marshal(b, m, deterministic)
}
func (m *PluginInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfo.Merge(m, src)
}
func (m *PluginInfo) XXX_Size() int {
	return xxx_messageInfo_PluginInfo.Size(m)
}
func (m *PluginInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfo proto.InternalMessageInfo

func (m *PluginInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginInfo) GetVersion() []int32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *PluginInfo) GetMD5() string {
	if m != nil {
		return m.MD5
	}
	return ""
}

func (m *PluginInfo) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

type PluginConf struct {
	// 插件名唯一
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// json字符串
	Conf                 string   `protobuf:"bytes,2,opt,name=Conf,proto3" json:"Conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginConf) Reset()         { *m = PluginConf{} }
func (m *PluginConf) String() string { return proto.CompactTextString(m) }
func (*PluginConf) ProtoMessage()    {}
func (*PluginConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_4368bc2f41172306, []int{1}
}

func (m *PluginConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginConf.Unmarshal(m, b)
}
func (m *PluginConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginConf.Marshal(b, m, deterministic)
}
func (m *PluginConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginConf.Merge(m, src)
}
func (m *PluginConf) XXX_Size() int {
	return xxx_messageInfo_PluginConf.Size(m)
}
func (m *PluginConf) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginConf.DiscardUnknown(m)
}

var xxx_messageInfo_PluginConf proto.InternalMessageInfo

func (m *PluginConf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginConf) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

func init() {
	proto.RegisterType((*PluginInfo)(nil), "proto.PluginInfo")
	proto.RegisterType((*PluginConf)(nil), "proto.PluginConf")
}

func init() { proto.RegisterFile("proto/plugin.proto", fileDescriptor_4368bc2f41172306) }

var fileDescriptor_4368bc2f41172306 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x61, 0xff, 0xc6, 0xb9, 0x28, 0x13, 0x0f, 0xcd, 0x9e, 0xc8, 0x9e, 0x38, 0xa1, 0x2e,
	0xf8, 0x05, 0x94, 0x8b, 0x07, 0x8d, 0x29, 0xd1, 0xb3, 0x45, 0x06, 0xb2, 0x09, 0x74, 0x9a, 0xdd,
	0xaa, 0x5f, 0xcf, 0x8f, 0x66, 0xda, 0xb2, 0x12, 0x13, 0x3c, 0xec, 0xa9, 0xf3, 0x5e, 0xdf, 0x2f,
	0x6f, 0x06, 0xd0, 0x34, 0x6c, 0xf9, 0xca, 0xec, 0x3e, 0xb6, 0xb5, 0x9e, 0x7a, 0x81, 0xa9, 0x7f,
	0x8a, 0x51, 0xf8, 0x5a, 0xa9, 0xb6, 0x7e, 0x0f, 0x3f, 0xe5, 0x1b, 0xc0, 0xb3, 0x4f, 0x3e, 0xe8,
	0x0d, 0x23, 0x42, 0xf2, 0xa4, 0xf6, 0x24, 0x86, 0xe3, 0xe1, 0xe4, 0x4c, 0xfa, 0x19, 0x05, 0xe4,
	0xaf, 0xd4, 0xb4, 0x35, 0x6b, 0x11, 0x8d, 0xe3, 0x49, 0x2a, 0x3b, 0x89, 0x17, 0x10, 0x3f, 0x2e,
	0x6e, 0x45, 0xec, 0xc3, 0x6e, 0x74, 0xfc, 0x3d, 0xeb, 0x8d, 0x48, 0x02, 0xef, 0xe6, 0x72, 0xde,
	0x35, 0x38, 0x75, 0xb2, 0xa1, 0xa3, 0xa2, 0x23, 0x55, 0x7d, 0x47, 0x90, 0x05, 0x0c, 0x2b, 0xc8,
	0x5f, 0xcc, 0xb6, 0x51, 0x6b, 0xc2, 0xf3, 0xb0, 0xf5, 0x74, 0xc1, 0x5f, 0x7a, 0xc7, 0x6a, 0x5d,
	0x5c, 0x1e, 0x8c, 0x3b, 0x77, 0x92, 0xa4, 0xd6, 0xb0, 0x6e, 0xa9, 0x1c, 0xe0, 0xf5, 0x2f, 0x3d,
	0x3a, 0x24, 0x8e, 0x3b, 0x14, 0x7f, 0x2d, 0x77, 0x78, 0x39, 0xc0, 0x19, 0x64, 0x92, 0xf6, 0xfc,
	0x49, 0xa7, 0x88, 0xff, 0x6a, 0x2a, 0x48, 0x97, 0x56, 0x35, 0xb6, 0x0f, 0x33, 0x87, 0x5c, 0x52,
	0x6f, 0xea, 0x06, 0x92, 0xa5, 0x65, 0xd3, 0x03, 0x59, 0x65, 0xde, 0x9e, 0xfd, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0xe2, 0xfc, 0x87, 0x11, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*BasicResponse, error)
	Plugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*PluginInfo, error)
	Remove(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error)
	Start(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error)
	ReStart(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error)
	Stop(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Plugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Plugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Remove(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Start(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ReStart(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/ReStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Stop(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	Upgrade(context.Context, *Download) (*BasicResponse, error)
	Plugin(context.Context, *PluginConf) (*PluginInfo, error)
	Remove(context.Context, *PluginConf) (*BasicResponse, error)
	Start(context.Context, *PluginConf) (*BasicResponse, error)
	ReStart(context.Context, *PluginConf) (*BasicResponse, error)
	Stop(context.Context, *PluginConf) (*BasicResponse, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Download)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Upgrade(ctx, req.(*Download))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Plugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Plugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Plugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Plugin(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Remove(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Start(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ReStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ReStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/ReStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ReStart(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Stop(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upgrade",
			Handler:    _Plugin_Upgrade_Handler,
		},
		{
			MethodName: "Plugin",
			Handler:    _Plugin_Plugin_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Plugin_Remove_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Plugin_Start_Handler,
		},
		{
			MethodName: "ReStart",
			Handler:    _Plugin_ReStart_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Plugin_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/plugin.proto",
}
