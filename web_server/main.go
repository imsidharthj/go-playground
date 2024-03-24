package main

import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "appliaction/json")
	w.Write([]byte(`{"name": "SJ"}`))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", rootHandler)
	r.HandleFunc("GET /profile", getProfileHandler)
	http.ListenAndServe("localhost:8080", r)
}
