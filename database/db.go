package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func connectDb() {
	host := "some-postgres"
	user := "postgres"
	password := "productDB"
	dbname := "postgres"
	ssl := "disable"

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("error ping", err.Error())
	}
}

func createTables() {
	productTableQ := `
	CREATE TABLE IF NOT EXISTS "products" (
		id serial,
		name varchar(150) NOT NULL,
		price int NOT NULL,
		description text NOT NULL,
		quantity int NOT NULL,
		createdAt timestamp default now(),
		updatedAt timestamp default now(),

		PRIMARY KEY (id)
	);`

	_, err := db.Exec(productTableQ)
	if err != nil {
		log.Panicln("error create table", err.Error())
	}
}

func mockData() {
	query := `
	INSERT INTO products (name,price,description,quantity)
	values 
	('botol gaming',20000,'botol gaming ciamik', 5),
	('standing desk',700000,'meja bisa berdiri', 3)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Panicln("error insert mock data", err.Error())
	}
}

func NewPG() *sql.DB {
	connectDb()
	createTables()
	mockData()
	return db
}
