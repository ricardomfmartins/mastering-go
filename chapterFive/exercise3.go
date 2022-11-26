package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}
	host := arguments[1]
	port := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	database := arguments[5]
	// connection string
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connString := fmt.Sprintf("p%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database)
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	rows, err := conn.Query(`SHOW DATABASE`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()
	// query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
	// rows, err = conn.Query(query)
	// if err != nil {
	// 	fmt.Println("Query", err)
	// 	return
	// }
	// for rows.Next() {
	// 	var name string
	// 	err = rows.Scan(&name)
	// 	if err != nil {
	// 		fmt.Println("Scan", err)
	// 		return
	// 	}
	// 	fmt.Println("+T", name)
	// }
	// defer rows.Close()
}
