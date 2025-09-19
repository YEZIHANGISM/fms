package database

import (
	"fms/configs"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	secretFile := "/etc/pg-secret/postgres-password"
	password, err := os.ReadFile(secretFile)
	if err != nil {
		panic(err)
	}
	formatPwd := strings.TrimSpace(string(password))

	dataSource := configs.Config.DB.DataSourceName(formatPwd)
	DB = sqlx.MustConnect("pgx", dataSource)
}
