package main
	
	/*import (
	"strconv"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"bufio"
	"strings"
	//"time"
	)*/
		
	/*type BddResultE struct
	{
		cheminbdd string
		chemincsv string
		resultat *sql.Rows
		db *sql.DB
		err error
	}*/
	
	/*
	* 		Bdd.connection:
	* Description: 
	* 		Méthode permettant de se connecter à la base de 
	* 		données située au chemin contenu dans l'attribut "cheminBdd"
	*/	
	
	/*func (base BddResultE) connection() (){
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}*/
	
	/*
	* 		Bdd.reset:
	* Description: 		
	*		Méthode permettant de supprimer tous les compétiteurs contenus dans la base de
	*		données.
	*/
	
	/*func (base BddResultE) reset(){
		_, base.err = base.db.Exec("DELETE FROM classementequipe")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
		} else {
			_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='classementequipe'")
			if base.err != nil {
				fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
				} else {
				fmt.Println("Remise à zéro de la base de données effectuée")
			
			}
		}
	}*/
	/*
	* 		Bdd.disp_comp:
	* Description: 	
	* 		Méthode permettant d'afficher l'integralité des
	* 		competiteurs de la base de données
	*/
	
	/*func (base BddResultE) displayCompetiteur(){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM classementequipe")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [3]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2])
		}
	}*/
	
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
	
	/*func (base BddResultE) searchCompetiteur(col_num int, value string){
		
		var id_col string
		var searchValue string
		
		searchValue = fmt.Sprint("'%",value,"%'")
		id_col, value = col_id2name(col_num, value)
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classementequipe WHERE ", id_col, " LIKE ", searchValue))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [3]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2])
		}
	}*/
	
	
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

/*	func (base BddResultE) addCompetiteur(boardE *ClassementEquipe){
		
		_, base.err = base.db.Exec("INSERT INTO classementequipe ( equipe, point) VALUES('" +
		boardE.equipe + "'," +
		strconv.Itoa(boardE.point) + ")")
		
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : "+ boardE.equipe, base.err)
			} else {
			fmt.Println("Ajout validé du resulat equipe " + boardE.equipe)
		}
	}*/
	
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

	/*func (base BddResultE) deleteCompetiteur(col_num int, value string){
		var id_col string
		value = fmt.Sprint("'",value,"'")
		
		if col_num==1 {
			id_col = "id"		
		} else if col_num==2{
			id_col = "equipe"		
		}
		
		if !(col_num < 1 && col_num > 2){
			_, base.err = base.db.Exec("DELETE FROM classementequipe WHERE " + id_col + " = " + value)
			if base.err != nil {
				fmt.Println("Echec lors de la suppression: \n", base.err)
				} else {
				fmt.Println("Suppression des competiteurs avec " + id_col + " = " + value)
			}
		} else {
			err := "Le numéro entré est invalide!"
			fmt.Println(err);
		}
	}*/
	
	/*
	* 		newBdd:
	* Paramètres:
	*	- cheminBdd:  Chemin et nom de la base de données à ouvrir.
	*
	* Description: 		
	*		Méthode permettant d'ouvrir une connection vers une base de données.
	*/
	/*func newBddResultE(cheminBdd string)(*BddResultE){
		base := new(BddResultE)
		base.cheminbdd = cheminBdd
		base.chemincsv = ""
		
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
		base.resultat = new(sql.Rows)
		
		return base
	}*/
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
	
	
	/*func col_id2name3(col_num int, value string)(string, string){
		var col_idr string
		
		switch col_num{
		    case 1:
				col_idr = "id"
				value = fmt.Sprint("'",value,"'")
			case 2:
				col_idr = "equipe"
				value = fmt.Sprint("'",value,"'")
			case 3:
				col_idr = "point"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return col_idr, value
	}*/
	
	// Enregistrer les épreuves dans le tableau
/*func getConfigurationEpreuve1()(*ConfigurationEpreuve){
	file, err := os.Open("config/ConfigurationEpreuve.csv")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier \"ConfigurationEpreuve\": " )
		log.Fatal(err)
	}
	defer file.Close()
	
	
	var firstCall bool
	firstCall = true
	var nextconfig *ConfigurationEpreuve
	
	scanner := bufio.NewScanner(file)
	//On clear l'ancien tableau:
	//p.cfgEpreuves = p.cfgEpreuves[:0]
	
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ";")
		if !firstCall{
		ordre, _ := strconv.Atoi(info[0])
		seuilMin, _ := strconv.Atoi(info[2])
		seuilMax, _ := strconv.Atoi(info[3])
		nbParPassage, _ := strconv.Atoi(info[4])
		dureeEchauffement, _ := strconv.Atoi(info[5])
		dureePassage, _ := strconv.Atoi(info[6])
		dureeAppel, _ := strconv.Atoi(info[7])
		surveillance, _ := strconv.Atoi(info[8])
		battementSerie, _ := strconv.Atoi(info[9])
		battementEpreuve, _ := strconv.Atoi(info[10])
	
		nextconfig = newConfigurationEpreuve(ordre, info[1], seuilMin, seuilMax, nbParPassage, dureeEchauffement, dureePassage, dureeAppel, surveillance,
												battementSerie,battementEpreuve, info[11])
		}
		firstCall = false
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nextconfig
}*/
	