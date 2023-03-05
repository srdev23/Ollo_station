// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/emissions/v1/emissions.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Emission struct {
	Id                    uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                 string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	AmountIn              github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in" yaml:"amount_in"`
	AmountOut             github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=amount_out,json=amountOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out" yaml:"amount_out"`
	CreatedAt             time.Time                              `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	InterestAccumulated   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=interest_accumulated,json=interestAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"interest_accumulated" yaml:"interest_accumulated"`
	ClosingFeeAccumulated github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=closing_fee_accumulated,json=closingFeeAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"closing_fee_accumulated" yaml:"interest_accumulated"`
	BlockHeight           int64                                  `protobuf:"varint,10,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	BlockTime             time.Time                              `protobuf:"bytes,11,opt,name=block_time,json=blockTime,proto3,stdtime" json:"block_time" yaml:"block_time"`
}

func (m *Emission) Reset()         { *m = Emission{} }
func (m *Emission) String() string { return proto.CompactTextString(m) }
func (*Emission) ProtoMessage()    {}
func (*Emission) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c5a0be93b8daef5, []int{0}
}
func (m *Emission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Emission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Emission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Emission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Emission.Merge(m, src)
}
func (m *Emission) XXX_Size() int {
	return m.Size()
}
func (m *Emission) XXX_DiscardUnknown() {
	xxx_messageInfo_Emission.DiscardUnknown(m)
}

var xxx_messageInfo_Emission proto.InternalMessageInfo

type EmissionsEmission struct {
	Id        uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	AmountIn  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount_in,json=amountIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in" yaml:"amount_in"`
	AmountOut github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount_out,json=amountOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out" yaml:"amount_out"`
	CreatedAt time.Time                              `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
}

func (m *EmissionsEmission) Reset()         { *m = EmissionsEmission{} }
func (m *EmissionsEmission) String() string { return proto.CompactTextString(m) }
func (*EmissionsEmission) ProtoMessage()    {}
func (*EmissionsEmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c5a0be93b8daef5, []int{1}
}
func (m *EmissionsEmission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmissionsEmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmissionsEmission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmissionsEmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmissionsEmission.Merge(m, src)
}
func (m *EmissionsEmission) XXX_Size() int {
	return m.Size()
}
func (m *EmissionsEmission) XXX_DiscardUnknown() {
	xxx_messageInfo_EmissionsEmission.DiscardUnknown(m)
}

var xxx_messageInfo_EmissionsEmission proto.InternalMessageInfo

type EmissionsEmissionRewards struct {
	User        string                                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" yaml:"user"`
	BlockHeight uint64                                 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	Amount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
}

func (m *EmissionsEmissionRewards) Reset()         { *m = EmissionsEmissionRewards{} }
func (m *EmissionsEmissionRewards) String() string { return proto.CompactTextString(m) }
func (*EmissionsEmissionRewards) ProtoMessage()    {}
func (*EmissionsEmissionRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c5a0be93b8daef5, []int{2}
}
func (m *EmissionsEmissionRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmissionsEmissionRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmissionsEmissionRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmissionsEmissionRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmissionsEmissionRewards.Merge(m, src)
}
func (m *EmissionsEmissionRewards) XXX_Size() int {
	return m.Size()
}
func (m *EmissionsEmissionRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_EmissionsEmissionRewards.DiscardUnknown(m)
}

var xxx_messageInfo_EmissionsEmissionRewards proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Emission)(nil), "ollo.emissions.v1.Emission")
	proto.RegisterType((*EmissionsEmission)(nil), "ollo.emissions.v1.EmissionsEmission")
	proto.RegisterType((*EmissionsEmissionRewards)(nil), "ollo.emissions.v1.EmissionsEmissionRewards")
}

func init() { proto.RegisterFile("ollo/emissions/v1/emissions.proto", fileDescriptor_8c5a0be93b8daef5) }

