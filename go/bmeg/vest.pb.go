// Code generated by protoc-gen-go.
// source: bmeg/vest.proto
// DO NOT EDIT!

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VestScore struct {
	AaMutation string          `protobuf:"bytes,1,opt,name=aa_mutation,json=aaMutation" json:"aa_mutation,omitempty"`
	Alt        string          `protobuf:"bytes,2,opt,name=alt" json:"alt,omitempty"`
	Chromosome string          `protobuf:"bytes,3,opt,name=chromosome" json:"chromosome,omitempty"`
	Gene       *VestScore_Gene `protobuf:"bytes,4,opt,name=gene" json:"gene,omitempty"`
	ProteinAlt string          `protobuf:"bytes,5,opt,name=protein_alt,json=proteinAlt" json:"protein_alt,omitempty"`
	ProteinRef string          `protobuf:"bytes,6,opt,name=protein_ref,json=proteinRef" json:"protein_ref,omitempty"`
	Ref        string          `protobuf:"bytes,7,opt,name=ref" json:"ref,omitempty"`
	Score      float64         `protobuf:"fixed64,8,opt,name=score" json:"score,omitempty"`
	Start      int64           `protobuf:"varint,9,opt,name=start" json:"start,omitempty"`
	Transcript string          `protobuf:"bytes,10,opt,name=transcript" json:"transcript,omitempty"`
}

func (m *VestScore) Reset()                    { *m = VestScore{} }
func (m *VestScore) String() string            { return proto.CompactTextString(m) }
func (*VestScore) ProtoMessage()               {}
func (*VestScore) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *VestScore) GetGene() *VestScore_Gene {
	if m != nil {
		return m.Gene
	}
	return nil
}

type VestScore_Gene struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *VestScore_Gene) Reset()                    { *m = VestScore_Gene{} }
func (m *VestScore_Gene) String() string            { return proto.CompactTextString(m) }
func (*VestScore_Gene) ProtoMessage()               {}
func (*VestScore_Gene) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

func init() {
	proto.RegisterType((*VestScore)(nil), "bmeg.VestScore")
	proto.RegisterType((*VestScore_Gene)(nil), "bmeg.VestScore.Gene")
}

func init() { proto.RegisterFile("bmeg/vest.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0x75, 0x3f, 0x09, 0xdc, 0x46, 0x02, 0x64, 0xa5, 0xb0, 0x28, 0x00, 0x51, 0x45, 0x14,
	0x87, 0x04, 0x4f, 0x40, 0x45, 0x45, 0x63, 0x24, 0xda, 0x68, 0x73, 0x6c, 0xc2, 0x49, 0x9c, 0x1d,
	0xf9, 0x16, 0x9e, 0x94, 0x07, 0x62, 0xd7, 0x3e, 0x11, 0xd2, 0x8d, 0xbf, 0x59, 0x7b, 0xc6, 0x0b,
	0xe7, 0x9b, 0x81, 0x76, 0xf7, 0xdf, 0x34, 0x72, 0xbb, 0x8f, 0x81, 0x83, 0xa9, 0x15, 0xdc, 0xfe,
	0x94, 0xd0, 0xbc, 0x09, 0x7c, 0xed, 0x42, 0x24, 0x73, 0x0d, 0x0b, 0xc4, 0xf5, 0xf0, 0xc5, 0xc8,
	0x7d, 0xf0, 0xb6, 0xb8, 0x29, 0x56, 0x8d, 0x03, 0xc4, 0x97, 0x89, 0x98, 0x0b, 0xa8, 0xf0, 0x93,
	0x6d, 0x99, 0x0c, 0x95, 0xe6, 0x0a, 0xa0, 0xfb, 0x88, 0x61, 0x08, 0x63, 0x18, 0xc8, 0x56, 0xf9,
	0xc6, 0x81, 0x98, 0x15, 0xd4, 0x3b, 0xf2, 0x64, 0x6b, 0x71, 0x16, 0x0f, 0xcb, 0x56, 0x53, 0xdb,
	0xbf, 0xc4, 0xf6, 0x59, 0x3c, 0x97, 0x26, 0x34, 0x5c, 0x9b, 0x51, 0xef, 0xd7, 0x9a, 0x31, 0xcb,
	0x4f, 0x4d, 0xe8, 0x49, 0xa2, 0xfe, 0x0d, 0x44, 0xda, 0xda, 0xf9, 0xd1, 0x80, 0xa3, 0xad, 0xb6,
	0x53, 0xe3, 0x24, 0xb7, 0x13, 0x69, 0x96, 0x30, 0x1b, 0x35, 0xc7, 0x9e, 0x0a, 0x2b, 0x5c, 0x3e,
	0x24, 0xca, 0x18, 0xd9, 0x36, 0x42, 0x2b, 0x97, 0x0f, 0xfa, 0x13, 0x8e, 0xe8, 0xc7, 0x2e, 0xf6,
	0x7b, 0xb6, 0x90, 0x5f, 0x3f, 0x90, 0xcb, 0x3b, 0xa8, 0xb5, 0xad, 0x39, 0x83, 0xb2, 0x7f, 0x9f,
	0x76, 0x23, 0xca, 0x18, 0xa8, 0x3d, 0xca, 0xdf, 0xf3, 0x52, 0x92, 0xde, 0xcc, 0xd3, 0x8e, 0x1f,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0xcf, 0x0f, 0x20, 0x76, 0x01, 0x00, 0x00,
}
