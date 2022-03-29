package messagesapp

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type MessagesStruct struct {
	Subject      string
	Body         string
	ReceiverHash string
}

type MessagesApp struct {
	Logger *zap.SugaredLogger
	Db     *pgxpool.Pool
}

const sqlInsertMessage = "INSERT INTO messages (id_from, id_to, subject, body) VALUES ($1, (SELECT id from users where hash=$2), $3, $4)"
const sqlGetMessages = "SELECT subject FROM messages WHERE id_to = $1"

func (app *MessagesApp) SaveMessages(ctx context.Context, messages []*MessagesStruct) error {

	for _, message := range messages {

		_, err := app.Db.Exec(ctx, sqlInsertMessage, ctx.Value("userID"), message.ReceiverHash, message.Subject, message.Body)
		if err != nil {
			app.Logger.Info("Can't save msg: ", err, message)
		}

	}

	// @todo decide about transaction usage
	return nil
}

func (app *MessagesApp) GetMessages(ctx context.Context) ([]*MessagesStruct, error) {

	var result []*MessagesStruct

	if err := pgxscan.Select(ctx, app.Db, &result, sqlGetMessages, ctx.Value("userID")); err != nil {
		app.Logger.Info("Error with getting messages", err)
		return nil, err
	}

	return result, nil

}
