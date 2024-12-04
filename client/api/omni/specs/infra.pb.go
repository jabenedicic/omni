// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: omni/specs/infra.proto

package specs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MachineRequestStatusSpec_Stage int32

const (
	MachineRequestStatusSpec_UNKNOWN      MachineRequestStatusSpec_Stage = 0
	MachineRequestStatusSpec_PROVISIONING MachineRequestStatusSpec_Stage = 1
	MachineRequestStatusSpec_PROVISIONED  MachineRequestStatusSpec_Stage = 2
	MachineRequestStatusSpec_FAILED       MachineRequestStatusSpec_Stage = 3
)

// Enum value maps for MachineRequestStatusSpec_Stage.
var (
	MachineRequestStatusSpec_Stage_name = map[int32]string{
		0: "UNKNOWN",
		1: "PROVISIONING",
		2: "PROVISIONED",
		3: "FAILED",
	}
	MachineRequestStatusSpec_Stage_value = map[string]int32{
		"UNKNOWN":      0,
		"PROVISIONING": 1,
		"PROVISIONED":  2,
		"FAILED":       3,
	}
)

func (x MachineRequestStatusSpec_Stage) Enum() *MachineRequestStatusSpec_Stage {
	p := new(MachineRequestStatusSpec_Stage)
	*p = x
	return p
}

func (x MachineRequestStatusSpec_Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineRequestStatusSpec_Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_omni_specs_infra_proto_enumTypes[0].Descriptor()
}

func (MachineRequestStatusSpec_Stage) Type() protoreflect.EnumType {
	return &file_omni_specs_infra_proto_enumTypes[0]
}

func (x MachineRequestStatusSpec_Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineRequestStatusSpec_Stage.Descriptor instead.
func (MachineRequestStatusSpec_Stage) EnumDescriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{1, 0}
}

type InfraMachineSpec_MachinePowerState int32

const (
	InfraMachineSpec_POWER_STATE_OFF InfraMachineSpec_MachinePowerState = 0
	InfraMachineSpec_POWER_STATE_ON  InfraMachineSpec_MachinePowerState = 1
)

// Enum value maps for InfraMachineSpec_MachinePowerState.
var (
	InfraMachineSpec_MachinePowerState_name = map[int32]string{
		0: "POWER_STATE_OFF",
		1: "POWER_STATE_ON",
	}
	InfraMachineSpec_MachinePowerState_value = map[string]int32{
		"POWER_STATE_OFF": 0,
		"POWER_STATE_ON":  1,
	}
)

func (x InfraMachineSpec_MachinePowerState) Enum() *InfraMachineSpec_MachinePowerState {
	p := new(InfraMachineSpec_MachinePowerState)
	*p = x
	return p
}

func (x InfraMachineSpec_MachinePowerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InfraMachineSpec_MachinePowerState) Descriptor() protoreflect.EnumDescriptor {
	return file_omni_specs_infra_proto_enumTypes[1].Descriptor()
}

func (InfraMachineSpec_MachinePowerState) Type() protoreflect.EnumType {
	return &file_omni_specs_infra_proto_enumTypes[1]
}

func (x InfraMachineSpec_MachinePowerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InfraMachineSpec_MachinePowerState.Descriptor instead.
func (InfraMachineSpec_MachinePowerState) EnumDescriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{2, 0}
}

type InfraMachineStatusSpec_MachinePowerState int32

const (
	InfraMachineStatusSpec_POWER_STATE_UNKNOWN InfraMachineStatusSpec_MachinePowerState = 0
	InfraMachineStatusSpec_POWER_STATE_OFF     InfraMachineStatusSpec_MachinePowerState = 1
	InfraMachineStatusSpec_POWER_STATE_ON      InfraMachineStatusSpec_MachinePowerState = 2
)

// Enum value maps for InfraMachineStatusSpec_MachinePowerState.
var (
	InfraMachineStatusSpec_MachinePowerState_name = map[int32]string{
		0: "POWER_STATE_UNKNOWN",
		1: "POWER_STATE_OFF",
		2: "POWER_STATE_ON",
	}
	InfraMachineStatusSpec_MachinePowerState_value = map[string]int32{
		"POWER_STATE_UNKNOWN": 0,
		"POWER_STATE_OFF":     1,
		"POWER_STATE_ON":      2,
	}
)

