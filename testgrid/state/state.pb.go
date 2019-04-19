/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package state

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Row_Result int32

const (
	Row_NO_RESULT        Row_Result = 0
	Row_PASS             Row_Result = 1
	Row_PASS_WITH_ERRORS Row_Result = 2
	Row_PASS_WITH_SKIPS  Row_Result = 3
	Row_RUNNING          Row_Result = 4
	Row_FAIL             Row_Result = 12
	Row_FLAKY            Row_Result = 13
)

var Row_Result_name = map[int32]string{
	0:  "NO_RESULT",
	1:  "PASS",
	2:  "PASS_WITH_ERRORS",
	3:  "PASS_WITH_SKIPS",
	4:  "RUNNING",
	12: "FAIL",
	13: "FLAKY",
}

var Row_Result_value = map[string]int32{
	"NO_RESULT":        0,
	"PASS":             1,
	"PASS_WITH_ERRORS": 2,
	"PASS_WITH_SKIPS":  3,
	"RUNNING":          4,
	"FAIL":             12,
	"FLAKY":            13,
}

func (x Row_Result) String() string {
	return proto.EnumName(Row_Result_name, int32(x))
}

func (Row_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{2, 0}
}

type Column struct {
	Build                string   `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Started              float64  `protobuf:"fixed64,3,opt,name=started,proto3" json:"started,omitempty"`
	Extra                []string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{0}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *Column) GetStarted() float64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Column) GetExtra() []string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type Metric struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Sparse encoding of values. So given:
	//   Indices: [0, 2, 6, 4]
	//   Values: [0.1,0.2,6.1,6.2,6.3,6.4]
	// Decoded 12-value equivalent is:
	// [0.1, 0.2, nil, nil, nil, nil, nil, 6.1, 6.2, 6.3, 6.4, nil, nil, ...]
	Indices              []int32   `protobuf:"varint,2,rep,packed,name=indices,proto3" json:"indices,omitempty"`
	Values               []float64 `protobuf:"fixed64,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{1}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *Metric) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Row struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Results for this row, run-length encoded to reduce size/improve performance.
	// Thus (encoded -> decoded equivalent):
	//   [0, 3, 5, 4] -> [0, 0, 0, 5, 5, 5, 5]
	//   [5, 1] -> [5]
	//   [1, 5] -> [1, 1, 1, 1, 1]
	// The decoded values are Result enums
	Results              []int32   `protobuf:"varint,3,rep,packed,name=results,proto3" json:"results,omitempty"`
	CellIds              []string  `protobuf:"bytes,4,rep,name=cell_ids,json=cellIds,proto3" json:"cell_ids,omitempty"`
	Messages             []string  `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	Metrics              []*Metric `protobuf:"bytes,8,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Icons                []string  `protobuf:"bytes,9,rep,name=icons,proto3" json:"icons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{2}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Row) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Row) GetResults() []int32 {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Row) GetCellIds() []string {
	if m != nil {
		return m.CellIds
	}
	return nil
}

func (m *Row) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Row) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Row) GetIcons() []string {
	if m != nil {
		return m.Icons
	}
	return nil
}

