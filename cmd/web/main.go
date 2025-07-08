package main

import (
	"log"
	"net/http"
)

func main() {
	// Mendeklarasikan sebuah variabel mux yang berisikan http.NewServeMux().
	// Terus memasukkan function home pada root "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// Menambahkan endpoint /snippet dan /snippet/create ke dalam webserver
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Disini menggunakan log untuk mengeluarkan log mulai di server pada port 3000
	// Untuk menjalankan webserver menggunakan http.ListenAndServe()
	// Mengisikan nilai portnya kemudian variabel yang berisikan ServeMux() => mux
	log.Println("Starting server di port : 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
