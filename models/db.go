package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB) string {
	sts := `
		create table if not exists movies(id INTEGER PRIMARY KEY, name TEXT, year INT);
		INSERT INTO movies(name, year) VALUES('Audi',52642);
		INSERT INTO movies(name, year) VALUES('Mercedes',57127);
		INSERT INTO movies(name, year) VALUES('Skoda',9000);
		INSERT INTO movies(name, year) VALUES('Volvo',29000);
		INSERT INTO movies(name, year) VALUES('Bentley',350000);
		INSERT INTO movies(name, year) VALUES('Citroen',21000);
		INSERT INTO movies(name, year) VALUES('Hummer',41400);
		INSERT INTO movies(name, year) VALUES('Volkswagen',21600);
		`

	_, err := db.Exec(sts)
	if err != nil {
		fmt.Println("Error while inserting the data")
		log.Panic("Error while inserting the data", err)
	}
	return sts
}

func InsertData(db *sql.DB, movies []*InMovies) error {
	//var v []interface
	var err error
	var r sql.Result
	for _, m := range movies {
		sts := "INSERT INTO movies(name,year) VALUES(?,?);"
		r, err = db.Exec(sts, m.Name, m.Year)
		if err != nil {
			log.Panicln("Error during insertion of data:", err)
		} else {

			ar, err := r.LastInsertId()
			if err != nil {
				fmt.Println("Error while checking rows affected")
				log.Fatal(err)
			}

			fmt.Println("value of ar", ar)
		}
	}
	return err
}

func SelectAll(db *sql.DB) []Movies {

	fmt.Println("In select section")
	sts := "Select * from movies"
	var movies = make([]Movies, 0)
	var movie Movies
	var (
		id   int
		name string
		year int
	)

	rows, err := db.Query(sts)
	if err != nil {
		log.Panicln("Error during select all", err)
	} else {
		//defer rows.Close()
		if !rows.Next() {
			fmt.Println("no tables")
		} else {
			for rows.Next() {
				err = rows.Scan(&id, &name, &year)
				if err != nil {
					fmt.Println("Error while scanning")
					log.Fatal(err)
				}

				movie = Movies{
					Id:   id,
					Name: name,
					Year: year,
				}
				movies = append(movies, movie)
			}
		}

	}

	return movies
}

func DeleteOnId(db *sql.DB, id int) {
	sts := "DELETE FROM movies WHERE id = ?"
	r, err := db.Exec(sts, id)

	if err != nil {
		log.Println("Error while deleting the data", err)
	} else {
		ar, err := r.RowsAffected()
		if err != nil {
			fmt.Println("Error while checking rows affected")
			log.Fatal(err)
		}

		fmt.Println("value of ar", ar)
	}
}

func SelectOnId(db *sql.DB, id int) Movies {
	sts := "Select * from movies where id = ?"

	var movie Movies
	var (
		name string
		year int
	)

	rows := db.QueryRow(sts, id)
	err := rows.Scan(&id, &name, &year)
	if err != nil {
		fmt.Println("Error while scanning")
		log.Fatal(err)
	}

	movie = Movies{
		Id:   id,
		Name: name,
		Year: year,
	}

	return movie
}