type Grid struct {
	Columns              []*Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*Row    `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Grid) Reset()         { *m = Grid{} }
func (m *Grid) String() string { return proto.CompactTextString(m) }
func (*Grid) ProtoMessage()    {}
func (*Grid) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{3}
}

func (m *Grid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grid.Unmarshal(m, b)
}
func (m *Grid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grid.Marshal(b, m, deterministic)
}
func (m *Grid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grid.Merge(m, src)
}
func (m *Grid) XXX_Size() int {
	return xxx_messageInfo_Grid.Size(m)
}
func (m *Grid) XXX_DiscardUnknown() {
	xxx_messageInfo_Grid.DiscardUnknown(m)
}

var xxx_messageInfo_Grid proto.InternalMessageInfo

func (m *Grid) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Grid) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterEnum("Row_Result", Row_Result_name, Row_Result_value)
	proto.RegisterType((*Column)(nil), "Column")
	proto.RegisterType((*Metric)(nil), "Metric")
	proto.RegisterType((*Row)(nil), "Row")
	proto.RegisterType((*Grid)(nil), "Grid")
}

func init() { proto.RegisterFile("state.proto", fileDescriptor_a888679467bb7853) }

var fileDescriptor_a888679467bb7853 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0xec, 0x24, 0xee, 0x84, 0x05, 0xcb, 0xac, 0x90, 0xe1, 0x14, 0x72, 0xca, 0xa9,
	0x87, 0xe5, 0x09, 0x2a, 0xb4, 0xbb, 0xa4, 0x5b, 0xb2, 0xab, 0x49, 0x57, 0x88, 0x53, 0x95, 0x26,
	0x16, 0xb2, 0x94, 0x36, 0x28, 0x76, 0x29, 0x2f, 0xc6, 0xfb, 0x21, 0x3b, 0x0d, 0x5c, 0xb8, 0xcd,
	0xf7, 0x67, 0xf4, 0x67, 0xfe, 0x19, 0x43, 0x6a, 0x6c, 0x63, 0xd5, 0xf2, 0xc7, 0x38, 0xd8, 0x21,
	0xdf, 0x42, 0xfc, 0x69, 0xe8, 0x4f, 0x87, 0xa3, 0xb8, 0x86, 0x68, 0x7f, 0xd2, 0x7d, 0x27, 0x83,
	0x2c, 0x28, 0x16, 0x38, 0x81, 0x90, 0x90, 0x18, 0xdb, 0x8c, 0x56, 0x75, 0x92, 0x64, 0x41, 0x11,
	0xe0, 0x8c, 0xae, 0x5f, 0xfd, 0xb2, 0x63, 0x23, 0x69, 0x46, 0x5c, 0xbf, 0x87, 0x35, 0x65, 0x21,
	0x27, 0x79, 0x05, 0xf1, 0x17, 0x65, 0x47, 0xdd, 0x0a, 0x01, 0xf4, 0xd8, 0x1c, 0xd4, 0xc5, 0xd4,
	0xd7, 0xce, 0x53, 0x1f, 0x3b, 0xdd, 0x2a, 0x23, 0xc3, 0x8c, 0x14, 0x11, 0xce, 0x28, 0xde, 0x42,
	0xfc, 0xb3, 0xe9, 0x4f, 0xca, 0x48, 0x92, 0x91, 0x22, 0xc0, 0x0b, 0xe5, 0xbf, 0x43, 0x20, 0x38,
	0x9c, 0xff, 0xeb, 0xf6, 0x0a, 0x42, 0xdd, 0xc9, 0xd0, 0x2b, 0xa1, 0xf6, 0x13, 0x8f, 0xca, 0x9c,
	0x7a, 0x3b, 0x99, 0x44, 0x38, 0xa3, 0x78, 0x07, 0xac, 0x55, 0x7d, 0xbf, 0xd3, 0x9d, 0xb9, 0x0c,
	0x9d, 0x38, 0x2e, 0x3b, 0x23, 0xde, 0x03, 0x3b, 0x28, 0x63, 0x9a, 0xef, 0xca, 0xc8, 0xc8, 0x7f,
	0xfa, 0xcb, 0xe2, 0x03, 0x24, 0x07, 0x1f, 0xc6, 0x48, 0x96, 0x91, 0x22, 0xbd, 0x49, 0x96, 0x53,
	0x38, 0x9c, 0x75, 0xb7, 0x0b, 0xdd, 0x0e, 0x47, 0x23, 0x17, 0xd3, 0x2e, 0x3c, 0xe4, 0x16, 0x62,
	0xf4, 0xbf, 0x16, 0x57, 0xb0, 0xa8, 0x1e, 0x77, 0x78, 0x5b, 0x3f, 0x6f, 0xb6, 0xfc, 0x85, 0x60,
	0x40, 0x9f, 0x56, 0x75, 0xcd, 0x03, 0x71, 0x0d, 0xdc, 0x55, 0xbb, 0xaf, 0xe5, 0xf6, 0xf3, 0xee,
	0x16, 0xf1, 0x11, 0x6b, 0x1e, 0x8a, 0x37, 0xf0, 0xfa, 0x9f, 0x5a, 0x3f, 0x94, 0x4f, 0x35, 0x27,
	0x22, 0x85, 0x04, 0x9f, 0xab, 0xaa, 0xac, 0xee, 0x39, 0x75, 0x0e, 0x77, 0xab, 0x72, 0xc3, 0x5f,
	0x8a, 0x05, 0x44, 0x77, 0x9b, 0xd5, 0xc3, 0x37, 0x7e, 0x95, 0x53, 0x16, 0xf1, 0x74, 0x4d, 0x59,
	0xcc, 0xd9, 0x9a, 0x32, 0xe0, 0x69, 0x5e, 0x02, 0xbd, 0x1f, 0x75, 0xe7, 0x22, 0xb4, 0xfe, 0xca,
	0x46, 0x06, 0x97, 0x08, 0xd3, 0xd5, 0x71, 0xd6, 0x85, 0x04, 0x3a, 0x0e, 0xe7, 0xe9, 0x22, 0xe9,
	0x0d, 0x5d, 0xe2, 0x70, 0x46, 0xaf, 0xac, 0x29, 0x23, 0x1c, 0xf6, 0xb1, 0x7f, 0x2f, 0x1f, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x92, 0x9d, 0x38, 0xff, 0x3e, 0x02, 0x00, 0x00,
}
