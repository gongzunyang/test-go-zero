package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovUserModel = (*customGovUserModel)(nil)

type (
	// GovUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovUserModel.
	GovUserModel interface {
		govUserModel
	}

	customGovUserModel struct {
		*defaultGovUserModel
	}
)

// NewGovUserModel returns a model for the database table.
func NewGovUserModel(conn sqlx.SqlConn) GovUserModel {
	return &customGovUserModel{
		defaultGovUserModel: newGovUserModel(conn),
	}
}
