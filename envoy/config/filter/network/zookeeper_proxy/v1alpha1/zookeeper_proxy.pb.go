// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto

package v1alpha1

import (
	_ "github.com/agentio/common-go/udpa"
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

type ZooKeeperProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_zookeeper_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing ZooKeeper access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Messages — requests, responses and events — that are bigger than this value will
	// be ignored. If it is not set, the default value is 1Mb.
	//
	// The value here should match the jute.maxbuffer property in your cluster configuration:
	//
	// https://zookeeper.apache.org/doc/r3.4.10/zookeeperAdmin.html#Unsafe+Options
	//
	// if that is set. If it isn't, ZooKeeper's default is also 1Mb.
	MaxPacketBytes *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
}

func (x *ZooKeeperProxy) Reset() {
	*x = ZooKeeperProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZooKeeperProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZooKeeperProxy) ProtoMessage() {}

func (x *ZooKeeperProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZooKeeperProxy.ProtoReflect.Descriptor instead.
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *ZooKeeperProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ZooKeeperProxy) GetAccessLog() string {
	if x != nil {
		return x.AccessLog
	}
	return ""
}

func (x *ZooKeeperProxy) GetMaxPacketBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxPacketBytes
	}
	return nil
}

var File_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x7a, 0x6f,
	0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0e, 0x5a,
	0x6f, 0x6f, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x46, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e,
	0x6d, 0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0xf1,
	0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x35, 0x12, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x01, 0x0a, 0x42, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13, 0x5a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescData = file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDesc
)

func file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescData)
	})
	return file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDescData
}

var file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_goTypes = []any{
	(*ZooKeeperProxy)(nil),         // 0: envoy.config.filter.network.zookeeper_proxy.v1alpha1.ZooKeeperProxy
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
}
var file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.network.zookeeper_proxy.v1alpha1.ZooKeeperProxy.max_packet_bytes:type_name -> google.protobuf.UInt32Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_init() }
func file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_init() {
	if File_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ZooKeeperProxy); i {
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
			RawDescriptor: file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto = out.File
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_rawDesc = nil
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_goTypes = nil
	file_envoy_config_filter_network_zookeeper_proxy_v1alpha1_zookeeper_proxy_proto_depIdxs = nil
}
