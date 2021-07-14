// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: carlibrary.proto

package carlibrary

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

//请求
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{0}
}

type CarSeriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarBand string `protobuf:"bytes,1,opt,name=car_band,json=carBand,proto3" json:"car_band,omitempty"`
}

func (x *CarSeriesRequest) Reset() {
	*x = CarSeriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarSeriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarSeriesRequest) ProtoMessage() {}

func (x *CarSeriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarSeriesRequest.ProtoReflect.Descriptor instead.
func (*CarSeriesRequest) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{1}
}

func (x *CarSeriesRequest) GetCarBand() string {
	if x != nil {
		return x.CarBand
	}
	return ""
}

type CarModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//品牌
	CarBand string `protobuf:"bytes,1,opt,name=car_band,json=carBand,proto3" json:"car_band,omitempty"`
	//车系
	CarSeries string `protobuf:"bytes,2,opt,name=car_series,json=carSeries,proto3" json:"car_series,omitempty"`
}

func (x *CarModelRequest) Reset() {
	*x = CarModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarModelRequest) ProtoMessage() {}

func (x *CarModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarModelRequest.ProtoReflect.Descriptor instead.
func (*CarModelRequest) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{2}
}

func (x *CarModelRequest) GetCarBand() string {
	if x != nil {
		return x.CarBand
	}
	return ""
}

func (x *CarModelRequest) GetCarSeries() string {
	if x != nil {
		return x.CarSeries
	}
	return ""
}

//响应
//车的品牌
type CarBand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//品牌
	CarBand string `protobuf:"bytes,1,opt,name=car_band,json=carBand,proto3" json:"car_band,omitempty"`
	//品牌logo
	CarBandLogo string `protobuf:"bytes,2,opt,name=car_band_logo,json=carBandLogo,proto3" json:"car_band_logo,omitempty"`
}

func (x *CarBand) Reset() {
	*x = CarBand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarBand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarBand) ProtoMessage() {}

func (x *CarBand) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarBand.ProtoReflect.Descriptor instead.
func (*CarBand) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{3}
}

func (x *CarBand) GetCarBand() string {
	if x != nil {
		return x.CarBand
	}
	return ""
}

func (x *CarBand) GetCarBandLogo() string {
	if x != nil {
		return x.CarBandLogo
	}
	return ""
}

//车系
type CarSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//车系
	CarSeries string `protobuf:"bytes,1,opt,name=car_series,json=carSeries,proto3" json:"car_series,omitempty"`
	//车系图片
	CarPicture string `protobuf:"bytes,2,opt,name=car_picture,json=carPicture,proto3" json:"car_picture,omitempty"`
	//车系评分
	CarPoint float32 `protobuf:"fixed32,3,opt,name=car_point,json=carPoint,proto3" json:"car_point,omitempty"`
	//车系参考价
	CarPrice string `protobuf:"bytes,4,opt,name=car_price,json=carPrice,proto3" json:"car_price,omitempty"`
}

func (x *CarSeries) Reset() {
	*x = CarSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarSeries) ProtoMessage() {}

func (x *CarSeries) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarSeries.ProtoReflect.Descriptor instead.
func (*CarSeries) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{4}
}

func (x *CarSeries) GetCarSeries() string {
	if x != nil {
		return x.CarSeries
	}
	return ""
}

func (x *CarSeries) GetCarPicture() string {
	if x != nil {
		return x.CarPicture
	}
	return ""
}

func (x *CarSeries) GetCarPoint() float32 {
	if x != nil {
		return x.CarPoint
	}
	return 0
}

func (x *CarSeries) GetCarPrice() string {
	if x != nil {
		return x.CarPrice
	}
	return ""
}

//车型
type CarModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarSeries *CarSeries `protobuf:"bytes,1,opt,name=car_series,json=carSeries,proto3" json:"car_series,omitempty"`
	//车型名
	CarModel string `protobuf:"bytes,2,opt,name=car_model,json=carModel,proto3" json:"car_model,omitempty"`
	//车型价格
	CarModelPrice string `protobuf:"bytes,3,opt,name=car_model_price,json=carModelPrice,proto3" json:"car_model_price,omitempty"`
}

