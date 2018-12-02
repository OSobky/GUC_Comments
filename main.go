package main

import (
	"database/sql"
	"fmt"

	"log"

	"os"
	"reflect"

	fb "github.com/huandu/facebook"
	_ "github.com/lib/pq"
)

var (
	dbname     = os.Getenv("DATABASE_NAME")
	dbpassword = os.Getenv("DATABASE_PASSWORD")
	dbuser     = os.Getenv("DATABASE_USER")
	dbhost     = os.Getenv("DATABASE_HOST")
	ac         = os.Getenv("ACCESS_TOKEN")
)

// const (
// 	dbname     = "GUC_Comments"
// 	dbpassword = "secret"
// 	dbuser     = "root"
// 	dbhost     = "db"
// )

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
						(DEFAULT,'CSEN 702 Microprocessors  '),
						(DEFAULT,'CSEN 703 Analysis and Design of Algorithms '),
						(DEFAULT,'CSEN 704 Advanced computer lab ');`

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
	/*	http.HandleFunc("/", defaultHandler) // default directory
		http.HandleFunc("/MICRO", microHandler)
		http.HandleFunc("/ANALYSIS", analysisHandler)
		http.HandleFunc("/ANDVANCED", advancedHandler)
		http.ListenAndServe(":3000", nil)*/

	// fmt.Print("hiii sobky")
	// http.HandleFunc("/", defaultHandler) // default directory
	// http.HandleFunc("/MICRO", microHandler)
	// http.HandleFunc("/ANALYSIS", analysisHandler)
	// http.HandleFunc("/ANDVANCED", advancedHandler)
	// http.ListenAndServe(":3000", nil)

}

// func defaultHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, Web!")
// }
// func microHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, micro!")
// }
// func analysisHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, analysis!")
// }
// func advancedHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, advanced!")
// }
