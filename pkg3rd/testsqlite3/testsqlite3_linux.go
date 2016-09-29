package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	defer db.Close()
	checkErr(err)

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	checkErr(err)

	tx, err := db.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	defer stmt.Close()
	checkErr(err)
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		checkErr(err)
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	checkErr(err)
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	checkErr(err)
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	checkErr(err)

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	checkErr(err)

	rows, err = db.Query("select id, name from foo")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
