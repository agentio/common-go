// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: envoy/api/v2/auth/tls.proto

package auth

import (
	_ "github.com/agentio/common-go/udpa"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpstreamTlsContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common TLS context settings.
	//
	// .. attention::
	//
	//	Server certificate verification is not enabled by default. Configure
	//	:ref:`trusted_ca<envoy_api_field_auth.CertificateValidationContext.trusted_ca>` to enable
	//	verification.
	CommonTlsContext *CommonTlsContext `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	// SNI string to use when creating TLS backend connections.
	Sni string `protobuf:"bytes,2,opt,name=sni,proto3" json:"sni,omitempty"`
	// If true, server-initiated TLS renegotiation will be allowed.
	//
	// .. attention::
	//
	//	TLS renegotiation is considered insecure and shouldn't be used unless absolutely necessary.
	AllowRenegotiation bool `protobuf:"varint,3,opt,name=allow_renegotiation,json=allowRenegotiation,proto3" json:"allow_renegotiation,omitempty"`
	// Maximum number of session keys (Pre-Shared Keys for TLSv1.3+, Session IDs and Session Tickets
	// for TLSv1.2 and older) to store for the purpose of session resumption.
	//
	// Defaults to 1, setting this to 0 disables session resumption.
	MaxSessionKeys *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=max_session_keys,json=maxSessionKeys,proto3" json:"max_session_keys,omitempty"`
}

func (x *UpstreamTlsContext) Reset() {
	*x = UpstreamTlsContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamTlsContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamTlsContext) ProtoMessage() {}

func (x *UpstreamTlsContext) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamTlsContext.ProtoReflect.Descriptor instead.
func (*UpstreamTlsContext) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_auth_tls_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if x != nil {
		return x.CommonTlsContext
	}
	return nil
}

func (x *UpstreamTlsContext) GetSni() string {
	if x != nil {
		return x.Sni
	}
	return ""
}

func (x *UpstreamTlsContext) GetAllowRenegotiation() bool {
	if x != nil {
		return x.AllowRenegotiation
	}
	return false
}

func (x *UpstreamTlsContext) GetMaxSessionKeys() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxSessionKeys
	}
	return nil
}

// [#next-free-field: 8]
type DownstreamTlsContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common TLS context settings.
	CommonTlsContext *CommonTlsContext `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	// If specified, Envoy will reject connections without a valid client
	// certificate.
	RequireClientCertificate *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=require_client_certificate,json=requireClientCertificate,proto3" json:"require_client_certificate,omitempty"`
	// If specified, Envoy will reject connections without a valid and matching SNI.
	// [#not-implemented-hide:]
	RequireSni *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=require_sni,json=requireSni,proto3" json:"require_sni,omitempty"`
	// Types that are assignable to SessionTicketKeysType:
	//
	//	*DownstreamTlsContext_SessionTicketKeys
	//	*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig
	//	*DownstreamTlsContext_DisableStatelessSessionResumption
	SessionTicketKeysType isDownstreamTlsContext_SessionTicketKeysType `protobuf_oneof:"session_ticket_keys_type"`
	// If specified, “session_timeout“ will change the maximum lifetime (in seconds) of the TLS session.
	// Currently this value is used as a hint for the `TLS session ticket lifetime (for TLSv1.2) <https://tools.ietf.org/html/rfc5077#section-5.6>`_.
	// Only seconds can be specified (fractional seconds are ignored).
	SessionTimeout *durationpb.Duration `protobuf:"bytes,6,opt,name=session_timeout,json=sessionTimeout,proto3" json:"session_timeout,omitempty"`
}

func (x *DownstreamTlsContext) Reset() {
	*x = DownstreamTlsContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownstreamTlsContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownstreamTlsContext) ProtoMessage() {}

func (x *DownstreamTlsContext) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownstreamTlsContext.ProtoReflect.Descriptor instead.
func (*DownstreamTlsContext) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_auth_tls_proto_rawDescGZIP(), []int{1}
}

func (x *DownstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if x != nil {
		return x.CommonTlsContext
	}
	return nil
}

func (x *DownstreamTlsContext) GetRequireClientCertificate() *wrapperspb.BoolValue {
	if x != nil {
		return x.RequireClientCertificate
	}
	return nil
}

