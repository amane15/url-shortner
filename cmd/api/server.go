package main

import (
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        ":4000",
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
	}

	app.infoLogger.Println("staring server on port: 4000")

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
