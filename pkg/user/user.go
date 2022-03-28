package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/randallmlough/pgxscan"
	"go.uber.org/zap"
)

type UserDataStruct struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
	Hash string `db:"hash"`
}

type UserStruct struct {
	Logger     *zap.SugaredLogger
	Db         *pgxpool.Pool
	LoggedUser UserDataStruct
}

func (user *UserStruct) Login(ctx context.Context, token string) error {

	rows, _ := user.Db.Query(ctx, "SELECT id, name, hash FROM users WHERE hash=$1 LIMIT 1", token)

	if err := pgxscan.NewScanner(rows).Scan(&user.LoggedUser); err != nil {
		user.Logger.Info("failed to fetch user: ", err)
		return err
	}

	user.Logger.Info("User fetched", zap.Any("user", user.LoggedUser))

	return nil
}
