// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: route-api/route_api.proto

package route_api

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

type RouteType int32

const (
	RouteType_ROUTE_TYPE_UNSPECIFIED RouteType = 0
	RouteType_ROUTE_TYPE_REQ_RESP    RouteType = 1
	RouteType_ROUTE_TYPE_WEBSOCKET   RouteType = 2
)

// Enum value maps for RouteType.
var (
	RouteType_name = map[int32]string{
		0: "ROUTE_TYPE_UNSPECIFIED",
		1: "ROUTE_TYPE_REQ_RESP",
		2: "ROUTE_TYPE_WEBSOCKET",
	}
	RouteType_value = map[string]int32{
		"ROUTE_TYPE_UNSPECIFIED": 0,
		"ROUTE_TYPE_REQ_RESP":    1,
		"ROUTE_TYPE_WEBSOCKET":   2,
	}
)

func (x RouteType) Enum() *RouteType {
	p := new(RouteType)
	*p = x
	return p
}

func (x RouteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteType) Descriptor() protoreflect.EnumDescriptor {
	return file_route_api_route_api_proto_enumTypes[0].Descriptor()
}

func (RouteType) Type() protoreflect.EnumType {
	return &file_route_api_route_api_proto_enumTypes[0]
}

func (x RouteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteType.Descriptor instead.
func (RouteType) EnumDescriptor() ([]byte, []int) {
	return file_route_api_route_api_proto_rawDescGZIP(), []int{0}
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server name is a unique identifier of the mock server.
	ServerName string `protobuf:"bytes,1,opt,name=server_name,proto3" json:"server_name,omitempty"`
	// URL is a relative URI of the route.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"` // The route type.
	// 1: ReqResp
	// 2: WebSocket
	Type RouteType `protobuf:"varint,3,opt,name=type,proto3,enum=statusmok.server.RouteService.RouteType" json:"type,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_api_proto_msgTypes[0]
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
	return file_route_api_route_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *CreateRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateRequest) GetType() RouteType {
	if x != nil {
		return x.Type
	}
	return RouteType_ROUTE_TYPE_UNSPECIFIED
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
		mi := &file_route_api_route_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_api_proto_msgTypes[1]
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
	return file_route_api_route_api_proto_rawDescGZIP(), []int{1}
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

	// Server name is a unique identifier of the mock server.
	ServerName string `protobuf:"bytes,1,opt,name=server_name,proto3" json:"server_name,omitempty"`
	// URL is a relative URI of the route.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_api_proto_msgTypes[2]
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
	return file_route_api_route_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *DeleteRequest) GetUrl() string {
	if x != nil {
		return x.Url
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
		mi := &file_route_api_route_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_api_proto_msgTypes[3]
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
	return file_route_api_route_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_route_api_route_api_proto protoreflect.FileDescriptor

var file_route_api_route_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x17, 0x4a, 0x12, 0x22,
	0x68, 0x74, 0x74, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x38, 0x30, 0x38, 0x30,
	0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x92, 0x41, 0x10, 0x4a, 0x0b, 0x22, 0x2f,
	0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x90, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x3a, 0x20, 0x92, 0x41, 0x1d, 0x0a, 0x1b, 0xd2, 0x01, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x03, 0x75, 0x72, 0x6c, 0xd2, 0x01, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0xa0, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x17, 0x4a, 0x12, 0x22, 0x68, 0x74,
	0x74, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x38, 0x30, 0x38, 0x30, 0x22, 0x80,
	0x01, 0x01, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x92, 0x41, 0x10, 0x4a, 0x0b, 0x22, 0x2f, 0x73, 0x6f,
	0x6d, 0x65, 0x2d, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90,
	0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x3a, 0x19, 0x92, 0x41, 0x16, 0x0a, 0x14, 0xd2, 0x01,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x35, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x09, 0x92, 0x41, 0x06, 0x4a, 0x04, 0x74, 0x72, 0x75, 0x65,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x5a, 0x0a, 0x09, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x4f, 0x43,
	0x4b, 0x45, 0x54, 0x10, 0x02, 0x32, 0x90, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x7f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x6d, 0x6f, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0xa8, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2d, 0x6d,
	0x6f, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x92, 0x41, 0x6f, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x45, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x12, 0x38, 0x0a, 0x14, 0x1a, 0x12, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x20, 0x7b, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x35, 0x2c, 0x22, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x22, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x22, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_api_route_api_proto_rawDescOnce sync.Once
	file_route_api_route_api_proto_rawDescData = file_route_api_route_api_proto_rawDesc
)

func file_route_api_route_api_proto_rawDescGZIP() []byte {
	file_route_api_route_api_proto_rawDescOnce.Do(func() {
		file_route_api_route_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_api_route_api_proto_rawDescData)
	})
	return file_route_api_route_api_proto_rawDescData
}

var file_route_api_route_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_route_api_route_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_route_api_route_api_proto_goTypes = []interface{}{
	(RouteType)(0),         // 0: statusmok.server.RouteService.RouteType
	(*CreateRequest)(nil),  // 1: statusmok.server.RouteService.CreateRequest
	(*CreateResponse)(nil), // 2: statusmok.server.RouteService.CreateResponse
	(*DeleteRequest)(nil),  // 3: statusmok.server.RouteService.DeleteRequest
	(*DeleteResponse)(nil), // 4: statusmok.server.RouteService.DeleteResponse
}
var file_route_api_route_api_proto_depIdxs = []int32{
	0, // 0: statusmok.server.RouteService.CreateRequest.type:type_name -> statusmok.server.RouteService.RouteType
	1, // 1: statusmok.server.RouteService.RouteService.Create:input_type -> statusmok.server.RouteService.CreateRequest
	3, // 2: statusmok.server.RouteService.RouteService.Delete:input_type -> statusmok.server.RouteService.DeleteRequest
	2, // 3: statusmok.server.RouteService.RouteService.Create:output_type -> statusmok.server.RouteService.CreateResponse
	4, // 4: statusmok.server.RouteService.RouteService.Delete:output_type -> statusmok.server.RouteService.DeleteResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_route_api_route_api_proto_init() }
func file_route_api_route_api_proto_init() {
	if File_route_api_route_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_api_route_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_route_api_route_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_route_api_route_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_route_api_route_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_api_route_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_api_route_api_proto_goTypes,
		DependencyIndexes: file_route_api_route_api_proto_depIdxs,
		EnumInfos:         file_route_api_route_api_proto_enumTypes,
		MessageInfos:      file_route_api_route_api_proto_msgTypes,
	}.Build()
	File_route_api_route_api_proto = out.File
	file_route_api_route_api_proto_rawDesc = nil
	file_route_api_route_api_proto_goTypes = nil
	file_route_api_route_api_proto_depIdxs = nil
}
