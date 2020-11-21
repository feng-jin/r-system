// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recall.proto

package recall

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	api "go_server/src/lib/proto/api"
	errmsg "go_server/src/lib/proto/errmsg"
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

type RecallReq struct {
	Req                  *api.RecReq `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	Bucket               string      `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RecallReq) Reset()         { *m = RecallReq{} }
func (m *RecallReq) String() string { return proto.CompactTextString(m) }
func (*RecallReq) ProtoMessage()    {}
func (*RecallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a73c1036b473dc8, []int{0}
}

func (m *RecallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecallReq.Unmarshal(m, b)
}
func (m *RecallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecallReq.Marshal(b, m, deterministic)
}
func (m *RecallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecallReq.Merge(m, src)
}
func (m *RecallReq) XXX_Size() int {
	return xxx_messageInfo_RecallReq.Size(m)
}
func (m *RecallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RecallReq.DiscardUnknown(m)
}

var xxx_messageInfo_RecallReq proto.InternalMessageInfo

func (m *RecallReq) GetReq() *api.RecReq {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *RecallReq) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

type RecallRespItems struct {
	ItemId               string   `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Weight               float32  `protobuf:"fixed32,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Labels               []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	Score                float32  `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecallRespItems) Reset()         { *m = RecallRespItems{} }
func (m *RecallRespItems) String() string { return proto.CompactTextString(m) }
func (*RecallRespItems) ProtoMessage()    {}
func (*RecallRespItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a73c1036b473dc8, []int{1}
}

func (m *RecallRespItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecallRespItems.Unmarshal(m, b)
}
func (m *RecallRespItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecallRespItems.Marshal(b, m, deterministic)
}
func (m *RecallRespItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecallRespItems.Merge(m, src)
}
func (m *RecallRespItems) XXX_Size() int {
	return xxx_messageInfo_RecallRespItems.Size(m)
}
func (m *RecallRespItems) XXX_DiscardUnknown() {
	xxx_messageInfo_RecallRespItems.DiscardUnknown(m)
}

var xxx_messageInfo_RecallRespItems proto.InternalMessageInfo

func (m *RecallRespItems) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *RecallRespItems) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *RecallRespItems) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *RecallRespItems) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type RecallResp struct {
	Err                  *errmsg.ErrMsg     `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Items                []*RecallRespItems `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RecallResp) Reset()         { *m = RecallResp{} }
func (m *RecallResp) String() string { return proto.CompactTextString(m) }
func (*RecallResp) ProtoMessage()    {}
func (*RecallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a73c1036b473dc8, []int{2}
}

func (m *RecallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecallResp.Unmarshal(m, b)
}
func (m *RecallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecallResp.Marshal(b, m, deterministic)
}
func (m *RecallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecallResp.Merge(m, src)
}
func (m *RecallResp) XXX_Size() int {
	return xxx_messageInfo_RecallResp.Size(m)
}
func (m *RecallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RecallResp.DiscardUnknown(m)
}

var xxx_messageInfo_RecallResp proto.InternalMessageInfo

func (m *RecallResp) GetErr() *errmsg.ErrMsg {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *RecallResp) GetItems() []*RecallRespItems {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*RecallReq)(nil), "recall.RecallReq")
	proto.RegisterType((*RecallRespItems)(nil), "recall.RecallRespItems")
	proto.RegisterType((*RecallResp)(nil), "recall.RecallResp")
}

func init() { proto.RegisterFile("recall.proto", fileDescriptor_1a73c1036b473dc8) }

