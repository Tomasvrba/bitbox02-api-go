// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bitboxbase.proto

package messages

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

type BitBoxBaseHeartbeatRequest_StateCode int32

const (
	BitBoxBaseHeartbeatRequest_IDLE    BitBoxBaseHeartbeatRequest_StateCode = 0
	BitBoxBaseHeartbeatRequest_WORKING BitBoxBaseHeartbeatRequest_StateCode = 1
	BitBoxBaseHeartbeatRequest_WARNING BitBoxBaseHeartbeatRequest_StateCode = 2
	BitBoxBaseHeartbeatRequest_ERROR   BitBoxBaseHeartbeatRequest_StateCode = 3
)

var BitBoxBaseHeartbeatRequest_StateCode_name = map[int32]string{
	0: "IDLE",
	1: "WORKING",
	2: "WARNING",
	3: "ERROR",
}

var BitBoxBaseHeartbeatRequest_StateCode_value = map[string]int32{
	"IDLE":    0,
	"WORKING": 1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x BitBoxBaseHeartbeatRequest_StateCode) String() string {
	return proto.EnumName(BitBoxBaseHeartbeatRequest_StateCode_name, int32(x))
}

func (BitBoxBaseHeartbeatRequest_StateCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{0, 0}
}

type BitBoxBaseHeartbeatRequest_DescriptionCode int32

const (
	BitBoxBaseHeartbeatRequest_EMPTY              BitBoxBaseHeartbeatRequest_DescriptionCode = 0
	BitBoxBaseHeartbeatRequest_INITIAL_BLOCK_SYNC BitBoxBaseHeartbeatRequest_DescriptionCode = 1
	BitBoxBaseHeartbeatRequest_DOWNLOAD_UPDATE    BitBoxBaseHeartbeatRequest_DescriptionCode = 2
	BitBoxBaseHeartbeatRequest_OUT_OF_DISK_SPACE  BitBoxBaseHeartbeatRequest_DescriptionCode = 3
)

var BitBoxBaseHeartbeatRequest_DescriptionCode_name = map[int32]string{
	0: "EMPTY",
	1: "INITIAL_BLOCK_SYNC",
	2: "DOWNLOAD_UPDATE",
	3: "OUT_OF_DISK_SPACE",
}

var BitBoxBaseHeartbeatRequest_DescriptionCode_value = map[string]int32{
	"EMPTY":              0,
	"INITIAL_BLOCK_SYNC": 1,
	"DOWNLOAD_UPDATE":    2,
	"OUT_OF_DISK_SPACE":  3,
}

func (x BitBoxBaseHeartbeatRequest_DescriptionCode) String() string {
	return proto.EnumName(BitBoxBaseHeartbeatRequest_DescriptionCode_name, int32(x))
}

func (BitBoxBaseHeartbeatRequest_DescriptionCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{0, 1}
}

type BitBoxBaseSetConfigRequest_StatusLedMode int32

const (
	BitBoxBaseSetConfigRequest_LED_ALWAYS     BitBoxBaseSetConfigRequest_StatusLedMode = 0
	BitBoxBaseSetConfigRequest_LED_ON_WORKING BitBoxBaseSetConfigRequest_StatusLedMode = 1
	BitBoxBaseSetConfigRequest_LED_ON_WARNING BitBoxBaseSetConfigRequest_StatusLedMode = 2
	BitBoxBaseSetConfigRequest_LED_ON_ERROR   BitBoxBaseSetConfigRequest_StatusLedMode = 3
)

var BitBoxBaseSetConfigRequest_StatusLedMode_name = map[int32]string{
	0: "LED_ALWAYS",
	1: "LED_ON_WORKING",
	2: "LED_ON_WARNING",
	3: "LED_ON_ERROR",
}

var BitBoxBaseSetConfigRequest_StatusLedMode_value = map[string]int32{
	"LED_ALWAYS":     0,
	"LED_ON_WORKING": 1,
	"LED_ON_WARNING": 2,
	"LED_ON_ERROR":   3,
}

