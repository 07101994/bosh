// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpi.proto

package cpi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
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

type Request struct {
	Type               string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	StemcellApiVersion int32           `protobuf:"varint,2,opt,name=stemcell_api_version,json=stemcellApiVersion,proto3" json:"stemcell_api_version,omitempty"`
	DirectorUuid       string          `protobuf:"bytes,3,opt,name=director_uuid,json=directorUuid,proto3" json:"director_uuid,omitempty"`
	Properties         *_struct.Struct `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	// Types that are valid to be assigned to Arguments:
	//	*Request_CreateVmArguments
	Arguments            isRequest_Arguments `protobuf_oneof:"arguments"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{0}
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

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetStemcellApiVersion() int32 {
	if m != nil {
		return m.StemcellApiVersion
	}
	return 0
}

func (m *Request) GetDirectorUuid() string {
	if m != nil {
		return m.DirectorUuid
	}
	return ""
}

func (m *Request) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type isRequest_Arguments interface {
	isRequest_Arguments()
}

type Request_CreateVmArguments struct {
	CreateVmArguments *CreateVMArguments `protobuf:"bytes,5,opt,name=create_vm_arguments,json=createVmArguments,proto3,oneof"`
}

func (*Request_CreateVmArguments) isRequest_Arguments() {}

func (m *Request) GetArguments() isRequest_Arguments {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *Request) GetCreateVmArguments() *CreateVMArguments {
	if x, ok := m.GetArguments().(*Request_CreateVmArguments); ok {
		return x.CreateVmArguments
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_CreateVmArguments)(nil),
	}
}

type CreateVMArguments struct {
	AgentId              string          `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	StemcellId           string          `protobuf:"bytes,2,opt,name=stemcell_id,json=stemcellId,proto3" json:"stemcell_id,omitempty"`
	DiskCids             []string        `protobuf:"bytes,3,rep,name=disk_cids,json=diskCids,proto3" json:"disk_cids,omitempty"`
	Networks             *_struct.Struct `protobuf:"bytes,4,opt,name=networks,proto3" json:"networks,omitempty"`
	CloudProperties      *_struct.Struct `protobuf:"bytes,5,opt,name=cloud_properties,json=cloudProperties,proto3" json:"cloud_properties,omitempty"`
	Env                  *_struct.Struct `protobuf:"bytes,6,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateVMArguments) Reset()         { *m = CreateVMArguments{} }
func (m *CreateVMArguments) String() string { return proto.CompactTextString(m) }
func (*CreateVMArguments) ProtoMessage()    {}
func (*CreateVMArguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{1}
}

func (m *CreateVMArguments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMArguments.Unmarshal(m, b)
}
func (m *CreateVMArguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMArguments.Marshal(b, m, deterministic)
}
func (m *CreateVMArguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMArguments.Merge(m, src)
}
func (m *CreateVMArguments) XXX_Size() int {
	return xxx_messageInfo_CreateVMArguments.Size(m)
}
func (m *CreateVMArguments) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMArguments.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMArguments proto.InternalMessageInfo

func (m *CreateVMArguments) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *CreateVMArguments) GetStemcellId() string {
	if m != nil {
		return m.StemcellId
	}
	return ""
}

func (m *CreateVMArguments) GetDiskCids() []string {
	if m != nil {
		return m.DiskCids
	}
	return nil
}

func (m *CreateVMArguments) GetNetworks() *_struct.Struct {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *CreateVMArguments) GetCloudProperties() *_struct.Struct {
	if m != nil {
		return m.CloudProperties
	}
	return nil
}

func (m *CreateVMArguments) GetEnv() *_struct.Struct {
	if m != nil {
		return m.Env
	}
	return nil
}

type Response struct {
	Error     *Response_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	RequestId string          `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Log       string          `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Response_InfoResult
	//	*Response_CreateVmResult
	Result               isResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{2}
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

func (m *Response) GetError() *Response_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Response) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type isResponse_Result interface {
	isResponse_Result()
}

type Response_InfoResult struct {
	InfoResult *InfoResult `protobuf:"bytes,5,opt,name=info_result,json=infoResult,proto3,oneof"`
}

type Response_CreateVmResult struct {
	CreateVmResult *CreateVMResult `protobuf:"bytes,6,opt,name=create_vm_result,json=createVmResult,proto3,oneof"`
}

func (*Response_InfoResult) isResponse_Result() {}

