// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipinfo.proto

package api

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//These are the error codes which can be returned by this endpoint. In case of an error, an ErrorResponse will be returned
type WhatsMyIpResponse_Error int32

const (
	WhatsMyIpResponse_INVALID_IP     WhatsMyIpResponse_Error = 0
	WhatsMyIpResponse_INTERNAL_ERROR WhatsMyIpResponse_Error = 1
)

var WhatsMyIpResponse_Error_name = map[int32]string{
	0: "INVALID_IP",
	1: "INTERNAL_ERROR",
}

var WhatsMyIpResponse_Error_value = map[string]int32{
	"INVALID_IP":     0,
	"INTERNAL_ERROR": 1,
}

func (x WhatsMyIpResponse_Error) String() string {
	return proto.EnumName(WhatsMyIpResponse_Error_name, int32(x))
}

func (WhatsMyIpResponse_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{1, 0}
}

//These are the error codes which can be returned by this endpoint. In case of an error, an ErrorResponse will be returned
type IpLookupResponse_Error int32

const (
	IpLookupResponse_INVALID_IP     IpLookupResponse_Error = 0
	IpLookupResponse_INTERNAL_ERROR IpLookupResponse_Error = 1
)

var IpLookupResponse_Error_name = map[int32]string{
	0: "INVALID_IP",
	1: "INTERNAL_ERROR",
}

var IpLookupResponse_Error_value = map[string]int32{
	"INVALID_IP":     0,
	"INTERNAL_ERROR": 1,
}

func (x IpLookupResponse_Error) String() string {
	return proto.EnumName(IpLookupResponse_Error_name, int32(x))
}

func (IpLookupResponse_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{3, 0}
}

//
//GET /ip
type WhatsMyIpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhatsMyIpRequest) Reset()         { *m = WhatsMyIpRequest{} }
func (m *WhatsMyIpRequest) String() string { return proto.CompactTextString(m) }
func (*WhatsMyIpRequest) ProtoMessage()    {}
func (*WhatsMyIpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{0}
}

