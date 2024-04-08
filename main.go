package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	// MySQL veritabanına bağlan
	err := ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	defer db.Close()

	// HTTP sunucusunu başlat
	StartHTTPServer()
}

// ConnectDB MySQL veritabanına bağlanır
func ConnectDB() error {
	// MySQL DSN bilgileri
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Veritabanı bağlantısını test et
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to MySQL")
	return nil
}

// StartHTTPServer HTTP sunucusunu başlatır
func StartHTTPServer() {
	// HTTP sunucusu başlatılıyor
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HomeHandler örnek bir HTTP handler fonksiyonudur
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the server!")
}