var fileDescriptor_8c5a0be93b8daef5 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0xde, 0x59, 0x60, 0xa5, 0xb3, 0x60, 0xa4, 0x60, 0xa8, 0x18, 0xda, 0x75, 0x0e, 0xba, 0x17,
	0xda, 0xa0, 0x37, 0x4c, 0x34, 0xd4, 0x68, 0xe4, 0x60, 0x34, 0xd5, 0x44, 0xe3, 0xa5, 0xe9, 0xb6,
	0x43, 0x99, 0xd0, 0x76, 0x48, 0x67, 0x0a, 0x72, 0xf3, 0xe2, 0x9d, 0xb3, 0xbf, 0xc0, 0x1f, 0xe0,
	0x8f, 0xe0, 0x88, 0x9e, 0x8c, 0x87, 0xaa, 0xcb, 0xd5, 0xd3, 0xfe, 0x02, 0xd3, 0x99, 0x59, 0xba,
	0x1b, 0x34, 0x71, 0x13, 0xf1, 0xb4, 0x7d, 0xdf, 0x9b, 0xf7, 0x7d, 0x6f, 0xbe, 0x79, 0x33, 0x0b,
	0x6f, 0xd0, 0x24, 0xa1, 0x0e, 0x4e, 0x09, 0x63, 0x84, 0x66, 0xcc, 0xd9, 0x5f, 0xaf, 0x03, 0x7b,
	0x2f, 0xa7, 0x9c, 0xea, 0x0b, 0xd5, 0x12, 0xbb, 0x46, 0xf7, 0xd7, 0x57, 0x96, 0x62, 0x1a, 0x53,
	0x91, 0x75, 0xaa, 0x2f, 0xb9, 0x70, 0xc5, 0x8a, 0x29, 0x8d, 0x13, 0xec, 0x88, 0xa8, 0x57, 0x6c,
	0x3b, 0x9c, 0xa4, 0x98, 0xf1, 0x20, 0xdd, 0x53, 0x0b, 0xae, 0x85, 0x94, 0xa5, 0x94, 0xf9, 0xb2,
	0x52, 0x06, 0x32, 0x85, 0xde, 0xb7, 0xe0, 0xec, 0x43, 0x25, 0xa1, 0x5f, 0x86, 0x4d, 0x12, 0x19,
	0xa0, 0x03, 0xba, 0xd3, 0x5e, 0x93, 0x44, 0xfa, 0x3d, 0x38, 0x43, 0x0f, 0x32, 0x9c, 0x1b, 0xcd,
	0x0e, 0xe8, 0x6a, 0x6e, 0x77, 0x50, 0x5a, 0x73, 0x87, 0x41, 0x9a, 0x6c, 0x20, 0x01, 0xa3, 0xcf,
	0x1f, 0xd7, 0x96, 0x14, 0xdb, 0x66, 0x14, 0xe5, 0x98, 0xb1, 0xe7, 0x3c, 0x27, 0x59, 0xec, 0xc9,
	0x32, 0xdd, 0x87, 0x5a, 0x90, 0xd2, 0x22, 0xe3, 0x3e, 0xc9, 0x8c, 0x19, 0xc1, 0xe1, 0x1e, 0x97,
	0x56, 0xe3, 0x6b, 0x69, 0xdd, 0x8c, 0x09, 0xdf, 0x29, 0x7a, 0x76, 0x48, 0x53, 0xd5, 0x90, 0xfa,
	0x59, 0x63, 0xd1, 0xae, 0xc3, 0x0f, 0xf7, 0x30, 0xb3, 0xb7, 0x32, 0x3e, 0x28, 0xad, 0x2b, 0x52,
	0xf1, 0x8c, 0x08, 0x79, 0xb3, 0xf2, 0x7b, 0x2b, 0xd3, 0x7b, 0x10, 0x2a, 0x9c, 0x16, 0xdc, 0x68,
	0x09, 0x85, 0x07, 0x13, 0x2b, 0x2c, 0x8c, 0x29, 0xd0, 0x82, 0x23, 0x4f, 0xf5, 0xfd, 0xb4, 0xe0,
	0xfa, 0x2b, 0x08, 0xc3, 0x1c, 0x07, 0x1c, 0x47, 0x7e, 0xc0, 0x8d, 0x4b, 0x1d, 0xd0, 0x6d, 0xdf,
	0x5e, 0xb1, 0xa5, 0xe5, 0xf6, 0xd0, 0x72, 0xfb, 0xc5, 0xd0, 0x72, 0x77, 0xb5, 0xd2, 0xaf, 0x59,
	0xeb, 0x5a, 0x74, 0xf4, 0xcd, 0x02, 0x9e, 0xa6, 0x80, 0x4d, 0xae, 0xbf, 0x05, 0x70, 0x89, 0x64,
	0x1c, 0xe7, 0x98, 0x71, 0x3f, 0x08, 0xc3, 0x22, 0x2d, 0x92, 0x2a, 0x65, 0xcc, 0x8a, 0x8d, 0x3c,
	0x99, 0x78, 0x23, 0xd7, 0xa5, 0xe4, 0xef, 0x38, 0x91, 0xb7, 0x38, 0x84, 0x37, 0x6b, 0x54, 0x7f,
	0x07, 0xe0, 0x72, 0x98, 0x50, 0x46, 0xb2, 0xd8, 0xdf, 0xc6, 0x78, 0xac, 0x0b, 0xed, 0x22, 0xba,
	0xb8, 0xaa, 0xd4, 0x1e, 0x61, 0x3c, 0xda, 0xc7, 0x06, 0x9c, 0xeb, 0x25, 0x34, 0xdc, 0xf5, 0x77,
	0x30, 0x89, 0x77, 0xb8, 0x01, 0x3b, 0xa0, 0x3b, 0xe5, 0x2e, 0x0f, 0x4a, 0x6b, 0x51, 0xb2, 0x8d,
	0x66, 0x91, 0xd7, 0x16, 0xe1, 0x63, 0x11, 0x55, 0x07, 0x24, 0xb3, 0xd5, 0xd8, 0x1b, 0xed, 0x49,
	0x0f, 0xa8, 0xae, 0x55, 0x07, 0x24, 0x80, 0x6a, 0x39, 0xfa, 0xd4, 0x84, 0x0b, 0xc3, 0xcb, 0xc1,
	0xce, 0x6e, 0xc9, 0x6a, 0x7d, 0x4b, 0xdc, 0xf9, 0x41, 0x69, 0x69, 0x6a, 0xbf, 0x11, 0x12, 0x97,
	0x66, 0x6c, 0xe8, 0x9b, 0x17, 0x3e, 0xf4, 0x53, 0xff, 0x61, 0xe8, 0xa7, 0xff, 0xdd, 0xd0, 0xa3,
	0x9f, 0x00, 0x1a, 0xe7, 0x3c, 0xf5, 0xf0, 0x41, 0x90, 0x47, 0x4c, 0xbf, 0x0b, 0xa7, 0x0b, 0x86,
	0x73, 0x61, 0xae, 0xe6, 0xde, 0x1a, 0x94, 0x56, 0x5b, 0x12, 0x56, 0xe8, 0x9f, 0x9f, 0x1b, 0x51,
	0x74, 0x6e, 0x86, 0x9a, 0xe2, 0x84, 0xfe, 0x6e, 0x86, 0x5e, 0xc2, 0x96, 0xdc, 0xbc, 0xf2, 0xf3,
	0xfe, 0xc4, 0x7e, 0xce, 0x8f, 0xfa, 0x89, 0x3c, 0x45, 0xe7, 0x3e, 0x3b, 0xfe, 0x61, 0x36, 0x3e,
	0xf4, 0xcd, 0xc6, 0x71, 0xdf, 0x04, 0x27, 0x7d, 0x13, 0x7c, 0xef, 0x9b, 0xe0, 0xe8, 0xd4, 0x6c,
	0x9c, 0x9c, 0x9a, 0x8d, 0x2f, 0xa7, 0x66, 0xe3, 0xb5, 0x3d, 0x22, 0x51, 0xbd, 0xf8, 0x6b, 0x8c,
	0x07, 0x9c, 0xd0, 0x4c, 0x04, 0xce, 0x9b, 0x91, 0xff, 0x08, 0x21, 0xd7, 0x6b, 0x09, 0xfb, 0xef,
	0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x1a, 0xc2, 0x54, 0x42, 0x06, 0x00, 0x00,
}

