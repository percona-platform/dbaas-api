// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller/psmdb_cluster.proto

package controllerv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
)

var PSMDBClusterState_name = map[int32]string{
	0: "PSMDB_CLUSTER_STATE_INVALID",
	1: "PSMDB_CLUSTER_STATE_CHANGING",
	2: "PSMDB_CLUSTER_STATE_READY",
	3: "PSMDB_CLUSTER_STATE_FAILED",
	4: "PSMDB_CLUSTER_STATE_DELETING",
}

var PSMDBClusterState_value = map[string]int32{
	"PSMDB_CLUSTER_STATE_INVALID":  0,
	"PSMDB_CLUSTER_STATE_CHANGING": 1,
	"PSMDB_CLUSTER_STATE_READY":    2,
	"PSMDB_CLUSTER_STATE_FAILED":   3,
	"PSMDB_CLUSTER_STATE_DELETING": 4,
}

func (x PSMDBClusterState) String() string {
	return proto.EnumName(PSMDBClusterState_name, int32(x))
}

func (PSMDBClusterState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04d59cae079be5, []int{0}
}

// PSMDBBackupState represents PSMDB backup CR state.
type PSMDBBackupState int32

const (
	// PSMDB_BACKUP_STATE_INVALID represents unknown state.
	PSMDBBackupState_PSMDB_BACKUP_STATE_INVALID PSMDBBackupState = 0
	// PSMDB_BACKUP_STATE_RUNNING represents running backup (Starting, Running).
	PSMDBBackupState_PSMDB_BACKUP_STATE_RUNNING PSMDBBackupState = 1
	// PSMDB_BACKUP_STATE_SUCCEEDED represents succeeded backup (Succeeded).
	PSMDBBackupState_PSMDB_BACKUP_STATE_SUCCEEDED PSMDBBackupState = 2
	// PSMDB_BACKUP_STATE_FAILED represents failed backup (Failed).
	PSMDBBackupState_PSMDB_BACKUP_STATE_FAILED PSMDBBackupState = 3
)

var PSMDBBackupState_name = map[int32]string{
	0: "PSMDB_BACKUP_STATE_INVALID",
	1: "PSMDB_BACKUP_STATE_RUNNING",
	2: "PSMDB_BACKUP_STATE_SUCCEEDED",
	3: "PSMDB_BACKUP_STATE_FAILED",
}

var PSMDBBackupState_value = map[string]int32{
	"PSMDB_BACKUP_STATE_INVALID":   0,
	"PSMDB_BACKUP_STATE_RUNNING":   1,
	"PSMDB_BACKUP_STATE_SUCCEEDED": 2,
	"PSMDB_BACKUP_STATE_FAILED":    3,
}

func (x PSMDBBackupState) String() string {
	return proto.EnumName(PSMDBBackupState_name, int32(x))
}

func (PSMDBBackupState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04d59cae079be5, []int{1}
}