var fileDescriptor_1a73c1036b473dc8 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4b, 0x84, 0x50,
	0x10, 0xc6, 0x53, 0x5b, 0xc3, 0x71, 0x29, 0x7a, 0x44, 0x2b, 0x42, 0xf1, 0xf0, 0xe4, 0x25, 0x25,
	0x3b, 0x76, 0x5b, 0xe8, 0xb0, 0x87, 0x2e, 0x73, 0x0c, 0x62, 0x51, 0x77, 0x30, 0x49, 0x51, 0xe7,
	0x59, 0xfd, 0xfb, 0xf1, 0x7c, 0xaf, 0x82, 0xed, 0xf6, 0x7e, 0xdf, 0x7c, 0xf3, 0xcd, 0xbc, 0x81,
	0x35, 0x53, 0x5d, 0x76, 0x5d, 0x36, 0xf2, 0x30, 0x0f, 0xc2, 0x37, 0x14, 0x07, 0xe5, 0xd8, 0x1a,
	0x29, 0x5e, 0x13, 0x73, 0xaf, 0x1a, 0x43, 0xc9, 0x16, 0x02, 0x5c, 0x2c, 0x48, 0x93, 0xb8, 0x01,
	0x8f, 0x69, 0x8a, 0x1c, 0xe9, 0xa4, 0x61, 0x11, 0x66, 0xba, 0x07, 0xa9, 0x46, 0x9a, 0x50, 0xeb,
	0xe2, 0x1a, 0xfc, 0xea, 0xa3, 0x7e, 0xa7, 0x39, 0x72, 0xa5, 0x93, 0x06, 0x68, 0x29, 0x19, 0xe1,
	0xe2, 0x27, 0x43, 0x8d, 0xbb, 0x99, 0x7a, 0x25, 0x36, 0x70, 0xd6, 0xce, 0xd4, 0xef, 0xdb, 0xc3,
	0x92, 0x16, 0xa0, 0xaf, 0x71, 0x77, 0xd0, 0x19, 0x5f, 0xd4, 0x36, 0x6f, 0x26, 0xc3, 0x45, 0x4b,
	0x5a, 0xef, 0xca, 0x8a, 0x3a, 0x15, 0x79, 0xd2, 0xd3, 0x7e, 0x43, 0xe2, 0x0a, 0x56, 0xaa, 0x1e,
	0x98, 0xa2, 0xd3, 0xc5, 0x6e, 0x20, 0x79, 0x05, 0xf8, 0x9b, 0x28, 0x24, 0x78, 0xc4, 0x6c, 0xd7,
	0x3e, 0xcf, 0xec, 0xff, 0x9e, 0x98, 0x9f, 0x55, 0x83, 0xba, 0x24, 0xee, 0x60, 0xa5, 0xe7, 0xab,
	0xc8, 0x95, 0x5e, 0x1a, 0x16, 0x9b, 0xcc, 0x1e, 0xe9, 0x68, 0x6d, 0x34, 0xae, 0xe2, 0x11, 0x7c,
	0x53, 0x11, 0xf7, 0xbf, 0xaf, 0xcb, 0xe3, 0x9e, 0x29, 0x16, 0xff, 0x63, 0x92, 0x93, 0xad, 0x7c,
	0xb9, 0x6d, 0x86, 0xbd, 0x22, 0xfe, 0x24, 0xce, 0x15, 0xd7, 0x79, 0xd7, 0x56, 0xf9, 0x72, 0xec,
	0xdc, 0xd8, 0x2b, 0x7f, 0xa1, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x47, 0x8b, 0x16,
	0xab, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecallClient is the client API for Recall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecallClient interface {
	Recall(ctx context.Context, in *RecallReq, opts ...grpc.CallOption) (*RecallResp, error)
}

type recallClient struct {
	cc *grpc.ClientConn
}

func NewRecallClient(cc *grpc.ClientConn) RecallClient {
	return &recallClient{cc}
}

func (c *recallClient) Recall(ctx context.Context, in *RecallReq, opts ...grpc.CallOption) (*RecallResp, error) {
	out := new(RecallResp)
	err := c.cc.Invoke(ctx, "/recall.Recall/Recall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecallServer is the server API for Recall service.
type RecallServer interface {
	Recall(context.Context, *RecallReq) (*RecallResp, error)
}

// UnimplementedRecallServer can be embedded to have forward compatible implementations.
type UnimplementedRecallServer struct {
}

func (*UnimplementedRecallServer) Recall(ctx context.Context, req *RecallReq) (*RecallResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recall not implemented")
}

func RegisterRecallServer(s *grpc.Server, srv RecallServer) {
	s.RegisterService(&_Recall_serviceDesc, srv)
}

func _Recall_Recall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecallServer).Recall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recall.Recall/Recall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecallServer).Recall(ctx, req.(*RecallReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recall.Recall",
	HandlerType: (*RecallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recall",
			Handler:    _Recall_Recall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recall.proto",
}
