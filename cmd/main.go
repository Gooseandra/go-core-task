package main

import (
	"fmt"
	"go-core-task/config"
	"go-core-task/internal/sliceAssembler/UseCase"
	"log"
)

func main() {
	log.Println("Starting...")
	log.Println("Loading config...")
	v, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	log.Println("Config loaded")

	sliceManipulate(*cfg)
}

func sliceManipulate(cfg config.Config) {
	sliceAssemblerUC := UseCase.NewSliceAssembler()

	originalSlice := sliceAssemblerUC.NewRandomSlice(cfg.Slice.Length, cfg.Slice.MinimalValue, cfg.Slice.MaximumValue)
	fmt.Println("\nOriginal slice: ")
	fmt.Println(originalSlice)

	evenSlice := sliceAssemblerUC.SliceExample(originalSlice)
	fmt.Println("\nEven slice: ")
	fmt.Println(evenSlice)

	sliceWithAdditionalNum := sliceAssemblerUC.AddElements(cfg.Slice.AdditionalNumber, originalSlice)
	fmt.Println("\nSlice with additional number ")
	fmt.Println(sliceWithAdditionalNum)

	copiedSlice := sliceAssemblerUC.CopySlice(originalSlice)
	fmt.Println("\nCopied slice: ")
	fmt.Println(copiedSlice)

	removedIndexSlice, err := sliceAssemblerUC.RemoveElement(cfg.Slice.IndexForRemove, originalSlice)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("\nSlice with removed index: ")
	fmt.Println(removedIndexSlice)
}
