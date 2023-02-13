package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GovServiceCategoryModel = (*customGovServiceCategoryModel)(nil)

type (
	// GovServiceCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGovServiceCategoryModel.
	GovServiceCategoryModel interface {
		govServiceCategoryModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GovServiceCategory, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindData(ctx context.Context, countBuilder squirrel.SelectBuilder) ([]*GovServiceCategory, error)
	}

	customGovServiceCategoryModel struct {
		*defaultGovServiceCategoryModel
	}
)

func (m customGovServiceCategoryModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GovServiceCategory, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, _, _ := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	var resp []*GovServiceCategory
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return resp, err
	}
}

func (m customGovServiceCategoryModel) FindData(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*GovServiceCategory, error) {
	query, _, _ := rowBuilder.ToSql()
	var resp []*GovServiceCategory
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return resp, err
	}
}

func (m customGovServiceCategoryModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
	countQuery, _, _ := countBuilder.ToSql()
	var countNum int64
	err := m.conn.QueryRowCtx(ctx, &countNum, countQuery)
	if err != nil {
		return countNum, err
	}
	return countNum, err
}

func (m *customGovServiceCategoryModel) RowBuilder() squirrel.SelectBuilder {
	//TODO implement me
	return squirrel.Select(govServiceCategoryRows).From(m.table)
	panic("implement me")
}

func (m *customGovServiceCategoryModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ") as bbb").From(m.table)
}

// NewGovServiceCategoryModel returns a model for the database table.
func NewGovServiceCategoryModel(conn sqlx.SqlConn) GovServiceCategoryModel {
	return &customGovServiceCategoryModel{
		defaultGovServiceCategoryModel: newGovServiceCategoryModel(conn),
	}
}
