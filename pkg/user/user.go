package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
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

func (user *UserStruct) CheckAuth(ctx context.Context) (context.Context, error) {

	headers, _ := metadata.FromIncomingContext(ctx)

	if len(headers["auth"]) == 0 || headers["auth"][0] == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Auth header missed")
	}
	token := headers["auth"][0]

	if err := user.validateToken(ctx, token); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "wrong Auth header")
	}

	ctx = context.WithValue(ctx, "userID", user.LoggedUser.ID)

	return ctx, nil

}

func (user *UserStruct) validateToken(ctx context.Context, token string) error {

	rows, _ := user.Db.Query(ctx, "SELECT id, name, hash FROM users WHERE hash=$1 LIMIT 1", token)

	if err := pgxscan.ScanOne(&user.LoggedUser, rows); err != nil {
		user.Logger.Info("failed to fetch user: ", err)
		return err
	}

	return nil
}
