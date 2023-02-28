// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/bet_odds.proto

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

// BetOdds is the type to store odds of an sport-event.
type BetOdds struct {
	// uid is universal unique identifier of odds.
	// Required | Unique | uuid-v4 or code
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// sport_event_uid is the parent, used for fast retrieving.
	// Required | NonUnique | -
	SportEventUID string `protobuf:"bytes,2,opt,name=sport_event_uid,proto3" json:"sport_event_uid"`
	// value of the odds in corresponding odds type proposed in bet placement
	// message. Required | NonUnique | "1.286" or "2/7" or "+500"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// max_loss_multiplier is the factor for calculating max loss for given odds
	MaxLossMultiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_loss_multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_loss_multiplier"`
}

func (m *BetOdds) Reset()         { *m = BetOdds{} }
func (m *BetOdds) String() string { return proto.CompactTextString(m) }
func (*BetOdds) ProtoMessage()    {}
func (*BetOdds) Descriptor() ([]byte, []int) {
	return fileDescriptor_2629a03d0a23fb04, []int{0}
}
func (m *BetOdds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BetOdds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetOdds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BetOdds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetOdds.Merge(m, src)
}
func (m *BetOdds) XXX_Size() int {
	return m.Size()
}
func (m *BetOdds) XXX_DiscardUnknown() {
	xxx_messageInfo_BetOdds.DiscardUnknown(m)
}

var xxx_messageInfo_BetOdds proto.InternalMessageInfo

func (m *BetOdds) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *BetOdds) GetSportEventUID() string {
	if m != nil {
		return m.SportEventUID
	}
	return ""
}

func (m *BetOdds) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*BetOdds)(nil), "sgenetwork.sge.bet.BetOdds")
}

func init() { proto.RegisterFile("sge/bet/bet_odds.proto", fileDescriptor_2629a03d0a23fb04) }

var fileDescriptor_2629a03d0a23fb04 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x4f, 0xd5,
	0x4f, 0x4a, 0x2d, 0x01, 0xe1, 0xf8, 0xfc, 0x94, 0x94, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x21, 0xa1, 0xe2, 0xf4, 0xd4, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0xbd, 0xe2, 0xf4, 0x54,
	0xbd, 0xa4, 0xd4, 0x12, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb4, 0x3e, 0x88, 0x05, 0x51,
	0xa9, 0xb4, 0x80, 0x89, 0x8b, 0xdd, 0x29, 0xb5, 0xc4, 0x3f, 0x25, 0xa5, 0x58, 0x48, 0x81, 0x8b,
	0xb9, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xef, 0xd1, 0x3d, 0x79, 0xe6,
	0x50, 0x4f, 0x97, 0x57, 0xf7, 0xe4, 0x41, 0xa2, 0x41, 0x20, 0x42, 0xc8, 0x8f, 0x8b, 0xbf, 0xb8,
	0x20, 0xbf, 0xa8, 0x24, 0x3e, 0xb5, 0x2c, 0x35, 0xaf, 0x24, 0x1e, 0xa4, 0x9a, 0x09, 0xac, 0x5a,
	0xe5, 0xd1, 0x3d, 0x79, 0xde, 0x60, 0x90, 0x94, 0x2b, 0x48, 0x06, 0xa2, 0x0f, 0x5d, 0x6d, 0x10,
	0xba, 0x80, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x33, 0xc8, 0x94, 0x20,
	0x08, 0x47, 0xa8, 0x87, 0x91, 0x4b, 0x38, 0x37, 0xb1, 0x22, 0x3e, 0x27, 0xbf, 0xb8, 0x38, 0x3e,
	0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x20, 0x27, 0x33, 0xb5, 0x48, 0x82, 0x05, 0x6c, 0x55, 0xd4, 0x89,
	0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92,
	0xca, 0x82, 0xd4, 0x62, 0x3d, 0x97, 0xd4, 0xe4, 0x47, 0xf7, 0xe4, 0x05, 0x7d, 0x13, 0x2b, 0x7c,
	0xf2, 0x8b, 0x8b, 0x7d, 0xe1, 0x46, 0xbd, 0xba, 0x27, 0x8f, 0xcd, 0x86, 0x20, 0x6c, 0x82, 0x4e,
	0x0e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x85, 0xec, 0x84, 0xe2, 0xf4,
	0x54, 0x5d, 0x68, 0x90, 0x83, 0xd8, 0xfa, 0x15, 0xe0, 0x78, 0x01, 0x3b, 0x23, 0x89, 0x0d, 0x1c,
	0xd6, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5f, 0xde, 0x75, 0xaf, 0x01, 0x00, 0x00,
}

func (m *BetOdds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetOdds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetOdds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLossMultiplier.Size()
		i -= size
		if _, err := m.MaxLossMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBetOdds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBetOdds(dAtA []byte, offset int, v uint64) int {
	offset -= sovBetOdds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BetOdds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = m.MaxLossMultiplier.Size()
	n += 1 + l + sovBetOdds(uint64(l))
	return n
}

func sovBetOdds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBetOdds(x uint64) (n int) {
	return sovBetOdds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BetOdds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBetOdds
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
			return fmt.Errorf("proto: BetOdds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetOdds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
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
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLossMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLossMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBetOdds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBetOdds
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
func skipBetOdds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
				return 0, ErrInvalidLengthBetOdds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBetOdds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBetOdds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBetOdds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBetOdds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBetOdds = fmt.Errorf("proto: unexpected end of group")
)
