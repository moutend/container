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
		`{"timestamp":%q,"method":%q,"url":%q,"remote":%q}`+"\n",
		time.Now().UTC().Format(time.RFC3339),
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
	magicWord := os.Getenv("MAGIC_WORD")
	fmt.Fprintf(w, "The secret keyword is %q", magicWord)
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
