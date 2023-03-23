// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/strategicreserve/sr_pool.proto

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

// SRPool defines the locked amount and the unlocked amount in the
// sr pool Account.
type SRPool struct {
	// locked_amount holds the amount of locked tokens in
	// sr pool module account balance.
	LockedAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=locked_amount,json=lockedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"locked_amount" yaml:"locked_amount"`
	// unlocked_amount holds the amount of unlocked tokens in
	// sr pool module account balance.
	UnlockedAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=unlocked_amount,json=unlockedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"unlocked_amount" yaml:"unlocked_amount"`
}

func (m *SRPool) Reset()      { *m = SRPool{} }
func (*SRPool) ProtoMessage() {}
func (*SRPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_c939b57bb279fe3c, []int{0}
}
func (m *SRPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SRPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SRPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SRPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRPool.Merge(m, src)
}
func (m *SRPool) XXX_Size() int {
	return m.Size()
}
func (m *SRPool) XXX_DiscardUnknown() {
	xxx_messageInfo_SRPool.DiscardUnknown(m)
}

var xxx_messageInfo_SRPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SRPool)(nil), "furysports.fury.strategicreserve.SRPool")
}

func init() {
	proto.RegisterFile("fury/strategicreserve/sr_pool.proto", fileDescriptor_c939b57bb279fe3c)
}

var fileDescriptor_c939b57bb279fe3c = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4e, 0x4f, 0xd5,
	0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xcf, 0x4c, 0x2e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b,
	0xd5, 0x2f, 0x2e, 0x8a, 0x2f, 0xc8, 0xcf, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x2f, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5, 0x43,
	0x57, 0x2e, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xab, 0x0f, 0x62, 0x41, 0xb4, 0x29, 0x7d,
	0x62, 0xe4, 0x62, 0x0b, 0x0e, 0x0a, 0xc8, 0xcf, 0xcf, 0x11, 0xca, 0xe6, 0xe2, 0xcd, 0xc9, 0x4f,
	0xce, 0x4e, 0x4d, 0x89, 0x4f, 0xcc, 0xcd, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x74, 0x72, 0x3b, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0x28, 0xa5, 0x5b, 0x9c,
	0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0xe7, 0x99, 0x57, 0xf2, 0xe9, 0x9e, 0xbc, 0x48,
	0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x8a, 0x61, 0x4a, 0x41, 0x3c, 0x10, 0xbe, 0x23, 0x98, 0x2b,
	0x54, 0xc8, 0xc5, 0x5f, 0x9a, 0x87, 0x6a, 0x1d, 0x13, 0xd8, 0x3a, 0x0f, 0x92, 0xad, 0x13, 0x83,
	0x58, 0x87, 0x66, 0x9c, 0x52, 0x10, 0x1f, 0x4c, 0x04, 0x62, 0xa5, 0x15, 0xc7, 0x8c, 0x05, 0xf2,
	0x0c, 0x2f, 0x16, 0xc8, 0x33, 0x3a, 0xf9, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x31, 0x92, 0xad, 0xc5, 0xe9, 0xa9, 0xba, 0xd0, 0x10, 0x05, 0xb1, 0xf5, 0x2b, 0x30,
	0xa3, 0x00, 0xec, 0x8c, 0x24, 0x36, 0x70, 0x50, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0x53, 0x8f, 0x12, 0xa7, 0x01, 0x00, 0x00,
}

func (this *SRPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SRPool)
	if !ok {
		that2, ok := that.(SRPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.LockedAmount.Equal(that1.LockedAmount) {
		return false
	}
	if !this.UnlockedAmount.Equal(that1.UnlockedAmount) {
		return false
	}
	return true
}
func (m *SRPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SRPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SRPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.UnlockedAmount.Size()
		i -= size
		if _, err := m.UnlockedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSrPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.LockedAmount.Size()
		i -= size
		if _, err := m.LockedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSrPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSrPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovSrPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SRPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LockedAmount.Size()
	n += 1 + l + sovSrPool(uint64(l))
	l = m.UnlockedAmount.Size()
	n += 1 + l + sovSrPool(uint64(l))
	return n
}

func sovSrPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSrPool(x uint64) (n int) {
	return sovSrPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SRPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrPool
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
			return fmt.Errorf("proto: SRPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SRPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrPool
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
				return ErrInvalidLengthSrPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSrPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrPool
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
				return ErrInvalidLengthSrPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSrPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnlockedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSrPool
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
func skipSrPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSrPool
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
					return 0, ErrIntOverflowSrPool
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
					return 0, ErrIntOverflowSrPool
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
				return 0, ErrInvalidLengthSrPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSrPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSrPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSrPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSrPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSrPool = fmt.Errorf("proto: unexpected end of group")
)