func (x *CarModel) Reset() {
	*x = CarModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carlibrary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarModel) ProtoMessage() {}

func (x *CarModel) ProtoReflect() protoreflect.Message {
	mi := &file_carlibrary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarModel.ProtoReflect.Descriptor instead.
func (*CarModel) Descriptor() ([]byte, []int) {
	return file_carlibrary_proto_rawDescGZIP(), []int{5}
}

func (x *CarModel) GetCarSeries() *CarSeries {
	if x != nil {
		return x.CarSeries
	}
	return nil
}

func (x *CarModel) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *CarModel) GetCarModelPrice() string {
	if x != nil {
		return x.CarModelPrice
	}
	return ""
}

var File_carlibrary_proto protoreflect.FileDescriptor

var file_carlibrary_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x61, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x61, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x72,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72,
	0x42, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x61, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x61, 0x72,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x61, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x09, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x72, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x61, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x63, 0x61, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0xe2, 0x01,
	0x0a, 0x11, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x4c, 0x4c, 0x43, 0x61,
	0x72, 0x42, 0x61, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x30, 0x01, 0x12,
	0x49, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x43,
	0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0f, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x2e,
	0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x72,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x61, 0x72, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carlibrary_proto_rawDescOnce sync.Once
	file_carlibrary_proto_rawDescData = file_carlibrary_proto_rawDesc
)

func file_carlibrary_proto_rawDescGZIP() []byte {
	file_carlibrary_proto_rawDescOnce.Do(func() {
		file_carlibrary_proto_rawDescData = protoimpl.X.CompressGZIP(file_carlibrary_proto_rawDescData)
	})
	return file_carlibrary_proto_rawDescData
}

var file_carlibrary_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_carlibrary_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: carlibrary.Empty
	(*CarSeriesRequest)(nil), // 1: carlibrary.CarSeriesRequest
	(*CarModelRequest)(nil),  // 2: carlibrary.CarModelRequest
	(*CarBand)(nil),          // 3: carlibrary.CarBand
	(*CarSeries)(nil),        // 4: carlibrary.CarSeries
	(*CarModel)(nil),         // 5: carlibrary.CarModel
}
var file_carlibrary_proto_depIdxs = []int32{
	4, // 0: carlibrary.CarModel.car_series:type_name -> carlibrary.CarSeries
	0, // 1: carlibrary.CarLibraryService.FindALLCarBand:input_type -> carlibrary.Empty
	1, // 2: carlibrary.CarLibraryService.FindAllCarSeries:input_type -> carlibrary.CarSeriesRequest
	2, // 3: carlibrary.CarLibraryService.FindAllCarModel:input_type -> carlibrary.CarModelRequest
	3, // 4: carlibrary.CarLibraryService.FindALLCarBand:output_type -> carlibrary.CarBand
	4, // 5: carlibrary.CarLibraryService.FindAllCarSeries:output_type -> carlibrary.CarSeries
	5, // 6: carlibrary.CarLibraryService.FindAllCarModel:output_type -> carlibrary.CarModel
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_carlibrary_proto_init() }
func file_carlibrary_proto_init() {
	if File_carlibrary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carlibrary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_carlibrary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarSeriesRequest); i {
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
		file_carlibrary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarModelRequest); i {
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
		file_carlibrary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarBand); i {
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
		file_carlibrary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarSeries); i {
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
		file_carlibrary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarModel); i {
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
			RawDescriptor: file_carlibrary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carlibrary_proto_goTypes,
		DependencyIndexes: file_carlibrary_proto_depIdxs,
		MessageInfos:      file_carlibrary_proto_msgTypes,
	}.Build()
	File_carlibrary_proto = out.File
	file_carlibrary_proto_rawDesc = nil
	file_carlibrary_proto_goTypes = nil
	file_carlibrary_proto_depIdxs = nil
}