package side_packages

import (
	"fmt"
	"net/http"
)

// Функция выводит надпись на экран
func index_handler(w http.ResponseWriter, r *http.Request){
	// Fprintf первый аргумент - куда записывать, второй - что записывать
	fmt.Fprintf(w, "<h1>Whoa, Go is neat</h1>")
}

// Функция активации сервера на локальной машине
func HTTPServer(){
	// указывается url и функция, его обрабатывающая
	http.HandleFunc("/test_route", index_handler)
	// указывается порт, который слушает сервер
	// второй аргумент - handler, обычно он nil
	http.ListenAndServe(":8000", nil)
}

// чтобы увидеть результаты: http://localhost:8000/test_route