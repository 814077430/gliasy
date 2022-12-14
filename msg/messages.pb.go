// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: messages.proto

package msg

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

// ---------------C2S MSG START---------------//
type C2S_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *string `protobuf:"bytes,1,opt,name=Account,proto3,oneof" json:"Account,omitempty"`
}

func (x *C2S_Login) Reset() {
	*x = C2S_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Login) ProtoMessage() {}

func (x *C2S_Login) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_Login.ProtoReflect.Descriptor instead.
func (*C2S_Login) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_Login) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

type C2S_Logout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *string `protobuf:"bytes,1,opt,name=Account,proto3,oneof" json:"Account,omitempty"`
}

func (x *C2S_Logout) Reset() {
	*x = C2S_Logout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Logout) ProtoMessage() {}

func (x *C2S_Logout) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_Logout.ProtoReflect.Descriptor instead.
func (*C2S_Logout) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *C2S_Logout) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

// ---------------S2C MSG START---------------//
type S2C_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code *int32    `protobuf:"varint,1,opt,name=Code,proto3,oneof" json:"Code,omitempty"`
	Data *UserData `protobuf:"bytes,2,opt,name=Data,proto3,oneof" json:"Data,omitempty"`
}

func (x *S2C_Login) Reset() {
	*x = S2C_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Login) ProtoMessage() {}

func (x *S2C_Login) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Login.ProtoReflect.Descriptor instead.
func (*S2C_Login) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_Login) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *S2C_Login) GetData() *UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

type S2C_Logout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code *int32 `protobuf:"varint,1,opt,name=Code,proto3,oneof" json:"Code,omitempty"`
}

func (x *S2C_Logout) Reset() {
	*x = S2C_Logout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Logout) ProtoMessage() {}

func (x *S2C_Logout) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Logout.ProtoReflect.Descriptor instead.
func (*S2C_Logout) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_Logout) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

// ---------------RPC MSG START---------------//
type G2D_PlayerIncrId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncrId *int32 `protobuf:"varint,1,opt,name=IncrId,proto3,oneof" json:"IncrId,omitempty"`
}

func (x *G2D_PlayerIncrId) Reset() {
	*x = G2D_PlayerIncrId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G2D_PlayerIncrId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G2D_PlayerIncrId) ProtoMessage() {}

func (x *G2D_PlayerIncrId) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G2D_PlayerIncrId.ProtoReflect.Descriptor instead.
func (*G2D_PlayerIncrId) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *G2D_PlayerIncrId) GetIncrId() int32 {
	if x != nil && x.IncrId != nil {
		return *x.IncrId
	}
	return 0
}

type G2D_CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *UserData `protobuf:"bytes,1,opt,name=Data,proto3,oneof" json:"Data,omitempty"`
}

func (x *G2D_CreatePlayer) Reset() {
	*x = G2D_CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G2D_CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G2D_CreatePlayer) ProtoMessage() {}

func (x *G2D_CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G2D_CreatePlayer.ProtoReflect.Descriptor instead.
func (*G2D_CreatePlayer) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{5}
}

func (x *G2D_CreatePlayer) GetData() *UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

// ---------------COMMON START------------//
type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  *string   `protobuf:"bytes,1,opt,name=Account,proto3,oneof" json:"Account,omitempty"`
	Uid      *int32    `protobuf:"varint,2,opt,name=Uid,proto3,oneof" json:"Uid,omitempty"`
	RoleData *RoleData `protobuf:"bytes,3,opt,name=RoleData,proto3,oneof" json:"RoleData,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UserData) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *UserData) GetUid() int32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *UserData) GetRoleData() *RoleData {
	if x != nil {
		return x.RoleData
	}
	return nil
}

type RoleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
}

func (x *RoleData) Reset() {
	*x = RoleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleData) ProtoMessage() {}

func (x *RoleData) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleData.ProtoReflect.Descriptor instead.
func (*RoleData) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{7}
}

func (x *RoleData) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x36, 0x0a, 0x09, 0x43, 0x32, 0x53, 0x5f, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x0a,
	0x43, 0x32, 0x53, 0x5f, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x09, 0x53, 0x32, 0x43, 0x5f, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x0a, 0x53, 0x32, 0x43, 0x5f, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x32, 0x44, 0x5f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x63, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x49, 0x6e, 0x63, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x49, 0x6e, 0x63, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x49, 0x6e, 0x63, 0x72, 0x49, 0x64,
	0x22, 0x42, 0x0a, 0x10, 0x47, 0x32, 0x44, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x02, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x55, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x52,
	0x6f, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_messages_proto_goTypes = []interface{}{
	(*C2S_Login)(nil),        // 0: pb.C2S_Login
	(*C2S_Logout)(nil),       // 1: pb.C2S_Logout
	(*S2C_Login)(nil),        // 2: pb.S2C_Login
	(*S2C_Logout)(nil),       // 3: pb.S2C_Logout
	(*G2D_PlayerIncrId)(nil), // 4: pb.G2D_PlayerIncrId
	(*G2D_CreatePlayer)(nil), // 5: pb.G2D_CreatePlayer
	(*UserData)(nil),         // 6: pb.UserData
	(*RoleData)(nil),         // 7: pb.RoleData
}
var file_messages_proto_depIdxs = []int32{
	6, // 0: pb.S2C_Login.Data:type_name -> pb.UserData
	6, // 1: pb.G2D_CreatePlayer.Data:type_name -> pb.UserData
	7, // 2: pb.UserData.RoleData:type_name -> pb.RoleData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Login); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Logout); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Login); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Logout); i {
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G2D_PlayerIncrId); i {
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
		file_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G2D_CreatePlayer); i {
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
		file_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleData); i {
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
	file_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_messages_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
