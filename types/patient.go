package epres

import (
	"time"
)

type Patient interface {
	Version() string
	Id() string
	Name() string
	Hoken() string
	Address() string
	Birthday() time.Time
	Age() string
}
