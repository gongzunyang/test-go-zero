package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovTagModel = (*customGovTagModel)(nil)

type (
	// GovTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovTagModel.
	GovTagModel interface {
		govTagModel
	}

	customGovTagModel struct {
		*defaultGovTagModel
	}
)

// NewGovTagModel returns a model for the database table.
func NewGovTagModel(conn sqlx.SqlConn) GovTagModel {
	return &customGovTagModel{
		defaultGovTagModel: newGovTagModel(conn),
	}
}
