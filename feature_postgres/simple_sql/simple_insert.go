package simplesql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	title string,
	description string,
	completed bool,
	completedAt time.Time,
) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx, sqlQuery, title, description, completed, completedAt) // - !

	return err
}

// *INSERT INTO [название таблицы] ()
// **VALUES ()

// * - Вставляем в таблицу данные
// ** - По порядку, указанный в первой строчке, вводим значения

// ! - для того чтобы у нас был более динамичная запись значений в таблицу, нужно в VALUES () указывать $, после него число 1, 2, 3, 4...
// А в conn.Exec(), начиная с третьего аргумента, указываем значения
