package models

import (
	"fms/database"
	fmslog "fms/log"
)

type Bill struct {
	ID uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
}

func (Bill) TableName() string {
	return "bills"
}

func (b *Bill) ListBills(filter map[string]interface{}) ([]Bill, error) {
	var bills []Bill
	if err := database.DB.Find(&bills).Where(filter).Error; err != nil {
		fmslog.SLogger.Errorf("failed to list bills: %v", err)
		return nil, err
	}
	return bills, nil
}
