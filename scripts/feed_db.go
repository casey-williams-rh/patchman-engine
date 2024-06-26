package main

import (
	"app/base/database"
	"os"
)

func main() {
	database.InitDB()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "inventory_hosts":
			createInventoryHosts()
			return
		case "feed":
			feed()
			return
		}
	}
	// create inventory.hosts and feed db
	createInventoryHosts()
	feed()
}

func createInventoryHosts() {
	database.DBWait("empty")
	query, err := os.ReadFile("./dev/create_inventory_hosts.sql")
	if err != nil {
		panic(err)
	}
	err = database.DB.Exec(string(query)).Error
	if err != nil {
		panic(err)
	}
}

func feed() {
	database.DBWait("full")
	query, err := os.ReadFile("./dev/test_data.sql")
	if err != nil {
		panic(err)
	}
	err = database.DB.Exec(string(query)).Error
	if err != nil {
		panic(err)
	}
}
