package main

import "net/http"

func main() {
	x := "Gustavo Lopes"
	print(x)

	http.HandleFunc("/courses", listCourses)
	http.ListenAndServe(":8081", nil)
}

func listCourses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