func (x BitBoxBaseSetConfigRequest_StatusLedMode) String() string {
	return proto.EnumName(BitBoxBaseSetConfigRequest_StatusLedMode_name, int32(x))
}

func (BitBoxBaseSetConfigRequest_StatusLedMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{2, 0}
}

type BitBoxBaseSetConfigRequest_StatusScreenMode int32

const (
	BitBoxBaseSetConfigRequest_SCREEN_ALWAYS     BitBoxBaseSetConfigRequest_StatusScreenMode = 0
	BitBoxBaseSetConfigRequest_SCREEN_ON_WORKING BitBoxBaseSetConfigRequest_StatusScreenMode = 1
	BitBoxBaseSetConfigRequest_SCREEN_ON_WARNING BitBoxBaseSetConfigRequest_StatusScreenMode = 2
	BitBoxBaseSetConfigRequest_SCREEN_ON_ERROR   BitBoxBaseSetConfigRequest_StatusScreenMode = 3
)

var BitBoxBaseSetConfigRequest_StatusScreenMode_name = map[int32]string{
	0: "SCREEN_ALWAYS",
	1: "SCREEN_ON_WORKING",
	2: "SCREEN_ON_WARNING",
	3: "SCREEN_ON_ERROR",
}

var BitBoxBaseSetConfigRequest_StatusScreenMode_value = map[string]int32{
	"SCREEN_ALWAYS":     0,
	"SCREEN_ON_WORKING": 1,
	"SCREEN_ON_WARNING": 2,
	"SCREEN_ON_ERROR":   3,
}

func (x BitBoxBaseSetConfigRequest_StatusScreenMode) String() string {
	return proto.EnumName(BitBoxBaseSetConfigRequest_StatusScreenMode_name, int32(x))
}

func (BitBoxBaseSetConfigRequest_StatusScreenMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{2, 1}
}

