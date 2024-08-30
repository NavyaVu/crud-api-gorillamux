package main

import (
	"crud-api-gorillamux/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

func main() {

	db, err = getSqlConnection()

	if err != nil {
		log.Panic(err)
	}
	//else{
	// 	//fmt.Println("Connection success with db version: ", version)

	// 	// 		st := `
	// 	// 		SELECT name FROM sqlite_master
	// 	// WHERE type='table'
	// 	// ORDER BY name;`

	// 	// sts := `
	// 	// DROP table movies;`

	// 	// sts := `
	// 	// create table if not exists movies(id INTEGER PRIMARY KEY, name TEXT, year INT);;
	// 	//"DELETE FROM cars WHERE id IN (1, 2, 3)")

	// 	//UPDATE table_name
	// 	// SET column1 = value1, column2 = value2, ...
	// 	// WHERE condition;
	// 	// `
	// 	// sts := `
	// 	// CREATE TABLE IF NOT EXISTS movies(id INTEGER PRIMARY KEY, name TEXT, year INT);
	// 	// `
	// 	// sts := `
	// 	// select * from movies;

	// 	//db.Exec("INSERT INTO users(name) VALUES(?)","Dolly",1994)
	// 	// `
	// 	//var v []any
	// 	// v := []any{"bold", 1993}
	// 	// //"INSERT INTO movies(name,year) VALUES(?,?);", v...
	// 	// r, err := db.Exec(models.InsertData(), v...)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error while exec")
	// 	// 	log.Fatal(err)
	// 	// } else {
	// 	// 	ar, err := r.LastInsertId()
	// 	// 	if err != nil {
	// 	// 		fmt.Println("Error while checking rows affected")
	// 	// 		log.Fatal(err)
	// 	// 	}

	// 		fmt.Println("value of ar", ar)
	// 	}

	// 	sts = `
	// 	select * from movies;
	// 	`

	// 	rows, err := db.Query(sts)
	// 	if err != nil {
	// 		fmt.Println("Error while query")
	// 		log.Fatal(err)
	// 	}

	// 	defer rows.Close()

	// 	// if !rows.Next() {
	// 	// 	fmt.Println("no tables")
	// 	// } else {
	// 	for rows.Next() {
	// 		var id int
	// 		var name string
	// 		var year int

	// 		err = rows.Scan(&id, &name, &year)
	// 		if err != nil {
	// 			fmt.Println("Error while scanning")
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Printf("%d %s %d\n", id, name, year)
	// 	}
	// 	// 	}
	// 	// }

	// 	defer db.Close()

	// }

	fmt.Println("Starting the server on 8090 port")

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", insertMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	y, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "Id not recognised")
	} else {
		models.DeleteOnId(db, y)
	}
}
func insertMovies(w http.ResponseWriter, r *http.Request) {
	var movies []*models.InMovies
	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		log.Panic(err)
	} else {
		err := models.InsertData(db, movies)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			fmt.Fprint(w, "Error during insertion of data")
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, "Inserted the data: ", len(movies))
		}
		// 	fmt.Fprint(w, "Inserting the intial data into db")
		//fmt.Println("Length:", len(movies))
	}
}

//	func insertInitialData(w http.ResponseWriter, r *http.Request) {
//		models.CreateTable(db)
//		w.Header().Set("Content-Type", "application/json")
//		fmt.Fprint(w, "Inserting the intial data into db")
//	}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//request to display all movies

	//name := mux.Vars(r)["name"]
	//fmt.Println("name:", name)
	s := models.SelectAll(db)
	for _, structs := range s {
		//if structs.Name == name {
		err := json.NewEncoder(w).Encode(&structs)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
		//}
	}

}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("id1:", id)

	y, err := strconv.Atoi(id)
	fmt.Println("id:", y)

	//y := 2

	if err != nil {
		log.Println("Error during a to i convertion", err)
	} else {
		m := models.SelectOnId(db, y)
		err := json.NewEncoder(w).Encode(&m)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func getSqlConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version of db:", version)

	return db, err
}
