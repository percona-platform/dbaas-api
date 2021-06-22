// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller/kubernetes_cluster.proto

package controllerv1beta1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// OperatorsStatus defines status of operators installed in Kubernetes cluster.
type OperatorsStatus int32

const (
	// OPERATORS_STATUS_INVALID represents unknown state.
	OperatorsStatus_OPERATORS_STATUS_INVALID OperatorsStatus = 0
	// OPERATORS_STATUS_OK represents that operators are installed and have supported API version.
	OperatorsStatus_OPERATORS_STATUS_OK OperatorsStatus = 1
	// OPERATORS_STATUS_NOT_INSTALLED represents that operators are not installed.
	OperatorsStatus_OPERATORS_STATUS_NOT_INSTALLED OperatorsStatus = 3
)

// Enum value maps for OperatorsStatus.
var (
	OperatorsStatus_name = map[int32]string{
		0: "OPERATORS_STATUS_INVALID",
		1: "OPERATORS_STATUS_OK",
		3: "OPERATORS_STATUS_NOT_INSTALLED",
	}
	OperatorsStatus_value = map[string]int32{
		"OPERATORS_STATUS_INVALID":       0,
		"OPERATORS_STATUS_OK":            1,
		"OPERATORS_STATUS_NOT_INSTALLED": 3,
	}
)

func (x OperatorsStatus) Enum() *OperatorsStatus {
	p := new(OperatorsStatus)
	*p = x
	return p
}

func (x OperatorsStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatorsStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_kubernetes_cluster_proto_enumTypes[0].Descriptor()
}

func (OperatorsStatus) Type() protoreflect.EnumType {
	return &file_controller_kubernetes_cluster_proto_enumTypes[0]
}

func (x OperatorsStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatorsStatus.Descriptor instead.
func (OperatorsStatus) EnumDescriptor() ([]byte, []int) {
	return file_controller_kubernetes_cluster_proto_rawDescGZIP(), []int{0}
}

// KubernetesClusterStatus defines status of Kubernetes cluster.
type KubernetesClusterStatus int32

const (
	// KUBERNETES_CLUSTER_STATUS_INVALID represents unknown state.
	KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_INVALID KubernetesClusterStatus = 0
	// KUBERNETES_CLUSTER_STATUS_OK represents that Kubernetes cluster is accessible.
	KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_OK KubernetesClusterStatus = 1
	// KUBERNETES_CLUSTER_STATUS_UNAVAILABLE represents that Kubernetes cluster is not accessible.
	KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_UNAVAILABLE KubernetesClusterStatus = 2
)

// Enum value maps for KubernetesClusterStatus.
var (
	KubernetesClusterStatus_name = map[int32]string{
		0: "KUBERNETES_CLUSTER_STATUS_INVALID",
		1: "KUBERNETES_CLUSTER_STATUS_OK",
		2: "KUBERNETES_CLUSTER_STATUS_UNAVAILABLE",
	}
	KubernetesClusterStatus_value = map[string]int32{
		"KUBERNETES_CLUSTER_STATUS_INVALID":     0,
		"KUBERNETES_CLUSTER_STATUS_OK":          1,
		"KUBERNETES_CLUSTER_STATUS_UNAVAILABLE": 2,
	}
)

func (x KubernetesClusterStatus) Enum() *KubernetesClusterStatus {
	p := new(KubernetesClusterStatus)
	*p = x
	return p
}

func (x KubernetesClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_kubernetes_cluster_proto_enumTypes[1].Descriptor()
}

func (KubernetesClusterStatus) Type() protoreflect.EnumType {
	return &file_controller_kubernetes_cluster_proto_enumTypes[1]
}

func (x KubernetesClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesClusterStatus.Descriptor instead.
func (KubernetesClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_controller_kubernetes_cluster_proto_rawDescGZIP(), []int{1}
}

// Operator contains all information about operator installed in Kubernetes cluster.
type Operator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the operator.
	Status OperatorsStatus `protobuf:"varint,1,opt,name=status,proto3,enum=percona.platform.dbaas.controller.v1beta1.OperatorsStatus" json:"status,omitempty"`
	// Installed version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Operator) Reset() {
	*x = Operator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_kubernetes_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operator) ProtoMessage() {}

func (x *Operator) ProtoReflect() protoreflect.Message {
	mi := &file_controller_kubernetes_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operator.ProtoReflect.Descriptor instead.
func (*Operator) Descriptor() ([]byte, []int) {
	return file_controller_kubernetes_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Operator) GetStatus() OperatorsStatus {
	if x != nil {
		return x.Status
	}
	return OperatorsStatus_OPERATORS_STATUS_INVALID
}

func (x *Operator) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Operators contains list of operators installed in Kubernetes cluster.
type Operators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percona XtraDB Cluster Operator.
	Xtradb *Operator `protobuf:"bytes,1,opt,name=xtradb,proto3" json:"xtradb,omitempty"`
	// Percona Server for MongoDB Operator.
	Psmdb *Operator `protobuf:"bytes,2,opt,name=psmdb,proto3" json:"psmdb,omitempty"`
}

func (x *Operators) Reset() {
	*x = Operators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_kubernetes_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operators) ProtoMessage() {}

func (x *Operators) ProtoReflect() protoreflect.Message {
	mi := &file_controller_kubernetes_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operators.ProtoReflect.Descriptor instead.
func (*Operators) Descriptor() ([]byte, []int) {
	return file_controller_kubernetes_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Operators) GetXtradb() *Operator {
	if x != nil {
		return x.Xtradb
	}
	return nil
}

func (x *Operators) GetPsmdb() *Operator {
	if x != nil {
		return x.Psmdb
	}
	return nil
}

// Resources contains Kubernetes cluster resources.
type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Memory in bytes.
	MemoryBytes uint64 `protobuf:"varint,1,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CpuM uint64 `protobuf:"varint,2,opt,name=cpu_m,json=cpuM,proto3" json:"cpu_m,omitempty"`
	// Disk size in bytes.
	// NOTE: Values are large for AWS EBS because there is a rather loose limit of
	// how much storage one can request. Maximum is 39*16TiB per node. It gets out
	// of int64 range fast even for small clusters, that's why we use uint64.
	DiskSize uint64 `protobuf:"varint,3,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_kubernetes_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_controller_kubernetes_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_controller_kubernetes_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *Resources) GetMemoryBytes() uint64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

func (x *Resources) GetCpuM() uint64 {
	if x != nil {
		return x.CpuM
	}
	return 0
}

func (x *Resources) GetDiskSize() uint64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

var File_controller_kubernetes_cluster_proto protoreflect.FileDescriptor

var file_controller_kubernetes_cluster_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x22, 0x78, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x52, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x70,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x06, 0x78, 0x74, 0x72, 0x61,
	0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x78,
	0x74, 0x72, 0x61, 0x64, 0x62, 0x12, 0x49, 0x0a, 0x05, 0x70, 0x73, 0x6d, 0x64, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x70, 0x73, 0x6d, 0x64, 0x62,
	0x22, 0x60, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x13, 0x0a, 0x05, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x63, 0x70, 0x75, 0x4d, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69,
	0x7a, 0x65, 0x2a, 0x6c, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x53,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x2a, 0x8d, 0x01, 0x0a, 0x17, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x21,
	0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45,
	0x53, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45,
	0x54, 0x45, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02,
	0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x3b, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_kubernetes_cluster_proto_rawDescOnce sync.Once
	file_controller_kubernetes_cluster_proto_rawDescData = file_controller_kubernetes_cluster_proto_rawDesc
)

func file_controller_kubernetes_cluster_proto_rawDescGZIP() []byte {
	file_controller_kubernetes_cluster_proto_rawDescOnce.Do(func() {
		file_controller_kubernetes_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_kubernetes_cluster_proto_rawDescData)
	})
	return file_controller_kubernetes_cluster_proto_rawDescData
}

