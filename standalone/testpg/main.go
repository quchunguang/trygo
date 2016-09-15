/*
Manuals
=======

Install DB
----------

pi@host$ sudo apt-get install postgresql

Create database and its account
-------------------------------

Notice: user/password must be same with that in OS !!!

pi@host$ sudo -u postgres psql -U postgres
CREATE USER pi WITH PASSWORD 'xxxxxx';
CREATE DATABASE snsrobot;
GRANT ALL PRIVILEGES ON DATABASE snsrobot TO pi;
\q

# raspbian with raspberry pi 2:
pi@host$ sudo -u postgres cat <<_END >>/var/lib/postgresql/data/pg_hba.conf
# ubuntu:
pi@host$ #sudo -u postgres cat <<_END >>/etc/postgresql/9.4/main/pg_hba.conf
host all "pi" 0.0.0.0/0 trust
_END

pi@host$ sudo service postgresql restart

Create tables in host
---------------------

pi@host$ psql snsrobot
snsrobot => CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    departname character varying(500) NOT NULL,
    Created date,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);
snsrobot => \q

Configure the remote access
---------------------------

pi@host$ sudo -u postgres cat <<_END >>/etc/postgresql/9.4/main/postgresql.conf
listen_addresses='*'
_END

pi@host$ sudo service postgresql restart
user@client$ psql -h 192.168.1.109 -U pi -W -d snsrobot

Usage Example
-------------
user@client$ go build
user@client$ ./testpg.exe -c pi:xxxxx@192.168.1.109/snsrobot

[Demo program]: http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html
[Remote access]: http://www.cyberciti.biz/tips/postgres-allow-remote-access-tcp-connection.html
*/
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	var db_conn string
	flag.StringVar(&db_conn, "c", "", "like user:pass@ip/dbname")
	flag.Parse()
	if db_conn == "" {
		flag.Usage()
		return
	}

	dbinfo := fmt.Sprintf("postgres://%s?sslmode=disable",
		db_conn)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("astaxieupdate", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
