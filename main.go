package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	hello := os.Getenv("MY_VARIABLE")
	fmt.Fprintf(w, "SERVICE2 MY_VARIABLE111:%s", hello)
}
