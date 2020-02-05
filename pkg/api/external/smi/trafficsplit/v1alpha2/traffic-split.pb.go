// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/external/smi/trafficsplit/v1alpha2/traffic-split.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TrafficSplit allows users to incrementally direct percentages of traffic
// between various services. It will be used by clients such as ingress
// controllers or service mesh sidecars to split the outgoing traffic to
// different destinations.
type TrafficSplit struct {
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by controllers during validation
	Status core.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec                 *TrafficSplitSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TrafficSplit) Reset()         { *m = TrafficSplit{} }
func (m *TrafficSplit) String() string { return proto.CompactTextString(m) }
func (*TrafficSplit) ProtoMessage()    {}
func (*TrafficSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{0}
}
func (m *TrafficSplit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplit.Unmarshal(m, b)
}
func (m *TrafficSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplit.Marshal(b, m, deterministic)
}
func (m *TrafficSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplit.Merge(m, src)
}
func (m *TrafficSplit) XXX_Size() int {
	return xxx_messageInfo_TrafficSplit.Size(m)
}
func (m *TrafficSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplit.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplit proto.InternalMessageInfo

func (m *TrafficSplit) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *TrafficSplit) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *TrafficSplit) GetSpec() *TrafficSplitSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// TrafficSplitSpec is the specification for a TrafficSplit
type TrafficSplitSpec struct {
	Service              string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Backends             []*TrafficSplitBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TrafficSplitSpec) Reset()         { *m = TrafficSplitSpec{} }