func (m *Emission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Emission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Emission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEmissions(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	if m.BlockHeight != 0 {
		i = encodeVarintEmissions(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.ClosingFeeAccumulated.Size()
		i -= size
		if _, err := m.ClosingFeeAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.InterestAccumulated.Size()
		i -= size
		if _, err := m.InterestAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEmissions(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	{
		size := m.AmountOut.Size()
		i -= size
		if _, err := m.AmountOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.AmountIn.Size()
		i -= size
		if _, err := m.AmountIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEmissions(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEmissions(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmissionsEmission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmissionsEmission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmissionsEmission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEmissions(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	{
		size := m.AmountOut.Size()
		i -= size
		if _, err := m.AmountOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AmountIn.Size()
		i -= size
		if _, err := m.AmountIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintEmissions(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmissionsEmissionRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmissionsEmissionRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmissionsEmissionRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEmissions(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.BlockHeight != 0 {
		i = encodeVarintEmissions(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEmissions(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEmissions(dAtA []byte, offset int, v uint64) int {
	offset -= sovEmissions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Emission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEmissions(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEmissions(uint64(l))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovEmissions(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovEmissions(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovEmissions(uint64(l))
	l = m.InterestAccumulated.Size()
	n += 1 + l + sovEmissions(uint64(l))
	l = m.ClosingFeeAccumulated.Size()
	n += 1 + l + sovEmissions(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovEmissions(uint64(m.BlockHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime)
	n += 1 + l + sovEmissions(uint64(l))
	return n
}

func (m *EmissionsEmission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEmissions(uint64(m.Id))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovEmissions(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovEmissions(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovEmissions(uint64(l))
	return n
}

func (m *EmissionsEmissionRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEmissions(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEmissions(uint64(m.BlockHeight))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEmissions(uint64(l))
	return n
}

func sovEmissions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEmissions(x uint64) (n int) {
	return sovEmissions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Emission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmissions
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
			return fmt.Errorf("proto: Emission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Emission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFeeAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFeeAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEmissions
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
func (m *EmissionsEmission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmissions
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
			return fmt.Errorf("proto: EmissionsEmission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmissionsEmission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEmissions
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
func (m *EmissionsEmissionRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmissions
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
			return fmt.Errorf("proto: EmissionsEmissionRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmissionsEmissionRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmissions
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
				return ErrInvalidLengthEmissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEmissions
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
func skipEmissions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEmissions
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
					return 0, ErrIntOverflowEmissions
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
					return 0, ErrIntOverflowEmissions
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
				return 0, ErrInvalidLengthEmissions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEmissions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEmissions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEmissions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEmissions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEmissions = fmt.Errorf("proto: unexpected end of group")
)