func (x *DownstreamTlsContext) GetRequireSni() *wrapperspb.BoolValue {
	if x != nil {
		return x.RequireSni
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeysType() isDownstreamTlsContext_SessionTicketKeysType {
	if m != nil {
		return m.SessionTicketKeysType
	}
	return nil
}

func (x *DownstreamTlsContext) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := x.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

func (x *DownstreamTlsContext) GetSessionTicketKeysSdsSecretConfig() *SdsSecretConfig {
	if x, ok := x.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig); ok {
		return x.SessionTicketKeysSdsSecretConfig
	}
	return nil
}

func (x *DownstreamTlsContext) GetDisableStatelessSessionResumption() bool {
	if x, ok := x.GetSessionTicketKeysType().(*DownstreamTlsContext_DisableStatelessSessionResumption); ok {
		return x.DisableStatelessSessionResumption
	}
	return false
}

func (x *DownstreamTlsContext) GetSessionTimeout() *durationpb.Duration {
	if x != nil {
		return x.SessionTimeout
	}
	return nil
}

type isDownstreamTlsContext_SessionTicketKeysType interface {
	isDownstreamTlsContext_SessionTicketKeysType()
}

type DownstreamTlsContext_SessionTicketKeys struct {
	// TLS session ticket key settings.
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,4,opt,name=session_ticket_keys,json=sessionTicketKeys,proto3,oneof"`
}

type DownstreamTlsContext_SessionTicketKeysSdsSecretConfig struct {
	// Config for fetching TLS session ticket keys via SDS API.
	SessionTicketKeysSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,5,opt,name=session_ticket_keys_sds_secret_config,json=sessionTicketKeysSdsSecretConfig,proto3,oneof"`
}

type DownstreamTlsContext_DisableStatelessSessionResumption struct {
	// Config for controlling stateless TLS session resumption: setting this to true will cause the TLS
	// server to not issue TLS session tickets for the purposes of stateless TLS session resumption.
	// If set to false, the TLS server will issue TLS session tickets and encrypt/decrypt them using
	// the keys specified through either :ref:`session_ticket_keys <envoy_api_field_auth.DownstreamTlsContext.session_ticket_keys>`
	// or :ref:`session_ticket_keys_sds_secret_config <envoy_api_field_auth.DownstreamTlsContext.session_ticket_keys_sds_secret_config>`.
	// If this config is set to false and no keys are explicitly configured, the TLS server will issue
	// TLS session tickets and encrypt/decrypt them using an internally-generated and managed key, with the
	// implication that sessions cannot be resumed across hot restarts or on different hosts.
	DisableStatelessSessionResumption bool `protobuf:"varint,7,opt,name=disable_stateless_session_resumption,json=disableStatelessSessionResumption,proto3,oneof"`
}

func (*DownstreamTlsContext_SessionTicketKeys) isDownstreamTlsContext_SessionTicketKeysType() {}

func (*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (*DownstreamTlsContext_DisableStatelessSessionResumption) isDownstreamTlsContext_SessionTicketKeysType() {
}

// TLS context shared by both client and server TLS contexts.
// [#next-free-field: 9]
type CommonTlsContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TLS protocol versions, cipher suites etc.
	TlsParams *TlsParameters `protobuf:"bytes,1,opt,name=tls_params,json=tlsParams,proto3" json:"tls_params,omitempty"`
	// :ref:`Multiple TLS certificates <arch_overview_ssl_cert_select>` can be associated with the
	// same context to allow both RSA and ECDSA certificates.
	//
	// Only a single TLS certificate is supported in client contexts. In server contexts, the first
	// RSA certificate is used for clients that only support RSA and the first ECDSA certificate is
	// used for clients that support ECDSA.
	TlsCertificates []*TlsCertificate `protobuf:"bytes,2,rep,name=tls_certificates,json=tlsCertificates,proto3" json:"tls_certificates,omitempty"`
	// Configs for fetching TLS certificates via SDS API.
	TlsCertificateSdsSecretConfigs []*SdsSecretConfig `protobuf:"bytes,6,rep,name=tls_certificate_sds_secret_configs,json=tlsCertificateSdsSecretConfigs,proto3" json:"tls_certificate_sds_secret_configs,omitempty"`
	// Types that are assignable to ValidationContextType:
	//
	//	*CommonTlsContext_ValidationContext
	//	*CommonTlsContext_ValidationContextSdsSecretConfig
	//	*CommonTlsContext_CombinedValidationContext
	ValidationContextType isCommonTlsContext_ValidationContextType `protobuf_oneof:"validation_context_type"`
	// Supplies the list of ALPN protocols that the listener should expose. In
	// practice this is likely to be set to one of two values (see the
	// :ref:`codec_type
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.codec_type>`
	// parameter in the HTTP connection manager for more information):
	//
	// * "h2,http/1.1" If the listener is going to support both HTTP/2 and HTTP/1.1.
	// * "http/1.1" If the listener is only going to support HTTP/1.1.
	//
	// There is no default for this parameter. If empty, Envoy will not expose ALPN.
	AlpnProtocols []string `protobuf:"bytes,4,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
}

func (x *CommonTlsContext) Reset() {
	*x = CommonTlsContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonTlsContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonTlsContext) ProtoMessage() {}

func (x *CommonTlsContext) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonTlsContext.ProtoReflect.Descriptor instead.
func (*CommonTlsContext) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_auth_tls_proto_rawDescGZIP(), []int{2}
}

