package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	// "os"
)

func main() {
	sqlite := true
	var db *sql.DB
	var err error
	if sqlite {
		db, err = sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal(err)
		}
	}
	postgres := false
	if postgres {
		// postgres db
	}
	defer db.Close()

	init := false

	if init {
		_, err := db.Exec("DROP TABLE sensors")
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
	GT11 := newSensor("GT11")
	fmt.Println(GT11)
	GT11.Write(db, 42.0)
	fmt.Println(GT11)
	GT11.Read(db)
	fmt.Println(GT11)
	err = GT11.writeHistorian(db)
	if err != nil {
		fmt.Println(err)
	}

	GT12 := newSensor("GT12")
	GT12.Write(db, 13.0)
	GT12.Read(db)
	GT12.writeHistorian(db)
}
