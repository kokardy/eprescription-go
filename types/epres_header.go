package epres

//Root is id/@root
type Root string

//Extension is id/@extension
type Extension string

//AttCode is code/@code
type AttCode string

//CodeSystem is code/@codeSystem
type CodeSystem string

//RealmCode 固定 "JP"
type RealmCode struct {
	Code AttCode `xml:"code,attr"`
}

const realmCodeAttCode AttCode = "JP"

//TypeIDRoot typeIdのroot
//固定:"2.16.840.1.113883.1.3"
type TypeIDRoot Root

const typeIDRoot TypeIDRoot = "2.16.840.1.113883.1.3"

//TypeIDExtension typeIdのextension
//固定:"POCD_HD000040"
type TypeIDExtension Extension

const typeIDExtension = "POCD_HD000040"

//TypeID HL7 CDA
type TypeID struct {
	TypeIDRoot      `xml:"root,attr"`
	TypeIDExtension `xml:"extension,attr"`
}

//PresRoot 処方箋ID@root 固定:"1.2.392.100495.20.3.11"
//PresRoot is ClinicalDocument/id/@root
type PresRoot Root

const presRoot PresRoot = "1.2.392.100495.20.3.11"

//PresExtension is ClinicalDocument/id/@extension
type PresExtension Extension

//PrescriptionID 処方箋ID ClinicalDocument/id
type PrescriptionID struct {
	PresRoot      `xml:"root,attr"`
	PresExtension `xml:"extension,attr"`
}

//DocCode is ClinialDocument/code/@code
type DocCode AttCode

//DocCodeSystem is ClinialDocument/code/@codeSystem
type DocCodeSystem CodeSystem

//DocumentCode is ClinicalDocument/code
type DocumentCode struct {
	DocCode       `xml:"code,attr"`
	DocCodeSystem `xml:"codeSystem,attr"`
}

//Title is ClinicalDocument/title/text()
type Title string

//EffectiveTime is ClinicalDocument/effectiveTime
type EffectiveTime struct {
	Value string `xml:"value,attr"`
}

//ConfidentialityAttCode is ClinicalDocument/confidentialCode/@code
type ConfidentialityAttCode AttCode

const confidentialityCode ConfidentialityAttCode = "N"

//ConfidentialityCodeSyStem is ClinicalDocument/confidentialCode/@codeSystem
//固定:"2.16.840.1.113883.5.25"
type ConfidentialityCodeSyStem CodeSystem

const confidentialityCodeSyStem ConfidentialityCodeSyStem = "2.16.840.1.113883.5.25"

//ConfidentialityCode is ClinicalDocument/confidentialityCode
//守備レベルコード
type ConfidentialityCode struct {
	ConfidentialityAttCode    `xml:"code,attr"`
	ConfidentialityCodeSyStem `xml:"codeSystem,attr"`
}

//VersionNumberValue is ClinicalDocument/versionNumber/@value
//固定:"100" v.1.00を示す
type VersionNumberValue string

const versionNumberValue VersionNumberValue = "100"

//VersionNumber is ClinicalDocument/versionNumber
type VersionNumber struct {
	Value string `xml:"value,attr"`
}

//CustodianIDNullFlavor is ClinicalDocument/custodian/assignedCustodian/representedCustodianOrganization/id/@nullFlavor
type CustodianIDNullFlavor string

//CustodianID is ClinicalDocument/custodian/assignedCustodian/representedCustodianOrganization/id
type CustodianID struct {
	CustodianIDNullFlavor `xml:"nullFlavor,attr"`
}

//RepresentedCustodianOrganization is ClinicalDocument/custodian/assignedCustodian/representedCustodianOrganization/
type RepresentedCustodianOrganization struct {
	CustodianID `xml:"id,attr"`
}

//AssignedCustodian is ClinicalDocument/custodian/assignedCustodian
type AssignedCustodian struct {
	RepresentedCustodianOrganization `xml:"representedCustodianOrganization"`
}

//Custodian is ClinicalDocument/custodian
type Custodian struct {
	AssignedCustodian `xml:"assignedCustodian"`
}

//Header 電子処方箋のヘッダ部
type Header struct {
	RealmCode           `xml:"realmCode"`
	TypeID              `xml:"typeId"`
	PrescriptionID      `xml:"id"`
	DocumentCode        `xml:"code"`
	Title               `xml:"title"`
	EffectiveTime       `xml:"effectiveTime"`
	ConfidentialityCode `xml:"confidentialityCode"`
	VersionNumber       `xml:"versionNumber"`
	RecordTarget        `xml:"recordTarget"`
	Author              `xml:"author"`
	Custodian           `xml:"custodian"`
	Component           `xml:"component"`
}