func (x InfraMachineStatusSpec_MachinePowerState) Enum() *InfraMachineStatusSpec_MachinePowerState {
	p := new(InfraMachineStatusSpec_MachinePowerState)
	*p = x
	return p
}

func (x InfraMachineStatusSpec_MachinePowerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InfraMachineStatusSpec_MachinePowerState) Descriptor() protoreflect.EnumDescriptor {
	return file_omni_specs_infra_proto_enumTypes[2].Descriptor()
}

func (InfraMachineStatusSpec_MachinePowerState) Type() protoreflect.EnumType {
	return &file_omni_specs_infra_proto_enumTypes[2]
}

func (x InfraMachineStatusSpec_MachinePowerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InfraMachineStatusSpec_MachinePowerState.Descriptor instead.
func (InfraMachineStatusSpec_MachinePowerState) EnumDescriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{4, 0}
}

type MachineRequestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TalosVersion string         `protobuf:"bytes,1,opt,name=talos_version,json=talosVersion,proto3" json:"talos_version,omitempty"`
	Overlay      *Overlay       `protobuf:"bytes,2,opt,name=overlay,proto3" json:"overlay,omitempty"`
	Extensions   []string       `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty"`
	KernelArgs   []string       `protobuf:"bytes,4,rep,name=kernel_args,json=kernelArgs,proto3" json:"kernel_args,omitempty"`
	MetaValues   []*MetaValue   `protobuf:"bytes,5,rep,name=meta_values,json=metaValues,proto3" json:"meta_values,omitempty"`
	ProviderData string         `protobuf:"bytes,6,opt,name=provider_data,json=providerData,proto3" json:"provider_data,omitempty"`
	GrpcTunnel   GrpcTunnelMode `protobuf:"varint,7,opt,name=grpc_tunnel,json=grpcTunnel,proto3,enum=specs.GrpcTunnelMode" json:"grpc_tunnel,omitempty"`
}

func (x *MachineRequestSpec) Reset() {
	*x = MachineRequestSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineRequestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineRequestSpec) ProtoMessage() {}

func (x *MachineRequestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineRequestSpec.ProtoReflect.Descriptor instead.
func (*MachineRequestSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{0}
}

func (x *MachineRequestSpec) GetTalosVersion() string {
	if x != nil {
		return x.TalosVersion
	}
	return ""
}

func (x *MachineRequestSpec) GetOverlay() *Overlay {
	if x != nil {
		return x.Overlay
	}
	return nil
}

func (x *MachineRequestSpec) GetExtensions() []string {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *MachineRequestSpec) GetKernelArgs() []string {
	if x != nil {
		return x.KernelArgs
	}
	return nil
}

func (x *MachineRequestSpec) GetMetaValues() []*MetaValue {
	if x != nil {
		return x.MetaValues
	}
	return nil
}

func (x *MachineRequestSpec) GetProviderData() string {
	if x != nil {
		return x.ProviderData
	}
	return ""
}

func (x *MachineRequestSpec) GetGrpcTunnel() GrpcTunnelMode {
	if x != nil {
		return x.GrpcTunnel
	}
	return GrpcTunnelMode_UNSET
}

type MachineRequestStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stage  MachineRequestStatusSpec_Stage `protobuf:"varint,2,opt,name=stage,proto3,enum=specs.MachineRequestStatusSpec_Stage" json:"stage,omitempty"`
	Error  string                         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Status string                         `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MachineRequestStatusSpec) Reset() {
	*x = MachineRequestStatusSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineRequestStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineRequestStatusSpec) ProtoMessage() {}

func (x *MachineRequestStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineRequestStatusSpec.ProtoReflect.Descriptor instead.
func (*MachineRequestStatusSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{1}
}

func (x *MachineRequestStatusSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MachineRequestStatusSpec) GetStage() MachineRequestStatusSpec_Stage {
	if x != nil {
		return x.Stage
	}
	return MachineRequestStatusSpec_UNKNOWN
}

func (x *MachineRequestStatusSpec) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MachineRequestStatusSpec) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InfraMachineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferredPowerState InfraMachineSpec_MachinePowerState `protobuf:"varint,1,opt,name=preferred_power_state,json=preferredPowerState,proto3,enum=specs.InfraMachineSpec_MachinePowerState" json:"preferred_power_state,omitempty"`
	Accepted            bool                               `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	ClusterTalosVersion string                             `protobuf:"bytes,3,opt,name=cluster_talos_version,json=clusterTalosVersion,proto3" json:"cluster_talos_version,omitempty"`
	Extensions          []string                           `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty"`
	// WipeId is set to a new id each time the machine gets unallocated.
	//
	// It is used by the provider to ensure that the machine is wiped between allocations.
	WipeId          string `protobuf:"bytes,5,opt,name=wipe_id,json=wipeId,proto3" json:"wipe_id,omitempty"`
	ExtraKernelArgs string `protobuf:"bytes,6,opt,name=extra_kernel_args,json=extraKernelArgs,proto3" json:"extra_kernel_args,omitempty"`
}

