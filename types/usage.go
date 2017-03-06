package epres

/*
Usageは用法interface.
標準用法マスタのフィールドは作るが、
標準用法マスタでは表現できないものがある前提
タイミングのリストで表現
*/
type Usage interface {
	Versioner
	MedisCode() string
	MedisName() string
	Times() []string
}
