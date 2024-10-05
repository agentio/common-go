// Copyright 2024 Agent IO. All Rights Reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: io/agent/proxy/v1/proxy.proto

package proxyv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Auth      *Auth       `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Upstreams []*Upstream `protobuf:"bytes,3,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Profile) GetUpstreams() []*Upstream {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AuthMechanism:
	//
	//	*Auth_Oidc
	AuthMechanism isAuth_AuthMechanism `protobuf_oneof:"auth_mechanism"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{1}
}

func (m *Auth) GetAuthMechanism() isAuth_AuthMechanism {
	if m != nil {
		return m.AuthMechanism
	}
	return nil
}

func (x *Auth) GetOidc() *OIDC {
	if x, ok := x.GetAuthMechanism().(*Auth_Oidc); ok {
		return x.Oidc
	}
	return nil
}

type isAuth_AuthMechanism interface {
	isAuth_AuthMechanism()
}

type Auth_Oidc struct {
	Oidc *OIDC `protobuf:"bytes,1,opt,name=oidc,proto3,oneof"`
}

func (*Auth_Oidc) isAuth_AuthMechanism() {}

type OIDC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwksUri string `protobuf:"bytes,1,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
}

func (x *OIDC) Reset() {
	*x = OIDC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OIDC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OIDC) ProtoMessage() {}

func (x *OIDC) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OIDC.ProtoReflect.Descriptor instead.
func (*OIDC) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *OIDC) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

type Upstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host       string       `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Credential *Credential  `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential,omitempty"`
	Operations []*Operation `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *Upstream) Reset() {
	*x = Upstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upstream) ProtoMessage() {}

func (x *Upstream) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upstream.ProtoReflect.Descriptor instead.
func (*Upstream) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *Upstream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Upstream) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Upstream) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *Upstream) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expires *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *Credential) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Credential) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_io_agent_proxy_v1_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_io_agent_proxy_v1_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *Operation) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Operation) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_io_agent_proxy_v1_proxy_proto protoreflect.FileDescriptor

var file_io_agent_proxy_v1_proxy_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x12, 0x39, 0x0a, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x47, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x69,
	0x64, 0x63, 0x42, 0x10, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x73, 0x6d, 0x22, 0x21, 0x0a, 0x04, 0x4f, 0x49, 0x44, 0x43, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x22, 0xaf, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_agent_proxy_v1_proxy_proto_rawDescOnce sync.Once
	file_io_agent_proxy_v1_proxy_proto_rawDescData = file_io_agent_proxy_v1_proxy_proto_rawDesc
)

func file_io_agent_proxy_v1_proxy_proto_rawDescGZIP() []byte {
	file_io_agent_proxy_v1_proxy_proto_rawDescOnce.Do(func() {
		file_io_agent_proxy_v1_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_agent_proxy_v1_proxy_proto_rawDescData)
	})
	return file_io_agent_proxy_v1_proxy_proto_rawDescData
}

var file_io_agent_proxy_v1_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_io_agent_proxy_v1_proxy_proto_goTypes = []any{
	(*Profile)(nil),               // 0: io.agent.proxy.v1.Profile
	(*Auth)(nil),                  // 1: io.agent.proxy.v1.Auth
	(*OIDC)(nil),                  // 2: io.agent.proxy.v1.OIDC
	(*Upstream)(nil),              // 3: io.agent.proxy.v1.Upstream
	(*Credential)(nil),            // 4: io.agent.proxy.v1.Credential
	(*Operation)(nil),             // 5: io.agent.proxy.v1.Operation
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_io_agent_proxy_v1_proxy_proto_depIdxs = []int32{
	1, // 0: io.agent.proxy.v1.Profile.auth:type_name -> io.agent.proxy.v1.Auth
	3, // 1: io.agent.proxy.v1.Profile.upstreams:type_name -> io.agent.proxy.v1.Upstream
	2, // 2: io.agent.proxy.v1.Auth.oidc:type_name -> io.agent.proxy.v1.OIDC
	4, // 3: io.agent.proxy.v1.Upstream.credential:type_name -> io.agent.proxy.v1.Credential
	5, // 4: io.agent.proxy.v1.Upstream.operations:type_name -> io.agent.proxy.v1.Operation
	6, // 5: io.agent.proxy.v1.Credential.expires:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_io_agent_proxy_v1_proxy_proto_init() }
func file_io_agent_proxy_v1_proxy_proto_init() {
	if File_io_agent_proxy_v1_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_io_agent_proxy_v1_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Profile); i {
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
		file_io_agent_proxy_v1_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Auth); i {
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
		file_io_agent_proxy_v1_proxy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OIDC); i {
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
		file_io_agent_proxy_v1_proxy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Upstream); i {
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
		file_io_agent_proxy_v1_proxy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Credential); i {
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
		file_io_agent_proxy_v1_proxy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Operation); i {
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
	file_io_agent_proxy_v1_proxy_proto_msgTypes[1].OneofWrappers = []any{
		(*Auth_Oidc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_agent_proxy_v1_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_agent_proxy_v1_proxy_proto_goTypes,
		DependencyIndexes: file_io_agent_proxy_v1_proxy_proto_depIdxs,
		MessageInfos:      file_io_agent_proxy_v1_proxy_proto_msgTypes,
	}.Build()
	File_io_agent_proxy_v1_proxy_proto = out.File
	file_io_agent_proxy_v1_proxy_proto_rawDesc = nil
	file_io_agent_proxy_v1_proxy_proto_goTypes = nil
	file_io_agent_proxy_v1_proxy_proto_depIdxs = nil
}