func (x *CommonTlsContext) GetTlsParams() *TlsParameters {
	if x != nil {
		return x.TlsParams
	}
	return nil
}

func (x *CommonTlsContext) GetTlsCertificates() []*TlsCertificate {
	if x != nil {
		return x.TlsCertificates
	}
	return nil
}

func (x *CommonTlsContext) GetTlsCertificateSdsSecretConfigs() []*SdsSecretConfig {
	if x != nil {
		return x.TlsCertificateSdsSecretConfigs
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContextType() isCommonTlsContext_ValidationContextType {
	if m != nil {
		return m.ValidationContextType
	}
	return nil
}

func (x *CommonTlsContext) GetValidationContext() *CertificateValidationContext {
	if x, ok := x.GetValidationContextType().(*CommonTlsContext_ValidationContext); ok {
		return x.ValidationContext
	}
	return nil
}

func (x *CommonTlsContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if x, ok := x.GetValidationContextType().(*CommonTlsContext_ValidationContextSdsSecretConfig); ok {
		return x.ValidationContextSdsSecretConfig
	}
	return nil
}

func (x *CommonTlsContext) GetCombinedValidationContext() *CommonTlsContext_CombinedCertificateValidationContext {
	if x, ok := x.GetValidationContextType().(*CommonTlsContext_CombinedValidationContext); ok {
		return x.CombinedValidationContext
	}
	return nil
}

func (x *CommonTlsContext) GetAlpnProtocols() []string {
	if x != nil {
		return x.AlpnProtocols
	}
	return nil
}

type isCommonTlsContext_ValidationContextType interface {
	isCommonTlsContext_ValidationContextType()
}

type CommonTlsContext_ValidationContext struct {
	// How to validate peer certificates.
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext,proto3,oneof"`
}

type CommonTlsContext_ValidationContextSdsSecretConfig struct {
	// Config for fetching validation context via SDS API.
	ValidationContextSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,7,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3,oneof"`
}

type CommonTlsContext_CombinedValidationContext struct {
	// Combined certificate validation context holds a default CertificateValidationContext
	// and SDS config. When SDS server returns dynamic CertificateValidationContext, both dynamic
	// and default CertificateValidationContext are merged into a new CertificateValidationContext
	// for validation. This merge is done by Message::MergeFrom(), so dynamic
	// CertificateValidationContext overwrites singular fields in default
	// CertificateValidationContext, and concatenates repeated fields to default
	// CertificateValidationContext, and logical OR is applied to boolean fields.
	CombinedValidationContext *CommonTlsContext_CombinedCertificateValidationContext `protobuf:"bytes,8,opt,name=combined_validation_context,json=combinedValidationContext,proto3,oneof"`
}

func (*CommonTlsContext_ValidationContext) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_ValidationContextSdsSecretConfig) isCommonTlsContext_ValidationContextType() {
}

func (*CommonTlsContext_CombinedValidationContext) isCommonTlsContext_ValidationContextType() {}

type CommonTlsContext_CombinedCertificateValidationContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How to validate peer certificates.
	DefaultValidationContext *CertificateValidationContext `protobuf:"bytes,1,opt,name=default_validation_context,json=defaultValidationContext,proto3" json:"default_validation_context,omitempty"`
	// Config for fetching validation context via SDS API.
	ValidationContextSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,2,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3" json:"validation_context_sds_secret_config,omitempty"`
}

