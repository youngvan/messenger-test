package messagesapp

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/randallmlough/pgxscan"
	"go.uber.org/zap"

	pb "messenger/gen/proto/go/messagesproto/v1"
)

type UserStruct struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
	Hash string `db:"hash"`
}

type MessagesApp struct {
	logger *zap.SugaredLogger
	db     *pgxpool.Pool
	user   UserStruct
}

func NewMessagesApp(logger *zap.SugaredLogger, db *pgxpool.Pool) *MessagesApp {
	app := new(MessagesApp)
	app.logger = logger
	app.db = db
	return app
}

func (app *MessagesApp) Login(token string) error {

	rows, _ := app.db.Query(context.Background(), "SELECT id, name, hash FROM users WHERE hash=$1 LIMIT 1", token)

	if err := pgxscan.NewScanner(rows).Scan(&app.user); err != nil {
		app.logger.Info("failed to fetch user: ", err)
		return err
	}

	app.logger.Info("User fetched", zap.Any("user", app.user))

	return nil
}

func (app *MessagesApp) SaveMessages(messages []*pb.Message) error {

	sqlStatement := fmt.Sprintf("INSERT INTO messages (id_from, id_to, subject, body) VALUES (%d, (SELECT id from users where hash=$1), $2, $3)", app.user.ID)

	for _, message := range messages {

		_, err := app.db.Exec(context.Background(), sqlStatement, message.ReceiverHash, message.Subject, message.Body)
		if err != nil {
			app.logger.Info("Can't save msg: ", err, message)
		}

	}

	// @todo decide about transaction usage
	return nil
}

func (app *MessagesApp) GetMessages(messages *[]*pb.Message) error {
	rows, _ := app.db.Query(context.Background(), "select subject from messages where id_to = $1", app.user.ID)

	if err := pgxscan.NewScanner(rows).Scan(&messages); err != nil {
		app.logger.Info("Error with getting messages", err)
		return err
	}

	return nil

}
