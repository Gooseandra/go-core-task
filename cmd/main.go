package main

import (
	"fmt"
	"go-core-task/config"
	StringAssemblerUseCase "go-core-task/internal/stringAssembler/UseCase"
	TypeIdentifierUseCase "go-core-task/internal/typeIdentifier/UseCase"
	"log"
	"reflect"
)

func main() {
	log.Println("Starting...")
	log.Println("Loading config...")
	v, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err.Error())
	}
	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Println("Starting server")
		v, err = config.LoadConfig()
		if err != nil {
			log.Fatal("Cannot load config: ", err.Error())
		}
		log.Fatal("Config parse error: ", err.Error())
	}
	log.Println("Config loaded")

	TypeIdentifierUC := TypeIdentifierUseCase.NewTypeIdentifier()
	StringAssemblerUC := StringAssemblerUseCase.NewStringAssembler(*cfg)

	values := reflect.ValueOf(cfg.Values)

	fmt.Println("\nTypes of variables: ")
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i).Interface()
		fmt.Println(TypeIdentifierUC.IdentifyType(field))
	}

	fmt.Println("\nAll data in one string:")
	var strings []string
	for i := 0; i < values.NumField(); i++ {
		str, err := StringAssemblerUC.ToString(values.Field(i))
		if err != nil {
			log.Println(err.Error())
		}
		strings = append(strings, str)
	}
	stringData := StringAssemblerUC.StringStream(strings...)
	fmt.Println(stringData)

	fmt.Println("\nHexed string: ")
	stringWithSalt := StringAssemblerUC.AddSalt(stringData, len(stringData)/2)
	runes := StringAssemblerUC.RuneSlice(stringWithSalt)
	fmt.Println(StringAssemblerUC.HexRunes(runes))
}
