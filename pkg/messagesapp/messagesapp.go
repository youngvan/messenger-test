package messagesapp

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"messenger/pkg/user"
	pb "messenger/proto/gen/proto/go/messagesproto/v1"
)

type MessagesStruct struct {
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

type MessagesApp struct {
	Logger *zap.SugaredLogger
	Db     *pgxpool.Pool
}

func (app *MessagesApp) SaveMessages(ctx context.Context, user *user.UserStruct, messages []*pb.Message) error {

	sqlStatement := "INSERT INTO messages (id_from, id_to, subject, body) VALUES ($1, (SELECT id from users where hash=$2), $3, $4)"

	for _, message := range messages {

		_, err := app.Db.Exec(ctx, sqlStatement, user.LoggedUser.ID, message.ReceiverHash, message.Subject, message.Body)
		if err != nil {
			app.Logger.Info("Can't save msg: ", err, message)
		}

	}

	// @todo decide about transaction usage
	return nil
}

func (app *MessagesApp) GetMessages(ctx context.Context, user *user.UserStruct) ([]MessagesStruct, error) {

	result := []MessagesStruct{}

	query := "select subject from messages where id_to = $1"

	if err := pgxscan.Select(ctx, app.Db, &result, query, user.LoggedUser.ID); err != nil {
		app.Logger.Info("Error with getting messages", err)
		return nil, err
	}

	return result, nil

}
