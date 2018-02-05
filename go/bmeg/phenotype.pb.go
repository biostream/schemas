// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bmeg/phenotype.proto

package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResponseSummary_SummaryType int32

const (
	ResponseSummary_UNKNOWN       ResponseSummary_SummaryType = 0
	ResponseSummary_EC50          ResponseSummary_SummaryType = 1
	ResponseSummary_IC50          ResponseSummary_SummaryType = 2
	ResponseSummary_AC50          ResponseSummary_SummaryType = 3
	ResponseSummary_LD50          ResponseSummary_SummaryType = 4
	ResponseSummary_GR50          ResponseSummary_SummaryType = 5
	ResponseSummary_AMAX          ResponseSummary_SummaryType = 6
	ResponseSummary_AUC           ResponseSummary_SummaryType = 7
	ResponseSummary_ACTIVITY_AREA ResponseSummary_SummaryType = 8
	ResponseSummary_RMSE          ResponseSummary_SummaryType = 9
)

var ResponseSummary_SummaryType_name = map[int32]string{
	0: "UNKNOWN",
	1: "EC50",
	2: "IC50",
	3: "AC50",
	4: "LD50",
	5: "GR50",
	6: "AMAX",
	7: "AUC",
	8: "ACTIVITY_AREA",
	9: "RMSE",
}
var ResponseSummary_SummaryType_value = map[string]int32{
	"UNKNOWN":       0,
	"EC50":          1,
	"IC50":          2,
	"AC50":          3,
	"LD50":          4,
	"GR50":          5,
	"AMAX":          6,
	"AUC":           7,
	"ACTIVITY_AREA": 8,
	"RMSE":          9,
}

func (x ResponseSummary_SummaryType) String() string {
	return proto.EnumName(ResponseSummary_SummaryType_name, int32(x))
}
func (ResponseSummary_SummaryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{6, 0}
}

type ResponseCurve_ResponseType int32

const (
	ResponseCurve_UNKNOWN  ResponseCurve_ResponseType = 0
	ResponseCurve_GROWTH   ResponseCurve_ResponseType = 1
	ResponseCurve_ACTIVITY ResponseCurve_ResponseType = 2
)

var ResponseCurve_ResponseType_name = map[int32]string{
	0: "UNKNOWN",
	1: "GROWTH",
	2: "ACTIVITY",
}
var ResponseCurve_ResponseType_value = map[string]int32{
	"UNKNOWN":  0,
	"GROWTH":   1,
	"ACTIVITY": 2,
}

func (x ResponseCurve_ResponseType) String() string {
	return proto.EnumName(ResponseCurve_ResponseType_name, int32(x))
}
func (ResponseCurve_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{8, 0}
}

//
// Gene Ontology
type GeneOntologyTerm struct {
	Id         string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace  string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Definition string   `protobuf:"bytes,4,opt,name=definition" json:"definition,omitempty"`
	Comment    string   `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
	Synonym    []string `protobuf:"bytes,6,rep,name=synonym" json:"synonym,omitempty"`
	IsA        []string `protobuf:"bytes,7,rep,name=is_a,json=isA" json:"is_a,omitempty"`
	AltId      []string `protobuf:"bytes,8,rep,name=alt_id,json=altId" json:"alt_id,omitempty"`
	Subset     []string `protobuf:"bytes,9,rep,name=subset" json:"subset,omitempty"`
	Xref       []string `protobuf:"bytes,10,rep,name=xref" json:"xref,omitempty"`
	IsObsolete bool     `protobuf:"varint,11,opt,name=is_obsolete,json=isObsolete" json:"is_obsolete,omitempty"`
	Consider   []string `protobuf:"bytes,12,rep,name=consider" json:"consider,omitempty"`
}

func (m *GeneOntologyTerm) Reset()                    { *m = GeneOntologyTerm{} }
func (m *GeneOntologyTerm) String() string            { return proto.CompactTextString(m) }
func (*GeneOntologyTerm) ProtoMessage()               {}
func (*GeneOntologyTerm) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *GeneOntologyTerm) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GeneOntologyTerm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GeneOntologyTerm) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GeneOntologyTerm) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

func (m *GeneOntologyTerm) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *GeneOntologyTerm) GetSynonym() []string {
	if m != nil {
		return m.Synonym
	}
	return nil
}

func (m *GeneOntologyTerm) GetIsA() []string {
	if m != nil {
		return m.IsA
	}
	return nil
}

func (m *GeneOntologyTerm) GetAltId() []string {
	if m != nil {
		return m.AltId
	}
	return nil
}

func (m *GeneOntologyTerm) GetSubset() []string {
	if m != nil {
		return m.Subset
	}
	return nil
}

func (m *GeneOntologyTerm) GetXref() []string {
	if m != nil {
		return m.Xref
	}
	return nil
}

func (m *GeneOntologyTerm) GetIsObsolete() bool {
	if m != nil {
		return m.IsObsolete
	}
	return false
}

func (m *GeneOntologyTerm) GetConsider() []string {
	if m != nil {
		return m.Consider
	}
	return nil
}

type GeneOntologyAnnotation struct {
	Title      string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Genes      []string `protobuf:"bytes,2,rep,name=genes" json:"genes,omitempty"`
	Functions  []string `protobuf:"bytes,3,rep,name=functions" json:"functions,omitempty"`
	Evidence   []string `protobuf:"bytes,4,rep,name=evidence" json:"evidence,omitempty"`
	References []string `protobuf:"bytes,5,rep,name=references" json:"references,omitempty"`
}

func (m *GeneOntologyAnnotation) Reset()                    { *m = GeneOntologyAnnotation{} }
func (m *GeneOntologyAnnotation) String() string            { return proto.CompactTextString(m) }
func (*GeneOntologyAnnotation) ProtoMessage()               {}
func (*GeneOntologyAnnotation) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *GeneOntologyAnnotation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GeneOntologyAnnotation) GetGenes() []string {
	if m != nil {
		return m.Genes
	}
	return nil
}

func (m *GeneOntologyAnnotation) GetFunctions() []string {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *GeneOntologyAnnotation) GetEvidence() []string {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *GeneOntologyAnnotation) GetReferences() []string {
	if m != nil {
		return m.References
	}
	return nil
}

type Compound struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Target      []string `protobuf:"bytes,3,rep,name=target" json:"target,omitempty"`
	Description []string `protobuf:"bytes,4,rep,name=description" json:"description,omitempty"`
	Synonyms    []string `protobuf:"bytes,5,rep,name=synonyms" json:"synonyms,omitempty"`
	Pubmed      []string `protobuf:"bytes,6,rep,name=pubmed" json:"pubmed,omitempty"`
	Report      string   `protobuf:"bytes,7,opt,name=report" json:"report,omitempty"`
	Status      string   `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	Rationale   string   `protobuf:"bytes,9,opt,name=rationale" json:"rationale,omitempty"`
	Smiles      string   `protobuf:"bytes,10,opt,name=smiles" json:"smiles,omitempty"`
	Assays      []string `protobuf:"bytes,11,rep,name=assays" json:"assays,omitempty"`
	Source      []string `protobuf:"bytes,12,rep,name=source" json:"source,omitempty"`
	Sids        []string `protobuf:"bytes,13,rep,name=sids" json:"sids,omitempty"`
	ChebiId     string   `protobuf:"bytes,14,opt,name=chebi_id,json=chebiId" json:"chebi_id,omitempty"`
}

