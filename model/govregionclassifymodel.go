package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovRegionClassifyModel = (*customGovRegionClassifyModel)(nil)

type (
	// GovRegionClassifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovRegionClassifyModel.
	GovRegionClassifyModel interface {
		govRegionClassifyModel
	}

	customGovRegionClassifyModel struct {
		*defaultGovRegionClassifyModel
	}
)

// NewGovRegionClassifyModel returns a model for the database table.
func NewGovRegionClassifyModel(conn sqlx.SqlConn) GovRegionClassifyModel {
	return &customGovRegionClassifyModel{
		defaultGovRegionClassifyModel: newGovRegionClassifyModel(conn),
	}
}
