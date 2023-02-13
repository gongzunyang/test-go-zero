package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovDraftMasterModel = (*customGovDraftMasterModel)(nil)

type (
	// GovDraftMasterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovDraftMasterModel.
	GovDraftMasterModel interface {
		govDraftMasterModel
	}

	customGovDraftMasterModel struct {
		*defaultGovDraftMasterModel
	}
)

// NewGovDraftMasterModel returns a model for the database table.
func NewGovDraftMasterModel(conn sqlx.SqlConn) GovDraftMasterModel {
	return &customGovDraftMasterModel{
		defaultGovDraftMasterModel: newGovDraftMasterModel(conn),
	}
}
