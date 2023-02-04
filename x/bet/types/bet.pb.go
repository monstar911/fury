// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/bet.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Status of the Bet.
type Bet_Status int32

const (
	// the unknown status
	Bet_STATUS_INVALID Bet_Status = 0
	// bet is placed
	Bet_STATUS_PLACED Bet_Status = 1
	// bet is canceled by Bettor
	Bet_STATUS_CANCELLED Bet_Status = 2
	// bet is aborted
	Bet_STATUS_ABORTED Bet_Status = 3
	// pending for getting placed
	Bet_STATUS_PENDING Bet_Status = 4
	// bet result is declared
	Bet_STATUS_RESULT_DECLARED Bet_Status = 5
	// the bet is settled
	Bet_STATUS_SETTLED Bet_Status = 6
)

var Bet_Status_name = map[int32]string{
	0: "STATUS_INVALID",
	1: "STATUS_PLACED",
	2: "STATUS_CANCELLED",
	3: "STATUS_ABORTED",
	4: "STATUS_PENDING",
	5: "STATUS_RESULT_DECLARED",
	6: "STATUS_SETTLED",
}

var Bet_Status_value = map[string]int32{
	"STATUS_INVALID":         0,
	"STATUS_PLACED":          1,
	"STATUS_CANCELLED":       2,
	"STATUS_ABORTED":         3,
	"STATUS_PENDING":         4,
	"STATUS_RESULT_DECLARED": 5,
	"STATUS_SETTLED":         6,
}

func (x Bet_Status) String() string {
	return proto.EnumName(Bet_Status_name, int32(x))
}

func (Bet_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bc076bb1a4d9f6e, []int{0, 0}
}

// Result of the bet.
type Bet_Result int32

const (
	// the invalid or unknown
	Bet_RESULT_INVALID Bet_Result = 0
	// bet result is pending
	Bet_RESULT_PENDING Bet_Result = 1
	// bet won by the bettor
	Bet_RESULT_WON Bet_Result = 2
	// bet lost by the bettor
	Bet_RESULT_LOST Bet_Result = 3
	// bet is draw
	Bet_RESULT_DRAW Bet_Result = 4
	// bet is aborted
	Bet_RESULT_ABORTED Bet_Result = 5
)

var Bet_Result_name = map[int32]string{
	0: "RESULT_INVALID",
	1: "RESULT_PENDING",
	2: "RESULT_WON",
	3: "RESULT_LOST",
	4: "RESULT_DRAW",
	5: "RESULT_ABORTED",
}

var Bet_Result_value = map[string]int32{
	"RESULT_INVALID": 0,
	"RESULT_PENDING": 1,
	"RESULT_WON":     2,
	"RESULT_LOST":    3,
	"RESULT_DRAW":    4,
	"RESULT_ABORTED": 5,
}

func (x Bet_Result) String() string {
	return proto.EnumName(Bet_Result_name, int32(x))
}

func (Bet_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bc076bb1a4d9f6e, []int{0, 1}
}

