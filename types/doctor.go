package epres

//Doctorは医師
type Doctor interface {
	Versioner
	IDer
	Ka() string
	Name() string
	MayakuLicense() string
}
