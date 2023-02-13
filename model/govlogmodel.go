package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovLogModel = (*customGovLogModel)(nil)

type (
	// GovLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovLogModel.
	GovLogModel interface {
		govLogModel
	}

	customGovLogModel struct {
		*defaultGovLogModel
	}
)

// NewGovLogModel returns a model for the database table.
func NewGovLogModel(conn sqlx.SqlConn) GovLogModel {
	return &customGovLogModel{
		defaultGovLogModel: newGovLogModel(conn),
	}
}
