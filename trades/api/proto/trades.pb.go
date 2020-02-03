// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/trades.proto

package trades

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Asset struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	BookCost             int64    `protobuf:"varint,3,opt,name=book_cost,json=bookCost,proto3" json:"book_cost,omitempty"`
	Trades               []*Trade `protobuf:"bytes,4,rep,name=trades,proto3" json:"trades,omitempty"`
	CurrentQuantity      int64    `protobuf:"varint,5,opt,name=current_quantity,json=currentQuantity,proto3" json:"current_quantity,omitempty"`
	CurrentValue         int64    `protobuf:"varint,6,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c087062f7b627349, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Asset) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Asset) GetBookCost() int64 {
	if m != nil {
		return m.BookCost
	}
	return 0
}

func (m *Asset) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

func (m *Asset) GetCurrentQuantity() int64 {
	if m != nil {
		return m.CurrentQuantity
	}
	return 0
}

func (m *Asset) GetCurrentValue() int64 {
	if m != nil {
		return m.CurrentValue
	}
	return 0
}

type Trade struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Asset                *Asset   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Quantity             int64    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	UnitPrice            int64    `protobuf:"varint,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	TargetPrice          int64    `protobuf:"varint,6,opt,name=target_price,json=targetPrice,proto3" json:"target_price,omitempty"`
	Notes                string   `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
	ClientUuid           string   `protobuf:"bytes,8,opt,name=client_uuid,json=clientUuid,proto3" json:"client_uuid,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_c087062f7b627349, []int{1}
}

func (m *Trade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trade.Unmarshal(m, b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
}
func (m *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(m, src)
}
func (m *Trade) XXX_Size() int {
	return xxx_messageInfo_Trade.Size(m)
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Trade) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Trade) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Trade) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Trade) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Trade) GetTargetPrice() int64 {
	if m != nil {
		return m.TargetPrice
	}
	return 0
}

func (m *Trade) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Trade) GetClientUuid() string {
	if m != nil {
		return m.ClientUuid
	}
	return ""
}

func (m *Trade) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Point struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Close                float32  `protobuf:"fixed32,2,opt,name=close,proto3" json:"close,omitempty"`
	Volume               float32  `protobuf:"fixed32,3,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_c087062f7b627349, []int{2}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Point) GetClose() float32 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Point) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*Asset)(nil), "Asset")
	proto.RegisterType((*Trade)(nil), "Trade")
	proto.RegisterType((*Point)(nil), "Point")
}

func init() { proto.RegisterFile("proto/trades.proto", fileDescriptor_c087062f7b627349) }

var fileDescriptor_c087062f7b627349 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x86, 0x49, 0x20, 0xd9, 0x64, 0x60, 0xb5, 0xc8, 0x5a, 0xad, 0x2c, 0x16, 0x76, 0x69, 0xda,
	0x03, 0xbd, 0x50, 0x89, 0x3e, 0x01, 0xe2, 0x50, 0xf5, 0x50, 0x89, 0xa6, 0xb4, 0xd7, 0xc8, 0x24,
	0x56, 0x15, 0x35, 0xc4, 0x34, 0x99, 0x20, 0x71, 0xec, 0x7b, 0xf5, 0xe1, 0x2a, 0x8f, 0x1d, 0x54,
	0xa9, 0x3d, 0xd9, 0xff, 0xe7, 0x91, 0xfd, 0xff, 0x33, 0x06, 0xb6, 0xaf, 0x14, 0xaa, 0x2b, 0xac,
	0x44, 0x26, 0xeb, 0x39, 0x89, 0xe8, 0xdd, 0x01, 0x6f, 0x59, 0xd7, 0x12, 0x19, 0x83, 0x5e, 0xd3,
	0xe4, 0x19, 0x77, 0xa6, 0xce, 0x2c, 0x8c, 0x69, 0xaf, 0x19, 0x1e, 0xf7, 0x92, 0xbb, 0x86, 0xe9,
	0x3d, 0xfb, 0x0b, 0xe1, 0x56, 0xa9, 0x97, 0x24, 0x55, 0x35, 0xf2, 0xee, 0xd4, 0x99, 0x75, 0xe3,
	0x40, 0x83, 0x95, 0xaa, 0x91, 0xfd, 0x03, 0xdf, 0x5c, 0xcf, 0x7b, 0xd3, 0xee, 0xac, 0xbf, 0xf0,
	0xe7, 0x1b, 0x2d, 0x63, 0x4b, 0xd9, 0x25, 0x0c, 0xd3, 0xa6, 0xaa, 0x64, 0x89, 0xc9, 0x6b, 0x23,
	0x4a, 0xcc, 0xf1, 0xc8, 0x3d, 0xba, 0xe3, 0x97, 0xe5, 0xf7, 0x16, 0xb3, 0x73, 0xf8, 0xd9, 0x96,
	0x1e, 0x44, 0xd1, 0x48, 0xee, 0x53, 0xdd, 0xc0, 0xc2, 0x27, 0xcd, 0xa2, 0x37, 0x17, 0x3c, 0x7a,
	0xe1, 0x5b, 0xfb, 0x63, 0xf0, 0x84, 0xce, 0x46, 0xfe, 0xb5, 0x19, 0x4a, 0x1a, 0x1b, 0xc8, 0x46,
	0x10, 0x9c, 0x3c, 0xd8, 0x1c, 0xad, 0x3e, 0x05, 0xef, 0x7d, 0x0a, 0x3e, 0x01, 0x68, 0xca, 0x1c,
	0x93, 0x7d, 0x95, 0xa7, 0xd2, 0xba, 0x0e, 0x35, 0x59, 0x6b, 0xc0, 0xce, 0x60, 0x80, 0xa2, 0x7a,
	0x96, 0x6d, 0x81, 0xb1, 0xdb, 0x37, 0xcc, 0x94, 0xfc, 0x06, 0xaf, 0x54, 0x28, 0x6b, 0xfe, 0x83,
	0xae, 0x35, 0x82, 0xfd, 0x87, 0x7e, 0x5a, 0xe4, 0x3a, 0x27, 0x05, 0x08, 0xe8, 0x0c, 0x0c, 0x7a,
	0xd4, 0x31, 0x26, 0x00, 0x69, 0x25, 0x05, 0xca, 0x2c, 0x11, 0xc8, 0x43, 0x3a, 0x0f, 0x2d, 0x59,
	0x62, 0x74, 0x0b, 0xde, 0x5a, 0xe5, 0x25, 0x4d, 0x30, 0x13, 0x28, 0xdb, 0x16, 0xe8, 0xbd, 0x7e,
	0x32, 0x2d, 0x54, 0x6d, 0x46, 0xe8, 0xc6, 0x46, 0xb0, 0x3f, 0xe0, 0x1f, 0x54, 0xd1, 0xec, 0x24,
	0x05, 0x77, 0x63, 0xab, 0x16, 0x3b, 0xf0, 0x37, 0x66, 0x50, 0x63, 0x08, 0x6e, 0x24, 0x9a, 0x9f,
	0x61, 0xfb, 0x36, 0xb2, 0x6b, 0xd4, 0xd1, 0x96, 0x57, 0xf4, 0xbe, 0xe9, 0xbd, 0x9d, 0xf2, 0xc8,
	0xae, 0x51, 0x87, 0x5d, 0xc0, 0xf0, 0x41, 0x22, 0xa9, 0x3b, 0x89, 0x22, 0x13, 0x28, 0xbe, 0x56,
	0x6d, 0x7d, 0xfa, 0x83, 0xd7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0x74, 0x59, 0x9f, 0x99,
	0x02, 0x00, 0x00,
}