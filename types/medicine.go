package epres

/*
Medicineは薬品を表現するinterface.
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
type Medicine interface {
	Version() string
	YjCode() string
	HotCode() string
	GenericCode() string
	Name() string
	GenericOK() bool
	Amount() Amount
	Comment() string
}
