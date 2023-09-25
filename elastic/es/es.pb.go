// Code generated by protoc-gen-go. DO NOT EDIT.
// source: es.proto

package es

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FieldZeroValueQuery struct {
	QueryZero            string   `protobuf:"bytes,1,opt,name=query_zero,json=queryZero,proto3" json:"query_zero,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldZeroValueQuery) Reset()         { *m = FieldZeroValueQuery{} }
func (m *FieldZeroValueQuery) String() string { return proto.CompactTextString(m) }
func (*FieldZeroValueQuery) ProtoMessage()    {}
func (*FieldZeroValueQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_718db5c20d0f3738, []int{0}
}

func (m *FieldZeroValueQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldZeroValueQuery.Unmarshal(m, b)
}
func (m *FieldZeroValueQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldZeroValueQuery.Marshal(b, m, deterministic)
}
func (m *FieldZeroValueQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldZeroValueQuery.Merge(m, src)
}
func (m *FieldZeroValueQuery) XXX_Size() int {
	return xxx_messageInfo_FieldZeroValueQuery.Size(m)
}
func (m *FieldZeroValueQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldZeroValueQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FieldZeroValueQuery proto.InternalMessageInfo

func (m *FieldZeroValueQuery) GetQueryZero() string {
	if m != nil {
		return m.QueryZero
	}
	return ""
}

type FieldEs struct {
	Es                   string   `protobuf:"bytes,1,opt,name=es,proto3" json:"es,omitempty"`
	StoreZeroValue       string   `protobuf:"bytes,2,opt,name=store_zero_value,json=storeZeroValue,proto3" json:"store_zero_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldEs) Reset()         { *m = FieldEs{} }
func (m *FieldEs) String() string { return proto.CompactTextString(m) }
func (*FieldEs) ProtoMessage()    {}
func (*FieldEs) Descriptor() ([]byte, []int) {
	return fileDescriptor_718db5c20d0f3738, []int{1}
}

func (m *FieldEs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldEs.Unmarshal(m, b)
}
func (m *FieldEs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldEs.Marshal(b, m, deterministic)
}
func (m *FieldEs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldEs.Merge(m, src)
}
func (m *FieldEs) XXX_Size() int {
	return xxx_messageInfo_FieldEs.Size(m)
}
func (m *FieldEs) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldEs.DiscardUnknown(m)
}

var xxx_messageInfo_FieldEs proto.InternalMessageInfo

func (m *FieldEs) GetEs() string {
	if m != nil {
		return m.Es
	}
	return ""
}

func (m *FieldEs) GetStoreZeroValue() string {
	if m != nil {
		return m.StoreZeroValue
	}
	return ""
}

