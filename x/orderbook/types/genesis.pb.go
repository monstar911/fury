// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/orderbook/genesis.proto

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

// GenesisState defines the orderbook module's genesis state.
type GenesisState struct {
	// params defines all the parameters related to orderbook.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// books defines the books active at genesis.
	Books []OrderBook `protobuf:"bytes,2,rep,name=books,proto3" json:"books"`
	// book_participants defines the book participants active at genesis.
	BookParticipants []BookParticipant `protobuf:"bytes,3,rep,name=book_participants,json=bookParticipants,proto3" json:"book_participants"`
	// book_exposures defines the book exposures active at genesis.
	BookExposures []BookOddsExposure `protobuf:"bytes,4,rep,name=book_exposures,json=bookExposures,proto3" json:"book_exposures"`
	// participant_exposures defines the participant exposures active at genesis.
	ParticipantExposures []ParticipantExposure `protobuf:"bytes,5,rep,name=participant_exposures,json=participantExposures,proto3" json:"participant_exposures"`
	// stats is the statistics of the order-book
	Stats OrderBookStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54e9379cfb7d94d, []int{0}
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

func (m *GenesisState) GetBooks() []OrderBook {
	if m != nil {
		return m.Books
	}
	return nil
}

func (m *GenesisState) GetBookParticipants() []BookParticipant {
	if m != nil {
		return m.BookParticipants
	}
	return nil
}

func (m *GenesisState) GetBookExposures() []BookOddsExposure {
	if m != nil {
		return m.BookExposures
	}
	return nil
}

func (m *GenesisState) GetParticipantExposures() []ParticipantExposure {
	if m != nil {
		return m.ParticipantExposures
	}
	return nil
}

func (m *GenesisState) GetStats() OrderBookStats {
	if m != nil {
		return m.Stats
	}
	return OrderBookStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.orderbook.GenesisState")
}

func init() { proto.RegisterFile("sge/orderbook/genesis.proto", fileDescriptor_b54e9379cfb7d94d) }

var fileDescriptor_b54e9379cfb7d94d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x77, 0xf3, 0xcf, 0x61, 0xac, 0xa8, 0xc1, 0x60, 0x31, 0xda, 0xa4, 0x2e, 0x16, 0xb9,
	0x0b, 0x76, 0x2f, 0x90, 0xfe, 0x1c, 0x95, 0x3a, 0x04, 0x11, 0xc4, 0xae, 0x0e, 0xe3, 0x22, 0xfa,
	0x0e, 0xf3, 0x8e, 0x64, 0xdf, 0xa2, 0x8f, 0xe5, 0xd1, 0x63, 0xa7, 0x08, 0xfd, 0x0e, 0x9d, 0x63,
	0x66, 0x47, 0xd3, 0x42, 0xe9, 0xf6, 0xec, 0x3c, 0xcf, 0xf3, 0x7b, 0x97, 0x79, 0x87, 0xec, 0x23,
	0x67, 0x21, 0xc8, 0x36, 0x93, 0x31, 0x40, 0x37, 0xe4, 0xac, 0xcf, 0x30, 0xc1, 0x40, 0x48, 0x50,
	0x40, 0x3d, 0xd4, 0xdf, 0xea, 0x05, 0x64, 0x37, 0x40, 0xce, 0x82, 0x79, 0xae, 0x54, 0xe4, 0xc0,
	0xc1, 0x84, 0x42, 0xad, 0xd2, 0x7c, 0xe9, 0x60, 0x19, 0x36, 0x57, 0xa9, 0x7d, 0xf4, 0x95, 0x21,
	0x9b, 0xb7, 0xe9, 0x80, 0x7b, 0x15, 0x29, 0x46, 0x2f, 0x48, 0x5e, 0x44, 0x32, 0xea, 0xa1, 0xe7,
	0x96, 0xdd, 0x4a, 0xa1, 0x56, 0x0e, 0x56, 0x0d, 0x0c, 0x9a, 0x26, 0x57, 0xcf, 0x8e, 0x3e, 0x0e,
	0x9d, 0x3b, 0xdb, 0xa2, 0x97, 0x24, 0xa7, 0x4d, 0xf4, 0x36, 0xca, 0x99, 0x4a, 0xa1, 0x76, 0xbc,
	0xba, 0xde, 0xd0, 0xaa, 0x0e, 0xd0, 0xb5, 0x84, 0xb4, 0x47, 0x9f, 0xc8, 0xae, 0x16, 0xcf, 0x22,
	0x92, 0x2a, 0x69, 0x25, 0x22, 0xea, 0x2b, 0xf4, 0x32, 0x06, 0x76, 0xb2, 0x1a, 0xa6, 0x39, 0xcd,
	0x9f, 0x86, 0x45, 0xee, 0xc4, 0xcb, 0xc7, 0x48, 0x1f, 0xc8, 0xb6, 0xa1, 0xb3, 0xa1, 0x00, 0x1c,
	0x48, 0x86, 0x5e, 0xd6, 0xa0, 0x4f, 0xd7, 0xa3, 0x1b, 0xed, 0x36, 0x5e, 0xdb, 0x8a, 0x65, 0x6f,
	0x69, 0x73, 0x76, 0x86, 0xb4, 0x43, 0xf6, 0x16, 0xfe, 0x78, 0x81, 0x9f, 0x33, 0xfc, 0xea, 0xda,
	0x6b, 0x9c, 0xd5, 0x7e, 0x8d, 0x28, 0x8a, 0xbf, 0x16, 0xd2, 0x2b, 0x92, 0x43, 0x15, 0x29, 0xf4,
	0xf2, 0x66, 0x41, 0x95, 0x7f, 0xdc, 0xb0, 0x5e, 0xed, 0x6c, 0x51, 0x69, 0xb9, 0x7e, 0x33, 0x9a,
	0xf8, 0xee, 0x78, 0xe2, 0xbb, 0x9f, 0x13, 0xdf, 0x7d, 0x9b, 0xfa, 0xce, 0x78, 0xea, 0x3b, 0xef,
	0x53, 0xdf, 0x79, 0x3c, 0xe3, 0x89, 0xea, 0x0c, 0xe2, 0xa0, 0x05, 0xbd, 0x10, 0x39, 0xab, 0x5a,
	0xb6, 0xd6, 0xe1, 0x70, 0xe1, 0x29, 0xa9, 0x57, 0xc1, 0x30, 0xce, 0x9b, 0x77, 0x74, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x4b, 0x0b, 0xa4, 0xb5, 0x02, 0x00, 0x00,
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
	dAtA[i] = 0x32
	if len(m.ParticipantExposures) > 0 {
		for iNdEx := len(m.ParticipantExposures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipantExposures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BookExposures) > 0 {
		for iNdEx := len(m.BookExposures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BookExposures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BookParticipants) > 0 {
		for iNdEx := len(m.BookParticipants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BookParticipants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Books) > 0 {
		for iNdEx := len(m.Books) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Books[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Books) > 0 {
		for _, e := range m.Books {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BookParticipants) > 0 {
		for _, e := range m.BookParticipants {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BookExposures) > 0 {
		for _, e := range m.BookExposures {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipantExposures) > 0 {
		for _, e := range m.ParticipantExposures {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Books", wireType)
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
			m.Books = append(m.Books, OrderBook{})
			if err := m.Books[len(m.Books)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookParticipants", wireType)
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
			m.BookParticipants = append(m.BookParticipants, BookParticipant{})
			if err := m.BookParticipants[len(m.BookParticipants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookExposures", wireType)
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
			m.BookExposures = append(m.BookExposures, BookOddsExposure{})
			if err := m.BookExposures[len(m.BookExposures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantExposures", wireType)
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
			m.ParticipantExposures = append(m.ParticipantExposures, ParticipantExposure{})
			if err := m.ParticipantExposures[len(m.ParticipantExposures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
