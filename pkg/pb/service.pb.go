// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/example-app/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_404d2ad77d2f86eb, []int{0}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
}

func init() {
	proto.RegisterFile("github.com/example-app/pkg/pb/service.proto", fileDescriptor_404d2ad77d2f86eb)
}

var fileDescriptor_404d2ad77d2f86eb = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x87, 0xd5, 0x3b, 0xdc, 0x82, 0x17, 0x50, 0x06, 0x54, 0x15, 0x86, 0xaa, 0x13, 0x52, 0x89,
	0x2d, 0xc1, 0x06, 0x53, 0x91, 0x2a, 0x56, 0xd4, 0x81, 0x81, 0x01, 0xc9, 0x0e, 0xa7, 0xc6, 0x22,
	0xf1, 0x39, 0xb2, 0xdd, 0x90, 0xac, 0xbc, 0x02, 0x8f, 0xc6, 0x2b, 0xf0, 0x20, 0xa8, 0xb1, 0x83,
	0x22, 0x40, 0x6c, 0x3e, 0xff, 0x7e, 0x9f, 0xf4, 0x99, 0x2d, 0xb4, 0x09, 0x4f, 0x5b, 0xc5, 0x0b,
	0xac, 0x04, 0x34, 0xb2, 0xa2, 0x12, 0x72, 0x49, 0x24, 0xe8, 0x59, 0x0b, 0x52, 0xc2, 0x83, 0xab,
	0x4d, 0x01, 0x9c, 0x1c, 0x06, 0xcc, 0xc6, 0xa9, 0x9c, 0x1e, 0x6b, 0x44, 0x5d, 0x82, 0xe8, 0xda,
	0x6a, 0xbb, 0x11, 0x50, 0x51, 0x68, 0xe3, 0xd6, 0xf4, 0x24, 0x0d, 0x25, 0x19, 0x21, 0xad, 0xc5,
	0x20, 0x83, 0x41, 0xeb, 0xd3, 0x74, 0x39, 0x04, 0xda, 0x1a, 0x5b, 0x72, 0xd8, 0xb4, 0x31, 0xa9,
	0xc8, 0x35, 0xd8, 0xbc, 0x96, 0xa5, 0x79, 0x94, 0x01, 0xc4, 0x8f, 0x47, 0x8a, 0x38, 0x1b, 0x2c,
	0xfb, 0x17, 0xa9, 0x35, 0x38, 0x81, 0xd4, 0x41, 0x7e, 0x01, 0x5e, 0x0e, 0x80, 0xc6, 0x6e, 0x50,
	0x95, 0xd8, 0x20, 0x81, 0x1d, 0x22, 0x35, 0xba, 0xea, 0x2b, 0x62, 0x57, 0xc4, 0xdb, 0xf9, 0x82,
	0x1d, 0xdc, 0x81, 0xf3, 0x06, 0xed, 0x1a, 0x3c, 0xa1, 0xf5, 0x90, 0x4d, 0xd8, 0xb8, 0x8e, 0xad,
	0xc9, 0x68, 0x36, 0x3a, 0xdd, 0x5f, 0xf7, 0xe5, 0xf9, 0x03, 0x63, 0xab, 0x68, 0x70, 0x49, 0x94,
	0xdd, 0x32, 0x76, 0x03, 0x21, 0x5d, 0x67, 0x47, 0x3c, 0x4a, 0xe1, 0xbd, 0x31, 0xbe, 0xda, 0x19,
	0x9b, 0x4e, 0x78, 0x6f, 0xf8, 0x1b, 0x67, 0x7e, 0xf8, 0xfa, 0xfe, 0xf1, 0xf6, 0x8f, 0x65, 0x7b,
	0x22, 0xe5, 0x5f, 0xcf, 0xef, 0x67, 0x7f, 0x7e, 0xd6, 0x15, 0x29, 0xf5, 0xbf, 0xcb, 0xbf, 0xf8,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x85, 0x67, 0xb7, 0x39, 0xd7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleAppClient is the client API for ExampleApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleAppClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type exampleAppClient struct {
	cc *grpc.ClientConn
}

func NewExampleAppClient(cc *grpc.ClientConn) ExampleAppClient {
	return &exampleAppClient{cc}
}

func (c *exampleAppClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleAppServer is the server API for ExampleApp service.
type ExampleAppServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
}

func RegisterExampleAppServer(s *grpc.Server, srv ExampleAppServer) {
	s.RegisterService(&_ExampleApp_serviceDesc, srv)
}

func _ExampleApp_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ExampleApp",
	HandlerType: (*ExampleAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _ExampleApp_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/example-app/pkg/pb/service.proto",
}
