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
	deviceInfoFieldNames          = builder.RawFieldNames(&DeviceInfo{})
	deviceInfoRows                = strings.Join(deviceInfoFieldNames, ",")
	deviceInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(deviceInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	deviceInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(deviceInfoFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	deviceInfoModel interface {
		Insert(ctx context.Context, data *DeviceInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DeviceInfo, error)
		FindOneByProductIDDeviceName(ctx context.Context, productID string, deviceName string) (*DeviceInfo, error)
		Update(ctx context.Context, data *DeviceInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDeviceInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DeviceInfo struct {
		Id          int64        `db:"id"`
		ProductID   string       `db:"productID"`  // 产品id
		DeviceName  string       `db:"deviceName"` // 设备名称
		Secret      string       `db:"secret"`     // 设备秘钥
		FirstLogin  sql.NullTime `db:"firstLogin"` // 激活时间
		LastLogin   sql.NullTime `db:"lastLogin"`  // 最后上线时间
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime time.Time    `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
		Version     string       `db:"version"`  // 固件版本
		LogLevel    int64        `db:"logLevel"` // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试
		Cert        string       `db:"cert"`     // 设备证书
		IsOnline    int64        `db:"isOnline"` // 是否在线,1是2否
		Tags        string       `db:"tags"`     // 设备标签
		Address     string       `db:"address"`  // 所在地址
		Position    string       `db:"position"` // 设备的位置,默认百度坐标系BD09
	}
)

func newDeviceInfoModel(conn sqlx.SqlConn) *defaultDeviceInfoModel {
	return &defaultDeviceInfoModel{
		conn:  conn,
		table: "`device_info_test2`",
	}
}

func (m *defaultDeviceInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDeviceInfoModel) FindOne(ctx context.Context, id int64) (*DeviceInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", deviceInfoRows, m.table)
	var resp DeviceInfo
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

func (m *defaultDeviceInfoModel) FindOneByProductIDDeviceName(ctx context.Context, productID string, deviceName string) (*DeviceInfo, error) {
	var resp DeviceInfo
	query := fmt.Sprintf("select %s from %s where `productID` = ? and `deviceName` = ? limit 1", deviceInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, productID, deviceName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceInfoModel) Insert(ctx context.Context, data *DeviceInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, deviceInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.DeviceName, data.Secret, data.FirstLogin, data.LastLogin, data.Version, data.LogLevel, data.Cert, data.IsOnline, data.Tags, data.Address, data.Position)
	return ret, err
}

func (m *defaultDeviceInfoModel) Update(ctx context.Context, newData *DeviceInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, deviceInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ProductID, newData.DeviceName, newData.Secret, newData.FirstLogin, newData.LastLogin, newData.Version, newData.LogLevel, newData.Cert, newData.IsOnline, newData.Tags, newData.Address, newData.Position, newData.Id)
	return err
}

func (m *defaultDeviceInfoModel) tableName() string {
	return m.table
}
