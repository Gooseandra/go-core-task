package stringAssembler

type StringAssembler interface {
	ToString(arg any) (string, error)
	StringStream(args ...string) string
	RuneSlice(str string) []rune
	HexRunes(runes []rune) string
}
