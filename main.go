package main

import (
	// "fmt"
	"Databases/s5_dbase"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const DB_NAME = "freeblebase.db"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("HELP INFO....")
	}

	if os.Args[1] == "create" {
		err := s5_dbase.SetupDB(DB_NAME)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Database Created \n")
	}

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	switch os.Args[1] {

	case "add_user":
		if len(os.Args) < 4 {
			log.Fatal("add_user <username> <password>")
		}
		err := s5_dbase.AddUser(db, os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("User Added: %s\n", os.Args[2])

	case "login":
		if len(os.Args) < 4 {
			log.Fatal("login <username> <password>")
		}
		n, err := s5_dbase.CheckLogin(db, os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Login = %t\n", n)
	
	case "add_item":
		if len(os.Args) < 6 {
			log.Fatal("add_item <giver_id> <item_name> <description> <image_location>")
		}

		fdata, err := ioutil.ReadFile(os.Args[5])
		if err != nil {
			log.Fatal("Read Image Error: %w", err)
		}
		err = s5_dbase.AddItem(db, os.Args[2], os.Args[3], os.Args[4], fdata)
		if err != nil {
			log.Fatal("could not add item %w", err)
		}
		fmt.Printf("Item added: %s\n", os.Args[3])
	
	case "get_image":
		if len(os.Args) < 3 {
			log.Fatal("get_image <image_id>")
		}

		item_id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("image_id should be int")
		}

		err = s5_dbase.GetImage(db, item_id, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	
	case "search_items":
		if len(os.Args) < 3 {
			log.Fatal("search_items <search term>")
		}

		list, err := s5_dbase.SearchItem(db, os.Args[2])
		if err != nil {
			log.Fatal("could not find items: %w", err)
		}

		for k, v := range list {
			fmt.Printf("Item %d = %v\n", k, v)
		}
	case "set_receiver":
		if len(os.Args) < 4 {
			log.Fatal("set_receiver <item_id> <receiver_id>")
		}
		
		i_id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("item_id should be int")
		}

		if err = s5_dbase.SetReceiver(db, i_id, os.Args[3]); err != nil {
			log.Fatal("could not set receiver %w", err)
		}

		println("reciever set")
	default:
		fmt.Printf("TODO: Help Info....\n")
	}
}

