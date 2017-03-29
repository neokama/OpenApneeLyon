	package main
	
	import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"bufio"
	"strings"
	)
	
	
	type Bdd struct
	{
		cheminbdd string
		chemincsv string
		resultat *sql.Rows
		db *sql.DB
		err error
	}
	
	func (base Bdd) connection() (){
		
	
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}
	
	
	func (base Bdd) disp_requete(sqlQuery string){
	
		base.resultat, base.err = base.db.Query(sqlQuery)
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
		
		var nom string
		var prenom string
		
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&nom, &prenom)
			if base.err != nil {
			log.Fatal(base.err)
			}
		fmt.Println(nom, "|", prenom)
		}
		
		
	}
	
	func (base Bdd) addComp(nom string, prenom string){
		
		_, base.err = base.db.Exec("INSERT INTO competiteurs(nom, prenom) VALUES('" + nom + "','" + prenom + "')")
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout: \n", base.err)
			} else {
			fmt.Println("Ajout validé du compétiteur " + nom +" "+ prenom)
		}
	}
	
	func (base Bdd) delComp(col string, value string){
		
		_, base.err = base.db.Exec("DELETE FROM competiteurs WHERE " + col + " = " + value)
		if base.err != nil {
			fmt.Println("Echec lors de la suppression: \n", base.err)
			} else {
			fmt.Println("Suppression des competiteurs avec " + col + " = " + value)
		}
	}
	
	
	func (base Bdd) requete_export(sqlQuery string){
	
		base.resultat, base.err = base.db.Query(sqlQuery)
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
		
		var nom string
		var prenom string
		file, err := os.Create("fichierTest.csv")
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier")
				log.Fatal(err)
			}
		

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&nom, &prenom)
			if base.err != nil {
			log.Fatal(base.err)
			}
		file.WriteString(fmt.Sprint(nom, ";", prenom,"\r\n"))
		}	

	}
	
	
	
	func (base Bdd) importComp(chemin string){
		file, err := os.Open(chemin)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		var info []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			info = strings.Split(scanner.Text(), ";")
			base.addComp(info[0], info[1])
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	
	
	
	
	
	
	func newBdd(cheminBdd string, chemincsv string)(*Bdd){
		base := new(Bdd)
		base.cheminbdd = cheminBdd
		base.chemincsv = chemincsv
		
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
		base.resultat = new(sql.Rows)
		
		return base
	}
	