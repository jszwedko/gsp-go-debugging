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
	http.HandleFunc("/factorial/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")

		if len(pathParts) < 3 || pathParts[2] == "" {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		n, err := strconv.ParseInt(pathParts[2], 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not parse as integer: %s", pathParts[2]), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%d", math2.Factorial(n))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
