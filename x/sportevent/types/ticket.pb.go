// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/sportevent/ticket.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// SportEventAddTicketPayload indicates data of add sport-event ticket
type SportEventAddTicketPayload struct {
	// uid is the universal unique identifier of the sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the sport-event.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// status is the current status of the sport-event.
	Status SportEventStatus `protobuf:"varint,5,opt,name=status,proto3,enum=sgenetwork.sge.sportevent.SportEventStatus" json:"status,omitempty"`
	// creator is the address of the creator of sport-event.
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// active is the status of active or inactive sport-event.
	Active bool `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
	// meta contains human-readable metadata of the sport-event.
	Meta string `protobuf:"bytes,10,opt,name=meta,proto3" json:"meta,omitempty"`
	// sr_contribution_for_house is the portion of payout that should be paid by sr
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
}

func (m *SportEventAddTicketPayload) Reset()         { *m = SportEventAddTicketPayload{} }
func (m *SportEventAddTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventAddTicketPayload) ProtoMessage()    {}
func (*SportEventAddTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{0}
}
func (m *SportEventAddTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventAddTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventAddTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventAddTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventAddTicketPayload.Merge(m, src)
}
func (m *SportEventAddTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventAddTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventAddTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventAddTicketPayload proto.InternalMessageInfo

func (m *SportEventAddTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *SportEventAddTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func (m *SportEventAddTicketPayload) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *SportEventAddTicketPayload) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// SportEventUpdateTicketPayload indicates data of update sport-event ticket
type SportEventUpdateTicketPayload struct {
	// uid is the uuid of the sport-event
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// active is the status of active or inactive sport-event
	Active bool `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *SportEventUpdateTicketPayload) Reset()         { *m = SportEventUpdateTicketPayload{} }
func (m *SportEventUpdateTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventUpdateTicketPayload) ProtoMessage()    {}
func (*SportEventUpdateTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{1}
}
func (m *SportEventUpdateTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventUpdateTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventUpdateTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventUpdateTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventUpdateTicketPayload.Merge(m, src)
}
func (m *SportEventUpdateTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventUpdateTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventUpdateTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventUpdateTicketPayload proto.InternalMessageInfo

func (m *SportEventUpdateTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventUpdateTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

// SportEventResolutionTicketPayload indicates data of the
// resolution of the sport-event ticket.
type SportEventResolutionTicketPayload struct {
	// uid is the universal unique identifier of sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// resolution_ts is the resolution timestamp of the sport-event.
	ResolutionTS uint64 `protobuf:"varint,2,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// winner_odds_uids is the universal unique identifier list of the winner
	// odds.
	WinnerOddsUIDs []string `protobuf:"bytes,3,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the status of the resolution.
	Status SportEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=sgenetwork.sge.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *SportEventResolutionTicketPayload) Reset()         { *m = SportEventResolutionTicketPayload{} }
func (m *SportEventResolutionTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventResolutionTicketPayload) ProtoMessage()    {}
func (*SportEventResolutionTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e970e0fce71fcc, []int{2}
}
func (m *SportEventResolutionTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventResolutionTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventResolutionTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventResolutionTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventResolutionTicketPayload.Merge(m, src)
}
func (m *SportEventResolutionTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventResolutionTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventResolutionTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventResolutionTicketPayload proto.InternalMessageInfo

func (m *SportEventResolutionTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventResolutionTicketPayload) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *SportEventResolutionTicketPayload) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *SportEventResolutionTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SportEventAddTicketPayload)(nil), "sgenetwork.sge.sportevent.SportEventAddTicketPayload")
	proto.RegisterType((*SportEventUpdateTicketPayload)(nil), "sgenetwork.sge.sportevent.SportEventUpdateTicketPayload")
	proto.RegisterType((*SportEventResolutionTicketPayload)(nil), "sgenetwork.sge.sportevent.SportEventResolutionTicketPayload")
}

func init() { proto.RegisterFile("sge/sportevent/ticket.proto", fileDescriptor_b7e970e0fce71fcc) }

