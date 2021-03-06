// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ultimate_torrent.proto

package ultimate_torrent

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
	return fileDescriptor_aa1b505f29ce2a88, []int{0}
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
	return fileDescriptor_aa1b505f29ce2a88, []int{1}
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
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownRequest) Reset()         { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()    {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{2}
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

func (m *ShutdownRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ShutdownReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownReply) Reset()         { *m = ShutdownReply{} }
func (m *ShutdownReply) String() string { return proto.CompactTextString(m) }
func (*ShutdownReply) ProtoMessage()    {}
func (*ShutdownReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{3}
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

func (m *ShutdownReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddMagnetUriRequest struct {
	Magnet               string   `protobuf:"bytes,1,opt,name=magnet,proto3" json:"magnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMagnetUriRequest) Reset()         { *m = AddMagnetUriRequest{} }
func (m *AddMagnetUriRequest) String() string { return proto.CompactTextString(m) }
func (*AddMagnetUriRequest) ProtoMessage()    {}
func (*AddMagnetUriRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{4}
}

func (m *AddMagnetUriRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMagnetUriRequest.Unmarshal(m, b)
}
func (m *AddMagnetUriRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMagnetUriRequest.Marshal(b, m, deterministic)
}
func (m *AddMagnetUriRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMagnetUriRequest.Merge(m, src)
}
func (m *AddMagnetUriRequest) XXX_Size() int {
	return xxx_messageInfo_AddMagnetUriRequest.Size(m)
}
func (m *AddMagnetUriRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMagnetUriRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMagnetUriRequest proto.InternalMessageInfo

func (m *AddMagnetUriRequest) GetMagnet() string {
	if m != nil {
		return m.Magnet
	}
	return ""
}

type AddMagnetUriReply struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMagnetUriReply) Reset()         { *m = AddMagnetUriReply{} }
func (m *AddMagnetUriReply) String() string { return proto.CompactTextString(m) }
func (*AddMagnetUriReply) ProtoMessage()    {}
func (*AddMagnetUriReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{5}
}

func (m *AddMagnetUriReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMagnetUriReply.Unmarshal(m, b)
}
func (m *AddMagnetUriReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMagnetUriReply.Marshal(b, m, deterministic)
}
func (m *AddMagnetUriReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMagnetUriReply.Merge(m, src)
}
func (m *AddMagnetUriReply) XXX_Size() int {
	return xxx_messageInfo_AddMagnetUriReply.Size(m)
}
func (m *AddMagnetUriReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMagnetUriReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddMagnetUriReply proto.InternalMessageInfo

func (m *AddMagnetUriReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type RemoveTorrentRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTorrentRequest) Reset()         { *m = RemoveTorrentRequest{} }
func (m *RemoveTorrentRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTorrentRequest) ProtoMessage()    {}
func (*RemoveTorrentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{6}
}

func (m *RemoveTorrentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTorrentRequest.Unmarshal(m, b)
}
func (m *RemoveTorrentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTorrentRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTorrentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTorrentRequest.Merge(m, src)
}
func (m *RemoveTorrentRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTorrentRequest.Size(m)
}
func (m *RemoveTorrentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTorrentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTorrentRequest proto.InternalMessageInfo

func (m *RemoveTorrentRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type RemoveTorrentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTorrentReply) Reset()         { *m = RemoveTorrentReply{} }
func (m *RemoveTorrentReply) String() string { return proto.CompactTextString(m) }
func (*RemoveTorrentReply) ProtoMessage()    {}
func (*RemoveTorrentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{7}
}

func (m *RemoveTorrentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTorrentReply.Unmarshal(m, b)
}
func (m *RemoveTorrentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTorrentReply.Marshal(b, m, deterministic)
}
func (m *RemoveTorrentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTorrentReply.Merge(m, src)
}
func (m *RemoveTorrentReply) XXX_Size() int {
	return xxx_messageInfo_RemoveTorrentReply.Size(m)
}
func (m *RemoveTorrentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTorrentReply.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTorrentReply proto.InternalMessageInfo

type PauseTorrentRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseTorrentRequest) Reset()         { *m = PauseTorrentRequest{} }
func (m *PauseTorrentRequest) String() string { return proto.CompactTextString(m) }
func (*PauseTorrentRequest) ProtoMessage()    {}
func (*PauseTorrentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{8}
}

func (m *PauseTorrentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseTorrentRequest.Unmarshal(m, b)
}
func (m *PauseTorrentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseTorrentRequest.Marshal(b, m, deterministic)
}
func (m *PauseTorrentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseTorrentRequest.Merge(m, src)
}
func (m *PauseTorrentRequest) XXX_Size() int {
	return xxx_messageInfo_PauseTorrentRequest.Size(m)
}
func (m *PauseTorrentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseTorrentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseTorrentRequest proto.InternalMessageInfo

func (m *PauseTorrentRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type PauseTorrentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseTorrentReply) Reset()         { *m = PauseTorrentReply{} }
func (m *PauseTorrentReply) String() string { return proto.CompactTextString(m) }
func (*PauseTorrentReply) ProtoMessage()    {}
func (*PauseTorrentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{9}
}

func (m *PauseTorrentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseTorrentReply.Unmarshal(m, b)
}
func (m *PauseTorrentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseTorrentReply.Marshal(b, m, deterministic)
}
func (m *PauseTorrentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseTorrentReply.Merge(m, src)
}
func (m *PauseTorrentReply) XXX_Size() int {
	return xxx_messageInfo_PauseTorrentReply.Size(m)
}
func (m *PauseTorrentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseTorrentReply.DiscardUnknown(m)
}

var xxx_messageInfo_PauseTorrentReply proto.InternalMessageInfo

type ResumeTorrentRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeTorrentRequest) Reset()         { *m = ResumeTorrentRequest{} }
func (m *ResumeTorrentRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeTorrentRequest) ProtoMessage()    {}
func (*ResumeTorrentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{10}
}

func (m *ResumeTorrentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeTorrentRequest.Unmarshal(m, b)
}
func (m *ResumeTorrentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeTorrentRequest.Marshal(b, m, deterministic)
}
func (m *ResumeTorrentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeTorrentRequest.Merge(m, src)
}
func (m *ResumeTorrentRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeTorrentRequest.Size(m)
}
func (m *ResumeTorrentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeTorrentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeTorrentRequest proto.InternalMessageInfo

func (m *ResumeTorrentRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ResumeTorrentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeTorrentReply) Reset()         { *m = ResumeTorrentReply{} }
func (m *ResumeTorrentReply) String() string { return proto.CompactTextString(m) }
func (*ResumeTorrentReply) ProtoMessage()    {}
func (*ResumeTorrentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{11}
}

func (m *ResumeTorrentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeTorrentReply.Unmarshal(m, b)
}
func (m *ResumeTorrentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeTorrentReply.Marshal(b, m, deterministic)
}
func (m *ResumeTorrentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeTorrentReply.Merge(m, src)
}
func (m *ResumeTorrentReply) XXX_Size() int {
	return xxx_messageInfo_ResumeTorrentReply.Size(m)
}
func (m *ResumeTorrentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeTorrentReply.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeTorrentReply proto.InternalMessageInfo

type TorrentStatus struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Progress             float32  `protobuf:"fixed32,3,opt,name=progress,proto3" json:"progress,omitempty"`
	DownloadRate         int32    `protobuf:"zigzag32,4,opt,name=download_rate,json=downloadRate,proto3" json:"download_rate,omitempty"`
	UploadRate           int32    `protobuf:"zigzag32,5,opt,name=upload_rate,json=uploadRate,proto3" json:"upload_rate,omitempty"`
	Peers                int32    `protobuf:"zigzag32,6,opt,name=peers,proto3" json:"peers,omitempty"`
	Seeds                int32    `protobuf:"zigzag32,7,opt,name=seeds,proto3" json:"seeds,omitempty"`
	State                string   `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	TotalSize            uint64   `protobuf:"varint,9,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentStatus) Reset()         { *m = TorrentStatus{} }
func (m *TorrentStatus) String() string { return proto.CompactTextString(m) }
func (*TorrentStatus) ProtoMessage()    {}
func (*TorrentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{12}
}

func (m *TorrentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentStatus.Unmarshal(m, b)
}
func (m *TorrentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentStatus.Marshal(b, m, deterministic)
}
func (m *TorrentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentStatus.Merge(m, src)
}
func (m *TorrentStatus) XXX_Size() int {
	return xxx_messageInfo_TorrentStatus.Size(m)
}
func (m *TorrentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentStatus proto.InternalMessageInfo

func (m *TorrentStatus) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TorrentStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TorrentStatus) GetProgress() float32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *TorrentStatus) GetDownloadRate() int32 {
	if m != nil {
		return m.DownloadRate
	}
	return 0
}

func (m *TorrentStatus) GetUploadRate() int32 {
	if m != nil {
		return m.UploadRate
	}
	return 0
}

func (m *TorrentStatus) GetPeers() int32 {
	if m != nil {
		return m.Peers
	}
	return 0
}

func (m *TorrentStatus) GetSeeds() int32 {
	if m != nil {
		return m.Seeds
	}
	return 0
}

func (m *TorrentStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TorrentStatus) GetTotalSize() uint64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

type AllTorrentStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllTorrentStatusRequest) Reset()         { *m = AllTorrentStatusRequest{} }
func (m *AllTorrentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AllTorrentStatusRequest) ProtoMessage()    {}
func (*AllTorrentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{13}
}

func (m *AllTorrentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTorrentStatusRequest.Unmarshal(m, b)
}
func (m *AllTorrentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTorrentStatusRequest.Marshal(b, m, deterministic)
}
func (m *AllTorrentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTorrentStatusRequest.Merge(m, src)
}
func (m *AllTorrentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_AllTorrentStatusRequest.Size(m)
}
func (m *AllTorrentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTorrentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllTorrentStatusRequest proto.InternalMessageInfo

type AllTorrentStatusReply struct {
	AllTorrentStatus     []*TorrentStatus `protobuf:"bytes,1,rep,name=all_torrent_status,json=allTorrentStatus,proto3" json:"all_torrent_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AllTorrentStatusReply) Reset()         { *m = AllTorrentStatusReply{} }
func (m *AllTorrentStatusReply) String() string { return proto.CompactTextString(m) }
func (*AllTorrentStatusReply) ProtoMessage()    {}
func (*AllTorrentStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1b505f29ce2a88, []int{14}
}

func (m *AllTorrentStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTorrentStatusReply.Unmarshal(m, b)
}
func (m *AllTorrentStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTorrentStatusReply.Marshal(b, m, deterministic)
}
func (m *AllTorrentStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTorrentStatusReply.Merge(m, src)
}
func (m *AllTorrentStatusReply) XXX_Size() int {
	return xxx_messageInfo_AllTorrentStatusReply.Size(m)
}
func (m *AllTorrentStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTorrentStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_AllTorrentStatusReply proto.InternalMessageInfo

func (m *AllTorrentStatusReply) GetAllTorrentStatus() []*TorrentStatus {
	if m != nil {
		return m.AllTorrentStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "ultimate_torrent.PingRequest")
	proto.RegisterType((*PingReply)(nil), "ultimate_torrent.PingReply")
	proto.RegisterType((*ShutdownRequest)(nil), "ultimate_torrent.ShutdownRequest")
	proto.RegisterType((*ShutdownReply)(nil), "ultimate_torrent.ShutdownReply")
	proto.RegisterType((*AddMagnetUriRequest)(nil), "ultimate_torrent.AddMagnetUriRequest")
	proto.RegisterType((*AddMagnetUriReply)(nil), "ultimate_torrent.AddMagnetUriReply")
	proto.RegisterType((*RemoveTorrentRequest)(nil), "ultimate_torrent.RemoveTorrentRequest")
	proto.RegisterType((*RemoveTorrentReply)(nil), "ultimate_torrent.RemoveTorrentReply")
	proto.RegisterType((*PauseTorrentRequest)(nil), "ultimate_torrent.PauseTorrentRequest")
	proto.RegisterType((*PauseTorrentReply)(nil), "ultimate_torrent.PauseTorrentReply")
	proto.RegisterType((*ResumeTorrentRequest)(nil), "ultimate_torrent.ResumeTorrentRequest")
	proto.RegisterType((*ResumeTorrentReply)(nil), "ultimate_torrent.ResumeTorrentReply")
	proto.RegisterType((*TorrentStatus)(nil), "ultimate_torrent.TorrentStatus")
	proto.RegisterType((*AllTorrentStatusRequest)(nil), "ultimate_torrent.AllTorrentStatusRequest")
	proto.RegisterType((*AllTorrentStatusReply)(nil), "ultimate_torrent.AllTorrentStatusReply")
}

func init() { proto.RegisterFile("ultimate_torrent.proto", fileDescriptor_aa1b505f29ce2a88) }

var fileDescriptor_aa1b505f29ce2a88 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0x95, 0xad, 0xed, 0xda, 0xdb, 0x55, 0x6b, 0xdd, 0x7e, 0xf7, 0x0d, 0x41, 0x53, 0x43,
	0xc6, 0x68, 0x0a, 0x62, 0x42, 0xe3, 0x2f, 0x18, 0xef, 0x43, 0x53, 0xca, 0x5e, 0x90, 0x50, 0x65,
	0x94, 0x4b, 0x1b, 0xc9, 0xf9, 0x41, 0xec, 0x80, 0xba, 0x77, 0xfe, 0x69, 0x9e, 0x90, 0x9d, 0x84,
	0x35, 0x89, 0x69, 0xfb, 0x96, 0x7b, 0xfc, 0xb9, 0x47, 0x67, 0x9e, 0x4f, 0xe1, 0x3c, 0x63, 0x22,
	0x08, 0xa9, 0xc0, 0xa5, 0x88, 0xd3, 0x14, 0x23, 0x71, 0x9d, 0xa4, 0xb1, 0x88, 0xc9, 0xb0, 0xae,
	0x3b, 0x33, 0xe8, 0xdf, 0x07, 0xd1, 0xca, 0xc3, 0xef, 0x19, 0x72, 0x41, 0x4c, 0x38, 0x09, 0x91,
	0x73, 0xba, 0x42, 0xd3, 0xb0, 0x0d, 0xb7, 0xe7, 0x95, 0xa3, 0x73, 0x05, 0xbd, 0x1c, 0x4c, 0xd8,
	0x66, 0x07, 0xf6, 0x06, 0xce, 0x16, 0xeb, 0x4c, 0xf8, 0xf1, 0xcf, 0x68, 0xbf, 0xe7, 0x1c, 0x06,
	0x4f, 0xf0, 0x6e, 0xdf, 0xb7, 0x30, 0xbe, 0xf5, 0xfd, 0x3b, 0xba, 0x8a, 0x50, 0x3c, 0xa4, 0x41,
	0xe9, 0x7d, 0x0e, 0x9d, 0x50, 0x69, 0x05, 0x5f, 0x4c, 0xce, 0x0c, 0x46, 0x55, 0x5c, 0xba, 0x13,
	0x68, 0xad, 0x29, 0x5f, 0x17, 0xa8, 0xfa, 0x76, 0x5e, 0xc3, 0xc4, 0xc3, 0x30, 0xfe, 0x81, 0x9f,
	0xf2, 0x0b, 0x29, 0x8d, 0x75, 0xec, 0x04, 0x48, 0x8d, 0x4d, 0xd8, 0xc6, 0x99, 0xc3, 0xf8, 0x9e,
	0x66, 0xfc, 0x10, 0x83, 0x31, 0x8c, 0xaa, 0xa8, 0xdc, 0x57, 0x09, 0x78, 0x16, 0x1e, 0x9c, 0xa0,
	0xc2, 0x4a, 0x87, 0xdf, 0x06, 0x0c, 0x0a, 0x61, 0x21, 0xa8, 0xc8, 0xb8, 0x6e, 0x57, 0x6a, 0x11,
	0x0d, 0xd1, 0x3c, 0xca, 0x35, 0xf9, 0x4d, 0x2c, 0xe8, 0x26, 0x69, 0xbc, 0x4a, 0x91, 0x73, 0xf3,
	0xd8, 0x36, 0xdc, 0x23, 0xef, 0xef, 0x4c, 0x2e, 0x61, 0x20, 0xff, 0x31, 0x2c, 0xa6, 0xfe, 0x32,
	0xa5, 0x02, 0xcd, 0x96, 0x6d, 0xb8, 0x23, 0xef, 0xb4, 0x14, 0x3d, 0x2a, 0x90, 0x4c, 0xa1, 0x9f,
	0x25, 0x4f, 0x48, 0x5b, 0x21, 0x90, 0x4b, 0x0a, 0x98, 0x40, 0x3b, 0x41, 0x4c, 0xb9, 0xd9, 0x51,
	0x47, 0xf9, 0x20, 0x55, 0x8e, 0xe8, 0x73, 0xf3, 0x24, 0x57, 0xd5, 0xa0, 0x54, 0x21, 0x6d, 0xba,
	0x2a, 0x62, 0x3e, 0x90, 0x0b, 0x00, 0x11, 0x0b, 0xca, 0x96, 0x3c, 0x78, 0x44, 0xb3, 0x67, 0x1b,
	0x6e, 0xcb, 0xeb, 0x29, 0x65, 0x11, 0x3c, 0xa2, 0xf3, 0x0c, 0xfe, 0xbf, 0x65, 0xac, 0xf2, 0xe7,
	0x17, 0x37, 0xe8, 0x7c, 0x83, 0xff, 0x9a, 0x47, 0xf2, 0x21, 0xdc, 0x01, 0xa1, 0x8c, 0x95, 0x1d,
	0x58, 0x72, 0x75, 0x64, 0x1a, 0xf6, 0xb1, 0xdb, 0xbf, 0x99, 0x5e, 0x37, 0xba, 0x53, 0x75, 0x18,
	0xd2, 0x9a, 0xe7, 0xcd, 0xaf, 0x36, 0x9c, 0x3d, 0x14, 0x4b, 0xc5, 0x09, 0xf9, 0x00, 0x2d, 0x59,
	0x17, 0x72, 0xd1, 0xb4, 0xdb, 0xea, 0x9b, 0xf5, 0xfc, 0x5f, 0xc7, 0x32, 0xe6, 0x47, 0xe8, 0x96,
	0xf5, 0x20, 0x2f, 0x9a, 0x60, 0xad, 0x67, 0xd6, 0x74, 0x17, 0x22, 0xfd, 0x3e, 0xc3, 0xe9, 0x76,
	0x29, 0xc8, 0x55, 0x73, 0x41, 0xd3, 0x31, 0xeb, 0x72, 0x1f, 0x26, 0xbd, 0xbf, 0xc0, 0xa0, 0xd2,
	0x0d, 0xf2, 0xaa, 0xb9, 0xa5, 0x2b, 0x9a, 0xf5, 0x72, 0x2f, 0x57, 0x44, 0xdf, 0x6e, 0x8e, 0x2e,
	0xba, 0xa6, 0x84, 0xba, 0xe8, 0x8d, 0x02, 0xe6, 0xd1, 0xb7, 0x4a, 0xa5, 0x8f, 0xde, 0x6c, 0xa8,
	0x3e, 0x7a, 0xbd, 0x9d, 0x64, 0x0d, 0xc3, 0xfa, 0x2b, 0x24, 0x73, 0xcd, 0x95, 0xea, 0x1f, 0xb1,
	0x35, 0x3b, 0x04, 0x4d, 0xd8, 0xe6, 0x9d, 0xf1, 0xb5, 0xa3, 0x7e, 0xe4, 0xdf, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0xb4, 0x64, 0x09, 0xfe, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UltimateTorrentClient is the client API for UltimateTorrent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UltimateTorrentClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownReply, error)
	AddMagnetUri(ctx context.Context, in *AddMagnetUriRequest, opts ...grpc.CallOption) (*AddMagnetUriReply, error)
	RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, opts ...grpc.CallOption) (*RemoveTorrentReply, error)
	PauseTorrent(ctx context.Context, in *PauseTorrentRequest, opts ...grpc.CallOption) (*PauseTorrentReply, error)
	ResumeTorrent(ctx context.Context, in *ResumeTorrentRequest, opts ...grpc.CallOption) (*ResumeTorrentReply, error)
	AllTorrentStatus(ctx context.Context, in *AllTorrentStatusRequest, opts ...grpc.CallOption) (UltimateTorrent_AllTorrentStatusClient, error)
}

type ultimateTorrentClient struct {
	cc *grpc.ClientConn
}

func NewUltimateTorrentClient(cc *grpc.ClientConn) UltimateTorrentClient {
	return &ultimateTorrentClient{cc}
}

func (c *ultimateTorrentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownReply, error) {
	out := new(ShutdownReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) AddMagnetUri(ctx context.Context, in *AddMagnetUriRequest, opts ...grpc.CallOption) (*AddMagnetUriReply, error) {
	out := new(AddMagnetUriReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/AddMagnetUri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, opts ...grpc.CallOption) (*RemoveTorrentReply, error) {
	out := new(RemoveTorrentReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/RemoveTorrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) PauseTorrent(ctx context.Context, in *PauseTorrentRequest, opts ...grpc.CallOption) (*PauseTorrentReply, error) {
	out := new(PauseTorrentReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/PauseTorrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) ResumeTorrent(ctx context.Context, in *ResumeTorrentRequest, opts ...grpc.CallOption) (*ResumeTorrentReply, error) {
	out := new(ResumeTorrentReply)
	err := c.cc.Invoke(ctx, "/ultimate_torrent.UltimateTorrent/ResumeTorrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultimateTorrentClient) AllTorrentStatus(ctx context.Context, in *AllTorrentStatusRequest, opts ...grpc.CallOption) (UltimateTorrent_AllTorrentStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UltimateTorrent_serviceDesc.Streams[0], "/ultimate_torrent.UltimateTorrent/AllTorrentStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultimateTorrentAllTorrentStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltimateTorrent_AllTorrentStatusClient interface {
	Recv() (*AllTorrentStatusReply, error)
	grpc.ClientStream
}

type ultimateTorrentAllTorrentStatusClient struct {
	grpc.ClientStream
}

func (x *ultimateTorrentAllTorrentStatusClient) Recv() (*AllTorrentStatusReply, error) {
	m := new(AllTorrentStatusReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UltimateTorrentServer is the server API for UltimateTorrent service.
type UltimateTorrentServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownReply, error)
	AddMagnetUri(context.Context, *AddMagnetUriRequest) (*AddMagnetUriReply, error)
	RemoveTorrent(context.Context, *RemoveTorrentRequest) (*RemoveTorrentReply, error)
	PauseTorrent(context.Context, *PauseTorrentRequest) (*PauseTorrentReply, error)
	ResumeTorrent(context.Context, *ResumeTorrentRequest) (*ResumeTorrentReply, error)
	AllTorrentStatus(*AllTorrentStatusRequest, UltimateTorrent_AllTorrentStatusServer) error
}

// UnimplementedUltimateTorrentServer can be embedded to have forward compatible implementations.
type UnimplementedUltimateTorrentServer struct {
}

func (*UnimplementedUltimateTorrentServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedUltimateTorrentServer) Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (*UnimplementedUltimateTorrentServer) AddMagnetUri(ctx context.Context, req *AddMagnetUriRequest) (*AddMagnetUriReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMagnetUri not implemented")
}
func (*UnimplementedUltimateTorrentServer) RemoveTorrent(ctx context.Context, req *RemoveTorrentRequest) (*RemoveTorrentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTorrent not implemented")
}
func (*UnimplementedUltimateTorrentServer) PauseTorrent(ctx context.Context, req *PauseTorrentRequest) (*PauseTorrentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseTorrent not implemented")
}
func (*UnimplementedUltimateTorrentServer) ResumeTorrent(ctx context.Context, req *ResumeTorrentRequest) (*ResumeTorrentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeTorrent not implemented")
}
func (*UnimplementedUltimateTorrentServer) AllTorrentStatus(req *AllTorrentStatusRequest, srv UltimateTorrent_AllTorrentStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method AllTorrentStatus not implemented")
}

func RegisterUltimateTorrentServer(s *grpc.Server, srv UltimateTorrentServer) {
	s.RegisterService(&_UltimateTorrent_serviceDesc, srv)
}

func _UltimateTorrent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_AddMagnetUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMagnetUriRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).AddMagnetUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/AddMagnetUri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).AddMagnetUri(ctx, req.(*AddMagnetUriRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_RemoveTorrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).RemoveTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/RemoveTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).RemoveTorrent(ctx, req.(*RemoveTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_PauseTorrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).PauseTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/PauseTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).PauseTorrent(ctx, req.(*PauseTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_ResumeTorrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltimateTorrentServer).ResumeTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultimate_torrent.UltimateTorrent/ResumeTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltimateTorrentServer).ResumeTorrent(ctx, req.(*ResumeTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltimateTorrent_AllTorrentStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllTorrentStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltimateTorrentServer).AllTorrentStatus(m, &ultimateTorrentAllTorrentStatusServer{stream})
}

type UltimateTorrent_AllTorrentStatusServer interface {
	Send(*AllTorrentStatusReply) error
	grpc.ServerStream
}

type ultimateTorrentAllTorrentStatusServer struct {
	grpc.ServerStream
}

func (x *ultimateTorrentAllTorrentStatusServer) Send(m *AllTorrentStatusReply) error {
	return x.ServerStream.SendMsg(m)
}

var _UltimateTorrent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ultimate_torrent.UltimateTorrent",
	HandlerType: (*UltimateTorrentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UltimateTorrent_Ping_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _UltimateTorrent_Shutdown_Handler,
		},
		{
			MethodName: "AddMagnetUri",
			Handler:    _UltimateTorrent_AddMagnetUri_Handler,
		},
		{
			MethodName: "RemoveTorrent",
			Handler:    _UltimateTorrent_RemoveTorrent_Handler,
		},
		{
			MethodName: "PauseTorrent",
			Handler:    _UltimateTorrent_PauseTorrent_Handler,
		},
		{
			MethodName: "ResumeTorrent",
			Handler:    _UltimateTorrent_ResumeTorrent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AllTorrentStatus",
			Handler:       _UltimateTorrent_AllTorrentStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ultimate_torrent.proto",
}