type FieldQuery struct {
	MatchPhrasePrefix      string   `protobuf:"bytes,1,opt,name=match_phrase_prefix,json=matchPhrasePrefix,proto3" json:"match_phrase_prefix,omitempty"`
	MatchPhrasePrefixLeft  string   `protobuf:"bytes,2,opt,name=match_phrase_prefix_left,json=matchPhrasePrefixLeft,proto3" json:"match_phrase_prefix_left,omitempty"`
	MatchPhrasePrefixRight string   `protobuf:"bytes,3,opt,name=match_phrase_prefix_right,json=matchPhrasePrefixRight,proto3" json:"match_phrase_prefix_right,omitempty"`
	Wildcard               string   `protobuf:"bytes,4,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
	WildcardLeft           string   `protobuf:"bytes,5,opt,name=wildcard_left,json=wildcardLeft,proto3" json:"wildcard_left,omitempty"`
	WildcardRight          string   `protobuf:"bytes,6,opt,name=wildcard_right,json=wildcardRight,proto3" json:"wildcard_right,omitempty"`
	Terms                  string   `protobuf:"bytes,7,opt,name=terms,proto3" json:"terms,omitempty"`
	Match                  string   `protobuf:"bytes,8,opt,name=match,proto3" json:"match,omitempty"`
	MatchPrefix            string   `protobuf:"bytes,9,opt,name=match_prefix,json=matchPrefix,proto3" json:"match_prefix,omitempty"`
	NotTerms               string   `protobuf:"bytes,10,opt,name=not_terms,json=notTerms,proto3" json:"not_terms,omitempty"`
	Gte                    string   `protobuf:"bytes,11,opt,name=gte,proto3" json:"gte,omitempty"`
	Gt                     string   `protobuf:"bytes,12,opt,name=gt,proto3" json:"gt,omitempty"`
	Lte                    string   `protobuf:"bytes,13,opt,name=lte,proto3" json:"lte,omitempty"`
	Lt                     string   `protobuf:"bytes,14,opt,name=lt,proto3" json:"lt,omitempty"`
	TermsZero              string   `protobuf:"bytes,15,opt,name=terms_zero,json=termsZero,proto3" json:"terms_zero,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *FieldQuery) Reset()         { *m = FieldQuery{} }
func (m *FieldQuery) String() string { return proto.CompactTextString(m) }
func (*FieldQuery) ProtoMessage()    {}
func (*FieldQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_718db5c20d0f3738, []int{2}
}

func (m *FieldQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldQuery.Unmarshal(m, b)
}
func (m *FieldQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldQuery.Marshal(b, m, deterministic)
}
func (m *FieldQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldQuery.Merge(m, src)
}
func (m *FieldQuery) XXX_Size() int {
	return xxx_messageInfo_FieldQuery.Size(m)
}
func (m *FieldQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FieldQuery proto.InternalMessageInfo

func (m *FieldQuery) GetMatchPhrasePrefix() string {
	if m != nil {
		return m.MatchPhrasePrefix
	}
	return ""
}

func (m *FieldQuery) GetMatchPhrasePrefixLeft() string {
	if m != nil {
		return m.MatchPhrasePrefixLeft
	}
	return ""
}

func (m *FieldQuery) GetMatchPhrasePrefixRight() string {
	if m != nil {
		return m.MatchPhrasePrefixRight
	}
	return ""
}

func (m *FieldQuery) GetWildcard() string {
	if m != nil {
		return m.Wildcard
	}
	return ""
}

func (m *FieldQuery) GetWildcardLeft() string {
	if m != nil {
		return m.WildcardLeft
	}
	return ""
}

func (m *FieldQuery) GetWildcardRight() string {
	if m != nil {
		return m.WildcardRight
	}
	return ""
}

func (m *FieldQuery) GetTerms() string {
	if m != nil {
		return m.Terms
	}
	return ""
}

func (m *FieldQuery) GetMatch() string {
	if m != nil {
		return m.Match
	}
	return ""
}

func (m *FieldQuery) GetMatchPrefix() string {
	if m != nil {
		return m.MatchPrefix
	}
	return ""
}

func (m *FieldQuery) GetNotTerms() string {
	if m != nil {
		return m.NotTerms
	}
	return ""
}

func (m *FieldQuery) GetGte() string {
	if m != nil {
		return m.Gte
	}
	return ""
}

func (m *FieldQuery) GetGt() string {
	if m != nil {
		return m.Gt
	}
	return ""
}

func (m *FieldQuery) GetLte() string {
	if m != nil {
		return m.Lte
	}
	return ""
}

func (m *FieldQuery) GetLt() string {
	if m != nil {
		return m.Lt
	}
	return ""
}

func (m *FieldQuery) GetTermsZero() string {
	if m != nil {
		return m.TermsZero
	}
	return ""
}

var E_Es = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*FieldEs)(nil),
	Field:         66664,
	Name:          "es_proto.es",
	Tag:           "bytes,66664,opt,name=es",
	Filename:      "es.proto",
}

var E_Query = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*FieldQuery)(nil),
	Field:         66665,
	Name:          "es_proto.query",
	Tag:           "bytes,66665,opt,name=query",
	Filename:      "es.proto",
}

var E_QueryZero = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*FieldZeroValueQuery)(nil),
	Field:         66666,
	Name:          "es_proto.query_zero",
	Tag:           "bytes,66666,opt,name=query_zero",
	Filename:      "es.proto",
}

func init() {
	proto.RegisterType((*FieldZeroValueQuery)(nil), "es_proto.FieldZeroValueQuery")
	proto.RegisterType((*FieldEs)(nil), "es_proto.FieldEs")
	proto.RegisterType((*FieldQuery)(nil), "es_proto.FieldQuery")
	proto.RegisterExtension(E_Es)
	proto.RegisterExtension(E_Query)
	proto.RegisterExtension(E_QueryZero)
}

func init() { proto.RegisterFile("es.proto", fileDescriptor_718db5c20d0f3738) }

