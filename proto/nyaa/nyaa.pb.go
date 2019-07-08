// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nyaa.proto

package nyaa

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The ping message containing the user's name.
type PingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The response message containing the greetings
type PingReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{1}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ShutdownRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownRequest) Reset()         { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()    {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{2}
}

func (m *ShutdownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownRequest.Unmarshal(m, b)
}
func (m *ShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownRequest.Marshal(b, m, deterministic)
}
func (m *ShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownRequest.Merge(m, src)
}
func (m *ShutdownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownRequest.Size(m)
}
func (m *ShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownRequest proto.InternalMessageInfo

type ShutdownReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownReply) Reset()         { *m = ShutdownReply{} }
func (m *ShutdownReply) String() string { return proto.CompactTextString(m) }
func (*ShutdownReply) ProtoMessage()    {}
func (*ShutdownReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{3}
}

func (m *ShutdownReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownReply.Unmarshal(m, b)
}
func (m *ShutdownReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownReply.Marshal(b, m, deterministic)
}
func (m *ShutdownReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownReply.Merge(m, src)
}
func (m *ShutdownReply) XXX_Size() int {
	return xxx_messageInfo_ShutdownReply.Size(m)
}
func (m *ShutdownReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownReply.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownReply proto.InternalMessageInfo

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Order                string   `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Page                 uint32   `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{4}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *SearchRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SearchRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *SearchRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Result struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Size                 string   `protobuf:"bytes,4,opt,name=size,proto3" json:"size,omitempty"`
	Date                 string   `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Seeders              uint32   `protobuf:"varint,6,opt,name=seeders,proto3" json:"seeders,omitempty"`
	Leechers             uint32   `protobuf:"varint,7,opt,name=leechers,proto3" json:"leechers,omitempty"`
	Downloads            uint32   `protobuf:"varint,8,opt,name=downloads,proto3" json:"downloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{5}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Result) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Result) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Result) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *Result) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Result) GetSeeders() uint32 {
	if m != nil {
		return m.Seeders
	}
	return 0
}

func (m *Result) GetLeechers() uint32 {
	if m != nil {
		return m.Leechers
	}
	return 0
}

func (m *Result) GetDownloads() uint32 {
	if m != nil {
		return m.Downloads
	}
	return 0
}

type SearchReply struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SearchReply) Reset()         { *m = SearchReply{} }
func (m *SearchReply) String() string { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()    {}
func (*SearchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{6}
}

func (m *SearchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReply.Unmarshal(m, b)
}
func (m *SearchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReply.Marshal(b, m, deterministic)
}
func (m *SearchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReply.Merge(m, src)
}
func (m *SearchReply) XXX_Size() int {
	return xxx_messageInfo_SearchReply.Size(m)
}
func (m *SearchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReply proto.InternalMessageInfo

func (m *SearchReply) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type CurrentResultsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentResultsRequest) Reset()         { *m = CurrentResultsRequest{} }
func (m *CurrentResultsRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentResultsRequest) ProtoMessage()    {}
func (*CurrentResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{7}
}

func (m *CurrentResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentResultsRequest.Unmarshal(m, b)
}
func (m *CurrentResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentResultsRequest.Marshal(b, m, deterministic)
}
func (m *CurrentResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentResultsRequest.Merge(m, src)
}
func (m *CurrentResultsRequest) XXX_Size() int {
	return xxx_messageInfo_CurrentResultsRequest.Size(m)
}
func (m *CurrentResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentResultsRequest proto.InternalMessageInfo

type CurrentResultsReply struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CurrentResultsReply) Reset()         { *m = CurrentResultsReply{} }
func (m *CurrentResultsReply) String() string { return proto.CompactTextString(m) }
func (*CurrentResultsReply) ProtoMessage()    {}
func (*CurrentResultsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fe92bf9d869316, []int{8}
}

func (m *CurrentResultsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentResultsReply.Unmarshal(m, b)
}
func (m *CurrentResultsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentResultsReply.Marshal(b, m, deterministic)
}
func (m *CurrentResultsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentResultsReply.Merge(m, src)
}
func (m *CurrentResultsReply) XXX_Size() int {
	return xxx_messageInfo_CurrentResultsReply.Size(m)
}
func (m *CurrentResultsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentResultsReply.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentResultsReply proto.InternalMessageInfo

func (m *CurrentResultsReply) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "nyaa.PingRequest")
	proto.RegisterType((*PingReply)(nil), "nyaa.PingReply")
	proto.RegisterType((*ShutdownRequest)(nil), "nyaa.ShutdownRequest")
	proto.RegisterType((*ShutdownReply)(nil), "nyaa.ShutdownReply")
	proto.RegisterType((*SearchRequest)(nil), "nyaa.SearchRequest")
	proto.RegisterType((*Result)(nil), "nyaa.Result")
	proto.RegisterType((*SearchReply)(nil), "nyaa.SearchReply")
	proto.RegisterType((*CurrentResultsRequest)(nil), "nyaa.CurrentResultsRequest")
	proto.RegisterType((*CurrentResultsReply)(nil), "nyaa.CurrentResultsReply")
}

