package dbinit

import (
	"database/sql"
	"fmt"
	"http-db/myrnd"
	"log"
)
import _ "github.com/lib/pq"

/* Создать в БД таблицу (tableName) */
func createTable(tableName string, db *sql.DB) error {
	DBCreate := fmt.Sprintf(`
	CREATE TABLE public.%v (
		id integer,
		firstname character varying COLLATE pg_catalog."default",
		lastname character varying COLLATE pg_catalog."default",
		age integer
	) WITH ( OIDS = FALSE )
	TABLESPACE pg_default;
	ALTER TABLE public.%v OWNER to postgres;`, tableName, tableName)
	_, err := db.Exec(DBCreate)
	return err
}

/* Наполнить таблицу случайными записями */
func populateTable(tableName string, db *sql.DB) error {

	qry := fmt.Sprintf("INSERT INTO public.%v (id, firstname, lastname, age) VALUES ($1, $2, $3, $4)", tableName)
	insert, err := db.Prepare(qry)
	if err != nil {
		log.Panic(err)
	}
	defer insert.Close()

	for i := 0; i < 100; i++ {
		firstname := myrnd.GetRandString(8)
		lastname := myrnd.GetRandString(16)
		age := myrnd.GetRandAge()

		_, err = insert.Exec(i, firstname, lastname, age)
		if err != nil {
			return err
		}
	}
	return nil
}

/* Создать в БД таблицу (tableName) и заполнить ее некоторым количеством записей (entityAmount) */
func DBInit(tableName string) error {
	// Подключиться к БД
	db, err := sql.Open("postgres", "user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	defer db.Close()

	// Создать таблицу 'tableName'
	err = createTable(tableName, db)
	if err != nil {
		log.Panic(err)
		return err
	} else {
		log.Printf("The table %v was successfully created!\n", tableName)
	}

	// Наполнить таблицу 'tableName' случайными значениями
	err = populateTable(tableName, db)
	if err != nil {
		log.Panic(err)
		return err
	} else {
		log.Printf("The table %v was successfully populated!\n", tableName)
	}

	return nil
}