// Should be sent every X seconds (TBD) unless the firmware already is busy with a command.
type BitBoxBaseHeartbeatRequest struct {
	StateCode            BitBoxBaseHeartbeatRequest_StateCode       `protobuf:"varint,1,opt,name=state_code,json=stateCode,proto3,enum=BitBoxBaseHeartbeatRequest_StateCode" json:"state_code,omitempty"`
	DescriptionCode      BitBoxBaseHeartbeatRequest_DescriptionCode `protobuf:"varint,2,opt,name=description_code,json=descriptionCode,proto3,enum=BitBoxBaseHeartbeatRequest_DescriptionCode" json:"description_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BitBoxBaseHeartbeatRequest) Reset()         { *m = BitBoxBaseHeartbeatRequest{} }
func (m *BitBoxBaseHeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*BitBoxBaseHeartbeatRequest) ProtoMessage()    {}
func (*BitBoxBaseHeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{0}
}

func (m *BitBoxBaseHeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitBoxBaseHeartbeatRequest.Unmarshal(m, b)
}
func (m *BitBoxBaseHeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitBoxBaseHeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *BitBoxBaseHeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitBoxBaseHeartbeatRequest.Merge(m, src)
}
func (m *BitBoxBaseHeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_BitBoxBaseHeartbeatRequest.Size(m)
}
func (m *BitBoxBaseHeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitBoxBaseHeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitBoxBaseHeartbeatRequest proto.InternalMessageInfo

func (m *BitBoxBaseHeartbeatRequest) GetStateCode() BitBoxBaseHeartbeatRequest_StateCode {
	if m != nil {
		return m.StateCode
	}
	return BitBoxBaseHeartbeatRequest_IDLE
}

func (m *BitBoxBaseHeartbeatRequest) GetDescriptionCode() BitBoxBaseHeartbeatRequest_DescriptionCode {
	if m != nil {
		return m.DescriptionCode
	}
	return BitBoxBaseHeartbeatRequest_EMPTY
}

// This will display the first 20 characters of the base32 encoded version of
// the provided msg
type BitBoxBaseConfirmPairingRequest struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitBoxBaseConfirmPairingRequest) Reset()         { *m = BitBoxBaseConfirmPairingRequest{} }
func (m *BitBoxBaseConfirmPairingRequest) String() string { return proto.CompactTextString(m) }
func (*BitBoxBaseConfirmPairingRequest) ProtoMessage()    {}
func (*BitBoxBaseConfirmPairingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{1}
}

func (m *BitBoxBaseConfirmPairingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitBoxBaseConfirmPairingRequest.Unmarshal(m, b)
}
func (m *BitBoxBaseConfirmPairingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitBoxBaseConfirmPairingRequest.Marshal(b, m, deterministic)
}
func (m *BitBoxBaseConfirmPairingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitBoxBaseConfirmPairingRequest.Merge(m, src)
}
func (m *BitBoxBaseConfirmPairingRequest) XXX_Size() int {
	return xxx_messageInfo_BitBoxBaseConfirmPairingRequest.Size(m)
}
func (m *BitBoxBaseConfirmPairingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitBoxBaseConfirmPairingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitBoxBaseConfirmPairingRequest proto.InternalMessageInfo

func (m *BitBoxBaseConfirmPairingRequest) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// Optional fields can be represented by a "oneof" with only one field in it.
// All fields are technically optional. But in reality the default value for the type will be set.
// It is therefore impossible to distinguish between the default value and if the value wasn't set.
// So any fields that have a default value which also is a valid value can use this method to send
// an empty value.
type BitBoxBaseSetConfigRequest struct {
	StatusLedMode    BitBoxBaseSetConfigRequest_StatusLedMode    `protobuf:"varint,1,opt,name=status_led_mode,json=statusLedMode,proto3,enum=BitBoxBaseSetConfigRequest_StatusLedMode" json:"status_led_mode,omitempty"`
	StatusScreenMode BitBoxBaseSetConfigRequest_StatusScreenMode `protobuf:"varint,2,opt,name=status_screen_mode,json=statusScreenMode,proto3,enum=BitBoxBaseSetConfigRequest_StatusScreenMode" json:"status_screen_mode,omitempty"`
	// 0.0.0.0 which is the default value of ip is also a valid IP, use the oneof-trick to determine
	// if IP wasn't set in the message.
	//
	// Types that are valid to be assigned to IpOption:
	//	*BitBoxBaseSetConfigRequest_Ip
	IpOption             isBitBoxBaseSetConfigRequest_IpOption `protobuf_oneof:"ip_option"`
	Hostname             string                                `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *BitBoxBaseSetConfigRequest) Reset()         { *m = BitBoxBaseSetConfigRequest{} }
func (m *BitBoxBaseSetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*BitBoxBaseSetConfigRequest) ProtoMessage()    {}
func (*BitBoxBaseSetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{2}
}

func (m *BitBoxBaseSetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitBoxBaseSetConfigRequest.Unmarshal(m, b)
}
func (m *BitBoxBaseSetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitBoxBaseSetConfigRequest.Marshal(b, m, deterministic)
}
func (m *BitBoxBaseSetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitBoxBaseSetConfigRequest.Merge(m, src)
}
func (m *BitBoxBaseSetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_BitBoxBaseSetConfigRequest.Size(m)
}
func (m *BitBoxBaseSetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitBoxBaseSetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitBoxBaseSetConfigRequest proto.InternalMessageInfo

func (m *BitBoxBaseSetConfigRequest) GetStatusLedMode() BitBoxBaseSetConfigRequest_StatusLedMode {
	if m != nil {
		return m.StatusLedMode
	}
	return BitBoxBaseSetConfigRequest_LED_ALWAYS
}

func (m *BitBoxBaseSetConfigRequest) GetStatusScreenMode() BitBoxBaseSetConfigRequest_StatusScreenMode {
	if m != nil {
		return m.StatusScreenMode
	}
	return BitBoxBaseSetConfigRequest_SCREEN_ALWAYS
}

type isBitBoxBaseSetConfigRequest_IpOption interface {
	isBitBoxBaseSetConfigRequest_IpOption()
}

type BitBoxBaseSetConfigRequest_Ip struct {
	Ip []byte `protobuf:"bytes,3,opt,name=ip,proto3,oneof"`
}

func (*BitBoxBaseSetConfigRequest_Ip) isBitBoxBaseSetConfigRequest_IpOption() {}

func (m *BitBoxBaseSetConfigRequest) GetIpOption() isBitBoxBaseSetConfigRequest_IpOption {
	if m != nil {
		return m.IpOption
	}
	return nil
}

func (m *BitBoxBaseSetConfigRequest) GetIp() []byte {
	if x, ok := m.GetIpOption().(*BitBoxBaseSetConfigRequest_Ip); ok {
		return x.Ip
	}
	return nil
}

func (m *BitBoxBaseSetConfigRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BitBoxBaseSetConfigRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BitBoxBaseSetConfigRequest_Ip)(nil),
	}
}

