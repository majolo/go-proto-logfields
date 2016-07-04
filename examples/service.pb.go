// Code generated by protoc-gen-go.
// source: examples/service.proto
// DO NOT EDIT!

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	examples/service.proto

It has these top-level messages:
	Note
	Request
	Response
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/improbable-io/go-proto-logfields"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Note struct {
	Author string   `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Text   []string `protobuf:"bytes,2,rep,name=text" json:"text,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Note *Note  `protobuf:"bytes,2,opt,name=note" json:"note,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

type Response struct {
	DidStuff    bool    `protobuf:"varint,1,opt,name=did_stuff,json=didStuff" json:"did_stuff,omitempty"`
	ChangedNote *Note   `protobuf:"bytes,2,opt,name=changed_note,json=changedNote" json:"changed_note,omitempty"`
	Notes       []*Note `protobuf:"bytes,3,rep,name=notes" json:"notes,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetChangedNote() *Note {
	if m != nil {
		return m.ChangedNote
	}
	return nil
}

func (m *Response) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Note)(nil), "example.Note")
	proto.RegisterType((*Request)(nil), "example.Request")
	proto.RegisterType((*Response)(nil), "example.Response")
}

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0x1a, 0xd2, 0xd4, 0x2d, 0x8b, 0x07, 0x14, 0x31, 0xa0, 0x52, 0x18, 0xca, 0x90,
	0x04, 0xc1, 0xc8, 0xc6, 0x8c, 0x3a, 0x98, 0x0f, 0xa8, 0x9c, 0xe4, 0x25, 0xb1, 0x94, 0xc4, 0x21,
	0x7e, 0x41, 0xdd, 0xf8, 0x01, 0xbe, 0x0c, 0x89, 0x7f, 0xc2, 0x76, 0x5c, 0x16, 0xd4, 0x25, 0x2f,
	0xba, 0xf7, 0xbc, 0x6b, 0x5f, 0x93, 0x4b, 0x38, 0xf0, 0xb6, 0x6f, 0x40, 0xa5, 0x0a, 0x86, 0x0f,
	0x91, 0x43, 0xd2, 0x0f, 0x12, 0x25, 0x9d, 0x3b, 0xfd, 0xea, 0xb9, 0x12, 0x58, 0x8f, 0x59, 0x92,
	0xcb, 0x36, 0x15, 0xad, 0x36, 0x33, 0x9e, 0x35, 0x10, 0x0b, 0x99, 0x56, 0x32, 0xb6, 0x6c, 0xdc,
	0xc8, 0xaa, 0x14, 0xd0, 0x14, 0x2a, 0xfd, 0xfb, 0x9b, 0x52, 0x36, 0xaf, 0xc4, 0xdf, 0x49, 0x04,
	0x7a, 0x47, 0x02, 0x3e, 0x62, 0x2d, 0x87, 0xc8, 0x5b, 0x7b, 0xdb, 0xc5, 0xcb, 0xea, 0xfb, 0xe7,
	0x33, 0x3c, 0x6a, 0xcc, 0x4d, 0x7a, 0x4d, 0x7c, 0x84, 0x03, 0x46, 0x67, 0xeb, 0x99, 0x66, 0x88,
	0x66, 0x02, 0xe2, 0x77, 0x7a, 0x9f, 0x59, 0x5d, 0xa7, 0xcd, 0x19, 0xbc, 0x8f, 0xa0, 0xd0, 0xa0,
	0x3d, 0xc7, 0xda, 0xc5, 0x39, 0xd4, 0x28, 0xcc, 0x7e, 0xe9, 0xcd, 0xb4, 0xa8, 0xa3, 0xbc, 0xed,
	0xf2, 0xf1, 0x22, 0x71, 0x6d, 0x92, 0x9d, 0x4d, 0x33, 0xd6, 0xe6, 0xcb, 0x23, 0x21, 0x03, 0xd5,
	0xcb, 0x4e, 0x01, 0xbd, 0x27, 0x8b, 0x42, 0x14, 0x7b, 0x85, 0x63, 0x59, 0xda, 0xd0, 0xf0, 0x78,
	0x47, 0x23, 0x0b, 0x64, 0xa1, 0x9e, 0x6f, 0xc6, 0xa5, 0x0f, 0x64, 0x95, 0xd7, 0xbc, 0xab, 0xa0,
	0xd8, 0x9f, 0x3e, 0x62, 0xe9, 0x10, 0xdb, 0xfe, 0x96, 0x9c, 0x1b, 0x52, 0x45, 0x33, 0x5d, 0xec,
	0x1f, 0x3a, 0x79, 0x59, 0x60, 0x5f, 0xec, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x33, 0x90,
	0xb2, 0x91, 0x01, 0x00, 0x00,
}
