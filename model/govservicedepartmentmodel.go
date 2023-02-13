package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovServiceDepartmentModel = (*customGovServiceDepartmentModel)(nil)

type (
	// GovServiceDepartmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovServiceDepartmentModel.
	GovServiceDepartmentModel interface {
		govServiceDepartmentModel
	}

	customGovServiceDepartmentModel struct {
		*defaultGovServiceDepartmentModel
	}
)

// NewGovServiceDepartmentModel returns a model for the database table.
func NewGovServiceDepartmentModel(conn sqlx.SqlConn) GovServiceDepartmentModel {
	return &customGovServiceDepartmentModel{
		defaultGovServiceDepartmentModel: newGovServiceDepartmentModel(conn),
	}
}
