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
	govNationFieldNames          = builder.RawFieldNames(&GovNation{})
	govNationRows                = strings.Join(govNationFieldNames, ",")
	govNationRowsExpectAutoSet   = strings.Join(stringx.Remove(govNationFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	govNationRowsWithPlaceHolder = strings.Join(stringx.Remove(govNationFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	govNationModel interface {
		Insert(ctx context.Context, data *GovNation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GovNation, error)
		Update(ctx context.Context, data *GovNation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGovNationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GovNation struct {
		Id   int64          `db:"id"`
		Name sql.NullString `db:"name"`
	}
)

func newGovNationModel(conn sqlx.SqlConn) *defaultGovNationModel {
	return &defaultGovNationModel{
		conn:  conn,
		table: "`gov_nation`",
	}
}

func (m *defaultGovNationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGovNationModel) FindOne(ctx context.Context, id int64) (*GovNation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", govNationRows, m.table)
	var resp GovNation
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

func (m *defaultGovNationModel) Insert(ctx context.Context, data *GovNation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, govNationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name)
	return ret, err
}

func (m *defaultGovNationModel) Update(ctx context.Context, data *GovNation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, govNationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Id)
	return err
}

func (m *defaultGovNationModel) tableName() string {
	return m.table
}
