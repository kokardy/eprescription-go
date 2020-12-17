package epres

//PatientRoleID is ClinicalDocument/recordTarget/patientRole/id
type PatientRoleID struct {
	//患者ID発行主体のOID
	Root `xml:"root,attr"`
	//患者ID
	Extension `xml:"extension,attr"`
}

//PostalCode 郵便番号
type PostalCode string

//StreetAddressLine 住所
type StreetAddressLine string

//Address 患者住所、医療機関住所
type Address struct {
	PostalCode        `xml:"postalCode"`
	StreetAddressLine `xml:"streetAddressLine"`
}

//PatientAddress ClinicalDocument/recordTarget/patientRole/addr
type PatientAddress Address

//NameUse is ClinicalDocument/recordTarget/patient/name/@use
type NameUse string

const (
	//NameUseKanji 漢字
	NameUseKanji NameUse = "IDE"
	//NameUseKana カナ
	NameUseKana NameUse = "SYL"
)

//Name is ClinicalDocument/recordTarget/patient/name
type Name struct {
	Use NameUse `xml:"use,attr"`
	//名字
	Family string `xml:"family"`
	//名前
	Given string `xml:"given"`
}

//PatientNameKanji 漢字氏名
type PatientNameKanji Name

//PatientNameKana カナ氏名
type PatientNameKana Name

//PatientNames 漢字とカナの氏名
//NameKanji NameKana のスライス
type PatientNames []Name

//GenderCode  is ClinicalDocument/recordTarget/patientRole/patient/administrativeGenderCode/@code
type GenderCode AttCode

//GenderCodeSystem is ClinicalDocument/recordTarget/patientRole/patient/administrativeGenderCode/@codeSystem
type GenderCodeSystem CodeSystem

//AdministrativeGenderCode is ClinicalDocument/recordTarget/patientRole/patient/administrativeGenderCode
type AdministrativeGenderCode struct {
	GenderCode       `xml:"code,attr"`
	GenderCodeSystem `xml:"codeSystem,attr"`
}

//BirthTime is ClinicalDocument/recordTarget/patientRole/patient/birthTime
type BirthTime struct {
	Value string `xml:"value,attr"`
}

//Patient ClinicalDocument/recordTarget/patient
type Patient struct {
	PatientNames             `xml:"name"`
	AdministrativeGenderCode `xml:"administrativeGenderCode"`
	BirthTime                `xml:"birthTime"`
}

//PatientRole is ClinicalDocument/recordTarget/patientRole
type PatientRole struct {
	PatientRoleID  `xml:"id"`
	PatientAddress `xml:"addr"`
	Patient        `xml:"patient"`
}

//RecordTarget is ClinicalDocument/recordTarget
//患者情報
type RecordTarget struct {
	PatientRole `xml:"patientRole"`
}
