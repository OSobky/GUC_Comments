package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//import "net/http"

func main() {
	db, err := sql.Open("postgres",
		"user=root password=secret dbname=GUC_Comments sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("connected")

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