func (x *InfraMachineSpec) Reset() {
	*x = InfraMachineSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfraMachineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraMachineSpec) ProtoMessage() {}

func (x *InfraMachineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraMachineSpec.ProtoReflect.Descriptor instead.
func (*InfraMachineSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{2}
}

func (x *InfraMachineSpec) GetPreferredPowerState() InfraMachineSpec_MachinePowerState {
	if x != nil {
		return x.PreferredPowerState
	}
	return InfraMachineSpec_POWER_STATE_OFF
}

func (x *InfraMachineSpec) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *InfraMachineSpec) GetClusterTalosVersion() string {
	if x != nil {
		return x.ClusterTalosVersion
	}
	return ""
}

func (x *InfraMachineSpec) GetExtensions() []string {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *InfraMachineSpec) GetWipeId() string {
	if x != nil {
		return x.WipeId
	}
	return ""
}

func (x *InfraMachineSpec) GetExtraKernelArgs() string {
	if x != nil {
		return x.ExtraKernelArgs
	}
	return ""
}

type InfraMachineStateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installed bool `protobuf:"varint,1,opt,name=installed,proto3" json:"installed,omitempty"`
}

func (x *InfraMachineStateSpec) Reset() {
	*x = InfraMachineStateSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfraMachineStateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraMachineStateSpec) ProtoMessage() {}

func (x *InfraMachineStateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraMachineStateSpec.ProtoReflect.Descriptor instead.
func (*InfraMachineStateSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{3}
}

func (x *InfraMachineStateSpec) GetInstalled() bool {
	if x != nil {
		return x.Installed
	}
	return false
}

type InfraMachineStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PowerState InfraMachineStatusSpec_MachinePowerState `protobuf:"varint,1,opt,name=power_state,json=powerState,proto3,enum=specs.InfraMachineStatusSpec_MachinePowerState" json:"power_state,omitempty"`
	ReadyToUse bool                                     `protobuf:"varint,2,opt,name=ready_to_use,json=readyToUse,proto3" json:"ready_to_use,omitempty"`
}

func (x *InfraMachineStatusSpec) Reset() {
	*x = InfraMachineStatusSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfraMachineStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraMachineStatusSpec) ProtoMessage() {}

func (x *InfraMachineStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraMachineStatusSpec.ProtoReflect.Descriptor instead.
func (*InfraMachineStatusSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{4}
}

func (x *InfraMachineStatusSpec) GetPowerState() InfraMachineStatusSpec_MachinePowerState {
	if x != nil {
		return x.PowerState
	}
	return InfraMachineStatusSpec_POWER_STATE_UNKNOWN
}

func (x *InfraMachineStatusSpec) GetReadyToUse() bool {
	if x != nil {
		return x.ReadyToUse
	}
	return false
}

type InfraProviderStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema      string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Icon        string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *InfraProviderStatusSpec) Reset() {
	*x = InfraProviderStatusSpec{}
	mi := &file_omni_specs_infra_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfraProviderStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraProviderStatusSpec) ProtoMessage() {}