func init() { proto.RegisterFile("nyaa.proto", fileDescriptor_35fe92bf9d869316) }

var fileDescriptor_35fe92bf9d869316 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0xd7, 0x49, 0x26, 0x49, 0x43, 0xb6, 0xa4, 0x31, 0xa6, 0x48, 0x91, 0x25, 0x20,
	0x42, 0xa2, 0x91, 0x8a, 0xb8, 0x20, 0xc1, 0x85, 0x13, 0x97, 0x0a, 0xb9, 0x5f, 0xb0, 0xc4, 0x83,
	0x63, 0xe1, 0x7a, 0xdd, 0xdd, 0xb5, 0x90, 0x7b, 0xe4, 0x17, 0x38, 0xf0, 0x4b, 0x88, 0x2b, 0xbf,
	0xc0, 0x87, 0xa0, 0x9d, 0xf5, 0xa6, 0x49, 0xa8, 0x90, 0x7a, 0x9b, 0x79, 0x6f, 0xe6, 0xcd, 0xec,
	0xf8, 0x19, 0xa0, 0x6c, 0x38, 0x3f, 0xab, 0xa4, 0xd0, 0x82, 0xf9, 0x26, 0x8e, 0x4e, 0x33, 0x21,
	0xb2, 0x02, 0x97, 0xbc, 0xca, 0x97, 0xbc, 0x2c, 0x85, 0xe6, 0x3a, 0x17, 0xa5, 0xb2, 0x35, 0xf1,
	0x73, 0x18, 0x7c, 0xcc, 0xcb, 0x2c, 0xc1, 0xeb, 0x1a, 0x95, 0x66, 0x21, 0x74, 0xaf, 0x50, 0x29,
	0x9e, 0x61, 0xe8, 0xcd, 0xbd, 0x45, 0x3f, 0x71, 0x69, 0xfc, 0x14, 0xfa, 0xb6, 0xb0, 0x2a, 0x9a,
	0xff, 0x94, 0x4d, 0x60, 0x7c, 0xb9, 0xae, 0x75, 0x2a, 0xbe, 0x96, 0xad, 0x66, 0x3c, 0x86, 0xd1,
	0x2d, 0x54, 0x15, 0x4d, 0xfc, 0xc3, 0x83, 0xd1, 0x25, 0x72, 0xb9, 0x5a, 0xbb, 0xb1, 0x0f, 0xe1,
	0xf0, 0xba, 0x46, 0xd9, 0xb4, 0x6a, 0x36, 0x61, 0x27, 0x10, 0x7c, 0xce, 0x0b, 0x8d, 0x32, 0xec,
	0x10, 0xdc, 0x66, 0x2c, 0x82, 0xde, 0x8a, 0x6b, 0xcc, 0x84, 0x6c, 0xc2, 0x03, 0x62, 0x36, 0x39,
	0x63, 0xe0, 0x2b, 0x21, 0x75, 0xe8, 0x13, 0x4e, 0xb1, 0x51, 0x17, 0x32, 0x45, 0x19, 0x1e, 0x5a,
	0x75, 0x4a, 0x4c, 0x65, 0x65, 0x1e, 0x10, 0xcc, 0xbd, 0xc5, 0x28, 0xa1, 0x38, 0xfe, 0xe9, 0x41,
	0x90, 0xa0, 0xaa, 0x0b, 0xbd, 0x33, 0xc4, 0xfb, 0x77, 0x48, 0xc9, 0xaf, 0xb0, 0x5d, 0x8b, 0x62,
	0x83, 0x15, 0x79, 0xf9, 0xa5, 0x5d, 0x88, 0x62, 0x5a, 0x26, 0xbf, 0xc1, 0xcd, 0x32, 0xf9, 0x0d,
	0xd5, 0xa5, 0x5c, 0x63, 0xbb, 0x0b, 0xc5, 0xe6, 0x9c, 0x0a, 0x31, 0x45, 0xa9, 0xda, 0x6d, 0x5c,
	0x6a, 0xb6, 0x28, 0x10, 0x57, 0x6b, 0x43, 0x75, 0x89, 0xda, 0xe4, 0xec, 0x14, 0xfa, 0xe6, 0xa6,
	0x85, 0xe0, 0xa9, 0x0a, 0x7b, 0x44, 0xde, 0x02, 0xf1, 0x6b, 0x18, 0xb8, 0x1b, 0x9b, 0x2f, 0xf6,
	0x0c, 0xba, 0x92, 0x1e, 0xa6, 0x42, 0x6f, 0x7e, 0xb0, 0x18, 0x9c, 0x0f, 0xcf, 0xc8, 0x29, 0xf6,
	0xb5, 0x89, 0x23, 0xe3, 0x19, 0x4c, 0xdf, 0xd7, 0x52, 0x62, 0xa9, 0x2d, 0xa3, 0xdc, 0x57, 0x7c,
	0x0b, 0xc7, 0xfb, 0xc4, 0x3d, 0x74, 0xcf, 0x7f, 0x75, 0xc0, 0xbf, 0x68, 0x38, 0x67, 0xef, 0xc0,
	0x37, 0x3e, 0x62, 0x13, 0x5b, 0xb7, 0x65, 0xbe, 0x68, 0xbc, 0x0d, 0x19, 0xa3, 0xb0, 0x6f, 0xbf,
	0xff, 0x7c, 0xef, 0x0c, 0x19, 0x2c, 0x0d, 0xb1, 0xac, 0x4c, 0xdf, 0x05, 0xf4, 0x9c, 0x9b, 0xd8,
	0xd4, 0x36, 0xec, 0x19, 0x2e, 0x3a, 0xde, 0x87, 0x8d, 0xd6, 0x09, 0x69, 0x3d, 0x60, 0x47, 0x56,
	0x4b, 0x39, 0x8d, 0x0f, 0x10, 0xd8, 0x3b, 0x31, 0xd7, 0xb6, 0xed, 0xcc, 0x68, 0xb2, 0x0b, 0x1a,
	0xa5, 0x19, 0x29, 0x4d, 0xe2, 0x61, 0xab, 0x44, 0xd4, 0x1b, 0xef, 0x05, 0xcb, 0xe0, 0x68, 0xf7,
	0x44, 0xec, 0xb1, 0xed, 0xbe, 0xf3, 0xa2, 0xd1, 0xa3, 0xbb, 0x49, 0x33, 0xe2, 0x09, 0x8d, 0x98,
	0xb1, 0xa9, 0x1d, 0xb1, 0xb2, 0x25, 0x2f, 0xdb, 0x63, 0x7e, 0x0a, 0xe8, 0xdf, 0x7d, 0xf5, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xdf, 0x05, 0x24, 0xed, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NyaaClient is the client API for Nyaa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NyaaClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownReply, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	CurrentResults(ctx context.Context, in *CurrentResultsRequest, opts ...grpc.CallOption) (*CurrentResultsReply, error)
}

