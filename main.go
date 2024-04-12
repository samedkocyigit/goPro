package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"goProject/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// MySQL veritabanına bağlan
	err := ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	defer db.Close()

	routes.SetupRoutes()
	// HTTP sunucusunu başlat
	StartHTTPServer()
}

// ConnectDB MySQL veritabanına bağlanır
func ConnectDB() error {
	// MySQL DSN bilgileri
	dsn := "samed123:samed123@tcp(127.0.0.1:3306)/goProject"

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

	if err := godotenv.Load("config.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	// HTTP sunucusu başlatılıyor
	port := os.Getenv("PORT")
	host := os.Getenv("LOCALHOST")

	// Değerleri kullanarak HTTP sunucusunu başlat
	log.Printf("Starting server on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
