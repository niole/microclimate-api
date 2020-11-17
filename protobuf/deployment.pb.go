// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: deployment.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Deployment_Status int32

const (
	Deployment_STARTING_UP   Deployment_Status = 0
	Deployment_FAILED        Deployment_Status = 1
	Deployment_RUNNING       Deployment_Status = 2
	Deployment_SHUTTING_DOWN Deployment_Status = 3
	Deployment_UNREACHABLE   Deployment_Status = 4
)

// Enum value maps for Deployment_Status.
var (
	Deployment_Status_name = map[int32]string{
		0: "STARTING_UP",
		1: "FAILED",
		2: "RUNNING",
		3: "SHUTTING_DOWN",
		4: "UNREACHABLE",
	}
	Deployment_Status_value = map[string]int32{
		"STARTING_UP":   0,
		"FAILED":        1,
		"RUNNING":       2,
		"SHUTTING_DOWN": 3,
		"UNREACHABLE":   4,
	}
)

func (x Deployment_Status) Enum() *Deployment_Status {
	p := new(Deployment_Status)
	*p = x
	return p
}

func (x Deployment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Deployment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_deployment_proto_enumTypes[0].Descriptor()
}

func (Deployment_Status) Type() protoreflect.EnumType {
	return &file_deployment_proto_enumTypes[0]
}

func (x Deployment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Deployment_Status.Descriptor instead.
func (Deployment_Status) EnumDescriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0, 0}
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerUserId string            `protobuf:"bytes,2,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	Status      Deployment_Status `protobuf:"varint,4,opt,name=status,proto3,enum=api.Deployment_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *Deployment) GetStatus() Deployment_Status {
	if x != nil {
		return x.Status
	}
	return Deployment_STARTING_UP
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserId string `protobuf:"bytes,1,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewDeployment) Reset() {
	*x = NewDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDeployment) ProtoMessage() {}

func (x *NewDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDeployment.ProtoReflect.Descriptor instead.
func (*NewDeployment) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *NewDeployment) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *NewDeployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDeploymentsForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetDeploymentsForUserRequest) Reset() {
	*x = GetDeploymentsForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentsForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentsForUserRequest) ProtoMessage() {}

func (x *GetDeploymentsForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentsForUserRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentsForUserRequest) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeploymentsForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserId  string `protobuf:"bytes,1,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	DeploymentId string `protobuf:"bytes,2,opt,name=deploymentId,proto3" json:"deploymentId,omitempty"`
}

func (x *GetDeploymentRequest) Reset() {
	*x = GetDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentRequest) ProtoMessage() {}

func (x *GetDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeploymentRequest) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *GetDeploymentRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

type RemoveDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string `protobuf:"bytes,1,opt,name=deploymentId,proto3" json:"deploymentId,omitempty"`
}

func (x *RemoveDeploymentRequest) Reset() {
	*x = RemoveDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDeploymentRequest) ProtoMessage() {}

func (x *RemoveDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDeploymentRequest.ProtoReflect.Descriptor instead.
func (*RemoveDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveDeploymentRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

var File_deployment_proto protoreflect.FileDescriptor

var file_deployment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x55, 0x50, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x48, 0x55, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x04, 0x22, 0x45, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3d,
	0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32, 0xa0, 0x02,
	0x0a, 0x1b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deployment_proto_rawDescOnce sync.Once
	file_deployment_proto_rawDescData = file_deployment_proto_rawDesc
)

func file_deployment_proto_rawDescGZIP() []byte {
	file_deployment_proto_rawDescOnce.Do(func() {
		file_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_proto_rawDescData)
	})
	return file_deployment_proto_rawDescData
}

var file_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_deployment_proto_goTypes = []interface{}{
	(Deployment_Status)(0),               // 0: api.Deployment.Status
	(*Deployment)(nil),                   // 1: api.Deployment
	(*NewDeployment)(nil),                // 2: api.NewDeployment
	(*GetDeploymentsForUserRequest)(nil), // 3: api.GetDeploymentsForUserRequest
	(*GetDeploymentRequest)(nil),         // 4: api.GetDeploymentRequest
	(*RemoveDeploymentRequest)(nil),      // 5: api.RemoveDeploymentRequest
	(*Empty)(nil),                        // 6: api.Empty
}
var file_deployment_proto_depIdxs = []int32{
	0, // 0: api.Deployment.status:type_name -> api.Deployment.Status
	3, // 1: api.DeploymentManagementService.GetDeploymentsForUser:input_type -> api.GetDeploymentsForUserRequest
	2, // 2: api.DeploymentManagementService.CreateDeployment:input_type -> api.NewDeployment
	4, // 3: api.DeploymentManagementService.GetDeployment:input_type -> api.GetDeploymentRequest
	5, // 4: api.DeploymentManagementService.RemoveDeployment:input_type -> api.RemoveDeploymentRequest
	1, // 5: api.DeploymentManagementService.GetDeploymentsForUser:output_type -> api.Deployment
	1, // 6: api.DeploymentManagementService.CreateDeployment:output_type -> api.Deployment
	1, // 7: api.DeploymentManagementService.GetDeployment:output_type -> api.Deployment
	6, // 8: api.DeploymentManagementService.RemoveDeployment:output_type -> api.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_deployment_proto_init() }
func file_deployment_proto_init() {
	if File_deployment_proto != nil {
		return
	}
	file_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewDeployment); i {
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
		file_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentsForUserRequest); i {
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
		file_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentRequest); i {
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
		file_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDeploymentRequest); i {
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
			RawDescriptor: file_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deployment_proto_goTypes,
		DependencyIndexes: file_deployment_proto_depIdxs,
		EnumInfos:         file_deployment_proto_enumTypes,
		MessageInfos:      file_deployment_proto_msgTypes,
	}.Build()
	File_deployment_proto = out.File
	file_deployment_proto_rawDesc = nil
	file_deployment_proto_goTypes = nil
	file_deployment_proto_depIdxs = nil
}
