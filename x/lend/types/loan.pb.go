// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/lend/v1/loan.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type LoanStatus int32

const (
	LoanStatusUnspecified LoanStatus = 0
	LoanStatusActive      LoanStatus = 1
	LoanStatusRepaid      LoanStatus = 2
	LoanStatusCancelled   LoanStatus = 3
	LoanStatusRequested   LoanStatus = 4
)

var LoanStatus_name = map[int32]string{
	0: "LOAN_STATUS_UNSPECIFIED",
	1: "LOAN_STATUS_ACTIVE",
	2: "LOAN_STATUS_REPAID",
	3: "LOAN_STATUS_CANCELLED",
	4: "LOAN_STATUS_REQUESTED",
}

var LoanStatus_value = map[string]int32{
	"LOAN_STATUS_UNSPECIFIED": 0,
	"LOAN_STATUS_ACTIVE":      1,
	"LOAN_STATUS_REPAID":      2,
	"LOAN_STATUS_CANCELLED":   3,
	"LOAN_STATUS_REQUESTED":   4,
}

func (x LoanStatus) String() string {
	return proto.EnumName(LoanStatus_name, int32(x))
}

func (LoanStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d765500d80eeb3d, []int{0}
}

type Loan struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Borrower   string `protobuf:"bytes,2,opt,name=borrower,proto3" json:"borrower,omitempty" yaml:"borrower"`
	Lender     string `protobuf:"bytes,3,opt,name=lender,proto3" json:"lender,omitempty" yaml:"lender"`
	Amount     string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Collateral string `protobuf:"bytes,5,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Fee        string `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Deadline   string `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status     string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Loan) Reset()         { *m = Loan{} }
func (m *Loan) String() string { return proto.CompactTextString(m) }
func (*Loan) ProtoMessage()    {}
func (*Loan) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d765500d80eeb3d, []int{0}
}
func (m *Loan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loan.Merge(m, src)
}
func (m *Loan) XXX_Size() int {
	return m.Size()
}
func (m *Loan) XXX_DiscardUnknown() {
	xxx_messageInfo_Loan.DiscardUnknown(m)
}

var xxx_messageInfo_Loan proto.InternalMessageInfo

func (m *Loan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Loan) GetBorrower() string {
	if m != nil {
		return m.Borrower
	}
	return ""
}

func (m *Loan) GetLender() string {
	if m != nil {
		return m.Lender
	}
	return ""
}

func (m *Loan) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Loan) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *Loan) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Loan) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Loan) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type LoanEntry struct {
	Lender    string                                  `protobuf:"bytes,1,opt,name=lender,proto3" json:"lender,omitempty" yaml:"lender"`
	CreatedAt time.Time                               `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	AmountIn  github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	asset_id  uint64                                  `protobuf:"varint,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *LoanEntry) Reset()         { *m = LoanEntry{} }
func (m *LoanEntry) String() string { return proto.CompactTextString(m) }
func (*LoanEntry) ProtoMessage()    {}
func (*LoanEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d765500d80eeb3d, []int{1}
}
func (m *LoanEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoanEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoanEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoanEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoanEntry.Merge(m, src)
}
func (m *LoanEntry) XXX_Size() int {
	return m.Size()
}
func (m *LoanEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LoanEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LoanEntry proto.InternalMessageInfo

func (m *LoanEntry) GetLender() string {
	if m != nil {
		return m.Lender
	}
	return ""
}

func (m *LoanEntry) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *LoanEntry) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *LoanEntry) Getasset_id() uint64 {
	if m != nil {
		return m.asset_id
	}
	return 0
}

func init() {
	proto.RegisterEnum("ollo.lend.v1.LoanStatus", LoanStatus_name, LoanStatus_value)
	proto.RegisterType((*Loan)(nil), "ollo.lend.v1.Loan")
	proto.RegisterType((*LoanEntry)(nil), "ollo.lend.v1.LoanEntry")
}

func init() { proto.RegisterFile("ollo/lend/v1/loan.proto", fileDescriptor_8d765500d80eeb3d) }

var fileDescriptor_8d765500d80eeb3d = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xd3, 0x7c, 0x69, 0x32, 0xfd, 0x7e, 0xfc, 0x4d, 0x5b, 0xea, 0x7a, 0x61, 0x47, 0xd9,
	0xd0, 0x22, 0x6a, 0x2b, 0x45, 0x62, 0xd1, 0x9d, 0xe3, 0x18, 0x29, 0x52, 0x54, 0x8a, 0x93, 0xb0,
	0x60, 0x13, 0x4d, 0xec, 0x69, 0x18, 0xe1, 0x78, 0x82, 0x67, 0x12, 0xa8, 0x78, 0x01, 0xd4, 0x05,
	0xea, 0x0b, 0x74, 0xc5, 0x2b, 0xf0, 0x10, 0x5d, 0xb0, 0xa8, 0x58, 0xb1, 0x0a, 0x28, 0x7d, 0x01,
	0xd4, 0x25, 0x2b, 0xe4, 0x19, 0xe7, 0x47, 0xed, 0x2a, 0x73, 0xee, 0x3d, 0xe7, 0xce, 0x9d, 0x7b,
	0x6e, 0x0c, 0x76, 0x68, 0x14, 0x51, 0x3b, 0xc2, 0x71, 0x68, 0x4f, 0x6a, 0x76, 0x44, 0x51, 0x6c,
	0x8d, 0x12, 0xca, 0x29, 0xfc, 0x3b, 0x4d, 0x58, 0x69, 0xc2, 0x9a, 0xd4, 0xf4, 0xdd, 0x80, 0xb2,
	0x21, 0x65, 0x3d, 0x91, 0xb3, 0x25, 0x90, 0x44, 0xdd, 0x90, 0xc8, 0xee, 0x23, 0x86, 0xed, 0x49,
	0xad, 0x8f, 0x39, 0xaa, 0xd9, 0x01, 0x25, 0x59, 0x21, 0xdd, 0x1c, 0x50, 0x3a, 0x88, 0xb0, 0x2d,
	0x50, 0x7f, 0x7c, 0x6a, 0x73, 0x32, 0xc4, 0x8c, 0xa3, 0xe1, 0x28, 0x23, 0x6c, 0x0d, 0xe8, 0x80,
	0xca, 0xc2, 0xe9, 0x49, 0x46, 0xab, 0xbf, 0x14, 0x50, 0x68, 0x51, 0x14, 0xc3, 0x7f, 0x41, 0x9e,
	0x84, 0x9a, 0x52, 0x51, 0xf6, 0x0a, 0x7e, 0x9e, 0x84, 0xd0, 0x06, 0xa5, 0x3e, 0x4d, 0x12, 0xfa,
	0x0e, 0x27, 0x5a, 0xbe, 0xa2, 0xec, 0x95, 0xeb, 0x9b, 0xb7, 0x53, 0xf3, 0xbf, 0x33, 0x34, 0x8c,
	0x8e, 0xaa, 0xf3, 0x4c, 0xd5, 0x5f, 0x90, 0xe0, 0x3e, 0x28, 0xa6, 0xcf, 0xc0, 0x89, 0xb6, 0x26,
	0xe8, 0xff, 0xdf, 0x4e, 0xcd, 0x7f, 0x24, 0x5d, 0xc6, 0xab, 0x7e, 0x46, 0x80, 0x0f, 0x40, 0x11,
	0x0d, 0xe9, 0x38, 0xe6, 0x5a, 0x21, 0xa5, 0xfa, 0x19, 0x82, 0x06, 0x00, 0x01, 0x8d, 0x22, 0xc4,
	0x71, 0x82, 0x22, 0xed, 0x2f, 0x91, 0x5b, 0x89, 0x40, 0x15, 0xac, 0x9d, 0x62, 0xac, 0x15, 0x45,
	0x22, 0x3d, 0x42, 0x1d, 0x94, 0x42, 0x8c, 0xc2, 0x88, 0xc4, 0x58, 0x5b, 0x17, 0xe1, 0x05, 0x4e,
	0x6f, 0x61, 0x1c, 0xf1, 0x31, 0xd3, 0x4a, 0xf2, 0x16, 0x89, 0xaa, 0x5f, 0xf3, 0xa0, 0x9c, 0x3e,
	0xd9, 0x8b, 0x79, 0x72, 0x06, 0x9d, 0x45, 0xdb, 0x8a, 0x68, 0x7b, 0xff, 0x5e, 0xdb, 0xdf, 0xbe,
	0x1c, 0x6c, 0x65, 0x56, 0x38, 0x61, 0x98, 0x60, 0xc6, 0xda, 0x3c, 0x21, 0xf1, 0x60, 0xf1, 0x1c,
	0x17, 0x80, 0x20, 0xc1, 0x88, 0xe3, 0xb0, 0x87, 0xb8, 0x18, 0xd6, 0xc6, 0xa1, 0x6e, 0x49, 0x3f,
	0xac, 0xb9, 0x1f, 0x56, 0x67, 0xee, 0x47, 0xbd, 0x74, 0x35, 0x35, 0x73, 0x17, 0x3f, 0x4c, 0xc5,
	0x2f, 0x67, 0x3a, 0x87, 0xc3, 0x0f, 0xa0, 0x2c, 0xa7, 0xd0, 0x23, 0xb1, 0x98, 0xe0, 0xc6, 0xe1,
	0xae, 0x95, 0x5d, 0x9b, 0x7a, 0x6e, 0x65, 0x9e, 0x5b, 0x2e, 0x25, 0x71, 0xdd, 0x4d, 0x4b, 0xdc,
	0x4e, 0x4d, 0x55, 0x76, 0xba, 0x50, 0x56, 0x7f, 0x4f, 0xcd, 0x87, 0x03, 0xc2, 0x5f, 0x8f, 0xfb,
	0x56, 0x40, 0x87, 0xd9, 0x0a, 0x65, 0x3f, 0x07, 0x2c, 0x7c, 0x63, 0xf3, 0xb3, 0x11, 0x66, 0xa2,
	0x88, 0x5f, 0x92, 0xb2, 0x66, 0x0c, 0x8f, 0x40, 0x09, 0x31, 0x86, 0x79, 0x8f, 0x84, 0xc2, 0x92,
	0x42, 0xdd, 0x9c, 0x4d, 0xcd, 0x45, 0x6c, 0x69, 0xfc, 0x3c, 0x52, 0xf5, 0xd7, 0xc5, 0xb1, 0x19,
	0x3e, 0xfa, 0x94, 0x07, 0x20, 0x1d, 0x67, 0x5b, 0x4c, 0x17, 0x3e, 0x05, 0x3b, 0xad, 0xe7, 0xce,
	0x71, 0xaf, 0xdd, 0x71, 0x3a, 0xdd, 0x76, 0xaf, 0x7b, 0xdc, 0x3e, 0xf1, 0xdc, 0xe6, 0xb3, 0xa6,
	0xd7, 0x50, 0x73, 0xfa, 0xee, 0xf9, 0x65, 0x65, 0x7b, 0x49, 0xee, 0xc6, 0x6c, 0x84, 0x03, 0x72,
	0x4a, 0x70, 0x08, 0x1f, 0x03, 0xb8, 0xaa, 0x73, 0xdc, 0x4e, 0xf3, 0xa5, 0xa7, 0x2a, 0xfa, 0xd6,
	0xf9, 0x65, 0x45, 0x5d, 0x4a, 0x9c, 0x80, 0x93, 0x09, 0xbe, 0xcb, 0xf6, 0xbd, 0x13, 0xa7, 0xd9,
	0x50, 0xf3, 0x77, 0xd9, 0x3e, 0x1e, 0x21, 0x12, 0xc2, 0x43, 0xb0, 0xbd, 0xca, 0x76, 0x9d, 0x63,
	0xd7, 0x6b, 0xb5, 0xbc, 0x86, 0xba, 0xa6, 0xef, 0x9c, 0x5f, 0x56, 0x36, 0x97, 0x02, 0x17, 0xc5,
	0x01, 0x8e, 0x22, 0x7c, 0x4f, 0xe3, 0x7b, 0x2f, 0xba, 0x5e, 0xbb, 0xe3, 0x35, 0xd4, 0xc2, 0x5d,
	0x8d, 0x8f, 0xdf, 0x8e, 0x31, 0xe3, 0x38, 0xd4, 0x0b, 0x1f, 0x3f, 0x1b, 0xb9, 0xba, 0x7b, 0x35,
	0x33, 0x94, 0xeb, 0x99, 0xa1, 0xfc, 0x9c, 0x19, 0xca, 0xc5, 0x8d, 0x91, 0xbb, 0xbe, 0x31, 0x72,
	0xdf, 0x6f, 0x8c, 0xdc, 0xab, 0xfd, 0x15, 0x67, 0xd2, 0xff, 0xfd, 0x41, 0xba, 0x91, 0x84, 0xc6,
	0x02, 0xd8, 0xef, 0xe5, 0xf7, 0x41, 0x18, 0xd4, 0x2f, 0x8a, 0xbd, 0x79, 0xf2, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x17, 0xc5, 0xa4, 0x39, 0x04, 0x00, 0x00,
}

func (m *Loan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Lender) > 0 {
		i -= len(m.Lender)
		copy(dAtA[i:], m.Lender)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Lender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Borrower) > 0 {
		i -= len(m.Borrower)
		copy(dAtA[i:], m.Borrower)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Borrower)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LoanEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoanEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoanEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.asset_id != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.asset_id))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLoan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLoan(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Lender) > 0 {
		i -= len(m.Lender)
		copy(dAtA[i:], m.Lender)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Lender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoan(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Loan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLoan(uint64(m.Id))
	}
	l = len(m.Borrower)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Lender)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	return n
}

func (m *LoanEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Lender)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovLoan(uint64(l))
	l = m.AmountIn.Size()
	n += 1 + l + sovLoan(uint64(l))
	if m.asset_id != 0 {
		n += 1 + sovLoan(uint64(m.asset_id))
	}
	return n
}

func sovLoan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoan(x uint64) (n int) {
	return sovLoan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Loan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoan
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
			return fmt.Errorf("proto: Loan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoan
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
func (m *LoanEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoan
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
			return fmt.Errorf("proto: LoanEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoanEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field asset_id", wireType)
			}
			m.asset_id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.asset_id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoan
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
func skipLoan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoan
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
					return 0, ErrIntOverflowLoan
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
					return 0, ErrIntOverflowLoan
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
				return 0, ErrInvalidLengthLoan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoan = fmt.Errorf("proto: unexpected end of group")
)