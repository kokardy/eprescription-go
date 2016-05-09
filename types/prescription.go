package epres

import (
	"fmt"
)

/*
Numberは有理数型.
1/3個などに対応できるように
*/
type Number struct {
	Child, Mother int
}

//Stringは分数表記
func (n Number) String() string {
	return fmt.Sprintf("%d/%d", n.Child, n.Mother)
}

//Floatは小数表記
func (n Number) Float() float64 {
	return (float64)(n.Child) / (float64)(n.Mother)
}

/*
Unitは単位構造体.
単位はコード化するべき.
さもないと投与量チェックや点数計算などで再利用しづらくなる
*/
type Unit struct {
	Code, Name string
}

/*
Amountは投与量を示す有理数型と単位構造体からなる.
勧告にしたがって倍散などでは秤取量(製剤量)で記載
かつ1回量表記
*/
type Amount struct {
	Dose []Number
	Unit Unit
}

/*
Medicineは薬品を表現する構造体.
YJコードが同じ薬品を区別するためにHOTコードも用意.
想定されるバターンは

*銘柄指定処方
	YjCode or HotCode指定
*銘柄指定かつ後発変更可
	GenericOK = true
*一般名処方
	GenericCode指定


Name:名称変更もあるのである程度好きな名前にしてもいい
ただ、他のコードと矛盾するような名前は混乱を呼ぶ
(秤量と予包など)

*/
type Medicine struct {
	YjCode, HotCode, GenericCode, Name string
	GenericOK                          bool
	Amount                             Amount
	Comment                            string
}

/*
Usageは用法構造体.
標準用法マスタのフィールドは作るが、
標準用法マスタでは表現できないものがある前提
タイミングのリストで表現
*/
type Usage struct {
	MedisCode, MedisName string
	Times                []string
}

/*
Rp構造体.
 Medicines:薬品のリスト
 Amounts:用量のリスト.勧告にしたがって1回量で表現
 Usage:用法
 Comment:Rp（用法）に対するコメント
*/
type Rp struct {
	Medicines []Medicine
	Usage     Usage
	Comment   string
}

/*
Patientは患者構造体.
	Address:麻薬処方箋の場合必要
*/
type Patient struct {
	Id, Name, Hoken, Address string
}

//Doctorは医師
type Doctor struct {
	Ka, Name, MayakuLicense string
}

//Hospitalは病院・診療所
type Hospital struct {
	Code, Name string
}

/*
Prescriptionは処方箋を表す
処方箋はRpの集合で表現
処方箋全体としてのコメントや
患者情報・医療機関情報・処方医などの処方箋の必要項目
*/
type Prescription struct {
	Pt              Patient
	Dr              Doctor
	Rps             []Rp
	Comment         string
	ContainsMayaku  bool
	GenericOK       bool
	GenericNGReason string
}
