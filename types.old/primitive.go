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
