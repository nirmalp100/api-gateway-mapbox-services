// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: message.proto

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

type ReverseGeoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  string `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude string `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *ReverseGeoReq) Reset() {
	*x = ReverseGeoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseGeoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseGeoReq) ProtoMessage() {}

func (x *ReverseGeoReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseGeoReq.ProtoReflect.Descriptor instead.
func (*ReverseGeoReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReverseGeoReq) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *ReverseGeoReq) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type ReverseGeoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result `protobuf:"bytes,16,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ReverseGeoResp) Reset() {
	*x = ReverseGeoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseGeoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseGeoResp) ProtoMessage() {}

func (x *ReverseGeoResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseGeoResp.ProtoReflect.Descriptor instead.
func (*ReverseGeoResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ReverseGeoResp) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locality string `protobuf:"bytes,2,opt,name=locality,proto3" json:"locality,omitempty"`
	Place    string `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
	Dsitrict string `protobuf:"bytes,4,opt,name=dsitrict,proto3" json:"dsitrict,omitempty"`
	State    string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Country  string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *Result) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Result) GetDsitrict() string {
	if x != nil {
		return x.Dsitrict
	}
	return ""
}

func (x *Result) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Result) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type ForwardGeoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place string `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *ForwardGeoReq) Reset() {
	*x = ForwardGeoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardGeoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardGeoReq) ProtoMessage() {}

func (x *ForwardGeoReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardGeoReq.ProtoReflect.Descriptor instead.
func (*ForwardGeoReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *ForwardGeoReq) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

type ForwardGeoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *ForwardGeoResp) Reset() {
	*x = ForwardGeoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardGeoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardGeoResp) ProtoMessage() {}

func (x *ForwardGeoResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardGeoResp.ProtoReflect.Descriptor instead.
func (*ForwardGeoResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *ForwardGeoResp) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ForwardGeoResp) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type SuggestApiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place string `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *SuggestApiReq) Reset() {
	*x = SuggestApiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestApiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestApiReq) ProtoMessage() {}

func (x *SuggestApiReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestApiReq.ProtoReflect.Descriptor instead.
func (*SuggestApiReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *SuggestApiReq) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

type SuggestApiResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggestion []*Suggestions `protobuf:"bytes,1,rep,name=suggestion,proto3" json:"suggestion,omitempty"`
}

func (x *SuggestApiResp) Reset() {
	*x = SuggestApiResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestApiResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestApiResp) ProtoMessage() {}

func (x *SuggestApiResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestApiResp.ProtoReflect.Descriptor instead.
func (*SuggestApiResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *SuggestApiResp) GetSuggestion() []*Suggestions {
	if x != nil {
		return x.Suggestion
	}
	return nil
}

type Suggestions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Details  string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ActionID string `protobuf:"bytes,4,opt,name=actionID,proto3" json:"actionID,omitempty"`
}

func (x *Suggestions) Reset() {
	*x = Suggestions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggestions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestions) ProtoMessage() {}

func (x *Suggestions) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestions.ProtoReflect.Descriptor instead.
func (*Suggestions) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *Suggestions) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Suggestions) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Suggestions) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Suggestions) GetActionID() string {
	if x != nil {
		return x.ActionID
	}
	return ""
}

type RetrieveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveReq) Reset() {
	*x = RetrieveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReq) ProtoMessage() {}

func (x *RetrieveReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveReq.ProtoReflect.Descriptor instead.
func (*RetrieveReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RetrieveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Details  string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RetrieveResp) Reset() {
	*x = RetrieveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResp) ProtoMessage() {}

func (x *RetrieveResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResp.ProtoReflect.Descriptor instead.
func (*RetrieveResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *RetrieveResp) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RetrieveResp) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *RetrieveResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x47, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x36, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x47, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x73, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x73, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x25, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x47, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x43, 0x0a, 0x0e, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a,
	0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x79, 0x0a, 0x0b, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x0b, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x0c, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xe8, 0x01, 0x0a, 0x07, 0x47,
	0x65, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x47, 0x65, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x37, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x47, 0x65, 0x6f, 0x12, 0x13, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x47, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x11,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_message_proto_goTypes = []interface{}{
	(*ReverseGeoReq)(nil),  // 0: main.ReverseGeoReq
	(*ReverseGeoResp)(nil), // 1: main.ReverseGeoResp
	(*Result)(nil),         // 2: main.Result
	(*ForwardGeoReq)(nil),  // 3: main.ForwardGeoReq
	(*ForwardGeoResp)(nil), // 4: main.ForwardGeoResp
	(*SuggestApiReq)(nil),  // 5: main.SuggestApiReq
	(*SuggestApiResp)(nil), // 6: main.SuggestApiResp
	(*Suggestions)(nil),    // 7: main.Suggestions
	(*RetrieveReq)(nil),    // 8: main.RetrieveReq
	(*RetrieveResp)(nil),   // 9: main.RetrieveResp
}
var file_message_proto_depIdxs = []int32{
	2, // 0: main.ReverseGeoResp.result:type_name -> main.Result
	7, // 1: main.SuggestApiResp.suggestion:type_name -> main.Suggestions
	0, // 2: main.GeoCode.ReverseGeo:input_type -> main.ReverseGeoReq
	3, // 3: main.GeoCode.ForwardGeo:input_type -> main.ForwardGeoReq
	5, // 4: main.GeoCode.Suggestions:input_type -> main.SuggestApiReq
	8, // 5: main.GeoCode.Retrieve:input_type -> main.RetrieveReq
	1, // 6: main.GeoCode.ReverseGeo:output_type -> main.ReverseGeoResp
	4, // 7: main.GeoCode.ForwardGeo:output_type -> main.ForwardGeoResp
	6, // 8: main.GeoCode.Suggestions:output_type -> main.SuggestApiResp
	9, // 9: main.GeoCode.Retrieve:output_type -> main.RetrieveResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseGeoReq); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseGeoResp); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardGeoReq); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardGeoResp); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestApiReq); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestApiResp); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggestions); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveReq); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResp); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}