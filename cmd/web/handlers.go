package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	// Inisialisasi sebuah array yang berisikan string
	// yang isinya adalah slice dari setiap template html
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Pertama kita baca dulu dan parsing terlebih dahulu file home.page.tmpl
	// Jika ternyata hasilnya bukan nil maka tampilkan errornya
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

// Ini sebuah function yang bertindak sebagai handler
// untuk routing ke endpoint /snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Disini menggunakan strconv.Atoi untuk mengekstrasi urlnya
	// terus mendapatkan query dengan prameter id
	// kemudian perlu pengecekan apakah errnya bukan nil
	// dan nilai id harus lebih dari 0 atau positif.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Print response sesuai dengan id yang didapatkan lewat query parameter
	fmt.Fprintf(w, "Display spesific snippet from id %d", id)
	// Menuliskan basic pada route saja
	// w.Write([]byte("Hellow from show snippet route!"))
}

// Ini sebuah function yang bertindak sebagai handler
// untuk routing ke endpoint /snippet/create
func createSnippet(w http.ResponseWriter, r *http.Request) {

	// Menggunakan r.Method untuk mengecek method apa yang digunakan
	// jika method yang digunakan bukan POST maka akan
	// mengembalikan response 405 (method not allowed)
	if r.Method != "POST" {
		// Tambahkan di header kalau yang diperbolehkan hanya method POST
		// Maka ketika user mendapatkan response 405 (method not allowed)
		// Dia mengetahui bahwa hanya bisa menggunakan method POST saja
		w.Header().Set("Allow", "POST")
		// Harus menuliskan WriteHeader terlebih dahulu baru Write
		// Jika tidak maka akan otomatis dibaca sebagai status code 200 (success)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed!"))
		// menggunakan http.Error() untuk langsung memberikan response code, dan body messagenya
		http.Error(w, "Method not allowed!", 405)
		return
	}
	w.Write([]byte("Creating new snippet..."))
}
