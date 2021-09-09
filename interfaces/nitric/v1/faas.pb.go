// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faas/v1/faas.proto

package v1

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

// Messages the client is able to send to the server
type ClientMessage struct {
	// Client message ID, used to pair requests/responses
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*ClientMessage_InitRequest
	//	*ClientMessage_TriggerResponse
	Content              isClientMessage_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientMessage) Reset()         { *m = ClientMessage{} }
func (m *ClientMessage) String() string { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()    {}
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{0}
}

func (m *ClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMessage.Unmarshal(m, b)
}
func (m *ClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMessage.Marshal(b, m, deterministic)
}
func (m *ClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMessage.Merge(m, src)
}
func (m *ClientMessage) XXX_Size() int {
	return xxx_messageInfo_ClientMessage.Size(m)
}
func (m *ClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMessage proto.InternalMessageInfo

func (m *ClientMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isClientMessage_Content interface {
	isClientMessage_Content()
}

type ClientMessage_InitRequest struct {
	InitRequest *InitRequest `protobuf:"bytes,2,opt,name=init_request,json=initRequest,proto3,oneof"`
}

type ClientMessage_TriggerResponse struct {
	TriggerResponse *TriggerResponse `protobuf:"bytes,3,opt,name=trigger_response,json=triggerResponse,proto3,oneof"`
}

func (*ClientMessage_InitRequest) isClientMessage_Content() {}

func (*ClientMessage_TriggerResponse) isClientMessage_Content() {}

func (m *ClientMessage) GetContent() isClientMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ClientMessage) GetInitRequest() *InitRequest {
	if x, ok := m.GetContent().(*ClientMessage_InitRequest); ok {
		return x.InitRequest
	}
	return nil
}

func (m *ClientMessage) GetTriggerResponse() *TriggerResponse {
	if x, ok := m.GetContent().(*ClientMessage_TriggerResponse); ok {
		return x.TriggerResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientMessage_InitRequest)(nil),
		(*ClientMessage_TriggerResponse)(nil),
	}
}

// Messages the server is able to send to the client
type ServerMessage struct {
	// Server message ID, used to pair requests/responses
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*ServerMessage_InitResponse
	//	*ServerMessage_TriggerRequest
	Content              isServerMessage_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{1}
}

func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (m *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(m, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

func (m *ServerMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isServerMessage_Content interface {
	isServerMessage_Content()
}

type ServerMessage_InitResponse struct {
	InitResponse *InitResponse `protobuf:"bytes,2,opt,name=init_response,json=initResponse,proto3,oneof"`
}

type ServerMessage_TriggerRequest struct {
	TriggerRequest *TriggerRequest `protobuf:"bytes,3,opt,name=trigger_request,json=triggerRequest,proto3,oneof"`
}

func (*ServerMessage_InitResponse) isServerMessage_Content() {}

func (*ServerMessage_TriggerRequest) isServerMessage_Content() {}

func (m *ServerMessage) GetContent() isServerMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ServerMessage) GetInitResponse() *InitResponse {
	if x, ok := m.GetContent().(*ServerMessage_InitResponse); ok {
		return x.InitResponse
	}
	return nil
}

func (m *ServerMessage) GetTriggerRequest() *TriggerRequest {
	if x, ok := m.GetContent().(*ServerMessage_TriggerRequest); ok {
		return x.TriggerRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerMessage_InitResponse)(nil),
		(*ServerMessage_TriggerRequest)(nil),
	}
}

// Placeholder message
type InitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{2}
}

