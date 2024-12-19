package main

import (
	"net/http"
	//"testapp/styles"
	//"testapp/scripts"
)

func main() {
	// Serve HTML files
	http.HandleFunc("/", rootHandler)

	// Serve CSS and JavaScript files from /home/larrry/laura2/
	http.Handle("/testapp/styles/", http.StripPrefix("/testapp/styles/", http.FileServer(http.Dir("testapp/styles"))))
	// Serve JavaScript file
	http.Handle("/testapp/scripts/", http.StripPrefix("/testapp/scripts/", http.FileServer(http.Dir("testapp/scripts"))))
	/*
	   http.HandleFunc("/scripts/script.js", func(w http.ResponseWriter, r *http.Request) {
	       w.Header().Set("Content-Type", "application/javascript")
	       http.ServeFile(w, r, "scripts/script.js")
	   })*/

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "testapp/index.html")
}
