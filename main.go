// Latihan bahasa pemrograman GO di Sekolah SMP ASTRA AGRO LESTARI
// Materi : Membuat sebuah web app sederhana

package main

import (
         "fmt" 
         "net/http"
         "os"
        )


func main() {

 port := Getenv("PORT")

 http.HandleFunc("/", func (w.httpResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is a simple app!")
 }

 fs := http.FileServer(http.Dir("public/"))
 http.Handle("/public/", http.StripPrefix("/public/", fs))
 
 http.ListenAndServe(":%d", port)
}