type Bet struct {
	// uid is the unique uuid assigned to bet
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// sport_event_uid is the unique uuid of the sportevent on which bet is placed
	SportEventUID string `protobuf:"bytes,2,opt,name=sport_event_uid,proto3" json:"sport_event_uid"`
	// odds_uid is the unique uuid of the odds on which bet is placed
	OddsUID string `protobuf:"bytes,3,opt,name=odds_uid,proto3" json:"odds_uid"`
	// odds_value is the odds on which bet is placed
	OddsValue string `protobuf:"bytes,4,opt,name=odds_value,json=oddsValue,proto3" json:"odds_value,omitempty"`
	// amount is the wager amount
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// betFee is the betting fee
	BetFee types.Coin `protobuf:"bytes,6,opt,name=bet_fee,json=betFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"bet_fee"`
	// status is the status of the bet, such as `pending` or `settled`
	Status Bet_Status `protobuf:"varint,7,opt,name=status,proto3,enum=sgenetwork.sge.bet.Bet_Status" json:"status,omitempty"`
	// result is the result of bet, such as `won` or `lost`
	Result Bet_Result `protobuf:"varint,8,opt,name=result,proto3,enum=sgenetwork.sge.bet.Bet_Result" json:"result,omitempty"`
	// verified shows bet is verified or not
	Verified bool `protobuf:"varint,9,opt,name=verified,proto3" json:"verified,omitempty"`
	// ticket is a signed string containing important info such as `oddsValue`
	Ticket string `protobuf:"bytes,10,opt,name=ticket,proto3" json:"ticket,omitempty"`
	// creator is the bettor address
	Creator string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
	// created_at is bet placement timestamp
	CreatedAt int64 `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// odds_type is the type of odds user chose such as decimal, fractional
	OddsType OddsType `protobuf:"varint,13,opt,name=odds_type,json=oddsType,proto3,enum=sgenetwork.sge.bet.OddsType" json:"odds_type,omitempty"`
}

func (m *Bet) Reset()         { *m = Bet{} }
func (m *Bet) String() string { return proto.CompactTextString(m) }
func (*Bet) ProtoMessage()    {}
func (*Bet) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc076bb1a4d9f6e, []int{0}
}
func (m *Bet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bet.Merge(m, src)
}
func (m *Bet) XXX_Size() int {
	return m.Size()
}
func (m *Bet) XXX_DiscardUnknown() {
	xxx_messageInfo_Bet.DiscardUnknown(m)
}

var xxx_messageInfo_Bet proto.InternalMessageInfo

func (m *Bet) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Bet) GetSportEventUID() string {
	if m != nil {
		return m.SportEventUID
	}
	return ""
}

func (m *Bet) GetOddsUID() string {
	if m != nil {
		return m.OddsUID
	}
	return ""
}

func (m *Bet) GetOddsValue() string {
	if m != nil {
		return m.OddsValue
	}
	return ""
}

func (m *Bet) GetBetFee() types.Coin {
	if m != nil {
		return m.BetFee
	}
	return types.Coin{}
}

func (m *Bet) GetStatus() Bet_Status {
	if m != nil {
		return m.Status
	}
	return Bet_STATUS_INVALID
}

func (m *Bet) GetResult() Bet_Result {
	if m != nil {
		return m.Result
	}
	return Bet_RESULT_INVALID
}

func (m *Bet) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Bet) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *Bet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Bet) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Bet) GetOddsType() OddsType {
	if m != nil {
		return m.OddsType
	}
	return OddsType_ODD_TYPE_INVALID
}

type UID2ID struct {
	// uid is the unique uuid assigned to bet
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// id is a autogenerated sequencial id for a bet
	ID uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id"`
}

func (m *UID2ID) Reset()         { *m = UID2ID{} }
func (m *UID2ID) String() string { return proto.CompactTextString(m) }
func (*UID2ID) ProtoMessage()    {}
func (*UID2ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc076bb1a4d9f6e, []int{1}
}
func (m *UID2ID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UID2ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UID2ID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UID2ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UID2ID.Merge(m, src)
}
func (m *UID2ID) XXX_Size() int {
	return m.Size()
}
func (m *UID2ID) XXX_DiscardUnknown() {
	xxx_messageInfo_UID2ID.DiscardUnknown(m)
}

var xxx_messageInfo_UID2ID proto.InternalMessageInfo

func (m *UID2ID) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *UID2ID) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterEnum("sgenetwork.sge.bet.Bet_Status", Bet_Status_name, Bet_Status_value)
	proto.RegisterEnum("sgenetwork.sge.bet.Bet_Result", Bet_Result_name, Bet_Result_value)
	proto.RegisterType((*Bet)(nil), "sgenetwork.sge.bet.Bet")
	proto.RegisterType((*UID2ID)(nil), "sgenetwork.sge.bet.UID2ID")
}

func init() { proto.RegisterFile("sge/bet/bet.proto", fileDescriptor_9bc076bb1a4d9f6e) }

