// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: api.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VideoState int32

const (
	VideoState_PROCESSING VideoState = 0
	VideoState_READY      VideoState = 1
)

// Enum value maps for VideoState.
var (
	VideoState_name = map[int32]string{
		0: "PROCESSING",
		1: "READY",
	}
	VideoState_value = map[string]int32{
		"PROCESSING": 0,
		"READY":      1,
	}
)

func (x VideoState) Enum() *VideoState {
	p := new(VideoState)
	*p = x
	return p
}

func (x VideoState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (VideoState) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x VideoState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoState.Descriptor instead.
func (VideoState) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId         string     `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Title           string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AccountId       int64      `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	VideoState      VideoState `protobuf:"varint,4,opt,name=video_state,json=videoState,proto3,enum=VideoState" json:"video_state,omitempty"`
	Filepath        string     `protobuf:"bytes,5,opt,name=filepath,proto3" json:"filepath,omitempty"`
	CreateTimestamp int64      `protobuf:"varint,6,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"`
	UpdateTimestamp int64      `protobuf:"varint,7,opt,name=update_timestamp,json=updateTimestamp,proto3" json:"update_timestamp,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Video) GetVideoState() VideoState {
	if x != nil {
		return x.VideoState
	}
	return VideoState_PROCESSING
}

func (x *Video) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *Video) GetCreateTimestamp() int64 {
	if x != nil {
		return x.CreateTimestamp
	}
	return 0
}

func (x *Video) GetUpdateTimestamp() int64 {
	if x != nil {
		return x.UpdateTimestamp
	}
	return 0
}

type UploadVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UploadVideoRequest) Reset() {
	*x = UploadVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoRequest) ProtoMessage() {}

func (x *UploadVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *UploadVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UploadVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadVideoResponse) Reset() {
	*x = UploadVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoResponse) ProtoMessage() {}

func (x *UploadVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetVideoRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type GetVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoResponse) Reset() {
	*x = GetVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoResponse) ProtoMessage() {}

func (x *GetVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoResponse.ProtoReflect.Descriptor instead.
func (*GetVideoResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideoResponse) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x05,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x2a, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2a, 0x27, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x01, 0x32, 0x79, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x50, 0x49, 0x12, 0x3a, 0x0a,
	0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x13, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(VideoState)(0),             // 0: VideoState
	(*Video)(nil),               // 1: Video
	(*UploadVideoRequest)(nil),  // 2: UploadVideoRequest
	(*UploadVideoResponse)(nil), // 3: UploadVideoResponse
	(*GetVideoRequest)(nil),     // 4: GetVideoRequest
	(*GetVideoResponse)(nil),    // 5: GetVideoResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: Video.video_state:type_name -> VideoState
	1, // 1: GetVideoResponse.video:type_name -> Video
	2, // 2: VideoAPI.UploadVideo:input_type -> UploadVideoRequest
	4, // 3: VideoAPI.GetVideo:input_type -> GetVideoRequest
	3, // 4: VideoAPI.UploadVideo:output_type -> UploadVideoResponse
	5, // 5: VideoAPI.GetVideo:output_type -> GetVideoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoAPIClient is the client API for VideoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoAPIClient interface {
	UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error)
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error)
}

type videoAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoAPIClient(cc grpc.ClientConnInterface) VideoAPIClient {
	return &videoAPIClient{cc}
}

func (c *videoAPIClient) UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error) {
	out := new(UploadVideoResponse)
	err := c.cc.Invoke(ctx, "/VideoAPI/UploadVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoAPIClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error) {
	out := new(GetVideoResponse)
	err := c.cc.Invoke(ctx, "/VideoAPI/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoAPIServer is the server API for VideoAPI service.
type VideoAPIServer interface {
	UploadVideo(context.Context, *UploadVideoRequest) (*UploadVideoResponse, error)
	GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error)
}

// UnimplementedVideoAPIServer can be embedded to have forward compatible implementations.
type UnimplementedVideoAPIServer struct {
}

func (*UnimplementedVideoAPIServer) UploadVideo(context.Context, *UploadVideoRequest) (*UploadVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (*UnimplementedVideoAPIServer) GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}

func RegisterVideoAPIServer(s *grpc.Server, srv VideoAPIServer) {
	s.RegisterService(&_VideoAPI_serviceDesc, srv)
}

func _VideoAPI_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoAPIServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoAPI/UploadVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoAPIServer).UploadVideo(ctx, req.(*UploadVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoAPI_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoAPIServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoAPI/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoAPIServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VideoAPI",
	HandlerType: (*VideoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadVideo",
			Handler:    _VideoAPI_UploadVideo_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _VideoAPI_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
