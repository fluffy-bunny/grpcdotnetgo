// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: internal/grpcContracts/helloworld/helloworld.proto

package helloworld

import (
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

type HelloDirectives int32

const (
	HelloDirectives_HELLO_DIRECTIVES_UNKNOWN HelloDirectives = 0
	HelloDirectives_HELLO_DIRECTIVES_PANIC   HelloDirectives = 1
	HelloDirectives_HELLO_DIRECTIVES_ERROR   HelloDirectives = 2
)

// Enum value maps for HelloDirectives.
var (
	HelloDirectives_name = map[int32]string{
		0: "HELLO_DIRECTIVES_UNKNOWN",
		1: "HELLO_DIRECTIVES_PANIC",
		2: "HELLO_DIRECTIVES_ERROR",
	}
	HelloDirectives_value = map[string]int32{
		"HELLO_DIRECTIVES_UNKNOWN": 0,
		"HELLO_DIRECTIVES_PANIC":   1,
		"HELLO_DIRECTIVES_ERROR":   2,
	}
)

func (x HelloDirectives) Enum() *HelloDirectives {
	p := new(HelloDirectives)
	*p = x
	return p
}

func (x HelloDirectives) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloDirectives) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_grpcContracts_helloworld_helloworld_proto_enumTypes[0].Descriptor()
}

func (HelloDirectives) Type() protoreflect.EnumType {
	return &file_internal_grpcContracts_helloworld_helloworld_proto_enumTypes[0]
}

func (x HelloDirectives) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloDirectives.Descriptor instead.
func (HelloDirectives) EnumDescriptor() ([]byte, []int) {
	return file_internal_grpcContracts_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Directive HelloDirectives `protobuf:"varint,2,opt,name=Directive,proto3,enum=helloworld.HelloDirectives" json:"Directive,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcContracts_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetDirective() HelloDirectives {
	if x != nil {
		return x.Directive
	}
	return HelloDirectives_HELLO_DIRECTIVES_UNKNOWN
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` //error.Error error = 999;
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_internal_grpcContracts_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloReply2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply2) Reset() {
	*x = HelloReply2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply2) ProtoMessage() {}

func (x *HelloReply2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply2.ProtoReflect.Descriptor instead.
func (*HelloReply2) Descriptor() ([]byte, []int) {
	return file_internal_grpcContracts_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *HelloReply2) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_grpcContracts_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_internal_grpcContracts_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x22, 0x5d, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x73, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22,
	0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2a, 0x67, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x53, 0x5f, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x49, 0x0a, 0x07, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x32, 0x4b, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x32,
	0x12, 0x3f, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x22,
	0x00, 0x42, 0x67, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_grpcContracts_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_internal_grpcContracts_helloworld_helloworld_proto_rawDescData = file_internal_grpcContracts_helloworld_helloworld_proto_rawDesc
)

func file_internal_grpcContracts_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_internal_grpcContracts_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_internal_grpcContracts_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpcContracts_helloworld_helloworld_proto_rawDescData)
	})
	return file_internal_grpcContracts_helloworld_helloworld_proto_rawDescData
}

var file_internal_grpcContracts_helloworld_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_grpcContracts_helloworld_helloworld_proto_goTypes = []interface{}{
	(HelloDirectives)(0), // 0: helloworld.HelloDirectives
	(*HelloRequest)(nil), // 1: helloworld.HelloRequest
	(*HelloReply)(nil),   // 2: helloworld.HelloReply
	(*HelloReply2)(nil),  // 3: helloworld.HelloReply2
}
var file_internal_grpcContracts_helloworld_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.HelloRequest.Directive:type_name -> helloworld.HelloDirectives
	1, // 1: helloworld.Greeter.SayHello:input_type -> helloworld.HelloRequest
	1, // 2: helloworld.Greeter2.SayHello:input_type -> helloworld.HelloRequest
	2, // 3: helloworld.Greeter.SayHello:output_type -> helloworld.HelloReply
	3, // 4: helloworld.Greeter2.SayHello:output_type -> helloworld.HelloReply2
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_grpcContracts_helloworld_helloworld_proto_init() }
func file_internal_grpcContracts_helloworld_helloworld_proto_init() {
	if File_internal_grpcContracts_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply2); i {
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
			RawDescriptor: file_internal_grpcContracts_helloworld_helloworld_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_internal_grpcContracts_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_internal_grpcContracts_helloworld_helloworld_proto_depIdxs,
		EnumInfos:         file_internal_grpcContracts_helloworld_helloworld_proto_enumTypes,
		MessageInfos:      file_internal_grpcContracts_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_internal_grpcContracts_helloworld_helloworld_proto = out.File
	file_internal_grpcContracts_helloworld_helloworld_proto_rawDesc = nil
	file_internal_grpcContracts_helloworld_helloworld_proto_goTypes = nil
	file_internal_grpcContracts_helloworld_helloworld_proto_depIdxs = nil
}
