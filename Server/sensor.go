package main

import (
	"database/sql"
	"log"
)

type sensor struct {
	name  string
	value float32
}

func (s sensor) Read(db database) error {
	// Read from SQL
	rows, err := db.Query("select value from sensors where name = ?", s.name)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&s.value)
		if err != nil {
			return err
		}
	}
	return nil

}
