package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovLocationModel = (*customGovLocationModel)(nil)

type (
	// GovLocationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovLocationModel.
	GovLocationModel interface {
		govLocationModel
	}

	customGovLocationModel struct {
		*defaultGovLocationModel
	}
)

// NewGovLocationModel returns a model for the database table.
func NewGovLocationModel(conn sqlx.SqlConn) GovLocationModel {
	return &customGovLocationModel{
		defaultGovLocationModel: newGovLocationModel(conn),
	}
}
