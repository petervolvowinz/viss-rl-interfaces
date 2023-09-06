// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: network_api.proto

package base

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

/// Parameters for subscription.
type SubscriberConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Identifier for subscription, will not receive subscriptions.
	ClientId *ClientId `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	/// List of signals in subscription.
	Signals *SignalIds `protobuf:"bytes,2,opt,name=signals,proto3" json:"signals,omitempty"`
	/// Only trigger callback when value changes when set to true.
	OnChange bool `protobuf:"varint,3,opt,name=onChange,proto3" json:"onChange,omitempty"`
}

func (x *SubscriberConfig) Reset() {
	*x = SubscriberConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriberConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberConfig) ProtoMessage() {}

func (x *SubscriberConfig) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberConfig.ProtoReflect.Descriptor instead.
func (*SubscriberConfig) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriberConfig) GetClientId() *ClientId {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *SubscriberConfig) GetSignals() *SignalIds {
	if x != nil {
		return x.Signals
	}
	return nil
}

func (x *SubscriberConfig) GetOnChange() bool {
	if x != nil {
		return x.OnChange
	}
	return false
}

type SubscriberWithMappingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Identifier for subscription, will not receive subscriptions.
	ClientId *ClientId `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	/// Custom Lua mapping code
	MappingCode string `protobuf:"bytes,2,opt,name=mappingCode,proto3" json:"mappingCode,omitempty"`
	/// Only trigger callback when value changes when set to true.
	OnChange bool `protobuf:"varint,3,opt,name=onChange,proto3" json:"onChange,omitempty"`
}

func (x *SubscriberWithMappingConfig) Reset() {
	*x = SubscriberWithMappingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriberWithMappingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberWithMappingConfig) ProtoMessage() {}

func (x *SubscriberWithMappingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberWithMappingConfig.ProtoReflect.Descriptor instead.
func (*SubscriberWithMappingConfig) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriberWithMappingConfig) GetClientId() *ClientId {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *SubscriberWithMappingConfig) GetMappingCode() string {
	if x != nil {
		return x.MappingCode
	}
	return ""
}

func (x *SubscriberWithMappingConfig) GetOnChange() bool {
	if x != nil {
		return x.OnChange
	}
	return false
}

/// List of signal identifiers
type SignalIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalId []*SignalId `protobuf:"bytes,1,rep,name=signalId,proto3" json:"signalId,omitempty"` /// List of signal identifiers.
}

func (x *SignalIds) Reset() {
	*x = SignalIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalIds) ProtoMessage() {}

func (x *SignalIds) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalIds.ProtoReflect.Descriptor instead.
func (*SignalIds) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{2}
}

func (x *SignalIds) GetSignalId() []*SignalId {
	if x != nil {
		return x.SignalId
	}
	return nil
}

/// Signals with values
type Signals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signal []*Signal `protobuf:"bytes,1,rep,name=signal,proto3" json:"signal,omitempty"` /// List of signals with values.
}

func (x *Signals) Reset() {
	*x = Signals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signals) ProtoMessage() {}

func (x *Signals) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signals.ProtoReflect.Descriptor instead.
func (*Signals) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{3}
}

func (x *Signals) GetSignal() []*Signal {
	if x != nil {
		return x.Signal
	}
	return nil
}

/// Parameters for publishing values.
type PublisherConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signals *Signals `protobuf:"bytes,1,opt,name=signals,proto3" json:"signals,omitempty"` /// Signals with values.
	//*
	// Identifier of publisher, typically your app identifier.
	// Subscribers with same identifier will not trigger callback.
	ClientId *ClientId `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	/// Specify frequency of publication. Specify 0 to only publish once.
	Frequency int32 `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *PublisherConfig) Reset() {
	*x = PublisherConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherConfig) ProtoMessage() {}

func (x *PublisherConfig) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherConfig.ProtoReflect.Descriptor instead.
func (*PublisherConfig) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{4}
}

func (x *PublisherConfig) GetSignals() *Signals {
	if x != nil {
		return x.Signals
	}
	return nil
}

func (x *PublisherConfig) GetClientId() *ClientId {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *PublisherConfig) GetFrequency() int32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

/// Signal with value.
type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *SignalId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` /// Identifier of signal to publish.
	/// Value signal. Union of supported data types.
	//
	// Types that are assignable to Payload:
	//	*Signal_Integer
	//	*Signal_Double
	//	*Signal_Arbitration
	//	*Signal_Empty
	Payload   isSignal_Payload `protobuf_oneof:"payload"`
	Raw       []byte           `protobuf:"bytes,5,opt,name=raw,proto3" json:"raw,omitempty"`              /// Binary data of value, relevant for Frames.
	Timestamp int64            `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"` /// time stamp in micro seconds, set when first seen.
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_network_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_network_api_proto_rawDescGZIP(), []int{5}
}

func (x *Signal) GetId() *SignalId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (m *Signal) GetPayload() isSignal_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Signal) GetInteger() int64 {
	if x, ok := x.GetPayload().(*Signal_Integer); ok {
		return x.Integer
	}
	return 0
}

