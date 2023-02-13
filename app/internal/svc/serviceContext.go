package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gov-cms/app/internal/config"
	"gov-cms/model"
)

type ServiceContext struct {
	Config                  config.Config
	GovServiceCategoryModel model.GovServiceCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		GovServiceCategoryModel: model.NewGovServiceCategoryModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
