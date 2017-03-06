package epres

type Generic interface {
	GenericOK() bool
	GenericNGReason() string
}
