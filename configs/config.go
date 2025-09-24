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
	Type    string `mapstructure:"type"`
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

func (db *DB) Dsn(password *string) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Dbname, db.SSLMode)
	if password != nil {
		dsn += fmt.Sprintf(" password=%s", *password)
	}
	return dsn
}
