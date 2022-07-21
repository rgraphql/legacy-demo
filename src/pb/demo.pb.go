// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.3
// source: github.com/rgraphql/magellan-soyuz-demo/src/pb/demo.proto

package demo

import (
	reflect "reflect"
	sync "sync"

	rgraphql "github.com/rgraphql/rgraphql"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RPC is the rpc type enum.
type RPC int32

const (
	RPC_RPC_Unknown           RPC = 0
	RPC_RPC_Ping              RPC = 1
	RPC_RPC_RGQLClientMessage RPC = 2
	RPC_RPC_RGQLServerMessage RPC = 3
)

// Enum value maps for RPC.
var (
	RPC_name = map[int32]string{
		0: "RPC_Unknown",
		1: "RPC_Ping",
		2: "RPC_RGQLClientMessage",
		3: "RPC_RGQLServerMessage",
	}
	RPC_value = map[string]int32{
		"RPC_Unknown":           0,
		"RPC_Ping":              1,
		"RPC_RGQLClientMessage": 2,
		"RPC_RGQLServerMessage": 3,
	}
)

func (x RPC) Enum() *RPC {
	p := new(RPC)
	*p = x
	return p
}

func (x RPC) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RPC) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_enumTypes[0].Descriptor()
}

func (RPC) Type() protoreflect.EnumType {
	return &file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_enumTypes[0]
}

func (x RPC) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RPC.Descriptor instead.
func (RPC) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescGZIP(), []int{0}
}

// RPCMessage contains a RPC over the message bus.
type RPCMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RPCId is the RPC identifier.
	RpcId RPC `protobuf:"varint,1,opt,name=rpc_id,json=rpcId,proto3,enum=demo.RPC" json:"rpc_id,omitempty"`
	// RGQLClientMessage is a client message.
	RgqlClientMessage *rgraphql.RGQLClientMessage `protobuf:"bytes,2,opt,name=rgql_client_message,json=rgqlClientMessage,proto3" json:"rgql_client_message,omitempty"`
	// RGQLServerMessage is a server message.
	RgqlServerMessage *rgraphql.RGQLServerMessage `protobuf:"bytes,3,opt,name=rgql_server_message,json=rgqlServerMessage,proto3" json:"rgql_server_message,omitempty"`
}

func (x *RPCMessage) Reset() {
	*x = RPCMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMessage) ProtoMessage() {}

func (x *RPCMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMessage.ProtoReflect.Descriptor instead.
func (*RPCMessage) Descriptor() ([]byte, []int) {
	return file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescGZIP(), []int{0}
}

func (x *RPCMessage) GetRpcId() RPC {
	if x != nil {
		return x.RpcId
	}
	return RPC_RPC_Unknown
}

func (x *RPCMessage) GetRgqlClientMessage() *rgraphql.RGQLClientMessage {
	if x != nil {
		return x.RgqlClientMessage
	}
	return nil
}

func (x *RPCMessage) GetRgqlServerMessage() *rgraphql.RGQLServerMessage {
	if x != nil {
		return x.RgqlServerMessage
	}
	return nil
}

var File_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto protoreflect.FileDescriptor

var file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2d, 0x73,
	0x6f, 0x79, 0x75, 0x7a, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62,
	0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d,
	0x6f, 0x1a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f,
	0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8,
	0x01, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x06, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x52, 0x50, 0x43, 0x52, 0x05, 0x72, 0x70, 0x63, 0x49, 0x64, 0x12,
	0x4b, 0x0a, 0x13, 0x72, 0x67, 0x71, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x52, 0x47, 0x51, 0x4c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x11, 0x72, 0x67, 0x71, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x13,
	0x72, 0x67, 0x71, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x2e, 0x52, 0x47, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x11, 0x72, 0x67, 0x71, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x5a, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x50, 0x43, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x50, 0x43, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x52, 0x50, 0x43, 0x5f, 0x52, 0x47, 0x51, 0x4c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x50,
	0x43, 0x5f, 0x52, 0x47, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescOnce sync.Once
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescData = file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDesc
)

func file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescGZIP() []byte {
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescOnce.Do(func() {
		file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescData)
	})
	return file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDescData
}

var file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_goTypes = []interface{}{
	(RPC)(0),                           // 0: demo.RPC
	(*RPCMessage)(nil),                 // 1: demo.RPCMessage
	(*rgraphql.RGQLClientMessage)(nil), // 2: rgraphql.RGQLClientMessage
	(*rgraphql.RGQLServerMessage)(nil), // 3: rgraphql.RGQLServerMessage
}
var file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_depIdxs = []int32{
	0, // 0: demo.RPCMessage.rpc_id:type_name -> demo.RPC
	2, // 1: demo.RPCMessage.rgql_client_message:type_name -> rgraphql.RGQLClientMessage
	3, // 2: demo.RPCMessage.rgql_server_message:type_name -> rgraphql.RGQLServerMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_init() }
func file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_init() {
	if File_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMessage); i {
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
			RawDescriptor: file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_goTypes,
		DependencyIndexes: file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_depIdxs,
		EnumInfos:         file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_enumTypes,
		MessageInfos:      file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_msgTypes,
	}.Build()
	File_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto = out.File
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_rawDesc = nil
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_goTypes = nil
	file_github_com_rgraphql_magellan_soyuz_demo_src_pb_demo_proto_depIdxs = nil
}
