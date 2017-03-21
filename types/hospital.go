package epres

//Hospitalは病院・診療所
type Hospital interface {
	Versioner
	IDer
	Name() string
	Address() string
}
