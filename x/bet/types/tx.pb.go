// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgPlaceBet defines a message to place a bet with the given data
type MsgPlaceBet struct {
	// creator is the bettor address
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// PlaceBetFields contains bet fields
	Bet *PlaceBetFields `protobuf:"bytes,2,opt,name=bet,proto3" json:"bet,omitempty"`
}

func (m *MsgPlaceBet) Reset()         { *m = MsgPlaceBet{} }
func (m *MsgPlaceBet) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBet) ProtoMessage()    {}
func (*MsgPlaceBet) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b4167f68c2a7f8, []int{0}
}
func (m *MsgPlaceBet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBet.Merge(m, src)
}
func (m *MsgPlaceBet) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBet proto.InternalMessageInfo

func (m *MsgPlaceBet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPlaceBet) GetBet() *PlaceBetFields {
	if m != nil {
		return m.Bet
	}
	return nil
}

// MsgPlaceBetResponse is the returning value in the response of MsgPlaceBet
// request
type MsgPlaceBetResponse struct {
	Error string          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Bet   *PlaceBetFields `protobuf:"bytes,2,opt,name=bet,proto3" json:"bet,omitempty"`
}

func (m *MsgPlaceBetResponse) Reset()         { *m = MsgPlaceBetResponse{} }
func (m *MsgPlaceBetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBetResponse) ProtoMessage()    {}
func (*MsgPlaceBetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b4167f68c2a7f8, []int{1}
}
func (m *MsgPlaceBetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBetResponse.Merge(m, src)
}
func (m *MsgPlaceBetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBetResponse proto.InternalMessageInfo

func (m *MsgPlaceBetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MsgPlaceBetResponse) GetBet() *PlaceBetFields {
	if m != nil {
		return m.Bet
	}
	return nil
}

// MsgSettleBet defines a message to settle the given bet
type MsgSettleBet struct {
	// creator is the bettor address
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// bet_uid is the unique uuid of the bet to settle
	BetUID string `protobuf:"bytes,2,opt,name=bet_uid,proto3" json:"bet_uid"`
	// bettor_address is sthe bec32 address of the bettor
	BettorAddress string `protobuf:"bytes,3,opt,name=bettor_address,json=bettorAddress,proto3" json:"bettor_address,omitempty"`
}

func (m *MsgSettleBet) Reset()         { *m = MsgSettleBet{} }
func (m *MsgSettleBet) String() string { return proto.CompactTextString(m) }
func (*MsgSettleBet) ProtoMessage()    {}
func (*MsgSettleBet) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b4167f68c2a7f8, []int{2}
}
func (m *MsgSettleBet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSettleBet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSettleBet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSettleBet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSettleBet.Merge(m, src)
}
func (m *MsgSettleBet) XXX_Size() int {
	return m.Size()
}
func (m *MsgSettleBet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSettleBet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSettleBet proto.InternalMessageInfo

func (m *MsgSettleBet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSettleBet) GetBetUID() string {
	if m != nil {
		return m.BetUID
	}
	return ""
}

func (m *MsgSettleBet) GetBettorAddress() string {
	if m != nil {
		return m.BettorAddress
	}
	return ""
}

// MsgSettleBetResponse is the returning value in the response of MsgSettleBet
// request
type MsgSettleBetResponse struct {
	// Error contains any error while settlement
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// BetUID is the uuid of the bet to be settled
	BetUID string `protobuf:"bytes,2,opt,name=bet_uid,proto3" json:"bet_uid"`
}

func (m *MsgSettleBetResponse) Reset()         { *m = MsgSettleBetResponse{} }
func (m *MsgSettleBetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSettleBetResponse) ProtoMessage()    {}
func (*MsgSettleBetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b4167f68c2a7f8, []int{3}
}
func (m *MsgSettleBetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSettleBetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSettleBetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSettleBetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSettleBetResponse.Merge(m, src)
}
func (m *MsgSettleBetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSettleBetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSettleBetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSettleBetResponse proto.InternalMessageInfo

func (m *MsgSettleBetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MsgSettleBetResponse) GetBetUID() string {
	if m != nil {
		return m.BetUID
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgPlaceBet)(nil), "sgenetwork.sge.bet.MsgPlaceBet")
	proto.RegisterType((*MsgPlaceBetResponse)(nil), "sgenetwork.sge.bet.MsgPlaceBetResponse")
	proto.RegisterType((*MsgSettleBet)(nil), "sgenetwork.sge.bet.MsgSettleBet")
	proto.RegisterType((*MsgSettleBetResponse)(nil), "sgenetwork.sge.bet.MsgSettleBetResponse")
}

func init() { proto.RegisterFile("sge/bet/tx.proto", fileDescriptor_38b4167f68c2a7f8) }

var fileDescriptor_38b4167f68c2a7f8 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0xed, 0xbc, 0xf2, 0xda, 0xd7, 0xe9, 0x7b, 0x0f, 0x89, 0x15, 0x42, 0x16, 0x69, 0x09, 0xa8,
	0xdd, 0x98, 0x40, 0xf5, 0x03, 0x34, 0x88, 0xe0, 0xa2, 0x20, 0x51, 0x11, 0x84, 0x52, 0x32, 0xcd,
	0x75, 0x2c, 0xd6, 0x4e, 0x98, 0xb9, 0xc5, 0xba, 0xf3, 0x13, 0xfc, 0x18, 0x3f, 0xc2, 0x65, 0x97,
	0xae, 0x8a, 0xa4, 0x3b, 0xbf, 0x42, 0x92, 0x98, 0x58, 0x50, 0x09, 0xba, 0xbb, 0x33, 0xe7, 0xcc,
	0xb9, 0xe7, 0x9e, 0xb9, 0x74, 0x45, 0x71, 0x70, 0x18, 0xa0, 0x83, 0x53, 0x3b, 0x94, 0x02, 0x85,
	0xa6, 0x29, 0x0e, 0x63, 0xc0, 0x1b, 0x21, 0xaf, 0x6c, 0xc5, 0xc1, 0x66, 0x80, 0x46, 0x83, 0x0b,
	0x2e, 0x12, 0xd8, 0x89, 0xab, 0x94, 0x69, 0x98, 0xd9, 0xdb, 0x70, 0xe4, 0x0f, 0xa0, 0xcf, 0x00,
	0xfb, 0x17, 0x43, 0x18, 0x05, 0x2a, 0xc5, 0xad, 0x1e, 0xad, 0x77, 0x15, 0x3f, 0x8a, 0x41, 0x17,
	0x50, 0xd3, 0x69, 0x75, 0x20, 0xc1, 0x47, 0x21, 0x75, 0xd2, 0x22, 0xed, 0x9a, 0x97, 0x1d, 0xb5,
	0x1d, 0x5a, 0x66, 0x80, 0xfa, 0xaf, 0x16, 0x69, 0xd7, 0x3b, 0x96, 0xfd, 0xd1, 0x80, 0x9d, 0x89,
	0x1c, 0x24, 0xfa, 0x5e, 0x4c, 0xb7, 0x7c, 0xba, 0xba, 0x24, 0xef, 0x81, 0x0a, 0xc5, 0x58, 0x81,
	0xd6, 0xa0, 0xbf, 0x41, 0xca, 0xbc, 0x49, 0x7a, 0xf8, 0x61, 0x8b, 0x3b, 0x42, 0xff, 0x76, 0x15,
	0x3f, 0x06, 0xc4, 0x51, 0xc1, 0x0c, 0x0e, 0xad, 0xc6, 0x01, 0x4c, 0x86, 0x41, 0xd2, 0xa4, 0xe6,
	0xae, 0x45, 0xf3, 0x66, 0xc5, 0x05, 0x3c, 0x3d, 0xdc, 0x7f, 0x99, 0x37, 0x33, 0xd0, 0xcb, 0x0a,
	0x6d, 0x9d, 0xfe, 0x67, 0x80, 0x28, 0x64, 0xdf, 0x0f, 0x02, 0x09, 0x4a, 0xe9, 0xe5, 0x44, 0xf1,
	0x5f, 0x7a, 0xbb, 0x97, 0x5e, 0x5a, 0x3d, 0xda, 0x58, 0x76, 0x50, 0x30, 0xe6, 0x77, 0x5d, 0x74,
	0x1e, 0x08, 0x2d, 0x77, 0x15, 0xd7, 0x4e, 0xe8, 0x9f, 0xfc, 0xa3, 0x9a, 0x9f, 0xc5, 0xb3, 0x14,
	0xb5, 0xb1, 0x59, 0x40, 0xc8, 0x4d, 0x9e, 0xd1, 0xda, 0x7b, 0x76, 0xad, 0x2f, 0x5e, 0xe5, 0x0c,
	0xa3, 0x5d, 0xc4, 0xc8, 0x84, 0xdd, 0xdd, 0xc7, 0xc8, 0x24, 0xb3, 0xc8, 0x24, 0xcf, 0x91, 0x49,
	0xee, 0x17, 0x66, 0x69, 0xb6, 0x30, 0x4b, 0x4f, 0x0b, 0xb3, 0x74, 0xbe, 0xc1, 0x87, 0x78, 0x39,
	0x61, 0xf6, 0x40, 0x5c, 0x3b, 0x8a, 0xc3, 0xd6, 0x9b, 0x5c, 0x5c, 0x3b, 0xd3, 0x74, 0xd3, 0x6f,
	0x43, 0x50, 0xac, 0x92, 0xec, 0xe8, 0xf6, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x7f, 0xeb,
	0x97, 0x01, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PlaceBet defines a method to place a bet with the given data
	PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error)
	// SettleBet defines a method to settle the given bet
	SettleBet(ctx context.Context, in *MsgSettleBet, opts ...grpc.CallOption) (*MsgSettleBetResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error) {
	out := new(MsgPlaceBetResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.bet.Msg/PlaceBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SettleBet(ctx context.Context, in *MsgSettleBet, opts ...grpc.CallOption) (*MsgSettleBetResponse, error) {
	out := new(MsgSettleBetResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.bet.Msg/SettleBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PlaceBet defines a method to place a bet with the given data
	PlaceBet(context.Context, *MsgPlaceBet) (*MsgPlaceBetResponse, error)
	// SettleBet defines a method to settle the given bet
	SettleBet(context.Context, *MsgSettleBet) (*MsgSettleBetResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlaceBet(ctx context.Context, req *MsgPlaceBet) (*MsgPlaceBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBet not implemented")
}
func (*UnimplementedMsgServer) SettleBet(ctx context.Context, req *MsgSettleBet) (*MsgSettleBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettleBet not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlaceBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceBet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.bet.Msg/PlaceBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceBet(ctx, req.(*MsgPlaceBet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SettleBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSettleBet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SettleBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.bet.Msg/SettleBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SettleBet(ctx, req.(*MsgSettleBet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgenetwork.sge.bet.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler:    _Msg_PlaceBet_Handler,
		},
		{
			MethodName: "SettleBet",
			Handler:    _Msg_SettleBet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sge/bet/tx.proto",
}

func (m *MsgPlaceBet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bet != nil {
		{
			size, err := m.Bet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceBetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bet != nil {
		{
			size, err := m.Bet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSettleBet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSettleBet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSettleBet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BettorAddress) > 0 {
		i -= len(m.BettorAddress)
		copy(dAtA[i:], m.BettorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BettorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BetUID) > 0 {
		i -= len(m.BetUID)
		copy(dAtA[i:], m.BetUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BetUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSettleBetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSettleBetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSettleBetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BetUID) > 0 {
		i -= len(m.BetUID)
		copy(dAtA[i:], m.BetUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BetUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPlaceBet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bet != nil {
		l = m.Bet.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPlaceBetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bet != nil {
		l = m.Bet.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSettleBet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BetUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BettorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSettleBetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BetUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPlaceBet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPlaceBet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bet == nil {
				m.Bet = &PlaceBetFields{}
			}
			if err := m.Bet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPlaceBetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPlaceBetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bet == nil {
				m.Bet = &PlaceBetFields{}
			}
			if err := m.Bet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSettleBet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSettleBet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSettleBet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BetUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BettorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BettorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSettleBetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSettleBetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSettleBetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BetUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