func (m *Compound) Reset()                    { *m = Compound{} }
func (m *Compound) String() string            { return proto.CompactTextString(m) }
func (*Compound) ProtoMessage()               {}
func (*Compound) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Compound) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Compound) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Compound) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Compound) GetDescription() []string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Compound) GetSynonyms() []string {
	if m != nil {
		return m.Synonyms
	}
	return nil
}

func (m *Compound) GetPubmed() []string {
	if m != nil {
		return m.Pubmed
	}
	return nil
}

func (m *Compound) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *Compound) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Compound) GetRationale() string {
	if m != nil {
		return m.Rationale
	}
	return ""
}

func (m *Compound) GetSmiles() string {
	if m != nil {
		return m.Smiles
	}
	return ""
}

func (m *Compound) GetAssays() []string {
	if m != nil {
		return m.Assays
	}
	return nil
}

func (m *Compound) GetSource() []string {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Compound) GetSids() []string {
	if m != nil {
		return m.Sids
	}
	return nil
}

func (m *Compound) GetChebiId() string {
	if m != nil {
		return m.ChebiId
	}
	return ""
}

type Assay struct {
	Id           string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Sid          string           `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Name         string           `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Cellline     string           `protobuf:"bytes,4,opt,name=cellline" json:"cellline,omitempty"`
	Pubmed       []string         `protobuf:"bytes,5,rep,name=pubmed" json:"pubmed,omitempty"`
	Drugresponse []*ResponseCurve `protobuf:"bytes,6,rep,name=drugresponse" json:"drugresponse,omitempty"`
	GeneTarget   []string         `protobuf:"bytes,7,rep,name=gene_target,json=geneTarget" json:"gene_target,omitempty"`
}

func (m *Assay) Reset()                    { *m = Assay{} }
func (m *Assay) String() string            { return proto.CompactTextString(m) }
func (*Assay) ProtoMessage()               {}
func (*Assay) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Assay) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Assay) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *Assay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assay) GetCellline() string {
	if m != nil {
		return m.Cellline
	}
	return ""
}

func (m *Assay) GetPubmed() []string {
	if m != nil {
		return m.Pubmed
	}
	return nil
}

func (m *Assay) GetDrugresponse() []*ResponseCurve {
	if m != nil {
		return m.Drugresponse
	}
	return nil
}

func (m *Assay) GetGeneTarget() []string {
	if m != nil {
		return m.GeneTarget
	}
	return nil
}

type FDAApproval struct {
	CompoundId            string `protobuf:"bytes,1,opt,name=compound_id,json=compoundId" json:"compound_id,omitempty"`
	Approved              string `protobuf:"bytes,2,opt,name=approved" json:"approved,omitempty"`
	ApprovedData          string `protobuf:"bytes,3,opt,name=approved_data,json=approvedData" json:"approved_data,omitempty"`
	Company               string `protobuf:"bytes,4,opt,name=company" json:"company,omitempty"`
	Summary               string `protobuf:"bytes,5,opt,name=summary" json:"summary,omitempty"`
	GeneralInformation    string `protobuf:"bytes,6,opt,name=general_information,json=generalInformation" json:"general_information,omitempty"`
	ClinicalResults       string `protobuf:"bytes,7,opt,name=clinical_results,json=clinicalResults" json:"clinical_results,omitempty"`
	SideEffects           string `protobuf:"bytes,8,opt,name=side_effects,json=sideEffects" json:"side_effects,omitempty"`
	MechanismOfAction     string `protobuf:"bytes,9,opt,name=mechanism_of_action,json=mechanismOfAction" json:"mechanism_of_action,omitempty"`
	AdditionalInformation string `protobuf:"bytes,10,opt,name=additional_information,json=additionalInformation" json:"additional_information,omitempty"`
	TheraputicAreas       string `protobuf:"bytes,11,opt,name=theraputic_areas,json=theraputicAreas" json:"theraputic_areas,omitempty"`
}

