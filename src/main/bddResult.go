package main
	
	import (
	"strconv"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"bufio"
	"strings"
	//"regexp"
	//"time"
	)
	
	
	type BddResult struct
	{
		cheminbdd string
		chemincsv string
		resultat *sql.Rows
		db *sql.DB
		err error
	}
	
	/*
	* 		Bdd.connection:
	* Description: 
	* 		Méthode permettant de se connecter à la base de 
	* 		données située au chemin contenu dans l'attribut "cheminBdd"
	*/	
	
	func (base BddResult) connection() (){
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}
	
	/*
	* 		Bdd.reset:
	* Description: 		
	*		Méthode permettant de supprimer tous les compétiteurs contenus dans la base de
	*		données.
	*/
	
	func (base BddResult) reset(){
		_, base.err = base.db.Exec("DELETE FROM classement")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
		} else {
			_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='classement'")
			if base.err != nil {
				fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
				} else {
				fmt.Println("Remise à zéro de la base de données effectuée")
			
			}
		}
	}
	/*
	* 		Bdd.disp_comp:
	* Description: 	
	* 		Méthode permettant d'afficher l'integralité des
	* 		competiteurs de la base de données
	*/
	
	func (base BddResult) displayCompetiteur(){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM classement")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [7]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6])
		}
	}
	
	/*
	* 		Bdd.searchCompetiteur:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de rechercher un compétiteur en
	* 		competiteurs de la base de données
	*/
	
	func (base BddResult) searchCompetiteur(col_num int, value string){
		
		var id_col string
		var searchValue string
		
		searchValue = fmt.Sprint("'%",value,"%'")
		id_col, value = col_id2name(col_num, value)
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE ", id_col, " LIKE ", searchValue))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [7]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6])
		}
	}
	
	
		/*
	* 		Bdd.addCompetiteur:
	* Paramètres:
	*	- comp: 	Les informations du compétiteur à ajouter sous la
	*				forme d'une structure de type "competiteur"
	*
	* Description: 		
	*		Méthode permettant d'ajouter un compétiteur dans la 
	* 		base de données
	*/

	func (base BddResult) addCompetiteur(board *Classement){
		
		_, base.err = base.db.Exec("INSERT INTO classement ( prenom, nom, sexe, equipe, epreuve, resultat) VALUES('" +
		board.prenom + "','" +
		board.nom + "','" +
		board.sexe + "','" +
		board.equipe + "','" +
		board.epreuve + "'," +
		strconv.Itoa(board.resultat) + ")")
		
		
		
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : "+ board.nom +" "+ board.prenom, base.err)
			} else {
			fmt.Println("Ajout validé du resulat compétiteur " + board.nom +" "+ board.prenom)
		}
	}
	
	/*
	* 		Bdd.deleteCompetiteur:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de supprimer les compétiteurs en fonction des critères
	*		en entrée.
	*/

	func (base BddResult) deleteCompetiteur(col_num int, value string){
		var id_col string
		value = fmt.Sprint("'",value,"'")
		
		if col_num==1 {
			id_col = "id"		
		} else if col_num==2{
			id_col = "equipe"		
		}
		
		if !(col_num < 1 && col_num > 2){
			_, base.err = base.db.Exec("DELETE FROM classement WHERE " + id_col + " = " + value)
			if base.err != nil {
				fmt.Println("Echec lors de la suppression: \n", base.err)
				} else {
				fmt.Println("Suppression des competiteurs avec " + id_col + " = " + value)
			}
		} else {
			err := "Le numéro entré est invalide!"
			fmt.Println(err);
		}
	}
	
	
	/*
	* 		Bdd.importCompetiteur:
	* Paramètres:
	*	- chemin: 	Chemin du fichier à importer avec le nom du fichier et l'extension.
	*
	* Description: 		
	*		Méthode permettant d'importer les compétiteurs contenu dans un fichier CSV
	*/
	
	func (base BddResult) importResultat(){
		file, err := os.Open("import/classement.csv")
		if err != nil {
			println("Impossible d'ouvrir le fichier \"classement.csv\" dans le dossier import")
			log.Fatal(err)
		}
		defer file.Close()
	
		var firstCall bool	
		firstCall = true
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ";")
			if !firstCall{
			temps,er := strconv.Atoi(info[6])
			idd,errr := strconv.Atoi(info[0])
			if er != nil {
			log.Fatal(er)
			}
			if errr != nil {
			log.Fatal(errr)
			}
			classemt := newClassement(idd, info[1], info[2], info[3],info[4], info[5], temps)
			base.addCompetiteur(classemt)
			}
			firstCall = false
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// Verification de l'unicité
		//fmt.Println("Unicité:")
		//base.uniqueness()
		//fmt.Println("\r\n")
		
	}
	/*
	* 		newBdd:
	* Paramètres:
	*	- cheminBdd:  Chemin et nom de la base de données à ouvrir.
	*
	* Description: 		
	*		Méthode permettant d'ouvrir une connection vers une base de données.
	*/
	
	
	func newBddResult(cheminBdd string)(*BddResult){
		base := new(BddResult)
		base.cheminbdd = cheminBdd
		base.chemincsv = ""
		
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
		base.resultat = new(sql.Rows)
		
		return base
	}
	/*
	* 		col_id2name:
	* Paramètres:
	*	- col_num:  Numéro de la colonne sur laquelle effectuer la modification (ex: 2 => prénom).
	*	- value:	Nouvelle valeur à entrée pour la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant à partir d'un numéro de colonne, de retourner le nom de la colonne.
	*		De plus, la valeur entrée ("value") est retournée au format adéquat pour une requête SQL
	*		(Ex: "VariableString" => "'VariableString'")
	*/
	
	
	func col_id2name2(col_num int, value string)(string, string){
		var col_idr string
		
		switch col_num{
		    case 1:
				col_idr = "id"
				value = fmt.Sprint("'",value,"'")
			case 2:
				col_idr = "prenom"
				value = fmt.Sprint("'",value,"'")
			case 3:
				col_idr = "nom"
				value = fmt.Sprint("'",value,"'")
			case 4:
				col_idr = "sexe"
				value = fmt.Sprint("'",value,"'")
			case 5:
				col_idr = "equipe"
				value = fmt.Sprint("'",value,"'")
			case 6:
				col_idr = "epreuve"
				value = fmt.Sprint("'",value,"'")
			case 7:
				col_idr = "resultat"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return col_idr, value
	}
	
	
	/*
	*
	*
	*
	*
	*/
	func (base BddResult) exportClassement(value string, sexe string){
		var id_col string 
		id_col, value = col_id2name2(6, value)
		
		var id_col2 string 
		id_col2, sexe = col_id2name2(4, sexe)
		fmt.Println(value," - Classement ",sexe)
		
	base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE ", id_col, " = ", value," AND ", id_col2, " = ", sexe," ORDER BY resultat"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
	var info [7]string
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
		}
		fmt.Println(info[0],";",info[1],";", info[2],";", info[3],";", info[4],";", info[5],";", info[6])
		}
	
	}
	
	
	