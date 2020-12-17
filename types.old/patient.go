package epres

type Patient interface {
	Versioner
	IDer
	Birthdayer

	Name() string
	Hoken() string
	Address() string
	Age() Duration
}
