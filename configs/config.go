package configs

import (
	"fmt"
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
	DB    DB    `mapstructure:"database"`
}

type DB struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Dbname  string `mapstructure:"dbname"`
	SSLMode string `mapstructure:"sslmode"`
}

type Log struct {
	OutputPaths []string `mapstructure:"outputPaths"`
	Level       string   `mapstructure:"level"`
	Encoding    string   `mapstructure:"encoding"`
}

type Basic struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
	Debug   bool   `mapstructure:"debug"`
}

func (b *Basic) FullAddr() string {
	return b.Address + ":" + b.Port
}

func (d *DB) DataSourceName(password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User, password, d.Host, d.Port, d.Dbname, d.SSLMode)
}