func (m *FDAApproval) Reset()                    { *m = FDAApproval{} }
func (m *FDAApproval) String() string            { return proto.CompactTextString(m) }
func (*FDAApproval) ProtoMessage()               {}
func (*FDAApproval) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *FDAApproval) GetCompoundId() string {
	if m != nil {
		return m.CompoundId
	}
	return ""
}

func (m *FDAApproval) GetApproved() string {
	if m != nil {
		return m.Approved
	}
	return ""
}

func (m *FDAApproval) GetApprovedData() string {
	if m != nil {
		return m.ApprovedData
	}
	return ""
}

func (m *FDAApproval) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *FDAApproval) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *FDAApproval) GetGeneralInformation() string {
	if m != nil {
		return m.GeneralInformation
	}
	return ""
}

func (m *FDAApproval) GetClinicalResults() string {
	if m != nil {
		return m.ClinicalResults
	}
	return ""
}

func (m *FDAApproval) GetSideEffects() string {
	if m != nil {
		return m.SideEffects
	}
	return ""
}

func (m *FDAApproval) GetMechanismOfAction() string {
	if m != nil {
		return m.MechanismOfAction
	}
	return ""
}

func (m *FDAApproval) GetAdditionalInformation() string {
	if m != nil {
		return m.AdditionalInformation
	}
	return ""
}

func (m *FDAApproval) GetTheraputicAreas() string {
	if m != nil {
		return m.TheraputicAreas
	}
	return ""
}

type DoseResponse struct {
	Dose     float64 `protobuf:"fixed64,1,opt,name=dose" json:"dose,omitempty"`
	Response float64 `protobuf:"fixed64,2,opt,name=response" json:"response,omitempty"`
}

func (m *DoseResponse) Reset()                    { *m = DoseResponse{} }
func (m *DoseResponse) String() string            { return proto.CompactTextString(m) }
func (*DoseResponse) ProtoMessage()               {}
func (*DoseResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *DoseResponse) GetDose() float64 {
	if m != nil {
		return m.Dose
	}
	return 0
}

func (m *DoseResponse) GetResponse() float64 {
	if m != nil {
		return m.Response
	}
	return 0
}

