package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovDraftModel = (*customGovDraftModel)(nil)

type (
	// GovDraftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovDraftModel.
	GovDraftModel interface {
		govDraftModel
	}

	customGovDraftModel struct {
		*defaultGovDraftModel
	}
)

// NewGovDraftModel returns a model for the database table.
func NewGovDraftModel(conn sqlx.SqlConn) GovDraftModel {
	return &customGovDraftModel{
		defaultGovDraftModel: newGovDraftModel(conn),
	}
}