func (m *TrafficSplitSpec) String() string { return proto.CompactTextString(m) }
func (*TrafficSplitSpec) ProtoMessage()    {}
func (*TrafficSplitSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{1}
}
func (m *TrafficSplitSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplitSpec.Unmarshal(m, b)
}
func (m *TrafficSplitSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplitSpec.Marshal(b, m, deterministic)
}
func (m *TrafficSplitSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplitSpec.Merge(m, src)
}
func (m *TrafficSplitSpec) XXX_Size() int {
	return xxx_messageInfo_TrafficSplitSpec.Size(m)
}
func (m *TrafficSplitSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplitSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplitSpec proto.InternalMessageInfo

func (m *TrafficSplitSpec) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *TrafficSplitSpec) GetBackends() []*TrafficSplitBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// TrafficSplitBackend defines a backend
type TrafficSplitBackend struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Weight               int32    `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficSplitBackend) Reset()         { *m = TrafficSplitBackend{} }
func (m *TrafficSplitBackend) String() string { return proto.CompactTextString(m) }
func (*TrafficSplitBackend) ProtoMessage()    {}
func (*TrafficSplitBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{2}
}
func (m *TrafficSplitBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplitBackend.Unmarshal(m, b)
}
func (m *TrafficSplitBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplitBackend.Marshal(b, m, deterministic)
}
func (m *TrafficSplitBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplitBackend.Merge(m, src)
}
func (m *TrafficSplitBackend) XXX_Size() int {
	return xxx_messageInfo_TrafficSplitBackend.Size(m)
}
func (m *TrafficSplitBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplitBackend.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplitBackend proto.InternalMessageInfo

func (m *TrafficSplitBackend) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *TrafficSplitBackend) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*TrafficSplit)(nil), "smi.trafficsplit.v1alpha2.TrafficSplit")
	proto.RegisterType((*TrafficSplitSpec)(nil), "smi.trafficsplit.v1alpha2.TrafficSplitSpec")
	proto.RegisterType((*TrafficSplitBackend)(nil), "smi.trafficsplit.v1alpha2.TrafficSplitBackend")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/external/smi/trafficsplit/v1alpha2/traffic-split.proto", fileDescriptor_fb32cb495875581b)
}

var fileDescriptor_fb32cb495875581b = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x09, 0x13, 0xca, 0xe0, 0x19, 0xfe, 0xcc, 0x68, 0x94, 0xa9, 0xd0, 0x0c, 0x8a, 0x84,
	0x54, 0x09, 0xd5, 0x56, 0xcb, 0x06, 0x75, 0x83, 0x94, 0x4d, 0x25, 0x10, 0x9b, 0x14, 0x36, 0xb0,
	0x72, 0xdd, 0x9b, 0xc4, 0x34, 0xa9, 0xad, 0xd8, 0x2d, 0x59, 0xf7, 0x69, 0x78, 0x04, 0x1e, 0x81,
	0xa7, 0xe8, 0x82, 0x37, 0x28, 0x12, 0x0b, 0x76, 0x28, 0x4e, 0x52, 0x85, 0x9f, 0x32, 0xdd, 0xf9,
	0xde, 0x73, 0xbf, 0xe3, 0x63, 0xeb, 0xa2, 0x0f, 0xb1, 0x30, 0xc9, 0x72, 0x4a, 0xb8, 0xcc, 0xa8,
	0x96, 0xa9, 0xec, 0x0b, 0x49, 0x33, 0xd0, 0x49, 0x5f, 0xe5, 0xf2, 0x23, 0x70, 0xa3, 0x29, 0x53,
	0x82, 0x42, 0x61, 0x20, 0x5f, 0xb0, 0x94, 0xea, 0x4c, 0x50, 0x93, 0xb3, 0x28, 0x12, 0x5c, 0xab,
	0x54, 0x18, 0xba, 0x1a, 0xb0, 0x54, 0x25, 0x6c, 0xd8, 0x74, 0xfb, 0xb6, 0x4d, 0x54, 0x2e, 0x8d,
	0xc4, 0x17, 0x3a, 0x13, 0xa4, 0x3d, 0x4e, 0x9a, 0xf1, 0xee, 0x59, 0x2c, 0x63, 0x69, 0xa7, 0x68,
	0x79, 0xaa, 0x80, 0x2e, 0x86, 0xc2, 0x54, 0x4d, 0x28, 0x6a, 0x93, 0xee, 0xa5, 0x8d, 0x35, 0x17,
	0xc6, 0x26, 0x59, 0x0d, 0x68, 0x06, 0x86, 0xcd, 0x98, 0x61, 0xfb, 0xf4, 0xa6, 0xae, 0xf5, 0xc7,
	0x7f, 0xe9, 0x86, 0x99, 0xa5, 0xae, 0x54, 0xff, 0xa7, 0x83, 0x4e, 0xdf, 0x56, 0x09, 0x27, 0x65,
	0x42, 0xfc, 0x02, 0x1d, 0x37, 0x17, 0x78, 0xce, 0x13, 0xa7, 0x77, 0x32, 0x3c, 0x27, 0x5c, 0xe6,
	0x40, 0x4a, 0x1b, 0x22, 0x24, 0x79, 0x53, 0xab, 0x81, 0xfb, 0x75, 0x73, 0x75, 0x23, 0xdc, 0x4d,
	0xe3, 0x31, 0xea, 0x54, 0xd6, 0xde, 0x4d, 0xcb, 0x9d, 0xfd, 0xce, 0x4d, 0xac, 0x16, 0x5c, 0x94,
	0xd4, 0xf7, 0xcd, 0xd5, 0x43, 0x03, 0xda, 0xcc, 0x44, 0x14, 0x8d, 0x7c, 0x11, 0x2f, 0x64, 0x0e,
	0x7e, 0x58, 0xe3, 0xf8, 0x25, 0x72, 0xb5, 0x02, 0xee, 0x1d, 0x59, 0x9b, 0x67, 0x64, 0xef, 0x2f,
	0x92, 0x76, 0xf2, 0x89, 0x02, 0x1e, 0x5a, 0x70, 0xf4, 0x74, 0xbd, 0x75, 0xef, 0xa1, 0xd3, 0x36,
	0xb3, 0xde, 0xba, 0xf7, 0xf1, 0xdd, 0x76, 0x47, 0xfb, 0x05, 0x7a, 0xf0, 0xa7, 0x01, 0xf6, 0xd0,
	0x6d, 0x0d, 0xf9, 0x4a, 0x70, 0xb0, 0xaf, 0xbf, 0x13, 0x36, 0x25, 0x7e, 0x85, 0x8e, 0xa7, 0x8c,
	0xcf, 0x61, 0x31, 0x2b, 0x1f, 0x78, 0xd4, 0x3b, 0x19, 0x92, 0x03, 0x93, 0x05, 0x15, 0x16, 0xee,
	0x78, 0x7f, 0x8c, 0x1e, 0xfd, 0x63, 0xe0, 0x3f, 0x97, 0x9f, 0xa3, 0xce, 0x27, 0x10, 0x71, 0x62,
	0xec, 0xdf, 0xde, 0x0a, 0xeb, 0x2a, 0x78, 0xf7, 0xe5, 0x87, 0xeb, 0x7c, 0xfe, 0x76, 0xe9, 0xbc,
	0x7f, 0x7d, 0xed, 0x22, 0xab, 0x79, 0x7c, 0xe0, 0x32, 0x4f, 0x3b, 0x76, 0x39, 0x9e, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xcb, 0xff, 0xb7, 0xad, 0x1e, 0x03, 0x00, 0x00,
}

func (this *TrafficSplit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplit)
	if !ok {
		that2, ok := that.(TrafficSplit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficSplitSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplitSpec)
	if !ok {
		that2, ok := that.(TrafficSplitSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if len(this.Backends) != len(that1.Backends) {
		return false
	}
	for i := range this.Backends {
		if !this.Backends[i].Equal(that1.Backends[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficSplitBackend) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplitBackend)
	if !ok {
		that2, ok := that.(TrafficSplitBackend)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if this.Weight != that1.Weight {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}