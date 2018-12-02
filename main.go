package main

import (
	"fmt"
	"net/http"
	"os"

	"html/template"

	//"os"
	//"reflect"

	_ "github.com/lib/pq"
)

type guccomments struct {
	Title    string
	Comments string
}

// var (
// 	dbname     = os.Getenv("DATABASE_NAME")
// 	dbpassword = os.Getenv("DATABASE_PASSWORD")
// 	dbuser     = os.Getenv("DATABASE_USER")
// 	dbhost     = os.Getenv("DATABASE_HOST")
//
//)
var ac = os.Getenv("ACCESS_TOKEN")

/*const (
	dbname     = "GUC_Comments"
	dbpassword = "secret"
	dbuser     = "root"
	dbhost     = "db"
)
*/
func main() {
	/*
		q := ` SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'`
		qcreate := `CREATE TABLE Courses
					(
						CourseId SERIAL PRIMARY KEY,
						CourseName VARCHAR(30)

					);

					CREATE TABLE Comments
					(
						CourseId INTEGER REFERENCES Courses(CourseId),
						Comment VARCHAR(50)

					);`
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

		Result, err := db.Exec(qcreate)
		if (err) != nil {
			log.Fatal(err)
		}

		log.Print(Result)

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
		log.Printf("hey again")*/
	/*
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

			fmt.Println(msg["message"])
		}

		fmt.Println("")

		res2, _ := fb.Get("582313518881669_582750698837951/comments", fb.Params{
			"fields":       "message",
			"access_token": ac})

		r, _ := res2["data"].([]interface{})

		for i := 0; i < reflect.ValueOf(r).Len(); i++ {

			msg := r[i].(map[string]interface{})

			fmt.Println(msg["message"])
		}

		res3, _ := fb.Get("582313518881669_582751015504586/comments", fb.Params{
			"fields":       "message",
			"access_token": ac})

		v, _ := res3["data"].([]interface{})

		for i := 0; i < reflect.ValueOf(v).Len(); i++ {

			msg := v[i].(map[string]interface{})

			fmt.Println(msg["message"])
		}*/

	fmt.Print("hiii sobky")
	http.HandleFunc("/", defaultHandler) // default directory
	http.HandleFunc("/MICRO", microHandler)
	http.HandleFunc("/ANALYSIS", analysisHandler)
	http.HandleFunc("/ANDVANCED", advancedHandler)
	http.ListenAndServe(":3000", nil)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func microHandler(w http.ResponseWriter, r *http.Request) {
	p := guccomments{Title: "Guc Comments", Comments: "reviews"}
	t, _ := template.ParseFiles("guc_comments.html")
	t.Execute(w, p)
}
func analysisHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, analysis!")
}
func advancedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, advanced!")
}
