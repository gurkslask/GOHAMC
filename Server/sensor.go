package main

import "database/sql"
import "log"
import "fmt"

type Sensor struct {
	name  string
	value float32
}

func newSensor(name string) *Sensor {
	return &Sensor{name: name}
}

func (s *Sensor) Read(db *sql.DB) error {
	// Read from SQL
	rows, err := db.Query("select value from sensors where name = ?", s.name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&s.value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Sensor) Write(db *sql.DB, value float32) error {
	//Write to SQL

	s.value = value

	//Check if value exists
	err := s.Read(db)
	if err != nil {
		log.Print("Value doesnt exist")
		// If vkalue doesnt exist
		stmt, err := db.Prepare("INSERT INTO sensors(name, value) VALUES(?, ?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(s.name, s.value)
		if err != nil {
			return err
		}
	}
	// Value exists
	log.Print("Value exist")

	_, err = db.Exec("UPDATE sensors SET value = {} WHERE name = {}", s.value, s.name)
	if err != nil {
		return err
	}
	return nil

}

func (s Sensor) String() string {
	return fmt.Sprintf("Name: %v, value: %v\n", s.name, s.value)
}
