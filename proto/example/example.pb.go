// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_todo_list_micro

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

type Todo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Completed            bool     `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

// Request data to create a todo
type CreateRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

// Contains data of the created todo
type CreateResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request data to get a todo
type GetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains data of a specific todo by id
type GetResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

// Request data of all todos
type GetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

// Contains data of all todos
type GetAllResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

// Request data to update a todo
type UpdateRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

// Contains data of the updated todo
type UpdateResponse struct {
	Updated              int32    `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUpdated() int32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Request data to delete a todo
type DeleteRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains data of the deleted todo
type DeleteResponse struct {
	Deleted              int32    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetDeleted() int32 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type Message struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{11}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Todo)(nil), "go.micro.srv.todo.list.micro.Todo")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.todo.list.micro.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.todo.list.micro.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.todo.list.micro.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.todo.list.micro.GetResponse")
	proto.RegisterType((*GetAllRequest)(nil), "go.micro.srv.todo.list.micro.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "go.micro.srv.todo.list.micro.GetAllResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.todo.list.micro.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.todo.list.micro.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.todo.list.micro.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.todo.list.micro.DeleteResponse")
	proto.RegisterType((*Message)(nil), "go.micro.srv.todo.list.micro.Message")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x9d, 0x4e, 0xdb, 0x59, 0xe7, 0x96, 0xa9, 0x12, 0x7c, 0x28, 0xb3, 0x03, 0x96, 0x80, 0x50,
	0x57, 0xa9, 0xb0, 0x82, 0xf8, 0x2a, 0x2a, 0x15, 0xc1, 0x97, 0xaa, 0xf8, 0xe2, 0xcb, 0xd8, 0x5c,
	0x86, 0x40, 0x66, 0x53, 0x9b, 0xec, 0xe2, 0xe7, 0xfa, 0x29, 0x92, 0xa4, 0x71, 0xec, 0x82, 0x6d,
	0x91, 0x7d, 0x9a, 0xc9, 0xcd, 0xc9, 0x39, 0xe7, 0xde, 0x7b, 0x28, 0x9c, 0xb7, 0x9d, 0xd4, 0xf2,
	0x39, 0xfe, 0xdc, 0x1f, 0x5b, 0x81, 0xfe, 0xb7, 0xb4, 0x55, 0xb2, 0x3b, 0xc8, 0xf2, 0xc8, 0x9b,
	0x4e, 0x96, 0xaa, 0xbb, 0x29, 0xb5, 0x64, 0xb2, 0x14, 0x5c, 0x69, 0x57, 0xa3, 0x02, 0xa2, 0xcf,
	0x92, 0x49, 0x92, 0xc2, 0x92, 0xb3, 0x2c, 0xc8, 0x83, 0x22, 0xae, 0x97, 0x9c, 0x91, 0x87, 0x10,
	0x6b, 0xae, 0x05, 0x66, 0xcb, 0x3c, 0x28, 0xd6, 0xb5, 0x3b, 0x90, 0x1c, 0x12, 0x86, 0xaa, 0xe9,
	0x78, 0xab, 0xb9, 0xbc, 0xca, 0x42, 0x7b, 0xf7, 0x77, 0x89, 0xec, 0x60, 0xdd, 0x48, 0xa3, 0xae,
	0x91, 0x65, 0x51, 0x1e, 0x14, 0xf7, 0xea, 0x53, 0x81, 0x56, 0xb0, 0x79, 0xd3, 0xe1, 0x5e, 0x63,
	0x8d, 0x3f, 0xae, 0x51, 0x69, 0xf2, 0x12, 0x22, 0xe3, 0xc8, 0x0a, 0x27, 0x97, 0xb4, 0x1c, 0xf3,
	0x5a, 0x1a, 0xa3, 0xb5, 0xc5, 0xd3, 0x1c, 0x52, 0x4f, 0xa4, 0x5a, 0x79, 0xa5, 0xf0, 0x76, 0x03,
	0x74, 0x07, 0x50, 0xa1, 0xf6, 0x3a, 0xb7, 0x6f, 0xdf, 0x41, 0x62, 0x6f, 0xfb, 0xc7, 0xff, 0x6b,
	0xe3, 0x3e, 0x6c, 0x2a, 0xd4, 0xaf, 0x85, 0xe8, 0x75, 0xe8, 0x07, 0x48, 0x7d, 0xa1, 0xa7, 0x7e,
	0x05, 0xb1, 0x81, 0xaa, 0x2c, 0xc8, 0xc3, 0x99, 0xdc, 0xee, 0x81, 0x19, 0xd6, 0x97, 0x96, 0xdd,
	0xc1, 0xb0, 0x2e, 0x20, 0xf5, 0x44, 0xbd, 0xa9, 0x0c, 0xce, 0xae, 0x6d, 0xc5, 0xcf, 0xc4, 0x1f,
	0xe9, 0x23, 0xd8, 0xbc, 0x45, 0xb3, 0xac, 0x7f, 0x4d, 0xee, 0x02, 0x52, 0x0f, 0x38, 0x91, 0x31,
	0x74, 0x0b, 0xef, 0xc9, 0xfa, 0x23, 0x3d, 0x87, 0xb3, 0x8f, 0xa8, 0xd4, 0xfe, 0x80, 0xe4, 0x01,
	0x84, 0x47, 0x75, 0xb0, 0x80, 0x75, 0x6d, 0xfe, 0x5e, 0xfe, 0x8a, 0x20, 0x31, 0x26, 0x3f, 0x61,
	0x77, 0xc3, 0x1b, 0x24, 0x08, 0x2b, 0xb7, 0x52, 0xf2, 0x74, 0xbc, 0xb3, 0x41, 0x82, 0xb6, 0xcf,
	0xe6, 0x81, 0x9d, 0x57, 0xba, 0x20, 0xdf, 0x20, 0xac, 0x50, 0x93, 0x62, 0xfc, 0xd9, 0x29, 0x3a,
	0xdb, 0x27, 0x33, 0x90, 0x7f, 0xd8, 0x11, 0x56, 0x6e, 0xff, 0x53, 0x4d, 0x0c, 0x62, 0x33, 0xd5,
	0xc4, 0x30, 0x52, 0x4e, 0xc6, 0x6d, 0x74, 0x4a, 0x66, 0x10, 0xa0, 0x29, 0x99, 0x61, 0x48, 0x9c,
	0x8c, 0xdb, 0xf5, 0x94, 0xcc, 0x20, 0x32, 0x53, 0x32, 0xc3, 0xf8, 0xd0, 0x05, 0xf9, 0x0a, 0xf1,
	0x7b, 0x14, 0x42, 0x92, 0xc7, 0xe3, 0x0f, 0xfb, 0x2c, 0x6d, 0xe7, 0xc1, 0xe8, 0xe2, 0xfb, 0xca,
	0x7e, 0x01, 0x5f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xd0, 0x80, 0x8b, 0x20, 0x05, 0x00,
	0x00,
}