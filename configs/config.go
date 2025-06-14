package configs

import (
	"os"

	"github.com/spf13/viper"
)

func InitFmsConfig() {
	path, _ := os.Getwd()
	configPath := path + "/configs/fms.yaml"
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}

var Config FMSConfig

type FMSConfig struct {
	Basic Basic `mapstructure:"basic"`
	Log   Log   `mapstructure:"log"`
}

type Log struct {
	LogFile string `mapstructure:"logFile"`
	Level   string `mapstructure:"level"`
}

type Basic struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
	Debug   bool   `mapstructure:"debug"`
}

func (b *Basic) FullAddr() string {
	return b.Address + ":" + b.Port
}