func (*Response_CreateVmResult) isResponse_Result() {}

func (m *Response) GetResult() isResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Response) GetInfoResult() *InfoResult {
	if x, ok := m.GetResult().(*Response_InfoResult); ok {
		return x.InfoResult
	}
	return nil
}

func (m *Response) GetCreateVmResult() *CreateVMResult {
	if x, ok := m.GetResult().(*Response_CreateVmResult); ok {
		return x.CreateVmResult
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_InfoResult)(nil),
		(*Response_CreateVmResult)(nil),
	}
}

type Response_Error struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	OkToRetry            bool     `protobuf:"varint,3,opt,name=ok_to_retry,json=okToRetry,proto3" json:"ok_to_retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Error) Reset()         { *m = Response_Error{} }
func (m *Response_Error) String() string { return proto.CompactTextString(m) }
func (*Response_Error) ProtoMessage()    {}
func (*Response_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{2, 0}
}

func (m *Response_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Error.Unmarshal(m, b)
}
func (m *Response_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Error.Marshal(b, m, deterministic)
}
func (m *Response_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Error.Merge(m, src)
}
func (m *Response_Error) XXX_Size() int {
	return xxx_messageInfo_Response_Error.Size(m)
}
func (m *Response_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Error proto.InternalMessageInfo

func (m *Response_Error) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Response_Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response_Error) GetOkToRetry() bool {
	if m != nil {
		return m.OkToRetry
	}
	return false
}

type InfoResult struct {
	ApiVersion           int32    `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	StemcellFormats      []string `protobuf:"bytes,2,rep,name=stemcell_formats,json=stemcellFormats,proto3" json:"stemcell_formats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResult) Reset()         { *m = InfoResult{} }
func (m *InfoResult) String() string { return proto.CompactTextString(m) }
func (*InfoResult) ProtoMessage()    {}
func (*InfoResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{3}
}

func (m *InfoResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResult.Unmarshal(m, b)
}
func (m *InfoResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResult.Marshal(b, m, deterministic)
}
func (m *InfoResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResult.Merge(m, src)
}
func (m *InfoResult) XXX_Size() int {
	return xxx_messageInfo_InfoResult.Size(m)
}
func (m *InfoResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResult.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResult proto.InternalMessageInfo

func (m *InfoResult) GetApiVersion() int32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *InfoResult) GetStemcellFormats() []string {
	if m != nil {
		return m.StemcellFormats
	}
	return nil
}

type CreateVMResult struct {
	VmCid                string          `protobuf:"bytes,1,opt,name=vm_cid,json=vmCid,proto3" json:"vm_cid,omitempty"`
	Networks             *_struct.Struct `protobuf:"bytes,2,opt,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateVMResult) Reset()         { *m = CreateVMResult{} }
func (m *CreateVMResult) String() string { return proto.CompactTextString(m) }
func (*CreateVMResult) ProtoMessage()    {}
func (*CreateVMResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{4}
}

