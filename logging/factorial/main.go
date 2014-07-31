package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../."
)

func main() {
	log.SetFlags(LstdFlags | Lshortfile)

	http.HandleFunc("/factorial/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received request for %s", r.URL)

		pathParts := strings.Split(r.URL.Path, "/")

		if len(pathParts) < 3 || pathParts[2] == "" {
			log.Printf("returning 404")
			http.Error(w, "", http.StatusNotFound)
			return
		}

		log.Printf("parsing resource %s as integer", pathParts[1])
		n, err := strconv.ParseInt(pathParts[1], 10, 64)
		if err != nil {
			log.Printf("returning 400: %v", err)
			http.Error(w, fmt.Sprintf("could not parse as integer: %s", pathParts[1]), http.StatusBadRequest)
			return
		}

		// time.Sleep(200 * time.Millisecond)

		log.Printf("calculating factorial for %d", n)

		f := math2.Factorial(n)

		// time.Sleep(200 * time.Millisecond)

		log.Printf("responding with %d", f)

		fmt.Fprintf(w, "%d", f)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println()
}
