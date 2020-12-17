package epres

//TimeType is ClinicalDocument/author/time/@xi:type
type TimeType string

const timeType TimeType = "IVL_TS"

//Date is ClinicalDocument/author/time/low or high
type Date struct {
	//"YYYYMMDD"
	Value string `xml:"value,attr"`
}

//Time is ClinicalDocument/author/time
type Time struct {
	//固定値 const timeType="IVL_TS"
	TimeType `xml:"xi:type,attr"`
	//交付年月日
	Low Date `xml:"low"`
	//処方箋有効期限
	High Date `xml:"high"`
}

//AuthorID is ClinicalDocument/author/assignedAuthor/id
type AuthorID struct {
	Root      `xml:"root,attr"`
	Extension `xml:"extension,attr"`
}

//DoctorID is ClinicalDocument/author/assignedAuthor/id
type DoctorID AuthorID

//NarcoticLisenceID is ClinicalDocument/author/assignedAuthor/id
type NarcoticLisenceID AuthorID

//AuthorIDs is DoctorID and NarcoticLicenseID
//ClinicalDocument/author/assignedAuthor/id
type AuthorIDs []AuthorID

//DoctorNames 漢字氏名とカナ氏名のslice
type DoctorNames []Name

//AssignedParson is ClinicalDocument/author/assignedAuthor/assignedParson
type AssignedParson struct {
	DoctorNames `xml:"name"`
}

//OrganizationID is prefID,HospitalTypeID, HospitalID
type OrganizationID struct {
	Root      `xml:"root,attr"`
	Extension `xml:"extension,attr"`
}

//PrefID 都道府県コード
type PrefID OrganizationID

//固定:"1.2.392.100495.20.3.21"
const prefIDRoot Root = "1.2.392.100495.20.3.21"

//HospitalTypeID 医療機関種別
type HospitalTypeID OrganizationID

const (
	//固定:"1.2.392.100495.20.3.22"
	hospitalTypeIDRoot           = "1.2.392.100495.20.3.22"
	hospitalTypeExtenxionMedical = "1"
	hospitalTypeExtenxionDental  = "3"
)

//HospitalID 医療機関コード
type HospitalID OrganizationID

//固定:"1.2.392.100495.20.3.23"
const hospitalIDRoot = "1.2.392.100495.20.3.23"

//OrganizationIDs is [prefID,HospitalTypeID, HospitalID]
type OrganizationIDs []OrganizationID

//OrganizationName is ClinicalDocument/author/representedOrganization/name
type OrganizationName struct {
	NameUse `xml:"use,attr"`
	Name    `xml:"text()"`
}

//Telecom is ClinicalDocument/author/representedOrganization/telecom
type Telecom struct {
	//Value should be prefixed "tel:" or "fax:" or ...
	Value string `xml:"value,attr"`
}

//OarganizationAddress is ClinicalDocument/author/representedOrganization/addr
type OarganizationAddress Address

//DisplayName is ClinicalDocument/author/representedOrganization/asOrganizationPartOf/code/@displayName
//診療科
type DisplayName string

//OrganizationPartCode is ClinicalDocument/author/representedOrganization/asOrganizationPartOf/code
type OrganizationPartCode struct {
	AttCode     `xml:"code,attr"`
	CodeSystem  `xml:"codeSystem,attr"`
	DisplayName `xml:"displayName,attr"`
}

//AsOarganizationPartOf  is ClinicalDocument/author/representedOrganization/asOrganizationPartOf
type AsOarganizationPartOf struct {
	OrganizationPartCode `xml:"code"`
}

//AuthorRepresentedOrganization is ClinicalDocument/author/representedOrganization
type AuthorRepresentedOrganization struct {
	OrganizationIDs       `xml:"id"`
	OrganizationName      `xml:"name"`
	Telecom               `xml:"telecom"`
	OarganizationAddress  `xml:"addr"`
	AsOarganizationPartOf `xml:"asOrganizationPartOf"`
}

//AssignedAuthor is ClinicalDocument/author/assignedAuthor
type AssignedAuthor struct {
	//[DoctorID, NarcoticLicenseID]
	AuthorIDs                     `xml:"id"`
	AssignedParson                `xml:"assignedParson"`
	AuthorRepresentedOrganization `xml:"representedOrganization"`
}

//Author is ClinicalDocument/author
//処方医
type Author struct {
	Time           `xml:"time"`
	AssignedAuthor `xml:"assignedAuthor"`
}
