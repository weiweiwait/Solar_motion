package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	Es *Es `yaml:"es"`
}

type Es struct {
	EsHost  string `yaml:"esHost"`
	EsPort  string `yaml:"esPort"`
	EsIndex string `yaml:"esIndex"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/locales")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
