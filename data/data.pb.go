// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: data/data.proto

package data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UsersOrder_Role int32

const (
	UsersOrder_USER     UsersOrder_Role = 0
	UsersOrder_VIP_USER UsersOrder_Role = 1
)

// Enum value maps for UsersOrder_Role.
var (
	UsersOrder_Role_name = map[int32]string{
		0: "USER",
		1: "VIP_USER",
	}
	UsersOrder_Role_value = map[string]int32{
		"USER":     0,
		"VIP_USER": 1,
	}
)

func (x UsersOrder_Role) Enum() *UsersOrder_Role {
	p := new(UsersOrder_Role)
	*p = x
	return p
}

func (x UsersOrder_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UsersOrder_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_data_data_proto_enumTypes[0].Descriptor()
}

func (UsersOrder_Role) Type() protoreflect.EnumType {
	return &file_data_data_proto_enumTypes[0]
}

func (x UsersOrder_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UsersOrder_Role.Descriptor instead.
func (UsersOrder_Role) EnumDescriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0, 0}
}

type UsersOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         *UsersOrder_UUID        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TimeStamp    int64                   `protobuf:"varint,2,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	Address      *UsersOrder_UserAddress `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	User         *UsersOrder_User        `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Order        []*UsersOrder_Product   `protobuf:"bytes,5,rep,name=Order,proto3" json:"Order,omitempty"`
	Payed        bool                    `protobuf:"varint,6,opt,name=Payed,proto3" json:"Payed,omitempty"`
	AccountTotal float32                 `protobuf:"fixed32,7,opt,name=accountTotal,proto3" json:"accountTotal,omitempty"`
}

func (x *UsersOrder) Reset() {
	*x = UsersOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersOrder) ProtoMessage() {}

func (x *UsersOrder) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersOrder.ProtoReflect.Descriptor instead.
func (*UsersOrder) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *UsersOrder) GetUuid() *UsersOrder_UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *UsersOrder) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *UsersOrder) GetAddress() *UsersOrder_UserAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UsersOrder) GetUser() *UsersOrder_User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UsersOrder) GetOrder() []*UsersOrder_Product {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *UsersOrder) GetPayed() bool {
	if x != nil {
		return x.Payed
	}
	return false
}

func (x *UsersOrder) GetAccountTotal() float32 {
	if x != nil {
		return x.AccountTotal
	}
	return 0
}

type UsersOrder_UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UsersOrder_UUID) Reset() {
	*x = UsersOrder_UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersOrder_UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersOrder_UUID) ProtoMessage() {}

func (x *UsersOrder_UUID) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersOrder_UUID.ProtoReflect.Descriptor instead.
func (*UsersOrder_UUID) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UsersOrder_UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UsersOrder_UserAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZipCode         int32  `protobuf:"varint,1,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	Country         string `protobuf:"bytes,2,opt,name=Country,proto3" json:"Country,omitempty"`
	State           string `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	City            string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	Street          string `protobuf:"bytes,5,opt,name=Street,proto3" json:"Street,omitempty"`
	NumberHouse     int32  `protobuf:"varint,6,opt,name=NumberHouse,proto3" json:"NumberHouse,omitempty"`
	NumberApartment int32  `protobuf:"varint,7,opt,name=NumberApartment,proto3" json:"NumberApartment,omitempty"`
}

func (x *UsersOrder_UserAddress) Reset() {
	*x = UsersOrder_UserAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersOrder_UserAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersOrder_UserAddress) ProtoMessage() {}

func (x *UsersOrder_UserAddress) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersOrder_UserAddress.ProtoReflect.Descriptor instead.
func (*UsersOrder_UserAddress) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *UsersOrder_UserAddress) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

func (x *UsersOrder_UserAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UsersOrder_UserAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UsersOrder_UserAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UsersOrder_UserAddress) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *UsersOrder_UserAddress) GetNumberHouse() int32 {
	if x != nil {
		return x.NumberHouse
	}
	return 0
}

func (x *UsersOrder_UserAddress) GetNumberApartment() int32 {
	if x != nil {
		return x.NumberApartment
	}
	return 0
}

type UsersOrder_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string          `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Age      int32           `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Role     UsersOrder_Role `protobuf:"varint,3,opt,name=role,proto3,enum=go_mq.data.UsersOrder_Role" json:"role,omitempty"`
}

func (x *UsersOrder_User) Reset() {
	*x = UsersOrder_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersOrder_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersOrder_User) ProtoMessage() {}

func (x *UsersOrder_User) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersOrder_User.ProtoReflect.Descriptor instead.
func (*UsersOrder_User) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0, 2}
}

func (x *UsersOrder_User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UsersOrder_User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UsersOrder_User) GetRole() UsersOrder_Role {
	if x != nil {
		return x.Role
	}
	return UsersOrder_USER
}

type UsersOrder_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price  float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Amount int32   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Cost   float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *UsersOrder_Product) Reset() {
	*x = UsersOrder_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersOrder_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersOrder_Product) ProtoMessage() {}

func (x *UsersOrder_Product) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersOrder_Product.ProtoReflect.Descriptor instead.
func (*UsersOrder_Product) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0, 3}
}

func (x *UsersOrder_Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UsersOrder_Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UsersOrder_Product) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UsersOrder_Product) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x5f, 0x6d, 0x71, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x06,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x5f,
	0x6d, 0x71, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67,
	0x6f, 0x5f, 0x6d, 0x71, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x6d, 0x71, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x6d,
	0x71, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x50, 0x61, 0x79, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x1a, 0x1c, 0x0a, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xcf, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x5a, 0x69, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x65, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x5f, 0x6d, 0x71, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x1a, 0x67, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x1e, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x56, 0x49, 0x50, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x6f,
	0x5f, 0x6d, 0x71, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData = file_data_data_proto_rawDesc
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_data_proto_rawDescData)
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_data_data_proto_goTypes = []interface{}{
	(UsersOrder_Role)(0),           // 0: go_mq.data.UsersOrder.Role
	(*UsersOrder)(nil),             // 1: go_mq.data.UsersOrder
	(*UsersOrder_UUID)(nil),        // 2: go_mq.data.UsersOrder.UUID
	(*UsersOrder_UserAddress)(nil), // 3: go_mq.data.UsersOrder.UserAddress
	(*UsersOrder_User)(nil),        // 4: go_mq.data.UsersOrder.User
	(*UsersOrder_Product)(nil),     // 5: go_mq.data.UsersOrder.Product
}
var file_data_data_proto_depIdxs = []int32{
	2, // 0: go_mq.data.UsersOrder.uuid:type_name -> go_mq.data.UsersOrder.UUID
	3, // 1: go_mq.data.UsersOrder.Address:type_name -> go_mq.data.UsersOrder.UserAddress
	4, // 2: go_mq.data.UsersOrder.user:type_name -> go_mq.data.UsersOrder.User
	5, // 3: go_mq.data.UsersOrder.Order:type_name -> go_mq.data.UsersOrder.Product
	0, // 4: go_mq.data.UsersOrder.User.role:type_name -> go_mq.data.UsersOrder.Role
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersOrder_UUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersOrder_UserAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersOrder_User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersOrder_Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		EnumInfos:         file_data_data_proto_enumTypes,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_rawDesc = nil
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}
