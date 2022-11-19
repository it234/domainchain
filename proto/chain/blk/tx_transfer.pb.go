// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx_transfer.proto

package blk

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TxTransfer struct {
	ToAddr               []byte   `protobuf:"bytes,1,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Sequence             uint64   `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxTransfer) Reset()         { *m = TxTransfer{} }
func (m *TxTransfer) String() string { return proto.CompactTextString(m) }
func (*TxTransfer) ProtoMessage()    {}
func (*TxTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91df75eb684d408, []int{0}
}
func (m *TxTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxTransfer.Merge(m, src)
}
func (m *TxTransfer) XXX_Size() int {
	return m.Size()
}
func (m *TxTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_TxTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_TxTransfer proto.InternalMessageInfo

func (m *TxTransfer) GetToAddr() []byte {
	if m != nil {
		return m.ToAddr
	}
	return nil
}

func (m *TxTransfer) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxTransfer) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*TxTransfer)(nil), "blk.TxTransfer")
}

func init() { proto.RegisterFile("tx_transfer.proto", fileDescriptor_a91df75eb684d408) }

var fileDescriptor_a91df75eb684d408 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x88, 0x2f,
	0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0xca, 0xc9, 0x56, 0x8a, 0xe4, 0xe2, 0x0a, 0xa9, 0x08, 0x81, 0x4a, 0x08, 0x89, 0x73, 0xb1, 0x97,
	0xe4, 0xc7, 0x27, 0xa6, 0xa4, 0x14, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0xb1, 0x95, 0xe4,
	0x3b, 0xa6, 0xa4, 0x14, 0x09, 0x89, 0x71, 0xb1, 0x25, 0xe6, 0xe6, 0x97, 0xe6, 0x95, 0x48, 0x30,
	0x29, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x52, 0x5c, 0x1c, 0xc5, 0xa9, 0x85, 0xa5, 0xa9,
	0x79, 0xc9, 0xa9, 0x12, 0xcc, 0x60, 0x19, 0x38, 0xdf, 0x49, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x0a, 0x64, 0x63, 0x12,
	0x1b, 0xd8, 0x76, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x9d, 0xdd, 0xd3, 0x92, 0x00,
	0x00, 0x00,
}

func (m *TxTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Sequence != 0 {
		i = encodeVarintTxTransfer(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if m.Amount != 0 {
		i = encodeVarintTxTransfer(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ToAddr) > 0 {
		i -= len(m.ToAddr)
		copy(dAtA[i:], m.ToAddr)
		i = encodeVarintTxTransfer(dAtA, i, uint64(len(m.ToAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ToAddr)
	if l > 0 {
		n += 1 + l + sovTxTransfer(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTxTransfer(uint64(m.Amount))
	}
	if m.Sequence != 0 {
		n += 1 + sovTxTransfer(uint64(m.Sequence))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTxTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxTransfer(x uint64) (n int) {
	return sovTxTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxTransfer
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
			return fmt.Errorf("proto: TxTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddr = append(m.ToAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ToAddr == nil {
				m.ToAddr = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTxTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTxTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxTransfer
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
					return 0, ErrIntOverflowTxTransfer
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
					return 0, ErrIntOverflowTxTransfer
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
				return 0, ErrInvalidLengthTxTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxTransfer = fmt.Errorf("proto: unexpected end of group")
)
