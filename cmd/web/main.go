package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")

	flag.Parse()

	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := ConnectDb()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
	}

	fileServer := http.FileServer(http.Dir("./static"))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("POST /create-wish", app.createWish)
	mux.HandleFunc("GET /success", app.wishCreated)
	mux.HandleFunc("GET /c/{slug}", app.retrieveWishMessage)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Server started running on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
