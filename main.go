package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	x := "Gustavo Lopes"
	print(x)

	http.HandleFunc("/courses", listCourses)
	http.ListenAndServe(":8081", nil)
}

func listCourses(w http.ResponseWriter, r *http.Request) {
	persistCourse()
	w.Write([]byte("Hello World"))
}

func persistCourse() error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("insert into courses values($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec("1", "Curso Full Cycle")
	if err != nil {
		return err
	}
	return nil
}
