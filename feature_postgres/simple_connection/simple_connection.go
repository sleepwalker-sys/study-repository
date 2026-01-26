package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourPassword@YourHostName:5432/YourDatabaseName"

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil { // проверочный запрос в БД для того, чтобы определить валидное ли подключение к БД, есть ли доступ к серверам БД?
		panic(err)
	}

	fmt.Println("Подключение к базе данных прошло успешно!")
}

// Для работы с PostgreSQL будем использовать стороннюю библиотеку "pgx"