type nyaaClient struct {
	cc *grpc.ClientConn
}

func NewNyaaClient(cc *grpc.ClientConn) NyaaClient {
	return &nyaaClient{cc}
}

func (c *nyaaClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/nyaa.Nyaa/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nyaaClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownReply, error) {
	out := new(ShutdownReply)
	err := c.cc.Invoke(ctx, "/nyaa.Nyaa/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nyaaClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/nyaa.Nyaa/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nyaaClient) CurrentResults(ctx context.Context, in *CurrentResultsRequest, opts ...grpc.CallOption) (*CurrentResultsReply, error) {
	out := new(CurrentResultsReply)
	err := c.cc.Invoke(ctx, "/nyaa.Nyaa/CurrentResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NyaaServer is the server API for Nyaa service.
type NyaaServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownReply, error)
	Search(context.Context, *SearchRequest) (*SearchReply, error)
	CurrentResults(context.Context, *CurrentResultsRequest) (*CurrentResultsReply, error)
}

// UnimplementedNyaaServer can be embedded to have forward compatible implementations.
type UnimplementedNyaaServer struct {
}

func (*UnimplementedNyaaServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedNyaaServer) Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (*UnimplementedNyaaServer) Search(ctx context.Context, req *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedNyaaServer) CurrentResults(ctx context.Context, req *CurrentResultsRequest) (*CurrentResultsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentResults not implemented")
}

func RegisterNyaaServer(s *grpc.Server, srv NyaaServer) {
	s.RegisterService(&_Nyaa_serviceDesc, srv)
}

func _Nyaa_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NyaaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nyaa.Nyaa/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NyaaServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nyaa_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NyaaServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nyaa.Nyaa/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NyaaServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nyaa_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NyaaServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nyaa.Nyaa/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NyaaServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nyaa_CurrentResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NyaaServer).CurrentResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nyaa.Nyaa/CurrentResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NyaaServer).CurrentResults(ctx, req.(*CurrentResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nyaa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nyaa.Nyaa",
	HandlerType: (*NyaaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Nyaa_Ping_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Nyaa_Shutdown_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Nyaa_Search_Handler,
		},
		{
			MethodName: "CurrentResults",
			Handler:    _Nyaa_CurrentResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nyaa.proto",
}