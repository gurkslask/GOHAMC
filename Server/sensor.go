package main

import "database/sql"
import "log"
import "fmt"
import "time"

type Sensor struct {
	name  string
	value float32
}

func newSensor(name string) *Sensor {
	return &Sensor{name: name}
}

func (s *Sensor) Read(db *sql.DB) error {
	// Read from SQL
	rows := db.QueryRow("select value from sensors where name = ?", s.name)
	err := rows.Scan(&s.value)
	if err != nil {
		return err
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

func (s *Sensor) writeHistorian(db *sql.DB) error {
	// Write current value in database to Historian
	s.Read(db)

	timeString := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeString)

	stmt, err := db.Prepare(`
	INSERT INTO historian(
		value,
		time,
		sensor_id
	)
	select 
		?,
		?,
		(select sensor_id from sensors where name = ?);
	
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(s.value, timeString, s.name)
	if err != nil {
		return err
	}
	return nil
}
