// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_postAvatar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// web->srv
type Request struct {
	// sessionId
	SessionId string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	// 文件大小
	FileSize int64 `protobuf:"varint,2,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	// 文件名
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	// 二进制图片
	Buffer               []byte   `protobuf:"bytes,4,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Request) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *Request) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Request) GetBuffer() []byte {
	if m != nil {
		return m.Buffer
	}
	return nil
}

// srv->web
type Response struct {
	// 错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	// 错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	// fileId
	FileId               string   `protobuf:"bytes,3,opt,name=fileId,proto3" json:"fileId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.postAvatar.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.postAvatar.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x09, 0x2d, 0x6d, 0x73, 0x62, 0xb2, 0x10, 0x44, 0x85, 0x21, 0x64, 0xca, 0x64, 0x24,
	0xf8, 0x05, 0x08, 0x31, 0x74, 0x00, 0x21, 0xb3, 0xb0, 0xba, 0xf4, 0x52, 0x45, 0x4a, 0x72, 0xe6,
	0xce, 0x2d, 0x88, 0x5f, 0x8f, 0xe2, 0xb8, 0x61, 0x82, 0xc9, 0xfa, 0x9e, 0xfd, 0xfc, 0xee, 0x1d,
	0x5c, 0x3a, 0x26, 0x4f, 0x37, 0xf8, 0x65, 0x5b, 0xd7, 0xe0, 0xe1, 0xd4, 0x41, 0x55, 0x17, 0x5b,
	0xd2, 0x6d, 0xfd, 0xce, 0xa4, 0x85, 0xf7, 0xda, 0x91, 0xf8, 0xfb, 0xbd, 0xf5, 0x96, 0x8b, 0x4f,
	0x98, 0x1b, 0xfc, 0xd8, 0xa1, 0x78, 0x75, 0x05, 0xa9, 0xa0, 0x48, 0x4d, 0xdd, 0x6a, 0x93, 0x25,
	0x79, 0x52, 0xa6, 0xe6, 0x57, 0x50, 0x4b, 0x58, 0x54, 0x75, 0x83, 0xaf, 0xf5, 0x37, 0x66, 0xc7,
	0x79, 0x52, 0x4e, 0xcc, 0xc8, 0x87, 0xbb, 0x67, 0xdb, 0x62, 0x36, 0x09, 0xc6, 0x91, 0xd5, 0x39,
	0xcc, 0xd6, 0xbb, 0xaa, 0x42, 0xce, 0xa6, 0x79, 0x52, 0x9e, 0x9a, 0x48, 0xc5, 0x0b, 0x2c, 0x0c,
	0x8a, 0xa3, 0x4e, 0x50, 0x9d, 0xc1, 0x09, 0x32, 0x77, 0x14, 0x53, 0x07, 0xe8, 0x9d, 0xc8, 0xdc,
	0xca, 0x36, 0xe4, 0xa5, 0x26, 0x52, 0xaf, 0xf7, 0xbf, 0xaf, 0x36, 0x31, 0x2b, 0xd2, 0xed, 0x1b,
	0xcc, 0x1f, 0x87, 0xd2, 0xea, 0x09, 0xa6, 0x0f, 0xb6, 0x69, 0x54, 0xae, 0xff, 0xe8, 0xad, 0x63,
	0xe9, 0xe5, 0xf5, 0x3f, 0x2f, 0x86, 0xe9, 0x8a, 0xa3, 0xf5, 0x2c, 0x2c, 0xf1, 0xee, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x2a, 0x77, 0x4d, 0x63, 0x01, 0x00, 0x00,
}