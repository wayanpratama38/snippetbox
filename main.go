package main

import (
	"log"
	"net/http"
)

// Ini handler home yang dibuat sebagai function.
// Isinya  hanya sebuah string saja.
func home(w http.ResponseWriter, r *http.Request) {

	// Karena bersifat catch-all maka setiap endpoint yang
	// diakhiri dengan "/" akan langsung di route ke dalam function home
	// bisa diatasi dengan conditional
	if r.URL.Path != "/" {
		// ini akan mengembalikan 404 ke user
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello From Snippetbox! Bro"))
}

// Ini sebuah function yang bertindak sebagai handler
// untuk routing ke endpoint /snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow from show snippet route!"))
}

// Ini sebuah function yang bertindak sebagai handler
// untuk routing ke endpoint /snippet/create
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating new snippet..."))
}

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
