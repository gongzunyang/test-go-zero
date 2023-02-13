package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovDraftDepartmentModel = (*customGovDraftDepartmentModel)(nil)

type (
	// GovDraftDepartmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovDraftDepartmentModel.
	GovDraftDepartmentModel interface {
		govDraftDepartmentModel
	}

	customGovDraftDepartmentModel struct {
		*defaultGovDraftDepartmentModel
	}
)

// NewGovDraftDepartmentModel returns a model for the database table.
func NewGovDraftDepartmentModel(conn sqlx.SqlConn) GovDraftDepartmentModel {
	return &customGovDraftDepartmentModel{
		defaultGovDraftDepartmentModel: newGovDraftDepartmentModel(conn),
	}
}
