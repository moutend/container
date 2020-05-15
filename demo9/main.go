package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func handler(w http.ResponseWriter, r *http.Request) {
	var count int

	row := db.QueryRow(`SELECT COUNT(id) FROM foo;`)
	err := row.Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		`{"logging.googleapis.com/labels":{"myLabel1":"value1","myLabel2":"value2"},"severity":"DEBUG","time":%q,"httpRequest":{"requestMethod":%q,"requestUrl":%q,"referrer":%q},"logging.googleapis.com/logName":"myLogName"}`+"\n",
		time.Now().UTC().Format(time.RFC3339Nano),
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
	magicWord := os.Getenv("MAGIC_WORD")
	fmt.Fprintf(w, "Count returns %d, %q!\n", count, magicWord)
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error

	db, err = sql.Open("mysql", source)

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
