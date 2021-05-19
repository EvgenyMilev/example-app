// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/EvgenyMilev/example-app/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	query "github.com/infobloxopen/atlas-app-toolkit/query"
	_ "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
	_ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Genre struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genre) Reset()         { *m = Genre{} }
func (m *Genre) String() string { return proto.CompactTextString(m) }
func (*Genre) ProtoMessage()    {}
func (*Genre) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{0}
}

func (m *Genre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre.Unmarshal(m, b)
}
func (m *Genre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre.Marshal(b, m, deterministic)
}
func (m *Genre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre.Merge(m, src)
}
func (m *Genre) XXX_Size() int {
	return xxx_messageInfo_Genre.Size(m)
}
func (m *Genre) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre.DiscardUnknown(m)
}

var xxx_messageInfo_Genre proto.InternalMessageInfo

func (m *Genre) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Genre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Book struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Genre                int32    `protobuf:"varint,4,opt,name=genre,proto3" json:"genre,omitempty"`
	Amount               int32    `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Book) GetGenre() int32 {
	if m != nil {
		return m.Genre
	}
	return 0
}

func (m *Book) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{2}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ListBookRequest struct {
	Filter               *query.Filtering      `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              *query.Sorting        `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Fields               *query.FieldSelection `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	Paging               *query.Pagination     `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListBookRequest) Reset()         { *m = ListBookRequest{} }
func (m *ListBookRequest) String() string { return proto.CompactTextString(m) }
func (*ListBookRequest) ProtoMessage()    {}
func (*ListBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{3}
}

func (m *ListBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBookRequest.Unmarshal(m, b)
}
func (m *ListBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBookRequest.Marshal(b, m, deterministic)
}
func (m *ListBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBookRequest.Merge(m, src)
}
func (m *ListBookRequest) XXX_Size() int {
	return xxx_messageInfo_ListBookRequest.Size(m)
}
func (m *ListBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBookRequest proto.InternalMessageInfo

func (m *ListBookRequest) GetFilter() *query.Filtering {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListBookRequest) GetOrderBy() *query.Sorting {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListBookRequest) GetFields() *query.FieldSelection {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListBookRequest) GetPaging() *query.Pagination {
	if m != nil {
		return m.Paging
	}
	return nil
}

type ListBookResponse struct {
	Results              []*Book         `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Page                 *query.PageInfo `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListBookResponse) Reset()         { *m = ListBookResponse{} }
func (m *ListBookResponse) String() string { return proto.CompactTextString(m) }
func (*ListBookResponse) ProtoMessage()    {}
func (*ListBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{4}
}

func (m *ListBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBookResponse.Unmarshal(m, b)
}
func (m *ListBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBookResponse.Marshal(b, m, deterministic)
}
func (m *ListBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBookResponse.Merge(m, src)
}
func (m *ListBookResponse) XXX_Size() int {
	return xxx_messageInfo_ListBookResponse.Size(m)
}
func (m *ListBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBookResponse proto.InternalMessageInfo

func (m *ListBookResponse) GetResults() []*Book {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListBookResponse) GetPage() *query.PageInfo {
	if m != nil {
		return m.Page
	}
	return nil
}

type ReadBookRequest struct {
	Id                   int32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *query.FieldSelection `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ReadBookRequest) Reset()         { *m = ReadBookRequest{} }
func (m *ReadBookRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBookRequest) ProtoMessage()    {}
func (*ReadBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{5}
}

