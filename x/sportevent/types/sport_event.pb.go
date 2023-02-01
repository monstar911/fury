// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/sportevent/sport_event.proto

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

// SportEventStatus is the sport event status enumeration
type SportEventStatus int32

const (
	// event is pending
	SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED SportEventStatus = 0
	// invalid event
	SportEventStatus_SPORT_EVENT_STATUS_INVALID SportEventStatus = 1
	// event cancelled
	SportEventStatus_SPORT_EVENT_STATUS_CANCELLED SportEventStatus = 2
	// event aborted
	SportEventStatus_SPORT_EVENT_STATUS_ABORTED SportEventStatus = 3
	// result of the event is declared
	SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED SportEventStatus = 4
)

var SportEventStatus_name = map[int32]string{
	0: "SPORT_EVENT_STATUS_UNSPECIFIED",
	1: "SPORT_EVENT_STATUS_INVALID",
	2: "SPORT_EVENT_STATUS_CANCELLED",
	3: "SPORT_EVENT_STATUS_ABORTED",
	4: "SPORT_EVENT_STATUS_RESULT_DECLARED",
}

var SportEventStatus_value = map[string]int32{
	"SPORT_EVENT_STATUS_UNSPECIFIED":     0,
	"SPORT_EVENT_STATUS_INVALID":         1,
	"SPORT_EVENT_STATUS_CANCELLED":       2,
	"SPORT_EVENT_STATUS_ABORTED":         3,
	"SPORT_EVENT_STATUS_RESULT_DECLARED": 4,
}

func (x SportEventStatus) String() string {
	return proto.EnumName(SportEventStatus_name, int32(x))
}

func (SportEventStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{0}
}

// SportEvent the representation of the sport event to be stored in sport-event
// state
type SportEvent struct {
	// UID is the uuid of the sport-event
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// StartTS is the start timestamp of the sport-event
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// EndTS is the end timestamp of the sport-event
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// Odds is the list of odds of the sport-event
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// WinnerOddsUIDs is the list of winner odds uuid
	WinnerOddsUIDs []string `protobuf:"bytes,5,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// Status is the current status of the sport-event
	Status SportEventStatus `protobuf:"varint,6,opt,name=status,proto3,enum=sgenetwork.sge.sportevent.SportEventStatus" json:"status,omitempty"`
	// ResolutionTS is the timestamp of respolution of sport-event
	ResolutionTS uint64 `protobuf:"varint,7,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// Creator is the address of creator of sport-event
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	// BetConstraints holds the constraints of sport-event to accept bets
	BetConstraints *EventBetConstraints `protobuf:"bytes,9,opt,name=bet_constraints,json=betConstraints,proto3" json:"bet_constraints,omitempty"`
	// Active is the status of active or inactive sport-event
	Active bool `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	// Meta contains human-readable metadata of the sport-event
	Meta string `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (m *SportEvent) Reset()         { *m = SportEvent{} }
func (m *SportEvent) String() string { return proto.CompactTextString(m) }
func (*SportEvent) ProtoMessage()    {}
func (*SportEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{0}
}
func (m *SportEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEvent.Merge(m, src)
}
func (m *SportEvent) XXX_Size() int {
	return m.Size()
}
func (m *SportEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SportEvent proto.InternalMessageInfo

func (m *SportEvent) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEvent) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEvent) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEvent) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *SportEvent) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *SportEvent) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func (m *SportEvent) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *SportEvent) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SportEvent) GetBetConstraints() *EventBetConstraints {
	if m != nil {
		return m.BetConstraints
	}
	return nil
}

func (m *SportEvent) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *SportEvent) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// ResolutionEvent is type of resolution for sport-event
type ResolutionEvent struct {
	// UID is the uuid of sport-event
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// ResolutionTS is the resolution timestamp of event
	ResolutionTS uint64 `protobuf:"varint,2,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// WinnerOddsUIDs is the uuid list of winner odds
	WinnerOddsUIDs []string `protobuf:"bytes,3,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// Status is the status of resolution
	Status SportEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=sgenetwork.sge.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *ResolutionEvent) Reset()         { *m = ResolutionEvent{} }
func (m *ResolutionEvent) String() string { return proto.CompactTextString(m) }
func (*ResolutionEvent) ProtoMessage()    {}
func (*ResolutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{1}
}
func (m *ResolutionEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResolutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResolutionEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResolutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolutionEvent.Merge(m, src)
}
func (m *ResolutionEvent) XXX_Size() int {
	return m.Size()
}
func (m *ResolutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResolutionEvent proto.InternalMessageInfo

func (m *ResolutionEvent) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *ResolutionEvent) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *ResolutionEvent) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *ResolutionEvent) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