func (m *CreateVMResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMResult.Unmarshal(m, b)
}
func (m *CreateVMResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMResult.Marshal(b, m, deterministic)
}
func (m *CreateVMResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMResult.Merge(m, src)
}
func (m *CreateVMResult) XXX_Size() int {
	return xxx_messageInfo_CreateVMResult.Size(m)
}
func (m *CreateVMResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMResult proto.InternalMessageInfo

func (m *CreateVMResult) GetVmCid() string {
	if m != nil {
		return m.VmCid
	}
	return ""
}

func (m *CreateVMResult) GetNetworks() *_struct.Struct {
	if m != nil {
		return m.Networks
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "cpi.Request")
	proto.RegisterType((*CreateVMArguments)(nil), "cpi.CreateVMArguments")
	proto.RegisterType((*Response)(nil), "cpi.Response")
	proto.RegisterType((*Response_Error)(nil), "cpi.Response.Error")
	proto.RegisterType((*InfoResult)(nil), "cpi.InfoResult")
	proto.RegisterType((*CreateVMResult)(nil), "cpi.CreateVMResult")
}

func init() { proto.RegisterFile("cpi.proto", fileDescriptor_27dcbb49f4ec00bf) }

var fileDescriptor_27dcbb49f4ec00bf = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x6d, 0xec, 0x3a, 0xb5, 0xaf, 0xfb, 0x48, 0xa7, 0x3c, 0x4c, 0x79, 0x45, 0xe9, 0x82, 0x74,
	0xe3, 0xa2, 0x74, 0xc1, 0x12, 0xb5, 0x11, 0x28, 0x59, 0x20, 0x55, 0x53, 0x1a, 0xb1, 0x40, 0xb2,
	0x5c, 0xcf, 0xc4, 0x1a, 0xc5, 0xf6, 0x98, 0x99, 0x71, 0x50, 0xf7, 0x7c, 0x00, 0x1f, 0xc4, 0xc7,
	0x21, 0x8f, 0x1f, 0x49, 0x04, 0x0a, 0xec, 0x3c, 0xe7, 0xde, 0xe3, 0x39, 0xf7, 0xdc, 0x33, 0xe0,
	0x44, 0x39, 0xf3, 0x73, 0xc1, 0x15, 0x47, 0x66, 0x94, 0xb3, 0xd3, 0x17, 0x31, 0xe7, 0x71, 0x42,
	0x2f, 0x34, 0x74, 0x5f, 0xcc, 0x2f, 0xa4, 0x12, 0x45, 0xa4, 0xaa, 0x96, 0xc1, 0x0f, 0x03, 0xf6,
	0x30, 0xfd, 0x56, 0x50, 0xa9, 0x10, 0x82, 0x5d, 0xf5, 0x90, 0x53, 0xaf, 0xd3, 0xef, 0x0c, 0x1d,
	0xac, 0xbf, 0xd1, 0x5b, 0x78, 0x24, 0x15, 0x4d, 0x23, 0x9a, 0x24, 0x41, 0x98, 0xb3, 0x60, 0x49,
	0x85, 0x64, 0x3c, 0xf3, 0x8c, 0x7e, 0x67, 0x68, 0x61, 0xd4, 0xd4, 0xae, 0x72, 0x36, 0xab, 0x2a,
	0xe8, 0x0c, 0x0e, 0x08, 0x13, 0x34, 0x52, 0x5c, 0x04, 0x45, 0xc1, 0x88, 0x67, 0xea, 0xdf, 0xed,
	0x37, 0xe0, 0x5d, 0xc1, 0x08, 0x7a, 0x07, 0x90, 0x0b, 0x9e, 0x53, 0xa1, 0x18, 0x95, 0xde, 0x6e,
	0xbf, 0x33, 0x74, 0x47, 0x4f, 0xfd, 0x4a, 0xa9, 0xdf, 0x28, 0xf5, 0x6f, 0xb5, 0x52, 0xbc, 0xd6,
	0x8a, 0x26, 0x70, 0x12, 0x09, 0x1a, 0x2a, 0x1a, 0x2c, 0xd3, 0x20, 0x14, 0x71, 0x91, 0xd2, 0x4c,
	0x49, 0xcf, 0xd2, 0x7f, 0x78, 0xe2, 0x97, 0xb3, 0x8f, 0x75, 0x7d, 0xf6, 0xe9, 0xaa, 0xa9, 0x4e,
	0x76, 0xf0, 0x71, 0x45, 0x9a, 0xa5, 0x2d, 0x78, 0xed, 0x82, 0xd3, 0xf2, 0x07, 0x3f, 0x0d, 0x38,
	0xfe, 0x83, 0x87, 0x9e, 0x81, 0x1d, 0xc6, 0x34, 0x53, 0x01, 0x23, 0xb5, 0x29, 0x7b, 0xfa, 0x3c,
	0x25, 0xe8, 0x35, 0xb8, 0xad, 0x2f, 0x8c, 0x68, 0x3b, 0x1c, 0x0c, 0x0d, 0x34, 0x25, 0xe8, 0x39,
	0x38, 0x84, 0xc9, 0x45, 0x10, 0x31, 0x22, 0x3d, 0xb3, 0x6f, 0x0e, 0x1d, 0x6c, 0x97, 0xc0, 0x98,
	0x11, 0x89, 0x2e, 0xc1, 0xce, 0xa8, 0xfa, 0xce, 0xc5, 0xe2, 0x9f, 0xc3, 0xb7, 0x8d, 0xe8, 0x1a,
	0x7a, 0x51, 0xc2, 0x0b, 0x12, 0xac, 0x39, 0x67, 0x6d, 0x27, 0x1f, 0x69, 0xc2, 0xcd, 0xca, 0xbe,
	0x73, 0x30, 0x69, 0xb6, 0xf4, 0xba, 0xdb, 0x69, 0x65, 0xcf, 0xe0, 0x97, 0x01, 0x36, 0xa6, 0x32,
	0xe7, 0x99, 0xa4, 0xe8, 0x1c, 0x2c, 0x2a, 0x04, 0x17, 0xda, 0x06, 0x77, 0x74, 0xa2, 0x8d, 0x6e,
	0xaa, 0xfe, 0x87, 0xb2, 0x84, 0xab, 0x0e, 0xf4, 0x12, 0x40, 0x54, 0x81, 0x5a, 0x19, 0xe3, 0xd4,
	0xc8, 0x94, 0xa0, 0x1e, 0x98, 0x09, 0x8f, 0xeb, 0x50, 0x94, 0x9f, 0x68, 0x04, 0x2e, 0xcb, 0xe6,
	0x3c, 0x10, 0x54, 0x16, 0x89, 0xaa, 0x47, 0x3a, 0xd2, 0x37, 0x4c, 0xb3, 0x39, 0xc7, 0x1a, 0x9e,
	0xec, 0x60, 0x60, 0xed, 0x09, 0xbd, 0x87, 0xde, 0x2a, 0x06, 0x35, 0xb1, 0xbb, 0x26, 0xad, 0xd9,
	0x65, 0x4b, 0x3e, 0x6c, 0x02, 0x50, 0x21, 0xa7, 0x77, 0x60, 0x69, 0xd5, 0x7f, 0x0d, 0xbd, 0x07,
	0x7b, 0x29, 0x95, 0x32, 0x8c, 0x69, 0xad, 0xbf, 0x39, 0xa2, 0x57, 0xe0, 0xf2, 0x45, 0xa0, 0x4a,
	0xb1, 0x4a, 0x3c, 0xe8, 0x29, 0x6c, 0xec, 0xf0, 0xc5, 0x67, 0x8e, 0x4b, 0xe0, 0xda, 0x86, 0x6e,
	0xa5, 0x66, 0xf0, 0x05, 0x60, 0xa5, 0xbe, 0x8c, 0xcb, 0xfa, 0xeb, 0xe9, 0xe8, 0xd7, 0x03, 0xe1,
	0xea, 0xd5, 0x9c, 0x43, 0xaf, 0xcd, 0xd3, 0x9c, 0x8b, 0x34, 0x54, 0xd2, 0x33, 0x74, 0x6a, 0x8e,
	0x1a, 0xfc, 0x63, 0x05, 0x0f, 0xbe, 0xc2, 0xe1, 0xe6, 0x78, 0xe8, 0x31, 0x74, 0x97, 0x69, 0x99,
	0xb4, 0x7a, 0x0a, 0x6b, 0x99, 0x8e, 0x19, 0xd9, 0x48, 0x99, 0xf1, 0x9f, 0x29, 0x1b, 0xdd, 0x82,
	0x39, 0xbe, 0x99, 0xa2, 0x33, 0xd8, 0x2d, 0xe5, 0xa3, 0xfd, 0x7a, 0xd3, 0x7a, 0x7d, 0xa7, 0x07,
	0x1b, 0x7b, 0x47, 0x6f, 0xc0, 0x6e, 0x94, 0x6c, 0x6d, 0xbc, 0xef, 0xea, 0xfb, 0x2e, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x3b, 0x84, 0xb9, 0x9c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CPIClient is the client API for CPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CPIClient interface {
	Info(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CreateVM(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type cPIClient struct {
	cc *grpc.ClientConn
}

func NewCPIClient(cc *grpc.ClientConn) CPIClient {
	return &cPIClient{cc}
}

func (c *cPIClient) Info(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cpi.CPI/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIClient) CreateVM(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cpi.CPI/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPIServer is the server API for CPI service.
type CPIServer interface {
	Info(context.Context, *Request) (*Response, error)
	CreateVM(context.Context, *Request) (*Response, error)
}

func RegisterCPIServer(s *grpc.Server, srv CPIServer) {
	s.RegisterService(&_CPI_serviceDesc, srv)
}

func _CPI_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpi.CPI/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIServer).Info(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPI_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpi.CPI/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIServer).CreateVM(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cpi.CPI",
	HandlerType: (*CPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _CPI_Info_Handler,
		},
		{
			MethodName: "CreateVM",
			Handler:    _CPI_CreateVM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cpi.proto",
}