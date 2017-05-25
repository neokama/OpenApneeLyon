package main
	
	import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	)
	
	
	type Bdd struct
	{
		cheminbdd string
		resultat *sql.Rows
		db *sql.DB
		err error
	}

	
	/*
	* 		Bdd.connection:
	* Description: 
	* 		Méthode permettant de se connecter à la base de 
	* 		données située au chemin contenu dans l'attribut "cheminBdd".
	*/
	
	func (base Bdd) connection() (){
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}
	
	
	
	/*
	* 		Bdd.reset:
	* Description: 		
	*		Méthode permettant de supprimer tous les compétiteurs, classement, et classements par équipe contenus dans la base de
	*		données.
	*/
	
	func (base Bdd) reset(){
		_, base.err = base.db.Exec("DELETE FROM competiteurs")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
		} else {
			_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='competiteurs'")
			if base.err != nil {
				fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
				} else {
					_, base.err = base.db.Exec("DELETE FROM classement")
					if base.err != nil {
						fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
					} else {
							_, base.err = base.db.Exec("DELETE FROM classementequipe")
						if base.err != nil {
							fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
						} else {
							_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='classementequipe'")
							if base.err != nil {
								fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
								}else{
								fmt.Println("Remise à zéro de la base de données effectuée")
							}
						}
					}
				
				}
			}
		}
	
	
	/*
	* 		newBdd:
	* Paramètres:
	*	- cheminBdd:  Chemin et nom de la base de données à ouvrir.
	*
	* Description: 		
	*		Méthode permettant d'ouvrir une connection vers une base de données.
	*/

	func newBdd(cheminBdd string)(*Bdd){
		base := new(Bdd)
		base.cheminbdd = cheminBdd
		
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
		base.resultat = new(sql.Rows)
		
		return base
	}