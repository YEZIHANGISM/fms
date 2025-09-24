package database

import (
	"fms/configs"
	"os"
	"strings"

	fmslog "fms/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	password := getPassword()
	dsn := configs.Config.DB.Dsn(&password)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmslog.SLogger.Fatalf("failed to connect database: %v", err)
		panic(err)
	}
}

func getPassword() string {
	secret := "/etc/pg-secret/postgres-password"
	password, err := os.ReadFile(secret)
	if err != nil {
		fmslog.SLogger.Fatalf("failed to read db password from secret: %v", err)
		panic(err)
	}
	fmtPwd := strings.TrimSpace(string(password))
	return fmtPwd
}
