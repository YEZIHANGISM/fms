package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fms/service/fms/model"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 数据库假数据生成工具

var (
	// i 与 d可以同时指定
	// i and d => 生成前先删除旧数据
	// !i and d => 删除旧数据
	// i and !d => 只生成新数据，不删除旧数据
	i = flag.Bool("insert", false, "生成mock数据")
	d = flag.Bool("delete", false, "清空旧数据")
	n = flag.Int("n", 100, "数据条数")
	t = flag.String("table", "", "指定表插入mock数据，如果不指定，默认所有表都生成mock数据")
)

var (
	tables = []string{
		"bill",
		"billbook",
		"asset",
		"asset_category",
		"category",
		"debt",
	}
)

var (
	conn      = sqlx.NewMysql("root:mysql@tcp(host.docker.internal:3306)/fms?parseTime=true")
	bill      = model.NewBillModel(conn)
	billbook  = model.NewBillbookModel(conn)
	asset     = model.NewAssetModel(conn)
	assetCate = model.NewAssetCategoryModel(conn)
	category  = model.NewCategoryModel(conn)
	debt      = model.NewDebtModel(conn)
)

var (
	err = errors.New("mock failed")
	ctx = context.Background()
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	flag.Parse()

	if *i {
		// 生成mock数据
		genMockData()
		return
	}

	if *d {
		clearMockData()
	}
}

func randStr() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// 生成所有表数据，需要考虑表的关联关系
func genAllTableLogicData() {
	// 1. debt
	debtData := &model.Debt{
		Id:        uuid.NewString(),
		DebtType:  randStr(),
		Object:    randStr(),
		CountInto: sql.NullBool{Bool: true, Valid: true},
		Ended:     sql.NullBool{Bool: true, Valid: true},
	}
	insertDebt(debtData)

	// 2. billbook
	bbData := &model.Billbook{
		Id:   uuid.NewString(),
		Name: randStr(),
	}
	insertBillbook(bbData)

	// 3. asset_category
	acData := &model.AssetCategory{
		Id:   uuid.NewString(),
		Name: randStr(),
	}
	insertAssetCategory(acData)

	// 4. asset with ac.Id
	assetData := &model.Asset{
		Id:          uuid.NewString(),
		Name:        randStr(),
		AssetCateId: acData.Id,
		Hide:        sql.NullBool{Bool: true, Valid: true},
		CountInto:   sql.NullBool{Bool: true, Valid: true},
		Balance:     100000,
	}
	insertAsset(assetData)

	// 5. category with bb.Id
	categoryData := &model.Category{
		Id:         uuid.NewString(),
		ParentId:   sql.NullString{String: "", Valid: false},
		Name:       randStr(),
		Level:      1,
		BillbookId: bbData.Id,
	}
	insertCategory(categoryData)

	// 6. bill with all exists table's Id
	billData := &model.Bill{
		Id:         uuid.NewString(),
		CategoryId: categoryData.Id,
		BillbookId: bbData.Id,
		AssetOutId: sql.NullString{String: uuid.NewString(), Valid: true},
		AssetInId:  sql.NullString{String: uuid.NewString(), Valid: true},
		DebtId:     sql.NullString{String: uuid.NewString(), Valid: true},
		BillType:   randStr(),
		Amount:     100,
		Remark:     sql.NullString{String: uuid.NewString(), Valid: true},
	}
	insertBill(billData)

}

func genMockData() {
	if *d {
		// 清空旧数据
		clearMockData()
	}

	switch *t {
	case "":
		// 生成所有表数据
		for j := 0; j < *n; j++ {
			genAllTableLogicData()
		}
	case "debt":
		for j := 0; j < *n; j++ {
			data := &model.Debt{
				Id:        uuid.NewString(),
				DebtType:  randStr(),
				Object:    randStr(),
				CountInto: sql.NullBool{Bool: true, Valid: true},
				Ended:     sql.NullBool{Bool: true, Valid: true},
			}
			insertDebt(data)
		}
	case "bill":
		for j := 0; j < *n; j++ {
			data := &model.Bill{
				Id:         uuid.NewString(),
				CategoryId: uuid.NewString(),
				BillbookId: uuid.NewString(),
				AssetOutId: sql.NullString{String: uuid.NewString(), Valid: true},
				AssetInId:  sql.NullString{String: uuid.NewString(), Valid: true},
				DebtId:     sql.NullString{String: uuid.NewString(), Valid: true},
				BillType:   randStr(),
				Amount:     100,
				Remark:     sql.NullString{String: uuid.NewString(), Valid: true},
				BillTime:   time.Now().UTC(),
			}
			insertBill(data)
		}
	case "billbook":
		for j := 0; j < *n; j++ {
			data := &model.Billbook{
				Id:   uuid.NewString(),
				Name: randStr(),
			}
			insertBillbook(data)
		}
	case "category":
		for j := 0; j < *n; j++ {
			data := &model.Category{
				Id:         uuid.NewString(),
				ParentId:   sql.NullString{String: "", Valid: false},
				Name:       randStr(),
				Level:      1,
				BillbookId: uuid.NewString(),
			}
			insertCategory(data)
		}
	case "asset":
		for j := 0; j < *n; j++ {
			data := &model.Asset{
				Id:          uuid.NewString(),
				Name:        randStr(),
				AssetCateId: uuid.NewString(),
				Hide:        sql.NullBool{Bool: true, Valid: true},
				CountInto:   sql.NullBool{Bool: true, Valid: true},
				Balance:     100000,
			}
			insertAsset(data)
		}
	case "asset_category":
		for j := 0; j < *n; j++ {
			data := &model.AssetCategory{
				Id:   uuid.NewString(),
				Name: randStr(),
			}
			insertAssetCategory(data)
		}
	}
}

func insertDebt(data *model.Debt) {
	_, err = debt.Insert(ctx, data)
	checkErr(err)
}

func insertBill(data *model.Bill) {
	_, err = bill.Insert(ctx, data)
	checkErr(err)
}

func insertBillbook(data *model.Billbook) {
	_, err = billbook.Insert(ctx, data)
	checkErr(err)
}

func insertCategory(data *model.Category) {
	_, err = category.Insert(ctx, data)
	checkErr(err)
}

func insertAsset(data *model.Asset) {
	_, err = asset.Insert(ctx, data)
	checkErr(err)
}

func insertAssetCategory(data *model.AssetCategory) {
	_, err = assetCate.Insert(ctx, data)
	checkErr(err)
}

func clearMockData() {
	if *t != "" {
		// 指定表删除数据
		delete(*t)
		return
	}
	// 删除所有表数据
	for _, t := range tables {
		delete(t)
	}
}

func delete(t string) {
	query := fmt.Sprintf("DELETE FROM %s", t)
	session, err := conn.Prepare(query)
	checkErr(err)

	_, err = session.Exec()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
