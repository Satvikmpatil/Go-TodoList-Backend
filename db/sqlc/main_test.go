package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open("postgres","postgresql://postgres:mypassword@localhost:5500/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("can not connect")
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}