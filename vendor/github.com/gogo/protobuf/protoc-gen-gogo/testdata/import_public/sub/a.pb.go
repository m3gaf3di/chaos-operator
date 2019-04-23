// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type E int32

const (
	E_ZERO E = 0
)

var E_name = map[int32]string{
	0: "ZERO",
}

var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}

func (x *E) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_value, data, "E")
	if err != nil {
		return err
	}
	*x = E(value)
	return nil
}

func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	p := new(M_Subenum)
	*p = x
	return p
}

func (x M_Subenum) String() string {
	return proto.EnumName(M_Subenum_name, int32(x))
}

func (x *M_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Subenum_value, data, "M_Subenum")
	if err != nil {
		return err
	}
	*x = M_Subenum(value)
	return nil
}

func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	p := new(M_Submessage_Submessage_Subenum)
	*p = x
	return p
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return proto.EnumName(M_Submessage_Submessage_Subenum_name, int32(x))
}

func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Submessage_Submessage_Subenum_value, data, "M_Submessage_Submessage_Subenum")
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(value)
	return nil
}

func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2 `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField           isM_OneofField `protobuf_oneof:"oneof_field"`
	Grouping             *M_Grouping    `protobuf:"group,4,opt,name=Grouping,json=grouping" json:"grouping,omitempty"`
	DefaultField         *string        `protobuf:"bytes,6,opt,name=default_field,json=defaultField,def=def" json:"default_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

const Default_M_DefaultField string = "def"

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof"`
}
type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof"`
}

func (*M_OneofInt32) isM_OneofField() {}
func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

func (m *M) GetGrouping() *M_Grouping {
	if m != nil {
		return m.Grouping
	}
	return nil
}

func (m *M) GetDefaultField() string {
	if m != nil && m.DefaultField != nil {
		return *m.DefaultField
	}
	return Default_M_DefaultField
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*M) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _M_OneofMarshaler, _M_OneofUnmarshaler, _M_OneofSizer, []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

func _M_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*M)
	// oneof_field
	switch x := m.OneofField.(type) {
	case *M_OneofInt32:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.OneofInt32))
	case *M_OneofInt64:
		_ = b.EncodeVarint(3<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.OneofInt64))
	case nil:
	default:
		return fmt.Errorf("M.OneofField has unexpected type %T", x)
	}
	return nil
}

func _M_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*M)
	switch tag {
	case 2: // oneof_field.oneof_int32
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OneofField = &M_OneofInt32{int32(x)}
		return true, err
	case 3: // oneof_field.oneof_int64
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OneofField = &M_OneofInt64{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _M_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*M)
	// oneof_field
	switch x := m.OneofField.(type) {
	case *M_OneofInt32:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.OneofInt32))
	case *M_OneofInt64:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.OneofInt64))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type M_Grouping struct {
	GroupField           *string  `protobuf:"bytes,5,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M_Grouping) Reset()         { *m = M_Grouping{} }
func (m *M_Grouping) String() string { return proto.CompactTextString(m) }
func (*M_Grouping) ProtoMessage()    {}
func (*M_Grouping) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}
func (m *M_Grouping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Grouping.Unmarshal(m, b)
}
func (m *M_Grouping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Grouping.Marshal(b, m, deterministic)
}
func (m *M_Grouping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Grouping.Merge(m, src)
}
func (m *M_Grouping) XXX_Size() int {
	return xxx_messageInfo_M_Grouping.Size(m)
}
func (m *M_Grouping) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Grouping.DiscardUnknown(m)
}

var xxx_messageInfo_M_Grouping proto.InternalMessageInfo

func (m *M_Grouping) GetGroupField() string {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return ""
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 1}
}
func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof"`
}
type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}
func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _M_Submessage_OneofMarshaler, _M_Submessage_OneofUnmarshaler, _M_Submessage_OneofSizer, []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

func _M_Submessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*M_Submessage)
	// submessage_oneof_field
	switch x := m.SubmessageOneofField.(type) {
	case *M_Submessage_SubmessageOneofInt32:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.SubmessageOneofInt32))
	case *M_Submessage_SubmessageOneofInt64:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.SubmessageOneofInt64))
	case nil:
	default:
		return fmt.Errorf("M_Submessage.SubmessageOneofField has unexpected type %T", x)
	}
	return nil
}

func _M_Submessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*M_Submessage)
	switch tag {
	case 1: // submessage_oneof_field.submessage_oneof_int32
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.SubmessageOneofField = &M_Submessage_SubmessageOneofInt32{int32(x)}
		return true, err
	case 2: // submessage_oneof_field.submessage_oneof_int64
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.SubmessageOneofField = &M_Submessage_SubmessageOneofInt64{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _M_Submessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*M_Submessage)
	// submessage_oneof_field
	switch x := m.SubmessageOneofField.(type) {
	case *M_Submessage_SubmessageOneofInt32:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.SubmessageOneofInt32))
	case *M_Submessage_SubmessageOneofInt64:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.SubmessageOneofInt64))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_ExtensionField = &proto.ExtensionDesc{
	ExtendedType:  (*M2)(nil),
	ExtensionType: (*string)(nil),
	Field:         1,
	Name:          "goproto.test.import_public.sub.extension_field",
	Tag:           "bytes,1,opt,name=extension_field",
	Filename:      "import_public/sub/a.proto",
}

func init() {
	proto.RegisterEnum("goproto.test.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.test.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.test.import_public.sub.M")
	proto.RegisterType((*M_Grouping)(nil), "goproto.test.import_public.sub.M.Grouping")
	proto.RegisterType((*M_Submessage)(nil), "goproto.test.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
}

func init() { proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_382f7805394b5c4e) }

