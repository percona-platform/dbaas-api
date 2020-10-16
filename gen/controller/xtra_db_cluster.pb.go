// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller/xtra_db_cluster.proto

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

// XtraDBClusterState represents XtraDB cluster CR state.
type XtraDBClusterState int32

const (
	// XTRA_DB_CLUSTER_STATE_INVALID represents unknown state.
	XtraDBClusterState_XTRA_DB_CLUSTER_STATE_INVALID XtraDBClusterState = 0
	// XTRA_DB_CLUSTER_STATE_CHANGING represents a cluster being changed (initializing).
	XtraDBClusterState_XTRA_DB_CLUSTER_STATE_CHANGING XtraDBClusterState = 1
	// XTRA_DB_CLUSTER_STATE_READY represents a cluster without pending changes (ready).
	XtraDBClusterState_XTRA_DB_CLUSTER_STATE_READY XtraDBClusterState = 2
	// XTRA_DB_CLUSTER_STATE_FAILED represents a failed cluster (error).
	XtraDBClusterState_XTRA_DB_CLUSTER_STATE_FAILED XtraDBClusterState = 3
	// XTRA_DB_CLUSTER_STATE_DELETING represents a cluster being deleting.
	XtraDBClusterState_XTRA_DB_CLUSTER_STATE_DELETING XtraDBClusterState = 4
)

var XtraDBClusterState_name = map[int32]string{
	0: "XTRA_DB_CLUSTER_STATE_INVALID",
	1: "XTRA_DB_CLUSTER_STATE_CHANGING",
	2: "XTRA_DB_CLUSTER_STATE_READY",
	3: "XTRA_DB_CLUSTER_STATE_FAILED",
	4: "XTRA_DB_CLUSTER_STATE_DELETING",
}

var XtraDBClusterState_value = map[string]int32{
	"XTRA_DB_CLUSTER_STATE_INVALID":  0,
	"XTRA_DB_CLUSTER_STATE_CHANGING": 1,
	"XTRA_DB_CLUSTER_STATE_READY":    2,
	"XTRA_DB_CLUSTER_STATE_FAILED":   3,
	"XTRA_DB_CLUSTER_STATE_DELETING": 4,
}

func (x XtraDBClusterState) String() string {
	return proto.EnumName(XtraDBClusterState_name, int32(x))
}

func (XtraDBClusterState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1effda7bec99251f, []int{0}
}

// XtraDBBackupState represents XtraDB backup CR state.
type XtraDBBackupState int32

const (
	// XTRA_DB_BACKUP_STATE_INVALID represents unknown state.
	XtraDBBackupState_XTRA_DB_BACKUP_STATE_INVALID XtraDBBackupState = 0
	// XTRA_DB_BACKUP_STATE_RUNNING represents running backup (Starting, Running).
	XtraDBBackupState_XTRA_DB_BACKUP_STATE_RUNNING XtraDBBackupState = 1
	// XTRA_DB_BACKUP_STATE_SUCCEEDED represents succeeded backup (Succeeded).
	XtraDBBackupState_XTRA_DB_BACKUP_STATE_SUCCEEDED XtraDBBackupState = 2
	// XTRA_DB_BACKUP_STATE_FAILED represents failed backup (Failed).
	XtraDBBackupState_XTRA_DB_BACKUP_STATE_FAILED XtraDBBackupState = 3
)

var XtraDBBackupState_name = map[int32]string{
	0: "XTRA_DB_BACKUP_STATE_INVALID",
	1: "XTRA_DB_BACKUP_STATE_RUNNING",
	2: "XTRA_DB_BACKUP_STATE_SUCCEEDED",
	3: "XTRA_DB_BACKUP_STATE_FAILED",
}

var XtraDBBackupState_value = map[string]int32{
	"XTRA_DB_BACKUP_STATE_INVALID":   0,
	"XTRA_DB_BACKUP_STATE_RUNNING":   1,
	"XTRA_DB_BACKUP_STATE_SUCCEEDED": 2,
	"XTRA_DB_BACKUP_STATE_FAILED":    3,
}

func (x XtraDBBackupState) String() string {
	return proto.EnumName(XtraDBBackupState_name, int32(x))
}

func (XtraDBBackupState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1effda7bec99251f, []int{1}
}