var fileDescriptor_9bc076bb1a4d9f6e = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6e, 0xda, 0x4c,
	0x10, 0xc7, 0x40, 0x0c, 0x6c, 0x12, 0xe2, 0xac, 0xa2, 0xc4, 0x1f, 0xca, 0x67, 0x23, 0x54, 0xb5,
	0x5c, 0x62, 0x2b, 0x89, 0x54, 0xa9, 0xb7, 0x1a, 0xec, 0xb4, 0x48, 0x08, 0xa2, 0xc5, 0x24, 0x52,
	0x2f, 0x08, 0xe3, 0x0d, 0xb5, 0x92, 0xb0, 0x88, 0x5d, 0x68, 0xf3, 0x16, 0x7d, 0x81, 0xde, 0xab,
	0x3e, 0x49, 0x8e, 0x39, 0x56, 0x3d, 0xd0, 0xca, 0xb9, 0xf5, 0x29, 0xaa, 0x5d, 0x2f, 0x04, 0xa5,
	0x55, 0x95, 0x03, 0x30, 0xf3, 0xfb, 0x33, 0x3b, 0xcc, 0x78, 0x0d, 0xb6, 0xe9, 0x10, 0xdb, 0x01,
	0x66, 0xfc, 0x63, 0x8d, 0x27, 0x84, 0x11, 0x08, 0xe9, 0x10, 0x8f, 0x30, 0xfb, 0x40, 0x26, 0x97,
	0x16, 0x1d, 0x62, 0x2b, 0xc0, 0xac, 0xb4, 0x33, 0x24, 0x43, 0x22, 0x68, 0x9b, 0x47, 0x89, 0xb2,
	0x64, 0x0c, 0x08, 0xbd, 0x26, 0xd4, 0x0e, 0xfa, 0x14, 0xdb, 0xb3, 0xc3, 0x00, 0xb3, 0xfe, 0xa1,
	0x3d, 0x20, 0xd1, 0x48, 0xf2, 0x7b, 0x8b, 0xe2, 0x24, 0x0c, 0x69, 0x8f, 0xdd, 0x8c, 0x71, 0x42,
	0x54, 0xbe, 0xe4, 0x40, 0xa6, 0x86, 0x19, 0x2c, 0x83, 0xcc, 0x34, 0x0a, 0x75, 0xa5, 0xac, 0x54,
	0x0b, 0xb5, 0x62, 0x3c, 0x37, 0x33, 0xdd, 0x86, 0xfb, 0x6b, 0x6e, 0x72, 0x14, 0xf1, 0x2f, 0xd8,
	0x02, 0x5b, 0x74, 0x4c, 0x26, 0xac, 0x87, 0x67, 0x78, 0xc4, 0x7a, 0x5c, 0x9d, 0x16, 0xea, 0x67,
	0xf1, 0xdc, 0xdc, 0xec, 0x70, 0xca, 0xe3, 0x4c, 0xe2, 0x7b, 0xac, 0x45, 0x8f, 0x01, 0x78, 0x0c,
	0xf2, 0xa2, 0x19, 0x5e, 0x28, 0x23, 0x0a, 0xed, 0xc5, 0x73, 0x33, 0xd7, 0x0e, 0x43, 0x9a, 0x94,
	0x58, 0xd2, 0x68, 0x19, 0xc1, 0xff, 0x01, 0x10, 0xf1, 0xac, 0x7f, 0x35, 0xc5, 0x7a, 0x96, 0xdb,
	0x50, 0x81, 0x23, 0x67, 0x1c, 0x80, 0x27, 0x40, 0xed, 0x5f, 0x93, 0xe9, 0x88, 0xe9, 0x6b, 0xa2,
	0xa2, 0x75, 0x3b, 0x37, 0x53, 0xdf, 0xe7, 0xe6, 0xf3, 0x61, 0xc4, 0xde, 0x4f, 0x03, 0x6b, 0x40,
	0xae, 0x6d, 0x39, 0xa9, 0xe4, 0xe7, 0x80, 0x86, 0x97, 0x36, 0x1f, 0x07, 0xb5, 0x1a, 0x23, 0x86,
	0xa4, 0x1b, 0x0e, 0x40, 0x2e, 0xc0, 0xac, 0x77, 0x81, 0xb1, 0xae, 0x96, 0x95, 0xea, 0xfa, 0xd1,
	0x7f, 0x56, 0xa2, 0xb7, 0xf8, 0x80, 0x2d, 0x39, 0x60, 0xab, 0x4e, 0xa2, 0x51, 0xcd, 0xe6, 0x67,
	0x7c, 0xfd, 0x61, 0xbe, 0x78, 0xc2, 0x19, 0xdc, 0x80, 0xd4, 0x00, 0xb3, 0x13, 0x8c, 0xe1, 0x4b,
	0xa0, 0x52, 0xd6, 0x67, 0x53, 0xaa, 0xe7, 0xca, 0x4a, 0xb5, 0x78, 0x64, 0x58, 0x7f, 0xae, 0xdb,
	0xaa, 0x61, 0x66, 0x75, 0x84, 0x0a, 0x49, 0x35, 0xf7, 0x4d, 0x30, 0x9d, 0x5e, 0x31, 0x3d, 0xff,
	0x6f, 0x1f, 0x12, 0x2a, 0x24, 0xd5, 0xb0, 0x04, 0xf2, 0x33, 0x3c, 0x89, 0x2e, 0x22, 0x1c, 0xea,
	0x85, 0xb2, 0x52, 0xcd, 0xa3, 0x65, 0x0e, 0x77, 0x81, 0xca, 0xa2, 0xc1, 0x25, 0x66, 0x3a, 0x10,
	0x33, 0x95, 0x19, 0xd4, 0x41, 0x6e, 0x30, 0xc1, 0x7d, 0x46, 0x26, 0xfa, 0xba, 0x20, 0x16, 0x29,
	0xdf, 0x84, 0x08, 0x71, 0xd8, 0xeb, 0x33, 0x7d, 0xa3, 0xac, 0x54, 0x33, 0xa8, 0x20, 0x11, 0x87,
	0xc1, 0x57, 0xa0, 0xb0, 0x7c, 0xd4, 0xf4, 0x4d, 0xd1, 0xe7, 0xfe, 0xdf, 0xfa, 0xe4, 0xeb, 0xf6,
	0x6f, 0xc6, 0x38, 0xd9, 0x31, 0x8f, 0x2a, 0x9f, 0x15, 0xa0, 0x26, 0x7f, 0x19, 0x42, 0x50, 0xec,
	0xf8, 0x8e, 0xdf, 0xed, 0xf4, 0x1a, 0xad, 0x33, 0xa7, 0xd9, 0x70, 0xb5, 0x14, 0xdc, 0x06, 0x9b,
	0x12, 0x3b, 0x6d, 0x3a, 0x75, 0xcf, 0xd5, 0x14, 0xb8, 0x03, 0x34, 0x09, 0xd5, 0x9d, 0x56, 0xdd,
	0x6b, 0x36, 0x3d, 0x57, 0x4b, 0xaf, 0x98, 0x9d, 0x5a, 0x1b, 0xf9, 0x9e, 0xab, 0x65, 0x56, 0xb0,
	0x53, 0xaf, 0xe5, 0x36, 0x5a, 0x6f, 0xb4, 0x2c, 0x2c, 0x81, 0x5d, 0x89, 0x21, 0xaf, 0xd3, 0x6d,
	0xfa, 0x3d, 0xd7, 0xab, 0x37, 0x1d, 0xe4, 0xb9, 0xda, 0xda, 0x8a, 0xbe, 0xe3, 0xf9, 0x3e, 0xaf,
	0xab, 0x56, 0x66, 0x40, 0x4d, 0x26, 0xcb, 0x59, 0x69, 0x79, 0x68, 0xef, 0x01, 0x5b, 0x9c, 0xa0,
	0xc0, 0x22, 0x00, 0x12, 0x3b, 0x6f, 0xb7, 0xb4, 0x34, 0xdc, 0x02, 0xeb, 0x32, 0x6f, 0xb6, 0x3b,
	0xbe, 0x96, 0x59, 0x01, 0x5c, 0xe4, 0x9c, 0x6b, 0xd9, 0x95, 0x2a, 0x8b, 0xde, 0xd7, 0x2a, 0x6f,
	0x81, 0xda, 0x6d, 0xb8, 0x47, 0x0d, 0xf7, 0x09, 0x97, 0x75, 0x1f, 0xa4, 0xe5, 0xfd, 0xcc, 0xd6,
	0x36, 0xe2, 0xb9, 0x99, 0x16, 0x7c, 0x3a, 0x0a, 0x51, 0x3a, 0x0a, 0x6b, 0xaf, 0x6f, 0x63, 0x43,
	0xb9, 0x8b, 0x0d, 0xe5, 0x67, 0x6c, 0x28, 0x9f, 0xee, 0x8d, 0xd4, 0xdd, 0xbd, 0x91, 0xfa, 0x76,
	0x6f, 0xa4, 0xde, 0xad, 0x5e, 0x14, 0x3a, 0xc4, 0x07, 0x72, 0x5d, 0x3c, 0xb6, 0x3f, 0x8a, 0x17,
	0x88, 0x78, 0x90, 0x03, 0x55, 0xbc, 0x3d, 0x8e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xf9,
	0x9d, 0x0e, 0xb5, 0x04, 0x00, 0x00,
}

