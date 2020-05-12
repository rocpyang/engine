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
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Md5     string `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5,omitempty"`
	// json字符串
	Conf                 string   `protobuf:"bytes,4,opt,name=conf,proto3" json:"conf,omitempty"`
	Code                 int32    `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
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

func (m *PluginInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PluginInfo) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *PluginInfo) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

func (m *PluginInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PluginInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PluginConf struct {
	// 插件名唯一
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// json字符串
	Conf                 string   `protobuf:"bytes,2,opt,name=conf,proto3" json:"conf,omitempty"`
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
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x6e, 0xf2, 0x30,
	0x10, 0x85, 0x09, 0x90, 0xe4, 0xff, 0x67, 0xd3, 0x32, 0xea, 0xc2, 0xca, 0x0a, 0x65, 0xc5, 0x0a,
	0x0a, 0x94, 0x0b, 0xb4, 0xdd, 0x54, 0xea, 0xa2, 0x0a, 0xea, 0x01, 0x0c, 0x19, 0xa2, 0x48, 0xc4,
	0x13, 0xd9, 0x29, 0xbd, 0x45, 0x2f, 0xd1, 0x8b, 0x56, 0xb6, 0x93, 0xb2, 0xa0, 0x8b, 0xb0, 0xca,
	0x9b, 0x67, 0x7f, 0xf3, 0x5e, 0x0c, 0x58, 0x6b, 0x6e, 0x78, 0x51, 0x1f, 0x3f, 0x8a, 0x52, 0xcd,
	0xdd, 0x80, 0xa1, 0xfb, 0x24, 0x13, 0x7f, 0xb4, 0x93, 0xa6, 0xdc, 0xfb, 0x93, 0xf4, 0x2b, 0x00,
	0x78, 0x73, 0x57, 0x5f, 0xd4, 0x81, 0x11, 0x61, 0xac, 0x64, 0x45, 0x22, 0x98, 0x06, 0xb3, 0xff,
	0x99, 0xd3, 0x28, 0x20, 0x3e, 0x91, 0x36, 0x25, 0x2b, 0x31, 0x74, 0x76, 0x37, 0xe2, 0x2d, 0x8c,
	0xaa, 0x7c, 0x23, 0x46, 0xce, 0xb5, 0xd2, 0xf2, 0x7b, 0x56, 0x07, 0x31, 0xf6, 0xbc, 0xd5, 0xde,
	0xcb, 0x49, 0x84, 0xd3, 0x60, 0x16, 0x66, 0x4e, 0xdb, 0x9d, 0x15, 0x19, 0x23, 0x0b, 0x12, 0x91,
	0xdf, 0xd9, 0x8e, 0xe9, 0x43, 0xd7, 0xe7, 0xa9, 0x65, 0x2f, 0xfa, 0x74, 0x19, 0xc3, 0x73, 0xc6,
	0xea, 0x7b, 0x04, 0x91, 0xc7, 0x70, 0x0d, 0xff, 0x9e, 0xf9, 0x53, 0x1d, 0x59, 0xe6, 0x78, 0xe3,
	0xff, 0x72, 0xde, 0x19, 0xc9, 0x5d, 0x6b, 0x3c, 0xda, 0x27, 0xc8, 0xc8, 0xd4, 0xac, 0x0c, 0xa5,
	0x03, 0xdc, 0x00, 0xbc, 0xb2, 0xcc, 0xdb, 0x15, 0x93, 0xf6, 0xd6, 0xb9, 0x48, 0x82, 0x9d, 0xa5,
	0xb9, 0xd0, 0xb2, 0xb2, 0x8f, 0x95, 0x0e, 0xf0, 0x1e, 0xe2, 0xf7, 0xba, 0xd0, 0x32, 0xa7, 0xcb,
	0xa8, 0xbf, 0x89, 0xe5, 0x6f, 0xcf, 0xde, 0x21, 0x4b, 0x88, 0x32, 0xaa, 0xf8, 0x44, 0xd7, 0xf4,
	0x0a, 0xb7, 0x8d, 0xd4, 0x4d, 0x7f, 0x62, 0x05, 0x71, 0x46, 0x57, 0x32, 0x0b, 0x18, 0x6f, 0x1b,
	0xae, 0x7b, 0x03, 0xbb, 0xc8, 0x99, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x25, 0x86,
	0x8b, 0xa3, 0x02, 0x00, 0x00,
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
	Download(ctx context.Context, in *Download, opts ...grpc.CallOption) (*BasicResponse, error)
	LoadPlugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
	Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*ProgramInfo, error)
	Plugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
	Remove(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
	Start(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
	ReStart(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
	Stop(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Download(ctx context.Context, in *Download, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) LoadPlugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/LoadPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Upgrade(ctx context.Context, in *Download, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Plugin(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Plugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Remove(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Start(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ReStart(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/ReStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Stop(ctx context.Context, in *PluginConf, opts ...grpc.CallOption) (*ProgramInfo, error) {
	out := new(ProgramInfo)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	Download(context.Context, *Download) (*BasicResponse, error)
	LoadPlugin(context.Context, *PluginConf) (*ProgramInfo, error)
	Upgrade(context.Context, *Download) (*ProgramInfo, error)
	Plugin(context.Context, *PluginConf) (*ProgramInfo, error)
	Remove(context.Context, *PluginConf) (*ProgramInfo, error)
	Start(context.Context, *PluginConf) (*ProgramInfo, error)
	ReStart(context.Context, *PluginConf) (*ProgramInfo, error)
	Stop(context.Context, *PluginConf) (*ProgramInfo, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Download)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Download(ctx, req.(*Download))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_LoadPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).LoadPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/LoadPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).LoadPlugin(ctx, req.(*PluginConf))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "Download",
			Handler:    _Plugin_Download_Handler,
		},
		{
			MethodName: "LoadPlugin",
			Handler:    _Plugin_LoadPlugin_Handler,
		},
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
