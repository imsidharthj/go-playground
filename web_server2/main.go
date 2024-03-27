package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/sj", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte)("web server"))
	})
	http.ListenAndServe(":8080", nil)
}
