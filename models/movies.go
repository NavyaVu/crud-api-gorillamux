package models

type Movies struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Year int    `json:"year" db:"year"`
}
type AM struct {
	InMovies []InMovies
}
type InMovies struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}
