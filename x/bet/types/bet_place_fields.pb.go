// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/bet_place_fields.proto

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

// PlaceBetFields contains necessary fields which come in BetPlacement and BetSlipPlacement TX requests
type BetPlaceFields struct {
	// uid is the unique uuid assigned to bet
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// amount is the wagger amount
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// ticket is a signed string containing important info such as `oddsValue`
	Ticket string `protobuf:"bytes,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	// odds_type is the type of odds used has chosen
	OddsType OddsType `protobuf:"varint,4,opt,name=odds_type,json=oddsType,proto3,enum=sgenetwork.sge.bet.OddsType" json:"odds_type,omitempty"`
}

func (m *BetPlaceFields) Reset()         { *m = BetPlaceFields{} }
func (m *BetPlaceFields) String() string { return proto.CompactTextString(m) }
func (*BetPlaceFields) ProtoMessage()    {}
func (*BetPlaceFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c2bcc9c0d2fd18c, []int{0}
}
func (m *BetPlaceFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BetPlaceFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetPlaceFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BetPlaceFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetPlaceFields.Merge(m, src)
}
func (m *BetPlaceFields) XXX_Size() int {
	return m.Size()
}
func (m *BetPlaceFields) XXX_DiscardUnknown() {
	xxx_messageInfo_BetPlaceFields.DiscardUnknown(m)
}

var xxx_messageInfo_BetPlaceFields proto.InternalMessageInfo

func (m *BetPlaceFields) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *BetPlaceFields) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *BetPlaceFields) GetOddsType() OddsType {
	if m != nil {
		return m.OddsType
	}
	return OddsType_ODD_TYPE_INVALID
}

func init() {
	proto.RegisterType((*BetPlaceFields)(nil), "BetPlaceFields")
}

func init() { proto.RegisterFile("sge/bet/bet_place_fields.proto", fileDescriptor_7c2bcc9c0d2fd18c) }

var fileDescriptor_7c2bcc9c0d2fd18c = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x4f, 0xd5,
	0x4f, 0x4a, 0x2d, 0x01, 0xe1, 0xf8, 0x82, 0x9c, 0xc4, 0xe4, 0xd4, 0xf8, 0xb4, 0xcc, 0xd4, 0x9c,
	0x94, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53,
	0x1f, 0xc4, 0x82, 0x8a, 0x8a, 0xc3, 0x74, 0xe5, 0xa7, 0xa4, 0x14, 0xc7, 0x97, 0x54, 0x16, 0xa4,
	0x42, 0x24, 0x94, 0xce, 0x32, 0x72, 0xf1, 0x39, 0xa5, 0x96, 0x04, 0x80, 0x0c, 0x72, 0x03, 0x9b,
	0x23, 0xa4, 0xc0, 0xc5, 0x5c, 0x9a, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xf7,
	0xe8, 0x9e, 0x3c, 0x73, 0xa8, 0xa7, 0xcb, 0xab, 0x7b, 0xf2, 0x20, 0xd1, 0x20, 0x10, 0x21, 0xe4,
	0xc6, 0xc5, 0x96, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0x04, 0x56, 0xa4, 0x77, 0xe2, 0x9e,
	0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0x1f, 0x64, 0x6b,
	0xb1, 0x9e, 0x67, 0x5e, 0x49, 0x10, 0x54, 0xb7, 0x90, 0x18, 0x17, 0x5b, 0x49, 0x66, 0x72, 0x76,
	0x6a, 0x89, 0x04, 0x33, 0xc8, 0x9c, 0x20, 0x28, 0x4f, 0xc8, 0x92, 0x8b, 0x13, 0xee, 0x4e, 0x09,
	0x16, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x19, 0xbd, 0xe2, 0xf4, 0xd4, 0xbc, 0xd4, 0x92, 0xf2, 0xfc,
	0xa2, 0x6c, 0x10, 0x53, 0x2f, 0x29, 0xb5, 0x44, 0xcf, 0x3f, 0x25, 0xa5, 0x38, 0xa4, 0xb2, 0x20,
	0x35, 0x88, 0x23, 0x1f, 0xca, 0x72, 0x72, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x64, 0xc7, 0x15, 0xa7, 0xa7, 0xea, 0x42, 0x0d, 0x03, 0xb1, 0xf5, 0x2b, 0xc0, 0x61,
	0x03, 0x76, 0x60, 0x12, 0x1b, 0x38, 0x60, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xb0,
	0x63, 0x63, 0x69, 0x01, 0x00, 0x00,
}

func (m *BetPlaceFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetPlaceFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetPlaceFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OddsType != 0 {
		i = encodeVarintBetPlaceFields(dAtA, i, uint64(m.OddsType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintBetPlaceFields(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBetPlaceFields(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBetPlaceFields(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBetPlaceFields(dAtA []byte, offset int, v uint64) int {
	offset -= sovBetPlaceFields(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BetPlaceFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBetPlaceFields(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovBetPlaceFields(uint64(l))
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovBetPlaceFields(uint64(l))
	}
	if m.OddsType != 0 {
		n += 1 + sovBetPlaceFields(uint64(m.OddsType))
	}
	return n
}

func sovBetPlaceFields(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBetPlaceFields(x uint64) (n int) {
	return sovBetPlaceFields(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BetPlaceFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBetPlaceFields
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
			return fmt.Errorf("proto: BetPlaceFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetPlaceFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetPlaceFields
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
				return ErrInvalidLengthBetPlaceFields
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetPlaceFields
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetPlaceFields
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
				return ErrInvalidLengthBetPlaceFields
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetPlaceFields
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetPlaceFields
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
				return ErrInvalidLengthBetPlaceFields
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetPlaceFields
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsType", wireType)
			}
			m.OddsType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetPlaceFields
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
			skippy, err := skipBetPlaceFields(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBetPlaceFields
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
func skipBetPlaceFields(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBetPlaceFields
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
					return 0, ErrIntOverflowBetPlaceFields
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
					return 0, ErrIntOverflowBetPlaceFields
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
				return 0, ErrInvalidLengthBetPlaceFields
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBetPlaceFields
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBetPlaceFields
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBetPlaceFields        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBetPlaceFields          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBetPlaceFields = fmt.Errorf("proto: unexpected end of group")
)
