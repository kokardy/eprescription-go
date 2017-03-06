package epres

//Hospitalは病院・診療所
type Hospital interface {
	Versioner
	Code() string
	Name() string
	Address() string
}
