// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productRemoteConfigFieldNames          = builder.RawFieldNames(&ProductRemoteConfig{})
	productRemoteConfigRows                = strings.Join(productRemoteConfigFieldNames, ",")
	productRemoteConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(productRemoteConfigFieldNames, "`id`", "`createdTime`", "`updatedTime`", "`deletedTime`"), ",")
	productRemoteConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(productRemoteConfigFieldNames, "`id`", "`createdTime`", "`updatedTime`", "`deletedTime`"), "=?,") + "=?"
)

type (
	productRemoteConfigModel interface {
		Insert(ctx context.Context, data *ProductRemoteConfig) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductRemoteConfig, error)
		Update(ctx context.Context, data *ProductRemoteConfig) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductRemoteConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ProductRemoteConfig struct {
		Id          int64        `db:"id"`
		ProductID   string       `db:"productID"`   // 产品id
		Content     string       `db:"content"`     // 配置内容
		CreatedTime time.Time    `db:"createdTime"` // 创建时间
		UpdatedTime time.Time    `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime `db:"deletedTime"` // 删除时间
	}
)

func newProductRemoteConfigModel(conn sqlx.SqlConn) *defaultProductRemoteConfigModel {
	return &defaultProductRemoteConfigModel{
		conn:  conn,
		table: "`product_remote_config`",
	}
}

func (m *defaultProductRemoteConfigModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultProductRemoteConfigModel) FindOne(ctx context.Context, id int64) (*ProductRemoteConfig, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRemoteConfigRows, m.table)
	var resp ProductRemoteConfig
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

func (m *defaultProductRemoteConfigModel) Insert(ctx context.Context, data *ProductRemoteConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, productRemoteConfigRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.Content)
	return ret, err
}

func (m *defaultProductRemoteConfigModel) Update(ctx context.Context, data *ProductRemoteConfig) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRemoteConfigRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.Content, data.Id)
	return err
}

func (m *defaultProductRemoteConfigModel) tableName() string {
	return m.table
}
