package main

//_"github.com/mattn/go-sqlite3"

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var db *sqlx.DB

	// exactly the same as the built-in
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())

	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())

	}

	createScheame(db)
}

type user struct {
	ID   int    "db:id"
	name string "db:name"
}

func createScheame(db *sqlx.DB) {
	schema := `CREATE TABLE user (
			id integer primary key AUTOINCREMENT,
			name varchar[56]);`

	//ejecutar esto
	db.Exec(schema)

	db.MustExec("INSERT INTO user(name) VAlUES(?)", "tomi")

	rows, err := db.Query("SELECT id , name from user ")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var u user
		//le paso la posicion en memoria
		err = rows.Scan(&u.ID, &u.name)
		fmt.Println(u)

	}
}
