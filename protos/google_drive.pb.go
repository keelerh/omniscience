// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google_drive.proto

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

func init() {
	proto.RegisterType((*GetAllDocumentsRequest)(nil), "omniscience.GetAllDocumentsRequest")
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
	GetAll(ctx context.Context, in *GetAllDocumentsRequest, opts ...grpc.CallOption) (GoogleDrive_GetAllClient, error)
}

type googleDriveClient struct {
	cc *grpc.ClientConn
}

func NewGoogleDriveClient(cc *grpc.ClientConn) GoogleDriveClient {
	return &googleDriveClient{cc}
}

func (c *googleDriveClient) GetAll(ctx context.Context, in *GetAllDocumentsRequest, opts ...grpc.CallOption) (GoogleDrive_GetAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleDrive_serviceDesc.Streams[0], c.cc, "/omniscience.GoogleDrive/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleDriveGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleDrive_GetAllClient interface {
	Recv() (*Document, error)
	grpc.ClientStream
}

type googleDriveGetAllClient struct {
	grpc.ClientStream
}

func (x *googleDriveGetAllClient) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GoogleDrive service

type GoogleDriveServer interface {
	// Gets all of a user's Google Drive documents.
	GetAll(*GetAllDocumentsRequest, GoogleDrive_GetAllServer) error
}

func RegisterGoogleDriveServer(s *grpc.Server, srv GoogleDriveServer) {
	s.RegisterService(&_GoogleDrive_serviceDesc, srv)
}

func _GoogleDrive_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllDocumentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleDriveServer).GetAll(m, &googleDriveGetAllServer{stream})
}

type GoogleDrive_GetAllServer interface {
	Send(*Document) error
	grpc.ServerStream
}

type googleDriveGetAllServer struct {
	grpc.ServerStream
}

func (x *googleDriveGetAllServer) Send(m *Document) error {
	return x.ServerStream.SendMsg(m)
}

var _GoogleDrive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omniscience.GoogleDrive",
	HandlerType: (*GoogleDriveServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _GoogleDrive_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google_drive.proto",
}

func init() { proto.RegisterFile("google_drive.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0x8d, 0x4f, 0x29, 0xca, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xce, 0xcf, 0xcd, 0xcb, 0x2c, 0x4e, 0xce, 0x4c, 0xcd, 0x4b, 0x4e, 0x95, 0xe2, 0x4b, 0xc9, 0x4f,
	0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x81, 0x48, 0x4a, 0xc9, 0x43, 0x34, 0xe8, 0x83, 0x79, 0x49, 0xa5,
	0x69, 0xfa, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0x05, 0x4a, 0xd1, 0x5c,
	0x62, 0xee, 0xa9, 0x25, 0x8e, 0x39, 0x39, 0x2e, 0x50, 0x8d, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x8e, 0x5c, 0x7c, 0xb9, 0xf9, 0x29, 0x99, 0x69, 0x99, 0xa9, 0x29, 0xf1, 0xc5,
	0x99, 0x79, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0x33, 0xf5,
	0x60, 0x66, 0xea, 0x85, 0xc0, 0xcc, 0x0c, 0xe2, 0x85, 0xe9, 0x08, 0x06, 0x69, 0x30, 0x0a, 0xe7,
	0xe2, 0x76, 0x07, 0xab, 0x75, 0x01, 0xb9, 0x57, 0xc8, 0x83, 0x8b, 0x0d, 0x62, 0x97, 0x90, 0xb2,
	0x1e, 0x92, 0xa3, 0xf5, 0xb0, 0x3b, 0x40, 0x4a, 0x14, 0x45, 0x11, 0x4c, 0x5a, 0x89, 0xc1, 0x80,
	0x31, 0x89, 0x0d, 0x6c, 0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xee, 0xdc, 0xf8, 0xf2, 0x10,
	0x01, 0x00, 0x00,
}