func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (m *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(m, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

// Placeholder message
type InitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitResponse) Reset()         { *m = InitResponse{} }
func (m *InitResponse) String() string { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()    {}
func (*InitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{3}
}

func (m *InitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResponse.Unmarshal(m, b)
}
func (m *InitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResponse.Marshal(b, m, deterministic)
}
func (m *InitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResponse.Merge(m, src)
}
func (m *InitResponse) XXX_Size() int {
	return xxx_messageInfo_InitResponse.Size(m)
}
func (m *InitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitResponse proto.InternalMessageInfo

// The server has a trigger for the client to handle
type TriggerRequest struct {
	// The data in the trigger
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Should we supply a mime type for the data?
	// Or rely on context?
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// The context of the trigger
	//
	// Types that are valid to be assigned to Context:
	//	*TriggerRequest_Http
	//	*TriggerRequest_Topic
	Context              isTriggerRequest_Context `protobuf_oneof:"context"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{4}
}

func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (m *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(m, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

func (m *TriggerRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TriggerRequest) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

type isTriggerRequest_Context interface {
	isTriggerRequest_Context()
}

type TriggerRequest_Http struct {
	Http *HttpTriggerContext `protobuf:"bytes,3,opt,name=http,proto3,oneof"`
}

type TriggerRequest_Topic struct {
	Topic *TopicTriggerContext `protobuf:"bytes,4,opt,name=topic,proto3,oneof"`
}

func (*TriggerRequest_Http) isTriggerRequest_Context() {}

func (*TriggerRequest_Topic) isTriggerRequest_Context() {}

func (m *TriggerRequest) GetContext() isTriggerRequest_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TriggerRequest) GetHttp() *HttpTriggerContext {
	if x, ok := m.GetContext().(*TriggerRequest_Http); ok {
		return x.Http
	}
	return nil
}

func (m *TriggerRequest) GetTopic() *TopicTriggerContext {
	if x, ok := m.GetContext().(*TriggerRequest_Topic); ok {
		return x.Topic
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TriggerRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TriggerRequest_Http)(nil),
		(*TriggerRequest_Topic)(nil),
	}
}

type HeaderValue struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValue) Reset()         { *m = HeaderValue{} }
func (m *HeaderValue) String() string { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()    {}
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{5}
}

func (m *HeaderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValue.Unmarshal(m, b)
}
func (m *HeaderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValue.Marshal(b, m, deterministic)
}
func (m *HeaderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValue.Merge(m, src)
}
func (m *HeaderValue) XXX_Size() int {
	return xxx_messageInfo_HeaderValue.Size(m)
}
func (m *HeaderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValue proto.InternalMessageInfo

func (m *HeaderValue) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type HttpTriggerContext struct {
	// The request method
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// The path of the request
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// The old request headers (preserving for backwards compatibility)
	// TODO: Remove in 1.0
	HeadersOld map[string]string `protobuf:"bytes,3,rep,name=headers_old,json=headersOld,proto3" json:"headers_old,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	// The query params (if parseable by the membrane)
	QueryParams map[string]string `protobuf:"bytes,4,rep,name=query_params,json=queryParams,proto3" json:"query_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request headers
	Headers              map[string]*HeaderValue `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpTriggerContext) Reset()         { *m = HttpTriggerContext{} }
func (m *HttpTriggerContext) String() string { return proto.CompactTextString(m) }
func (*HttpTriggerContext) ProtoMessage()    {}
func (*HttpTriggerContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{6}
}

func (m *HttpTriggerContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpTriggerContext.Unmarshal(m, b)
}
func (m *HttpTriggerContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpTriggerContext.Marshal(b, m, deterministic)
}
func (m *HttpTriggerContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpTriggerContext.Merge(m, src)
}
func (m *HttpTriggerContext) XXX_Size() int {
	return xxx_messageInfo_HttpTriggerContext.Size(m)
}
func (m *HttpTriggerContext) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpTriggerContext.DiscardUnknown(m)
}

var xxx_messageInfo_HttpTriggerContext proto.InternalMessageInfo

func (m *HttpTriggerContext) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpTriggerContext) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Deprecated: Do not use.
func (m *HttpTriggerContext) GetHeadersOld() map[string]string {
	if m != nil {
		return m.HeadersOld
	}
	return nil
}

func (m *HttpTriggerContext) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

func (m *HttpTriggerContext) GetHeaders() map[string]*HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

type TopicTriggerContext struct {
	// The topic the message was published for
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicTriggerContext) Reset()         { *m = TopicTriggerContext{} }
func (m *TopicTriggerContext) String() string { return proto.CompactTextString(m) }
func (*TopicTriggerContext) ProtoMessage()    {}
func (*TopicTriggerContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{7}
}

func (m *TopicTriggerContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicTriggerContext.Unmarshal(m, b)
}
func (m *TopicTriggerContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicTriggerContext.Marshal(b, m, deterministic)
}
func (m *TopicTriggerContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicTriggerContext.Merge(m, src)
}
func (m *TopicTriggerContext) XXX_Size() int {
	return xxx_messageInfo_TopicTriggerContext.Size(m)
}
func (m *TopicTriggerContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicTriggerContext.DiscardUnknown(m)
}

var xxx_messageInfo_TopicTriggerContext proto.InternalMessageInfo

func (m *TopicTriggerContext) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// The worker has successfully processed a trigger
type TriggerResponse struct {
	// The data returned in the response
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// The context of the request response
	// Typically this will be one to one with the Trigger Context
	// i.e. if you receive http context you may return http context
	//
	// Types that are valid to be assigned to Context:
	//	*TriggerResponse_Http
	//	*TriggerResponse_Topic
	Context              isTriggerResponse_Context `protobuf_oneof:"context"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TriggerResponse) Reset()         { *m = TriggerResponse{} }
func (m *TriggerResponse) String() string { return proto.CompactTextString(m) }
func (*TriggerResponse) ProtoMessage()    {}
func (*TriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{8}
}

func (m *TriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerResponse.Unmarshal(m, b)
}
func (m *TriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerResponse.Marshal(b, m, deterministic)
}
func (m *TriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerResponse.Merge(m, src)
}
func (m *TriggerResponse) XXX_Size() int {
	return xxx_messageInfo_TriggerResponse.Size(m)
}
func (m *TriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerResponse proto.InternalMessageInfo

func (m *TriggerResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type isTriggerResponse_Context interface {
	isTriggerResponse_Context()
}

type TriggerResponse_Http struct {
	Http *HttpResponseContext `protobuf:"bytes,10,opt,name=http,proto3,oneof"`
}

type TriggerResponse_Topic struct {
	Topic *TopicResponseContext `protobuf:"bytes,11,opt,name=topic,proto3,oneof"`
}

func (*TriggerResponse_Http) isTriggerResponse_Context() {}

func (*TriggerResponse_Topic) isTriggerResponse_Context() {}

func (m *TriggerResponse) GetContext() isTriggerResponse_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TriggerResponse) GetHttp() *HttpResponseContext {
	if x, ok := m.GetContext().(*TriggerResponse_Http); ok {
		return x.Http
	}
	return nil
}

func (m *TriggerResponse) GetTopic() *TopicResponseContext {
	if x, ok := m.GetContext().(*TriggerResponse_Topic); ok {
		return x.Topic
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TriggerResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TriggerResponse_Http)(nil),
		(*TriggerResponse_Topic)(nil),
	}
}

// Specific HttpResponse message
// Note this does not have to be handled by the
// User at all but they will have the option of control
// If they choose...
type HttpResponseContext struct {
	// Old HTTP response headers (deprecated)
	// TODO: Remove in 1.0
	HeadersOld map[string]string `protobuf:"bytes,1,rep,name=headers_old,json=headersOld,proto3" json:"headers_old,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	// The HTTP status of the request
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// HTTP response headers
	Headers              map[string]*HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpResponseContext) Reset()         { *m = HttpResponseContext{} }
func (m *HttpResponseContext) String() string { return proto.CompactTextString(m) }
func (*HttpResponseContext) ProtoMessage()    {}
func (*HttpResponseContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{9}
}

func (m *HttpResponseContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpResponseContext.Unmarshal(m, b)
}
func (m *HttpResponseContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpResponseContext.Marshal(b, m, deterministic)
}
func (m *HttpResponseContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpResponseContext.Merge(m, src)
}
func (m *HttpResponseContext) XXX_Size() int {
	return xxx_messageInfo_HttpResponseContext.Size(m)
}
func (m *HttpResponseContext) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpResponseContext.DiscardUnknown(m)
}

