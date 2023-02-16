package main

import (
	"database/sql"
	"fmt"
	_"github.com/mattn/go-sqlite3"
	_"log"
)

// Create model here

type Person struct {
	id         int
	first_name string
	last_name  string
	email      string
	ip_address string
}

func AddPerson(db *sql.DB, newPerson Person) {
	stmt, _ := db.Prepare("insert into people (id, first_name, last_name, email, ip_address) values (?,?,?,?,?)")
	stmt.Exec(nil, newPerson.first_name, newPerson.last_name, newPerson.email, newPerson.ip_address)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newPerson.first_name, newPerson.last_name)
}

func searchForPerson(db *sql.DB, searchString string) []Person {
	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")

	// err = rows.Err()

	handleErr(err)
	defer rows.Close()
	people := make([]Person, 0)
	for rows.Next() {
		ourPerson := Person{}
		err = rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
		handleErr(err)
		people = append(people, ourPerson)

	}
	err = rows.Err()
	handleErr(err)
	return people
}

func getPersonById(db *sql.DB, ourID string) Person {
	rows, err := db.Query("select id, first_name, last_name, email, ip_address from people where id = '" + ourID + "' ")
	handleErr(err)
	defer rows.Close()
	ourPerson := Person{}
	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
	}
	return ourPerson

}

func updatePerson(db *sql.DB, currentPerson Person) int64 {
	stmt, err := db.Prepare("update people set first_name=?, last_name=?, email=?, ip_address=? where id=?")
	handleErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(currentPerson.first_name, currentPerson.last_name, currentPerson.email, currentPerson.ip_address, currentPerson.id)
	handleErr(err)

	affected, err := res.RowsAffected()
	handleErr(err)
	return affected
}

func deletePerson(db *sql.DB, idToDelete string) int64 {
	stmt, err := db.Prepare("delete from people where id=?")
	handleErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(idToDelete)
	handleErr(err)
	affected, err := res.RowsAffected()
	handleErr(err)
	return affected
}
