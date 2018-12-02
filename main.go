package main

import (
	"database/sql"
	"fmt"

	"log"

	"net/http"
	"os"

	"html/template"
	"reflect"

	fb "github.com/huandu/facebook"
	_ "github.com/lib/pq"
)

type comment struct {
	Title    string
	Reviews  string
	Courses  []string
	Comments []string
}

var (
	dbname     = os.Getenv("DATABASE_NAME")
	dbpassword = os.Getenv("DATABASE_PASSWORD")
	dbuser     = os.Getenv("DATABASE_USER")
	dbhost     = os.Getenv("DATABASE_HOST")
	ac         = os.Getenv("ACCESS_TOKEN")
	db         *sql.DB
)

func main() {
	q := ` SELECT table_name 
			FROM information_schema.tables
			 WHERE table_schema='public' AND table_type='BASE TABLE'`
	q1 := `SELECT CourseId,CourseName
			FROM Courses `

	qcreate := `CREATE TABLE Courses
				(
					CourseId SERIAL PRIMARY KEY,
					CourseName TEXT

				);

				CREATE TABLE Comments
				(
					CourseId INTEGER REFERENCES Courses(CourseId),
					Comment TEXT

				);`
	qinsertCourses := `INSERT INTO Courses(CourseId ,CourseName ) VALUES
						(DEFAULT,'CSEN 702 Microprocessors'),
						(DEFAULT,'CSEN 703 Analysis and Design of Algorithms'),
						(DEFAULT,'CSEN 704 Advanced computer lab');`

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbuser, dbpassword, dbname)

	db, err := sql.Open("postgres",
		dbInfo)
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("connected")

	db.Exec(qcreate)

	Result, err := db.Exec(qinsertCourses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Result)

	rows, err := db.Query(q)
	if (err) != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var (
			tableName string
		)
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		log.Printf("Table name is %s", tableName)
	}
	rows2, err := db.Query(q1)

	for rows2.Next() {
		var (
			courseName string
			courseid   int
		)
		if err := rows2.Scan(&courseid, &courseName); err != nil {
			log.Fatal(err)
		}
		log.Printf("course name is %s and id is %d", courseName, courseid)
	}
	log.Printf("hey again")

	res, err := fb.Get("582313518881669_582751342171220/comments", fb.Params{
		"fields":       "message",
		"access_token": ac,
	})

	if (err) != nil {
		log.Fatal(err)

	}

	s, _ := res["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(s).Len(); i++ {

		msg := s[i].(map[string]interface{})

		qinsertcomment := fmt.Sprintf(`INSERT INTO Comments(CourseId, Comment) VALUES (1 ,'%s')`,
			msg["message"])

		Result, err := db.Exec(qinsertcomment)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Result)
		fmt.Println(qinsertcomment)
	}

	fmt.Println("")

	res2, _ := fb.Get("582313518881669_582750698837951/comments", fb.Params{
		"fields":       "message",
		"access_token": ac})

	r, _ := res2["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(r).Len(); i++ {

		msg := r[i].(map[string]interface{})

		qinsertcomment := fmt.Sprintf(`INSERT INTO Comments(CourseId, Comment) VALUES (2 ,'%s')`,
			msg["message"])

		Result, err := db.Exec(qinsertcomment)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Result)
		fmt.Println(qinsertcomment)
	}

	fmt.Println("")

	res3, _ := fb.Get("582313518881669_582751015504586/comments", fb.Params{
		"fields":       "message",
		"access_token": ac})

	v, _ := res3["data"].([]interface{})

	for i := 0; i < reflect.ValueOf(v).Len(); i++ {

		msg := v[i].(map[string]interface{})

		qinsertcomment := fmt.Sprintf(`INSERT INTO Comments(CourseId, Comment) VALUES (3 ,'%s')`,
			msg["message"])

		Result, err := db.Exec(qinsertcomment)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Result)
		fmt.Println(qinsertcomment)
	}
	fmt.Println("3ash ya sobkyy")

	fmt.Print("hiii sobky")
	http.HandleFunc("/", defaultHandler) // default directory
	http.HandleFunc("/MICRO", microHandler)
	http.HandleFunc("/ANALYSIS", analysisHandler)
	http.HandleFunc("/ANDVANCED", advancedHandler)
	http.ListenAndServe(":8080", nil)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	courses2 := gettingCourseFromCourses()
	comments2 = 
	p := guccomments{Title: "Guc Comments", Reviews: "Reviews", Courses: courses2, Comments: []string{}}
	t, _ := template.ParseFiles("guc_comments.html")
	t.Execute(w, p)
}

func microHandler(w http.ResponseWriter, r *http.Request) {
	p := guccomments{Title: "Guc Comments", Reviews: "Reviews", Comments: []string{"asdasds", "asdasdas", "asdasd"}}
	t, _ := template.ParseFiles("guc_comments.html")
	t.Execute(w, p)
}
func analysisHandler(w http.ResponseWriter, r *http.Request) {
	p := guccomments{Title: "Guc Comments", Reviews: "Reviews", Comments: []string{"asdasds", "asdasdas", "asdasd"}}
	t, _ := template.ParseFiles("guc_comments.html")
	t.Execute(w, p)
}
func advancedHandler(w http.ResponseWriter, r *http.Request) {
	p := guccomments{Title: "Guc Comments", Reviews: "Reviews", Comments: []string{"asdasds", "asdasdas", "asdasd"}}
	t, _ := template.ParseFiles("guc_comments.html")
	t.Execute(w, p)
}

func gettingCommentsFromCourse(CourseName string) []string {
	var comments []string
	q := fmt.Sprintf(`SELECT Comment
			 FROM Comments  
			 INNER JOIN Courses ON  Courses.CourseId = Comments.CourseId 
			 WHERE Courses.CourseName = '%s' `, CourseName)
	rows, err := db.Query(q)

	if (err) != nil {
		log.Fatal(err)
	}
	i := 0
	for rows.Next() {
		var (
			comment string
		)
		if err := rows.Scan(&comment); err != nil {
			log.Fatal(err)
		}
		comments[i] = comment
		i++

	}
	return comments
}

func gettingCourseFromCourses() []string {
	var courses []string
	q := fmt.Sprintf(`SELECT Course
			 		FROM Courses `)

	rows, err := db.Query(q)

	if (err) != nil {
		log.Fatal(err)
	}
	i := 0
	for rows.Next() {
		var (
			course string
		)
		if err := rows.Scan(&course); err != nil {
			log.Fatal(err)
		}
		courses[i] = course
		i++

	}
	return courses
}
