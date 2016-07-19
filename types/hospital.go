package epres

//Hospitalは病院・診療所
type Hospital interface {
	Version() string
	Code() string
	Name() string
}
