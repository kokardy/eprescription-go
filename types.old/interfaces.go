package epres

type Versioner interface {
	Version() string
}

type IDer interface {
	ID() string
}

type GenericFlager interface {
	GenericOK() bool
	GenericNGReason() string
}
