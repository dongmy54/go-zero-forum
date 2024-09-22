package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		PracticeQuery(ctx context.Context) error
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn, c, opts...),
	}
}

func (m *customPostModel) PracticeQuery(ctx context.Context) error {
	// 事务
	m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		// 只要其中一个报错 则失败
		_, err := s.ExecCtx(ctx, "insert into post (title, content, user_id) values (?,?, ?)", "标题1", "内容1", 1)
		if err != nil {
			return err
		}

		// 这里user_id必填会为空
		_, err = s.ExecCtx(ctx, "insert into post (title, content, user_id) values (?,?,?)", "标题1", "内容1", 11)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}
