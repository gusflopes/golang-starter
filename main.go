package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
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
	e := echo.New()
	e.GET("/courses", listCourses)
	e.Logger.Fatal(e.Start(":8081"))

}

func listCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
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
