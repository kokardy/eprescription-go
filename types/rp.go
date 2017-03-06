package epres

/*
Rp inteface.
 Medicines:薬品のリスト
 Amounts:用量のリスト.勧告にしたがって1回量で表現
 Usage:用法
 Comment:Rp（用法）に対するコメント
*/
type Rp interface {
	Versioner
	Medicines() []Medicine
	Usage() Usage
	Comment() string
}
