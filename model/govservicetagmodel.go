package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovServiceTagModel = (*customGovServiceTagModel)(nil)

type (
	// GovServiceTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovServiceTagModel.
	GovServiceTagModel interface {
		govServiceTagModel
	}

	customGovServiceTagModel struct {
		*defaultGovServiceTagModel
	}
)

// NewGovServiceTagModel returns a model for the database table.
func NewGovServiceTagModel(conn sqlx.SqlConn) GovServiceTagModel {
	return &customGovServiceTagModel{
		defaultGovServiceTagModel: newGovServiceTagModel(conn),
	}
}
