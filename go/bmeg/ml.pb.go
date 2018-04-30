// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bmeg/ml.proto

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ModelPerformance struct {
	ModelId string           `protobuf:"bytes,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	Metrics []*MetricSummary `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty"`
	Cutoffs []*CutoffMetrics `protobuf:"bytes,3,rep,name=cutoffs" json:"cutoffs,omitempty"`
}

func (m *ModelPerformance) Reset()                    { *m = ModelPerformance{} }
func (m *ModelPerformance) String() string            { return proto.CompactTextString(m) }
func (*ModelPerformance) ProtoMessage()               {}
func (*ModelPerformance) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ModelPerformance) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *ModelPerformance) GetMetrics() []*MetricSummary {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ModelPerformance) GetCutoffs() []*CutoffMetrics {
	if m != nil {
		return m.Cutoffs
	}
	return nil
}

type Model struct {
	Id          string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Structure   *ModelStructure     `protobuf:"bytes,2,opt,name=structure" json:"structure,omitempty"`
	Performance []*ModelPerformance `protobuf:"bytes,3,rep,name=performance" json:"performance,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Model) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Model) GetStructure() *ModelStructure {
	if m != nil {
		return m.Structure
	}
	return nil
}

func (m *Model) GetPerformance() []*ModelPerformance {
	if m != nil {
		return m.Performance
	}
	return nil
}

type MetricSummary struct {
	Type  string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty"`
}

func (m *MetricSummary) Reset()                    { *m = MetricSummary{} }
func (m *MetricSummary) String() string            { return proto.CompactTextString(m) }
func (*MetricSummary) ProtoMessage()               {}
func (*MetricSummary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *MetricSummary) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetricSummary) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfusionMatrix struct {
	TruePositive  float32 `protobuf:"fixed32,1,opt,name=true_positive,json=truePositive" json:"true_positive,omitempty"`
	TrueNegative  float32 `protobuf:"fixed32,2,opt,name=true_negative,json=trueNegative" json:"true_negative,omitempty"`
	FalsePositive float32 `protobuf:"fixed32,3,opt,name=false_positive,json=falsePositive" json:"false_positive,omitempty"`
	FalseNegative float32 `protobuf:"fixed32,4,opt,name=false_negative,json=falseNegative" json:"false_negative,omitempty"`
}

func (m *ConfusionMatrix) Reset()                    { *m = ConfusionMatrix{} }
func (m *ConfusionMatrix) String() string            { return proto.CompactTextString(m) }
func (*ConfusionMatrix) ProtoMessage()               {}
func (*ConfusionMatrix) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ConfusionMatrix) GetTruePositive() float32 {
	if m != nil {
		return m.TruePositive
	}
	return 0
}

func (m *ConfusionMatrix) GetTrueNegative() float32 {
	if m != nil {
		return m.TrueNegative
	}
	return 0
}

func (m *ConfusionMatrix) GetFalsePositive() float32 {
	if m != nil {
		return m.FalsePositive
	}
	return 0
}

func (m *ConfusionMatrix) GetFalseNegative() float32 {
	if m != nil {
		return m.FalseNegative
	}
	return 0
}

type CutoffMetrics struct {
	Value float32          `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
	Truth *ConfusionMatrix `protobuf:"bytes,2,opt,name=truth" json:"truth,omitempty"`
}

func (m *CutoffMetrics) Reset()                    { *m = CutoffMetrics{} }
func (m *CutoffMetrics) String() string            { return proto.CompactTextString(m) }
func (*CutoffMetrics) ProtoMessage()               {}
func (*CutoffMetrics) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *CutoffMetrics) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CutoffMetrics) GetTruth() *ConfusionMatrix {
	if m != nil {
		return m.Truth
	}
	return nil
}

type ModelStructure struct {
	Components []*ModelComponent `protobuf:"bytes,1,rep,name=components" json:"components,omitempty"`
}

func (m *ModelStructure) Reset()                    { *m = ModelStructure{} }
func (m *ModelStructure) String() string            { return proto.CompactTextString(m) }
func (*ModelStructure) ProtoMessage()               {}
func (*ModelStructure) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *ModelStructure) GetComponents() []*ModelComponent {
	if m != nil {
		return m.Components
	}
	return nil
}

type ModelComponent struct {
	Coeff float32 `protobuf:"fixed32,1,opt,name=coeff" json:"coeff,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*ModelComponent_RandomForest
	//	*ModelComponent_LinearCoeff
	Data isModelComponent_Data `protobuf_oneof:"Data"`
}