func (x *CommonTlsContext_CombinedCertificateValidationContext) Reset() {
	*x = CommonTlsContext_CombinedCertificateValidationContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonTlsContext_CombinedCertificateValidationContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonTlsContext_CombinedCertificateValidationContext) ProtoMessage() {}

func (x *CommonTlsContext_CombinedCertificateValidationContext) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_auth_tls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonTlsContext_CombinedCertificateValidationContext.ProtoReflect.Descriptor instead.
func (*CommonTlsContext_CombinedCertificateValidationContext) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_auth_tls_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CommonTlsContext_CombinedCertificateValidationContext) GetDefaultValidationContext() *CertificateValidationContext {
	if x != nil {
		return x.DefaultValidationContext
	}
	return nil
}

func (x *CommonTlsContext_CombinedCertificateValidationContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if x != nil {
		return x.ValidationContextSdsSecretConfig
	}
	return nil
}

var File_envoy_api_v2_auth_tls_proto protoreflect.FileDescriptor

var file_envoy_api_v2_auth_tls_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x1a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x51, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x73, 0x6e, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x28, 0xff, 0x01, 0x52, 0x03, 0x73, 0x6e, 0x69, 0x12, 0x2f,
	0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x97, 0x05, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x51, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x58, 0x0a, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x6e, 0x69, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x53, 0x6e, 0x69, 0x12, 0x59, 0x0a, 0x13, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6c, 0x73, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x48, 0x00, 0x52, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x75, 0x0a, 0x25, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x73, 0x64, 0x73,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x20, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x64, 0x73,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x24,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x73,
	0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x21, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x54, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0xaa, 0x01, 0x0a, 0x1a, 0x06, 0x08, 0x80, 0x80,
	0x80, 0x80, 0x10, 0x32, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x1a, 0x0a, 0x18, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xe8, 0x07, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x09, 0x74, 0x6c,
	0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x74, 0x6c, 0x73, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x78, 0x0a, 0x22, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x1e, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x60, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x11,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x74, 0x0a, 0x24, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8a, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x19, 0x63, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x70, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c,
	0x70, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x9d, 0x02, 0x0a, 0x24,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x77, 0x0a, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x7c, 0x0a,
	0x24, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x73, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x64, 0x73, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x19, 0x0a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x42, 0x96, 0x01, 0xf2,
	0x98, 0xfe, 0x8f, 0x05, 0x2b, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x42, 0x08, 0x54, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_auth_tls_proto_rawDescOnce sync.Once
	file_envoy_api_v2_auth_tls_proto_rawDescData = file_envoy_api_v2_auth_tls_proto_rawDesc
)

func file_envoy_api_v2_auth_tls_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_auth_tls_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_auth_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_auth_tls_proto_rawDescData)
	})
	return file_envoy_api_v2_auth_tls_proto_rawDescData
}

