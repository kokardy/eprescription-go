package epres

/*
Prescriptionは処方箋を表す
処方内容はRpの集合で表現
処方箋全体としてのコメントや
患者情報・医療機関情報・処方医などの処方箋の必要項目
*/
type Prescription interface {
	Versioner
	GenericFlager //後発切り替え可不可

	Hospital() Hospital
	Pt() Patient
	Dr() Doctor
	Rps() []Rp
	Comment() string
	ContainsMayaku() bool
}
