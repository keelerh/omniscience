// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gdrive.proto

package omniscience

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetAllDocumentsRequest struct {
	ModifiedSince *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=modified_since,json=modifiedSince" json:"modified_since,omitempty"`
}

func (m *GetAllDocumentsRequest) Reset()                    { *m = GetAllDocumentsRequest{} }
func (m *GetAllDocumentsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllDocumentsRequest) ProtoMessage()               {}
func (*GetAllDocumentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetAllDocumentsRequest) GetModifiedSince() *google_protobuf.Timestamp {
	if m != nil {
		return m.ModifiedSince
	}
	return nil
}

type GetAllDocumentsResponse struct {
	Documents []*Document `protobuf:"bytes,1,rep,name=documents" json:"documents,omitempty"`
}

func (m *GetAllDocumentsResponse) Reset()                    { *m = GetAllDocumentsResponse{} }
func (m *GetAllDocumentsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllDocumentsResponse) ProtoMessage()               {}
func (*GetAllDocumentsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetAllDocumentsResponse) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllDocumentsRequest)(nil), "omniscience_server.GetAllDocumentsRequest")
	proto.RegisterType((*GetAllDocumentsResponse)(nil), "omniscience_server.GetAllDocumentsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoogleDrive service

type GoogleDriveClient interface {
	// Gets all of a user's Google Drive documents.
	GetAll(ctx context.Context, in *GetAllDocumentsRequest, opts ...grpc.CallOption) (*GetAllDocumentsResponse, error)
}

type googleDriveClient struct {
	cc *grpc.ClientConn
}

func NewGoogleDriveClient(cc *grpc.ClientConn) GoogleDriveClient {
	return &googleDriveClient{cc}
}

func (c *googleDriveClient) GetAll(ctx context.Context, in *GetAllDocumentsRequest, opts ...grpc.CallOption) (*GetAllDocumentsResponse, error) {
	out := new(GetAllDocumentsResponse)
	err := grpc.Invoke(ctx, "/omniscience_server.GoogleDrive/GetAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoogleDrive service

type GoogleDriveServer interface {
	// Gets all of a user's Google Drive documents.
	GetAll(context.Context, *GetAllDocumentsRequest) (*GetAllDocumentsResponse, error)
}

func RegisterGoogleDriveServer(s *grpc.Server, srv GoogleDriveServer) {
	s.RegisterService(&_GoogleDrive_serviceDesc, srv)
}

func _GoogleDrive_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleDriveServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omniscience_server.GoogleDrive/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleDriveServer).GetAll(ctx, req.(*GetAllDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoogleDrive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omniscience_server.GoogleDrive",
	HandlerType: (*GoogleDriveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _GoogleDrive_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gdrive.proto",
}

func init() { proto.RegisterFile("gdrive.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x2d, 0xc2, 0x82, 0x09, 0xee, 0x21, 0xa0, 0x2e, 0xb9, 0xb8, 0x54, 0x0f, 0x7b, 0xca,
	0x42, 0xf7, 0x17, 0x2c, 0x2c, 0xf4, 0xe6, 0xa1, 0xea, 0xc9, 0x43, 0xb1, 0xc9, 0xb4, 0x04, 0x9a,
	0x4c, 0xed, 0xa4, 0xfe, 0x7e, 0x69, 0x63, 0x50, 0x51, 0x3c, 0x86, 0xbc, 0xef, 0x9b, 0xf7, 0x98,
	0xe8, 0x10, 0xbb, 0x1e, 0x6a, 0x33, 0xda, 0x77, 0x50, 0xc3, 0x88, 0x01, 0x05, 0x47, 0xe7, 0x2d,
	0x69, 0x0b, 0x5e, 0x83, 0x5c, 0x1b, 0xd4, 0x93, 0x03, 0x1f, 0xe2, 0xa7, 0xbc, 0x8d, 0xc0, 0x7e,
	0x79, 0x35, 0x53, 0xbb, 0x0f, 0xd6, 0x01, 0x85, 0x57, 0x37, 0xc4, 0x40, 0xfe, 0xc2, 0xae, 0x4b,
	0x08, 0xc7, 0xbe, 0x3f, 0x7d, 0x82, 0x54, 0xc1, 0xdb, 0x04, 0x14, 0xc4, 0x91, 0xad, 0x1d, 0x1a,
	0xdb, 0x5a, 0x30, 0x35, 0x59, 0xaf, 0x61, 0x93, 0x6d, 0xb3, 0x1d, 0x2f, 0xa4, 0x8a, 0x4e, 0x95,
	0x9c, 0xea, 0x29, 0x39, 0xab, 0xcb, 0x44, 0x3c, 0xce, 0x40, 0xfe, 0xc0, 0x6e, 0x7e, 0xc9, 0x69,
	0x40, 0x4f, 0x20, 0x0e, 0xec, 0x22, 0x55, 0xa5, 0x4d, 0xb6, 0x3d, 0xdf, 0xf1, 0xe2, 0x4a, 0x7d,
	0x5b, 0xa2, 0x12, 0x52, 0x7d, 0xe5, 0x0a, 0xc3, 0x78, 0xb9, 0xdc, 0x3e, 0xcd, 0xfb, 0xc5, 0x33,
	0x5b, 0x45, 0xbd, 0xb8, 0xfb, 0x81, 0xfe, 0x3d, 0x48, 0xde, 0xff, 0x1f, 0x8a, 0xc5, 0xf2, 0xb3,
	0x66, 0xb5, 0x0c, 0x3b, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0xfa, 0x9c, 0xa3, 0x6d, 0x01,
	0x00, 0x00,
}