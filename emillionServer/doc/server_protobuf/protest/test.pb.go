// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package protest is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package protest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Reps             []int64 `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "protest.Test")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x51, 0xa9, 0xc5, 0x25, 0x4a, 0x2e, 0x5c,
	0x2c, 0x21, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x89, 0x49, 0xa9, 0x39, 0x12, 0x8c,
	0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x0d, 0x12, 0x2b, 0x4a, 0x2d, 0x28, 0x96, 0x60, 0x56,
	0x60, 0xd6, 0x60, 0x0e, 0x02, 0xb3, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x14, 0xe4, 0x50,
	0x5b, 0x00, 0x00, 0x00,
}