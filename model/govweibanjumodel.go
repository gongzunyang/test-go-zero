package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovWeibanjuModel = (*customGovWeibanjuModel)(nil)

type (
	// GovWeibanjuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovWeibanjuModel.
	GovWeibanjuModel interface {
		govWeibanjuModel
	}

	customGovWeibanjuModel struct {
		*defaultGovWeibanjuModel
	}
)

// NewGovWeibanjuModel returns a model for the database table.
func NewGovWeibanjuModel(conn sqlx.SqlConn) GovWeibanjuModel {
	return &customGovWeibanjuModel{
		defaultGovWeibanjuModel: newGovWeibanjuModel(conn),
	}
}