var fileDescriptor_b7e970e0fce71fcc = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x4e, 0x1b, 0x3d,
	0x14, 0xcd, 0x90, 0x64, 0x02, 0x86, 0x2f, 0xfa, 0xe4, 0x56, 0x74, 0xa0, 0x6a, 0x3c, 0xb0, 0x40,
	0x91, 0x10, 0x13, 0x09, 0x9e, 0x80, 0xf0, 0x57, 0x56, 0xad, 0x1c, 0x50, 0xa5, 0x6e, 0x46, 0x93,
	0xf8, 0x32, 0x8c, 0x20, 0xe3, 0xc8, 0xf6, 0x40, 0x79, 0x8b, 0xbe, 0x43, 0x55, 0xf5, 0x55, 0x58,
	0xb2, 0xac, 0x58, 0x58, 0xd5, 0xb0, 0xcb, 0x53, 0x54, 0x36, 0x53, 0x12, 0x40, 0xa9, 0x04, 0x62,
	0xd1, 0x4d, 0xe6, 0xfa, 0xde, 0x73, 0x4e, 0xec, 0xe3, 0xeb, 0x8b, 0xde, 0xca, 0x18, 0x5a, 0x72,
	0xc0, 0x85, 0x82, 0x33, 0x48, 0x55, 0x4b, 0x25, 0xbd, 0x13, 0x50, 0xc1, 0x40, 0x70, 0xc5, 0xf1,
	0x82, 0x8c, 0x21, 0x05, 0x75, 0xce, 0xc5, 0x49, 0x20, 0x63, 0x08, 0x46, 0xb8, 0x45, 0xff, 0x01,
	0xcf, 0x86, 0xa1, 0x8d, 0x6f, 0xc9, 0x8b, 0x0b, 0x0f, 0x10, 0x9c, 0x31, 0x59, 0x94, 0x5e, 0xc7,
	0x3c, 0xe6, 0x36, 0x6c, 0x99, 0xe8, 0x36, 0xbb, 0xfc, 0xbd, 0x8a, 0x16, 0x3b, 0x06, 0xbf, 0x63,
	0xf0, 0x9b, 0x8c, 0x1d, 0xd8, 0xbd, 0x7c, 0x8c, 0x2e, 0x4e, 0x79, 0xc4, 0xb0, 0x8f, 0xca, 0x59,
	0xc2, 0x3c, 0xc7, 0x77, 0x9a, 0x33, 0xed, 0x7a, 0xae, 0x49, 0xf9, 0x70, 0x7f, 0x7b, 0xa8, 0x89,
	0xc9, 0x52, 0xf3, 0x83, 0x37, 0xd0, 0xb4, 0x54, 0x91, 0x50, 0xa1, 0x92, 0xde, 0x94, 0xef, 0x34,
	0x2b, 0xed, 0x37, 0xb9, 0x26, 0xb5, 0x8e, 0xc9, 0x1d, 0x74, 0x86, 0x9a, 0xdc, 0x95, 0xe9, 0x5d,
	0x84, 0x57, 0x91, 0x0b, 0x29, 0x33, 0x94, 0xb2, 0xa5, 0xbc, 0xca, 0x35, 0xa9, 0xee, 0xa4, 0xcc,
	0x12, 0x8a, 0x12, 0x2d, 0xbe, 0x78, 0x03, 0x55, 0xcc, 0x31, 0xbc, 0x8a, 0x5f, 0x6e, 0xce, 0xae,
	0x93, 0x60, 0xa2, 0x3f, 0xc1, 0x07, 0xc6, 0x24, 0xb5, 0x60, 0xbc, 0x85, 0x5c, 0xa9, 0x22, 0x95,
	0x49, 0xaf, 0xea, 0x3b, 0xcd, 0xfa, 0xfa, 0xea, 0x5f, 0x68, 0xa3, 0xf3, 0x77, 0x2c, 0x85, 0x16,
	0x54, 0xec, 0xa1, 0x5a, 0x4f, 0x40, 0xa4, 0xb8, 0xf0, 0x5c, 0xe3, 0x00, 0xfd, 0xb3, 0xc4, 0x07,
	0xa8, 0xde, 0x4f, 0xd2, 0xb0, 0x0b, 0x2a, 0x8c, 0xfa, 0x3c, 0x4b, 0x95, 0x57, 0xb3, 0x16, 0x05,
	0x97, 0x9a, 0x94, 0xae, 0x35, 0x59, 0x89, 0x13, 0x75, 0x9c, 0x75, 0x83, 0x1e, 0xef, 0xb7, 0x7a,
	0x5c, 0xf6, 0xb9, 0x2c, 0x3e, 0x6b, 0x92, 0x9d, 0xb4, 0xd4, 0xc5, 0x00, 0x64, 0xb0, 0x9f, 0x2a,
	0x3a, 0xd7, 0x4f, 0xd2, 0x36, 0xa8, 0x4d, 0xab, 0x81, 0xf7, 0x50, 0xcd, 0x28, 0x1e, 0x01, 0x78,
	0xd3, 0xcf, 0x92, 0x73, 0xbb, 0xa0, 0x76, 0x01, 0xf0, 0x3c, 0x72, 0xa3, 0x9e, 0x4a, 0xce, 0xc0,
	0x9b, 0xf1, 0x9d, 0xe6, 0x34, 0x2d, 0x56, 0x18, 0xa3, 0x4a, 0x1f, 0x54, 0xe4, 0x21, 0x7b, 0x1a,
	0x1b, 0xe3, 0x6f, 0x0e, 0x5a, 0x90, 0x22, 0xec, 0xf1, 0x54, 0x89, 0xa4, 0x9b, 0xa9, 0x84, 0xa7,
	0xe1, 0x11, 0x17, 0xe1, 0x31, 0xcf, 0x24, 0x78, 0xb3, 0x76, 0x1f, 0xf0, 0xb4, 0x7d, 0xe4, 0x9a,
	0xcc, 0x77, 0xc4, 0xd6, 0x98, 0xe2, 0x2e, 0x17, 0xef, 0x8d, 0xde, 0x50, 0x93, 0xc9, 0x7f, 0x46,
	0x27, 0x97, 0x96, 0xaf, 0xa7, 0xd0, 0xbb, 0xd1, 0x3d, 0x1d, 0x0e, 0x58, 0xa4, 0xe0, 0xdf, 0x6b,
	0xd5, 0xc7, 0x6d, 0x51, 0x79, 0xd9, 0xb6, 0xa8, 0xbe, 0x50, 0x5b, 0xb8, 0xe3, 0x6d, 0xb1, 0xfc,
	0x63, 0x0a, 0x2d, 0x8d, 0xcc, 0xa5, 0x20, 0xf9, 0xa9, 0xf5, 0xff, 0xa9, 0x06, 0xef, 0xa1, 0xff,
	0xc4, 0x1d, 0x79, 0xe4, 0xf2, 0x52, 0xae, 0xc9, 0xdc, 0x98, 0xaa, 0x71, 0xee, 0x3e, 0x90, 0xde,
	0x5f, 0x62, 0x8a, 0xfe, 0x3f, 0x4f, 0xd2, 0x14, 0x44, 0x68, 0x1e, 0x73, 0x98, 0x25, 0xcc, 0xd8,
	0x5f, 0x6e, 0xce, 0xb4, 0x57, 0x72, 0x4d, 0xea, 0x9f, 0x6c, 0xcd, 0xbc, 0xf6, 0xc3, 0xfd, 0x6d,
	0x39, 0xd4, 0xe4, 0x11, 0x9a, 0x3e, 0xca, 0x8c, 0x4d, 0x84, 0xca, 0xb3, 0x27, 0x42, 0x7b, 0xef,
	0x32, 0x6f, 0x38, 0x57, 0x79, 0xc3, 0xf9, 0x95, 0x37, 0x9c, 0xaf, 0x37, 0x8d, 0xd2, 0xd5, 0x4d,
	0xa3, 0xf4, 0xf3, 0xa6, 0x51, 0xfa, 0xbc, 0x36, 0x76, 0x17, 0x32, 0x86, 0xb5, 0x42, 0xd9, 0xc4,
	0xad, 0x2f, 0xf7, 0x86, 0xbd, 0xb9, 0x96, 0xae, 0x6b, 0xc7, 0xef, 0xc6, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x4d, 0xe9, 0x18, 0x0b, 0x06, 0x00, 0x00,
}

func (m *SportEventAddTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventAddTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventAddTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x52
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventUpdateTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventUpdateTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventUpdateTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventResolutionTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventResolutionTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventResolutionTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SportEventAddTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.Active {
		n += 2
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func (m *SportEventUpdateTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.Active {
		n += 2
	}
	return n
}

func (m *SportEventResolutionTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovTicket(uint64(m.ResolutionTS))
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SportEventAddTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventAddTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventAddTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Odds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Odds = append(m.Odds, &Odds{})
			if err := m.Odds[len(m.Odds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.Active = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *SportEventUpdateTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *SportEventResolutionTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolutionTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