func (m *ModelComponent) Reset()                    { *m = ModelComponent{} }
func (m *ModelComponent) String() string            { return proto.CompactTextString(m) }
func (*ModelComponent) ProtoMessage()               {}
func (*ModelComponent) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type isModelComponent_Data interface {
	isModelComponent_Data()
}

type ModelComponent_RandomForest struct {
	RandomForest *RandomForestData `protobuf:"bytes,2,opt,name=random_forest,json=randomForest,oneof"`
}
type ModelComponent_LinearCoeff struct {
	LinearCoeff *LinearCoeffData `protobuf:"bytes,3,opt,name=linear_coeff,json=linearCoeff,oneof"`
}

func (*ModelComponent_RandomForest) isModelComponent_Data() {}
func (*ModelComponent_LinearCoeff) isModelComponent_Data()  {}

func (m *ModelComponent) GetData() isModelComponent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ModelComponent) GetCoeff() float32 {
	if m != nil {
		return m.Coeff
	}
	return 0
}

func (m *ModelComponent) GetRandomForest() *RandomForestData {
	if x, ok := m.GetData().(*ModelComponent_RandomForest); ok {
		return x.RandomForest
	}
	return nil
}

func (m *ModelComponent) GetLinearCoeff() *LinearCoeffData {
	if x, ok := m.GetData().(*ModelComponent_LinearCoeff); ok {
		return x.LinearCoeff
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModelComponent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModelComponent_OneofMarshaler, _ModelComponent_OneofUnmarshaler, _ModelComponent_OneofSizer, []interface{}{
		(*ModelComponent_RandomForest)(nil),
		(*ModelComponent_LinearCoeff)(nil),
	}
}

func _ModelComponent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModelComponent)
	// Data
	switch x := m.Data.(type) {
	case *ModelComponent_RandomForest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RandomForest); err != nil {
			return err
		}
	case *ModelComponent_LinearCoeff:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinearCoeff); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ModelComponent.Data has unexpected type %T", x)
	}
	return nil
}

