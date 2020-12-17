// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller/psmdb_cluster.proto

package controllerv1beta1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// PSMDBClusterState represents PSMDB cluster CR state.
type PSMDBClusterState int32

const (
	// PSMDB_CLUSTER_STATE_INVALID represents unknown state.
	PSMDBClusterState_PSMDB_CLUSTER_STATE_INVALID PSMDBClusterState = 0
	// PSMDB_CLUSTER_STATE_CHANGING represents a cluster being changed (initializing).
	PSMDBClusterState_PSMDB_CLUSTER_STATE_CHANGING PSMDBClusterState = 1
	// PSMDB_CLUSTER_STATE_READY represents a cluster without pending changes (ready).
	PSMDBClusterState_PSMDB_CLUSTER_STATE_READY PSMDBClusterState = 2
	// PSMDB_CLUSTER_STATE_FAILED represents a failed cluster (error).
	PSMDBClusterState_PSMDB_CLUSTER_STATE_FAILED PSMDBClusterState = 3
	// PSMDB_CLUSTER_STATE_DELETING represents a cluster being deleting.
	PSMDBClusterState_PSMDB_CLUSTER_STATE_DELETING PSMDBClusterState = 4
	// PSMDB_CLUSTER_STATE_PAUSED represents a paused cluster state.
	PSMDBClusterState_PSMDB_CLUSTER_STATE_PAUSED PSMDBClusterState = 5
)

// Enum value maps for PSMDBClusterState.
var (
	PSMDBClusterState_name = map[int32]string{
		0: "PSMDB_CLUSTER_STATE_INVALID",
		1: "PSMDB_CLUSTER_STATE_CHANGING",
		2: "PSMDB_CLUSTER_STATE_READY",
		3: "PSMDB_CLUSTER_STATE_FAILED",
		4: "PSMDB_CLUSTER_STATE_DELETING",
		5: "PSMDB_CLUSTER_STATE_PAUSED",
	}
	PSMDBClusterState_value = map[string]int32{
		"PSMDB_CLUSTER_STATE_INVALID":  0,
		"PSMDB_CLUSTER_STATE_CHANGING": 1,
		"PSMDB_CLUSTER_STATE_READY":    2,
		"PSMDB_CLUSTER_STATE_FAILED":   3,
		"PSMDB_CLUSTER_STATE_DELETING": 4,
		"PSMDB_CLUSTER_STATE_PAUSED":   5,
	}
)

func (x PSMDBClusterState) Enum() *PSMDBClusterState {
	p := new(PSMDBClusterState)
	*p = x
	return p
}

func (x PSMDBClusterState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PSMDBClusterState) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_psmdb_cluster_proto_enumTypes[0].Descriptor()
}

func (PSMDBClusterState) Type() protoreflect.EnumType {
	return &file_controller_psmdb_cluster_proto_enumTypes[0]
}

func (x PSMDBClusterState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PSMDBClusterState.Descriptor instead.
func (PSMDBClusterState) EnumDescriptor() ([]byte, []int) {
	return file_controller_psmdb_cluster_proto_rawDescGZIP(), []int{0}
}

// PSMDBBackupState represents PSMDB backup CR state.
type PSMDBBackupState int32

const (
	// PSMDB_BACKUP_STATE_INVALID represents unknown state.
	PSMDBBackupState_PSMDB_BACKUP_STATE_INVALID PSMDBBackupState = 0
	// PSMDB_BACKUP_STATE_RUNNING represents running backup (waiting, requested, running, FIXME check it).
	PSMDBBackupState_PSMDB_BACKUP_STATE_RUNNING PSMDBBackupState = 1
	// PSMDB_BACKUP_STATE_SUCCEEDED represents succeeded backup (ready, FIXME check it).
	PSMDBBackupState_PSMDB_BACKUP_STATE_SUCCEEDED PSMDBBackupState = 2
	// PSMDB_BACKUP_STATE_FAILED represents failed backup (rejected, error, FIXME check it).
	PSMDBBackupState_PSMDB_BACKUP_STATE_FAILED PSMDBBackupState = 3
)

// Enum value maps for PSMDBBackupState.
var (
	PSMDBBackupState_name = map[int32]string{
		0: "PSMDB_BACKUP_STATE_INVALID",
		1: "PSMDB_BACKUP_STATE_RUNNING",
		2: "PSMDB_BACKUP_STATE_SUCCEEDED",
		3: "PSMDB_BACKUP_STATE_FAILED",
	}
	PSMDBBackupState_value = map[string]int32{
		"PSMDB_BACKUP_STATE_INVALID":   0,
		"PSMDB_BACKUP_STATE_RUNNING":   1,
		"PSMDB_BACKUP_STATE_SUCCEEDED": 2,
		"PSMDB_BACKUP_STATE_FAILED":    3,
	}
)

func (x PSMDBBackupState) Enum() *PSMDBBackupState {
	p := new(PSMDBBackupState)
	*p = x
	return p
}

func (x PSMDBBackupState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PSMDBBackupState) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_psmdb_cluster_proto_enumTypes[1].Descriptor()
}

func (PSMDBBackupState) Type() protoreflect.EnumType {
	return &file_controller_psmdb_cluster_proto_enumTypes[1]
}

