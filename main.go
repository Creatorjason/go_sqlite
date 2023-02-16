package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite3", "./names.db")
	handleErr(err)
	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")
	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })
	menu.Option("Add new person", 0, true, nil)
	menu.Option("Find a person", 1, false, nil)
	menu.Option("Update a person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menu.Option("Quit application", 4, false, nil)
	err = menu.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {
	case 0:
		fmt.Println("Adding new person")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your first name:")
		firstName, _ := reader.ReadString('\n')
		firstName = strings.TrimSuffix(firstName, "\n")
		fmt.Println("Enter your last name:")
		lastName, _ := reader.ReadString('\n')
		lastName = strings.TrimSuffix(lastName, "\n")
		fmt.Println("Enter your email address")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSuffix(email, "\n")
		fmt.Println("Enter you IP address")
		ipAddr, _ := reader.ReadString('\n')
		ipAddr = strings.TrimSuffix(ipAddr, "\n")

		newPerson := Person{
			first_name: firstName,
			last_name:  lastName,
			email:      email,
			ip_address: ipAddr,
		}
		AddPerson(db, newPerson)

	case 1:
		fmt.Println("Finding person information")
		fmt.Println("Enter name of person:")
		reader := bufio.NewReader(os.Stdin)
		searchPerson, _ := reader.ReadString('\n')
		searchPerson = strings.TrimSuffix(searchPerson, "\n")
		people := searchForPerson(db, searchPerson)
		fmt.Printf("Found %v results\n", len(people))
		for _, ourPerson := range people {
			fmt.Printf("\n----\nFirst Name: %s\nLast Name: %s\nEmail: %s\nIP Address: %s\n", ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address)
		}
	case 2:
		fmt.Println("updating person information")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter an ID to update:")
		updateId, _ := reader.ReadString('\n')
		currentPerson := getPersonById(db, updateId)
		fmt.Printf("First name (Currently is %s)", currentPerson.first_name)
		first_name, _ := reader.ReadString('\n')
		if first_name != "\n" {
			currentPerson.first_name = strings.TrimSuffix(first_name, "\n")
		}
		fmt.Printf("Last name (Currently is %s)", currentPerson.last_name)
		last_name, _ := reader.ReadString('\n')
		if last_name != "\n" {
			currentPerson.last_name = strings.TrimSuffix(last_name, "\n")
		}
		fmt.Printf("Email (Currently is %s)", currentPerson.email)
		email, _ := reader.ReadString('\n')
		if email != "\n" {
			currentPerson.email = strings.TrimSuffix(email, "\n")
		}
		fmt.Printf("IP Address (Currently %s):", currentPerson.ip_address)
		ipAddress, _ := reader.ReadString('\n')
		if ipAddress != "\n" {
			currentPerson.ip_address = strings.TrimSuffix(ipAddress, "\n")
		}

		affected := updatePerson(db, currentPerson)
		if affected == 1 {
			fmt.Println("One row affected")
		}

	case 3:
		fmt.Println("deleting person information")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the ID you want to delete : ")
		searchString, _ := reader.ReadString('\n')

		idToDelete := strings.TrimSuffix(searchString, "\n")

		affected := deletePerson(db, idToDelete)

		if affected == 1 {
			fmt.Println("Deleted person from database")
		}
	case 4:
		fmt.Println("Goodbye")
		os.Exit(1)

	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