func _ModelComponent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModelComponent)
	switch tag {
	case 2: // Data.random_forest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RandomForestData)
		err := b.DecodeMessage(msg)
		m.Data = &ModelComponent_RandomForest{msg}
		return true, err
	case 3: // Data.linear_coeff
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinearCoeffData)
		err := b.DecodeMessage(msg)
		m.Data = &ModelComponent_LinearCoeff{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ModelComponent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModelComponent)
	// Data
	switch x := m.Data.(type) {
	case *ModelComponent_RandomForest:
		s := proto.Size(x.RandomForest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModelComponent_LinearCoeff:
		s := proto.Size(x.LinearCoeff)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DecisionTree struct {
	Nodes []*DecisionTree_DecisionNode `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *DecisionTree) Reset()                    { *m = DecisionTree{} }
func (m *DecisionTree) String() string            { return proto.CompactTextString(m) }
func (*DecisionTree) ProtoMessage()               {}
func (*DecisionTree) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *DecisionTree) GetNodes() []*DecisionTree_DecisionNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type DecisionTree_DecisionNode struct {
	AboveChild    int32   `protobuf:"varint,1,opt,name=above_child,json=aboveChild" json:"above_child,omitempty"`
	BelowChild    int32   `protobuf:"varint,2,opt,name=below_child,json=belowChild" json:"below_child,omitempty"`
	SplitVariable string  `protobuf:"bytes,3,opt,name=split_variable,json=splitVariable" json:"split_variable,omitempty"`
	SplitValue    float64 `protobuf:"fixed64,4,opt,name=split_value,json=splitValue" json:"split_value,omitempty"`
	LeafNode      bool    `protobuf:"varint,5,opt,name=leaf_node,json=leafNode" json:"leaf_node,omitempty"`
	Label         bool    `protobuf:"varint,6,opt,name=label" json:"label,omitempty"`
}

func (m *DecisionTree_DecisionNode) Reset()                    { *m = DecisionTree_DecisionNode{} }
func (m *DecisionTree_DecisionNode) String() string            { return proto.CompactTextString(m) }
func (*DecisionTree_DecisionNode) ProtoMessage()               {}
func (*DecisionTree_DecisionNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7, 0} }

func (m *DecisionTree_DecisionNode) GetAboveChild() int32 {
	if m != nil {
		return m.AboveChild
	}
	return 0
}

func (m *DecisionTree_DecisionNode) GetBelowChild() int32 {
	if m != nil {
		return m.BelowChild
	}
	return 0
}

func (m *DecisionTree_DecisionNode) GetSplitVariable() string {
	if m != nil {
		return m.SplitVariable
	}
	return ""
}

func (m *DecisionTree_DecisionNode) GetSplitValue() float64 {
	if m != nil {
		return m.SplitValue
	}
	return 0
}

func (m *DecisionTree_DecisionNode) GetLeafNode() bool {
	if m != nil {
		return m.LeafNode
	}
	return false
}

func (m *DecisionTree_DecisionNode) GetLabel() bool {
	if m != nil {
		return m.Label
	}
	return false
}

type RandomForestData struct {
	Forest []*DecisionTree `protobuf:"bytes,1,rep,name=forest" json:"forest,omitempty"`
}

func (m *RandomForestData) Reset()                    { *m = RandomForestData{} }
func (m *RandomForestData) String() string            { return proto.CompactTextString(m) }
func (*RandomForestData) ProtoMessage()               {}
func (*RandomForestData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *RandomForestData) GetForest() []*DecisionTree {
	if m != nil {
		return m.Forest
	}
	return nil
}

type LinearCoeffData struct {
	Intercept float32            `protobuf:"fixed32,1,opt,name=intercept" json:"intercept,omitempty"`
	Alpha     float32            `protobuf:"fixed32,2,opt,name=alpha" json:"alpha,omitempty"`
	Coeff     map[string]float64 `protobuf:"bytes,3,rep,name=coeff" json:"coeff,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *LinearCoeffData) Reset()                    { *m = LinearCoeffData{} }
func (m *LinearCoeffData) String() string            { return proto.CompactTextString(m) }
func (*LinearCoeffData) ProtoMessage()               {}
func (*LinearCoeffData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *LinearCoeffData) GetIntercept() float32 {
	if m != nil {
		return m.Intercept
	}
	return 0
}

func (m *LinearCoeffData) GetAlpha() float32 {
	if m != nil {
		return m.Alpha
	}
	return 0
}

func (m *LinearCoeffData) GetCoeff() map[string]float64 {
	if m != nil {
		return m.Coeff
	}
	return nil
}

type FeatureCoefficient struct {
	Feature string  `protobuf:"bytes,1,opt,name=feature" json:"feature,omitempty"`
	Coeff   float32 `protobuf:"fixed32,2,opt,name=coeff" json:"coeff,omitempty"`
}

func (m *FeatureCoefficient) Reset()                    { *m = FeatureCoefficient{} }
func (m *FeatureCoefficient) String() string            { return proto.CompactTextString(m) }
func (*FeatureCoefficient) ProtoMessage()               {}
func (*FeatureCoefficient) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *FeatureCoefficient) GetFeature() string {
	if m != nil {
		return m.Feature
	}
	return ""
}

func (m *FeatureCoefficient) GetCoeff() float32 {
	if m != nil {
		return m.Coeff
	}
	return 0
}

func init() {
	proto.RegisterType((*ModelPerformance)(nil), "bmeg.ModelPerformance")
	proto.RegisterType((*Model)(nil), "bmeg.Model")
	proto.RegisterType((*MetricSummary)(nil), "bmeg.MetricSummary")
	proto.RegisterType((*ConfusionMatrix)(nil), "bmeg.ConfusionMatrix")
	proto.RegisterType((*CutoffMetrics)(nil), "bmeg.CutoffMetrics")
	proto.RegisterType((*ModelStructure)(nil), "bmeg.ModelStructure")
	proto.RegisterType((*ModelComponent)(nil), "bmeg.ModelComponent")
	proto.RegisterType((*DecisionTree)(nil), "bmeg.DecisionTree")
	proto.RegisterType((*DecisionTree_DecisionNode)(nil), "bmeg.DecisionTree.DecisionNode")
	proto.RegisterType((*RandomForestData)(nil), "bmeg.RandomForestData")
	proto.RegisterType((*LinearCoeffData)(nil), "bmeg.LinearCoeffData")
	proto.RegisterType((*FeatureCoefficient)(nil), "bmeg.FeatureCoefficient")
}

func init() { proto.RegisterFile("bmeg/ml.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0xb7, 0x9d, 0x43, 0x9b, 0x95, 0x43, 0xab, 0xd9, 0xdd, 0x5b, 0xa6, 0x20, 0x35, 0x32,
	0xaa, 0x14, 0x81, 0x08, 0x52, 0x39, 0xa8, 0x54, 0x82, 0x0b, 0x52, 0x2a, 0x90, 0x68, 0x55, 0x4d,
	0x11, 0xb7, 0xd1, 0xc4, 0x1e, 0xb7, 0x23, 0xc6, 0x1e, 0x6b, 0x3c, 0x0e, 0xe4, 0x01, 0xb8, 0xe2,
	0x92, 0xa7, 0xe0, 0x09, 0x78, 0x08, 0x5e, 0x0a, 0xcd, 0xc1, 0xb1, 0x53, 0xb8, 0xf3, 0xfa, 0xd7,
	0x37, 0xcb, 0xeb, 0x34, 0x03, 0xc3, 0x45, 0x4a, 0xaf, 0x1f, 0xa7, 0x7c, 0x9a, 0x4b, 0xa1, 0x04,
	0x6a, 0x6b, 0x33, 0xfc, 0xe6, 0xc1, 0xee, 0xb9, 0x88, 0x29, 0xbf, 0xa4, 0x32, 0x11, 0x32, 0x25,
	0x59, 0x44, 0xd1, 0x1d, 0xd8, 0x4e, 0xb5, 0x36, 0x67, 0x71, 0xe0, 0x8d, 0xbd, 0x49, 0x0f, 0x6f,
	0x19, 0xfb, 0x5d, 0x8c, 0x1e, 0xc1, 0x56, 0x4a, 0x95, 0x64, 0x51, 0x11, 0xf8, 0xe3, 0xd6, 0xa4,
	0x7f, 0xf4, 0xef, 0x54, 0xc7, 0x99, 0x9e, 0x1b, 0xf1, 0xaa, 0x4c, 0x53, 0x22, 0x57, 0xb8, 0x62,
	0x34, 0x1e, 0x95, 0x4a, 0x24, 0x49, 0x11, 0xb4, 0x9a, 0xf8, 0xcc, 0x88, 0xf6, 0x50, 0x81, 0x2b,
	0x26, 0xfc, 0xea, 0x41, 0xc7, 0x64, 0x83, 0x46, 0xe0, 0xaf, 0x7f, 0xee, 0xb3, 0x18, 0x1d, 0x41,
	0xaf, 0x50, 0xb2, 0x8c, 0x54, 0x29, 0x69, 0xe0, 0x8f, 0xbd, 0x49, 0xff, 0x68, 0xcf, 0xfd, 0x59,
	0xf3, 0x57, 0x95, 0x0f, 0xd7, 0x18, 0x3a, 0x86, 0x7e, 0x5e, 0x57, 0xe5, 0x12, 0xf8, 0xbf, 0x71,
	0xaa, 0x51, 0x33, 0x6e, 0xa2, 0xe1, 0x0b, 0x18, 0x6e, 0x14, 0x84, 0x10, 0xb4, 0xd5, 0x2a, 0xa7,
	0x2e, 0x21, 0xf3, 0x8d, 0xf6, 0xa0, 0xb3, 0x24, 0xbc, 0xb4, 0xe9, 0xf8, 0xd8, 0x1a, 0xe1, 0x0f,
	0x0f, 0x76, 0x66, 0x22, 0x4b, 0xca, 0x82, 0x89, 0xec, 0x9c, 0x28, 0xc9, 0xbe, 0xa0, 0xfb, 0x30,
	0x54, 0xb2, 0xa4, 0xf3, 0x5c, 0x14, 0x4c, 0xb1, 0xa5, 0x0d, 0xe3, 0xe3, 0x81, 0x16, 0x2f, 0x9d,
	0xb6, 0x86, 0x32, 0x7a, 0x4d, 0x0c, 0xe4, 0xd7, 0xd0, 0x85, 0xd3, 0xd0, 0x21, 0x8c, 0x12, 0xc2,
	0x8b, 0x46, 0xa8, 0x96, 0xa1, 0x86, 0x46, 0x5d, 0xc7, 0x5a, 0x63, 0xeb, 0x60, 0xed, 0x06, 0x56,
	0x45, 0x0b, 0x31, 0x0c, 0x37, 0x06, 0x51, 0x97, 0xe4, 0x35, 0x4a, 0x42, 0x0f, 0xa1, 0xa3, 0x64,
	0xa9, 0x6e, 0x5c, 0xdf, 0xff, 0x73, 0x23, 0xdc, 0x2c, 0x12, 0x5b, 0x26, 0x3c, 0x83, 0xd1, 0xe6,
	0x44, 0xd0, 0x53, 0x80, 0x48, 0xa4, 0xb9, 0xc8, 0x68, 0xa6, 0x8a, 0xc0, 0x33, 0x53, 0x68, 0xce,
	0x6e, 0x56, 0x39, 0x71, 0x83, 0xd3, 0x7d, 0x1c, 0x6d, 0xba, 0x75, 0x76, 0x91, 0xa0, 0x49, 0x52,
	0x65, 0x67, 0x0c, 0xf4, 0x12, 0x86, 0x92, 0x64, 0xb1, 0x48, 0xe7, 0x89, 0x90, 0xb4, 0x50, 0x2e,
	0x4b, 0x37, 0x67, 0x6c, 0x5c, 0x67, 0xc6, 0x73, 0x4a, 0x14, 0x79, 0xfb, 0x0f, 0x1e, 0xc8, 0x86,
	0x86, 0x4e, 0x60, 0xc0, 0x59, 0x46, 0x89, 0x9c, 0xdb, 0xd8, 0xad, 0x66, 0x8d, 0xef, 0x8d, 0x67,
	0xa6, 0x1d, 0xee, 0x70, 0x9f, 0xd7, 0xd2, 0xeb, 0x2e, 0xb4, 0xb5, 0x1c, 0x7e, 0xf7, 0x61, 0x70,
	0x4a, 0x23, 0xa6, 0xbb, 0xf1, 0x41, 0x52, 0x8a, 0x9e, 0x41, 0x27, 0x13, 0x31, 0xad, 0xaa, 0x3d,
	0xb0, 0xd1, 0x9a, 0xc8, 0xda, 0xb8, 0x10, 0x31, 0xc5, 0x96, 0xde, 0xff, 0xe5, 0xd5, 0x71, 0xb4,
	0x8e, 0x0e, 0xa0, 0x4f, 0x16, 0x62, 0x49, 0xe7, 0xd1, 0x0d, 0xe3, 0xf6, 0x3a, 0x74, 0x30, 0x18,
	0x69, 0xa6, 0x15, 0x0d, 0x2c, 0x28, 0x17, 0x9f, 0x1d, 0xe0, 0x5b, 0xc0, 0x48, 0x16, 0x38, 0x84,
	0x51, 0x91, 0x73, 0xa6, 0xe6, 0x4b, 0x22, 0x19, 0x59, 0x70, 0xbb, 0x30, 0x3d, 0x3c, 0x34, 0xea,
	0x47, 0x27, 0xea, 0x38, 0x15, 0xa6, 0xc7, 0xaf, 0xb7, 0xc5, 0xc3, 0xe0, 0x18, 0xbd, 0x03, 0x77,
	0xa1, 0xc7, 0x29, 0x49, 0xe6, 0x3a, 0xd1, 0xa0, 0x33, 0xf6, 0x26, 0xdb, 0x78, 0x5b, 0x0b, 0x26,
	0xcd, 0x3d, 0xe8, 0x70, 0xb2, 0xa0, 0x3c, 0xe8, 0x1a, 0x87, 0x35, 0xc2, 0x57, 0xb0, 0x7b, 0xbb,
	0xfb, 0xe8, 0x01, 0x74, 0xdd, 0x94, 0x6c, 0x67, 0xd0, 0x9f, 0x9d, 0xc1, 0x8e, 0x08, 0x7f, 0x7a,
	0xb0, 0x73, 0x6b, 0x00, 0xe8, 0x1e, 0xf4, 0x58, 0xa6, 0xa8, 0x8c, 0x68, 0xae, 0xdc, 0x1a, 0xd4,
	0x82, 0xce, 0x83, 0xf0, 0xfc, 0x86, 0x54, 0x37, 0xd2, 0x18, 0xe8, 0x79, 0xb5, 0x36, 0xf6, 0x01,
	0x18, 0xff, 0x75, 0xb4, 0x53, 0xf3, 0xf5, 0x26, 0x53, 0x72, 0xe5, 0x16, 0x6b, 0xff, 0x18, 0xa0,
	0x16, 0xd1, 0x2e, 0xb4, 0x3e, 0xd1, 0x95, 0x7b, 0x00, 0xf4, 0xe7, 0xe6, 0xfd, 0xf7, 0xdc, 0x65,
	0x39, 0xf1, 0x8f, 0xbd, 0xf0, 0x14, 0xd0, 0x19, 0x25, 0x7a, 0xf9, 0x4d, 0x00, 0x16, 0x31, 0xbd,
	0xbe, 0x01, 0x6c, 0x25, 0x56, 0xad, 0x1e, 0x55, 0x67, 0xd6, 0x8b, 0xed, 0x37, 0x16, 0x7b, 0xd1,
	0x35, 0xef, 0xf4, 0x93, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0x18, 0xbc, 0xf7, 0xb8, 0x05,
	0x00, 0x00,
}
