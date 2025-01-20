// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pkg/testing/remote/loader.proto

package remote

import (
	server "github.com/linuxsuren/api-testing/pkg/server"
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

type TestSuites struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TestSuite `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TestSuites) Reset() {
	*x = TestSuites{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSuites) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSuites) ProtoMessage() {}

func (x *TestSuites) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSuites.ProtoReflect.Descriptor instead.
func (*TestSuites) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{0}
}

func (x *TestSuites) GetData() []*TestSuite {
	if x != nil {
		return x.Data
	}
	return nil
}

type TestSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Api   string             `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	Param []*server.Pair     `protobuf:"bytes,3,rep,name=param,proto3" json:"param,omitempty"`
	Spec  *server.APISpec    `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Items []*server.TestCase `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	Full  bool               `protobuf:"varint,6,opt,name=full,proto3" json:"full,omitempty"`
}

func (x *TestSuite) Reset() {
	*x = TestSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSuite) ProtoMessage() {}

func (x *TestSuite) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSuite.ProtoReflect.Descriptor instead.
func (*TestSuite) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{1}
}

func (x *TestSuite) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestSuite) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *TestSuite) GetParam() []*server.Pair {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *TestSuite) GetSpec() *server.APISpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TestSuite) GetItems() []*server.TestCase {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TestSuite) GetFull() bool {
	if x != nil {
		return x.Full
	}
	return false
}

type HistoryTestSuites struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*HistoryTestSuite `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HistoryTestSuites) Reset() {
	*x = HistoryTestSuites{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTestSuites) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTestSuites) ProtoMessage() {}

func (x *HistoryTestSuites) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTestSuites.ProtoReflect.Descriptor instead.
func (*HistoryTestSuites) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryTestSuites) GetData() []*HistoryTestSuite {
	if x != nil {
		return x.Data
	}
	return nil
}

type HistoryTestSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistorySuiteName string                    `protobuf:"bytes,1,opt,name=historySuiteName,proto3" json:"historySuiteName,omitempty"`
	Items            []*server.HistoryTestCase `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *HistoryTestSuite) Reset() {
	*x = HistoryTestSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTestSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTestSuite) ProtoMessage() {}

func (x *HistoryTestSuite) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTestSuite.ProtoReflect.Descriptor instead.
func (*HistoryTestSuite) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{3}
}

func (x *HistoryTestSuite) GetHistorySuiteName() string {
	if x != nil {
		return x.HistorySuiteName
	}
	return ""
}

func (x *HistoryTestSuite) GetItems() []*server.HistoryTestCase {
	if x != nil {
		return x.Items
	}
	return nil
}

type Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Config `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Configs) Reset() {
	*x = Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{4}
}

func (x *Configs) GetData() []*Config {
	if x != nil {
		return x.Data
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value       string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_testing_remote_loader_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_testing_remote_loader_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_pkg_testing_remote_loader_proto_rawDescGZIP(), []int{5}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Config) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_pkg_testing_remote_loader_proto protoreflect.FileDescriptor

var file_pkg_testing_remote_loader_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x22, 0x0a, 0x05, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x23, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x75, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x75, 0x6c, 0x6c,
	0x22, 0x41, 0x0a, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x6d, 0x0a, 0x10, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x53, 0x75, 0x69, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x75, 0x69, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x22, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x54, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc8, 0x0a, 0x0a, 0x06, 0x4c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75,
	0x69, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x0d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x11, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x44,
	0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x11, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x0d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x0d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x41, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x05, 0x50, 0x50, 0x72, 0x6f, 0x66, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x50, 0x72, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x50, 0x72, 0x6f, 0x66, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x32, 0x96, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x32, 0x9e, 0x02, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x0d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0e,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x73, 0x75, 0x72, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_testing_remote_loader_proto_rawDescOnce sync.Once
	file_pkg_testing_remote_loader_proto_rawDescData = file_pkg_testing_remote_loader_proto_rawDesc
)

func file_pkg_testing_remote_loader_proto_rawDescGZIP() []byte {
	file_pkg_testing_remote_loader_proto_rawDescOnce.Do(func() {
		file_pkg_testing_remote_loader_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_testing_remote_loader_proto_rawDescData)
	})
	return file_pkg_testing_remote_loader_proto_rawDescData
}

var file_pkg_testing_remote_loader_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_testing_remote_loader_proto_goTypes = []interface{}{
	(*TestSuites)(nil),                // 0: remote.TestSuites
	(*TestSuite)(nil),                 // 1: remote.TestSuite
	(*HistoryTestSuites)(nil),         // 2: remote.HistoryTestSuites
	(*HistoryTestSuite)(nil),          // 3: remote.HistoryTestSuite
	(*Configs)(nil),                   // 4: remote.Configs
	(*Config)(nil),                    // 5: remote.Config
	(*server.Pair)(nil),               // 6: server.Pair
	(*server.APISpec)(nil),            // 7: server.APISpec
	(*server.TestCase)(nil),           // 8: server.TestCase
	(*server.HistoryTestCase)(nil),    // 9: server.HistoryTestCase
	(*server.Empty)(nil),              // 10: server.Empty
	(*server.TestSuiteDuplicate)(nil), // 11: server.TestSuiteDuplicate
	(*server.TestCaseDuplicate)(nil),  // 12: server.TestCaseDuplicate
	(*server.HistoryTestResult)(nil),  // 13: server.HistoryTestResult
	(*server.PProfRequest)(nil),       // 14: server.PProfRequest
	(*server.Secret)(nil),             // 15: server.Secret
	(*server.SimpleName)(nil),         // 16: server.SimpleName
	(*server.HelloReply)(nil),         // 17: server.HelloReply
	(*server.TestCases)(nil),          // 18: server.TestCases
	(*server.HistoryTestCases)(nil),   // 19: server.HistoryTestCases
	(*server.Version)(nil),            // 20: server.Version
	(*server.ExtensionStatus)(nil),    // 21: server.ExtensionStatus
	(*server.PProfData)(nil),          // 22: server.PProfData
	(*server.Secrets)(nil),            // 23: server.Secrets
	(*server.CommonResult)(nil),       // 24: server.CommonResult
}
var file_pkg_testing_remote_loader_proto_depIdxs = []int32{
	1,  // 0: remote.TestSuites.data:type_name -> remote.TestSuite
	6,  // 1: remote.TestSuite.param:type_name -> server.Pair
	7,  // 2: remote.TestSuite.spec:type_name -> server.APISpec
	8,  // 3: remote.TestSuite.items:type_name -> server.TestCase
	3,  // 4: remote.HistoryTestSuites.data:type_name -> remote.HistoryTestSuite
	9,  // 5: remote.HistoryTestSuite.items:type_name -> server.HistoryTestCase
	5,  // 6: remote.Configs.data:type_name -> remote.Config
	10, // 7: remote.Loader.ListTestSuite:input_type -> server.Empty
	1,  // 8: remote.Loader.CreateTestSuite:input_type -> remote.TestSuite
	1,  // 9: remote.Loader.GetTestSuite:input_type -> remote.TestSuite
	1,  // 10: remote.Loader.UpdateTestSuite:input_type -> remote.TestSuite
	1,  // 11: remote.Loader.DeleteTestSuite:input_type -> remote.TestSuite
	11, // 12: remote.Loader.RenameTestSuite:input_type -> server.TestSuiteDuplicate
	1,  // 13: remote.Loader.ListTestCases:input_type -> remote.TestSuite
	8,  // 14: remote.Loader.CreateTestCase:input_type -> server.TestCase
	8,  // 15: remote.Loader.GetTestCase:input_type -> server.TestCase
	8,  // 16: remote.Loader.UpdateTestCase:input_type -> server.TestCase
	8,  // 17: remote.Loader.DeleteTestCase:input_type -> server.TestCase
	12, // 18: remote.Loader.RenameTestCase:input_type -> server.TestCaseDuplicate
	10, // 19: remote.Loader.ListHistoryTestSuite:input_type -> server.Empty
	13, // 20: remote.Loader.CreateTestCaseHistory:input_type -> server.HistoryTestResult
	9,  // 21: remote.Loader.GetHistoryTestCaseWithResult:input_type -> server.HistoryTestCase
	9,  // 22: remote.Loader.GetHistoryTestCase:input_type -> server.HistoryTestCase
	9,  // 23: remote.Loader.DeleteHistoryTestCase:input_type -> server.HistoryTestCase
	9,  // 24: remote.Loader.DeleteAllHistoryTestCase:input_type -> server.HistoryTestCase
	8,  // 25: remote.Loader.GetTestCaseAllHistory:input_type -> server.TestCase
	10, // 26: remote.Loader.GetVersion:input_type -> server.Empty
	10, // 27: remote.Loader.Verify:input_type -> server.Empty
	14, // 28: remote.Loader.PProf:input_type -> server.PProfRequest
	15, // 29: remote.SecretService.GetSecret:input_type -> server.Secret
	10, // 30: remote.SecretService.GetSecrets:input_type -> server.Empty
	15, // 31: remote.SecretService.CreateSecret:input_type -> server.Secret
	15, // 32: remote.SecretService.DeleteSecret:input_type -> server.Secret
	15, // 33: remote.SecretService.UpdateSecret:input_type -> server.Secret
	10, // 34: remote.ConfigService.GetConfigs:input_type -> server.Empty
	16, // 35: remote.ConfigService.GetConfig:input_type -> server.SimpleName
	5,  // 36: remote.ConfigService.CreateConfig:input_type -> remote.Config
	5,  // 37: remote.ConfigService.UpdateConfig:input_type -> remote.Config
	16, // 38: remote.ConfigService.DeleteConfig:input_type -> server.SimpleName
	0,  // 39: remote.Loader.ListTestSuite:output_type -> remote.TestSuites
	10, // 40: remote.Loader.CreateTestSuite:output_type -> server.Empty
	1,  // 41: remote.Loader.GetTestSuite:output_type -> remote.TestSuite
	1,  // 42: remote.Loader.UpdateTestSuite:output_type -> remote.TestSuite
	10, // 43: remote.Loader.DeleteTestSuite:output_type -> server.Empty
	17, // 44: remote.Loader.RenameTestSuite:output_type -> server.HelloReply
	18, // 45: remote.Loader.ListTestCases:output_type -> server.TestCases
	10, // 46: remote.Loader.CreateTestCase:output_type -> server.Empty
	8,  // 47: remote.Loader.GetTestCase:output_type -> server.TestCase
	8,  // 48: remote.Loader.UpdateTestCase:output_type -> server.TestCase
	10, // 49: remote.Loader.DeleteTestCase:output_type -> server.Empty
	17, // 50: remote.Loader.RenameTestCase:output_type -> server.HelloReply
	2,  // 51: remote.Loader.ListHistoryTestSuite:output_type -> remote.HistoryTestSuites
	10, // 52: remote.Loader.CreateTestCaseHistory:output_type -> server.Empty
	13, // 53: remote.Loader.GetHistoryTestCaseWithResult:output_type -> server.HistoryTestResult
	9,  // 54: remote.Loader.GetHistoryTestCase:output_type -> server.HistoryTestCase
	10, // 55: remote.Loader.DeleteHistoryTestCase:output_type -> server.Empty
	10, // 56: remote.Loader.DeleteAllHistoryTestCase:output_type -> server.Empty
	19, // 57: remote.Loader.GetTestCaseAllHistory:output_type -> server.HistoryTestCases
	20, // 58: remote.Loader.GetVersion:output_type -> server.Version
	21, // 59: remote.Loader.Verify:output_type -> server.ExtensionStatus
	22, // 60: remote.Loader.PProf:output_type -> server.PProfData
	15, // 61: remote.SecretService.GetSecret:output_type -> server.Secret
	23, // 62: remote.SecretService.GetSecrets:output_type -> server.Secrets
	24, // 63: remote.SecretService.CreateSecret:output_type -> server.CommonResult
	24, // 64: remote.SecretService.DeleteSecret:output_type -> server.CommonResult
	24, // 65: remote.SecretService.UpdateSecret:output_type -> server.CommonResult
	4,  // 66: remote.ConfigService.GetConfigs:output_type -> remote.Configs
	5,  // 67: remote.ConfigService.GetConfig:output_type -> remote.Config
	24, // 68: remote.ConfigService.CreateConfig:output_type -> server.CommonResult
	24, // 69: remote.ConfigService.UpdateConfig:output_type -> server.CommonResult
	24, // 70: remote.ConfigService.DeleteConfig:output_type -> server.CommonResult
	39, // [39:71] is the sub-list for method output_type
	7,  // [7:39] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_testing_remote_loader_proto_init() }
func file_pkg_testing_remote_loader_proto_init() {
	if File_pkg_testing_remote_loader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_testing_remote_loader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSuites); i {
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
		file_pkg_testing_remote_loader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSuite); i {
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
		file_pkg_testing_remote_loader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTestSuites); i {
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
		file_pkg_testing_remote_loader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTestSuite); i {
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
		file_pkg_testing_remote_loader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configs); i {
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
		file_pkg_testing_remote_loader_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_pkg_testing_remote_loader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_pkg_testing_remote_loader_proto_goTypes,
		DependencyIndexes: file_pkg_testing_remote_loader_proto_depIdxs,
		MessageInfos:      file_pkg_testing_remote_loader_proto_msgTypes,
	}.Build()
	File_pkg_testing_remote_loader_proto = out.File
	file_pkg_testing_remote_loader_proto_rawDesc = nil
	file_pkg_testing_remote_loader_proto_goTypes = nil
	file_pkg_testing_remote_loader_proto_depIdxs = nil
}
