package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovPublishServiceClassifyModel = (*customGovPublishServiceClassifyModel)(nil)

type (
	// GovPublishServiceClassifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovPublishServiceClassifyModel.
	GovPublishServiceClassifyModel interface {
		govPublishServiceClassifyModel
	}

	customGovPublishServiceClassifyModel struct {
		*defaultGovPublishServiceClassifyModel
	}
)

// NewGovPublishServiceClassifyModel returns a model for the database table.
func NewGovPublishServiceClassifyModel(conn sqlx.SqlConn) GovPublishServiceClassifyModel {
	return &customGovPublishServiceClassifyModel{
		defaultGovPublishServiceClassifyModel: newGovPublishServiceClassifyModel(conn),
	}
}
