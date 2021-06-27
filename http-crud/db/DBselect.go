package db

import (
	"database/sql"
	"http-db/model"
	"log"
)
import _ "github.com/lib/pq"

const QRY = "SELECT id, firstname, lastname, age FROM public.people"

func GetPerson(id int) (model.Person, error) {
	person := model.Person{}
	db, err := sql.Open("postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		return person, err
	}
	defer db.Close()

	qryrow, err := db.Prepare(QRY + " WHERE id=$1")
	if err != nil {
		return person, err
	}

	err = qryrow.QueryRow(id).Scan(&person.Id, &person.Firstname, &person.Lastname, &person.Age)
	if err != nil {
		return person, err
	}
	return person, nil
}

func GetPeople() ([]model.Person, error) {
	db, err := sql.Open("postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(QRY)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, age int
	var firstname, lastname string
	var result []model.Person
	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname, &age)
		if err != nil {
			return nil, err
		}
		tmp := model.Person{
			Id:        id,
			Firstname: firstname,
			Lastname:  lastname,
			Age:       age,
		}
		result = append(result, tmp)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	log.Printf("Retrieved: %v people\n", len(result))
	return result, nil
}
