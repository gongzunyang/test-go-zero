package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovManagerClassifyModel = (*customGovManagerClassifyModel)(nil)

type (
	// GovManagerClassifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovManagerClassifyModel.
	GovManagerClassifyModel interface {
		govManagerClassifyModel
	}

	customGovManagerClassifyModel struct {
		*defaultGovManagerClassifyModel
	}
)

// NewGovManagerClassifyModel returns a model for the database table.
func NewGovManagerClassifyModel(conn sqlx.SqlConn) GovManagerClassifyModel {
	return &customGovManagerClassifyModel{
		defaultGovManagerClassifyModel: newGovManagerClassifyModel(conn),
	}
}
