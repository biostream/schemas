// Code generated by protoc-gen-go.
// source: bmeg/clinical.proto
// DO NOT EDIT!

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An ontology term describing an attribute. (e.g. the phenotype attribute
// 'polydactyly' from HPO)
type OntologyTerm struct {
	// Ontology term identifier - the CURIE for an ontology term. It
	// differs from the standard GA4GH schema's :ref:`id <apidesign_object_ids>`
	// in that it is a CURIE pointing to an information resource outside of the
	// scope of the schema or its resource implementation.
	TermId string `protobuf:"bytes,1,opt,name=term_id,json=termId" json:"term_id,omitempty"`
	// Ontology term - the label of the ontology term the termId is pointing to.
	Term string `protobuf:"bytes,2,opt,name=term" json:"term,omitempty"`
}

func (m *OntologyTerm) Reset()                    { *m = OntologyTerm{} }
func (m *OntologyTerm) String() string            { return proto.CompactTextString(m) }
func (*OntologyTerm) ProtoMessage()               {}
func (*OntologyTerm) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Individual struct {
	// The Individual's :ref:`id <apidesign_object_ids>`. This is unique in the
	// context of the server instance.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The ID of the dataset this Individual belongs to.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	// The Individual's name. This is a label or symbolic identifier for the individual.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The Individual's description. This attribute contains human readable text.
	// The "description" attributes should not contain any structured data.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The :ref:`ISO 8601<metadata_date_time>` time at which this Individual's record
	// was created.
	Created string `protobuf:"bytes,5,opt,name=created" json:"created,omitempty"`
	// The :ref:`ISO 8601<metadata_date_time>` time at which this Individual record
	// was updated.
	Updated string `protobuf:"bytes,6,opt,name=updated" json:"updated,omitempty"`
	// For a representation of an NCBI Taxon ID as an OntologyTerm, see
	// NCBITaxon Ontology
	//   http://www.obofoundry.org/ontology/ncbitaxon.html
	// For example, 'Homo sapiens' has the ID 9606. The NCBITaxon ontology ID for
	// this is NCBITaxon:9606, and therefore:
	//   species : { term_id: "NCBITaxon:9606", term : "Homo sapiens" }
	Species *OntologyTerm `protobuf:"bytes,7,opt,name=species" json:"species,omitempty"`
	// The genetic sex of this individual.
	// Use `null` when unknown or not applicable.
	// Recommended: PATO http://purl.obolibrary.org/obo/PATO_0020000 ...
	// Example:
	//   sex : { term_id: "PATO:0020000", term : "female genetic sex" }
	Sex *OntologyTerm `protobuf:"bytes,8,opt,name=sex" json:"sex,omitempty"`
	// the source of the data
	Source string `protobuf:"bytes,9,opt,name=source" json:"source,omitempty"`
	// A map of additional information regarding the Individual.
	Attributes *google_protobuf.Struct `protobuf:"bytes,10,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *Individual) Reset()                    { *m = Individual{} }
func (m *Individual) String() string            { return proto.CompactTextString(m) }
func (*Individual) ProtoMessage()               {}
func (*Individual) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Individual) GetSpecies() *OntologyTerm {
	if m != nil {
		return m.Species
	}
	return nil
}

func (m *Individual) GetSex() *OntologyTerm {
	if m != nil {
		return m.Sex
	}
	return nil
}

func (m *Individual) GetAttributes() *google_protobuf.Struct {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Biosample struct {
	// The Biosample :ref:`id <apidesign_object_ids>`. This is unique in the
	// context of the server instance.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The ID of the dataset this Biosample belongs to.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	// The Biosample's name This is a label or symbolic identifier for the biosample.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The biosample's description. This attribute contains human readable text.
	// The "description" attributes should not contain any structured data.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// OntologyTerm describing the primary disease associated with this Biosample.
	Disease *OntologyTerm `protobuf:"bytes,5,opt,name=disease" json:"disease,omitempty"`
	// The :ref:`ISO 8601<metadata_date_time>` time at which this Biosample record
	// was created.
	Created string `protobuf:"bytes,6,opt,name=created" json:"created,omitempty"`
	// The :ref:`ISO 8601<metadata_date_time>` time at which this Biosample record was
	// updated.
	Updated string `protobuf:"bytes,7,opt,name=updated" json:"updated,omitempty"`
	// The individual this biosample was derived from.
	IndividualId string `protobuf:"bytes,8,opt,name=individual_id,json=individualId" json:"individual_id,omitempty"`
	// the source of the data
	Source string `protobuf:"bytes,9,opt,name=source" json:"source,omitempty"`
	// A map of additional information about the Biosample.
	Attributes *google_protobuf.Struct `protobuf:"bytes,10,opt,name=attributes" json:"attributes,omitempty"`
	// An age object describing the age of the individual this biosample was
	// derived from at the time of collection. The Age object allows the encoding
	// of the age either as ISO8601 duraion or time interval (preferred), or
	// as ontology term object.
	IndividualAgeAtCollection string `protobuf:"bytes,11,opt,name=individual_age_at_collection,json=individualAgeAtCollection" json:"individual_age_at_collection,omitempty"`
}

func (m *Biosample) Reset()                    { *m = Biosample{} }
func (m *Biosample) String() string            { return proto.CompactTextString(m) }
func (*Biosample) ProtoMessage()               {}
func (*Biosample) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Biosample) GetDisease() *OntologyTerm {
	if m != nil {
		return m.Disease
	}
	return nil
}

func (m *Biosample) GetAttributes() *google_protobuf.Struct {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type DrugTherapy struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IndividualId   string `protobuf:"bytes,3,opt,name=individual_id,json=individualId" json:"individual_id,omitempty"`
	DrugName       string `protobuf:"bytes,4,opt,name=drug_name,json=drugName" json:"drug_name,omitempty"`
	PubchemId      string `protobuf:"bytes,5,opt,name=pubchem_id,json=pubchemId" json:"pubchem_id,omitempty"`
	PrescribedDose string `protobuf:"bytes,6,opt,name=prescribed_dose,json=prescribedDose" json:"prescribed_dose,omitempty"`
	Source         string `protobuf:"bytes,7,opt,name=source" json:"source,omitempty"`
}

func (m *DrugTherapy) Reset()                    { *m = DrugTherapy{} }
func (m *DrugTherapy) String() string            { return proto.CompactTextString(m) }
func (*DrugTherapy) ProtoMessage()               {}
func (*DrugTherapy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type RadiationTherapy struct {
	Id           string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IndividualId string  `protobuf:"bytes,3,opt,name=individual_id,json=individualId" json:"individual_id,omitempty"`
	TotalDose    float32 `protobuf:"fixed32,4,opt,name=total_dose,json=totalDose" json:"total_dose,omitempty"`
	Source       string  `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
}

func (m *RadiationTherapy) Reset()                    { *m = RadiationTherapy{} }
func (m *RadiationTherapy) String() string            { return proto.CompactTextString(m) }
func (*RadiationTherapy) ProtoMessage()               {}
func (*RadiationTherapy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type ClinicalFollowup struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IndividualId string `protobuf:"bytes,3,opt,name=individual_id,json=individualId" json:"individual_id,omitempty"`
	Date         string `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
	Event        string `protobuf:"bytes,5,opt,name=event" json:"event,omitempty"`
	DaysToDeath  int32  `protobuf:"varint,6,opt,name=days_to_death,json=daysToDeath" json:"days_to_death,omitempty"`
	VitalStatus  string `protobuf:"bytes,7,opt,name=vital_status,json=vitalStatus" json:"vital_status,omitempty"`
	Source       string `protobuf:"bytes,8,opt,name=source" json:"source,omitempty"`
}

func (m *ClinicalFollowup) Reset()                    { *m = ClinicalFollowup{} }
func (m *ClinicalFollowup) String() string            { return proto.CompactTextString(m) }
func (*ClinicalFollowup) ProtoMessage()               {}
func (*ClinicalFollowup) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Cohort is for groups of Biosamples
type Cohort struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Location    string   `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	HasSample   []string `protobuf:"bytes,7,rep,name=hasSample" json:"hasSample,omitempty"`
}

func (m *Cohort) Reset()                    { *m = Cohort{} }
func (m *Cohort) String() string            { return proto.CompactTextString(m) }
func (*Cohort) ProtoMessage()               {}
func (*Cohort) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// IndividualCohort is for groups of Individuals
type IndividualCohort struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Location    string   `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	HasMember   []string `protobuf:"bytes,7,rep,name=hasMember" json:"hasMember,omitempty"`
}

func (m *IndividualCohort) Reset()                    { *m = IndividualCohort{} }
func (m *IndividualCohort) String() string            { return proto.CompactTextString(m) }
func (*IndividualCohort) ProtoMessage()               {}
func (*IndividualCohort) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*OntologyTerm)(nil), "bmeg.OntologyTerm")
	proto.RegisterType((*Individual)(nil), "bmeg.Individual")
	proto.RegisterType((*Biosample)(nil), "bmeg.Biosample")
	proto.RegisterType((*DrugTherapy)(nil), "bmeg.DrugTherapy")
	proto.RegisterType((*RadiationTherapy)(nil), "bmeg.RadiationTherapy")
	proto.RegisterType((*ClinicalFollowup)(nil), "bmeg.ClinicalFollowup")
	proto.RegisterType((*Cohort)(nil), "bmeg.Cohort")
	proto.RegisterType((*IndividualCohort)(nil), "bmeg.IndividualCohort")
}

func init() { proto.RegisterFile("bmeg/clinical.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x6e, 0xd4, 0x3e,
	0x10, 0xd6, 0xfe, 0xcb, 0x6e, 0x66, 0xdb, 0xfe, 0x7e, 0x32, 0x88, 0x86, 0x52, 0xa4, 0xb2, 0x20,
	0xc1, 0x01, 0x6d, 0x25, 0x38, 0x70, 0xe0, 0x80, 0x4a, 0x2b, 0xa4, 0x1e, 0x00, 0x69, 0xdb, 0x7b,
	0xe4, 0xc4, 0x43, 0xd6, 0x52, 0x36, 0x8e, 0x62, 0xa7, 0xd0, 0x37, 0x80, 0x2b, 0x8f, 0xc5, 0x95,
	0x57, 0xe0, 0x41, 0xb0, 0xc7, 0xd9, 0xdd, 0xb4, 0x65, 0x6f, 0x55, 0x4f, 0xb1, 0xbf, 0x19, 0x4f,
	0xbe, 0xf9, 0xbe, 0xb1, 0xe1, 0x5e, 0xb2, 0xc0, 0xec, 0x30, 0xcd, 0x65, 0x21, 0x53, 0x9e, 0x4f,
	0xcb, 0x4a, 0x19, 0xc5, 0xfa, 0x0e, 0xdc, 0xdb, 0xcf, 0x94, 0xca, 0x72, 0x3c, 0x24, 0x2c, 0xa9,
	0xbf, 0x1c, 0x6a, 0x53, 0xd5, 0xa9, 0xf1, 0x39, 0x93, 0xb7, 0xb0, 0xf5, 0xb9, 0x30, 0x2a, 0x57,
	0xd9, 0xe5, 0x39, 0x56, 0x0b, 0xb6, 0x0b, 0x43, 0x63, 0xbf, 0xb1, 0x14, 0x51, 0xe7, 0xa0, 0xf3,
	0x22, 0x9c, 0x05, 0x6e, 0x7b, 0x2a, 0x18, 0x83, 0xbe, 0x5b, 0x45, 0x5d, 0x42, 0x69, 0x3d, 0xf9,
	0xd5, 0x05, 0x38, 0x2d, 0x84, 0xbc, 0x90, 0xa2, 0xe6, 0x39, 0xdb, 0x81, 0xee, 0xea, 0x98, 0x5d,
	0xb1, 0xc7, 0x00, 0x82, 0x1b, 0xae, 0xd1, 0xb8, 0x72, 0xfe, 0x60, 0xd8, 0x20, 0xbe, 0x62, 0xc1,
	0x17, 0x18, 0xf5, 0x7c, 0x45, 0xb7, 0x66, 0x07, 0x30, 0x16, 0xa8, 0xd3, 0x4a, 0x96, 0x46, 0xaa,
	0x22, 0xea, 0x53, 0xa8, 0x0d, 0xb1, 0x08, 0x86, 0x69, 0x85, 0xdc, 0xa0, 0x88, 0x06, 0x14, 0x5d,
	0x6e, 0x5d, 0xa4, 0x2e, 0x05, 0x45, 0x02, 0x1f, 0x69, 0xb6, 0xec, 0x25, 0x0c, 0x75, 0x89, 0xa9,
	0x44, 0x1d, 0x0d, 0x6d, 0x64, 0xfc, 0x8a, 0x4d, 0x9d, 0x34, 0xd3, 0x76, 0xe7, 0xb3, 0x65, 0x0a,
	0x7b, 0x06, 0x3d, 0x8d, 0xdf, 0xa2, 0xd1, 0xc6, 0x4c, 0x17, 0x66, 0x0f, 0x20, 0xd0, 0xaa, 0xae,
	0x52, 0x8c, 0x42, 0xaf, 0x93, 0xdf, 0xb1, 0x37, 0x00, 0xdc, 0x98, 0x4a, 0x26, 0xb5, 0xb1, 0xbf,
	0x03, 0x2a, 0xb2, 0x3b, 0xf5, 0x1e, 0x4c, 0x97, 0x1e, 0x4c, 0xcf, 0xc8, 0x83, 0x59, 0x2b, 0x75,
	0xf2, 0xa3, 0x07, 0xe1, 0x7b, 0xa9, 0x34, 0x5f, 0x94, 0x39, 0xde, 0x8d, 0x96, 0x56, 0x17, 0x21,
	0x35, 0xda, 0x1a, 0xa4, 0xe5, 0x06, 0x5d, 0x9a, 0x94, 0xb6, 0xf2, 0xc1, 0x46, 0xe5, 0x87, 0x57,
	0x95, 0x7f, 0x0a, 0xdb, 0x72, 0x35, 0x20, 0x8e, 0xf9, 0x88, 0xe2, 0x5b, 0x6b, 0xd0, 0x92, 0xbf,
	0x6d, 0x29, 0xd9, 0x3b, 0xd8, 0x6f, 0xfd, 0x95, 0x67, 0x18, 0x73, 0x13, 0xa7, 0x2a, 0xcf, 0x31,
	0x25, 0x29, 0xc6, 0xf4, 0x9b, 0x87, 0xeb, 0x9c, 0xa3, 0x0c, 0x8f, 0xcc, 0xf1, 0x2a, 0x61, 0xf2,
	0xbb, 0x03, 0xe3, 0x93, 0xaa, 0xce, 0xce, 0xe7, 0x58, 0xf1, 0xf2, 0xf2, 0x86, 0x1b, 0x4b, 0xb9,
	0xbb, 0x2d, 0xb9, 0x6f, 0xb4, 0xda, 0xfb, 0x47, 0xab, 0x8f, 0x20, 0x14, 0xb6, 0x6e, 0x4c, 0xa7,
	0xbd, 0x23, 0x23, 0x07, 0x7c, 0x72, 0x15, 0xac, 0xc7, 0x65, 0x9d, 0xa4, 0x73, 0xa4, 0xeb, 0xe7,
	0xa7, 0x3b, 0x6c, 0x10, 0x7b, 0xf6, 0x39, 0xfc, 0x57, 0x56, 0xe4, 0x5e, 0x82, 0x22, 0x16, 0xca,
	0xba, 0xe6, 0x7d, 0xd8, 0x59, 0xc3, 0x27, 0x16, 0x6d, 0xe9, 0x39, 0x6c, 0xeb, 0x39, 0xf9, 0xd9,
	0x81, 0xff, 0x67, 0x5c, 0x48, 0xee, 0x7a, 0xbc, 0xf5, 0xd6, 0x2c, 0x7b, 0xa3, 0x8c, 0x8d, 0x13,
	0x33, 0xd7, 0x5b, 0x77, 0x16, 0x12, 0x72, 0x8d, 0xd4, 0xe0, 0x0a, 0xa9, 0x3f, 0x96, 0xd4, 0x71,
	0xf3, 0x6e, 0x7d, 0xb0, 0x0e, 0xa8, 0xaf, 0x75, 0x79, 0x7b, 0xa4, 0xec, 0x41, 0x37, 0x88, 0x8d,
	0xd4, 0xb4, 0x66, 0xf7, 0x61, 0x80, 0x17, 0x58, 0x98, 0x86, 0x88, 0xdf, 0xb0, 0x09, 0x6c, 0x0b,
	0x7e, 0xa9, 0x63, 0xa3, 0x62, 0x61, 0xa7, 0x7a, 0x4e, 0xda, 0x0e, 0xec, 0x7d, 0xb1, 0xe0, 0xb9,
	0x3a, 0x71, 0x10, 0x7b, 0x02, 0x5b, 0x17, 0xd2, 0xb5, 0xa8, 0x0d, 0x37, 0xb5, 0x6e, 0xe4, 0x1d,
	0x13, 0x76, 0x46, 0x50, 0xab, 0xcd, 0xd1, 0x95, 0x36, 0xbf, 0x77, 0x20, 0x38, 0x56, 0x73, 0x55,
	0x99, 0x8d, 0xcd, 0xf5, 0x5b, 0xcd, 0xed, 0xc1, 0x28, 0x57, 0x29, 0x19, 0xd5, 0xd0, 0x5c, 0xed,
	0xaf, 0xdf, 0xeb, 0xe0, 0xe6, 0xbd, 0xde, 0x87, 0x70, 0xce, 0xf5, 0x19, 0xbd, 0x24, 0x96, 0x64,
	0xcf, 0xcd, 0xd1, 0x0a, 0xa0, 0x31, 0x58, 0xbf, 0xda, 0x77, 0x4c, 0xea, 0x23, 0x2e, 0x12, 0xac,
	0x5a, 0xa4, 0x3c, 0x90, 0x04, 0x74, 0x9f, 0x5f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x51,
	0xee, 0x70, 0xc9, 0x06, 0x00, 0x00,
}
