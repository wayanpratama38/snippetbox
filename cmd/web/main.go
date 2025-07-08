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

	// buat sebuah variabel fileServer dengan http.FileServer yang langsung ke
	// File static untuk UI menggunakan http.Dir() bisa langsung panggil dari root
	// page.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Menambahkan endpoint /static/ untuk menyimpan file statis
	// disini ketika sebuah permintaan berupa file di dalam static
	// maka prefix /static akan dihilangkan
	// contoh : ui/static/gambar.jpg
	// maka karena StripPrefix("/static")
	// hasilnya menjadi /gambar.jpg
	// sehingga jika diteruskan ke fileServer
	// maka akan mencari file dengan directory ./ui/static/gambar.jpg
	// (mengabaikan // pada gambar.jpg)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Disini menggunakan log untuk mengeluarkan log mulai di server pada port 3000
	// Untuk menjalankan webserver menggunakan http.ListenAndServe()
	// Mengisikan nilai portnya kemudian variabel yang berisikan ServeMux() => mux
	log.Println("Starting server di port : 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
