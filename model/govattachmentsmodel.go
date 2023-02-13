package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovAttachmentsModel = (*customGovAttachmentsModel)(nil)

type (
	// GovAttachmentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovAttachmentsModel.
	GovAttachmentsModel interface {
		govAttachmentsModel
	}

	customGovAttachmentsModel struct {
		*defaultGovAttachmentsModel
	}
)

// NewGovAttachmentsModel returns a model for the database table.
func NewGovAttachmentsModel(conn sqlx.SqlConn) GovAttachmentsModel {
	return &customGovAttachmentsModel{
		defaultGovAttachmentsModel: newGovAttachmentsModel(conn),
	}
}
