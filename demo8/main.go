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
		`{"logging.googleapis.com/labels":{"myLabel1":"value1","myLabel2":"value2"},"severity":"DEBUG","time":%q,"httpRequest":{"requestMethod":%q,"requestUrl":%q,"referrer":%q},"logging.googleapis.com/logName":"myLogName"}`+"\n",
		time.Now().UTC().Format(time.RFC3339Nano),
		r.Method,
		r.URL,
		r.RemoteAddr,
	)
	magicWord := os.Getenv("MAGIC_WORD")
	fmt.Fprintf(w, "Hola, %q!\n", magicWord)
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
