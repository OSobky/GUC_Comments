package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// var (
// 	dbname     = os.Getenv("DATABASE_NAME")
// 	dbpassword = os.Getenv("DATABASE_PASSWORD")
// 	dbuser     = os.Getenv("DATABASE_USER")
// 	dbhost     = os.Getenv("DATABASE_HOST")
// )

const (
	dbname     = "GUC_Comments"
	dbpassword = "secret"
	dbuser     = "root"
	dbhost     = "db"
)

func main() {
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
	log.Printf("hey again")

	/*	http.HandleFunc("/", defaultHandler) // default directory
		http.HandleFunc("/MICRO", microHandler)
		http.HandleFunc("/ANALYSIS", analysisHandler)
		http.HandleFunc("/ANDVANCED", advancedHandler)
		http.ListenAndServe(":3000", nil)*/

}

/*func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Web!")
}
func microHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, micro!")
}
func analysisHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, analysis!")
}
func advancedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, advanced!")*/
//}