var xxx_messageInfo_HttpResponseContext proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *HttpResponseContext) GetHeadersOld() map[string]string {
	if m != nil {
		return m.HeadersOld
	}
	return nil
}

func (m *HttpResponseContext) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpResponseContext) GetHeaders() map[string]*HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Specific event response message
// We do not accept responses for events
// only whether or not they were successfully processed
type TopicResponseContext struct {
	// Success status of the handled event
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicResponseContext) Reset()         { *m = TopicResponseContext{} }
func (m *TopicResponseContext) String() string { return proto.CompactTextString(m) }
func (*TopicResponseContext) ProtoMessage()    {}
func (*TopicResponseContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4fcab4ed1b17145, []int{10}
}

func (m *TopicResponseContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicResponseContext.Unmarshal(m, b)
}
func (m *TopicResponseContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicResponseContext.Marshal(b, m, deterministic)
}
func (m *TopicResponseContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicResponseContext.Merge(m, src)
}
func (m *TopicResponseContext) XXX_Size() int {
	return xxx_messageInfo_TopicResponseContext.Size(m)
}
func (m *TopicResponseContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicResponseContext.DiscardUnknown(m)
}

var xxx_messageInfo_TopicResponseContext proto.InternalMessageInfo

func (m *TopicResponseContext) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ClientMessage)(nil), "nitric.faas.v1.ClientMessage")
	proto.RegisterType((*ServerMessage)(nil), "nitric.faas.v1.ServerMessage")
	proto.RegisterType((*InitRequest)(nil), "nitric.faas.v1.InitRequest")
	proto.RegisterType((*InitResponse)(nil), "nitric.faas.v1.InitResponse")
	proto.RegisterType((*TriggerRequest)(nil), "nitric.faas.v1.TriggerRequest")
	proto.RegisterType((*HeaderValue)(nil), "nitric.faas.v1.HeaderValue")
	proto.RegisterType((*HttpTriggerContext)(nil), "nitric.faas.v1.HttpTriggerContext")
	proto.RegisterMapType((map[string]*HeaderValue)(nil), "nitric.faas.v1.HttpTriggerContext.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "nitric.faas.v1.HttpTriggerContext.HeadersOldEntry")
	proto.RegisterMapType((map[string]string)(nil), "nitric.faas.v1.HttpTriggerContext.QueryParamsEntry")
	proto.RegisterType((*TopicTriggerContext)(nil), "nitric.faas.v1.TopicTriggerContext")
	proto.RegisterType((*TriggerResponse)(nil), "nitric.faas.v1.TriggerResponse")
	proto.RegisterType((*HttpResponseContext)(nil), "nitric.faas.v1.HttpResponseContext")
	proto.RegisterMapType((map[string]*HeaderValue)(nil), "nitric.faas.v1.HttpResponseContext.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "nitric.faas.v1.HttpResponseContext.HeadersOldEntry")
	proto.RegisterType((*TopicResponseContext)(nil), "nitric.faas.v1.TopicResponseContext")
}

func init() { proto.RegisterFile("faas/v1/faas.proto", fileDescriptor_b4fcab4ed1b17145) }

var fileDescriptor_b4fcab4ed1b17145 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x4e, 0xdb, 0x4c,
	0x10, 0xc6, 0x39, 0x00, 0x19, 0xe7, 0x80, 0x16, 0xc4, 0x6f, 0xc1, 0xdf, 0x36, 0x32, 0xbd, 0x88,
	0x54, 0x29, 0x21, 0xe1, 0x86, 0x96, 0xb6, 0xaa, 0x82, 0x8a, 0x42, 0xd5, 0x03, 0x18, 0x04, 0x55,
	0x85, 0x94, 0x6e, 0x9d, 0x2d, 0x59, 0x35, 0xb1, 0x8d, 0x77, 0x13, 0x91, 0xf7, 0xe8, 0x4b, 0xb4,
	0x8f, 0xd0, 0xdb, 0xde, 0xf5, 0x55, 0xfa, 0x12, 0xd5, 0xee, 0xda, 0xc4, 0x76, 0x8c, 0x00, 0xa9,
	0xbd, 0xca, 0xce, 0x78, 0xbe, 0x99, 0x6f, 0x66, 0xbf, 0xc9, 0x02, 0xfa, 0x8c, 0x31, 0x6b, 0x8c,
	0x9b, 0x0d, 0xf1, 0x5b, 0xf7, 0x7c, 0x97, 0xbb, 0xa8, 0xec, 0x50, 0xee, 0x53, 0xbb, 0x2e, 0x5d,
	0xe3, 0xa6, 0xf9, 0x43, 0x83, 0xd2, 0xee, 0x80, 0x12, 0x87, 0xbf, 0x21, 0x8c, 0xe1, 0x73, 0x82,
	0xca, 0x90, 0xa1, 0x3d, 0x43, 0xab, 0x6a, 0xb5, 0x82, 0x95, 0xa1, 0x3d, 0xf4, 0x02, 0x8a, 0xd4,
	0xa1, 0xbc, 0xeb, 0x93, 0x8b, 0x11, 0x61, 0xdc, 0xc8, 0x54, 0xb5, 0x9a, 0xde, 0x5a, 0xaf, 0xc7,
	0x13, 0xd5, 0xf7, 0x1d, 0xca, 0x2d, 0x15, 0xd2, 0x99, 0xb3, 0x74, 0x3a, 0x35, 0xd1, 0x6b, 0x58,
	0xe2, 0x3e, 0x3d, 0x3f, 0x27, 0x7e, 0xd7, 0x27, 0xcc, 0x73, 0x1d, 0x46, 0x8c, 0xac, 0xcc, 0xf2,
	0x20, 0x99, 0xe5, 0x58, 0xc5, 0x59, 0x41, 0x58, 0x67, 0xce, 0xaa, 0xf0, 0xb8, 0xab, 0x5d, 0x80,
	0x05, 0xdb, 0x75, 0x38, 0x71, 0xb8, 0x24, 0x7f, 0x44, 0xfc, 0x31, 0xf1, 0xaf, 0x23, 0xbf, 0x0b,
	0xa5, 0x80, 0x7c, 0x50, 0x57, 0xb1, 0xff, 0x3f, 0x9d, 0xfd, 0x55, 0xd1, 0x22, 0x8d, 0xd8, 0x68,
	0x1f, 0x2a, 0x53, 0xfe, 0x6a, 0x08, 0x8a, 0xfe, 0xfd, 0x6b, 0xe9, 0x87, 0x73, 0x28, 0xf3, 0x98,
	0x27, 0x4a, 0xbe, 0x04, 0x7a, 0x64, 0x66, 0x66, 0x19, 0x8a, 0x51, 0x12, 0xe6, 0x4f, 0x0d, 0xca,
	0xf1, 0x74, 0x08, 0x41, 0xae, 0x87, 0x39, 0x96, 0xed, 0x15, 0x2d, 0x79, 0x46, 0xeb, 0x50, 0x18,
	0xd2, 0x21, 0xe9, 0xf2, 0x89, 0xa7, 0x9a, 0x2b, 0x58, 0x8b, 0xc2, 0x71, 0x3c, 0xf1, 0x08, 0xda,
	0x86, 0x5c, 0x9f, 0x73, 0x2f, 0x60, 0x6b, 0x26, 0xd9, 0x76, 0x38, 0xf7, 0x82, 0x12, 0xbb, 0x82,
	0xd4, 0xa5, 0x60, 0x2c, 0x11, 0x68, 0x07, 0xf2, 0xdc, 0xf5, 0xa8, 0x6d, 0xe4, 0x24, 0x74, 0x63,
	0xa6, 0x51, 0xf1, 0x71, 0x06, 0xab, 0x30, 0x57, 0x4d, 0x5e, 0x72, 0x73, 0x03, 0xf4, 0x0e, 0xc1,
	0x3d, 0xe2, 0x9f, 0xe0, 0xc1, 0x88, 0xa0, 0x15, 0xc8, 0x8f, 0xc5, 0xc1, 0xd0, 0xaa, 0xd9, 0x5a,
	0xc1, 0x52, 0x86, 0xf9, 0x35, 0x07, 0x68, 0x96, 0x0b, 0x5a, 0x85, 0xf9, 0x21, 0xe1, 0x7d, 0x37,
	0xbc, 0xcf, 0xc0, 0x12, 0x63, 0xf0, 0x30, 0xef, 0x07, 0xdd, 0xca, 0x33, 0x3a, 0x05, 0xbd, 0x2f,
	0xeb, 0xb0, 0xae, 0x3b, 0xe8, 0x19, 0xd9, 0x6a, 0xb6, 0xa6, 0xb7, 0x5a, 0x37, 0x37, 0x5c, 0x57,
	0xec, 0xd8, 0xbb, 0x41, 0xef, 0xa5, 0xc3, 0xfd, 0x49, 0x3b, 0x63, 0x68, 0x16, 0xf4, 0xaf, 0x9c,
	0xe8, 0x04, 0x8a, 0x17, 0x23, 0xe2, 0x4f, 0xba, 0x1e, 0xf6, 0xf1, 0x90, 0x19, 0x39, 0x99, 0x79,
	0xeb, 0x16, 0x99, 0x0f, 0x05, 0xec, 0x40, 0xa2, 0x64, 0x6a, 0x4b, 0xbf, 0x98, 0x7a, 0xd0, 0x3e,
	0x2c, 0x04, 0x55, 0x8c, 0xbc, 0x4c, 0xd9, 0xb8, 0x3d, 0x59, 0x95, 0x2e, 0xc4, 0xaf, 0x3d, 0x83,
	0x4a, 0xa2, 0x0b, 0xb4, 0x04, 0xd9, 0x2f, 0x64, 0x12, 0xcc, 0x4d, 0x1c, 0xa7, 0x93, 0x57, 0x53,
	0x53, 0xc6, 0x93, 0xcc, 0xb6, 0xb6, 0xf6, 0x1c, 0x96, 0x92, 0x54, 0xef, 0x84, 0x3f, 0x85, 0x62,
	0x94, 0x57, 0x0a, 0xb6, 0x19, 0xc5, 0xa6, 0xfc, 0x75, 0x44, 0x14, 0x12, 0x49, 0x6c, 0x3e, 0x82,
	0xe5, 0x14, 0x99, 0x09, 0x26, 0x4a, 0x9a, 0xaa, 0x82, 0x32, 0xcc, 0x6f, 0x1a, 0x54, 0x12, 0x7f,
	0x1e, 0xa9, 0xfb, 0xf2, 0x38, 0x58, 0x09, 0x48, 0xd7, 0xb5, 0x18, 0x7a, 0x88, 0x4f, 0xee, 0xc4,
	0xd3, 0xb0, 0xb0, 0x2e, 0xb1, 0x0f, 0x53, 0x77, 0x62, 0x16, 0x3c, 0xbb, 0x14, 0xbf, 0x33, 0xb0,
	0x9c, 0x52, 0x08, 0xbd, 0x8f, 0x8b, 0x58, 0xbb, 0x5e, 0x6a, 0x09, 0xe4, 0x8d, 0x2a, 0x5e, 0x85,
	0x79, 0xc6, 0x31, 0x1f, 0x31, 0x79, 0x05, 0x79, 0x2b, 0xb0, 0xd0, 0xab, 0xa9, 0x0a, 0xd5, 0xca,
	0x6c, 0xde, 0xa1, 0xda, 0xdf, 0x95, 0xe1, 0x3f, 0x93, 0xd1, 0x26, 0xac, 0xa4, 0xdd, 0x0c, 0x32,
	0x60, 0x81, 0x8d, 0x6c, 0x9b, 0x30, 0x26, 0x8b, 0x2c, 0x5a, 0xa1, 0xd9, 0xfa, 0x08, 0xfa, 0x1e,
	0xc6, 0x4c, 0xbc, 0x2c, 0xd4, 0x26, 0xe8, 0x10, 0x4a, 0x81, 0xb2, 0x8e, 0xb8, 0x4f, 0xf0, 0x10,
	0xdd, 0x4b, 0x56, 0x8e, 0x3d, 0xa0, 0x6b, 0x33, 0x9f, 0x63, 0x4f, 0x54, 0x4d, 0xdb, 0xd4, 0xda,
	0x36, 0xfc, 0x47, 0xdd, 0x30, 0x4c, 0x3e, 0xcc, 0x61, 0x70, 0x1b, 0xde, 0x4a, 0xaf, 0x20, 0x70,
	0xa0, 0x7d, 0x28, 0xaa, 0x98, 0xc6, 0xb8, 0xb9, 0x33, 0x6e, 0x7e, 0xcf, 0xac, 0xa8, 0x8f, 0xf5,
	0x03, 0x09, 0xd9, 0x53, 0x90, 0x5f, 0xa1, 0xfb, 0x4c, 0xba, 0xcf, 0x84, 0xfb, 0xec, 0xa4, 0xf9,
	0x69, 0x5e, 0x26, 0xde, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x89, 0x94, 0xca, 0x07, 0x08,
	0x00, 0x00,
}
