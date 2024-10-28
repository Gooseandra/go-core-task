package typeIdentifier

type TypeIdentifier interface {
	IdentifyType(arg any) string
}
