package UseCase

import (
	"reflect"
)

type TypeIdentifier struct{}

func NewTypeIdentifier() *TypeIdentifier {
	return &TypeIdentifier{}
}

func (t TypeIdentifier) IdentifyType(arg any) string {
	if arg == nil {
		return "nil"
	}
	return reflect.TypeOf(arg).String()
}
