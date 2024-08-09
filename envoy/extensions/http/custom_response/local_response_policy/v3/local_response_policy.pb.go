// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: envoy/extensions/http/custom_response/local_response_policy/v3/local_response_policy.proto

package local_response_policyv3

import (
	v3 "github.com/agentio/common-go/envoy/config/core/v3"
	_ "github.com/agentio/common-go/udpa"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Custom response policy to serve a locally stored response to the
// downstream.
type LocalResponsePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional new local reply body text. It will be used
	// in the “%LOCAL_REPLY_BODY%“ command operator in the “body_format“.
	Body *v3.DataSource `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Optional body format to be used for this response. If “body_format“ is  not
	// provided, and “body“ is, the contents of “body“ will be used to populate
	// the body of the local reply without formatting.
	BodyFormat *v3.SubstitutionFormatString `protobuf:"bytes,2,opt,name=body_format,json=bodyFormat,proto3" json:"body_format,omitempty"`
	// The new response status code if specified.
	StatusCode *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// HTTP headers to add to the response. This allows the
	// response policy to append, to add or to override headers of
	// the original response for local body, or the custom response from the
	// remote body, before it is sent to a downstream client.
	ResponseHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
}

func (x *LocalResponsePolicy) Reset() {
	*x = LocalResponsePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalResponsePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalResponsePolicy) ProtoMessage() {}

func (x *LocalResponsePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalResponsePolicy.ProtoReflect.Descriptor instead.
func (*LocalResponsePolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescGZIP(), []int{0}
}

func (x *LocalResponsePolicy) GetBody() *v3.DataSource {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LocalResponsePolicy) GetBodyFormat() *v3.SubstitutionFormatString {
	if x != nil {
		return x.BodyFormat
	}
	return nil
}

func (x *LocalResponsePolicy) GetStatusCode() *wrapperspb.UInt32Value {
	if x != nil {
		return x.StatusCode
	}
	return nil
}

func (x *LocalResponsePolicy) GetResponseHeadersToAdd() []*v3.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

var File_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x33,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02,
	0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x4a, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0b, 0xfa, 0x42, 0x08, 0x2a, 0x06, 0x10, 0xd8, 0x04, 0x28, 0xc8, 0x01, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x69, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x14, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f,
	0x41, 0x64, 0x64, 0x42, 0xef, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6,
	0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x4c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x76, 0x33, 0x42, 0x18, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescData = file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDesc
)

func file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescData)
	})
	return file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDescData
}

var file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_goTypes = []any{
	(*LocalResponsePolicy)(nil),         // 0: envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy
	(*v3.DataSource)(nil),               // 1: envoy.config.core.v3.DataSource
	(*v3.SubstitutionFormatString)(nil), // 2: envoy.config.core.v3.SubstitutionFormatString
	(*wrapperspb.UInt32Value)(nil),      // 3: google.protobuf.UInt32Value
	(*v3.HeaderValueOption)(nil),        // 4: envoy.config.core.v3.HeaderValueOption
}
var file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy.body:type_name -> envoy.config.core.v3.DataSource
	2, // 1: envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy.body_format:type_name -> envoy.config.core.v3.SubstitutionFormatString
	3, // 2: envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy.status_code:type_name -> google.protobuf.UInt32Value
	4, // 3: envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy.response_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_init()
}
func file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_init() {
	if File_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LocalResponsePolicy); i {
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
			RawDescriptor: file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto = out.File
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_rawDesc = nil
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_goTypes = nil
	file_envoy_extensions_http_custom_response_local_response_policy_v3_local_response_policy_proto_depIdxs = nil
}
