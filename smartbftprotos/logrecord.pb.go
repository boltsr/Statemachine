// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smartbftprotos/logrecord.proto

package smartbftprotos

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

type LogRecord_Type int32

const (
	LogRecord_ENTRY      LogRecord_Type = 0
	LogRecord_CONTROL    LogRecord_Type = 1
	LogRecord_CRC_ANCHOR LogRecord_Type = 2
)

var LogRecord_Type_name = map[int32]string{
	0: "ENTRY",
	1: "CONTROL",
	2: "CRC_ANCHOR",
}

var LogRecord_Type_value = map[string]int32{
	"ENTRY":      0,
	"CONTROL":    1,
	"CRC_ANCHOR": 2,
}

func (x LogRecord_Type) String() string {
	return proto.EnumName(LogRecord_Type_name, int32(x))
}

func (LogRecord_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6f88cf2e7d665db, []int{0, 0}
}

type LogRecord struct {
	Type                 LogRecord_Type `protobuf:"varint,1,opt,name=type,proto3,enum=smartbftprotos.LogRecord_Type" json:"type,omitempty"`
	TruncateTo           bool           `protobuf:"varint,2,opt,name=truncate_to,json=truncateTo,proto3" json:"truncate_to,omitempty"`
	Data                 []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f88cf2e7d665db, []int{0}
}

func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetType() LogRecord_Type {
	if m != nil {
		return m.Type
	}
	return LogRecord_ENTRY
}

func (m *LogRecord) GetTruncateTo() bool {
	if m != nil {
		return m.TruncateTo
	}
	return false
}

func (m *LogRecord) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("smartbftprotos.LogRecord_Type", LogRecord_Type_name, LogRecord_Type_value)
	proto.RegisterType((*LogRecord)(nil), "smartbftprotos.LogRecord")
}

func init() {
	proto.RegisterFile("smartbftprotos/logrecord.proto", fileDescriptor_d6f88cf2e7d665db)
}

var fileDescriptor_d6f88cf2e7d665db = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xce, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x2b, 0x29, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0xcf, 0xc9, 0x4f, 0x2f, 0x4a,
	0x4d, 0xce, 0x2f, 0x4a, 0xd1, 0x03, 0x0b, 0x08, 0xf1, 0xa1, 0xca, 0x2b, 0x2d, 0x61, 0xe4, 0xe2,
	0xf4, 0xc9, 0x4f, 0x0f, 0x02, 0xab, 0x11, 0x32, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd3, 0x43, 0x55, 0xac, 0x07, 0x57, 0xa8, 0x17, 0x52, 0x59,
	0x90, 0x1a, 0x04, 0x56, 0x2b, 0x24, 0xcf, 0xc5, 0x5d, 0x52, 0x54, 0x9a, 0x97, 0x9c, 0x58, 0x92,
	0x1a, 0x5f, 0x92, 0x2f, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x05, 0x13, 0x0a, 0xc9, 0x17,
	0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3,
	0x95, 0xf4, 0xb8, 0x58, 0x40, 0x46, 0x08, 0x71, 0x72, 0xb1, 0xba, 0xfa, 0x85, 0x04, 0x45, 0x0a,
	0x30, 0x08, 0x71, 0x73, 0xb1, 0x3b, 0xfb, 0xfb, 0x85, 0x04, 0xf9, 0xfb, 0x08, 0x30, 0x0a, 0xf1,
	0x71, 0x71, 0x39, 0x07, 0x39, 0xc7, 0x3b, 0xfa, 0x39, 0x7b, 0xf8, 0x07, 0x09, 0x30, 0x25, 0xb1,
	0x81, 0x5d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x40, 0xd3, 0x1a, 0xdf, 0x00, 0x00,
	0x00,
}
