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
	)
	
	
	type Bdd struct
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
	
	
	
	func (base Bdd) connection() (){
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}
	
	
	/*
	* 		Bdd.disp_comp:
	* Description: 	
	* 		Méthode permettant d'afficher l'integralité des
	* 		competiteurs de la base de données
	*/
	
	func (base Bdd) disp_comp(){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM competiteurs")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [10]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6]+ "|" + info[7]+ "|" + info[8]+ "|" + info[9])
		}
	}
	
	/*
	* 		Bdd.search_comp:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de rechercher un compétiteur en
	* 		competiteurs de la base de données
	*/
	
	func (base Bdd) search_comp(col_num int, value string){
		
		var id_col string
		id_col, value = col_id2name(col_num, value)
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs WHERE ", id_col, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [10]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6]+ "|" + info[7]+ "|" + info[8]+ "|" + info[9])
		}
	}
	
	/*
	* 		Bdd.delComp:
	* Paramètres:
	*	- comp: 	Les informations du compétiteur à ajouter sous la
	*				forme d'une structure de type "competiteur"
	*
	* Description: 		
	*		Méthode permettant d'ajouter un compétiteur dans la 
	* 		base de données
	*/

	func (base Bdd) addComp(comp *competiteur){
		
		_, base.err = base.db.Exec("INSERT INTO competiteurs VALUES('"+
		comp.id + "','" +
		comp.prenom + "','" +
		comp.nom + "','" +
		comp.sexe + "','" +
		comp.num_license + "','" +
		comp.equipe + "','" +
		comp.epreuve1 + "'," +
		strconv.FormatFloat(float64(comp.temps1),'f', -1,  32)+ ",'" +
		comp.epreuve2 + "'," +
		strconv.FormatFloat(float64(comp.temps2) ,'f', -1, 32) + ")")
		
		
		
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout: \n", base.err)
			} else {
			fmt.Println("Ajout validé du compétiteur " + comp.nom +" "+ comp.prenom)
		}
	}
	
	/*
	* 		Bdd.addComp:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de supprimer les compétiteurs en fonction des critères
	*		en entrée.
	*/

	func (base Bdd) delComp(col_num int, value string){
		var id_col string
		id_col, value = col_id2name(col_num, value)
	
		_, base.err = base.db.Exec("DELETE FROM competiteurs WHERE " + id_col + " = " + value)
		if base.err != nil {
			fmt.Println("Echec lors de la suppression: \n", base.err)
			} else {
			fmt.Println("Suppression des competiteurs avec " + id_col + " = " + value)
		}
	}
	
	
	/*
	* 		Bdd.reset:
	* Description: 		
	*		Méthode permettant de supprimer tous les compétiteurs contenus dans la base de
	*		données.
	*/
	
	func (base Bdd) reset(){
		_, base.err = base.db.Exec("DELETE FROM competiteurs")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la base: \n", base.err)
			} else {
			fmt.Println("Remise à zéro de la base de données effectuée")
		}
	}
	
		
	/*
	* 		Bdd.export_comp:
	* Paramètres:
	*	- cheminFichier: 	Chemin du fichier à exporter.
	*	- nomFichier:	Nom du fichier à exporter (sans ".CSV")
	*
	* Description: 		
	*		Méthode permettant d'exporter un fichier CSV contenant tous les
	*		compétiteurs de la base de données.
	*/
	func (base Bdd) export_comp(cheminFichier string, nomFichier string){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM competiteurs")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
		
		file, err := os.Create(fmt.Sprint(cheminFichier,nomFichier,".csv"))
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier")
				log.Fatal(err)
			}
		
			var info [10]string
			
			for base.resultat.Next() {
				base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
			}
		file.WriteString(fmt.Sprint(info[0],";", info[1],";", info[2],";", info[3],";", info[4],";", info[5],";", info[6],";", info[7],";", info[8],";", info[9],"\r\n"))
		}
	}
	
	
	/*
	* 		Bdd.import_comp:
	* Paramètres:
	*	- chemin: 	Chemin du fichier à importer avec le nom du fichier et l'extension.
	*
	* Description: 		
	*		Méthode permettant d'importer les compétiteurs contenu dans un fichier CSV
	*/
	
	func (base Bdd) import_comp(chemin string){
		file, err := os.Open(chemin)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		//var info []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ";")
			
			temps1,_ := strconv.ParseFloat(info[7], 64)
			temps2,_ := strconv.ParseFloat(info[9], 64)
			comp := newcomp2(info[0], info[1], info[2], info[3], info[4], info[5], info[6], float32(temps1), info[8],float32(temps2))
			base.addComp(comp)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	
	/*
	* 		Bdd.modif_comp:
	* Paramètres:
	*	- id_comp:	id du compétiteur à modifier
	*	- col_num:  Numéro de la colonne sur laquelle effectuer la modification (ex: 2 => prénom).
	*	- newvalue:	Nouvelle valeur à entrée pour la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de modifier une valeur d'un compétiteur de la base de données.
	*/
	
	func (base Bdd) modif_comp (id_comp string, col_num int, newvalue string){
		
		col_id, value := col_id2name(col_num, newvalue)
		id_comp = fmt.Sprint("'",id_comp,"'")
		
		_, base.err = base.db.Exec("UPDATE competiteurs SET " + col_id + " = " + value + " WHERE id = " + id_comp)
		
		if base.err != nil {
			fmt.Println("Echec lors de la modification: \n", base.err)
		} else {
			fmt.Println("Modification du competiteur " + id_comp + " avec " + col_id + " = " + value)
		}
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
	
	
	func col_id2name(col_num int, value string)(string, string){
		var col_id string
		
		switch col_num{
		    case 1:
				col_id = "id"
				value = fmt.Sprint("'",value,"'")
			case 2:
				col_id = "prenom"
				value = fmt.Sprint("'",value,"'")
			case 3:
				col_id = "nom"
				value = fmt.Sprint("'",value,"'")
			case 4:
				col_id = "sexe"
				value = fmt.Sprint("'",value,"'")
			case 5:
				col_id = "num_license"
				value = fmt.Sprint("'",value,"'")
			case 6:
				col_id = "equipe"
				value = fmt.Sprint("'",value,"'")
			case 7:
				col_id = "epreuve1"
				value = fmt.Sprint("'",value,"'")
			case 8:
				col_id = "temps1"
				value = fmt.Sprint("'",value,"'")
			case 9:
				col_id = "epreuve2"
				value = fmt.Sprint("'",value,"'")
			case 10:
				col_id = "temps2"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return col_id, value
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
		base.chemincsv = ""
		
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
		base.resultat = new(sql.Rows)
		
		return base
	}
	