// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/v1/packet.proto

package types

import (
	fmt "fmt"
	types "github.com/cheqd/cosmos-sdk/codec/types"
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

// Type defines a classification of message issued from a controller chain to its associated interchain accounts
// host
type Type int32

const (
	// Default zero value enumeration
	UNSPECIFIED Type = 0
	// Execute a transaction on an interchain accounts host chain
	EXECUTE_TX Type = 1
)

var Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_EXECUTE_TX",
}

var Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_EXECUTE_TX":  1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89a080d7401cd393, []int{0}
}

// InterchainAccountPacketData is comprised of a raw transaction, type of transaction and optional memo field.
type InterchainAccountPacketData struct {
	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.applications.interchain_accounts.v1.Type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *InterchainAccountPacketData) Reset()         { *m = InterchainAccountPacketData{} }
func (m *InterchainAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*InterchainAccountPacketData) ProtoMessage()    {}
func (*InterchainAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a080d7401cd393, []int{0}
}
func (m *InterchainAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainAccountPacketData.Merge(m, src)
}
func (m *InterchainAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *InterchainAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainAccountPacketData proto.InternalMessageInfo

func (m *InterchainAccountPacketData) GetType() Type {
	if m != nil {
		return m.Type
	}
	return UNSPECIFIED
}

func (m *InterchainAccountPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterchainAccountPacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// CosmosTx contains a list of sdk.Msg's. It should be used when sending transactions to an SDK host chain.
type CosmosTx struct {
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *CosmosTx) Reset()         { *m = CosmosTx{} }
func (m *CosmosTx) String() string { return proto.CompactTextString(m) }
func (*CosmosTx) ProtoMessage()    {}
func (*CosmosTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_89a080d7401cd393, []int{1}
}
func (m *CosmosTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosTx.Merge(m, src)
}
func (m *CosmosTx) XXX_Size() int {
	return m.Size()
}
func (m *CosmosTx) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosTx.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosTx proto.InternalMessageInfo

func (m *CosmosTx) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterEnum("ibc.applications.interchain_accounts.v1.Type", Type_name, Type_value)
	proto.RegisterType((*InterchainAccountPacketData)(nil), "ibc.applications.interchain_accounts.v1.InterchainAccountPacketData")
	proto.RegisterType((*CosmosTx)(nil), "ibc.applications.interchain_accounts.v1.CosmosTx")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/v1/packet.proto", fileDescriptor_89a080d7401cd393)
}

var fileDescriptor_89a080d7401cd393 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xcd, 0xb8, 0x41, 0xd6, 0x59, 0xd9, 0x2d, 0x61, 0x0f, 0x31, 0x42, 0x08, 0x2b, 0x62, 0x10,
	0x32, 0xe3, 0xb6, 0x82, 0x17, 0x2f, 0xb5, 0x8d, 0xd0, 0x8b, 0x94, 0x98, 0x42, 0xf5, 0x12, 0x26,
	0xd3, 0x31, 0x1d, 0x6c, 0x32, 0xa1, 0x33, 0x29, 0xe6, 0x1f, 0x94, 0x9e, 0xfc, 0x03, 0x3d, 0xf9,
	0x67, 0x3c, 0xf6, 0xe8, 0x51, 0xda, 0x3f, 0x22, 0x99, 0x60, 0xdb, 0x83, 0x07, 0x6f, 0x8f, 0xc7,
	0xf7, 0xde, 0xf7, 0xbd, 0xef, 0xc1, 0xd7, 0x3c, 0xa5, 0x98, 0x94, 0xe5, 0x82, 0x53, 0xa2, 0xb8,
	0x28, 0x24, 0xe6, 0x85, 0x62, 0x4b, 0x3a, 0x27, 0xbc, 0x48, 0x08, 0xa5, 0xa2, 0x2a, 0x94, 0xc4,
	0xab, 0x7b, 0x5c, 0x12, 0xfa, 0x95, 0x29, 0x54, 0x2e, 0x85, 0x12, 0xd6, 0x0b, 0x9e, 0x52, 0x74,
	0xae, 0x42, 0xff, 0x50, 0xa1, 0xd5, 0xbd, 0xf3, 0x24, 0x13, 0x22, 0x5b, 0x30, 0xac, 0x65, 0x69,
	0xf5, 0x05, 0x93, 0xa2, 0x6e, 0x3d, 0x9c, 0xdb, 0x4c, 0x64, 0x42, 0x43, 0xdc, 0xa0, 0x96, 0xbd,
	0x5b, 0x03, 0xf8, 0x74, 0x74, 0xf4, 0xea, 0xb7, 0x56, 0x63, 0xbd, 0x7b, 0x48, 0x14, 0xb1, 0xfa,
	0xd0, 0x54, 0x75, 0xc9, 0x6c, 0xe0, 0x01, 0xff, 0xba, 0x1b, 0xa0, 0xff, 0x3c, 0x04, 0xc5, 0x75,
	0xc9, 0x22, 0x2d, 0xb5, 0x2c, 0x68, 0xce, 0x88, 0x22, 0xf6, 0x03, 0x0f, 0xf8, 0x8f, 0x23, 0x8d,
	0x1b, 0x2e, 0x67, 0xb9, 0xb0, 0x2f, 0x3c, 0xe0, 0x3f, 0x8a, 0x34, 0xbe, 0x7b, 0x0b, 0x2f, 0x07,
	0x42, 0xe6, 0x42, 0xc6, 0xdf, 0xac, 0x57, 0xf0, 0x32, 0x67, 0x52, 0x92, 0x8c, 0x49, 0x1b, 0x78,
	0x17, 0xfe, 0x55, 0xf7, 0x16, 0xb5, 0xd1, 0xd0, 0xdf, 0x68, 0xa8, 0x5f, 0xd4, 0xd1, 0x71, 0xea,
	0xe5, 0x14, 0x9a, 0xcd, 0x4e, 0xeb, 0x39, 0xec, 0xc4, 0x9f, 0xc6, 0x61, 0x32, 0xf9, 0xf0, 0x71,
	0x1c, 0x0e, 0x46, 0xef, 0x47, 0xe1, 0xb0, 0x63, 0x38, 0x37, 0x9b, 0xad, 0x77, 0x75, 0x46, 0x59,
	0xcf, 0xe0, 0x8d, 0x1e, 0x0b, 0xa7, 0xe1, 0x60, 0x12, 0x87, 0x49, 0x3c, 0xed, 0x00, 0xe7, 0x7a,
	0xb3, 0xf5, 0xe0, 0x89, 0x71, 0xcc, 0xf5, 0x0f, 0xd7, 0x78, 0x97, 0xfc, 0xdc, 0xbb, 0x60, 0xb7,
	0x77, 0xc1, 0xef, 0xbd, 0x0b, 0xbe, 0x1f, 0x5c, 0x63, 0x77, 0x70, 0x8d, 0x5f, 0x07, 0xd7, 0xf8,
	0x1c, 0x66, 0x5c, 0xcd, 0xab, 0x14, 0x51, 0x91, 0x63, 0xaa, 0x4f, 0xc7, 0x3c, 0xa5, 0x41, 0x26,
	0xf0, 0xaa, 0x87, 0x73, 0x31, 0xab, 0x16, 0x4c, 0x36, 0x65, 0x4b, 0xdc, 0x7d, 0x13, 0x9c, 0x1e,
	0x15, 0x1c, 0x7b, 0x6e, 0xfe, 0x23, 0xd3, 0x87, 0x3a, 0x52, 0xef, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdd, 0x25, 0x3e, 0xc6, 0x1c, 0x02, 0x00, 0x00,
}

func (m *InterchainAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CosmosTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterchainAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPacket(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *CosmosTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterchainAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *CosmosTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: CosmosTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
