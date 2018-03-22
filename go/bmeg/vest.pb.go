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
	Gene       *VestScore_GENE `protobuf:"bytes,4,opt,name=gene" json:"gene,omitempty"`
	ProteinAlt string          `protobuf:"bytes,5,opt,name=protein_alt,json=proteinAlt" json:"protein_alt,omitempty"`
	ProteinRef string          `protobuf:"bytes,6,opt,name=protein_ref,json=proteinRef" json:"protein_ref,omitempty"`
	Ref        string          `protobuf:"bytes,7,opt,name=ref" json:"ref,omitempty"`
	Score      float64         `protobuf:"fixed64,8,opt,name=score" json:"score,omitempty"`
	Start      int64           `protobuf:"zigzag64,9,opt,name=start" json:"start,omitempty"`
	Transcript string          `protobuf:"bytes,10,opt,name=transcript" json:"transcript,omitempty"`
}

func (m *VestScore) Reset()                    { *m = VestScore{} }
func (m *VestScore) String() string            { return proto.CompactTextString(m) }
func (*VestScore) ProtoMessage()               {}
func (*VestScore) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *VestScore) GetGene() *VestScore_GENE {
	if m != nil {
		return m.Gene
	}
	return nil
}

type VestScore_GENE struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *VestScore_GENE) Reset()                    { *m = VestScore_GENE{} }
func (m *VestScore_GENE) String() string            { return proto.CompactTextString(m) }
func (*VestScore_GENE) ProtoMessage()               {}
func (*VestScore_GENE) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

func init() {
	proto.RegisterType((*VestScore)(nil), "bmeg.VestScore")
	proto.RegisterType((*VestScore_GENE)(nil), "bmeg.VestScore.GENE")
}

func init() { proto.RegisterFile("bmeg/vest.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0xc3, 0x30,
	0x0c, 0x86, 0xd5, 0xae, 0x1b, 0xd4, 0x93, 0x00, 0x59, 0x3b, 0x44, 0x1c, 0x00, 0x71, 0x9a, 0x38,
	0x14, 0x09, 0x9e, 0x80, 0xc3, 0xc4, 0x09, 0x0e, 0x41, 0xe2, 0x3a, 0x79, 0xc5, 0x1b, 0x95, 0x68,
	0x32, 0xa5, 0x86, 0x27, 0xe5, 0x81, 0x88, 0x93, 0x8a, 0xc1, 0xed, 0xcf, 0x67, 0xc7, 0xff, 0x6f,
	0xc3, 0xe9, 0xa6, 0xe7, 0xdd, 0xed, 0x17, 0x0f, 0xd2, 0xec, 0x83, 0x17, 0x8f, 0x95, 0x82, 0xeb,
	0xef, 0x12, 0xea, 0xd7, 0x08, 0x5f, 0x5a, 0x1f, 0x18, 0x2f, 0x61, 0x4e, 0xb4, 0xee, 0x3f, 0x85,
	0xa4, 0xf3, 0xce, 0x14, 0x57, 0xc5, 0xb2, 0xb6, 0x40, 0xf4, 0x34, 0x12, 0x3c, 0x83, 0x09, 0x7d,
	0x88, 0x29, 0x53, 0x41, 0x25, 0x5e, 0x00, 0xb4, 0xef, 0xc1, 0xf7, 0x7e, 0xf0, 0x3d, 0x9b, 0x49,
	0xfe, 0x71, 0x20, 0xb8, 0x84, 0x6a, 0xc7, 0x8e, 0x4d, 0x15, 0x2b, 0xf3, 0xbb, 0x45, 0xa3, 0xae,
	0xcd, 0xaf, 0x63, 0xf3, 0xb8, 0x7a, 0x5e, 0xd9, 0xd4, 0xa1, 0xe6, 0x9a, 0x8c, 0x3b, 0xb7, 0x56,
	0x8f, 0x69, 0x1e, 0x35, 0xa2, 0x87, 0x68, 0xf5, 0xa7, 0x21, 0xf0, 0xd6, 0xcc, 0xfe, 0x35, 0x58,
	0xde, 0x6a, 0x3a, 0x2d, 0x1c, 0xe5, 0x74, 0x51, 0xe2, 0x02, 0xa6, 0x83, 0xfa, 0x98, 0xe3, 0xc8,
	0x0a, 0x9b, 0x1f, 0x89, 0x0a, 0x05, 0x31, 0x75, 0xa4, 0x68, 0xf3, 0x43, 0x37, 0x91, 0x40, 0x6e,
	0x68, 0x43, 0xb7, 0x17, 0x03, 0x79, 0xfa, 0x81, 0x9c, 0xdf, 0x40, 0xa5, 0x69, 0xf1, 0x04, 0xca,
	0xee, 0x6d, 0xbc, 0x4d, 0x54, 0x88, 0x50, 0x39, 0x8a, 0xbb, 0xe7, 0xa3, 0x24, 0xbd, 0x99, 0xa5,
	0x1b, 0xdf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x1a, 0x5a, 0x97, 0x76, 0x01, 0x00, 0x00,
}