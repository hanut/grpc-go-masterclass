// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Reps                 uint32    `protobuf:"varint,2,opt,name=reps,proto3" json:"reps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

func (m *GreetManyTimesRequest) GetReps() uint32 {
	if m != nil {
		return m.Reps
	}
	return 0
}

type GreetManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManyTimesResponse) Reset()         { *m = GreetManyTimesResponse{} }
func (m *GreetManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesResponse) ProtoMessage()    {}
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{4}
}

func (m *GreetManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesResponse.Unmarshal(m, b)
}
func (m *GreetManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesResponse.Merge(m, src)
}
func (m *GreetManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesResponse.Size(m)
}
func (m *GreetManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesResponse proto.InternalMessageInfo

func (m *GreetManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetLongRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetLongRequest) Reset()         { *m = GreetLongRequest{} }
func (m *GreetLongRequest) String() string { return proto.CompactTextString(m) }
func (*GreetLongRequest) ProtoMessage()    {}
func (*GreetLongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{5}
}

func (m *GreetLongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetLongRequest.Unmarshal(m, b)
}
func (m *GreetLongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetLongRequest.Marshal(b, m, deterministic)
}
func (m *GreetLongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetLongRequest.Merge(m, src)
}
func (m *GreetLongRequest) XXX_Size() int {
	return xxx_messageInfo_GreetLongRequest.Size(m)
}
func (m *GreetLongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetLongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetLongRequest proto.InternalMessageInfo

func (m *GreetLongRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetLongResponse struct {
	Result               []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetLongResponse) Reset()         { *m = GreetLongResponse{} }
func (m *GreetLongResponse) String() string { return proto.CompactTextString(m) }
func (*GreetLongResponse) ProtoMessage()    {}
func (*GreetLongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{6}
}

func (m *GreetLongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetLongResponse.Unmarshal(m, b)
}
func (m *GreetLongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetLongResponse.Marshal(b, m, deterministic)
}
func (m *GreetLongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetLongResponse.Merge(m, src)
}
func (m *GreetLongResponse) XXX_Size() int {
	return xxx_messageInfo_GreetLongResponse.Size(m)
}
func (m *GreetLongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetLongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetLongResponse proto.InternalMessageInfo

func (m *GreetLongResponse) GetResult() []string {
	if m != nil {
		return m.Result
	}
	return nil
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{7}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{8}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDeadlineRequest) Reset()         { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()    {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{9}
}

func (m *GreetWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineRequest.Unmarshal(m, b)
}
func (m *GreetWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineRequest.Merge(m, src)
}
func (m *GreetWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineRequest.Size(m)
}
func (m *GreetWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineRequest proto.InternalMessageInfo

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{10}
}

func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManyTimesResponse)(nil), "greet.GreetManyTimesResponse")
	proto.RegisterType((*GreetLongRequest)(nil), "greet.GreetLongRequest")
	proto.RegisterType((*GreetLongResponse)(nil), "greet.GreetLongResponse")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

func init() {
	proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871)
}

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0x86, 0x4f, 0x0f, 0x07, 0x0e, 0x9d, 0x73, 0xf0, 0x63, 0x44, 0x84, 0x02, 0x91, 0xf4, 0x46,
	0x12, 0x12, 0x20, 0xe0, 0x9d, 0x17, 0x26, 0xf8, 0xc1, 0x8d, 0x1a, 0xad, 0x26, 0x12, 0x6f, 0x4c,
	0xd1, 0xb1, 0x36, 0x81, 0xb6, 0x76, 0x0b, 0x09, 0xbf, 0xc0, 0xbf, 0x6d, 0x58, 0xb6, 0xb0, 0x60,
	0x91, 0xa4, 0x37, 0xed, 0xee, 0xce, 0x3b, 0xcf, 0xbb, 0xbb, 0x33, 0x59, 0x28, 0x58, 0x3e, 0x51,
	0xd0, 0xe0, 0x5f, 0xaf, 0x3f, 0xfb, 0xd7, 0x3d, 0xdf, 0x0d, 0x5c, 0x4c, 0xf2, 0x89, 0x7e, 0x09,
	0xe9, 0xee, 0x74, 0x60, 0x3b, 0x16, 0x96, 0x01, 0xde, 0x6c, 0x9f, 0x05, 0xcf, 0x8e, 0x39, 0xa4,
	0xbc, 0x52, 0x51, 0xaa, 0xaa, 0xa1, 0xf2, 0x95, 0x1b, 0x73, 0x48, 0x58, 0x04, 0x75, 0x60, 0x86,
	0xd1, 0xdf, 0x3c, 0x9a, 0x9e, 0x2e, 0x4c, 0x83, 0xfa, 0x09, 0xfc, 0xe7, 0x1c, 0x83, 0x3e, 0x46,
	0xc4, 0x02, 0xac, 0x41, 0xda, 0x12, 0x5c, 0x4e, 0xfa, 0xd7, 0xda, 0xae, 0xcf, 0xec, 0x43, 0x3b,
	0x63, 0x2e, 0xd0, 0x8f, 0x20, 0x23, 0x92, 0x99, 0xe7, 0x3a, 0x8c, 0x30, 0x07, 0x29, 0x9f, 0xd8,
	0x68, 0x10, 0x88, 0x5d, 0x88, 0x99, 0xde, 0x83, 0x7d, 0x2e, 0xbc, 0x36, 0x9d, 0xc9, 0x83, 0x3d,
	0x24, 0x16, 0xc7, 0x0e, 0x11, 0xfe, 0xf8, 0xe4, 0x31, 0x7e, 0x86, 0x8c, 0xc1, 0xc7, 0x7a, 0x13,
	0x72, 0xab, 0xe4, 0x0d, 0x7b, 0x39, 0x85, 0x1d, 0x9e, 0x71, 0xe5, 0x3a, 0x56, 0xac, 0x53, 0xd7,
	0x60, 0x57, 0x02, 0x44, 0xb8, 0x25, 0x24, 0xb7, 0x33, 0xc8, 0x72, 0xf1, 0xc5, 0x98, 0xfc, 0x89,
	0xeb, 0x50, 0x2c, 0xc7, 0x86, 0xb8, 0xbe, 0x05, 0x64, 0xc3, 0x19, 0xbb, 0x90, 0xe7, 0x09, 0x8f,
	0x76, 0xf0, 0x7e, 0x4e, 0xe6, 0xeb, 0xc0, 0x8e, 0xe9, 0xdc, 0x86, 0x42, 0x04, 0xe8, 0x67, 0xf7,
	0xd6, 0x67, 0x42, 0x34, 0xd5, 0x3d, 0xf9, 0x63, 0xfb, 0x85, 0xf0, 0x18, 0x92, 0x7c, 0x8e, 0x7b,
	0xb2, 0x93, 0xd8, 0x90, 0x96, 0x5d, 0x5e, 0x9c, 0xc1, 0xf5, 0x5f, 0x78, 0x07, 0x5b, 0xcb, 0xa5,
	0xc5, 0x92, 0xac, 0x5c, 0xed, 0x25, 0xad, 0xbc, 0x26, 0x1a, 0x02, 0x9b, 0x0a, 0x76, 0x40, 0x9d,
	0x97, 0x0e, 0x0f, 0x64, 0xbd, 0xd4, 0x0d, 0x5a, 0xfe, 0x7b, 0x20, 0x64, 0x54, 0x15, 0xbc, 0x15,
	0x4d, 0x1f, 0x16, 0x03, 0x8b, 0xb2, 0x7c, 0xa5, 0xce, 0x5a, 0x29, 0x3a, 0xb8, 0xe0, 0x35, 0x15,
	0xec, 0x89, 0x86, 0x92, 0x2f, 0x19, 0x0f, 0xe5, 0xc4, 0x88, 0x3a, 0x6a, 0x95, 0xf5, 0x82, 0x90,
	0xde, 0x51, 0x9f, 0xfe, 0x8a, 0x37, 0xa4, 0x9f, 0xe2, 0xcf, 0x47, 0xfb, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x01, 0x80, 0xd9, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary API
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// Server Streaming API
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	// Client Streaming API
	GreetLong(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetLongClient, error)
	// Bi-di Streaming API
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
	// Deadlines!!!!
	GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetLong(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetLongClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/GreetLong", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetLongClient{stream}
	return x, nil
}

type GreetService_GreetLongClient interface {
	Send(*GreetLongRequest) error
	CloseAndRecv() (*GreetLongResponse, error)
	grpc.ClientStream
}

type greetServiceGreetLongClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetLongClient) Send(m *GreetLongRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetLongClient) CloseAndRecv() (*GreetLongResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetLongResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary API
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// Server Streaming API
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	// Client Streaming API
	GreetLong(GreetService_GreetLongServer) error
	// Bi-di Streaming API
	GreetEveryone(GreetService_GreetEveryoneServer) error
	// Deadlines!!!!
	GreetWithDeadline(context.Context, *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTimes(req *GreetManyTimesRequest, srv GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (*UnimplementedGreetServiceServer) GreetLong(srv GreetService_GreetLongServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetLong not implemented")
}
func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (*UnimplementedGreetServiceServer) GreetWithDeadline(ctx context.Context, req *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetWithDeadline not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_GreetLong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetLong(&greetServiceGreetLongServer{stream})
}

type GreetService_GreetLongServer interface {
	SendAndClose(*GreetLongResponse) error
	Recv() (*GreetLongRequest, error)
	grpc.ServerStream
}

type greetServiceGreetLongServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetLongServer) SendAndClose(m *GreetLongResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetLongServer) Recv() (*GreetLongRequest, error) {
	m := new(GreetLongRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetLong",
			Handler:       _GreetService_GreetLong_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
