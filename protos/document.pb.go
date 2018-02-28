// Code generated by protoc-gen-go. DO NOT EDIT.
// source: document.proto

/*
Package omniscience is a generated protocol buffer package.

It is generated from these files:
	document.proto
	google_drive.proto
	ingestion.proto
	search.proto

It has these top-level messages:
	Document
	DocumentId
	GetAllDocumentsRequest
	IndexDocumentServiceRequest
	QueryDocumentsRequest
*/
package omniscience

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DocumentService int32

const (
	DocumentService_ALL        DocumentService = 0
	DocumentService_GDRIVE     DocumentService = 1
	DocumentService_JIRA       DocumentService = 2
	DocumentService_CONFLUENCE DocumentService = 3
	DocumentService_GITHUB     DocumentService = 4
)

var DocumentService_name = map[int32]string{
	0: "ALL",
	1: "GDRIVE",
	2: "JIRA",
	3: "CONFLUENCE",
	4: "GITHUB",
}
var DocumentService_value = map[string]int32{
	"ALL":        0,
	"GDRIVE":     1,
	"JIRA":       2,
	"CONFLUENCE": 3,
	"GITHUB":     4,
}

func (x DocumentService) String() string {
	return proto.EnumName(DocumentService_name, int32(x))
}
func (DocumentService) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Document struct {
	Id           *DocumentId                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description  string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Service      DocumentService            `protobuf:"varint,4,opt,name=service,enum=omniscience.DocumentService" json:"service,omitempty"`
	Content      string                     `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	Url          string                     `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	LastModified *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Document) GetId() *DocumentId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Document) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Document) GetService() DocumentService {
	if m != nil {
		return m.Service
	}
	return DocumentService_ALL
}

func (m *Document) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Document) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Document) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type DocumentId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DocumentId) Reset()                    { *m = DocumentId{} }
func (m *DocumentId) String() string            { return proto.CompactTextString(m) }
func (*DocumentId) ProtoMessage()               {}
func (*DocumentId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DocumentId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Document)(nil), "omniscience.Document")
	proto.RegisterType((*DocumentId)(nil), "omniscience.DocumentId")
	proto.RegisterEnum("omniscience.DocumentService", DocumentService_name, DocumentService_value)
}

func init() { proto.RegisterFile("document.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xfe, 0x79, 0x29, 0x0c, 0x5a, 0x9b, 0xb9, 0xb8, 0x21, 0x24, 0x36, 0x5c, 0x24,
	0x1e, 0x4a, 0x82, 0x89, 0x57, 0x83, 0x80, 0x5a, 0x53, 0x31, 0xa9, 0xe0, 0xd5, 0xc0, 0xee, 0x42,
	0x36, 0x61, 0x77, 0x49, 0xbb, 0xf8, 0x31, 0xfc, 0xcc, 0x86, 0x85, 0x2a, 0x31, 0xde, 0x66, 0x67,
	0x7e, 0xfb, 0xcc, 0x3c, 0x0f, 0x84, 0x4c, 0xd3, 0xad, 0xe4, 0xca, 0x24, 0x9b, 0x42, 0x1b, 0x8d,
	0x4d, 0x2d, 0x95, 0x28, 0xa9, 0xe0, 0x8a, 0xf2, 0xd6, 0xc5, 0x4a, 0xeb, 0xd5, 0x9a, 0xf7, 0xec,
	0x68, 0xb1, 0x5d, 0xf6, 0x8c, 0x90, 0xbc, 0x34, 0x73, 0xb9, 0xd9, 0xd3, 0x9d, 0x4f, 0x17, 0xea,
	0xa3, 0x83, 0x00, 0x5e, 0x82, 0x2b, 0x18, 0x71, 0x62, 0xa7, 0xdb, 0xec, 0x9f, 0x27, 0x47, 0x3a,
	0x49, 0x85, 0xa4, 0x2c, 0x77, 0x05, 0x43, 0x04, 0x5f, 0xcd, 0x25, 0x27, 0x6e, 0xec, 0x74, 0x1b,
	0xb9, 0xad, 0x31, 0x86, 0x26, 0xe3, 0x25, 0x2d, 0xc4, 0xc6, 0x08, 0xad, 0x88, 0x67, 0x47, 0xc7,
	0x2d, 0xbc, 0x81, 0xa0, 0xe4, 0xc5, 0x87, 0xa0, 0x9c, 0xf8, 0xb1, 0xd3, 0x0d, 0xfb, 0xed, 0x3f,
	0x77, 0xbc, 0xee, 0x99, 0xbc, 0x82, 0x91, 0x40, 0x40, 0xb5, 0x32, 0x5c, 0x19, 0xf2, 0xdf, 0xaa,
	0x56, 0x4f, 0x8c, 0xc0, 0xdb, 0x16, 0x6b, 0x52, 0xb3, 0xdd, 0x5d, 0x89, 0xb7, 0x70, 0xba, 0x9e,
	0x97, 0xe6, 0x5d, 0x6a, 0x26, 0x96, 0x82, 0x33, 0x12, 0x58, 0x37, 0xad, 0x64, 0x1f, 0x44, 0x52,
	0x05, 0x91, 0x4c, 0xab, 0x20, 0xf2, 0x93, 0xdd, 0x87, 0xe7, 0x03, 0xdf, 0x69, 0x03, 0xfc, 0x98,
	0xc5, 0xf0, 0x3b, 0x91, 0xc6, 0xce, 0xf8, 0x55, 0x06, 0x67, 0xbf, 0xce, 0xc4, 0x00, 0xbc, 0x41,
	0x96, 0x45, 0xff, 0x10, 0xa0, 0xf6, 0x30, 0xca, 0xd3, 0xb7, 0x71, 0xe4, 0x60, 0x1d, 0xfc, 0xa7,
	0x34, 0x1f, 0x44, 0x2e, 0x86, 0x00, 0xc3, 0x97, 0xc9, 0x7d, 0x36, 0x1b, 0x4f, 0x86, 0xe3, 0xc8,
	0xb3, 0x54, 0x3a, 0x7d, 0x9c, 0xdd, 0x45, 0xfe, 0xa2, 0x66, 0xaf, 0xb9, 0xfe, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x9a, 0x7f, 0x21, 0xc3, 0x01, 0x00, 0x00,
}