// PSMDBClusterParams represents PSMDB cluster parameters that can be updated.
type PSMDBClusterParams struct {
	// Cluster size.
	ClusterSize int32 `protobuf:"varint,1,opt,name=cluster_size,json=clusterSize,proto3" json:"cluster_size,omitempty"`
	// Replicaset container parameters.
	Replicaset           *PSMDBClusterParams_ReplicaSet `protobuf:"bytes,2,opt,name=replicaset,proto3" json:"replicaset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PSMDBClusterParams) Reset()         { *m = PSMDBClusterParams{} }
func (m *PSMDBClusterParams) String() string { return proto.CompactTextString(m) }
func (*PSMDBClusterParams) ProtoMessage()    {}
func (*PSMDBClusterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04d59cae079be5, []int{0}
}

func (m *PSMDBClusterParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PSMDBClusterParams.Unmarshal(m, b)
}
func (m *PSMDBClusterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PSMDBClusterParams.Marshal(b, m, deterministic)
}
func (m *PSMDBClusterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PSMDBClusterParams.Merge(m, src)
}
func (m *PSMDBClusterParams) XXX_Size() int {
	return xxx_messageInfo_PSMDBClusterParams.Size(m)
}
func (m *PSMDBClusterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PSMDBClusterParams.DiscardUnknown(m)
}

var xxx_messageInfo_PSMDBClusterParams proto.InternalMessageInfo

func (m *PSMDBClusterParams) GetClusterSize() int32 {
	if m != nil {
		return m.ClusterSize
	}
	return 0
}

func (m *PSMDBClusterParams) GetReplicaset() *PSMDBClusterParams_ReplicaSet {
	if m != nil {
		return m.Replicaset
	}
	return nil
}

// ReplicaSet container parameters.
type PSMDBClusterParams_ReplicaSet struct {
	// Requested compute resources.
	ComputeResources     *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PSMDBClusterParams_ReplicaSet) Reset()         { *m = PSMDBClusterParams_ReplicaSet{} }
func (m *PSMDBClusterParams_ReplicaSet) String() string { return proto.CompactTextString(m) }
func (*PSMDBClusterParams_ReplicaSet) ProtoMessage()    {}
func (*PSMDBClusterParams_ReplicaSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04d59cae079be5, []int{0, 0}
}

func (m *PSMDBClusterParams_ReplicaSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PSMDBClusterParams_ReplicaSet.Unmarshal(m, b)
}
func (m *PSMDBClusterParams_ReplicaSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PSMDBClusterParams_ReplicaSet.Marshal(b, m, deterministic)
}
func (m *PSMDBClusterParams_ReplicaSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PSMDBClusterParams_ReplicaSet.Merge(m, src)
}
func (m *PSMDBClusterParams_ReplicaSet) XXX_Size() int {
	return xxx_messageInfo_PSMDBClusterParams_ReplicaSet.Size(m)
}
func (m *PSMDBClusterParams_ReplicaSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PSMDBClusterParams_ReplicaSet.DiscardUnknown(m)
}

var xxx_messageInfo_PSMDBClusterParams_ReplicaSet proto.InternalMessageInfo

func (m *PSMDBClusterParams_ReplicaSet) GetComputeResources() *ComputeResources {
	if m != nil {
		return m.ComputeResources
	}
	return nil
}

func init() {
	proto.RegisterEnum("percona.platform.dbaas.controller.v1beta1.PSMDBClusterState", PSMDBClusterState_name, PSMDBClusterState_value)
	proto.RegisterEnum("percona.platform.dbaas.controller.v1beta1.PSMDBBackupState", PSMDBBackupState_name, PSMDBBackupState_value)
	proto.RegisterType((*PSMDBClusterParams)(nil), "percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams")
	proto.RegisterType((*PSMDBClusterParams_ReplicaSet)(nil), "percona.platform.dbaas.controller.v1beta1.PSMDBClusterParams.ReplicaSet")
}

func init() {
	proto.RegisterFile("controller/psmdb_cluster.proto", fileDescriptor_ab04d59cae079be5)
}

var fileDescriptor_ab04d59cae079be5 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x97, 0x00, 0x3b, 0xb8, 0x1c, 0x3c, 0x5f, 0x18, 0x65, 0x74, 0x15, 0xa7, 0x6d, 0xd2,
	0x5c, 0x6d, 0x48, 0x5c, 0x76, 0xca, 0x1f, 0xb3, 0x55, 0x94, 0xa8, 0x4a, 0x5a, 0x24, 0xb8, 0x44,
	0x8e, 0x6b, 0x46, 0xb4, 0xb8, 0xb6, 0x6c, 0x67, 0x93, 0x76, 0xe0, 0x4b, 0xf0, 0x49, 0x38, 0xf1,
	0x71, 0x90, 0xf8, 0x24, 0x88, 0xb8, 0x6b, 0x3a, 0x5a, 0x24, 0x76, 0xb3, 0xf2, 0x3c, 0xef, 0xf3,
	0xfc, 0x5e, 0xbd, 0x01, 0x3d, 0x26, 0xe7, 0x56, 0xcb, 0xaa, 0xe2, 0x7a, 0xa0, 0x8c, 0x98, 0x15,
	0x39, 0xab, 0x6a, 0x63, 0xb9, 0xc6, 0x4a, 0x4b, 0x2b, 0xd1, 0xa1, 0xe2, 0x9a, 0xc9, 0x39, 0xc5,
	0xaa, 0xa2, 0xf6, 0xb3, 0xd4, 0x02, 0xcf, 0x0a, 0x4a, 0x0d, 0x6e, 0xc7, 0xf0, 0xf5, 0x49, 0xc1,
	0x2d, 0x3d, 0xe9, 0x3e, 0x5b, 0x89, 0x62, 0x52, 0x08, 0x39, 0x77, 0x19, 0xdd, 0x37, 0x97, 0xa5,
	0xfd, 0x52, 0x17, 0x98, 0x49, 0x31, 0x10, 0x37, 0xa5, 0xbd, 0x92, 0x37, 0x83, 0x4b, 0x79, 0xdc,
	0x88, 0xc7, 0xd7, 0xb4, 0x2a, 0x67, 0xd4, 0x4a, 0x6d, 0x06, 0xcb, 0xa7, 0x9b, 0x7b, 0xf5, 0xdd,
	0x07, 0x68, 0x9c, 0xbd, 0x8f, 0xc3, 0xc8, 0x21, 0x8d, 0xa9, 0xa6, 0xc2, 0xa0, 0x43, 0xf0, 0x74,
	0xc1, 0x98, 0x9b, 0xf2, 0x96, 0xef, 0x7a, 0x7d, 0xef, 0xe0, 0x49, 0xb8, 0xfd, 0xeb, 0xe7, 0xbe,
	0x0f, 0xb7, 0xd2, 0xce, 0x42, 0xcb, 0xca, 0x5b, 0x8e, 0x14, 0x00, 0x9a, 0xab, 0xaa, 0x64, 0xd4,
	0x70, 0xbb, 0xeb, 0xf7, 0xbd, 0x83, 0xce, 0xe9, 0x05, 0xfe, 0xef, 0x95, 0xf0, 0x7a, 0x3b, 0x4e,
	0x5d, 0x5e, 0xc6, 0xad, 0xab, 0xec, 0x7b, 0xe9, 0x4a, 0x47, 0xf7, 0x2b, 0x00, 0xad, 0x03, 0x29,
	0xb0, 0xc3, 0xa4, 0x50, 0xb5, 0xe5, 0xb9, 0xe6, 0x46, 0xd6, 0x9a, 0x71, 0xd3, 0xf0, 0x76, 0x4e,
	0xcf, 0x1e, 0x80, 0x11, 0xb9, 0x8c, 0xf4, 0x2e, 0x62, 0xd9, 0x0c, 0xd9, 0x5f, 0xca, 0xd1, 0x0f,
	0x0f, 0xec, 0xac, 0x52, 0x67, 0x96, 0x5a, 0x8e, 0xf6, 0xc1, 0x8b, 0xe6, 0x63, 0x1e, 0x8d, 0xa6,
	0xd9, 0x84, 0xa4, 0x79, 0x36, 0x09, 0x26, 0x24, 0x1f, 0x26, 0x1f, 0x82, 0xd1, 0x30, 0x86, 0x5b,
	0xa8, 0x0f, 0xf6, 0x36, 0x19, 0xa2, 0x8b, 0x20, 0x39, 0x1f, 0x26, 0xe7, 0xd0, 0x43, 0x2f, 0xc1,
	0xf3, 0x4d, 0x8e, 0x94, 0x04, 0xf1, 0x47, 0xe8, 0xa3, 0x1e, 0xe8, 0x6e, 0x92, 0xdf, 0x06, 0xc3,
	0x11, 0x89, 0xe1, 0xa3, 0x7f, 0x15, 0xc4, 0x64, 0x44, 0x26, 0x7f, 0x0a, 0x1e, 0x1f, 0x7d, 0xf3,
	0x00, 0x6c, 0x2c, 0x21, 0x65, 0x57, 0xb5, 0x72, 0xe0, 0xcb, 0xd8, 0x30, 0x88, 0xde, 0x4d, 0xc7,
	0x6b, 0xdc, 0x9b, 0xf5, 0x74, 0x9a, 0x24, 0x8e, 0x7a, 0x59, 0x7b, 0x4f, 0xcf, 0xa6, 0x51, 0x44,
	0x48, 0x4c, 0x62, 0xe8, 0xb7, 0x7b, 0xdd, 0x73, 0xdc, 0x71, 0x87, 0xbd, 0x4f, 0x7b, 0xed, 0x41,
	0xce, 0xda, 0xe7, 0xe2, 0x34, 0xc5, 0x76, 0xf3, 0xab, 0xbe, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0xf3, 0x9c, 0xe1, 0x48, 0x03, 0x00, 0x00,
}