func (m *ReadBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBookRequest.Unmarshal(m, b)
}
func (m *ReadBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBookRequest.Marshal(b, m, deterministic)
}
func (m *ReadBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBookRequest.Merge(m, src)
}
func (m *ReadBookRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBookRequest.Size(m)
}
func (m *ReadBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBookRequest proto.InternalMessageInfo

func (m *ReadBookRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadBookRequest) GetFields() *query.FieldSelection {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ReadBookResponse struct {
	Result               *Book    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBookResponse) Reset()         { *m = ReadBookResponse{} }
func (m *ReadBookResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBookResponse) ProtoMessage()    {}
func (*ReadBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{6}
}

func (m *ReadBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBookResponse.Unmarshal(m, b)
}
func (m *ReadBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBookResponse.Marshal(b, m, deterministic)
}
func (m *ReadBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBookResponse.Merge(m, src)
}
func (m *ReadBookResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBookResponse.Size(m)
}
func (m *ReadBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBookResponse proto.InternalMessageInfo

func (m *ReadBookResponse) GetResult() *Book {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateBookRequest struct {
	Payload              *Book    `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{7}
}

func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (m *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(m, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

func (m *CreateBookRequest) GetPayload() *Book {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateBookResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookResponse) Reset()         { *m = CreateBookResponse{} }
func (m *CreateBookResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBookResponse) ProtoMessage()    {}
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{8}
}

func (m *CreateBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookResponse.Unmarshal(m, b)
}
func (m *CreateBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookResponse.Marshal(b, m, deterministic)
}
func (m *CreateBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookResponse.Merge(m, src)
}
func (m *CreateBookResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBookResponse.Size(m)
}
func (m *CreateBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookResponse proto.InternalMessageInfo

func (m *CreateBookResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateBookRequest struct {
	Payload              *Book                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Fields               *field_mask.FieldMask `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateBookRequest) Reset()         { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()    {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{9}
}

func (m *UpdateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBookRequest.Unmarshal(m, b)
}
func (m *UpdateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBookRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBookRequest.Merge(m, src)
}
func (m *UpdateBookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBookRequest.Size(m)
}
func (m *UpdateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBookRequest proto.InternalMessageInfo

func (m *UpdateBookRequest) GetPayload() *Book {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UpdateBookRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

type UpdateBookResponse struct {
	Result               *Book    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBookResponse) Reset()         { *m = UpdateBookResponse{} }
func (m *UpdateBookResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBookResponse) ProtoMessage()    {}
func (*UpdateBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{10}
}

func (m *UpdateBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBookResponse.Unmarshal(m, b)
}
func (m *UpdateBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBookResponse.Marshal(b, m, deterministic)
}
func (m *UpdateBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBookResponse.Merge(m, src)
}
func (m *UpdateBookResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBookResponse.Size(m)
}
func (m *UpdateBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBookResponse proto.InternalMessageInfo

func (m *UpdateBookResponse) GetResult() *Book {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteBookRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookRequest) Reset()         { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()    {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{11}
}

func (m *DeleteBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookRequest.Unmarshal(m, b)
}
func (m *DeleteBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookRequest.Merge(m, src)
}
func (m *DeleteBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookRequest.Size(m)
}
func (m *DeleteBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookRequest proto.InternalMessageInfo

func (m *DeleteBookRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteBookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookResponse) Reset()         { *m = DeleteBookResponse{} }
func (m *DeleteBookResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBookResponse) ProtoMessage()    {}
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cb65a2a5e6d522, []int{12}
}

func (m *DeleteBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookResponse.Unmarshal(m, b)
}
func (m *DeleteBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookResponse.Merge(m, src)
}
func (m *DeleteBookResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBookResponse.Size(m)
}
func (m *DeleteBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Genre)(nil), "service.Genre")
	proto.RegisterType((*Book)(nil), "service.Book")
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
	proto.RegisterType((*ListBookRequest)(nil), "service.ListBookRequest")
	proto.RegisterType((*ListBookResponse)(nil), "service.ListBookResponse")
	proto.RegisterType((*ReadBookRequest)(nil), "service.ReadBookRequest")
	proto.RegisterType((*ReadBookResponse)(nil), "service.ReadBookResponse")
	proto.RegisterType((*CreateBookRequest)(nil), "service.CreateBookRequest")
	proto.RegisterType((*CreateBookResponse)(nil), "service.CreateBookResponse")
	proto.RegisterType((*UpdateBookRequest)(nil), "service.UpdateBookRequest")
	proto.RegisterType((*UpdateBookResponse)(nil), "service.UpdateBookResponse")
	proto.RegisterType((*DeleteBookRequest)(nil), "service.DeleteBookRequest")
	proto.RegisterType((*DeleteBookResponse)(nil), "service.DeleteBookResponse")
}

func init() {
	proto.RegisterFile("github.com/EvgenyMilev/example-app/pkg/pb/service.proto", fileDescriptor_02cb65a2a5e6d522)
}

var fileDescriptor_02cb65a2a5e6d522 = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5d, 0x8f, 0xdb, 0x44,
	0x17, 0x7e, 0x9d, 0x66, 0x9d, 0xec, 0xc9, 0xdb, 0x6e, 0x76, 0x0a, 0x6d, 0xe2, 0x2e, 0x55, 0x64,
	0x40, 0x5d, 0x95, 0xae, 0x5d, 0x05, 0x24, 0xc4, 0x16, 0x21, 0x35, 0x74, 0xa9, 0xf8, 0x28, 0x5a,
	0xa5, 0x7c, 0x48, 0x7b, 0xb3, 0x4c, 0xe2, 0x13, 0x33, 0x8a, 0xe3, 0x99, 0xda, 0x93, 0xb0, 0x11,
	0xea, 0x0d, 0x37, 0xfc, 0x00, 0x7e, 0x4d, 0x72, 0x01, 0x3f, 0x81, 0x0b, 0xee, 0x10, 0x97, 0xfc,
	0x10, 0xe4, 0x99, 0x71, 0xe2, 0x75, 0x16, 0x14, 0xb8, 0xb3, 0xcf, 0xc7, 0xf3, 0x9c, 0xf3, 0x9c,
	0x93, 0x13, 0xc3, 0xbb, 0x21, 0x93, 0xdf, 0x4e, 0x07, 0xde, 0x90, 0x4f, 0xfc, 0x93, 0x59, 0x88,
	0xf1, 0xfc, 0x19, 0x8b, 0x70, 0xe6, 0xe3, 0x05, 0x9d, 0x88, 0x08, 0x8f, 0xa8, 0x10, 0xbe, 0x18,
	0x87, 0xbe, 0x18, 0xf8, 0x29, 0x26, 0x33, 0x36, 0x44, 0x4f, 0x24, 0x5c, 0x72, 0x52, 0x33, 0xaf,
	0xce, 0x9d, 0x90, 0xf3, 0x30, 0x42, 0x5f, 0x99, 0x07, 0xd3, 0x91, 0x8f, 0x13, 0x21, 0xe7, 0x3a,
	0xca, 0x39, 0x30, 0x4e, 0x2a, 0x98, 0x4f, 0xe3, 0x98, 0x4b, 0x2a, 0x19, 0x8f, 0x53, 0xe3, 0x7d,
	0x5c, 0x20, 0xc7, 0x78, 0xc6, 0xe7, 0x22, 0xe1, 0x17, 0x73, 0x8d, 0x34, 0x3c, 0x0a, 0x31, 0x3e,
	0x9a, 0xd1, 0x88, 0x05, 0x54, 0xa2, 0xbf, 0xf1, 0x60, 0x20, 0x1e, 0x14, 0x82, 0xd3, 0xef, 0x68,
	0x18, 0x62, 0xe2, 0x73, 0xa1, 0x48, 0xae, 0x20, 0x3c, 0x2e, 0x10, 0xb2, 0x78, 0xc4, 0x07, 0x11,
	0xbf, 0xe0, 0x02, 0xe3, 0x22, 0x65, 0xc8, 0x93, 0xc9, 0x0a, 0x22, 0x7b, 0x31, 0xb9, 0x9f, 0xfc,
	0x5d, 0x2e, 0x95, 0x11, 0x4d, 0x33, 0xa1, 0x8e, 0x24, 0xe7, 0xd1, 0x98, 0x49, 0xff, 0xc5, 0x14,
	0x93, 0xb9, 0x3f, 0xe4, 0x51, 0x84, 0xc3, 0x0c, 0xe7, 0x9c, 0x0b, 0x4c, 0xa8, 0xe4, 0x49, 0x5e,
	0x47, 0xa7, 0xac, 0xd9, 0x88, 0x61, 0x14, 0x9c, 0x4f, 0x68, 0x3a, 0x36, 0x11, 0x5f, 0x6c, 0x51,
	0xa9, 0x26, 0x56, 0x64, 0x6b, 0xa1, 0xf2, 0xca, 0x95, 0xf9, 0xbc, 0xa4, 0xd6, 0xe7, 0x5b, 0xa3,
	0x6e, 0xe0, 0x29, 0x73, 0x09, 0xcf, 0x3d, 0x81, 0x9d, 0xa7, 0x18, 0x27, 0x48, 0x3a, 0x50, 0x61,
	0x41, 0xcb, 0xea, 0x58, 0x87, 0x3b, 0xbd, 0xe6, 0x72, 0xd1, 0xfe, 0x3f, 0x00, 0xb1, 0x53, 0x4c,
	0x18, 0x8d, 0x0e, 0xad, 0x7e, 0x85, 0x05, 0x84, 0x40, 0x35, 0xa6, 0x13, 0x6c, 0x55, 0x3a, 0xd6,
	0xe1, 0x6e, 0x5f, 0x3d, 0x1f, 0xdb, 0xcb, 0x45, 0xbb, 0x52, 0xb7, 0xdc, 0x9f, 0x2d, 0xa8, 0xf6,
	0x38, 0x1f, 0x6f, 0x01, 0xe3, 0x14, 0x61, 0x7a, 0xf6, 0xef, 0xbf, 0xb4, 0x2b, 0xcd, 0xff, 0x69,
	0x38, 0x72, 0x00, 0x3b, 0x22, 0x61, 0x43, 0x6c, 0x5d, 0xeb, 0x58, 0x87, 0x95, 0x95, 0x53, 0x1b,
	0x49, 0x17, 0x76, 0xc2, 0xac, 0xd6, 0x56, 0x55, 0xc1, 0x1f, 0x2c, 0x17, 0x6d, 0x02, 0xbb, 0xa4,
	0xc6, 0x62, 0x89, 0x21, 0x26, 0xc7, 0xba, 0x93, 0x3c, 0x47, 0x85, 0x92, 0xbb, 0x60, 0xd3, 0x09,
	0x9f, 0xc6, 0xb2, 0xb5, 0xa3, 0x92, 0x72, 0x48, 0x63, 0x5d, 0x35, 0xf0, 0x16, 0xec, 0x7d, 0x85,
	0x49, 0xca, 0x78, 0xdc, 0xc7, 0x54, 0xf0, 0x38, 0x45, 0xd2, 0x82, 0xda, 0x4c, 0x9b, 0x54, 0x3f,
	0xbb, 0xfd, 0xfc, 0xd5, 0xfd, 0xc3, 0x82, 0xbd, 0xcf, 0x58, 0x2a, 0xb3, 0x8e, 0xfb, 0xf8, 0x62,
	0x8a, 0xa9, 0x24, 0x3e, 0xd8, 0x23, 0x16, 0x49, 0x4c, 0x54, 0x70, 0xa3, 0x7b, 0xdb, 0xcb, 0xc7,
	0xe3, 0x51, 0xc1, 0xbc, 0x8f, 0x94, 0x8f, 0xc5, 0x61, 0xdf, 0x84, 0x91, 0x87, 0x50, 0xe7, 0x49,
	0x80, 0xc9, 0xf9, 0x60, 0xae, 0xb4, 0x68, 0x74, 0x5f, 0xbd, 0x9c, 0xf2, 0x9c, 0x27, 0x32, 0x4b,
	0xa8, 0xa9, 0xb0, 0xde, 0x9c, 0xbc, 0x93, 0x51, 0x60, 0x14, 0xa4, 0x4a, 0x9e, 0x46, 0xf7, 0xa0,
	0x4c, 0x81, 0x51, 0xf0, 0x1c, 0xcd, 0xc6, 0xf6, 0x4d, 0x2c, 0x79, 0x08, 0xb6, 0xa0, 0x21, 0x8b,
	0x43, 0x25, 0x5b, 0xa3, 0xdb, 0xba, 0x9c, 0x75, 0x9a, 0xf9, 0xa8, 0xce, 0xd0, 0x71, 0x6e, 0x08,
	0xcd, 0x75, 0x77, 0x46, 0x8c, 0x7b, 0x50, 0x4b, 0x30, 0x9d, 0x46, 0x32, 0x6d, 0x59, 0x9d, 0x6b,
	0x87, 0x8d, 0xee, 0x75, 0x2f, 0xbf, 0x26, 0x2a, 0x2e, 0xf7, 0x92, 0xfb, 0x50, 0x15, 0x34, 0x44,
	0xd3, 0xd2, 0xad, 0x0d, 0x32, 0xfc, 0x38, 0x1e, 0xf1, 0xbe, 0x8a, 0x71, 0xbf, 0x86, 0xbd, 0x3e,
	0xd2, 0xa0, 0x28, 0xe3, 0x8d, 0xf5, 0xfe, 0xa8, 0x6d, 0x59, 0xf7, 0x5c, 0xd9, 0xbe, 0x67, 0xf7,
	0x3d, 0x68, 0xae, 0x81, 0x4d, 0x07, 0x6f, 0x82, 0xad, 0x6b, 0x34, 0x03, 0x2a, 0x35, 0x60, 0x9c,
	0xee, 0xfb, 0xb0, 0xff, 0x61, 0x82, 0x54, 0x62, 0xb1, 0xaa, 0x7b, 0x50, 0x13, 0x74, 0x1e, 0x71,
	0x1a, 0x5c, 0x9d, 0x9c, 0x7b, 0xdd, 0x37, 0x80, 0x14, 0xb3, 0x0d, 0x75, 0xa9, 0x29, 0x57, 0xc0,
	0xfe, 0x97, 0x22, 0xf8, 0x8f, 0x1c, 0xa4, 0x5b, 0x92, 0xc4, 0xf1, 0xf4, 0x2d, 0xf2, 0xf2, 0x5b,
	0xa4, 0x55, 0x79, 0x46, 0xd3, 0xf1, 0x4a, 0x90, 0x47, 0x40, 0x8a, 0x8c, 0xff, 0x4e, 0x92, 0xd7,
	0x61, 0xff, 0x09, 0x46, 0x78, 0xb9, 0xdc, 0x72, 0x4f, 0xaf, 0x00, 0x29, 0x06, 0x69, 0x86, 0xee,
	0xaf, 0x55, 0x80, 0x13, 0xfd, 0x47, 0xf4, 0x58, 0x08, 0xf2, 0x29, 0x54, 0xb3, 0xcd, 0x22, 0xad,
	0x15, 0x51, 0xe9, 0x67, 0xe4, 0xb4, 0xaf, 0xf0, 0x68, 0x2c, 0xf7, 0xc6, 0x0f, 0xbf, 0xfd, 0xf9,
	0x53, 0xa5, 0x4e, 0x6c, 0x7f, 0xc0, 0xf9, 0x38, 0x25, 0xa7, 0x50, 0xcd, 0x86, 0x5c, 0x00, 0x2b,
	0x2d, 0x53, 0x01, 0xac, 0xbc, 0x0d, 0xee, 0x4d, 0x05, 0x76, 0x9d, 0x34, 0x34, 0x98, 0xff, 0x3d,
	0x0b, 0x5e, 0x92, 0x33, 0xb0, 0xf5, 0xf4, 0x88, 0xb3, 0xca, 0xdc, 0x58, 0x06, 0xe7, 0xce, 0x95,
	0x3e, 0x83, 0x7b, 0x5b, 0xe1, 0xee, 0xbb, 0xa6, 0xc8, 0xe3, 0xd5, 0xd4, 0x7e, 0xb4, 0xc0, 0xd6,
	0x23, 0x28, 0x80, 0x6f, 0x6c, 0x41, 0x01, 0x7c, 0x73, 0x5e, 0xee, 0x13, 0x05, 0xfe, 0x81, 0x73,
	0x33, 0x2f, 0xda, 0x80, 0x7b, 0x2c, 0x78, 0xb9, 0x62, 0x3a, 0xbb, 0xdb, 0xfd, 0x47, 0x3f, 0xf9,
	0x06, 0x6c, 0x3d, 0xa9, 0x42, 0x21, 0x1b, 0xf3, 0x2d, 0x14, 0xb2, 0x39, 0x56, 0xf7, 0xb5, 0xe5,
	0xa2, 0x6d, 0xeb, 0x8b, 0xaf, 0x75, 0xbc, 0x7f, 0x49, 0xc7, 0x53, 0x80, 0xa7, 0x28, 0xcd, 0x3d,
	0x25, 0xb7, 0x36, 0xf6, 0xf3, 0x24, 0xfb, 0xbe, 0x70, 0xd6, 0x73, 0x2b, 0x5d, 0x5e, 0xb7, 0xa9,
	0x40, 0x81, 0xd4, 0x7d, 0x73, 0x71, 0x9d, 0xfa, 0x72, 0xd1, 0xae, 0xd6, 0xad, 0xa6, 0xd5, 0xf3,
	0xce, 0x1e, 0x6c, 0xfd, 0xc1, 0xf3, 0x48, 0x0c, 0x06, 0xb6, 0x62, 0x7d, 0xfb, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x0a, 0x59, 0x04, 0x27, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleAppClient is the client API for ExampleApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleAppClient interface {
	List(ctx context.Context, in *ListBookRequest, opts ...grpc.CallOption) (*ListBookResponse, error)
	Read(ctx context.Context, in *ReadBookRequest, opts ...grpc.CallOption) (*ReadBookResponse, error)
	Create(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	Update(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
	Delete(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type exampleAppClient struct {
	cc *grpc.ClientConn
}

func NewExampleAppClient(cc *grpc.ClientConn) ExampleAppClient {
	return &exampleAppClient{cc}
}

func (c *exampleAppClient) List(ctx context.Context, in *ListBookRequest, opts ...grpc.CallOption) (*ListBookResponse, error) {
	out := new(ListBookResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAppClient) Read(ctx context.Context, in *ReadBookRequest, opts ...grpc.CallOption) (*ReadBookResponse, error) {
	out := new(ReadBookResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAppClient) Create(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAppClient) Update(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAppClient) Delete(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAppClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/service.ExampleApp/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleAppServer is the server API for ExampleApp service.
type ExampleAppServer interface {
	List(context.Context, *ListBookRequest) (*ListBookResponse, error)
	Read(context.Context, *ReadBookRequest) (*ReadBookResponse, error)
	Create(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	Update(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	Delete(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
}

func RegisterExampleAppServer(s *grpc.Server, srv ExampleAppServer) {
	s.RegisterService(&_ExampleApp_serviceDesc, srv)
}

func _ExampleApp_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).List(ctx, req.(*ListBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleApp_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).Read(ctx, req.(*ReadBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleApp_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).Create(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleApp_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).Update(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleApp_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).Delete(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleApp_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAppServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExampleApp/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAppServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ExampleApp",
	HandlerType: (*ExampleAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ExampleApp_List_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ExampleApp_Read_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ExampleApp_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ExampleApp_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ExampleApp_Delete_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ExampleApp_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/EvgenyMilev/example-app/pkg/pb/service.proto",
}