var fileDescriptor_382f7805394b5c4e = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0xb3, 0xc6, 0x49, 0x9d, 0x71, 0xdd, 0x3f, 0xab, 0xa6, 0xa2, 0x39, 0xb4, 0x94, 0x13,
	0x4a, 0x94, 0x45, 0xa2, 0x88, 0x43, 0x6e, 0x75, 0x45, 0xd2, 0xaa, 0x42, 0x91, 0x40, 0xbd, 0xe4,
	0x82, 0x58, 0xb3, 0x6c, 0x91, 0x0c, 0x8b, 0xcc, 0xae, 0xd4, 0x47, 0xe8, 0x7b, 0xf5, 0xc5, 0x2a,
	0x16, 0xb0, 0x63, 0xd5, 0x6d, 0x73, 0x5b, 0xe6, 0xfb, 0x7e, 0x33, 0x9a, 0x8f, 0x81, 0x37, 0x65,
	0xd5, 0x88, 0x8d, 0x4c, 0x1b, 0x45, 0xd7, 0xe5, 0xca, 0x6d, 0x15, 0x75, 0x33, 0xd2, 0x6c, 0x84,
	0x14, 0xf8, 0x2d, 0x17, 0xfa, 0x41, 0x24, 0x6b, 0x25, 0xd9, 0xf3, 0x91, 0x56, 0xd1, 0xf3, 0x03,
	0x28, 0xed, 0x51, 0xfb, 0xe7, 0x14, 0x50, 0x84, 0x3d, 0x98, 0x54, 0x9e, 0x89, 0x2c, 0xe4, 0xcc,
	0x3d, 0x9b, 0xfc, 0xbb, 0x1b, 0x89, 0xbc, 0x78, 0x52, 0x79, 0xf8, 0x3d, 0xcc, 0x45, 0xcd, 0x44,
	0x91, 0x96, 0xb5, 0xfc, 0xe0, 0x99, 0x13, 0x0b, 0x39, 0xc7, 0x9f, 0x8f, 0x62, 0xd0, 0xc5, 0x2f,
	0x5d, 0x6d, 0xcf, 0x12, 0xf8, 0xa6, 0x61, 0x21, 0xc7, 0x78, 0x68, 0x09, 0x7c, 0x7c, 0x03, 0x33,
	0xbe, 0x11, 0xaa, 0x29, 0x6b, 0x6e, 0x4e, 0x2d, 0xe4, 0x80, 0x77, 0xf1, 0xdf, 0xf9, 0xe4, 0x76,
	0x20, 0xe2, 0x2d, 0x8b, 0x1d, 0x58, 0xe4, 0xac, 0xc8, 0xd4, 0x5a, 0xa6, 0x45, 0xc9, 0xd6, 0xb9,
	0x79, 0x62, 0x21, 0xe7, 0xf4, 0xda, 0xc8, 0x59, 0x11, 0x3f, 0x1d, 0x94, 0x9b, 0x4e, 0x38, 0xbf,
	0x84, 0xd9, 0xc8, 0xe3, 0x77, 0x30, 0xd7, 0x1d, 0x06, 0xe6, 0xb8, 0x63, 0x62, 0xd0, 0xa5, 0xde,
	0xfc, 0x0b, 0x01, 0x24, 0x8a, 0x56, 0xac, 0x6d, 0x33, 0xce, 0x70, 0x00, 0xaf, 0xdb, 0xed, 0x57,
	0xfa, 0x70, 0x7d, 0x34, 0xac, 0xff, 0x6a, 0xa7, 0xdf, 0xed, 0x82, 0xf8, 0x0b, 0x17, 0xf8, 0x3a,
	0x36, 0xe3, 0x30, 0x17, 0xf8, 0xf6, 0x25, 0xe0, 0xdd, 0xf4, 0x34, 0x51, 0x94, 0xd5, 0xaa, 0xc2,
	0x67, 0xf0, 0x32, 0x4a, 0x93, 0x6f, 0xcb, 0x28, 0x4c, 0x92, 0x8f, 0xb7, 0x61, 0x7a, 0x1f, 0xc6,
	0x77, 0x2f, 0x8e, 0x96, 0xe6, 0x81, 0x21, 0x7a, 0x2f, 0xfb, 0x0c, 0x9e, 0x8c, 0x2c, 0xc0, 0x49,
	0x34, 0x02, 0x8b, 0xf1, 0xf7, 0x68, 0xd7, 0xc5, 0x02, 0x50, 0x88, 0x67, 0x30, 0xed, 0xd5, 0xeb,
	0xaf, 0xf0, 0x9c, 0xfd, 0x90, 0xac, 0x6e, 0x4b, 0x51, 0xf7, 0x0e, 0xfc, 0x88, 0xd3, 0xd0, 0x41,
	0x9c, 0xc6, 0xcf, 0xb6, 0xa8, 0xce, 0x71, 0x19, 0xde, 0x7f, 0xe2, 0xa5, 0xfc, 0xae, 0x28, 0x59,
	0x89, 0xca, 0xe5, 0x82, 0x0b, 0x57, 0x37, 0xa2, 0xaa, 0xe8, 0x1f, 0xab, 0x2b, 0xce, 0xea, 0x2b,
	0x2d, 0x74, 0xbd, 0xf3, 0x4c, 0x66, 0xee, 0x1f, 0x57, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4f,
	0x37, 0x1b, 0x7d, 0x04, 0x03, 0x00, 0x00,
}
