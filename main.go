package main

import (
	"log"
	"net/http"
)

// Ini handler home yang dibuat sebagai function.
// Isinya  hanya sebuah string saja.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Snippetbox! Bro"))
}

func main() {
	// Mendeklarasikan sebuah variabel mux yang berisikan http.NewServeMux().
	// Terus memasukkan function home pada root "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Disini menggunakan log untuk mengeluarkan log mulai di server pada port 3000
	// Untuk menjalankan webserver menggunakan http.ListenAndServe()
	// Mengisikan nilai portnya kemudian variabel yang berisikan ServeMux() => mux
	log.Println("Starting server di port : 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