func (m *Bet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OddsType != 0 {
		i = encodeVarintBet(dAtA, i, uint64(m.OddsType))
		i--
		dAtA[i] = 0x68
	}
	if m.CreatedAt != 0 {
		i = encodeVarintBet(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBet(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintBet(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x52
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Result != 0 {
		i = encodeVarintBet(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x40
	}
	if m.Status != 0 {
		i = encodeVarintBet(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.BetFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBet(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBet(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.OddsValue) > 0 {
		i -= len(m.OddsValue)
		copy(dAtA[i:], m.OddsValue)
		i = encodeVarintBet(dAtA, i, uint64(len(m.OddsValue)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OddsUID) > 0 {
		i -= len(m.OddsUID)
		copy(dAtA[i:], m.OddsUID)
		i = encodeVarintBet(dAtA, i, uint64(len(m.OddsUID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintBet(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBet(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UID2ID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UID2ID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UID2ID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintBet(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBet(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBet(dAtA []byte, offset int, v uint64) int {
	offset -= sovBet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	l = len(m.OddsUID)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	l = len(m.OddsValue)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovBet(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovBet(uint64(l))
	if m.Status != 0 {
		n += 1 + sovBet(uint64(m.Status))
	}
	if m.Result != 0 {
		n += 1 + sovBet(uint64(m.Result))
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovBet(uint64(m.CreatedAt))
	}
	if m.OddsType != 0 {
		n += 1 + sovBet(uint64(m.OddsType))
	}
	return n
}

func (m *UID2ID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBet(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovBet(uint64(m.ID))
	}
	return n
}

func sovBet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBet(x uint64) (n int) {
	return sovBet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBet
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
			return fmt.Errorf("proto: Bet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportEventUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Bet_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= Bet_Result(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsType", wireType)
			}
			m.OddsType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsType |= OddsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBet
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
func (m *UID2ID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBet
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
			return fmt.Errorf("proto: UID2ID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UID2ID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
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
				return ErrInvalidLengthBet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBet
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
func skipBet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBet
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
					return 0, ErrIntOverflowBet
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
					return 0, ErrIntOverflowBet
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
				return 0, ErrInvalidLengthBet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBet = fmt.Errorf("proto: unexpected end of group")
)