func (x *InfraProviderStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_omni_specs_infra_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraProviderStatusSpec.ProtoReflect.Descriptor instead.
func (*InfraProviderStatusSpec) Descriptor() ([]byte, []int) {
	return file_omni_specs_infra_proto_rawDescGZIP(), []int{5}
}

func (x *InfraProviderStatusSpec) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InfraProviderStatusSpec) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

var File_omni_specs_infra_proto protoreflect.FileDescriptor

var file_omni_specs_infra_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a,
	0x15, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x31, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xda, 0x01,
	0x0a, 0x18, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0xe4, 0x02, 0x0a, 0x10, 0x49,
	0x6e, 0x66, 0x72, 0x61, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x5d, 0x0a, 0x15, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x13, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x61, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x77, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x77, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x5f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41,
	0x72, 0x67, 0x73, 0x22, 0x3c, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x57, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x10,
	0x01, 0x22, 0x35, 0x0a, 0x15, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x50, 0x0a, 0x0b, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4f,
	0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x7b,
	0x0a, 0x17, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omni_specs_infra_proto_rawDescOnce sync.Once
	file_omni_specs_infra_proto_rawDescData = file_omni_specs_infra_proto_rawDesc
)

func file_omni_specs_infra_proto_rawDescGZIP() []byte {
	file_omni_specs_infra_proto_rawDescOnce.Do(func() {
		file_omni_specs_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_omni_specs_infra_proto_rawDescData)
	})
	return file_omni_specs_infra_proto_rawDescData
}

var file_omni_specs_infra_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_omni_specs_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omni_specs_infra_proto_goTypes = []any{
	(MachineRequestStatusSpec_Stage)(0),           // 0: specs.MachineRequestStatusSpec.Stage
	(InfraMachineSpec_MachinePowerState)(0),       // 1: specs.InfraMachineSpec.MachinePowerState
	(InfraMachineStatusSpec_MachinePowerState)(0), // 2: specs.InfraMachineStatusSpec.MachinePowerState
	(*MachineRequestSpec)(nil),                    // 3: specs.MachineRequestSpec
	(*MachineRequestStatusSpec)(nil),              // 4: specs.MachineRequestStatusSpec
	(*InfraMachineSpec)(nil),                      // 5: specs.InfraMachineSpec
	(*InfraMachineStateSpec)(nil),                 // 6: specs.InfraMachineStateSpec
	(*InfraMachineStatusSpec)(nil),                // 7: specs.InfraMachineStatusSpec
	(*InfraProviderStatusSpec)(nil),               // 8: specs.InfraProviderStatusSpec
	(*Overlay)(nil),                               // 9: specs.Overlay
	(*MetaValue)(nil),                             // 10: specs.MetaValue
	(GrpcTunnelMode)(0),                           // 11: specs.GrpcTunnelMode
}
var file_omni_specs_infra_proto_depIdxs = []int32{
	9,  // 0: specs.MachineRequestSpec.overlay:type_name -> specs.Overlay
	10, // 1: specs.MachineRequestSpec.meta_values:type_name -> specs.MetaValue
	11, // 2: specs.MachineRequestSpec.grpc_tunnel:type_name -> specs.GrpcTunnelMode
	0,  // 3: specs.MachineRequestStatusSpec.stage:type_name -> specs.MachineRequestStatusSpec.Stage
	1,  // 4: specs.InfraMachineSpec.preferred_power_state:type_name -> specs.InfraMachineSpec.MachinePowerState
	2,  // 5: specs.InfraMachineStatusSpec.power_state:type_name -> specs.InfraMachineStatusSpec.MachinePowerState
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_omni_specs_infra_proto_init() }
func file_omni_specs_infra_proto_init() {
	if File_omni_specs_infra_proto != nil {
		return
	}
	file_omni_specs_omni_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omni_specs_infra_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omni_specs_infra_proto_goTypes,
		DependencyIndexes: file_omni_specs_infra_proto_depIdxs,
		EnumInfos:         file_omni_specs_infra_proto_enumTypes,
		MessageInfos:      file_omni_specs_infra_proto_msgTypes,
	}.Build()
	File_omni_specs_infra_proto = out.File
	file_omni_specs_infra_proto_rawDesc = nil
	file_omni_specs_infra_proto_goTypes = nil
	file_omni_specs_infra_proto_depIdxs = nil
}