type ResponseSummary struct {
	Type  ResponseSummary_SummaryType `protobuf:"varint,1,opt,name=type,enum=bmeg.ResponseSummary_SummaryType" json:"type,omitempty"`
	Value float64                     `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	Unit  string                      `protobuf:"bytes,3,opt,name=unit" json:"unit,omitempty"`
}

func (m *ResponseSummary) Reset()                    { *m = ResponseSummary{} }
func (m *ResponseSummary) String() string            { return proto.CompactTextString(m) }
func (*ResponseSummary) ProtoMessage()               {}
func (*ResponseSummary) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ResponseSummary) GetType() ResponseSummary_SummaryType {
	if m != nil {
		return m.Type
	}
	return ResponseSummary_UNKNOWN
}

func (m *ResponseSummary) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ResponseSummary) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type CompoundElement struct {
	Compound string  `protobuf:"bytes,1,opt,name=compound" json:"compound,omitempty"`
	Ratio    float64 `protobuf:"fixed64,3,opt,name=ratio" json:"ratio,omitempty"`
}

func (m *CompoundElement) Reset()                    { *m = CompoundElement{} }
func (m *CompoundElement) String() string            { return proto.CompactTextString(m) }
func (*CompoundElement) ProtoMessage()               {}
func (*CompoundElement) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *CompoundElement) GetCompound() string {
	if m != nil {
		return m.Compound
	}
	return ""
}

func (m *CompoundElement) GetRatio() float64 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

type ResponseCurve struct {
	Gid            string                     `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	AssayId        string                     `protobuf:"bytes,2,opt,name=assay_id,json=assayId" json:"assay_id,omitempty"`
	ResponseType   ResponseCurve_ResponseType `protobuf:"varint,3,opt,name=responseType,enum=bmeg.ResponseCurve_ResponseType" json:"responseType,omitempty"`
	Values         []*DoseResponse            `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
	GrowthStandard float64                    `protobuf:"fixed64,5,opt,name=growthStandard" json:"growthStandard,omitempty"`
	Compounds      []*CompoundElement         `protobuf:"bytes,6,rep,name=compounds" json:"compounds,omitempty"`
	Sample         string                     `protobuf:"bytes,7,opt,name=sample" json:"sample,omitempty"`
	Summary        []*ResponseSummary         `protobuf:"bytes,8,rep,name=summary" json:"summary,omitempty"`
	Controls       []float64                  `protobuf:"fixed64,9,rep,packed,name=controls" json:"controls,omitempty"`
	Blanks         []float64                  `protobuf:"fixed64,10,rep,packed,name=blanks" json:"blanks,omitempty"`
	Source         string                     `protobuf:"bytes,11,opt,name=source" json:"source,omitempty"`
}

func (m *ResponseCurve) Reset()                    { *m = ResponseCurve{} }
func (m *ResponseCurve) String() string            { return proto.CompactTextString(m) }
func (*ResponseCurve) ProtoMessage()               {}
func (*ResponseCurve) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ResponseCurve) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *ResponseCurve) GetAssayId() string {
	if m != nil {
		return m.AssayId
	}
	return ""
}

func (m *ResponseCurve) GetResponseType() ResponseCurve_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ResponseCurve_UNKNOWN
}

func (m *ResponseCurve) GetValues() []*DoseResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ResponseCurve) GetGrowthStandard() float64 {
	if m != nil {
		return m.GrowthStandard
	}
	return 0
}

func (m *ResponseCurve) GetCompounds() []*CompoundElement {
	if m != nil {
		return m.Compounds
	}
	return nil
}

func (m *ResponseCurve) GetSample() string {
	if m != nil {
		return m.Sample
	}
	return ""
}

func (m *ResponseCurve) GetSummary() []*ResponseSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *ResponseCurve) GetControls() []float64 {
	if m != nil {
		return m.Controls
	}
	return nil
}

func (m *ResponseCurve) GetBlanks() []float64 {
	if m != nil {
		return m.Blanks
	}
	return nil
}

func (m *ResponseCurve) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type AssayData struct {
	AssayId    string             `protobuf:"bytes,1,opt,name=assay_id,json=assayId" json:"assay_id,omitempty"`
	CompoundId string             `protobuf:"bytes,2,opt,name=compound_id,json=compoundId" json:"compound_id,omitempty"`
	Active     bool               `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	Comment    string             `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	FloatVals  map[string]float64 `protobuf:"bytes,5,rep,name=float_vals,json=floatVals" json:"float_vals,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	StringVals map[string]string  `protobuf:"bytes,6,rep,name=string_vals,json=stringVals" json:"string_vals,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IntVals    map[string]int32   `protobuf:"bytes,7,rep,name=int_vals,json=intVals" json:"int_vals,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *AssayData) Reset()                    { *m = AssayData{} }
func (m *AssayData) String() string            { return proto.CompactTextString(m) }
func (*AssayData) ProtoMessage()               {}
func (*AssayData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *AssayData) GetAssayId() string {
	if m != nil {
		return m.AssayId
	}
	return ""
}

func (m *AssayData) GetCompoundId() string {
	if m != nil {
		return m.CompoundId
	}
	return ""
}

func (m *AssayData) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *AssayData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AssayData) GetFloatVals() map[string]float64 {
	if m != nil {
		return m.FloatVals
	}
	return nil
}

func (m *AssayData) GetStringVals() map[string]string {
	if m != nil {
		return m.StringVals
	}
	return nil
}

func (m *AssayData) GetIntVals() map[string]int32 {
	if m != nil {
		return m.IntVals
	}
	return nil
}

type Evidence struct {
	// placeholder
	Evidence string `protobuf:"bytes,1,opt,name=evidence" json:"evidence,omitempty"`
}

func (m *Evidence) Reset()                    { *m = Evidence{} }
func (m *Evidence) String() string            { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()               {}
func (*Evidence) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *Evidence) GetEvidence() string {
	if m != nil {
		return m.Evidence
	}
	return ""
}

type G2PAssociation struct {
	Genes         []string                `protobuf:"bytes,1,rep,name=genes" json:"genes,omitempty"`
	Environments  []string                `protobuf:"bytes,2,rep,name=environments" json:"environments,omitempty"`
	Phenotypes    []string                `protobuf:"bytes,3,rep,name=phenotypes" json:"phenotypes,omitempty"`
	Features      []string                `protobuf:"bytes,4,rep,name=features" json:"features,omitempty"`
	Publications  []string                `protobuf:"bytes,5,rep,name=publications" json:"publications,omitempty"`
	Evidence      []*Evidence             `protobuf:"bytes,7,rep,name=evidence" json:"evidence,omitempty"`
	Description   string                  `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	EvidenceLabel string                  `protobuf:"bytes,8,opt,name=evidence_label,json=evidenceLabel" json:"evidence_label,omitempty"`
	ResponseType  string                  `protobuf:"bytes,10,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	Original      *google_protobuf.Struct `protobuf:"bytes,9,opt,name=original" json:"original,omitempty"`
	Source        string                  `protobuf:"bytes,11,opt,name=source" json:"source,omitempty"`
}

func (m *G2PAssociation) Reset()                    { *m = G2PAssociation{} }
func (m *G2PAssociation) String() string            { return proto.CompactTextString(m) }
func (*G2PAssociation) ProtoMessage()               {}
func (*G2PAssociation) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *G2PAssociation) GetGenes() []string {
	if m != nil {
		return m.Genes
	}
	return nil
}

func (m *G2PAssociation) GetEnvironments() []string {
	if m != nil {
		return m.Environments
	}
	return nil
}

func (m *G2PAssociation) GetPhenotypes() []string {
	if m != nil {
		return m.Phenotypes
	}
	return nil
}

func (m *G2PAssociation) GetFeatures() []string {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *G2PAssociation) GetPublications() []string {
	if m != nil {
		return m.Publications
	}
	return nil
}

func (m *G2PAssociation) GetEvidence() []*Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *G2PAssociation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *G2PAssociation) GetEvidenceLabel() string {
	if m != nil {
		return m.EvidenceLabel
	}
	return ""
}

func (m *G2PAssociation) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *G2PAssociation) GetOriginal() *google_protobuf.Struct {
	if m != nil {
		return m.Original
	}
	return nil
}

func (m *G2PAssociation) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func init() {
	proto.RegisterType((*GeneOntologyTerm)(nil), "bmeg.GeneOntologyTerm")
	proto.RegisterType((*GeneOntologyAnnotation)(nil), "bmeg.GeneOntologyAnnotation")
	proto.RegisterType((*Compound)(nil), "bmeg.Compound")
	proto.RegisterType((*Assay)(nil), "bmeg.Assay")
	proto.RegisterType((*FDAApproval)(nil), "bmeg.FDAApproval")
	proto.RegisterType((*DoseResponse)(nil), "bmeg.DoseResponse")
	proto.RegisterType((*ResponseSummary)(nil), "bmeg.ResponseSummary")
	proto.RegisterType((*CompoundElement)(nil), "bmeg.CompoundElement")
	proto.RegisterType((*ResponseCurve)(nil), "bmeg.ResponseCurve")
	proto.RegisterType((*AssayData)(nil), "bmeg.AssayData")
	proto.RegisterType((*Evidence)(nil), "bmeg.Evidence")
	proto.RegisterType((*G2PAssociation)(nil), "bmeg.G2PAssociation")
	proto.RegisterEnum("bmeg.ResponseSummary_SummaryType", ResponseSummary_SummaryType_name, ResponseSummary_SummaryType_value)
	proto.RegisterEnum("bmeg.ResponseCurve_ResponseType", ResponseCurve_ResponseType_name, ResponseCurve_ResponseType_value)
}

func init() { proto.RegisterFile("bmeg/phenotype.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xdd, 0x6e, 0x1b, 0xc7,
	0x15, 0xee, 0xf2, 0x9f, 0x87, 0x14, 0x45, 0x8f, 0x6d, 0x75, 0x2b, 0x18, 0xb6, 0xcc, 0xa2, 0x86,
	0xeb, 0x0b, 0xca, 0x90, 0x21, 0xb8, 0x30, 0xea, 0xa2, 0x84, 0x24, 0xab, 0x44, 0x6d, 0xab, 0x58,
	0xc9, 0x76, 0x7b, 0xb5, 0x18, 0xee, 0x0e, 0xa9, 0x81, 0x97, 0xb3, 0x8b, 0x9d, 0x59, 0x35, 0x04,
	0xf2, 0x12, 0x79, 0x81, 0x3c, 0x45, 0xee, 0xf2, 0x08, 0x79, 0x80, 0x5c, 0xe7, 0x3e, 0x17, 0x79,
	0x84, 0xe0, 0xcc, 0xcf, 0x72, 0x49, 0x3b, 0x70, 0xae, 0x38, 0xdf, 0x37, 0xe7, 0xcc, 0x9e, 0x39,
	0x7f, 0x73, 0x08, 0x77, 0x66, 0x4b, 0xb6, 0x38, 0xcc, 0xae, 0x99, 0x48, 0xd5, 0x2a, 0x63, 0xe3,
	0x2c, 0x4f, 0x55, 0x4a, 0x1a, 0xc8, 0xee, 0xdf, 0x5b, 0xa4, 0xe9, 0x22, 0x61, 0x87, 0x9a, 0x9b,
	0x15, 0xf3, 0x43, 0xa9, 0xf2, 0x22, 0x52, 0x46, 0x66, 0xf4, 0x7d, 0x0d, 0x86, 0xe7, 0x4c, 0xb0,
	0x0b, 0xa1, 0xd2, 0x24, 0x5d, 0xac, 0xae, 0x58, 0xbe, 0x24, 0x03, 0xa8, 0xf1, 0xd8, 0xf7, 0x0e,
	0xbc, 0xc7, 0xdd, 0xa0, 0xc6, 0x63, 0x42, 0xa0, 0x21, 0xe8, 0x92, 0xf9, 0x35, 0xcd, 0xe8, 0x35,
	0xb9, 0x07, 0x5d, 0xfc, 0x95, 0x19, 0x8d, 0x98, 0x5f, 0xd7, 0x1b, 0x6b, 0x82, 0xdc, 0x07, 0x88,
	0xd9, 0x9c, 0x0b, 0xae, 0x78, 0x2a, 0xfc, 0x86, 0xde, 0xae, 0x30, 0xc4, 0x87, 0x76, 0x94, 0x2e,
	0x97, 0x4c, 0x28, 0xbf, 0xa9, 0x37, 0x1d, 0xc4, 0x1d, 0xb9, 0x12, 0xa9, 0x58, 0x2d, 0xfd, 0xd6,
	0x41, 0x1d, 0x77, 0x2c, 0x24, 0xb7, 0xa0, 0xc1, 0x65, 0x48, 0xfd, 0xb6, 0xa6, 0xeb, 0x5c, 0x4e,
	0xc8, 0x5d, 0x68, 0xd1, 0x44, 0x85, 0x3c, 0xf6, 0x3b, 0x9a, 0x6c, 0xd2, 0x44, 0x4d, 0x63, 0xb2,
	0x07, 0x2d, 0x59, 0xcc, 0x24, 0x53, 0x7e, 0x57, 0xd3, 0x16, 0xe1, 0x3d, 0xbe, 0xca, 0xd9, 0xdc,
	0x07, 0xcd, 0xea, 0x35, 0x79, 0x00, 0x3d, 0x2e, 0xc3, 0x74, 0x26, 0xd3, 0x84, 0x29, 0xe6, 0xf7,
	0x0e, 0xbc, 0xc7, 0x9d, 0x00, 0xb8, 0xbc, 0xb0, 0x0c, 0xd9, 0x87, 0x4e, 0x94, 0x0a, 0xc9, 0x63,
	0x96, 0xfb, 0x7d, 0xad, 0x58, 0xe2, 0xd1, 0xb7, 0x1e, 0xec, 0x55, 0xbd, 0x37, 0x11, 0x22, 0x55,
	0x54, 0xdf, 0xf0, 0x0e, 0x34, 0x15, 0x57, 0x09, 0xb3, 0x6e, 0x34, 0x00, 0xd9, 0x05, 0x13, 0x4c,
	0xfa, 0x35, 0x63, 0xaf, 0x06, 0xe8, 0xcb, 0x79, 0x21, 0x22, 0xd4, 0x93, 0x7e, 0x5d, 0xef, 0xac,
	0x09, 0x34, 0x80, 0xdd, 0xf0, 0x98, 0x89, 0x88, 0xf9, 0x0d, 0x63, 0x80, 0xc3, 0xe8, 0xe7, 0x9c,
	0xcd, 0x59, 0x8e, 0x40, 0xfa, 0x4d, 0xbd, 0x5b, 0x61, 0x46, 0x3f, 0xd5, 0xa0, 0x73, 0x92, 0x2e,
	0xb3, 0xb4, 0x10, 0xf1, 0xef, 0x0a, 0xeb, 0x1e, 0xb4, 0x14, 0xcd, 0x17, 0x4c, 0x59, 0x3b, 0x2c,
	0x22, 0x07, 0xd0, 0x8b, 0x99, 0x8c, 0x72, 0x9e, 0xd9, 0x88, 0xe2, 0x66, 0x95, 0x42, 0x33, 0x6d,
	0xa4, 0x9c, 0x21, 0x25, 0xc6, 0x53, 0xb3, 0x62, 0xb6, 0x64, 0xb1, 0x8d, 0xa9, 0x45, 0xc8, 0xe7,
	0x2c, 0x4b, 0x73, 0xe5, 0xb7, 0xb5, 0x0d, 0x16, 0xe9, 0x00, 0x2a, 0xaa, 0x0a, 0xe9, 0x77, 0x0c,
	0x6f, 0x10, 0x3a, 0x2a, 0xd7, 0xee, 0xa5, 0x09, 0xf3, 0xbb, 0x26, 0xe9, 0x4a, 0x42, 0x6b, 0x2d,
	0x79, 0xc2, 0xa4, 0x0f, 0x56, 0x4b, 0x23, 0xe4, 0xa9, 0x94, 0x74, 0x25, 0xfd, 0x9e, 0xf9, 0xba,
	0x41, 0x5a, 0x3e, 0x2d, 0xf2, 0x88, 0xd9, 0xb8, 0x5a, 0x84, 0x7e, 0x91, 0x3c, 0x96, 0xfe, 0x8e,
	0x49, 0x13, 0x5c, 0x93, 0x3f, 0x41, 0x27, 0xba, 0x66, 0x33, 0x8e, 0xb9, 0x36, 0xb0, 0x19, 0x8b,
	0x78, 0x1a, 0x8f, 0x7e, 0xf0, 0xa0, 0x39, 0xc1, 0x13, 0x3f, 0x71, 0xf0, 0x10, 0xea, 0x92, 0xc7,
	0xd6, 0xbf, 0xb8, 0x2c, 0x5d, 0x5e, 0xaf, 0xb8, 0x1c, 0x13, 0x8c, 0x25, 0x49, 0xc2, 0x05, 0xb3,
	0x95, 0x52, 0xe2, 0x8a, 0xe3, 0x9a, 0x1b, 0x8e, 0x7b, 0x0e, 0xfd, 0x38, 0x2f, 0x16, 0x39, 0x93,
	0x59, 0x2a, 0x24, 0xd3, 0x6e, 0xed, 0x1d, 0xdd, 0x1e, 0x63, 0xc5, 0x8f, 0x03, 0xcb, 0x9e, 0x14,
	0xf9, 0x0d, 0x0b, 0x36, 0x04, 0x31, 0xdd, 0x31, 0xe7, 0x42, 0x1b, 0x64, 0x53, 0x4b, 0x80, 0xd4,
	0x95, 0x66, 0x46, 0xdf, 0xd5, 0xa1, 0xf7, 0xea, 0x74, 0x32, 0xc9, 0xb2, 0x3c, 0xbd, 0xa1, 0x09,
	0x2a, 0x44, 0x36, 0x81, 0xc2, 0xf2, 0x72, 0xe0, 0xa8, 0x69, 0x8c, 0xe6, 0x53, 0x2d, 0xcc, 0xdc,
	0x4d, 0x4b, 0x4c, 0xfe, 0x0c, 0x3b, 0x6e, 0x1d, 0xc6, 0x54, 0x51, 0x7b, 0xef, 0xbe, 0x23, 0x4f,
	0xa9, 0xa2, 0xb6, 0x17, 0x64, 0x54, 0xac, 0xec, 0xf5, 0x1d, 0xd4, 0xbd, 0xa0, 0x58, 0x2e, 0x69,
	0xbe, 0x72, 0x5d, 0xc2, 0x42, 0x72, 0x08, 0xb7, 0xd1, 0xe6, 0x9c, 0x26, 0x21, 0x17, 0xf3, 0x34,
	0x5f, 0xea, 0x24, 0xf0, 0x5b, 0x5a, 0x8a, 0xd8, 0xad, 0xe9, 0x7a, 0x87, 0xfc, 0x15, 0x86, 0x51,
	0xc2, 0x05, 0x8f, 0x68, 0x12, 0xe6, 0x4c, 0x16, 0x89, 0x92, 0x36, 0xe7, 0x76, 0x1d, 0x1f, 0x18,
	0x9a, 0x3c, 0x84, 0x3e, 0x56, 0x77, 0xc8, 0xe6, 0x73, 0x16, 0x29, 0x97, 0x82, 0x3d, 0xe4, 0xce,
	0x0c, 0x45, 0xc6, 0x70, 0x7b, 0xc9, 0xa2, 0x6b, 0x2a, 0xb8, 0x5c, 0x86, 0xe9, 0x3c, 0xa4, 0xba,
	0x54, 0x6d, 0x46, 0xde, 0x2a, 0xb7, 0x2e, 0xe6, 0x13, 0xbd, 0x41, 0x8e, 0x61, 0x8f, 0xc6, 0x31,
	0x37, 0x89, 0xba, 0x61, 0xb1, 0xc9, 0xd4, 0xbb, 0xeb, 0xdd, 0x2d, 0xa3, 0xd5, 0x35, 0xcb, 0x69,
	0x56, 0x28, 0x1e, 0x85, 0x34, 0x67, 0x54, 0xea, 0x06, 0xd5, 0x0d, 0x76, 0xd7, 0xfc, 0x04, 0xe9,
	0xd1, 0x3f, 0xa0, 0x7f, 0x9a, 0x4a, 0xe6, 0x42, 0x8f, 0x89, 0x16, 0xa7, 0xd2, 0x74, 0x1f, 0x2f,
	0xd0, 0x6b, 0x8c, 0x54, 0x99, 0x30, 0x35, 0xcd, 0x97, 0x78, 0xf4, 0x8b, 0x07, 0xbb, 0x4e, 0xf9,
	0xd2, 0x3a, 0xf9, 0x18, 0x1a, 0xf8, 0x9a, 0xe8, 0x33, 0x06, 0x47, 0x0f, 0x37, 0x93, 0xcb, 0x0a,
	0x8d, 0xed, 0xef, 0xd5, 0x2a, 0x63, 0x81, 0x16, 0xc7, 0x1e, 0x77, 0x43, 0x93, 0xc2, 0x7d, 0xc3,
	0x00, 0x34, 0xa8, 0x10, 0x5c, 0xb9, 0xcc, 0xc7, 0xf5, 0xe8, 0x6b, 0xe8, 0x55, 0xd4, 0x49, 0x0f,
	0xda, 0xef, 0xde, 0xfe, 0xfb, 0xed, 0xc5, 0x87, 0xb7, 0xc3, 0x3f, 0x90, 0x0e, 0x34, 0xce, 0x4e,
	0x8e, 0x9f, 0x0e, 0x3d, 0x5c, 0x4d, 0x71, 0x55, 0xc3, 0xd5, 0x04, 0x57, 0x75, 0x5c, 0xbd, 0x3e,
	0x3d, 0x7e, 0x3a, 0x6c, 0xe0, 0xea, 0x3c, 0x38, 0x7e, 0x3a, 0x6c, 0xea, 0xdd, 0x37, 0x93, 0xff,
	0x0e, 0x5b, 0xa4, 0x0d, 0xf5, 0xc9, 0xbb, 0x93, 0x61, 0x9b, 0xdc, 0x82, 0x9d, 0xc9, 0xc9, 0xd5,
	0xf4, 0xfd, 0xf4, 0xea, 0x7f, 0xe1, 0x24, 0x38, 0x9b, 0x0c, 0x3b, 0x28, 0x15, 0xbc, 0xb9, 0x3c,
	0x1b, 0x76, 0x47, 0x27, 0xb0, 0xeb, 0x5a, 0xe3, 0x59, 0xc2, 0xf4, 0xe3, 0xa3, 0x7b, 0xbd, 0xa1,
	0x6c, 0xa6, 0x97, 0x18, 0xaf, 0xa5, 0x5b, 0x8d, 0xbe, 0x81, 0x17, 0x18, 0x30, 0xfa, 0xb1, 0x0e,
	0x3b, 0x1b, 0xf5, 0x86, 0x45, 0xbf, 0x28, 0x0b, 0x05, 0x97, 0xd8, 0x3b, 0x74, 0xc7, 0x09, 0xcb,
	0x5e, 0xd0, 0xd6, 0x78, 0x1a, 0x93, 0x53, 0xe8, 0xbb, 0x10, 0xa0, 0x0b, 0xf4, 0xd9, 0x83, 0xa3,
	0x83, 0xcf, 0xd4, 0x71, 0x89, 0xb4, 0xa7, 0x37, 0xb4, 0xc8, 0x13, 0x68, 0x69, 0x27, 0x4b, 0xdd,
	0x97, 0x7b, 0x47, 0xc4, 0xe8, 0x57, 0x13, 0x22, 0xb0, 0x12, 0xe4, 0x11, 0x0c, 0x16, 0x79, 0xfa,
	0x7f, 0x75, 0x7d, 0xa9, 0xa8, 0x88, 0x69, 0x1e, 0xeb, 0xd2, 0xf2, 0x82, 0x2d, 0x96, 0x3c, 0x83,
	0xae, 0xbb, 0xba, 0xb4, 0xed, 0xe5, 0xae, 0x39, 0x76, 0xcb, 0x69, 0xc1, 0x5a, 0x4e, 0x77, 0x54,
	0xba, 0xcc, 0x12, 0xe6, 0xfa, 0xb9, 0x41, 0xe4, 0x70, 0x5d, 0xc8, 0x9d, 0xea, 0x51, 0x5b, 0xc9,
	0xb4, 0xae, 0x6f, 0xf3, 0xe8, 0xaa, 0x3c, 0x4d, 0xa4, 0x7e, 0xc3, 0xbd, 0xa0, 0xc4, 0xf8, 0x91,
	0x59, 0x42, 0xc5, 0x47, 0xa9, 0xdf, 0x71, 0x2f, 0xb0, 0xa8, 0xd2, 0xce, 0x7b, 0xf6, 0xe3, 0x1a,
	0x8d, 0x8e, 0xa1, 0x5f, 0xf5, 0xdd, 0x66, 0x9a, 0x01, 0xb4, 0xce, 0x83, 0x8b, 0x0f, 0x57, 0xff,
	0x1a, 0x7a, 0xa4, 0x0f, 0x1d, 0x97, 0x2d, 0xc3, 0xda, 0xe8, 0xe7, 0x3a, 0x74, 0x75, 0x5b, 0xd7,
	0x4d, 0xaa, 0x1a, 0x43, 0x6f, 0x33, 0x86, 0x5b, 0x1d, 0xb2, 0xf6, 0x49, 0x87, 0xc4, 0xf7, 0x27,
	0x52, 0xfc, 0xc6, 0x84, 0xb7, 0x13, 0x58, 0x54, 0x1d, 0x82, 0x1a, 0x9b, 0x43, 0xd0, 0x4b, 0x80,
	0x79, 0x92, 0x52, 0x15, 0xde, 0xd0, 0xc4, 0xbc, 0xa6, 0xbd, 0xa3, 0xfb, 0xc6, 0x65, 0xa5, 0x49,
	0xe3, 0x57, 0x28, 0xf1, 0x9e, 0x26, 0xf2, 0x4c, 0xa8, 0x7c, 0x15, 0x74, 0xe7, 0x0e, 0x93, 0x7f,
	0x42, 0x4f, 0xaa, 0x9c, 0x8b, 0x85, 0xd1, 0x37, 0xd1, 0x7b, 0xb0, 0xad, 0x7f, 0xa9, 0x45, 0xd6,
	0x07, 0x80, 0x2c, 0x09, 0xf2, 0x1c, 0x3a, 0x5c, 0xd8, 0xcf, 0xb7, 0xb5, 0xfa, 0xbd, 0x6d, 0xf5,
	0xa9, 0xa8, 0x7c, 0xbc, 0xcd, 0x0d, 0xda, 0xff, 0x3b, 0x0c, 0x36, 0xed, 0xc2, 0x7a, 0xf8, 0xc8,
	0x56, 0xae, 0x1e, 0x3e, 0xb2, 0xd5, 0xe7, 0x1b, 0xc4, 0x8b, 0xda, 0xdf, 0xbc, 0xfd, 0x97, 0xb0,
	0xbb, 0x65, 0xd5, 0x97, 0xd4, 0xbb, 0x55, 0xf5, 0x17, 0xd0, 0xaf, 0x5a, 0xf5, 0x25, 0xdd, 0x66,
	0x45, 0x77, 0xf4, 0x08, 0x3a, 0x67, 0x6e, 0xaa, 0xaa, 0x4e, 0x5c, 0xb6, 0x0d, 0x38, 0x3c, 0xfa,
	0xa6, 0x0e, 0x83, 0xf3, 0xa3, 0xff, 0x4c, 0xa4, 0x4c, 0x23, 0x5e, 0x8e, 0x7a, 0x66, 0xa8, 0xf3,
	0xaa, 0x43, 0xdd, 0x08, 0xfa, 0x4c, 0xdc, 0xf0, 0x3c, 0x15, 0x18, 0x52, 0x37, 0xf1, 0x6d, 0x70,
	0x38, 0xbe, 0x95, 0x43, 0xbb, 0x9b, 0xfc, 0x2a, 0x0c, 0x1a, 0x32, 0x67, 0x54, 0x15, 0xb9, 0x2d,
	0xed, 0x6e, 0x50, 0x62, 0x3c, 0x3f, 0x2b, 0x66, 0x09, 0x8f, 0xa8, 0x99, 0x1b, 0xcd, 0x80, 0xb0,
	0xc1, 0x91, 0x27, 0x95, 0x8b, 0x98, 0x30, 0x0e, 0x4c, 0x18, 0xdd, 0x55, 0x2b, 0xa3, 0xe4, 0xd6,
	0x84, 0x67, 0x9e, 0xd2, 0x8d, 0x09, 0xef, 0x2f, 0x30, 0x70, 0xd2, 0x61, 0x42, 0x67, 0x2c, 0xb1,
	0x4f, 0xe3, 0x8e, 0x63, 0x5f, 0x23, 0x89, 0x8f, 0xbe, 0xeb, 0x4e, 0xa1, 0x7e, 0x3f, 0xcc, 0x1b,
	0xb7, 0xd9, 0xb2, 0x9e, 0x41, 0x27, 0xcd, 0xf9, 0x82, 0x0b, 0x9a, 0xe8, 0x67, 0xb3, 0x77, 0xf4,
	0xc7, 0xb1, 0xf9, 0xa3, 0x32, 0x76, 0x7f, 0x54, 0x30, 0x41, 0x8b, 0x48, 0x05, 0xa5, 0xe0, 0x6f,
	0x55, 0xf8, 0xac, 0xa5, 0x55, 0x9e, 0xfd, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x42, 0x68, 0xc1,
	0x07, 0x0d, 0x00, 0x00,
}