// XtraDBClusterParams represents XtraDB cluster parameters that can be updated.
type XtraDBClusterParams struct {
	// Cluster size.
	ClusterSize int32 `protobuf:"varint,1,opt,name=cluster_size,json=clusterSize,proto3" json:"cluster_size,omitempty"`
	// PXC container parameters.
	Pxc *XtraDBClusterParams_PXC `protobuf:"bytes,2,opt,name=pxc,proto3" json:"pxc,omitempty"`
	// ProxySQL container parameters.
	Proxysql             *XtraDBClusterParams_ProxySQL `protobuf:"bytes,3,opt,name=proxysql,proto3" json:"proxysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *XtraDBClusterParams) Reset()         { *m = XtraDBClusterParams{} }
func (m *XtraDBClusterParams) String() string { return proto.CompactTextString(m) }
func (*XtraDBClusterParams) ProtoMessage()    {}
func (*XtraDBClusterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1effda7bec99251f, []int{0}
}

func (m *XtraDBClusterParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtraDBClusterParams.Unmarshal(m, b)
}
func (m *XtraDBClusterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtraDBClusterParams.Marshal(b, m, deterministic)
}
func (m *XtraDBClusterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtraDBClusterParams.Merge(m, src)
}
func (m *XtraDBClusterParams) XXX_Size() int {
	return xxx_messageInfo_XtraDBClusterParams.Size(m)
}
func (m *XtraDBClusterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_XtraDBClusterParams.DiscardUnknown(m)
}

var xxx_messageInfo_XtraDBClusterParams proto.InternalMessageInfo

func (m *XtraDBClusterParams) GetClusterSize() int32 {
	if m != nil {
		return m.ClusterSize
	}
	return 0
}

func (m *XtraDBClusterParams) GetPxc() *XtraDBClusterParams_PXC {
	if m != nil {
		return m.Pxc
	}
	return nil
}

func (m *XtraDBClusterParams) GetProxysql() *XtraDBClusterParams_ProxySQL {
	if m != nil {
		return m.Proxysql
	}
	return nil
}

// PXC container parameters.
type XtraDBClusterParams_PXC struct {
	// Requested compute resources.
	ComputeResources     *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XtraDBClusterParams_PXC) Reset()         { *m = XtraDBClusterParams_PXC{} }
func (m *XtraDBClusterParams_PXC) String() string { return proto.CompactTextString(m) }
func (*XtraDBClusterParams_PXC) ProtoMessage()    {}
func (*XtraDBClusterParams_PXC) Descriptor() ([]byte, []int) {
	return fileDescriptor_1effda7bec99251f, []int{0, 0}
}

func (m *XtraDBClusterParams_PXC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtraDBClusterParams_PXC.Unmarshal(m, b)
}
func (m *XtraDBClusterParams_PXC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtraDBClusterParams_PXC.Marshal(b, m, deterministic)
}
func (m *XtraDBClusterParams_PXC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtraDBClusterParams_PXC.Merge(m, src)
}
func (m *XtraDBClusterParams_PXC) XXX_Size() int {
	return xxx_messageInfo_XtraDBClusterParams_PXC.Size(m)
}
func (m *XtraDBClusterParams_PXC) XXX_DiscardUnknown() {
	xxx_messageInfo_XtraDBClusterParams_PXC.DiscardUnknown(m)
}

var xxx_messageInfo_XtraDBClusterParams_PXC proto.InternalMessageInfo

func (m *XtraDBClusterParams_PXC) GetComputeResources() *ComputeResources {
	if m != nil {
		return m.ComputeResources
	}
	return nil
}

// ProxySQL container parameters.
type XtraDBClusterParams_ProxySQL struct {
	// Requested compute resources.
	ComputeResources     *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XtraDBClusterParams_ProxySQL) Reset()         { *m = XtraDBClusterParams_ProxySQL{} }
func (m *XtraDBClusterParams_ProxySQL) String() string { return proto.CompactTextString(m) }
func (*XtraDBClusterParams_ProxySQL) ProtoMessage()    {}
func (*XtraDBClusterParams_ProxySQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_1effda7bec99251f, []int{0, 1}
}

func (m *XtraDBClusterParams_ProxySQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtraDBClusterParams_ProxySQL.Unmarshal(m, b)
}
func (m *XtraDBClusterParams_ProxySQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtraDBClusterParams_ProxySQL.Marshal(b, m, deterministic)
}
func (m *XtraDBClusterParams_ProxySQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtraDBClusterParams_ProxySQL.Merge(m, src)
}
func (m *XtraDBClusterParams_ProxySQL) XXX_Size() int {
	return xxx_messageInfo_XtraDBClusterParams_ProxySQL.Size(m)
}
func (m *XtraDBClusterParams_ProxySQL) XXX_DiscardUnknown() {
	xxx_messageInfo_XtraDBClusterParams_ProxySQL.DiscardUnknown(m)
}

var xxx_messageInfo_XtraDBClusterParams_ProxySQL proto.InternalMessageInfo

func (m *XtraDBClusterParams_ProxySQL) GetComputeResources() *ComputeResources {
	if m != nil {
		return m.ComputeResources
	}
	return nil
}

func init() {
	proto.RegisterEnum("percona.platform.dbaas.controller.v1beta1.XtraDBClusterState", XtraDBClusterState_name, XtraDBClusterState_value)
	proto.RegisterEnum("percona.platform.dbaas.controller.v1beta1.XtraDBBackupState", XtraDBBackupState_name, XtraDBBackupState_value)
	proto.RegisterType((*XtraDBClusterParams)(nil), "percona.platform.dbaas.controller.v1beta1.XtraDBClusterParams")
	proto.RegisterType((*XtraDBClusterParams_PXC)(nil), "percona.platform.dbaas.controller.v1beta1.XtraDBClusterParams.PXC")
	proto.RegisterType((*XtraDBClusterParams_ProxySQL)(nil), "percona.platform.dbaas.controller.v1beta1.XtraDBClusterParams.ProxySQL")
}

func init() {
	proto.RegisterFile("controller/xtra_db_cluster.proto", fileDescriptor_1effda7bec99251f)
}

var fileDescriptor_1effda7bec99251f = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0x76, 0x4c, 0x93, 0xcb, 0x45, 0x66, 0x2e, 0xa8, 0xca, 0xd8, 0xc2, 0xae, 0xb6,
	0x49, 0x73, 0xb5, 0x21, 0x71, 0xb3, 0xab, 0xc4, 0x31, 0xa5, 0x22, 0x8a, 0x4a, 0x92, 0xa2, 0x82,
	0x90, 0x22, 0xc7, 0x35, 0x23, 0x6a, 0x52, 0x07, 0xc7, 0x59, 0xcb, 0xc4, 0x15, 0xcf, 0xc1, 0xd3,
	0xf0, 0x22, 0x48, 0x3c, 0x09, 0x6a, 0xd3, 0xbf, 0xd0, 0x4a, 0x20, 0x24, 0xee, 0x2c, 0x9d, 0xef,
	0x7c, 0xbf, 0xf3, 0x9d, 0x23, 0x03, 0x83, 0x89, 0xa1, 0x92, 0x22, 0x49, 0xb8, 0x6c, 0x8e, 0x95,
	0xa4, 0x61, 0x3f, 0x0a, 0x59, 0x52, 0xe4, 0x8a, 0x4b, 0x94, 0x49, 0xa1, 0x04, 0x3c, 0xcb, 0xb8,
	0x64, 0x62, 0x48, 0x51, 0x96, 0x50, 0xf5, 0x5e, 0xc8, 0x14, 0xf5, 0x23, 0x4a, 0x73, 0xb4, 0x6c,
	0x44, 0xb7, 0x97, 0x11, 0x57, 0xf4, 0xb2, 0xf1, 0x70, 0xc5, 0x8c, 0x89, 0x34, 0x15, 0xc3, 0xd2,
	0xa3, 0xf1, 0xec, 0x26, 0x56, 0x1f, 0x8a, 0x08, 0x31, 0x91, 0x36, 0xd3, 0x51, 0xac, 0x06, 0x62,
	0xd4, 0xbc, 0x11, 0x17, 0xd3, 0xe2, 0xc5, 0x2d, 0x4d, 0xe2, 0x3e, 0x55, 0x42, 0xe6, 0xcd, 0xc5,
	0xb3, 0xec, 0x3b, 0xf9, 0xb2, 0x0b, 0x1e, 0xf4, 0x94, 0xa4, 0xb6, 0x85, 0xcb, 0x99, 0x3a, 0x54,
	0xd2, 0x34, 0x87, 0x67, 0xe0, 0xfe, 0x6c, 0xc8, 0x30, 0x8f, 0xef, 0x78, 0x5d, 0x33, 0xb4, 0xd3,
	0x7b, 0xd6, 0xde, 0x8f, 0xef, 0xc7, 0x15, 0x7d, 0xc7, 0xab, 0xcd, 0x6a, 0x7e, 0x7c, 0xc7, 0xe1,
	0x3b, 0x50, 0xcd, 0xc6, 0xac, 0x5e, 0x31, 0xb4, 0xd3, 0xda, 0x95, 0x85, 0xfe, 0x38, 0x0c, 0xda,
	0xc0, 0x45, 0x9d, 0x1e, 0x2e, 0x29, 0x86, 0xe6, 0x4d, 0x6c, 0xe1, 0x00, 0xec, 0x67, 0x52, 0x8c,
	0x3f, 0xe5, 0x1f, 0x93, 0x7a, 0x75, 0x8a, 0x68, 0xfd, 0x2b, 0x62, 0x62, 0xe7, 0xbf, 0x72, 0x16,
	0x9c, 0x05, 0xa0, 0x31, 0x02, 0xd5, 0x4e, 0x0f, 0xc3, 0x0c, 0x1c, 0x30, 0x91, 0x66, 0x85, 0xe2,
	0xa1, 0xe4, 0xb9, 0x28, 0x24, 0xe3, 0xf9, 0x74, 0x03, 0xb5, 0xab, 0xeb, 0xbf, 0x80, 0xe3, 0xd2,
	0xc3, 0x9b, 0x5b, 0x2c, 0x80, 0x3a, 0xfb, 0xa5, 0xd2, 0xf8, 0x0c, 0xf6, 0xe7, 0x63, 0xfd, 0x7f,
	0xfa, 0xf9, 0x37, 0x0d, 0xc0, 0xb5, 0x4d, 0xf9, 0x8a, 0x2a, 0x0e, 0x9f, 0x80, 0xc7, 0xbd, 0xc0,
	0x33, 0x43, 0xdb, 0x0a, 0xb1, 0xd3, 0xf5, 0x03, 0xe2, 0x85, 0x7e, 0x60, 0x06, 0x24, 0x6c, 0xbb,
	0xaf, 0x4d, 0xa7, 0x6d, 0xeb, 0x3b, 0xf0, 0x04, 0x1c, 0x6d, 0x96, 0xe0, 0x17, 0xa6, 0xdb, 0x6a,
	0xbb, 0x2d, 0x5d, 0x83, 0xc7, 0xe0, 0xd1, 0x66, 0x8d, 0x47, 0x4c, 0xfb, 0x8d, 0x5e, 0x81, 0x06,
	0x38, 0xdc, 0x2c, 0x78, 0x6e, 0xb6, 0x1d, 0x62, 0xeb, 0xd5, 0xed, 0x18, 0x9b, 0x38, 0x24, 0x98,
	0x60, 0x76, 0xcf, 0xbf, 0x6a, 0xe0, 0xa0, 0x0c, 0x61, 0x51, 0x36, 0x28, 0xb2, 0x32, 0xc3, 0x8a,
	0xb7, 0x65, 0xe2, 0x97, 0xdd, 0xce, 0x6f, 0x11, 0xb6, 0x29, 0xbc, 0xae, 0xeb, 0x96, 0x01, 0x56,
	0xe8, 0x6b, 0x0a, 0xbf, 0x8b, 0x31, 0x21, 0x36, 0xb1, 0xf5, 0xca, 0x6a, 0xc8, 0x35, 0xcd, 0x3c,
	0x82, 0x75, 0xf4, 0xf6, 0x70, 0x79, 0xa4, 0xeb, 0xe5, 0x73, 0x76, 0xae, 0x68, 0x6f, 0xfa, 0x1f,
	0x9f, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc8, 0xba, 0xf5, 0x2f, 0x04, 0x00, 0x00,
}
