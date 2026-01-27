package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET description = 'testing'
	WHERE completed = false
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err

}