var file_envoy_api_v2_auth_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_api_v2_auth_tls_proto_goTypes = []any{
	(*UpstreamTlsContext)(nil),                                    // 0: envoy.api.v2.auth.UpstreamTlsContext
	(*DownstreamTlsContext)(nil),                                  // 1: envoy.api.v2.auth.DownstreamTlsContext
	(*CommonTlsContext)(nil),                                      // 2: envoy.api.v2.auth.CommonTlsContext
	(*CommonTlsContext_CombinedCertificateValidationContext)(nil), // 3: envoy.api.v2.auth.CommonTlsContext.CombinedCertificateValidationContext
	(*wrapperspb.UInt32Value)(nil),                                // 4: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),                                  // 5: google.protobuf.BoolValue
	(*TlsSessionTicketKeys)(nil),                                  // 6: envoy.api.v2.auth.TlsSessionTicketKeys
	(*SdsSecretConfig)(nil),                                       // 7: envoy.api.v2.auth.SdsSecretConfig
	(*durationpb.Duration)(nil),                                   // 8: google.protobuf.Duration
	(*TlsParameters)(nil),                                         // 9: envoy.api.v2.auth.TlsParameters
	(*TlsCertificate)(nil),                                        // 10: envoy.api.v2.auth.TlsCertificate
	(*CertificateValidationContext)(nil),                          // 11: envoy.api.v2.auth.CertificateValidationContext
}
var file_envoy_api_v2_auth_tls_proto_depIdxs = []int32{
	2,  // 0: envoy.api.v2.auth.UpstreamTlsContext.common_tls_context:type_name -> envoy.api.v2.auth.CommonTlsContext
	4,  // 1: envoy.api.v2.auth.UpstreamTlsContext.max_session_keys:type_name -> google.protobuf.UInt32Value
	2,  // 2: envoy.api.v2.auth.DownstreamTlsContext.common_tls_context:type_name -> envoy.api.v2.auth.CommonTlsContext
	5,  // 3: envoy.api.v2.auth.DownstreamTlsContext.require_client_certificate:type_name -> google.protobuf.BoolValue
	5,  // 4: envoy.api.v2.auth.DownstreamTlsContext.require_sni:type_name -> google.protobuf.BoolValue
	6,  // 5: envoy.api.v2.auth.DownstreamTlsContext.session_ticket_keys:type_name -> envoy.api.v2.auth.TlsSessionTicketKeys
	7,  // 6: envoy.api.v2.auth.DownstreamTlsContext.session_ticket_keys_sds_secret_config:type_name -> envoy.api.v2.auth.SdsSecretConfig
	8,  // 7: envoy.api.v2.auth.DownstreamTlsContext.session_timeout:type_name -> google.protobuf.Duration
	9,  // 8: envoy.api.v2.auth.CommonTlsContext.tls_params:type_name -> envoy.api.v2.auth.TlsParameters
	10, // 9: envoy.api.v2.auth.CommonTlsContext.tls_certificates:type_name -> envoy.api.v2.auth.TlsCertificate
	7,  // 10: envoy.api.v2.auth.CommonTlsContext.tls_certificate_sds_secret_configs:type_name -> envoy.api.v2.auth.SdsSecretConfig
	11, // 11: envoy.api.v2.auth.CommonTlsContext.validation_context:type_name -> envoy.api.v2.auth.CertificateValidationContext
	7,  // 12: envoy.api.v2.auth.CommonTlsContext.validation_context_sds_secret_config:type_name -> envoy.api.v2.auth.SdsSecretConfig
	3,  // 13: envoy.api.v2.auth.CommonTlsContext.combined_validation_context:type_name -> envoy.api.v2.auth.CommonTlsContext.CombinedCertificateValidationContext
	11, // 14: envoy.api.v2.auth.CommonTlsContext.CombinedCertificateValidationContext.default_validation_context:type_name -> envoy.api.v2.auth.CertificateValidationContext
	7,  // 15: envoy.api.v2.auth.CommonTlsContext.CombinedCertificateValidationContext.validation_context_sds_secret_config:type_name -> envoy.api.v2.auth.SdsSecretConfig
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_auth_tls_proto_init() }
func file_envoy_api_v2_auth_tls_proto_init() {
	if File_envoy_api_v2_auth_tls_proto != nil {
		return
	}
	file_envoy_api_v2_auth_common_proto_init()
	file_envoy_api_v2_auth_secret_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_auth_tls_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UpstreamTlsContext); i {
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
		file_envoy_api_v2_auth_tls_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DownstreamTlsContext); i {
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
		file_envoy_api_v2_auth_tls_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CommonTlsContext); i {
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
		file_envoy_api_v2_auth_tls_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CommonTlsContext_CombinedCertificateValidationContext); i {
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
	file_envoy_api_v2_auth_tls_proto_msgTypes[1].OneofWrappers = []any{
		(*DownstreamTlsContext_SessionTicketKeys)(nil),
		(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig)(nil),
		(*DownstreamTlsContext_DisableStatelessSessionResumption)(nil),
	}
	file_envoy_api_v2_auth_tls_proto_msgTypes[2].OneofWrappers = []any{
		(*CommonTlsContext_ValidationContext)(nil),
		(*CommonTlsContext_ValidationContextSdsSecretConfig)(nil),
		(*CommonTlsContext_CombinedValidationContext)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_api_v2_auth_tls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_auth_tls_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_auth_tls_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_auth_tls_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_auth_tls_proto = out.File
	file_envoy_api_v2_auth_tls_proto_rawDesc = nil
	file_envoy_api_v2_auth_tls_proto_goTypes = nil
	file_envoy_api_v2_auth_tls_proto_depIdxs = nil
}
