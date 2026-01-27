package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,

		UNIQUE(title)
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

// *CREATE TABLE **IF NOT EXISTS [название таблицы] (
// id ***SERIAL PRIMARY KEY
// ... [тип данных соответствующий стандартам СУБД] ****NOT NULL/-
// );

// * - Создаем таблицу
// ** - Указываем, что будем ее создавать если такой таблицы не существует
// *** - Первичный ключ, SERIAL означает, что при добавлении новой записи, ее идентификатор будет увеличиваться на +1
// **** - NOT NULL нужно указывать, если мы хотим чтобы в этом конкретном поле что-то было, если нам не принципиально, можно оставить как есть
// UNIQUE([название столбца]) - явно указываем что те или иные столбцы/столбец должны быть уникальными (исключение: id, она по умолчанию является уникальной)
