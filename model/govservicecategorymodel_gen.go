// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	govServiceCategoryFieldNames          = builder.RawFieldNames(&GovServiceCategory{})
	govServiceCategoryRows                = strings.Join(govServiceCategoryFieldNames, ",")
	govServiceCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(govServiceCategoryFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	govServiceCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(govServiceCategoryFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	govServiceCategoryModel interface {
		Insert(ctx context.Context, data *GovServiceCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GovServiceCategory, error)
		Update(ctx context.Context, data *GovServiceCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGovServiceCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GovServiceCategory struct {
		Id        int64          `db:"id"`
		Pid       int64          `db:"pid"`    // 父id
		Uid       sql.NullInt64  `db:"uid"`    // 创建人uid
		Name      sql.NullString `db:"name"`   // 类目名字
		Status    string         `db:"status"` // 状态 1启用；2禁用
		Remark    sql.NullString `db:"remark"` // 备注
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
	}
)

func newGovServiceCategoryModel(conn sqlx.SqlConn) *defaultGovServiceCategoryModel {
	return &defaultGovServiceCategoryModel{
		conn:  conn,
		table: "`gov_service_category`",
	}
}

func (m *defaultGovServiceCategoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGovServiceCategoryModel) FindOne(ctx context.Context, id int64) (*GovServiceCategory, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", govServiceCategoryRows, m.table)
	var resp GovServiceCategory
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGovServiceCategoryModel) Insert(ctx context.Context, data *GovServiceCategory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, govServiceCategoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Pid, data.Uid, data.Name, data.Status, data.Remark, data.DeletedAt)
	return ret, err
}

func (m *defaultGovServiceCategoryModel) Update(ctx context.Context, data *GovServiceCategory) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, govServiceCategoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Pid, data.Uid, data.Name, data.Status, data.Remark, data.DeletedAt, data.Id)
	return err
}

func (m *defaultGovServiceCategoryModel) tableName() string {
	return m.table
}
