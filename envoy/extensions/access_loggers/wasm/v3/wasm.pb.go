// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: envoy/extensions/access_loggers/wasm/v3/wasm.proto

package wasmv3

import (
	v3 "github.com/agentio/common-go/envoy/extensions/wasm/v3"
	_ "github.com/agentio/common-go/udpa"
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

// Custom configuration for an :ref:`AccessLog <envoy_v3_api_msg_config.accesslog.v3.AccessLog>`
// that calls into a WASM VM. Configures the built-in “envoy.access_loggers.wasm“
// AccessLog.
type WasmAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *v3.PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *WasmAccessLog) Reset() {
	*x = WasmAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmAccessLog) ProtoMessage() {}

func (x *WasmAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmAccessLog.ProtoReflect.Descriptor instead.
func (*WasmAccessLog) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *WasmAccessLog) GetConfig() *v3.PluginConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_extensions_access_loggers_wasm_v3_wasm_proto protoreflect.FileDescriptor

var file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x23, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x0d, 0x57, 0x61, 0x73, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x6f, 0x67, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x99, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x35, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x57, 0x61, 0x73, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2f,
	0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x3b, 0x77, 0x61, 0x73, 0x6d, 0x76, 0x33, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescOnce sync.Once
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescData = file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDesc
)

func file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescGZIP() []byte {
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescData)
	})
	return file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDescData
}

var file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_goTypes = []any{
	(*WasmAccessLog)(nil),   // 0: envoy.extensions.access_loggers.wasm.v3.WasmAccessLog
	(*v3.PluginConfig)(nil), // 1: envoy.extensions.wasm.v3.PluginConfig
}
var file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.access_loggers.wasm.v3.WasmAccessLog.config:type_name -> envoy.extensions.wasm.v3.PluginConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_init() }
func file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_init() {
	if File_envoy_extensions_access_loggers_wasm_v3_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WasmAccessLog); i {
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
			RawDescriptor: file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_msgTypes,
	}.Build()
	File_envoy_extensions_access_loggers_wasm_v3_wasm_proto = out.File
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_rawDesc = nil
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_goTypes = nil
	file_envoy_extensions_access_loggers_wasm_v3_wasm_proto_depIdxs = nil
}
