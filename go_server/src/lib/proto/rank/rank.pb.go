// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rank.proto

package rank

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	api "go_server/src/lib/proto/api"
	errmsg "go_server/src/lib/proto/errmsg"
	recall "go_server/src/lib/proto/recall"
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

type RankReq struct {
	Req                  *api.RecReq        `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	RecallResp           *recall.RecallResp `protobuf:"bytes,2,opt,name=recallResp,proto3" json:"recallResp,omitempty"`
	Bucket               string             `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RankReq) Reset()         { *m = RankReq{} }
func (m *RankReq) String() string { return proto.CompactTextString(m) }
func (*RankReq) ProtoMessage()    {}
func (*RankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_28127d302aca29e8, []int{0}
}

func (m *RankReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankReq.Unmarshal(m, b)
}
func (m *RankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankReq.Marshal(b, m, deterministic)
}
func (m *RankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankReq.Merge(m, src)
}
func (m *RankReq) XXX_Size() int {
	return xxx_messageInfo_RankReq.Size(m)
}
func (m *RankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankReq.DiscardUnknown(m)
}

var xxx_messageInfo_RankReq proto.InternalMessageInfo

func (m *RankReq) GetReq() *api.RecReq {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *RankReq) GetRecallResp() *recall.RecallResp {
	if m != nil {
		return m.RecallResp
	}
	return nil
}

func (m *RankReq) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

type RankResp struct {
	Err                  *errmsg.ErrMsg            `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Items                []*recall.RecallRespItems `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RankResp) Reset()         { *m = RankResp{} }
func (m *RankResp) String() string { return proto.CompactTextString(m) }
func (*RankResp) ProtoMessage()    {}
func (*RankResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_28127d302aca29e8, []int{1}
}

func (m *RankResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankResp.Unmarshal(m, b)
}
func (m *RankResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankResp.Marshal(b, m, deterministic)
}
func (m *RankResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankResp.Merge(m, src)
}
func (m *RankResp) XXX_Size() int {
	return xxx_messageInfo_RankResp.Size(m)
}
func (m *RankResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankResp.DiscardUnknown(m)
}

var xxx_messageInfo_RankResp proto.InternalMessageInfo

func (m *RankResp) GetErr() *errmsg.ErrMsg {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *RankResp) GetItems() []*recall.RecallRespItems {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*RankReq)(nil), "rank.RankReq")
	proto.RegisterType((*RankResp)(nil), "rank.RankResp")
}

func init() { proto.RegisterFile("rank.proto", fileDescriptor_28127d302aca29e8) }

var fileDescriptor_28127d302aca29e8 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x3d, 0x4f, 0x84, 0x40,
	0x10, 0x95, 0xe3, 0x3c, 0xbd, 0x41, 0xaf, 0xd8, 0x42, 0x09, 0x51, 0x43, 0x68, 0xa4, 0x71, 0x37,
	0xc1, 0x7f, 0x60, 0x62, 0x61, 0x61, 0x33, 0xa5, 0x16, 0x06, 0xc8, 0x84, 0x10, 0xb8, 0x63, 0x99,
	0x45, 0x7f, 0xbf, 0xd9, 0x0f, 0xcd, 0x25, 0x57, 0xed, 0xbc, 0x79, 0x6f, 0xdf, 0xdb, 0x7d, 0x00,
	0x5c, 0x1f, 0x06, 0xa9, 0x79, 0x5a, 0x26, 0xb1, 0xb6, 0x73, 0xb6, 0xad, 0x75, 0xef, 0x17, 0xd9,
	0x15, 0x53, 0x5b, 0x8f, 0xe3, 0x1f, 0x22, 0xe6, 0xbd, 0xe9, 0x3c, 0x2a, 0x16, 0xb8, 0xc0, 0xfa,
	0x30, 0x20, 0xcd, 0xe2, 0x1e, 0x62, 0xa6, 0x39, 0x8d, 0xf2, 0xa8, 0x4c, 0xaa, 0x44, 0xda, 0xfb,
	0x48, 0x2d, 0xd2, 0x8c, 0x76, 0x2f, 0x2a, 0x00, 0xef, 0x83, 0x64, 0x74, 0xba, 0x72, 0x2a, 0x21,
	0x83, 0x35, 0xfe, 0x33, 0x78, 0xa4, 0x12, 0x37, 0xb0, 0x69, 0xbe, 0xdb, 0x81, 0x96, 0x34, 0xce,
	0xa3, 0x72, 0x8b, 0x01, 0x15, 0x9f, 0x70, 0xe9, 0x53, 0x8d, 0x16, 0x39, 0xc4, 0xc4, 0x1c, 0x62,
	0x77, 0x32, 0xbc, 0xee, 0x95, 0xf9, 0xdd, 0x74, 0x68, 0x29, 0xf1, 0x04, 0xe7, 0xfd, 0x42, 0x7b,
	0x93, 0xae, 0xf2, 0xb8, 0x4c, 0xaa, 0xdb, 0xd3, 0xd0, 0x37, 0x4b, 0xa3, 0x57, 0x55, 0x0a, 0xd6,
	0xd6, 0x5c, 0x3c, 0x86, 0xf3, 0x5a, 0xba, 0x72, 0xc2, 0x37, 0xb3, 0xdd, 0x31, 0x34, 0xba, 0x38,
	0x7b, 0x79, 0xf8, 0xb8, 0xeb, 0xa6, 0x2f, 0x43, 0xfc, 0x43, 0xac, 0x0c, 0xb7, 0x6a, 0xec, 0x1b,
	0xe5, 0xea, 0x51, 0x56, 0xda, 0x6c, 0xdc, 0xfc, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x02, 0x91,
	0x6c, 0x54, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankClient interface {
	Rank(ctx context.Context, in *RankReq, opts ...grpc.CallOption) (*RankResp, error)
}

type rankClient struct {
	cc *grpc.ClientConn
}

func NewRankClient(cc *grpc.ClientConn) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) Rank(ctx context.Context, in *RankReq, opts ...grpc.CallOption) (*RankResp, error) {
	out := new(RankResp)
	err := c.cc.Invoke(ctx, "/rank.Rank/Rank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
type RankServer interface {
	Rank(context.Context, *RankReq) (*RankResp, error)
}

// UnimplementedRankServer can be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (*UnimplementedRankServer) Rank(ctx context.Context, req *RankReq) (*RankResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rank not implemented")
}

func RegisterRankServer(s *grpc.Server, srv RankServer) {
	s.RegisterService(&_Rank_serviceDesc, srv)
}

func _Rank_Rank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).Rank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rank.Rank/Rank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).Rank(ctx, req.(*RankReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rank.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rank",
			Handler:    _Rank_Rank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rank.proto",
}