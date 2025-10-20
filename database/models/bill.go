package models

import (
	"context"
	"fms/database"
	fmslog "fms/log"
	"time"
)

type Bill struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	BillbookID int       `gorm:"column:billbook_id;not null" json:"billbook_id"`
	CategoryID int       `gorm:"column:category_id;not null" json:"category_id"`
	AssetOutID int       `gorm:"column:asset_out_id;not null" json:"asset_out_id"`
	AssetInID  int       `gorm:"column:asset_in_id;not null" json:"asset_in_id"`
	DebtID     int       `gorm:"column:debt_id;not null" json:"debt_id"`
	BillDate   time.Time `gorm:"column:bill_date;not null" json:"bill_date"`
	Type       int       `gorm:"column:type;not null" json:"type"`
	Amount     float32   `gorm:"column:amount;not null" json:"amount"`
	Remark     string    `gorm:"column:remark;type:text" json:"remark"`
}

func (Bill) TableName() string {
	return "bill"
}

func (b *Bill) ListBills(filter map[string]interface{}) ([]Bill, error) {
	var bills []Bill
	if err := database.DB.Find(&bills).Where(filter).Error; err != nil {
		fmslog.SLogger.Errorf("failed to list bills: %v", err)
		return nil, err
	}
	return bills, nil
}

func (b *Bill) CreateBill(ctx context.Context, record *Bill) error {
	if err := database.DB.Create(record).Error; err != nil {
		fmslog.SLogger.Errorf("failed to create bill: %v", err)
		return err
	}
	return nil
}
