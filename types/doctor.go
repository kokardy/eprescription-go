package epres

//Doctorは医師
type Doctor interface {
	Version() string
	Ka() string
	Name() string
	MayakuLicense() string
}
