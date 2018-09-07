package main

import "net/http"

const port = ":9000"

func main() {
	http.HandleFunc("/", func(rec http.ResponseWriter, req *http.Request) {
		rec.WriteHeader(http.StatusOK)
		rec.Write([]byte("OK"))
	})
	http.ListenAndServe(port, nil)
}
