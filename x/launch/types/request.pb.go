// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/request.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type Request struct {
	ChainID   string     `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	RequestID uint64     `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Creator   string     `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedAt int64      `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Content   *types.Any `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e4b0ce31bf039, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Request) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *Request) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Request) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Request) GetContent() *types.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type AccountRemoval struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AccountRemoval) Reset()         { *m = AccountRemoval{} }
func (m *AccountRemoval) String() string { return proto.CompactTextString(m) }
func (*AccountRemoval) ProtoMessage()    {}
func (*AccountRemoval) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e4b0ce31bf039, []int{1}
}
func (m *AccountRemoval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountRemoval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountRemoval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountRemoval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRemoval.Merge(m, src)
}
func (m *AccountRemoval) XXX_Size() int {
	return m.Size()
}
func (m *AccountRemoval) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRemoval.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRemoval proto.InternalMessageInfo

func (m *AccountRemoval) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ValidatorRemoval struct {
	ValAddress string `protobuf:"bytes,1,opt,name=valAddress,proto3" json:"valAddress,omitempty"`
}

func (m *ValidatorRemoval) Reset()         { *m = ValidatorRemoval{} }
func (m *ValidatorRemoval) String() string { return proto.CompactTextString(m) }
func (*ValidatorRemoval) ProtoMessage()    {}
func (*ValidatorRemoval) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e4b0ce31bf039, []int{2}
}
func (m *ValidatorRemoval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorRemoval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorRemoval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorRemoval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorRemoval.Merge(m, src)
}
func (m *ValidatorRemoval) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorRemoval) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorRemoval.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorRemoval proto.InternalMessageInfo

func (m *ValidatorRemoval) GetValAddress() string {
	if m != nil {
		return m.ValAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "tendermint.spn.launch.Request")
	proto.RegisterType((*AccountRemoval)(nil), "tendermint.spn.launch.AccountRemoval")
	proto.RegisterType((*ValidatorRemoval)(nil), "tendermint.spn.launch.ValidatorRemoval")
}

func init() { proto.RegisterFile("launch/request.proto", fileDescriptor_028e4b0ce31bf039) }

var fileDescriptor_028e4b0ce31bf039 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xfb, 0x30,
	0x1c, 0x5d, 0xfe, 0xdb, 0xdf, 0xb1, 0x08, 0x43, 0xca, 0x84, 0x6e, 0x48, 0x28, 0x3b, 0xf5, 0x62,
	0x02, 0x7a, 0xf3, 0x20, 0xb4, 0x0e, 0x61, 0xd7, 0x1c, 0x3c, 0x78, 0x91, 0x2c, 0x8d, 0x5b, 0xa1,
	0x4d, 0x6a, 0x92, 0x0e, 0xf7, 0x2d, 0xfc, 0x30, 0x5e, 0xbd, 0x8b, 0xa7, 0xe1, 0xc9, 0xa3, 0x6c,
	0x5f, 0x44, 0xd6, 0xb4, 0x4c, 0xd1, 0x5b, 0xde, 0xcb, 0x7b, 0x8f, 0xf7, 0xf8, 0xc1, 0x41, 0xc6,
	0x4a, 0xc9, 0x17, 0x44, 0x8b, 0x87, 0x52, 0x18, 0x8b, 0x0b, 0xad, 0xac, 0xf2, 0x8e, 0xad, 0x90,
	0x89, 0xd0, 0x79, 0x2a, 0x2d, 0x36, 0x85, 0xc4, 0x4e, 0x34, 0x1a, 0xce, 0x95, 0x9a, 0x67, 0x82,
	0x54, 0xa2, 0x59, 0x79, 0x4f, 0x98, 0x5c, 0x39, 0xc7, 0x68, 0xc8, 0x95, 0xc9, 0x95, 0xb9, 0xab,
	0x10, 0x71, 0xc0, 0x7d, 0x8d, 0x5f, 0x00, 0xec, 0x52, 0x17, 0xef, 0xf9, 0xb0, 0xcb, 0x17, 0x2c,
	0x95, 0xd3, 0x89, 0x0f, 0x02, 0x10, 0xf6, 0x68, 0x03, 0xbd, 0x13, 0xd8, 0xab, 0x3b, 0x4c, 0x27,
	0xfe, 0xbf, 0x00, 0x84, 0x1d, 0xba, 0x27, 0x2a, 0x9f, 0x16, 0xcc, 0x2a, 0xed, 0xb7, 0x6b, 0x9f,
	0x83, 0x3b, 0x5f, 0xf5, 0x14, 0x49, 0x64, 0xfd, 0x4e, 0x00, 0xc2, 0x36, 0xdd, 0x13, 0x5e, 0x0c,
	0xbb, 0x5c, 0x49, 0x2b, 0xa4, 0xf5, 0xff, 0x07, 0x20, 0x3c, 0x3c, 0x1b, 0x60, 0xb7, 0x01, 0x37,
	0x1b, 0x70, 0x24, 0x57, 0xb1, 0xf7, 0xf6, 0x7c, 0xda, 0xaf, 0x3b, 0x5e, 0x39, 0x3d, 0x6d, 0x8c,
	0xe3, 0x4b, 0xd8, 0x8f, 0x38, 0x57, 0xa5, 0xb4, 0x54, 0xe4, 0x6a, 0xc9, 0xb2, 0x5d, 0x1b, 0x96,
	0x24, 0x5a, 0x18, 0xd3, 0xac, 0xa8, 0xe1, 0x85, 0xf7, 0xfe, 0x2b, 0x68, 0x7c, 0x0d, 0x8f, 0x6e,
	0x58, 0x96, 0x26, 0xbb, 0xba, 0x4d, 0x02, 0x82, 0x70, 0xc9, 0xb2, 0xe8, 0x47, 0xc8, 0x37, 0xe6,
	0xaf, 0x9c, 0x38, 0x7e, 0xdd, 0x20, 0xb0, 0xde, 0x20, 0xf0, 0xb9, 0x41, 0xe0, 0x69, 0x8b, 0x5a,
	0xeb, 0x2d, 0x6a, 0x7d, 0x6c, 0x51, 0xeb, 0x36, 0x9c, 0xa7, 0x76, 0x51, 0xce, 0x30, 0x57, 0x39,
	0xd9, 0x5f, 0x8e, 0x98, 0x42, 0x92, 0x47, 0x52, 0x1f, 0xd8, 0xae, 0x0a, 0x61, 0x66, 0x07, 0xd5,
	0xec, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xea, 0x82, 0xc0, 0xf7, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.CreatedAt != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RequestID != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountRemoval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountRemoval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountRemoval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorRemoval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorRemoval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorRemoval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValAddress) > 0 {
		i -= len(m.ValAddress)
		copy(dAtA[i:], m.ValAddress)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ValAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.RequestID != 0 {
		n += 1 + sovRequest(uint64(m.RequestID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovRequest(uint64(m.CreatedAt))
	}
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func (m *AccountRemoval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func (m *ValidatorRemoval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValAddress)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *AccountRemoval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: AccountRemoval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountRemoval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *ValidatorRemoval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: ValidatorRemoval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorRemoval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
