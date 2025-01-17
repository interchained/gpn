// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/validator.proto

package types

import (
	fmt "fmt"
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

type Validator struct {
	Address     string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Description *ValidatorDescription `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e2276f43ab77aa3, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetDescription() *ValidatorDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

type ValidatorDescription struct {
	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Moniker         string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Website         string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	SecurityContact string `protobuf:"bytes,4,opt,name=securityContact,proto3" json:"securityContact,omitempty"`
	Details         string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *ValidatorDescription) Reset()         { *m = ValidatorDescription{} }
func (m *ValidatorDescription) String() string { return proto.CompactTextString(m) }
func (*ValidatorDescription) ProtoMessage()    {}
func (*ValidatorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e2276f43ab77aa3, []int{1}
}
func (m *ValidatorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorDescription.Merge(m, src)
}
func (m *ValidatorDescription) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorDescription proto.InternalMessageInfo

func (m *ValidatorDescription) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *ValidatorDescription) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *ValidatorDescription) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *ValidatorDescription) GetSecurityContact() string {
	if m != nil {
		return m.SecurityContact
	}
	return ""
}

func (m *ValidatorDescription) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*Validator)(nil), "tendermint.spn.profile.Validator")
	proto.RegisterType((*ValidatorDescription)(nil), "tendermint.spn.profile.ValidatorDescription")
}

func init() { proto.RegisterFile("profile/validator.proto", fileDescriptor_8e2276f43ab77aa3) }

var fileDescriptor_8e2276f43ab77aa3 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x6b, 0xfe, 0xd7, 0x1d, 0x90, 0x22, 0x04, 0x16, 0x83, 0x55, 0x75, 0x2a, 0x12, 0x72,
	0x24, 0x78, 0x03, 0xca, 0xcc, 0xd0, 0x81, 0x81, 0x2d, 0x8d, 0x0f, 0x38, 0x91, 0xd8, 0x96, 0x7d,
	0x01, 0xf2, 0x16, 0x3c, 0x07, 0x4f, 0xc2, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x92, 0x26, 0x2d,
	0x42, 0x1d, 0x7f, 0xbe, 0xcf, 0x3f, 0x7d, 0xba, 0xe3, 0x67, 0xce, 0xdb, 0x47, 0xcc, 0x20, 0x7e,
	0x4d, 0x32, 0xd4, 0x09, 0x59, 0xaf, 0x9c, 0xb7, 0x64, 0xa3, 0x53, 0x02, 0xa3, 0xc1, 0xe7, 0x68,
	0x48, 0x05, 0x67, 0x54, 0x97, 0x9b, 0x14, 0x7c, 0x78, 0xdf, 0x47, 0x23, 0xc1, 0x0f, 0x13, 0xad,
	0x3d, 0x84, 0x20, 0xd8, 0x98, 0x4d, 0x87, 0xf3, 0x1e, 0xa3, 0x3b, 0x3e, 0xd2, 0x10, 0x52, 0x8f,
	0x8e, 0xd0, 0x1a, 0xb1, 0x33, 0x66, 0xd3, 0xd1, 0xd5, 0xa5, 0xda, 0x2e, 0x55, 0x6b, 0xe3, 0xed,
	0xe6, 0xcf, 0xfc, 0xaf, 0x60, 0xf2, 0xc9, 0xf8, 0xc9, 0xb6, 0x54, 0x74, 0xce, 0x8f, 0x50, 0x83,
	0x21, 0xa4, 0xb2, 0xeb, 0xb0, 0xe6, 0xa6, 0x5e, 0x6e, 0x0d, 0xbe, 0x80, 0x6f, 0x0b, 0x0c, 0xe7,
	0x3d, 0x36, 0x93, 0x37, 0x58, 0x04, 0x24, 0x10, 0xbb, 0xab, 0x49, 0x87, 0xd1, 0x94, 0x1f, 0x07,
	0x48, 0x0b, 0x8f, 0x54, 0xce, 0xac, 0xa1, 0x24, 0x25, 0xb1, 0xd7, 0x26, 0xfe, 0x3f, 0x37, 0x0e,
	0x0d, 0x94, 0x60, 0x16, 0xc4, 0xfe, 0xca, 0xd1, 0xe1, 0xcd, 0xec, 0xab, 0x92, 0x6c, 0x59, 0x49,
	0xf6, 0x53, 0x49, 0xf6, 0x51, 0xcb, 0xc1, 0xb2, 0x96, 0x83, 0xef, 0x5a, 0x0e, 0x1e, 0x2e, 0x9e,
	0x90, 0x9e, 0x8b, 0x85, 0x4a, 0x6d, 0x1e, 0x6f, 0x76, 0x11, 0x07, 0x67, 0xe2, 0xf7, 0xb8, 0x3f,
	0x05, 0x95, 0x0e, 0xc2, 0xe2, 0xa0, 0xbd, 0xc3, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5,
	0x20, 0x13, 0x9d, 0xa2, 0x01, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Description != nil {
		{
			size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorDescription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorDescription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Description != nil {
		l = m.Description.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func (m *ValidatorDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Description == nil {
				m.Description = &ValidatorDescription{}
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *ValidatorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
