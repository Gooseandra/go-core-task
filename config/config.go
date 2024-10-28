package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Config struct {
	Values struct {
		IntDecimal     int       `yaml:"int_decimal"`
		IntOctal       int       `yaml:"int_octal"`
		IntHexadecimal int       `yaml:"int_hexadecimal"`
		Float          float64   `yaml:"float"`
		Str            string    `yaml:"str"`
		Boolean        bool      `yaml:"boolean"`
		Complex        complex64 `yaml:"complex"`
	} `yaml:"Values"`
	Hex struct {
		Salt string `yaml:"salt"`
	} `yaml:"Hex"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	var err error
	c.Values.IntDecimal = v.GetInt("Values.int_decimal")

	if c.Values.IntOctal, err = parseOctalString(v.GetString("Values.int_octal")); err != nil {
		return nil, fmt.Errorf("error parsing IntOctal: %w", err)
	}

	if c.Values.IntHexadecimal, err = parseHexString(v.GetString("Values.int_hexadecimal")); err != nil {
		return nil, fmt.Errorf("error parsing IntHexadecimal: %w", err)
	}

	c.Values.Float = v.GetFloat64("Values.float")
	c.Values.Str = v.GetString("Values.str")
	c.Values.Boolean = v.GetBool("Values.boolean")

	complexStr := v.GetString("Values.complex")
	if c.Values.Complex, err = parseComplex(complexStr); err != nil {
		return nil, fmt.Errorf("error parsing Complex: %w", err)
	}

	c.Hex.Salt = v.GetString("Hex.salt") // Добавлено получение Salt

	return &c, nil
}

func parseOctalString(octalStr string) (int, error) {
	octalValue, err := strconv.ParseInt(octalStr, 8, 32)
	return int(octalValue), err
}

func parseHexString(hexStr string) (int, error) {
	hexValue, err := strconv.ParseInt(hexStr, 16, 32)
	return int(hexValue), err
}

func parseComplex(s string) (complex64, error) {
	s = strings.ReplaceAll(s, " ", "")
	parts := strings.Split(s, "+")
	if len(parts) != 2 || !strings.HasSuffix(parts[1], "i") {
		return 0, errors.New("invalid complex format")
	}

	realPart, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return 0, err
	}
	imaginaryPart, err := strconv.ParseFloat(parts[1][:len(parts[1])-1], 32)
	if err != nil {
		return 0, err
	}

	return complex(float32(realPart), float32(imaginaryPart)), nil
}
