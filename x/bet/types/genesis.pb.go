// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the bet module's genesis state.
type GenesisState struct {
	// params contains parameters of dvm module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// bet_list contains bet list in the genesis init.
	BetList []Bet `protobuf:"bytes,2,rep,name=bet_list,json=betList,proto3" json:"bet_list"`
	// uid2id_list contains bet to id list in the genesis init.
	Uid2IdList []UID2ID `protobuf:"bytes,3,rep,name=uid2id_list,json=uid2idList,proto3" json:"uid2id_list"`
	// stats contains statistics in the genesis init.
	Stats BetStats `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c49ebc0f2678a09, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBetList() []Bet {
	if m != nil {
		return m.BetList
	}
	return nil
}

func (m *GenesisState) GetUid2IdList() []UID2ID {
	if m != nil {
		return m.Uid2IdList
	}
	return nil
}

func (m *GenesisState) GetStats() BetStats {
	if m != nil {
		return m.Stats
	}
	return BetStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.bet.GenesisState")
}

func init() { proto.RegisterFile("sge/bet/genesis.proto", fileDescriptor_6c49ebc0f2678a09) }

var fileDescriptor_6c49ebc0f2678a09 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0xe3, 0x7f, 0xfb, 0x2f, 0xc8, 0x65, 0x21, 0x14, 0x51, 0x45, 0xc8, 0x54, 0x0c, 0xa8,
	0x0b, 0x8e, 0x14, 0x96, 0x8c, 0x10, 0x55, 0x42, 0x95, 0x18, 0x10, 0x88, 0x85, 0x05, 0xc5, 0xf4,
	0x64, 0x2c, 0x28, 0x8e, 0xe2, 0xab, 0x80, 0xb7, 0xe0, 0xb1, 0x3a, 0x76, 0x64, 0x42, 0x28, 0x79,
	0x10, 0x50, 0x6c, 0x67, 0xa2, 0x6c, 0xe7, 0xf3, 0xf7, 0xfb, 0xee, 0xbb, 0xa3, 0xbb, 0x46, 0x42,
	0x2c, 0x00, 0x63, 0x09, 0xcf, 0x60, 0x94, 0xe1, 0x45, 0xa9, 0x51, 0x87, 0xa1, 0x69, 0xde, 0xf8,
	0xa2, 0xcb, 0x47, 0x6e, 0x24, 0x70, 0x01, 0x18, 0x0d, 0xa4, 0x96, 0xda, 0x7e, 0xc7, 0x4d, 0xe5,
	0x94, 0xd1, 0xa0, 0x35, 0x28, 0xf2, 0x32, 0x9f, 0x7b, 0x3e, 0xda, 0x6e, 0xbb, 0x02, 0xd0, 0xb7,
	0x76, 0xda, 0x96, 0xc1, 0x1c, 0xbd, 0xee, 0xf0, 0x9b, 0xd0, 0xad, 0x73, 0x37, 0xf9, 0x1a, 0x73,
	0x84, 0x30, 0xa5, 0x3d, 0x67, 0x34, 0x24, 0x23, 0x32, 0xee, 0x27, 0x11, 0xff, 0x9d, 0x84, 0x5f,
	0x5a, 0x45, 0xd6, 0x5d, 0x7e, 0x1e, 0x04, 0x57, 0x5e, 0x1f, 0xa6, 0x74, 0x53, 0x00, 0xde, 0x3d,
	0x29, 0x83, 0xc3, 0x7f, 0xa3, 0xce, 0xb8, 0x9f, 0xec, 0xad, 0x63, 0x33, 0x40, 0x0f, 0x6e, 0x08,
	0xc0, 0x0b, 0x65, 0x30, 0x3c, 0xa3, 0xfd, 0x85, 0x9a, 0x25, 0x6a, 0xe6, 0xe0, 0x8e, 0x85, 0xd7,
	0x0e, 0xbe, 0x99, 0x4e, 0x92, 0xe9, 0xc4, 0xf3, 0xd4, 0x41, 0xd6, 0x22, 0xa5, 0xff, 0xed, 0x5a,
	0xc3, 0xae, 0x4d, 0xbd, 0xff, 0xc7, 0xe4, 0x66, 0xc7, 0x36, 0xb7, 0x03, 0xb2, 0xd3, 0x65, 0xc5,
	0xc8, 0xaa, 0x62, 0xe4, 0xab, 0x62, 0xe4, 0xbd, 0x66, 0xc1, 0xaa, 0x66, 0xc1, 0x47, 0xcd, 0x82,
	0xdb, 0x23, 0xa9, 0xf0, 0x61, 0x21, 0xf8, 0xbd, 0x9e, 0xc7, 0x46, 0xc2, 0xb1, 0xf7, 0x6b, 0xea,
	0xf8, 0xd5, 0x5e, 0x12, 0xdf, 0x0a, 0x30, 0xa2, 0x67, 0x4f, 0x79, 0xf2, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x0e, 0xd6, 0xf8, 0xcb, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Uid2IdList) > 0 {
		for iNdEx := len(m.Uid2IdList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Uid2IdList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BetList) > 0 {
		for iNdEx := len(m.BetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BetList) > 0 {
		for _, e := range m.BetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Uid2IdList) > 0 {
		for _, e := range m.Uid2IdList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Stats.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BetList = append(m.BetList, Bet{})
			if err := m.BetList[len(m.BetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid2IdList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid2IdList = append(m.Uid2IdList, UID2ID{})
			if err := m.Uid2IdList[len(m.Uid2IdList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
