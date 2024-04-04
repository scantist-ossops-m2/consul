// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pbmulticluster/v2beta1/sameness_group.proto

package multiclusterv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type SamenessGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultForFailover bool                   `protobuf:"varint,1,opt,name=default_for_failover,json=defaultForFailover,proto3" json:"default_for_failover,omitempty"`
	Members            []*SamenessGroupMember `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *SamenessGroup) Reset() {
	*x = SamenessGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamenessGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamenessGroup) ProtoMessage() {}

func (x *SamenessGroup) ProtoReflect() protoreflect.Message {
	mi := &file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamenessGroup.ProtoReflect.Descriptor instead.
func (*SamenessGroup) Descriptor() ([]byte, []int) {
	return file_pbmulticluster_v2beta1_sameness_group_proto_rawDescGZIP(), []int{0}
}

func (x *SamenessGroup) GetDefaultForFailover() bool {
	if x != nil {
		return x.DefaultForFailover
	}
	return false
}

func (x *SamenessGroup) GetMembers() []*SamenessGroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

type SamenessGroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Member:
	//
	//	*SamenessGroupMember_Peer
	//	*SamenessGroupMember_Partition
	Member isSamenessGroupMember_Member `protobuf_oneof:"member"`
}

func (x *SamenessGroupMember) Reset() {
	*x = SamenessGroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamenessGroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamenessGroupMember) ProtoMessage() {}

func (x *SamenessGroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamenessGroupMember.ProtoReflect.Descriptor instead.
func (*SamenessGroupMember) Descriptor() ([]byte, []int) {
	return file_pbmulticluster_v2beta1_sameness_group_proto_rawDescGZIP(), []int{1}
}

func (m *SamenessGroupMember) GetMember() isSamenessGroupMember_Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (x *SamenessGroupMember) GetPeer() string {
	if x, ok := x.GetMember().(*SamenessGroupMember_Peer); ok {
		return x.Peer
	}
	return ""
}

func (x *SamenessGroupMember) GetPartition() string {
	if x, ok := x.GetMember().(*SamenessGroupMember_Partition); ok {
		return x.Partition
	}
	return ""
}

type isSamenessGroupMember_Member interface {
	isSamenessGroupMember_Member()
}

type SamenessGroupMember_Peer struct {
	Peer string `protobuf:"bytes,1,opt,name=peer,proto3,oneof"`
}

type SamenessGroupMember_Partition struct {
	Partition string `protobuf:"bytes,2,opt,name=partition,proto3,oneof"`
}

func (*SamenessGroupMember_Peer) isSamenessGroupMember_Member() {}

func (*SamenessGroupMember_Partition) isSamenessGroupMember_Member() {}

var File_pbmulticluster_v2beta1_sameness_group_proto protoreflect.FileDescriptor

var file_pbmulticluster_v2beta1_sameness_group_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x62, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x46, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x06, 0xa2, 0x93,
	0x04, 0x02, 0x08, 0x02, 0x22, 0x55, 0x0a, 0x13, 0x53, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x73, 0x73,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0xcb, 0x02, 0x0a, 0x29,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x12, 0x53, 0x61, 0x6d, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x25, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x25, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x31, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x28, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pbmulticluster_v2beta1_sameness_group_proto_rawDescOnce sync.Once
	file_pbmulticluster_v2beta1_sameness_group_proto_rawDescData = file_pbmulticluster_v2beta1_sameness_group_proto_rawDesc
)

func file_pbmulticluster_v2beta1_sameness_group_proto_rawDescGZIP() []byte {
	file_pbmulticluster_v2beta1_sameness_group_proto_rawDescOnce.Do(func() {
		file_pbmulticluster_v2beta1_sameness_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmulticluster_v2beta1_sameness_group_proto_rawDescData)
	})
	return file_pbmulticluster_v2beta1_sameness_group_proto_rawDescData
}

var file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbmulticluster_v2beta1_sameness_group_proto_goTypes = []interface{}{
	(*SamenessGroup)(nil),       // 0: hashicorp.consul.multicluster.v2beta1.SamenessGroup
	(*SamenessGroupMember)(nil), // 1: hashicorp.consul.multicluster.v2beta1.SamenessGroupMember
}
var file_pbmulticluster_v2beta1_sameness_group_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.multicluster.v2beta1.SamenessGroup.members:type_name -> hashicorp.consul.multicluster.v2beta1.SamenessGroupMember
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pbmulticluster_v2beta1_sameness_group_proto_init() }
func file_pbmulticluster_v2beta1_sameness_group_proto_init() {
	if File_pbmulticluster_v2beta1_sameness_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamenessGroup); i {
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
		file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamenessGroupMember); i {
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
	file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SamenessGroupMember_Peer)(nil),
		(*SamenessGroupMember_Partition)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmulticluster_v2beta1_sameness_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmulticluster_v2beta1_sameness_group_proto_goTypes,
		DependencyIndexes: file_pbmulticluster_v2beta1_sameness_group_proto_depIdxs,
		MessageInfos:      file_pbmulticluster_v2beta1_sameness_group_proto_msgTypes,
	}.Build()
	File_pbmulticluster_v2beta1_sameness_group_proto = out.File
	file_pbmulticluster_v2beta1_sameness_group_proto_rawDesc = nil
	file_pbmulticluster_v2beta1_sameness_group_proto_goTypes = nil
	file_pbmulticluster_v2beta1_sameness_group_proto_depIdxs = nil
}
