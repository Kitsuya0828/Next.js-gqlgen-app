package services

import (
	"context"
	"database/sql"

	"github.com/Kitsuya0828/Next.js-gqlgen-app/docs/chapter03/graph/db"
	"github.com/Kitsuya0828/Next.js-gqlgen-app/docs/chapter03/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

type ContextExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users( // from users
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name), // select id, name
		db.UserWhere.Name.EQ(name), // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}