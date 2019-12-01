// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package pb

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

type GetSampleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSampleRequest) Reset()         { *m = GetSampleRequest{} }
func (m *GetSampleRequest) String() string { return proto.CompactTextString(m) }
func (*GetSampleRequest) ProtoMessage()    {}
func (*GetSampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *GetSampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSampleRequest.Unmarshal(m, b)
}
func (m *GetSampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSampleRequest.Marshal(b, m, deterministic)
}
func (m *GetSampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSampleRequest.Merge(m, src)
}
func (m *GetSampleRequest) XXX_Size() int {
	return xxx_messageInfo_GetSampleRequest.Size(m)
}
func (m *GetSampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSampleRequest proto.InternalMessageInfo

func (m *GetSampleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SampleResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleResponse) Reset()         { *m = SampleResponse{} }
func (m *SampleResponse) String() string { return proto.CompactTextString(m) }
func (*SampleResponse) ProtoMessage()    {}
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{1}
}

func (m *SampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleResponse.Unmarshal(m, b)
}
func (m *SampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleResponse.Marshal(b, m, deterministic)
}
func (m *SampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleResponse.Merge(m, src)
}
func (m *SampleResponse) XXX_Size() int {
	return xxx_messageInfo_SampleResponse.Size(m)
}
func (m *SampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SampleResponse proto.InternalMessageInfo

func (m *SampleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSampleRequest)(nil), "pb.GetSampleRequest")
	proto.RegisterType((*SampleResponse)(nil), "pb.SampleResponse")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe3, 0x12,
	0x70, 0x4f, 0x2d, 0x09, 0x06, 0x0b, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x5a,
	0x5c, 0x7c, 0x30, 0x45, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9,
	0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x85, 0x30, 0xae, 0x91, 0x23, 0x17, 0x1b, 0x44, 0xad, 0x90, 0x39,
	0x17, 0x27, 0xdc, 0x74, 0x21, 0x11, 0xbd, 0x82, 0x24, 0x3d, 0x74, 0xcb, 0xa4, 0x84, 0x40, 0xa2,
	0xa8, 0x46, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x5d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x09, 0x75, 0x0e, 0xb1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleClient interface {
	GetSample(ctx context.Context, in *GetSampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type sampleClient struct {
	cc *grpc.ClientConn
}

func NewSampleClient(cc *grpc.ClientConn) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) GetSample(ctx context.Context, in *GetSampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/pb.Sample/GetSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
type SampleServer interface {
	GetSample(context.Context, *GetSampleRequest) (*SampleResponse, error)
}

// UnimplementedSampleServer can be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (*UnimplementedSampleServer) GetSample(ctx context.Context, req *GetSampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSample not implemented")
}

func RegisterSampleServer(s *grpc.Server, srv SampleServer) {
	s.RegisterService(&_Sample_serviceDesc, srv)
}

func _Sample_GetSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).GetSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sample/GetSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).GetSample(ctx, req.(*GetSampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSample",
			Handler:    _Sample_GetSample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}