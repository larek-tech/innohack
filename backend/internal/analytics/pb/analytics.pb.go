// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: analytics/analytics.proto

package pb

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

type ChartType int32

const (
	ChartType_UNDEFINED  ChartType = 0
	ChartType_BAR_CHART  ChartType = 1
	ChartType_PIE_CHART  ChartType = 2
	ChartType_LINE_CHART ChartType = 3
)

// Enum value maps for ChartType.
var (
	ChartType_name = map[int32]string{
		0: "UNDEFINED",
		1: "BAR_CHART",
		2: "PIE_CHART",
		3: "LINE_CHART",
	}
	ChartType_value = map[string]int32{
		"UNDEFINED":  0,
		"BAR_CHART":  1,
		"PIE_CHART":  2,
		"LINE_CHART": 3,
	}
)

func (x ChartType) Enum() *ChartType {
	p := new(ChartType)
	*p = x
	return p
}

func (x ChartType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartType) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_analytics_proto_enumTypes[0].Descriptor()
}

func (ChartType) Type() protoreflect.EnumType {
	return &file_analytics_analytics_proto_enumTypes[0]
}

func (x ChartType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartType.Descriptor instead.
func (ChartType) EnumDescriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId   int64  `protobuf:"varint,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Prompt    string `protobuf:"bytes,4,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetQueryId() int64 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *Params) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Params) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Params) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

type DescriptionReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sources     []string `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Filenames   []string `protobuf:"bytes,2,rep,name=filenames,proto3" json:"filenames,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DescriptionReport) Reset() {
	*x = DescriptionReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptionReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptionReport) ProtoMessage() {}

func (x *DescriptionReport) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptionReport.ProtoReflect.Descriptor instead.
func (*DescriptionReport) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *DescriptionReport) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *DescriptionReport) GetFilenames() []string {
	if x != nil {
		return x.Filenames
	}
	return nil
}

func (x *DescriptionReport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ChartReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Charts      []*Chart      `protobuf:"bytes,1,rep,name=charts,proto3" json:"charts,omitempty"`
	Multipliers []*Multiplier `protobuf:"bytes,2,rep,name=multipliers,proto3" json:"multipliers,omitempty"`
	Description string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChartReport) Reset() {
	*x = ChartReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartReport) ProtoMessage() {}

func (x *ChartReport) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartReport.ProtoReflect.Descriptor instead.
func (*ChartReport) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *ChartReport) GetCharts() []*Chart {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *ChartReport) GetMultipliers() []*Multiplier {
	if x != nil {
		return x.Multipliers
	}
	return nil
}

func (x *ChartReport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type        ChartType `protobuf:"varint,2,opt,name=type,proto3,enum=analytics.ChartType" json:"type,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Records     []*Record `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *Chart) Reset() {
	*x = Chart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *Chart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chart) GetType() ChartType {
	if x != nil {
		return x.Type
	}
	return ChartType_UNDEFINED
}

func (x *Chart) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Chart) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Multiplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Multiplier) Reset() {
	*x = Multiplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Multiplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Multiplier) ProtoMessage() {}

func (x *Multiplier) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Multiplier.ProtoReflect.Descriptor instead.
func (*Multiplier) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{4}
}

func (x *Multiplier) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Multiplier) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string  `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{5}
}

func (x *Record) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *Record) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_analytics_analytics_proto protoreflect.FileDescriptor

var file_analytics_analytics_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x22, 0x75, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x6d, 0x0a,
	0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x06,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x96, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x24, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x2a, 0x48, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x41, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x10, 0x03,
	0x32, 0xd2, 0x01, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x11, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x1c, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_analytics_proto_rawDescOnce sync.Once
	file_analytics_analytics_proto_rawDescData = file_analytics_analytics_proto_rawDesc
)

func file_analytics_analytics_proto_rawDescGZIP() []byte {
	file_analytics_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_analytics_proto_rawDescData)
	})
	return file_analytics_analytics_proto_rawDescData
}

var file_analytics_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_analytics_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_analytics_analytics_proto_goTypes = []interface{}{
	(ChartType)(0),            // 0: analytics.ChartType
	(*Params)(nil),            // 1: analytics.Params
	(*DescriptionReport)(nil), // 2: analytics.DescriptionReport
	(*ChartReport)(nil),       // 3: analytics.ChartReport
	(*Chart)(nil),             // 4: analytics.Chart
	(*Multiplier)(nil),        // 5: analytics.Multiplier
	(*Record)(nil),            // 6: analytics.Record
}
var file_analytics_analytics_proto_depIdxs = []int32{
	4, // 0: analytics.ChartReport.charts:type_name -> analytics.Chart
	5, // 1: analytics.ChartReport.multipliers:type_name -> analytics.Multiplier
	0, // 2: analytics.Chart.type:type_name -> analytics.ChartType
	6, // 3: analytics.Chart.records:type_name -> analytics.Record
	1, // 4: analytics.Analytics.GetCharts:input_type -> analytics.Params
	1, // 5: analytics.Analytics.GetDescriptionStream:input_type -> analytics.Params
	1, // 6: analytics.Analytics.GetChartSummary:input_type -> analytics.Params
	3, // 7: analytics.Analytics.GetCharts:output_type -> analytics.ChartReport
	2, // 8: analytics.Analytics.GetDescriptionStream:output_type -> analytics.DescriptionReport
	3, // 9: analytics.Analytics.GetChartSummary:output_type -> analytics.ChartReport
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_analytics_analytics_proto_init() }
func file_analytics_analytics_proto_init() {
	if File_analytics_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_analytics_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptionReport); i {
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
		file_analytics_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartReport); i {
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
		file_analytics_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chart); i {
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
		file_analytics_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Multiplier); i {
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
		file_analytics_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_analytics_analytics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_analytics_proto_depIdxs,
		EnumInfos:         file_analytics_analytics_proto_enumTypes,
		MessageInfos:      file_analytics_analytics_proto_msgTypes,
	}.Build()
	File_analytics_analytics_proto = out.File
	file_analytics_analytics_proto_rawDesc = nil
	file_analytics_analytics_proto_goTypes = nil
	file_analytics_analytics_proto_depIdxs = nil
}