func (x *Signal) GetDouble() float64 {
	if x, ok := x.GetPayload().(*Signal_Double); ok {
		return x.Double
	}
	return 0
}

func (x *Signal) GetArbitration() bool {
	if x, ok := x.GetPayload().(*Signal_Arbitration); ok {
		return x.Arbitration
	}
	return false
}

func (x *Signal) GetEmpty() bool {
	if x, ok := x.GetPayload().(*Signal_Empty); ok {
		return x.Empty
	}
	return false
}

func (x *Signal) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *Signal) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type isSignal_Payload interface {
	isSignal_Payload()
}

type Signal_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,proto3,oneof"` /// Value with integer type (int64).
}

type Signal_Double struct {
	Double float64 `protobuf:"fixed64,3,opt,name=double,proto3,oneof"` /// Floating point data type (double).
}

type Signal_Arbitration struct {
	Arbitration bool `protobuf:"varint,4,opt,name=arbitration,proto3,oneof"` /// Exclusive for LIN bus, present as true when a header is presented on the bus.
}

type Signal_Empty struct {
	Empty bool `protobuf:"varint,6,opt,name=empty,proto3,oneof"` /// No data supplied.
}

func (*Signal_Integer) isSignal_Payload() {}

func (*Signal_Double) isSignal_Payload() {}

func (*Signal_Arbitration) isSignal_Payload() {}

func (*Signal_Empty) isSignal_Payload() {}

var File_network_api_proto protoreflect.FileDescriptor

var file_network_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22,
	0x87, 0x01, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73,
	0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xd5, 0x01, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x61,
	0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0b, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x32, 0x91, 0x02, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x12,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x73, 0x12, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x73, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x62, 0x61, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_api_proto_rawDescOnce sync.Once
	file_network_api_proto_rawDescData = file_network_api_proto_rawDesc
)

func file_network_api_proto_rawDescGZIP() []byte {
	file_network_api_proto_rawDescOnce.Do(func() {
		file_network_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_api_proto_rawDescData)
	})
	return file_network_api_proto_rawDescData
}

var file_network_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_network_api_proto_goTypes = []interface{}{
	(*SubscriberConfig)(nil),            // 0: base.SubscriberConfig
	(*SubscriberWithMappingConfig)(nil), // 1: base.SubscriberWithMappingConfig
	(*SignalIds)(nil),                   // 2: base.SignalIds
	(*Signals)(nil),                     // 3: base.Signals
	(*PublisherConfig)(nil),             // 4: base.PublisherConfig
	(*Signal)(nil),                      // 5: base.Signal
	(*ClientId)(nil),                    // 6: base.ClientId
	(*SignalId)(nil),                    // 7: base.SignalId
	(*Empty)(nil),                       // 8: base.Empty
}
var file_network_api_proto_depIdxs = []int32{
	6,  // 0: base.SubscriberConfig.clientId:type_name -> base.ClientId
	2,  // 1: base.SubscriberConfig.signals:type_name -> base.SignalIds
	6,  // 2: base.SubscriberWithMappingConfig.clientId:type_name -> base.ClientId
	7,  // 3: base.SignalIds.signalId:type_name -> base.SignalId
	5,  // 4: base.Signals.signal:type_name -> base.Signal
	3,  // 5: base.PublisherConfig.signals:type_name -> base.Signals
	6,  // 6: base.PublisherConfig.clientId:type_name -> base.ClientId
	7,  // 7: base.Signal.id:type_name -> base.SignalId
	0,  // 8: base.NetworkService.SubscribeToSignals:input_type -> base.SubscriberConfig
	1,  // 9: base.NetworkService.SubscribeToSignalsWithMapping:input_type -> base.SubscriberWithMappingConfig
	4,  // 10: base.NetworkService.PublishSignals:input_type -> base.PublisherConfig
	2,  // 11: base.NetworkService.ReadSignals:input_type -> base.SignalIds
	3,  // 12: base.NetworkService.SubscribeToSignals:output_type -> base.Signals
	3,  // 13: base.NetworkService.SubscribeToSignalsWithMapping:output_type -> base.Signals
	8,  // 14: base.NetworkService.PublishSignals:output_type -> base.Empty
	3,  // 15: base.NetworkService.ReadSignals:output_type -> base.Signals
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_network_api_proto_init() }
func file_network_api_proto_init() {
	if File_network_api_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_network_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriberConfig); i {
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
		file_network_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriberWithMappingConfig); i {
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
		file_network_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalIds); i {
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
		file_network_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signals); i {
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
		file_network_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherConfig); i {
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
		file_network_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
	file_network_api_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Signal_Integer)(nil),
		(*Signal_Double)(nil),
		(*Signal_Arbitration)(nil),
		(*Signal_Empty)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_api_proto_goTypes,
		DependencyIndexes: file_network_api_proto_depIdxs,
		MessageInfos:      file_network_api_proto_msgTypes,
	}.Build()
	File_network_api_proto = out.File
	file_network_api_proto_rawDesc = nil
	file_network_api_proto_goTypes = nil
	file_network_api_proto_depIdxs = nil
}
