package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"import-track-data/notification"

	_ "github.com/lib/pq"
)

func OpenConnectionDatabase() *sql.DB {
	connectionString := "user=postgres dbname=monteirobrena sslmode=disable"
	database, connectionError := sql.Open("postgres", connectionString)

	notification.Error(connectionError)

	return database
}

func CloseConnectionDatabase(database *sql.DB) {
	database.Close()
}

func WriteDatabase(rows [][]string) {
	database := OpenConnectionDatabase()

	for index := 0; index < len(rows); index++ {
		var values bytes.Buffer

		values.WriteString("'")
		values.WriteString(rows[index][0])
		values.WriteString("', '")
		values.WriteString(rows[index][1])
		values.WriteString("', ")
		values.WriteString(rows[index][2])
		values.WriteString(", ")
		values.WriteString(rows[index][3])

		result, insertError := database.Exec(fmt.Sprintf("INSERT INTO tracks (name, isrc, units, revenue) VALUES (%s)", values.String()))

		notification.Error(insertError)

		fmt.Println(result.RowsAffected())
	}

	CloseConnectionDatabase(database)
}
