package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/awilson506/med-bridge-hep/server"
	_ "github.com/lib/pq"
)

func main() {

	//TODO: should live in the env
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=med_bridge_hep sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//start the server
	s := server.NewServer(db)
	serverErr := s.Start()

	if serverErr != http.ErrServerClosed {
		log.Println(serverErr)
	}
}
