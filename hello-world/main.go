package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World! \n")
	})

	http.ListenAndServe(":3000", nil)
	fmt.Println("Listening on port 3000")
}
