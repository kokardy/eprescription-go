package epres

//Doctorは医師
type Doctor interface {
	Versioner
	Ka() string
	Name() string
	MayakuLicense() string
}