var fileDescriptor_718db5c20d0f3738 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xa5, 0x89, 0x69, 0x93, 0x9b, 0x34, 0xb6, 0xd3, 0x2a, 0x63, 0x25, 0x50, 0x23, 0x42, 0x1f,
	0x74, 0x17, 0xaa, 0x20, 0xc6, 0xb7, 0x8a, 0xbe, 0x28, 0x58, 0x83, 0xf8, 0x50, 0x90, 0x25, 0xd9,
	0xbd, 0xd9, 0x2c, 0x4c, 0x32, 0xdb, 0x99, 0x9b, 0x58, 0xfd, 0x05, 0xfe, 0x53, 0x3f, 0x7e, 0x85,
	0xcc, 0xbd, 0x9b, 0x50, 0xdb, 0xa2, 0x6f, 0xf7, 0xe3, 0x9c, 0x73, 0x67, 0xe6, 0x9e, 0x81, 0x26,
	0xfa, 0xa8, 0x74, 0x96, 0xac, 0x6a, 0xa2, 0x4f, 0x38, 0x3a, 0x38, 0xcc, 0xad, 0xcd, 0x0d, 0xc6,
	0x9c, 0x8d, 0x17, 0x93, 0x38, 0x43, 0x9f, 0xba, 0xa2, 0x24, 0xeb, 0x04, 0xdb, 0x7f, 0x06, 0x7b,
	0x6f, 0x0a, 0x34, 0xd9, 0x19, 0x3a, 0xfb, 0x69, 0x64, 0x16, 0xf8, 0x61, 0x81, 0xee, 0xab, 0xea,
	0x01, 0x9c, 0x87, 0x20, 0xf9, 0x86, 0xce, 0xea, 0x8d, 0xc3, 0x8d, 0xa3, 0xd6, 0xb0, 0xc5, 0x95,
	0x00, 0xec, 0xbf, 0x82, 0x2d, 0x66, 0xbd, 0xf6, 0xaa, 0x0b, 0x35, 0xf4, 0x15, 0xa2, 0x86, 0x5e,
	0x1d, 0xc1, 0x8e, 0x27, 0xeb, 0x90, 0x99, 0xc9, 0x32, 0x48, 0xea, 0x1a, 0x77, 0xbb, 0x5c, 0x5f,
	0x0f, 0xea, 0xff, 0xa8, 0x03, 0xb0, 0x8a, 0x8c, 0x8c, 0x60, 0x6f, 0x36, 0xa2, 0x74, 0x9a, 0x94,
	0x53, 0x37, 0xf2, 0x98, 0x94, 0x0e, 0x27, 0xc5, 0x45, 0xa5, 0xbc, 0xcb, 0xad, 0x53, 0xee, 0x9c,
	0x72, 0x43, 0x3d, 0x07, 0x7d, 0x03, 0x3e, 0x31, 0x38, 0xa1, 0x6a, 0xe0, 0x9d, 0x6b, 0xa4, 0x77,
	0x38, 0x21, 0xf5, 0x02, 0xee, 0xdd, 0x44, 0x74, 0x45, 0x3e, 0x25, 0x5d, 0x67, 0xe6, 0xdd, 0x6b,
	0xcc, 0x61, 0xe8, 0xaa, 0x03, 0x68, 0x7e, 0x29, 0x4c, 0x96, 0x8e, 0x5c, 0xa6, 0x6f, 0x31, 0x72,
	0x9d, 0xab, 0x87, 0xb0, 0xbd, 0x8a, 0xe5, 0x10, 0x0d, 0x06, 0x74, 0x56, 0x45, 0x9e, 0xfd, 0x08,
	0xba, 0x6b, 0x90, 0x0c, 0xdc, 0x64, 0xd4, 0x9a, 0x2a, 0x73, 0xf6, 0xa1, 0x41, 0xe8, 0x66, 0x5e,
	0x6f, 0x71, 0x57, 0x92, 0x50, 0xe5, 0x73, 0xe9, 0xa6, 0x54, 0x39, 0x51, 0x0f, 0xa0, 0x53, 0x5d,
	0x47, 0x1e, 0xac, 0xc5, 0xcd, 0xb6, 0xdc, 0x40, 0x9e, 0xea, 0x3e, 0xb4, 0xe6, 0x96, 0x12, 0x91,
	0x04, 0x39, 0xf7, 0xdc, 0xd2, 0x47, 0x56, 0xdd, 0x81, 0x7a, 0x4e, 0xa8, 0xdb, 0x5c, 0x0e, 0x61,
	0x58, 0x69, 0x4e, 0xba, 0x23, 0x2b, 0xcd, 0x29, 0x20, 0x0c, 0xa1, 0xde, 0x16, 0x84, 0x11, 0x84,
	0x21, 0xdd, 0x15, 0x84, 0xa1, 0x60, 0x17, 0x16, 0x17, 0xbb, 0xdc, 0x16, 0xbb, 0x70, 0x25, 0xac,
	0x7b, 0x70, 0x12, 0x3c, 0xa2, 0x7a, 0x91, 0xb8, 0x31, 0x5a, 0xb9, 0x31, 0xe2, 0xed, 0xbf, 0x2f,
	0xa9, 0xb0, 0x73, 0xaf, 0x7f, 0x7e, 0x0f, 0x2f, 0xda, 0x3e, 0xde, 0x8d, 0x56, 0xf6, 0x8d, 0x2a,
	0x8f, 0x05, 0x5f, 0x0d, 0xde, 0x42, 0xe3, 0x5c, 0xac, 0xf9, 0x6f, 0x99, 0x5f, 0x95, 0xcc, 0xfe,
	0x15, 0x19, 0x36, 0xd9, 0x50, 0x34, 0x06, 0x9f, 0x2f, 0xdb, 0xfb, 0x7f, 0x8a, 0xbf, 0x2b, 0xc5,
	0xde, 0x15, 0xc5, 0xbf, 0xbf, 0xcc, 0xa5, 0xef, 0x71, 0x12, 0x9d, 0x3d, 0xce, 0x0b, 0x9a, 0x2e,
	0xc6, 0x51, 0x6a, 0x67, 0x71, 0x86, 0x4b, 0xbc, 0x28, 0x7d, 0x9c, 0xdb, 0x27, 0xb3, 0x22, 0x75,
	0x36, 0x5e, 0x1e, 0xc7, 0x68, 0x46, 0x9e, 0x8a, 0x34, 0x46, 0xff, 0x12, 0xfd, 0x78, 0x93, 0x55,
	0x9f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x81, 0x40, 0x8e, 0xc3, 0x03, 0x00, 0x00,
}