func (x PSMDBBackupState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PSMDBBackupState.Descriptor instead.
func (PSMDBBackupState) EnumDescriptor() ([]byte, []int) {
	return file_controller_psmdb_cluster_proto_rawDescGZIP(), []int{1}
}

// PSMDBClusterParams represents PSMDB cluster parameters that can be updated.
type PSMDBClusterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster size.
	ClusterSize int32 `protobuf:"varint,1,opt,name=cluster_size,json=clusterSize,proto3" json:"cluster_size,omitempty"`
	// Replicaset container parameters.
	Replicaset *PSMDBClusterParams_ReplicaSet `protobuf:"bytes,2,opt,name=replicaset,proto3" json:"replicaset,omitempty"`
	// Paused is true when cluster is suspended.
	Paused bool `protobuf:"varint,3,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *PSMDBClusterParams) Reset() {
	*x = PSMDBClusterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_psmdb_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSMDBClusterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSMDBClusterParams) ProtoMessage() {}

func (x *PSMDBClusterParams) ProtoReflect() protoreflect.Message {
	mi := &file_controller_psmdb_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSMDBClusterParams.ProtoReflect.Descriptor instead.
func (*PSMDBClusterParams) Descriptor() ([]byte, []int) {
	return file_controller_psmdb_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *PSMDBClusterParams) GetClusterSize() int32 {
	if x != nil {
		return x.ClusterSize
	}
	return 0
}

func (x *PSMDBClusterParams) GetReplicaset() *PSMDBClusterParams_ReplicaSet {
	if x != nil {
		return x.Replicaset
	}
	return nil
}

func (x *PSMDBClusterParams) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

// ReplicaSet container parameters.
type PSMDBClusterParams_ReplicaSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested compute resources.
	ComputeResources *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	// Disk size in bytes.
	DiskSize int64 `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *PSMDBClusterParams_ReplicaSet) Reset() {
	*x = PSMDBClusterParams_ReplicaSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_psmdb_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSMDBClusterParams_ReplicaSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSMDBClusterParams_ReplicaSet) ProtoMessage() {}

func (x *PSMDBClusterParams_ReplicaSet) ProtoReflect() protoreflect.Message {
	mi := &file_controller_psmdb_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSMDBClusterParams_ReplicaSet.ProtoReflect.Descriptor instead.
func (*PSMDBClusterParams_ReplicaSet) Descriptor() ([]byte, []int) {
	return file_controller_psmdb_cluster_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PSMDBClusterParams_ReplicaSet) GetComputeResources() *ComputeResources {
	if x != nil {
		return x.ComputeResources
	}
	return nil
}

func (x *PSMDBClusterParams_ReplicaSet) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

var File_controller_psmdb_cluster_proto protoreflect.FileDescriptor

var file_controller_psmdb_cluster_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x73, 0x6d,
	0x64, 0x62, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x29, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x17, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a,
	0x12, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10,
	0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x70,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x48, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x53, 0x4d, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x1a, 0x9b, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x12, 0x68, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0xd7, 0x01, 0x0a, 0x11, 0x50, 0x53, 0x4d, 0x44, 0x42,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a,
	0x1c, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x1e,
	0x0a, 0x1a, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x05,
	0x2a, 0x93, 0x01, 0x0a, 0x10, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x5f, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x53, 0x4d, 0x44, 0x42,
	0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_psmdb_cluster_proto_rawDescOnce sync.Once
	file_controller_psmdb_cluster_proto_rawDescData = file_controller_psmdb_cluster_proto_rawDesc
)

func file_controller_psmdb_cluster_proto_rawDescGZIP() []byte {
	file_controller_psmdb_cluster_proto_rawDescOnce.Do(func() {
		file_controller_psmdb_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_psmdb_cluster_proto_rawDescData)
	})
	return file_controller_psmdb_cluster_proto_rawDescData
}

var file_controller_psmdb_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_controller_psmdb_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_psmdb_cluster_proto_goTypes = []interface{}{
	(PSMDBClusterState)(0),                // 0: percona.platform.dbaas.controller.v1beta1.PSMDBClusterState
	(PSMDBBackupState)(0),                 // 1: percona.platform.dbaas.controller.v1beta1.PSMDBBackupState
	(*PSMDBClusterParams)(nil),            // 2: percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams
	(*PSMDBClusterParams_ReplicaSet)(nil), // 3: percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams.ReplicaSet
	(*ComputeResources)(nil),              // 4: percona.platform.dbaas.controller.v1beta1.ComputeResources
}
var file_controller_psmdb_cluster_proto_depIdxs = []int32{
	3, // 0: percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams.replicaset:type_name -> percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams.ReplicaSet
	4, // 1: percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams.ReplicaSet.compute_resources:type_name -> percona.platform.dbaas.controller.v1beta1.ComputeResources
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_psmdb_cluster_proto_init() }
func file_controller_psmdb_cluster_proto_init() {
	if File_controller_psmdb_cluster_proto != nil {
		return
	}
	file_controller_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_psmdb_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PSMDBClusterParams); i {
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
		file_controller_psmdb_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PSMDBClusterParams_ReplicaSet); i {
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
			RawDescriptor: file_controller_psmdb_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_psmdb_cluster_proto_goTypes,
		DependencyIndexes: file_controller_psmdb_cluster_proto_depIdxs,
		EnumInfos:         file_controller_psmdb_cluster_proto_enumTypes,
		MessageInfos:      file_controller_psmdb_cluster_proto_msgTypes,
	}.Build()
	File_controller_psmdb_cluster_proto = out.File
	file_controller_psmdb_cluster_proto_rawDesc = nil
	file_controller_psmdb_cluster_proto_goTypes = nil
	file_controller_psmdb_cluster_proto_depIdxs = nil
}
