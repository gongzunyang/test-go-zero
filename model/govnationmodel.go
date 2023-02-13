package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GovNationModel = (*customGovNationModel)(nil)

type (
	// GovNationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovNationModel.
	GovNationModel interface {
		govNationModel
	}

	customGovNationModel struct {
		*defaultGovNationModel
	}
)

// NewGovNationModel returns a model for the database table.
func NewGovNationModel(conn sqlx.SqlConn) GovNationModel {
	return &customGovNationModel{
		defaultGovNationModel: newGovNationModel(conn),
	}
}
