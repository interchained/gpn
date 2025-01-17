// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/genesis.proto

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

// GenesisState defines the profile module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	ValidatorList            []Validator            `protobuf:"bytes,4,rep,name=validatorList,proto3" json:"validatorList"`
	CoordinatorByAddressList []CoordinatorByAddress `protobuf:"bytes,3,rep,name=coordinatorByAddressList,proto3" json:"coordinatorByAddressList"`
	CoordinatorList          []Coordinator          `protobuf:"bytes,1,rep,name=coordinatorList,proto3" json:"coordinatorList"`
	CoordinatorCount         uint64                 `protobuf:"varint,2,opt,name=coordinatorCount,proto3" json:"coordinatorCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4bc1562021cf42, []int{0}
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

func (m *GenesisState) GetValidatorList() []Validator {
	if m != nil {
		return m.ValidatorList
	}
	return nil
}

func (m *GenesisState) GetCoordinatorByAddressList() []CoordinatorByAddress {
	if m != nil {
		return m.CoordinatorByAddressList
	}
	return nil
}

func (m *GenesisState) GetCoordinatorList() []Coordinator {
	if m != nil {
		return m.CoordinatorList
	}
	return nil
}

func (m *GenesisState) GetCoordinatorCount() uint64 {
	if m != nil {
		return m.CoordinatorCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.profile.GenesisState")
}

func init() { proto.RegisterFile("profile/genesis.proto", fileDescriptor_db4bc1562021cf42) }

var fileDescriptor_db4bc1562021cf42 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2b, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e,
	0xc8, 0xd3, 0x83, 0xaa, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd1, 0x07, 0xb1, 0x20,
	0xaa, 0xa5, 0xc4, 0x61, 0x86, 0x94, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x41, 0x25,
	0x24, 0x61, 0x12, 0xc9, 0xf9, 0xf9, 0x45, 0x29, 0x99, 0x79, 0x08, 0x29, 0xa5, 0x73, 0x4c, 0x5c,
	0x3c, 0xee, 0x10, 0x3b, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0x7c, 0xb9, 0x78, 0xe1, 0xda, 0x7d,
	0x32, 0x8b, 0x4b, 0x24, 0x58, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x14, 0xf5, 0xb0, 0x3b, 0x45, 0x2f,
	0x0c, 0xa6, 0xd8, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x54, 0xdd, 0x42, 0x79, 0x5c, 0x12,
	0x48, 0x96, 0x3a, 0x55, 0x3a, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x83, 0x4d, 0x66, 0x06, 0x9b,
	0xac, 0x83, 0xcb, 0x64, 0x67, 0x2c, 0xfa, 0xa0, 0x96, 0xe0, 0x34, 0x53, 0x28, 0x98, 0x8b, 0x1f,
	0x49, 0x0e, 0x6c, 0x0d, 0x23, 0xd8, 0x1a, 0x65, 0x62, 0xac, 0x81, 0x98, 0x8e, 0x6e, 0x82, 0x90,
	0x16, 0x97, 0x00, 0x92, 0x90, 0x73, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b,
	0x10, 0x86, 0xb8, 0x93, 0xf3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0xdc, 0xa2, 0x5f, 0x5c,
	0x90, 0xa7, 0x5f, 0xa1, 0x0f, 0x8b, 0xa1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xe4,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x1b, 0x26, 0x74, 0x17, 0x02, 0x00, 0x00,
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
	if len(m.ValidatorList) > 0 {
		for iNdEx := len(m.ValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CoordinatorByAddressList) > 0 {
		for iNdEx := len(m.CoordinatorByAddressList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoordinatorByAddressList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.CoordinatorCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CoordinatorCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CoordinatorList) > 0 {
		for iNdEx := len(m.CoordinatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoordinatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.CoordinatorList) > 0 {
		for _, e := range m.CoordinatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CoordinatorCount != 0 {
		n += 1 + sovGenesis(uint64(m.CoordinatorCount))
	}
	if len(m.CoordinatorByAddressList) > 0 {
		for _, e := range m.CoordinatorByAddressList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorList) > 0 {
		for _, e := range m.ValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorList", wireType)
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
			m.CoordinatorList = append(m.CoordinatorList, Coordinator{})
			if err := m.CoordinatorList[len(m.CoordinatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorCount", wireType)
			}
			m.CoordinatorCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorByAddressList", wireType)
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
			m.CoordinatorByAddressList = append(m.CoordinatorByAddressList, CoordinatorByAddress{})
			if err := m.CoordinatorByAddressList[len(m.CoordinatorByAddressList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorList", wireType)
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
			m.ValidatorList = append(m.ValidatorList, Validator{})
			if err := m.ValidatorList[len(m.ValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
