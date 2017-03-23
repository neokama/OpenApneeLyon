package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	)

func main() {

    db, err := sql.Open("sqlite3", "../src/database/test")
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()
	
	
	rows, err := db.Query("SELECT * FROM competiteurs")
	if err != nil {
		fmt.Println("C'est ici qu'il y a une erreur je crois...")
		log.Fatal(err)
	}
	defer rows.Close();
	
	for rows.Next() {
		var nom string
		var prenom string
		err = rows.Scan(&nom, &prenom)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(nom, "|", prenom)
	}
	
	//Ne marche pas!!
	//db.Query(".separator \";\"")
	//db.Query(".import ../src/database/test.csv competiteurs")
	
	
}