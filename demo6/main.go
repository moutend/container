package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		`{"labels":{"myLabel1":"value1","myLabel2":"value2"},"severity":"DEBUG","timestamp":%q,"method":%q,"url":%q,"remote":%q}`+"\n",
		time.Now().UTC().Format(time.RFC3339Nano),
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
	magicWord := os.Getenv("MAGIC_WORD")
	fmt.Fprintf(w, "Hi, %q!\n", magicWord)
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