func (m *WhatsMyIpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhatsMyIpRequest.Unmarshal(m, b)
}
func (m *WhatsMyIpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhatsMyIpRequest.Marshal(b, m, deterministic)
}
func (m *WhatsMyIpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhatsMyIpRequest.Merge(m, src)
}
func (m *WhatsMyIpRequest) XXX_Size() int {
	return xxx_messageInfo_WhatsMyIpRequest.Size(m)
}
func (m *WhatsMyIpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhatsMyIpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhatsMyIpRequest proto.InternalMessageInfo

type WhatsMyIpResponse struct {
	Ip                   string    `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	ForwardedFor         string    `protobuf:"bytes,2,opt,name=forwardedFor,proto3" json:"forwardedFor,omitempty"`
	Hostnames            []string  `protobuf:"bytes,3,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	Geolocation          *Location `protobuf:"bytes,4,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WhatsMyIpResponse) Reset()         { *m = WhatsMyIpResponse{} }
func (m *WhatsMyIpResponse) String() string { return proto.CompactTextString(m) }
func (*WhatsMyIpResponse) ProtoMessage()    {}
func (*WhatsMyIpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{1}
}

func (m *WhatsMyIpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhatsMyIpResponse.Unmarshal(m, b)
}
func (m *WhatsMyIpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhatsMyIpResponse.Marshal(b, m, deterministic)
}
func (m *WhatsMyIpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhatsMyIpResponse.Merge(m, src)
}
func (m *WhatsMyIpResponse) XXX_Size() int {
	return xxx_messageInfo_WhatsMyIpResponse.Size(m)
}
func (m *WhatsMyIpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhatsMyIpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhatsMyIpResponse proto.InternalMessageInfo

func (m *WhatsMyIpResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *WhatsMyIpResponse) GetForwardedFor() string {
	if m != nil {
		return m.ForwardedFor
	}
	return ""
}

func (m *WhatsMyIpResponse) GetHostnames() []string {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

func (m *WhatsMyIpResponse) GetGeolocation() *Location {
	if m != nil {
		return m.Geolocation
	}
	return nil
}

//
//GET /ip/123.45.67.89
type IpLookupRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpLookupRequest) Reset()         { *m = IpLookupRequest{} }
func (m *IpLookupRequest) String() string { return proto.CompactTextString(m) }
func (*IpLookupRequest) ProtoMessage()    {}
func (*IpLookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{2}
}

func (m *IpLookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpLookupRequest.Unmarshal(m, b)
}
func (m *IpLookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpLookupRequest.Marshal(b, m, deterministic)
}
func (m *IpLookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpLookupRequest.Merge(m, src)
}
func (m *IpLookupRequest) XXX_Size() int {
	return xxx_messageInfo_IpLookupRequest.Size(m)
}
func (m *IpLookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IpLookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IpLookupRequest proto.InternalMessageInfo

func (m *IpLookupRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type IpLookupResponse struct {
	Ip                   string    `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Hostnames            []string  `protobuf:"bytes,2,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	Geolocation          *Location `protobuf:"bytes,3,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IpLookupResponse) Reset()         { *m = IpLookupResponse{} }
func (m *IpLookupResponse) String() string { return proto.CompactTextString(m) }
func (*IpLookupResponse) ProtoMessage()    {}
func (*IpLookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{3}
}

func (m *IpLookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpLookupResponse.Unmarshal(m, b)
}
func (m *IpLookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpLookupResponse.Marshal(b, m, deterministic)
}
func (m *IpLookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpLookupResponse.Merge(m, src)
}
func (m *IpLookupResponse) XXX_Size() int {
	return xxx_messageInfo_IpLookupResponse.Size(m)
}
func (m *IpLookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IpLookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IpLookupResponse proto.InternalMessageInfo

func (m *IpLookupResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IpLookupResponse) GetHostnames() []string {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

func (m *IpLookupResponse) GetGeolocation() *Location {
	if m != nil {
		return m.Geolocation
	}
	return nil
}

type Location struct {
	CountryCode          string   `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Lat                  float64  `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Long                 float64  `protobuf:"fixed64,5,opt,name=long,proto3" json:"long,omitempty"`
	Timezone             string   `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f776f3a264ad8cc, []int{4}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *Location) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Location) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Location) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Location) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.WhatsMyIpResponse_Error", WhatsMyIpResponse_Error_name, WhatsMyIpResponse_Error_value)
	proto.RegisterEnum("api.IpLookupResponse_Error", IpLookupResponse_Error_name, IpLookupResponse_Error_value)
	proto.RegisterType((*WhatsMyIpRequest)(nil), "api.WhatsMyIpRequest")
	proto.RegisterType((*WhatsMyIpResponse)(nil), "api.WhatsMyIpResponse")
	proto.RegisterType((*IpLookupRequest)(nil), "api.IpLookupRequest")
	proto.RegisterType((*IpLookupResponse)(nil), "api.IpLookupResponse")
	proto.RegisterType((*Location)(nil), "api.Location")
}

func init() { proto.RegisterFile("ipinfo.proto", fileDescriptor_3f776f3a264ad8cc) }

var fileDescriptor_3f776f3a264ad8cc = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x86, 0xed, 0x06, 0xc8, 0x0e, 0x88, 0xf3, 0x5c, 0x2d, 0xc6, 0x0b, 0xd8, 0x15, 0x89, 0x09,
	0x26, 0xfa, 0x04, 0x44, 0x31, 0x59, 0x32, 0xd1, 0x34, 0x46, 0x2f, 0xc9, 0xdc, 0x0a, 0x34, 0xc2,
	0x4e, 0xed, 0x4a, 0x0c, 0xbe, 0x8b, 0xbe, 0x8c, 0x2f, 0x66, 0x68, 0x86, 0x80, 0x89, 0xf1, 0xc6,
	0xbb, 0x73, 0xbe, 0xfe, 0x6d, 0xce, 0xd7, 0x16, 0x9a, 0x52, 0xc9, 0x7c, 0x4c, 0x3d, 0xa5, 0xc9,
	0x10, 0xba, 0x89, 0x92, 0x21, 0x82, 0xff, 0x38, 0x4d, 0x4c, 0x71, 0xb3, 0x8c, 0x14, 0x17, 0x2f,
	0x0b, 0x51, 0x98, 0xf0, 0x93, 0xc1, 0xd1, 0x16, 0x2c, 0x14, 0xe5, 0x85, 0xc0, 0x16, 0x38, 0x52,
	0x05, 0xac, 0xcd, 0xba, 0x1e, 0x77, 0xa4, 0xc2, 0x10, 0x9a, 0x63, 0xd2, 0xaf, 0x89, 0xce, 0x44,
	0x76, 0x4d, 0x3a, 0x70, 0xec, 0xca, 0x0e, 0xc3, 0x13, 0xf0, 0xa6, 0x54, 0x98, 0x3c, 0x99, 0x8b,
	0x22, 0x70, 0xdb, 0x6e, 0xd7, 0xe3, 0x1b, 0x80, 0x67, 0xd0, 0x98, 0x08, 0x9a, 0x51, 0x9a, 0x18,
	0x49, 0x79, 0x50, 0x69, 0xb3, 0x6e, 0xe3, 0xfc, 0xa0, 0x97, 0x28, 0xd9, 0x8b, 0x4b, 0xc8, 0xb7,
	0x13, 0xe1, 0x29, 0x54, 0x07, 0x5a, 0x93, 0xc6, 0x16, 0x40, 0x34, 0x7c, 0xe8, 0xc7, 0xd1, 0xd5,
	0x28, 0xba, 0xf3, 0xf7, 0x10, 0xa1, 0x15, 0x0d, 0xef, 0x07, 0x7c, 0xd8, 0x8f, 0x47, 0x03, 0xce,
	0x6f, 0xb9, 0xcf, 0xc2, 0x0e, 0x1c, 0x46, 0x2a, 0x26, 0x7a, 0x5e, 0xac, 0xc5, 0x7e, 0x2a, 0x84,
	0x1f, 0x0c, 0xfc, 0x4d, 0xe6, 0x17, 0xcf, 0x1d, 0x07, 0xe7, 0x0f, 0x07, 0xf7, 0x7f, 0x1d, 0xde,
	0x19, 0xd4, 0xd7, 0xc7, 0x60, 0x07, 0x9a, 0x29, 0x2d, 0x72, 0xa3, 0x97, 0xa3, 0x94, 0x32, 0x51,
	0x8e, 0xd8, 0x28, 0xd9, 0x25, 0x65, 0x02, 0x03, 0xd8, 0x2f, 0xdb, 0xf2, 0x39, 0xd6, 0x2d, 0x22,
	0x54, 0x52, 0x69, 0x96, 0x76, 0x40, 0x8f, 0xdb, 0x1a, 0x7d, 0x70, 0x67, 0x89, 0xb1, 0xf7, 0xce,
	0xf8, 0xaa, 0x5c, 0xa5, 0x66, 0x94, 0x4f, 0x82, 0xaa, 0x45, 0xb6, 0xc6, 0x63, 0xa8, 0x1b, 0x39,
	0x17, 0x6f, 0x94, 0x8b, 0xa0, 0x66, 0x77, 0x7f, 0xf7, 0x4f, 0x35, 0xfb, 0x93, 0x2e, 0xbe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x06, 0xdb, 0x9f, 0x59, 0x02, 0x00, 0x00,
}
