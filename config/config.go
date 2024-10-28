package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Slice struct {
		Length           int `yaml:"Length"`
		MinimalValue     int `yaml:"MinimalValue"`
		MaximumValue     int `yaml:"MaximumValue"`
		AdditionalNumber int `yaml:"AdditionalNumber"`
		IndexForRemove   int `yaml:"IndexForRemove"`
	} `yaml:"Slice"`
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
	err := v.Unmarshal(&c)
	if err != nil {
		log.Println(err.Error())
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