type BitBoxBaseDisplayStatusRequest struct {
	Duration             uint32   `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitBoxBaseDisplayStatusRequest) Reset()         { *m = BitBoxBaseDisplayStatusRequest{} }
func (m *BitBoxBaseDisplayStatusRequest) String() string { return proto.CompactTextString(m) }
func (*BitBoxBaseDisplayStatusRequest) ProtoMessage()    {}
func (*BitBoxBaseDisplayStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{3}
}

func (m *BitBoxBaseDisplayStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitBoxBaseDisplayStatusRequest.Unmarshal(m, b)
}
func (m *BitBoxBaseDisplayStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitBoxBaseDisplayStatusRequest.Marshal(b, m, deterministic)
}
func (m *BitBoxBaseDisplayStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitBoxBaseDisplayStatusRequest.Merge(m, src)
}
func (m *BitBoxBaseDisplayStatusRequest) XXX_Size() int {
	return xxx_messageInfo_BitBoxBaseDisplayStatusRequest.Size(m)
}
func (m *BitBoxBaseDisplayStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitBoxBaseDisplayStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitBoxBaseDisplayStatusRequest proto.InternalMessageInfo

func (m *BitBoxBaseDisplayStatusRequest) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type BitBoxBaseRequest struct {
	// Types that are valid to be assigned to Request:
	//	*BitBoxBaseRequest_Heartbeat
	//	*BitBoxBaseRequest_SetConfig
	//	*BitBoxBaseRequest_ConfirmPairing
	//	*BitBoxBaseRequest_DisplayStatus
	Request              isBitBoxBaseRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BitBoxBaseRequest) Reset()         { *m = BitBoxBaseRequest{} }
func (m *BitBoxBaseRequest) String() string { return proto.CompactTextString(m) }
func (*BitBoxBaseRequest) ProtoMessage()    {}
func (*BitBoxBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d283a3459010a, []int{4}
}

func (m *BitBoxBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitBoxBaseRequest.Unmarshal(m, b)
}
func (m *BitBoxBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitBoxBaseRequest.Marshal(b, m, deterministic)
}
func (m *BitBoxBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitBoxBaseRequest.Merge(m, src)
}
func (m *BitBoxBaseRequest) XXX_Size() int {
	return xxx_messageInfo_BitBoxBaseRequest.Size(m)
}
func (m *BitBoxBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitBoxBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitBoxBaseRequest proto.InternalMessageInfo

type isBitBoxBaseRequest_Request interface {
	isBitBoxBaseRequest_Request()
}

type BitBoxBaseRequest_Heartbeat struct {
	Heartbeat *BitBoxBaseHeartbeatRequest `protobuf:"bytes,1,opt,name=heartbeat,proto3,oneof"`
}

type BitBoxBaseRequest_SetConfig struct {
	SetConfig *BitBoxBaseSetConfigRequest `protobuf:"bytes,2,opt,name=set_config,json=setConfig,proto3,oneof"`
}

type BitBoxBaseRequest_ConfirmPairing struct {
	ConfirmPairing *BitBoxBaseConfirmPairingRequest `protobuf:"bytes,3,opt,name=confirm_pairing,json=confirmPairing,proto3,oneof"`
}

type BitBoxBaseRequest_DisplayStatus struct {
	DisplayStatus *BitBoxBaseDisplayStatusRequest `protobuf:"bytes,4,opt,name=display_status,json=displayStatus,proto3,oneof"`
}

func (*BitBoxBaseRequest_Heartbeat) isBitBoxBaseRequest_Request() {}

func (*BitBoxBaseRequest_SetConfig) isBitBoxBaseRequest_Request() {}

func (*BitBoxBaseRequest_ConfirmPairing) isBitBoxBaseRequest_Request() {}

func (*BitBoxBaseRequest_DisplayStatus) isBitBoxBaseRequest_Request() {}

func (m *BitBoxBaseRequest) GetRequest() isBitBoxBaseRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BitBoxBaseRequest) GetHeartbeat() *BitBoxBaseHeartbeatRequest {
	if x, ok := m.GetRequest().(*BitBoxBaseRequest_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (m *BitBoxBaseRequest) GetSetConfig() *BitBoxBaseSetConfigRequest {
	if x, ok := m.GetRequest().(*BitBoxBaseRequest_SetConfig); ok {
		return x.SetConfig
	}
	return nil
}

func (m *BitBoxBaseRequest) GetConfirmPairing() *BitBoxBaseConfirmPairingRequest {
	if x, ok := m.GetRequest().(*BitBoxBaseRequest_ConfirmPairing); ok {
		return x.ConfirmPairing
	}
	return nil
}

func (m *BitBoxBaseRequest) GetDisplayStatus() *BitBoxBaseDisplayStatusRequest {
	if x, ok := m.GetRequest().(*BitBoxBaseRequest_DisplayStatus); ok {
		return x.DisplayStatus
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BitBoxBaseRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BitBoxBaseRequest_Heartbeat)(nil),
		(*BitBoxBaseRequest_SetConfig)(nil),
		(*BitBoxBaseRequest_ConfirmPairing)(nil),
		(*BitBoxBaseRequest_DisplayStatus)(nil),
	}
}

func init() {
	proto.RegisterEnum("BitBoxBaseHeartbeatRequest_StateCode", BitBoxBaseHeartbeatRequest_StateCode_name, BitBoxBaseHeartbeatRequest_StateCode_value)
	proto.RegisterEnum("BitBoxBaseHeartbeatRequest_DescriptionCode", BitBoxBaseHeartbeatRequest_DescriptionCode_name, BitBoxBaseHeartbeatRequest_DescriptionCode_value)
	proto.RegisterEnum("BitBoxBaseSetConfigRequest_StatusLedMode", BitBoxBaseSetConfigRequest_StatusLedMode_name, BitBoxBaseSetConfigRequest_StatusLedMode_value)
	proto.RegisterEnum("BitBoxBaseSetConfigRequest_StatusScreenMode", BitBoxBaseSetConfigRequest_StatusScreenMode_name, BitBoxBaseSetConfigRequest_StatusScreenMode_value)
	proto.RegisterType((*BitBoxBaseHeartbeatRequest)(nil), "BitBoxBaseHeartbeatRequest")
	proto.RegisterType((*BitBoxBaseConfirmPairingRequest)(nil), "BitBoxBaseConfirmPairingRequest")
	proto.RegisterType((*BitBoxBaseSetConfigRequest)(nil), "BitBoxBaseSetConfigRequest")
	proto.RegisterType((*BitBoxBaseDisplayStatusRequest)(nil), "BitBoxBaseDisplayStatusRequest")
	proto.RegisterType((*BitBoxBaseRequest)(nil), "BitBoxBaseRequest")
}

func init() { proto.RegisterFile("bitboxbase.proto", fileDescriptor_9c0d283a3459010a) }

var fileDescriptor_9c0d283a3459010a = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x0d, 0xa1, 0x5b, 0x9b, 0x4b, 0x01, 0xe3, 0xaa, 0x53, 0xd5, 0x49, 0x6b, 0x15, 0x69, 0x52,
	0xa7, 0x4d, 0x3c, 0xd0, 0xb7, 0xad, 0x2f, 0x81, 0x64, 0x03, 0x91, 0x26, 0xcc, 0xa1, 0xab, 0xd8,
	0x8b, 0x17, 0x88, 0xd7, 0x46, 0x2a, 0x24, 0x8b, 0x8d, 0xd4, 0xfd, 0xd5, 0xfe, 0x63, 0x3f, 0xb0,
	0xcf, 0x99, 0xe2, 0x10, 0x42, 0x90, 0x4a, 0xdf, 0x72, 0x8f, 0x7d, 0xce, 0xb5, 0xcf, 0xb9, 0x0e,
	0xa0, 0x69, 0x28, 0xa6, 0xd1, 0xe3, 0xd4, 0xe7, 0xac, 0x1d, 0x27, 0x91, 0x88, 0xf4, 0x7f, 0x2a,
	0x9c, 0x76, 0x43, 0xd1, 0x8d, 0x1e, 0xbb, 0x3e, 0x67, 0x7d, 0xe6, 0x27, 0x62, 0xca, 0x7c, 0x41,
	0xd8, 0xaf, 0x25, 0xe3, 0x02, 0x9b, 0x00, 0x5c, 0xf8, 0x82, 0xd1, 0x59, 0x14, 0xb0, 0x93, 0xca,
	0x79, 0xe5, 0xa2, 0xd1, 0x79, 0xdb, 0x7e, 0x9a, 0xd0, 0xf6, 0xd2, 0xdd, 0xbd, 0x28, 0x60, 0x44,
	0xe3, 0xf9, 0x27, 0xfe, 0x06, 0x28, 0x60, 0x7c, 0x96, 0x84, 0xb1, 0x08, 0xa3, 0x45, 0xa6, 0xa5,
	0x4a, 0xad, 0xf7, 0xbb, 0xb4, 0xcc, 0x82, 0x23, 0x15, 0x9b, 0x41, 0x19, 0xd0, 0x3f, 0x82, 0xb6,
	0xee, 0x87, 0x0f, 0x60, 0x6f, 0x60, 0xda, 0x16, 0x52, 0x70, 0x0d, 0xf6, 0x6f, 0x5d, 0x32, 0x1c,
	0x38, 0x5f, 0x50, 0x45, 0x16, 0x06, 0x71, 0xd2, 0x42, 0xc5, 0x1a, 0xbc, 0xb0, 0x08, 0x71, 0x09,
	0xaa, 0xea, 0x3f, 0xa0, 0xb9, 0xa5, 0x2f, 0x57, 0xaf, 0x47, 0xe3, 0x09, 0x52, 0xf0, 0x2b, 0xc0,
	0x03, 0x67, 0x30, 0x1e, 0x18, 0x36, 0xed, 0xda, 0x6e, 0x6f, 0x48, 0xbd, 0x89, 0xd3, 0x43, 0x15,
	0x7c, 0x04, 0x4d, 0xd3, 0xbd, 0x75, 0x6c, 0xd7, 0x30, 0xe9, 0xcd, 0xc8, 0x34, 0xc6, 0x16, 0x52,
	0xf1, 0x31, 0xb4, 0xdc, 0x9b, 0x31, 0x75, 0x3f, 0x53, 0x73, 0xe0, 0x0d, 0xa9, 0x37, 0x32, 0x7a,
	0x16, 0xaa, 0xea, 0x97, 0x70, 0x56, 0x5c, 0xae, 0x17, 0x2d, 0x7e, 0x86, 0xc9, 0x7c, 0xe4, 0x87,
	0x49, 0xb8, 0xb8, 0xcb, 0xed, 0x45, 0x50, 0x9d, 0xf3, 0x3b, 0xe9, 0xeb, 0x21, 0x49, 0x3f, 0xf5,
	0xbf, 0xd5, 0xcd, 0x3c, 0x3c, 0x26, 0x24, 0x71, 0x4d, 0xf8, 0x0a, 0xcd, 0xd4, 0xd6, 0x25, 0xa7,
	0x0f, 0x2c, 0xa0, 0xf3, 0x22, 0x94, 0x77, 0xed, 0xa7, 0x59, 0x32, 0x94, 0x25, 0xb7, 0x59, 0x70,
	0x9d, 0xda, 0x58, 0xe7, 0x9b, 0x25, 0xfe, 0x0e, 0x78, 0x25, 0xc9, 0x67, 0x09, 0x63, 0x8b, 0x4c,
	0x35, 0x8b, 0xe7, 0xc3, 0xf3, 0xaa, 0x9e, 0x24, 0x49, 0x61, 0xc4, 0xb7, 0x10, 0x8c, 0x40, 0x0d,
	0xe3, 0x93, 0x6a, 0x7a, 0xbd, 0xbe, 0x42, 0xd4, 0x30, 0xc6, 0xa7, 0x70, 0x70, 0x1f, 0x71, 0xb1,
	0xf0, 0xe7, 0xec, 0x64, 0xef, 0xbc, 0x72, 0xa1, 0x91, 0x75, 0xad, 0x4f, 0xa0, 0x5e, 0x3a, 0x29,
	0x6e, 0x00, 0xd8, 0x96, 0x49, 0x0d, 0xfb, 0xd6, 0x98, 0x78, 0x48, 0xc1, 0x18, 0x1a, 0x69, 0xed,
	0x3a, 0xb4, 0xc8, 0x77, 0x03, 0x5b, 0xc7, 0x8c, 0xe0, 0x70, 0x85, 0xe5, 0x69, 0xdf, 0x03, 0xda,
	0x3e, 0x2e, 0x6e, 0x41, 0xdd, 0xeb, 0x11, 0xcb, 0x72, 0x8a, 0x06, 0xc7, 0xd0, 0x5a, 0x41, 0xa5,
	0x1e, 0x65, 0x78, 0xdd, 0xe6, 0x08, 0x9a, 0x05, 0xbc, 0xea, 0xd4, 0xad, 0x81, 0x16, 0xc6, 0x34,
	0x92, 0x53, 0xa5, 0x5f, 0xc1, 0x9b, 0xc2, 0x40, 0x33, 0xe4, 0xf1, 0x83, 0xff, 0x3b, 0x3b, 0x47,
	0x1e, 0xe8, 0x29, 0x1c, 0x04, 0xcb, 0xc4, 0x4f, 0x77, 0xcb, 0x24, 0xeb, 0x64, 0x5d, 0xeb, 0x7f,
	0x54, 0x68, 0x15, 0xf4, 0x9c, 0xf1, 0x09, 0xb4, 0xfb, 0xfc, 0xa5, 0x48, 0x4a, 0xad, 0xf3, 0x7a,
	0xc7, 0x2b, 0xea, 0x2b, 0xa4, 0xd8, 0x8f, 0xaf, 0x00, 0x38, 0x13, 0x74, 0x26, 0x83, 0x94, 0x21,
	0x97, 0xd9, 0xdb, 0x21, 0xa7, 0x6c, 0x9e, 0x63, 0x78, 0x08, 0xcd, 0x59, 0x36, 0xc7, 0x34, 0xce,
	0x06, 0x59, 0x66, 0x5b, 0xeb, 0x9c, 0xb7, 0x9f, 0x99, 0xf4, 0xbe, 0x42, 0x1a, 0xb3, 0xd2, 0x02,
	0xee, 0x43, 0x23, 0xc8, 0x1c, 0xa1, 0xd9, 0xdc, 0xc8, 0x79, 0xa8, 0x75, 0xce, 0xda, 0xbb, 0x2d,
	0xeb, 0x2b, 0xa4, 0x1e, 0x6c, 0xe2, 0x5d, 0x0d, 0xf6, 0x93, 0x6c, 0x6d, 0xfa, 0x52, 0xfe, 0xd5,
	0x2e, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x88, 0x15, 0xff, 0xe9, 0x04, 0x00, 0x00,
}
