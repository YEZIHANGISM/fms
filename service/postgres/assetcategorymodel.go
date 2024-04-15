package postgres

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AssetCategoryModel = (*customAssetCategoryModel)(nil)

type (
	// AssetCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetCategoryModel.
	AssetCategoryModel interface {
		assetCategoryModel
	}

	customAssetCategoryModel struct {
		*defaultAssetCategoryModel
	}
)

// NewAssetCategoryModel returns a model for the database table.
func NewAssetCategoryModel(conn sqlx.SqlConn) AssetCategoryModel {
	return &customAssetCategoryModel{
		defaultAssetCategoryModel: newAssetCategoryModel(conn),
	}
}
