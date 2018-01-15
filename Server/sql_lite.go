package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
)

func main() {
	fmt.Println("Hej")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	init := true

	if init {
		_, err = db.Exec("DROP TABLE sensors")
		if err != nil {
			log.Print(err)
		}

		_, err = db.Exec("DROP TABLE historian")
		if err != nil {
			log.Print(err)
		}

		stmt := `create table sensors (
			sensor_id integer not null ,
			name text,
			value integer,
			CONSTRAINT sensor_pk PRIMARY KEY (sensor_id)
		);`

		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
		stmt = `create table historian (
			historian_id integer not null ,
			value integer,
			time timestamp not null,
			sensor_id integer not null,
			CONSTRAINT historian_pk PRIMARY KEY (historian_id) 
			CONSTRAINT fk_historian_sensor_id FOREIGN KEY (sensor_id) REFERENCES sensors(sensor_id)
		);`

		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
