// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: expectation-api/expectation_api.proto

package expectation_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is a unique identifier of the server.
	ServerName string `protobuf:"bytes,1,opt,name=server_name,proto3" json:"server_name,omitempty"`
	// URL is a relative URI of the route.
	RouteUrl string `protobuf:"bytes,2,opt,name=route_url,proto3" json:"route_url,omitempty"`
	// ID is a unique identifier of expectation.
	Id *string `protobuf:"bytes,3,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expectation_api_expectation_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expectation_api_expectation_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_expectation_api_expectation_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *CreateRequest) GetRouteUrl() string {
	if x != nil {
		return x.RouteUrl
	}
	return ""
}

func (x *CreateRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expectation_api_expectation_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expectation_api_expectation_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_expectation_api_expectation_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is a unique identifier of the server.
	ServerName string `protobuf:"bytes,1,opt,name=server_name,proto3" json:"server_name,omitempty"`
	// URL is a relative URI of an route.
	RouteUrl string `protobuf:"bytes,2,opt,name=route_url,proto3" json:"route_url,omitempty"`
	// ID is a unique identifier of an expectation.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expectation_api_expectation_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expectation_api_expectation_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_expectation_api_expectation_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *DeleteRequest) GetRouteUrl() string {
	if x != nil {
		return x.RouteUrl
	}
	return ""
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expectation_api_expectation_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expectation_api_expectation_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_expectation_api_expectation_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_expectation_api_expectation_api_proto protoreflect.FileDescriptor

var file_expectation_api_expectation_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d,
	0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x17, 0x4a,
	0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x38, 0x30,
	0x38, 0x30, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1b, 0x92, 0x41, 0x10, 0x4a, 0x0b, 0x22, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x75, 0x72, 0x6c,
	0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90, 0x01, 0x01, 0x52, 0x09, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x4d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x38, 0x92, 0x41, 0x2b, 0x4a, 0x26, 0x22, 0x39, 0x35, 0x34, 0x32, 0x33,
	0x37, 0x30, 0x36, 0x2d, 0x31, 0x64, 0x37, 0x63, 0x2d, 0x34, 0x33, 0x31, 0x62, 0x2d, 0x62, 0x31,
	0x38, 0x32, 0x2d, 0x33, 0x62, 0x39, 0x35, 0x65, 0x62, 0x36, 0x33, 0x33, 0x31, 0x30, 0x62, 0x22,
	0x80, 0x01, 0x01, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0xd0, 0x01, 0x01, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x1f, 0x92, 0x41, 0x1c, 0x0a, 0x1a, 0xd2, 0x01, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x09, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x2a,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xfe, 0x01, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0x92, 0x41, 0x17, 0x4a, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x38, 0x30, 0x38, 0x30, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x92, 0x41, 0x10, 0x4a, 0x0b, 0x22, 0x2f, 0x73,
	0x6f, 0x6d, 0x65, 0x2d, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x90, 0x01, 0x01, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x45,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x92, 0x41, 0x2b, 0x4a,
	0x26, 0x22, 0x39, 0x35, 0x34, 0x32, 0x33, 0x37, 0x30, 0x36, 0x2d, 0x31, 0x64, 0x37, 0x63, 0x2d,
	0x34, 0x33, 0x31, 0x62, 0x2d, 0x62, 0x31, 0x38, 0x32, 0x2d, 0x33, 0x62, 0x39, 0x35, 0x65, 0x62,
	0x36, 0x33, 0x33, 0x31, 0x30, 0x62, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x24, 0x92, 0x41, 0x21, 0x0a, 0x1f, 0xd2, 0x01, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x09, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x09,
	0x92, 0x41, 0x06, 0x4a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xbc, 0x02, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x91, 0x01,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0xb4, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2d, 0x6d, 0x6f, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x92, 0x41, 0x6f, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0x45, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x12, 0x38,
	0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x20, 0x7b, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a,
	0x35, 0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x22, 0x6e, 0x6f, 0x74,
	0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_expectation_api_expectation_api_proto_rawDescOnce sync.Once
	file_expectation_api_expectation_api_proto_rawDescData = file_expectation_api_expectation_api_proto_rawDesc
)

func file_expectation_api_expectation_api_proto_rawDescGZIP() []byte {
	file_expectation_api_expectation_api_proto_rawDescOnce.Do(func() {
		file_expectation_api_expectation_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_expectation_api_expectation_api_proto_rawDescData)
	})
	return file_expectation_api_expectation_api_proto_rawDescData
}

var file_expectation_api_expectation_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_expectation_api_expectation_api_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: statusmok.server.ExpectationService.CreateRequest
	(*CreateResponse)(nil), // 1: statusmok.server.ExpectationService.CreateResponse
	(*DeleteRequest)(nil),  // 2: statusmok.server.ExpectationService.DeleteRequest
	(*DeleteResponse)(nil), // 3: statusmok.server.ExpectationService.DeleteResponse
}
var file_expectation_api_expectation_api_proto_depIdxs = []int32{
	0, // 0: statusmok.server.ExpectationService.ExpectationService.Create:input_type -> statusmok.server.ExpectationService.CreateRequest
	2, // 1: statusmok.server.ExpectationService.ExpectationService.Delete:input_type -> statusmok.server.ExpectationService.DeleteRequest
	1, // 2: statusmok.server.ExpectationService.ExpectationService.Create:output_type -> statusmok.server.ExpectationService.CreateResponse
	3, // 3: statusmok.server.ExpectationService.ExpectationService.Delete:output_type -> statusmok.server.ExpectationService.DeleteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_expectation_api_expectation_api_proto_init() }
func file_expectation_api_expectation_api_proto_init() {
	if File_expectation_api_expectation_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_expectation_api_expectation_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_expectation_api_expectation_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_expectation_api_expectation_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_expectation_api_expectation_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
	file_expectation_api_expectation_api_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_expectation_api_expectation_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_expectation_api_expectation_api_proto_goTypes,
		DependencyIndexes: file_expectation_api_expectation_api_proto_depIdxs,
		MessageInfos:      file_expectation_api_expectation_api_proto_msgTypes,
	}.Build()
	File_expectation_api_expectation_api_proto = out.File
	file_expectation_api_expectation_api_proto_rawDesc = nil
	file_expectation_api_expectation_api_proto_goTypes = nil
	file_expectation_api_expectation_api_proto_depIdxs = nil
}
