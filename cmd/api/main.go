package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/amane15/url-shortner/internal/data"
	_ "github.com/lib/pq"
)

type application struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	urlModel    data.URLDataModel
}

func main() {
	db, err := openDB("postgres://postgres:password@localhost/shortener?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("database connection established")

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		urlModel:    data.URLDataModel{DB: db},
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}

	err = app.serve()
	if err != nil {
		errorLogger.Fatal(err)
	}
	infoLogger.Println(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
