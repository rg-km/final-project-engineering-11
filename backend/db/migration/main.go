package migration

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Migrate() {
	db, err := sql.Open("sqlite3", "db/migration/app2.db?_loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		name varchar(255) not null,
		password varchar(255) not null,
		role varchar(255) not null,
		address varchar(255) not null,
		phone varchar(255) not null,
		email varchar(255) not null,
		created_at datetime not null
);

CREATE TABLE IF NOT EXISTS mentor (
    id integer not null primary key AUTOINCREMENT,
	skill varchar(255) not null,
	bio varchar(255) not null,
	name varchar(255) not null,
	address varchar(255) not null,
	phone varchar(255) not null,
	email varchar(255) not null,
	created_at datetime not null
);
CREATE TABLE IF NOT EXISTS artikel (
    id integer not null primary key AUTOINCREMENT,
	judul varchar(255) not null,
	content varchar(255) not null,
	created_at datetime not null
);

CREATE TABLE IF NOT EXISTS bookmentor (
	id integer not null primary key AUTOINCREMENT,
	mentor_id integer not null,
	book_id integer not null,
	created_at datetime not null
);


`)

	if err != nil {
		panic(err)
	}
}