// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: omni/specs/siderolink.proto

package specs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SiderolinkConfigSpec describes siderolink wireguard server state to persist it across restarts.
type SiderolinkConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey         string `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey          string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	WireguardEndpoint  string `protobuf:"bytes,3,opt,name=wireguard_endpoint,json=wireguardEndpoint,proto3" json:"wireguard_endpoint,omitempty"`
	Subnet             string `protobuf:"bytes,5,opt,name=subnet,proto3" json:"subnet,omitempty"`
	ServerAddress      string `protobuf:"bytes,6,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	JoinToken          string `protobuf:"bytes,7,opt,name=join_token,json=joinToken,proto3" json:"join_token,omitempty"`
	AdvertisedEndpoint string `protobuf:"bytes,8,opt,name=advertised_endpoint,json=advertisedEndpoint,proto3" json:"advertised_endpoint,omitempty"`
}

func (x *SiderolinkConfigSpec) Reset() {
	*x = SiderolinkConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_siderolink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiderolinkConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiderolinkConfigSpec) ProtoMessage() {}

func (x *SiderolinkConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_siderolink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiderolinkConfigSpec.ProtoReflect.Descriptor instead.
func (*SiderolinkConfigSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_siderolink_proto_rawDescGZIP(), []int{0}
}

func (x *SiderolinkConfigSpec) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetWireguardEndpoint() string {
	if x != nil {
		return x.WireguardEndpoint
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetJoinToken() string {
	if x != nil {
		return x.JoinToken
	}
	return ""
}

func (x *SiderolinkConfigSpec) GetAdvertisedEndpoint() string {
	if x != nil {
		return x.AdvertisedEndpoint
	}
	return ""
}

// SiderolinkConnectionSpec describes each node connection information.
type SiderolinkSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeSubnet      string `protobuf:"bytes,1,opt,name=node_subnet,json=nodeSubnet,proto3" json:"node_subnet,omitempty"`
	NodePublicKey   string `protobuf:"bytes,2,opt,name=node_public_key,json=nodePublicKey,proto3" json:"node_public_key,omitempty"`
	LastEndpoint    string `protobuf:"bytes,3,opt,name=last_endpoint,json=lastEndpoint,proto3" json:"last_endpoint,omitempty"`
	Connected       bool   `protobuf:"varint,4,opt,name=connected,proto3" json:"connected,omitempty"`
	VirtualAddrport string `protobuf:"bytes,7,opt,name=virtual_addrport,json=virtualAddrport,proto3" json:"virtual_addrport,omitempty"`
	// RemoteAddr is the machine address how it's visible from Omni
	// it is determined by reading X-Real-IP header coming from the gRPC API.
	RemoteAddr string `protobuf:"bytes,8,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
}

func (x *SiderolinkSpec) Reset() {
	*x = SiderolinkSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_siderolink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiderolinkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiderolinkSpec) ProtoMessage() {}

func (x *SiderolinkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_siderolink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiderolinkSpec.ProtoReflect.Descriptor instead.
func (*SiderolinkSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_siderolink_proto_rawDescGZIP(), []int{1}
}

func (x *SiderolinkSpec) GetNodeSubnet() string {
	if x != nil {
		return x.NodeSubnet
	}
	return ""
}

func (x *SiderolinkSpec) GetNodePublicKey() string {
	if x != nil {
		return x.NodePublicKey
	}
	return ""
}

func (x *SiderolinkSpec) GetLastEndpoint() string {
	if x != nil {
		return x.LastEndpoint
	}
	return ""
}

func (x *SiderolinkSpec) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *SiderolinkSpec) GetVirtualAddrport() string {
	if x != nil {
		return x.VirtualAddrport
	}
	return ""
}

func (x *SiderolinkSpec) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

// SiderolinkConnectionSpec describes each node connection information.
type SiderolinkCounterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BytesReceived int64                  `protobuf:"varint,1,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	BytesSent     int64                  `protobuf:"varint,2,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	LastAlive     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_alive,json=lastAlive,proto3" json:"last_alive,omitempty"`
}

func (x *SiderolinkCounterSpec) Reset() {
	*x = SiderolinkCounterSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_siderolink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiderolinkCounterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiderolinkCounterSpec) ProtoMessage() {}

func (x *SiderolinkCounterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_siderolink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiderolinkCounterSpec.ProtoReflect.Descriptor instead.
func (*SiderolinkCounterSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_siderolink_proto_rawDescGZIP(), []int{2}
}

func (x *SiderolinkCounterSpec) GetBytesReceived() int64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

func (x *SiderolinkCounterSpec) GetBytesSent() int64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *SiderolinkCounterSpec) GetLastAlive() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAlive
	}
	return nil
}