var file_controller_kubernetes_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_controller_kubernetes_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_kubernetes_cluster_proto_goTypes = []interface{}{
	(OperatorsStatus)(0),         // 0: percona.platform.dbaas.controller.v1beta1.OperatorsStatus
	(KubernetesClusterStatus)(0), // 1: percona.platform.dbaas.controller.v1beta1.KubernetesClusterStatus
	(*Operator)(nil),             // 2: percona.platform.dbaas.controller.v1beta1.Operator
	(*Operators)(nil),            // 3: percona.platform.dbaas.controller.v1beta1.Operators
	(*Resources)(nil),            // 4: percona.platform.dbaas.controller.v1beta1.Resources
}
var file_controller_kubernetes_cluster_proto_depIdxs = []int32{
	0, // 0: percona.platform.dbaas.controller.v1beta1.Operator.status:type_name -> percona.platform.dbaas.controller.v1beta1.OperatorsStatus
	2, // 1: percona.platform.dbaas.controller.v1beta1.Operators.xtradb:type_name -> percona.platform.dbaas.controller.v1beta1.Operator
	2, // 2: percona.platform.dbaas.controller.v1beta1.Operators.psmdb:type_name -> percona.platform.dbaas.controller.v1beta1.Operator
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controller_kubernetes_cluster_proto_init() }
func file_controller_kubernetes_cluster_proto_init() {
	if File_controller_kubernetes_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_kubernetes_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operator); i {
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
		file_controller_kubernetes_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operators); i {
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
		file_controller_kubernetes_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
			RawDescriptor: file_controller_kubernetes_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_kubernetes_cluster_proto_goTypes,
		DependencyIndexes: file_controller_kubernetes_cluster_proto_depIdxs,
		EnumInfos:         file_controller_kubernetes_cluster_proto_enumTypes,
		MessageInfos:      file_controller_kubernetes_cluster_proto_msgTypes,
	}.Build()
	File_controller_kubernetes_cluster_proto = out.File
	file_controller_kubernetes_cluster_proto_rawDesc = nil
	file_controller_kubernetes_cluster_proto_goTypes = nil
	file_controller_kubernetes_cluster_proto_depIdxs = nil
}
