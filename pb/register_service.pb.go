// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: register_service.proto

package pb

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

type RequestRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RequestRegister) Reset() {
	*x = RequestRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister) ProtoMessage() {}

func (x *RequestRegister) ProtoReflect() protoreflect.Message {
	mi := &file_register_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister.ProtoReflect.Descriptor instead.
func (*RequestRegister) Descriptor() ([]byte, []int) {
	return file_register_service_proto_rawDescGZIP(), []int{0}
}

func (x *RequestRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestRegister) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestRegister) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *RequestRegister) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestRegister) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_register_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_register_service_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Role    string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ResponseRegister) Reset() {
	*x = ResponseRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister) ProtoMessage() {}

func (x *ResponseRegister) ProtoReflect() protoreflect.Message {
	mi := &file_register_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister.ProtoReflect.Descriptor instead.
func (*ResponseRegister) Descriptor() ([]byte, []int) {
	return file_register_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseRegister) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponseRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResponseRegister) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ResponseRegister) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ResponseRegister) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RequestToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RequestToken) Reset() {
	*x = RequestToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToken) ProtoMessage() {}

func (x *RequestToken) ProtoReflect() protoreflect.Message {
	mi := &file_register_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToken.ProtoReflect.Descriptor instead.
func (*RequestToken) Descriptor() ([]byte, []int) {
	return file_register_service_proto_rawDescGZIP(), []int{3}
}

func (x *RequestToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ResponseToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *ResponseToken) Reset() {
	*x = ResponseToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseToken) ProtoMessage() {}

func (x *ResponseToken) ProtoReflect() protoreflect.Message {
	mi := &file_register_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseToken.ProtoReflect.Descriptor instead.
func (*ResponseToken) Descriptor() ([]byte, []int) {
	return file_register_service_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ResponseToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_register_service_proto protoreflect.FileDescriptor

var file_register_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x87, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7c, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x57, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa4, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65,
	0x67, 0x65, 0x6e, 0x6f, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x20, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_register_service_proto_rawDescOnce sync.Once
	file_register_service_proto_rawDescData = file_register_service_proto_rawDesc
)

func file_register_service_proto_rawDescGZIP() []byte {
	file_register_service_proto_rawDescOnce.Do(func() {
		file_register_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_service_proto_rawDescData)
	})
	return file_register_service_proto_rawDescData
}

var file_register_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_register_service_proto_goTypes = []interface{}{
	(*RequestRegister)(nil),  // 0: pb.RequestRegister
	(*Response)(nil),         // 1: pb.Response
	(*ResponseRegister)(nil), // 2: pb.ResponseRegister
	(*RequestToken)(nil),     // 3: pb.RequestToken
	(*ResponseToken)(nil),    // 4: pb.ResponseToken
}
var file_register_service_proto_depIdxs = []int32{
	0, // 0: pb.Register.SignUp:input_type -> pb.RequestRegister
	0, // 1: pb.Register.SignIn:input_type -> pb.RequestRegister
	3, // 2: pb.Register.RefreshToken:input_type -> pb.RequestToken
	1, // 3: pb.Register.SignUp:output_type -> pb.Response
	4, // 4: pb.Register.SignIn:output_type -> pb.ResponseToken
	4, // 5: pb.Register.RefreshToken:output_type -> pb.ResponseToken
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_register_service_proto_init() }
func file_register_service_proto_init() {
	if File_register_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRegister); i {
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
		file_register_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_register_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRegister); i {
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
		file_register_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestToken); i {
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
		file_register_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseToken); i {
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
			RawDescriptor: file_register_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_register_service_proto_goTypes,
		DependencyIndexes: file_register_service_proto_depIdxs,
		MessageInfos:      file_register_service_proto_msgTypes,
	}.Build()
	File_register_service_proto = out.File
	file_register_service_proto_rawDesc = nil
	file_register_service_proto_goTypes = nil
	file_register_service_proto_depIdxs = nil
}
