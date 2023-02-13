package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovServiceManagementModel = (*customGovServiceManagementModel)(nil)

type (
	// GovServiceManagementModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovServiceManagementModel.
	GovServiceManagementModel interface {
		govServiceManagementModel
	}

	customGovServiceManagementModel struct {
		*defaultGovServiceManagementModel
	}
)

// NewGovServiceManagementModel returns a model for the database table.
func NewGovServiceManagementModel(conn sqlx.SqlConn) GovServiceManagementModel {
	return &customGovServiceManagementModel{
		defaultGovServiceManagementModel: newGovServiceManagementModel(conn),
	}
}
