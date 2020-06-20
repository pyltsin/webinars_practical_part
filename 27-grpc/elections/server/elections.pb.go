// Code generated by protoc-gen-go. DO NOT EDIT.
// source: elections.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Vote struct {
	Passport             string               `protobuf:"bytes,1,opt,name=passport,proto3" json:"passport,omitempty"`
	CandidateId          uint32               `protobuf:"varint,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Note                 string               `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a2d3077d176f007, []int{0}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetPassport() string {
	if m != nil {
		return m.Passport
	}
	return ""
}

func (m *Vote) GetCandidateId() uint32 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *Vote) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Vote) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*Vote)(nil), "Vote")
}

func init() {
	proto.RegisterFile("elections.proto", fileDescriptor_1a2d3077d176f007)
}

var fileDescriptor_1a2d3077d176f007 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcd, 0x49, 0x4d,
	0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4f, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b,
	0x12, 0x73, 0x0b, 0xa0, 0x0a, 0xa4, 0xd1, 0x15, 0xa4, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x24, 0x95,
	0x3a, 0x19, 0xb9, 0x58, 0xc2, 0xf2, 0x4b, 0x52, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b,
	0x0b, 0xf2, 0x8b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x45, 0x2e,
	0x9e, 0xe4, 0xc4, 0xbc, 0x94, 0xcc, 0x94, 0xc4, 0x92, 0xd4, 0xf8, 0xcc, 0x14, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xde, 0x20, 0x6e, 0xb8, 0x98, 0x67, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x7e, 0x49,
	0xaa, 0x04, 0x33, 0x58, 0x2b, 0x98, 0x2d, 0xa4, 0xc7, 0xc5, 0x02, 0x72, 0x8b, 0x04, 0x8b, 0x02,
	0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x1d, 0x7a, 0x30, 0x77, 0xe8, 0x85, 0xc0, 0x1c, 0x1a,
	0x04, 0x56, 0x67, 0x64, 0xc5, 0xc5, 0xe9, 0x0a, 0xf3, 0x9c, 0x90, 0x2e, 0x17, 0x57, 0x70, 0x69,
	0x52, 0x6e, 0x66, 0x09, 0xd8, 0x75, 0xac, 0x7a, 0x20, 0x4a, 0x4a, 0x0c, 0xc3, 0x0c, 0x57, 0x90,
	0x5f, 0x94, 0x18, 0x9c, 0xd8, 0xa2, 0x58, 0x72, 0x13, 0x33, 0xf3, 0x92, 0xd8, 0xc0, 0x32, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x80, 0x03, 0xcc, 0x27, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ElectionsClient is the client API for Elections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElectionsClient interface {
	SubmitVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*empty.Empty, error)
}

type electionsClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionsClient(cc grpc.ClientConnInterface) ElectionsClient {
	return &electionsClient{cc}
}

func (c *electionsClient) SubmitVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Elections/SubmitVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionsServer is the server API for Elections service.
type ElectionsServer interface {
	SubmitVote(context.Context, *Vote) (*empty.Empty, error)
}

// UnimplementedElectionsServer can be embedded to have forward compatible implementations.
type UnimplementedElectionsServer struct {
}

func (*UnimplementedElectionsServer) SubmitVote(ctx context.Context, req *Vote) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVote not implemented")
}

func RegisterElectionsServer(s *grpc.Server, srv ElectionsServer) {
	s.RegisterService(&_Elections_serviceDesc, srv)
}

func _Elections_SubmitVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionsServer).SubmitVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Elections/SubmitVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionsServer).SubmitVote(ctx, req.(*Vote))
	}
	return interceptor(ctx, in, info, handler)
}

var _Elections_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Elections",
	HandlerType: (*ElectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitVote",
			Handler:    _Elections_SubmitVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elections.proto",
}