// Bet constraints parent group for a sport event
type EventBetConstraints struct {
	// MinAmount is the minimum allowed bet amount for a sport-event
	MinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_amount,json=minAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_amount"`
	// BetFee is the fee that bettor needs to pay to bet on the sport-event
	BetFee types.Coin `protobuf:"bytes,3,opt,name=bet_fee,json=betFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"bet_fee"`
}

func (m *EventBetConstraints) Reset()         { *m = EventBetConstraints{} }
func (m *EventBetConstraints) String() string { return proto.CompactTextString(m) }
func (*EventBetConstraints) ProtoMessage()    {}
func (*EventBetConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{2}
}
func (m *EventBetConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBetConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBetConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBetConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBetConstraints.Merge(m, src)
}
func (m *EventBetConstraints) XXX_Size() int {
	return m.Size()
}
func (m *EventBetConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBetConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_EventBetConstraints proto.InternalMessageInfo

func (m *EventBetConstraints) GetBetFee() types.Coin {
	if m != nil {
		return m.BetFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("sgenetwork.sge.sportevent.SportEventStatus", SportEventStatus_name, SportEventStatus_value)
	proto.RegisterType((*SportEvent)(nil), "sgenetwork.sge.sportevent.SportEvent")
	proto.RegisterType((*ResolutionEvent)(nil), "sgenetwork.sge.sportevent.ResolutionEvent")
	proto.RegisterType((*EventBetConstraints)(nil), "sgenetwork.sge.sportevent.EventBetConstraints")
}

func init() { proto.RegisterFile("sge/sportevent/sport_event.proto", fileDescriptor_f4c38f73099259f8) }

var fileDescriptor_f4c38f73099259f8 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x8d, 0x93, 0x90, 0x90, 0xc9, 0xf7, 0x85, 0x68, 0xa8, 0x5a, 0x83, 0x2a, 0xdb, 0xcd, 0x82,
	0x46, 0x45, 0xd8, 0x22, 0x3c, 0x41, 0x1c, 0x1b, 0x14, 0x29, 0x0d, 0x68, 0x9c, 0x80, 0xd4, 0x8d,
	0xe5, 0xc4, 0xd3, 0xd4, 0xa2, 0xf1, 0x20, 0xcf, 0x04, 0xda, 0xb7, 0xe8, 0xa6, 0x2f, 0xd1, 0xa7,
	0xe8, 0xaa, 0xa2, 0x3b, 0x96, 0x55, 0x17, 0x6e, 0x65, 0x76, 0x3c, 0x45, 0x35, 0x13, 0x97, 0xf0,
	0x13, 0x10, 0x42, 0xdd, 0xd8, 0xf7, 0xce, 0x39, 0xc7, 0x73, 0xe7, 0x5c, 0xcf, 0x05, 0x1a, 0x1d,
	0x61, 0x83, 0x1e, 0x91, 0x88, 0xe1, 0x63, 0x1c, 0xb2, 0x69, 0xe8, 0x8a, 0x58, 0x3f, 0x8a, 0x08,
	0x23, 0x70, 0x85, 0x8e, 0x70, 0x88, 0xd9, 0x09, 0x89, 0x0e, 0x75, 0x3a, 0xc2, 0xfa, 0x8c, 0xbc,
	0xfa, 0x64, 0x44, 0x46, 0x44, 0xb0, 0x0c, 0x1e, 0x4d, 0x05, 0xab, 0xca, 0x90, 0xd0, 0x31, 0xa1,
	0xc6, 0xc0, 0xa3, 0xd8, 0x38, 0xde, 0x1c, 0x60, 0xe6, 0x6d, 0x1a, 0x43, 0x12, 0x84, 0x29, 0xbe,
	0x72, 0x63, 0x4b, 0xe2, 0xfb, 0x74, 0x0a, 0xd5, 0xbe, 0xe5, 0x01, 0x70, 0x38, 0x62, 0x73, 0x04,
	0x6a, 0x20, 0x37, 0x09, 0x7c, 0x59, 0xd2, 0xa4, 0x7a, 0xc9, 0xac, 0x24, 0xb1, 0x9a, 0xeb, 0xb7,
	0xad, 0x8b, 0x58, 0xe5, 0xab, 0x88, 0x3f, 0xe0, 0x16, 0x58, 0xa4, 0xcc, 0x8b, 0x98, 0xcb, 0xa8,
	0x9c, 0xd5, 0xa4, 0x7a, 0xde, 0x7c, 0x96, 0xc4, 0x6a, 0xd1, 0xe1, 0x6b, 0x3d, 0xe7, 0x22, 0x56,
	0x2f, 0x61, 0x74, 0x19, 0xc1, 0x75, 0x50, 0xc0, 0xa1, 0xcf, 0x25, 0x39, 0x21, 0x59, 0x4e, 0x62,
	0x75, 0xc1, 0x0e, 0x7d, 0x21, 0x48, 0x21, 0x94, 0xbe, 0xe1, 0x16, 0xc8, 0xf3, 0x02, 0xe5, 0xbc,
	0x96, 0xab, 0x97, 0x1b, 0xaa, 0x7e, 0xa7, 0x1b, 0xfa, 0xae, 0xef, 0x53, 0x24, 0xc8, 0x10, 0x81,
	0xea, 0x49, 0x10, 0x86, 0x38, 0x72, 0x79, 0xea, 0x4e, 0x02, 0x9f, 0xca, 0x0b, 0x5a, 0xae, 0x5e,
	0x32, 0xd7, 0x92, 0x58, 0xad, 0x1c, 0x08, 0x8c, 0xf3, 0xfb, 0x6d, 0x8b, 0x5e, 0xc4, 0xea, 0x2d,
	0x36, 0xba, 0xb5, 0x02, 0x5b, 0xa0, 0x40, 0x99, 0xc7, 0x26, 0x54, 0x2e, 0x68, 0x52, 0xbd, 0xd2,
	0x58, 0xbf, 0xa7, 0x94, 0x99, 0x87, 0x8e, 0x90, 0xa0, 0x54, 0x0a, 0x77, 0xc0, 0xff, 0x11, 0xa6,
	0xe4, 0xfd, 0x84, 0x05, 0x24, 0xe4, 0x0e, 0x14, 0x85, 0x03, 0x2f, 0x92, 0x58, 0xfd, 0x0f, 0x5d,
	0x02, 0xc2, 0x88, 0xeb, 0x44, 0x74, 0x3d, 0x85, 0x32, 0x28, 0x0e, 0x23, 0xec, 0x31, 0x12, 0xc9,
	0x8b, 0xbc, 0x3d, 0xe8, 0x6f, 0x0a, 0x0f, 0xc0, 0xd2, 0x00, 0x33, 0x77, 0x48, 0x42, 0xca, 0x22,
	0x2f, 0x08, 0x19, 0x95, 0x4b, 0x9a, 0x54, 0x2f, 0x37, 0xf4, 0x7b, 0x0a, 0x16, 0xb5, 0x9a, 0x98,
	0xb5, 0x66, 0x2a, 0x54, 0x19, 0x5c, 0xcb, 0xe1, 0x53, 0x50, 0xf0, 0x86, 0x2c, 0x38, 0xc6, 0x32,
	0xd0, 0xa4, 0xfa, 0x22, 0x4a, 0x33, 0x08, 0x41, 0x7e, 0x8c, 0x99, 0x27, 0x97, 0x45, 0x1d, 0x22,
	0xae, 0x7d, 0xce, 0x82, 0xa5, 0xd9, 0x79, 0x1e, 0xfa, 0x37, 0xdd, 0x72, 0x27, 0xfb, 0x48, 0x77,
	0xe6, 0xf5, 0x3f, 0xf7, 0xcf, 0xfa, 0x9f, 0x7f, 0x74, 0xff, 0x6b, 0xdf, 0x25, 0xb0, 0x3c, 0xc7,
	0x6b, 0xf8, 0x1a, 0x80, 0x71, 0x10, 0xba, 0xde, 0x98, 0x4c, 0x42, 0x26, 0x8e, 0x5d, 0x32, 0xf5,
	0xd3, 0x58, 0xcd, 0xfc, 0x8c, 0xd5, 0xb5, 0x51, 0xc0, 0xde, 0x4d, 0x06, 0xfa, 0x90, 0x8c, 0x8d,
	0xf4, 0x6a, 0x4f, 0x5f, 0x1b, 0xd4, 0x3f, 0x34, 0xd8, 0xc7, 0x23, 0x4c, 0xf5, 0x76, 0xc8, 0x50,
	0x69, 0x1c, 0x84, 0x4d, 0xf1, 0x01, 0x38, 0x04, 0x45, 0xfe, 0x0f, 0xbc, 0xc5, 0x58, 0x5c, 0xb1,
	0x72, 0x63, 0x45, 0x9f, 0x4a, 0x74, 0x3e, 0x14, 0xf4, 0x74, 0x28, 0xe8, 0x2d, 0x12, 0x84, 0xa6,
	0xc1, 0xb7, 0xf9, 0xf2, 0x4b, 0x7d, 0xf9, 0x80, 0x6d, 0xb8, 0x00, 0x15, 0x06, 0x98, 0x6d, 0x63,
	0xfc, 0xea, 0xab, 0x04, 0xaa, 0x37, 0x0f, 0x0a, 0x6b, 0x40, 0x71, 0xf6, 0x76, 0x51, 0xcf, 0xb5,
	0xf7, 0xed, 0x6e, 0xcf, 0x75, 0x7a, 0xcd, 0x5e, 0xdf, 0x71, 0xfb, 0x5d, 0x67, 0xcf, 0x6e, 0xb5,
	0xb7, 0xdb, 0xb6, 0x55, 0xcd, 0x40, 0x05, 0xac, 0xce, 0xe1, 0xb4, 0xbb, 0xfb, 0xcd, 0x4e, 0xdb,
	0xaa, 0x4a, 0x50, 0x03, 0xcf, 0xe7, 0xe0, 0xad, 0x66, 0xb7, 0x65, 0x77, 0x3a, 0xb6, 0x55, 0xcd,
	0xde, 0xf1, 0x85, 0xa6, 0xb9, 0x8b, 0x7a, 0xb6, 0x55, 0xcd, 0xc1, 0x35, 0x50, 0x9b, 0x83, 0x23,
	0xdb, 0xe9, 0x77, 0x7a, 0xae, 0x65, 0xb7, 0x3a, 0x4d, 0x64, 0x5b, 0xd5, 0xbc, 0xb9, 0x73, 0x9a,
	0x28, 0xd2, 0x59, 0xa2, 0x48, 0xbf, 0x13, 0x45, 0xfa, 0x74, 0xae, 0x64, 0xce, 0xce, 0x95, 0xcc,
	0x8f, 0x73, 0x25, 0xf3, 0x66, 0xe3, 0x8a, 0x1b, 0x74, 0x84, 0x37, 0xd2, 0x46, 0xf3, 0xd8, 0xf8,
	0x70, 0x75, 0x7a, 0x0a, 0x63, 0x06, 0x05, 0x31, 0x3f, 0xb7, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x75, 0xb5, 0x4e, 0xd0, 0xcf, 0x05, 0x00, 0x00,
}

func (m *SportEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.BetConstraints != nil {
		{
			size, err := m.BetConstraints.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSportEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintSportEvent(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSportEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResolutionEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolutionEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResolutionEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintSportEvent(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBetConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBetConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBetConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BetFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSportEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinAmount.Size()
		i -= size
		if _, err := m.MinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSportEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintSportEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovSportEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SportEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovSportEvent(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovSportEvent(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovSportEvent(uint64(l))
		}
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovSportEvent(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovSportEvent(uint64(m.Status))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovSportEvent(uint64(m.ResolutionTS))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.BetConstraints != nil {
		l = m.BetConstraints.Size()
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.Active {
		n += 2
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	return n
}

func (m *ResolutionEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovSportEvent(uint64(m.ResolutionTS))
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovSportEvent(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovSportEvent(uint64(m.Status))
	}
	return n
}

func (m *EventBetConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinAmount.Size()
	n += 1 + l + sovSportEvent(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovSportEvent(uint64(l))
	return n
}

func sovSportEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSportEvent(x uint64) (n int) {
	return sovSportEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SportEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSportEvent
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
			return fmt.Errorf("proto: SportEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
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
					return ErrIntOverflowSportEvent
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
					return ErrIntOverflowSportEvent
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
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetConstraints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BetConstraints == nil {
				m.BetConstraints = &EventBetConstraints{}
			}
			if err := m.BetConstraints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSportEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSportEvent
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
func (m *ResolutionEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSportEvent
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
			return fmt.Errorf("proto: ResolutionEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolutionEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
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
					return ErrIntOverflowSportEvent
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
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
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
					return ErrIntOverflowSportEvent
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
			skippy, err := skipSportEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSportEvent
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
func (m *EventBetConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSportEvent
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
			return fmt.Errorf("proto: EventBetConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBetConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
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
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSportEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSportEvent
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
func skipSportEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSportEvent
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
					return 0, ErrIntOverflowSportEvent
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
					return 0, ErrIntOverflowSportEvent
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
				return 0, ErrInvalidLengthSportEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSportEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSportEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSportEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSportEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSportEvent = fmt.Errorf("proto: unexpected end of group")
)
