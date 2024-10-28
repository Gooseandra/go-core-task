package UseCase

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-core-task/config"
	"strings"
)

type StringAssembler struct {
	cfg config.Config
}

func NewStringAssembler(cfg config.Config) *StringAssembler {
	return &StringAssembler{cfg: cfg}
}

func (s StringAssembler) ToString(arg any) (string, error) {
	str := fmt.Sprintf("%v", arg)
	return str, nil
}

func (s StringAssembler) StringStream(args ...string) string {
	var str strings.Builder
	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}

func (s StringAssembler) RuneSlice(str string) []rune {
	return []rune(str)
}

func (s StringAssembler) AddSalt(str string, position int) string {
	return str[:position] + s.cfg.Hex.Salt + str[position:]
}

func (s StringAssembler) HexRunes(runes []rune) string {
	byteData := []byte(string(runes))
	hash := sha256.Sum256(byteData)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
