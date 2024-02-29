package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SeriesModel = (*customSeriesModel)(nil)

type (
	// SeriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSeriesModel.
	SeriesModel interface {
		seriesModel
		withSession(session sqlx.Session) SeriesModel
	}

	customSeriesModel struct {
		*defaultSeriesModel
	}
)

// NewSeriesModel returns a model for the database table.
func NewSeriesModel(conn sqlx.SqlConn) SeriesModel {
	return &customSeriesModel{
		defaultSeriesModel: newSeriesModel(conn),
	}
}

func (m *customSeriesModel) withSession(session sqlx.Session) SeriesModel {
	return NewSeriesModel(sqlx.NewSqlConnFromSession(session))
}
