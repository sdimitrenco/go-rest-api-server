package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Hello world, %q ", html.EscapeString(request.URL.Path))

	})

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