// ConnectionParamsSpec describes generated kernel parameters for connecting
// the Talos node to Omni.
type ConnectionParamsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Args keeps the generated kernel arguments string.
	Args string `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"`
	// ApiEndpoint is the service gRPC API endpoint (external domain/ip, schema and port).
	ApiEndpoint string `protobuf:"bytes,2,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
	// WireguardEndpoint is the service IP visible from the internal SideroLink network.
	WireguardEndpoint string `protobuf:"bytes,3,opt,name=wireguard_endpoint,json=wireguardEndpoint,proto3" json:"wireguard_endpoint,omitempty"`
	// JoinToken is a join token required to connect to SideroLink.
	JoinToken string `protobuf:"bytes,4,opt,name=join_token,json=joinToken,proto3" json:"join_token,omitempty"`
	// UseGRPCTunnel is a flag to enable gRPC tunneling.
	UseGrpcTunnel bool `protobuf:"varint,5,opt,name=use_grpc_tunnel,json=useGrpcTunnel,proto3" json:"use_grpc_tunnel,omitempty"`
}

func (x *ConnectionParamsSpec) Reset() {
	*x = ConnectionParamsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_specs_siderolink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionParamsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionParamsSpec) ProtoMessage() {}

func (x *ConnectionParamsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_siderolink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionParamsSpec.ProtoReflect.Descriptor instead.
func (*ConnectionParamsSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_siderolink_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionParamsSpec) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *ConnectionParamsSpec) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *ConnectionParamsSpec) GetWireguardEndpoint() string {
	if x != nil {
		return x.WireguardEndpoint
	}
	return ""
}

func (x *ConnectionParamsSpec) GetJoinToken() string {
	if x != nil {
		return x.JoinToken
	}
	return ""
}

func (x *ConnectionParamsSpec) GetUseGrpcTunnel() bool {
	if x != nil {
		return x.UseGrpcTunnel
	}
	return false
}

var File_omni_specs_siderolink_proto protoreflect.FileDescriptor

var file_omni_specs_siderolink_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x14, 0x53, 0x69, 0x64, 0x65, 0x72, 0x6f,
	0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x12, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x69, 0x72, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xf4, 0x01, 0x0a,
	0x0e, 0x53, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x07, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x53, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69,
	0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x22, 0xc3,
	0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x70, 0x69, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x69, 0x72, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0f,
	0x75, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x6d,
	0x6e, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d,
	0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omni_specs_siderolink_proto_rawDescOnce sync.Once
	file_omni_specs_siderolink_proto_rawDescData = file_omni_specs_siderolink_proto_rawDesc
)

func file_omni_specs_siderolink_proto_rawDescGZIP() []byte {
	file_omni_specs_siderolink_proto_rawDescOnce.Do(func() {
		file_omni_specs_siderolink_proto_rawDescData = protoimpl.X.CompressGZIP(file_omni_specs_siderolink_proto_rawDescData)
	})
	return file_omni_specs_siderolink_proto_rawDescData
}

var file_omni_specs_siderolink_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_omni_specs_siderolink_proto_goTypes = []any{
	(*SiderolinkConfigSpec)(nil),  // 0: specs.SiderolinkConfigSpec
	(*SiderolinkSpec)(nil),        // 1: specs.SiderolinkSpec
	(*SiderolinkCounterSpec)(nil), // 2: specs.SiderolinkCounterSpec
	(*ConnectionParamsSpec)(nil),  // 3: specs.ConnectionParamsSpec
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_omni_specs_siderolink_proto_depIdxs = []int32{
	4, // 0: specs.SiderolinkCounterSpec.last_alive:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_omni_specs_siderolink_proto_init() }
func file_omni_specs_siderolink_proto_init() {
	if File_omni_specs_siderolink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omni_specs_siderolink_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SiderolinkConfigSpec); i {
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
		file_omni_specs_siderolink_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SiderolinkSpec); i {
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
		file_omni_specs_siderolink_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SiderolinkCounterSpec); i {
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
		file_omni_specs_siderolink_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionParamsSpec); i {
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
			RawDescriptor: file_omni_specs_siderolink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omni_specs_siderolink_proto_goTypes,
		DependencyIndexes: file_omni_specs_siderolink_proto_depIdxs,
		MessageInfos:      file_omni_specs_siderolink_proto_msgTypes,
	}.Build()
	File_omni_specs_siderolink_proto = out.File
	file_omni_specs_siderolink_proto_rawDesc = nil
	file_omni_specs_siderolink_proto_goTypes = nil
	file_omni_specs_siderolink_proto_depIdxs = nil
}
