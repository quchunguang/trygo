/*
Manuals
=======

Install DB
----------

pi@host$ sudo apt-get install postgresql

Create database and its account
-------------------------------

Notice: user/password must match exist accounts in OS !!!

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
)

func main() {
	var db_conn string
	flag.StringVar(&db_conn, "c", "", "like pi:pass@192.168.1.109/snsrobot")
	flag.Parse()
	if db_conn == "" {
		flag.Usage()
		return
	}

	dbinfo := fmt.Sprintf("postgres://%s?sslmode=disable", db_conn)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("\n# Inserting values")

	res, err := db.Exec(`INSERT INTO emp(data) VALUES('{"id": 1, "name": "raju", "description": "HR", "salary": 25000.00 }')`)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("RowsAffected", affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
