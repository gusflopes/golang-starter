package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var courses []Course

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"course_name"`
}

func generateCourses() {
	course1 := Course{
		ID:   1,
		Name: "Full Cycle",
	}
	course2 := Course{
		ID:   2,
		Name: "Bonus Full Cycle",
	}
	courses = append(courses, course1, course2)
}

func main() {
	generateCourses()
	http.HandleFunc("/courses", listCourses)
	http.ListenAndServe(":8081", nil)
}

func listCourses(w http.ResponseWriter, r *http.Request) {
	jsonCourses, err := json.Marshal(courses)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(jsonCourses))
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
