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
	"regexp"
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
	
	func (base Bdd) displayCompetiteur(){
	
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
	* 		Bdd.searchCompetiteur:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de rechercher un compétiteur en
	* 		competiteurs de la base de données
	*/
	
	func (base Bdd) searchCompetiteur(col_num int, value string){
		
		var id_col string
		var searchValue string
		
		searchValue = fmt.Sprint("'%",value,"%'")
		id_col, value = col_id2name(col_num, value)
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs WHERE ", id_col, " LIKE ", searchValue))
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
	* 		Bdd.addCompetiteur:
	* Paramètres:
	*	- comp: 	Les informations du compétiteur à ajouter sous la
	*				forme d'une structure de type "competiteur"
	*
	* Description: 		
	*		Méthode permettant d'ajouter un compétiteur dans la 
	* 		base de données
	*/

	func (base Bdd) addCompetiteur(comp *Competiteur){
		
		_, base.err = base.db.Exec("INSERT INTO competiteurs VALUES('"+
		comp.id + "','" +
		comp.prenom + "','" +
		comp.nom + "','" +
		comp.sexe + "','" +
		comp.num_license + "','" +
		comp.equipe + "','" +
		comp.epreuve1 + "'," +
		strconv.Itoa(comp.temps1) + ",'" +
		comp.epreuve2 + "'," +
		strconv.Itoa(comp.temps2) + ")")
		
		
		
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout: \n", base.err)
			} else {
			fmt.Println("Ajout validé du compétiteur " + comp.nom +" "+ comp.prenom)
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

	func (base Bdd) deleteCompetiteur(col_num int, value string){
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
	* 		Bdd.exportCompetiteur:
	* Paramètres:
	*	- cheminFichier: 	Chemin du fichier à exporter.
	*	- nomFichier:	Nom du fichier à exporter (sans ".CSV")
	*
	* Description: 		
	*		Méthode permettant d'exporter un fichier CSV contenant tous les
	*		compétiteurs de la base de données.
	*/
	func (base Bdd) exportCompetiteur(cheminFichier string){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM competiteurs")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
		
		file, err := os.Create(fmt.Sprint(cheminFichier))
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier")
				log.Fatal(err)
			}
		
			var info [10]string
			
			
			file.WriteString(fmt.Sprint("Identifiant; Prenom; Nom; Sexe; Num_License; Equipe; Epreuve1; temps1; Epreuve2; temps2\r\n"))
			
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
	* 		Bdd.importCompetiteur:
	* Paramètres:
	*	- chemin: 	Chemin du fichier à importer avec le nom du fichier et l'extension.
	*
	* Description: 		
	*		Méthode permettant d'importer les compétiteurs contenu dans un fichier CSV
	*/
	
	func (base Bdd) importCompetiteur(chemin string){
		file, err := os.Open(chemin)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		var firstCall bool
		var num_comp int
		
		firstCall = true
		num_comp = 0;
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			
			info := strings.Split(scanner.Text(), ";")
			if !firstCall{
			num_comp = num_comp + 1
			temps1,_ := strconv.Atoi(info[7])
			temps2,_ := strconv.Atoi(info[9])
			comp := newCompetiteur(fmt.Sprint(info[0],num_comp), info[1], info[2], info[3], info[4], info[5], info[6], temps1, info[8],temps2)
			base.addCompetiteur(comp)
			}
			firstCall = false
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	
	/*
	* 		Bdd.modifCompetiteur:
	* Paramètres:
	*	- id_comp:	id du compétiteur à modifier
	*	- col_num:  Numéro de la colonne sur laquelle effectuer la modification (ex: 2 => prénom).
	*	- newvalue:	Nouvelle valeur à entrée pour la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de modifier une valeur d'un compétiteur de la base de données.
	*/
	
	func (base Bdd) modifCompetiteur (id_comp string, col_num int, newvalue string){
		
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
		/*
	* 		Bdd.orderby_comp:
	* 
	* Description: 		
	*		Méthode permettant de trier les compétiteurs 
	* par équipe
	*/
	
	func (base Bdd) orderby_comp(){
		
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs ORDER BY equipe "))
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
	* 		Bdd.count_comp:
	* 
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe	
	*		
	*/
	func (base Bdd) count_comp(col_num int, value string)(string){
	var id_col string
		id_col, value = col_id2name(col_num, value)
		
	
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", id_col, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		
		}
		return info[0]
	}
	
	/*
	* 		Bdd.count_epreuve_comp:
	* 
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe	
	*		
	*/
	func (base Bdd) count_epreuve_comp(col_num int, value string){
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
			switch (info[6]){
			case "Statique":
			if (info[8]=="DWF"){
			fmt.Println("Erreur temps de repos pour "+ info[0])
			}
			case "DWF":
			if (info[8]=="Speed 100"){
			fmt.Println("Erreur temps de repos pour "+ info[0])
			}
			case "Speed 100":
			if (info[8]=="DNF"){
			fmt.Println("Erreur temps de repos pour "+ info[0])
			}
			case "DNF":
			if (info[8]=="16*50"){
			fmt.Println("Erreur temps de repos pour "+ info[0])
			}
			
			}
		
		}

	}
	
	/*
	* 		Bdd.count_sexe_comp:
	* 
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe	
	*		
	*/
	func (base Bdd) count_sexe_comp(col_num int, value string)(string,string){
	var id_col string
	var id_col2 string
	var col_num2 int = 4
	var valueH string = "H"
	var valueF string = "F"
		id_col, value = col_id2name(col_num, value)
		id_col2, valueH = col_id2name(col_num2, valueH)
		id_col2, valueF = col_id2name(col_num2, valueF)
		
	
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", id_col, " = ", value," AND ", id_col2, " = ", valueH))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var infoH [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&infoH[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		}
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", id_col, " = ", value," AND ", id_col2, " = ", valueF))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var infoF [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&infoF[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		}
		return infoH[0],infoF[0]
	}
	
	
	/*
	* 		Bdd.check_team:
	* 
	* Description: 		
	*		Méthode permettant de vérifier la validité des equipes
	*/
	
	func (base Bdd) check_team(){
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT DISTINCT equipe FROM competiteurs "))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
	
		// Verification de l'unicité
		fmt.Println("Unicité:")
		base.uniqueness()
		fmt.Println("\n")
		
		// Verification de la validité des champs
		fmt.Println("Validité champs:")
		base.valeur()
		fmt.Println("\n")
		
		//fmt.Println(base.resultat)
		var info [1]string
		var nb_sexeH string ="0"
		var nb_sexeF string ="0"
		
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
			var nb_comp string = base.count_comp(6,info[0])
			nb_sexeH,nb_sexeF=base.count_sexe_comp(6,info[0])
			base.count_epreuve_comp(6,info[0])
			
			fmt.Println(info[0] + "|" + nb_comp + "|" + "Homme : "+ nb_sexeH + "|" + "Femme : "+ nb_sexeF+ "|" )
			
			if (nb_comp!="5"){
				fmt.Println("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nb_comp + " compétiteurs !")
			}
			
			if (nb_sexeH != "3"){
				fmt.Println("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !")
			}
			
			if (nb_sexeF != "2"){
				fmt.Println("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !")
			}
		}
		
	}
	
		
	/*
	* 		Bdd.uniqueness:
	* 
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe	
	*		
	*/
	func (base Bdd) uniqueness(){	
	
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs"))
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
			base.verif(info[0],1)
			base.verif(info[4],2)
		}	
	}
		
	func (base Bdd) verif(val string, num int ){
		var id_col string
			var value string = val
		id_col, value = col_id2name(1, value)
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", id_col, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 2")
			log.Fatal(base.err)}
			
			var inf [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&inf[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats 2: \n")
				log.Fatal(base.err)
			}
			if inf[0]!="1" && num == 1{
			fmt.Println("Erreur doublons sur "+value )
			
			} else if inf[0]!="0" && num == 2 {
				fmt.Println("Erreur doublons sur "+value )
			}
		}
			
	}
	
	/*
	* 		Bdd.valeur:
	* 
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe	
	*		
	*/
	func (base Bdd) valeur(){	
	base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs"))
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
			for n := 0; n < 9; n++{
			
			switch(n){
			case 0 : match, _ := regexp.MatchString("([[:alnum:]]{5,7})", info[0] )
			 if(match){
            //([:digit:]{1,2})
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 1 : 
			match, _ := regexp.MatchString("([:alpha:]*)([:digit:]{0})", info[1] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			
			case 2:  
			match, _ := regexp.MatchString("([:alpha:]*)([:digit:]{0})", info[2] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 3 : 
			match, _ := regexp.MatchString("([H|F]?)", info[3] )
			 if(match){
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 4 : 
			match, _ := regexp.MatchString("([:digit:]*)+([:alpha:]*)", info[4] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 5 : 
			match, _ := regexp.MatchString("([:alpha:]*)", info[5] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 6 : 
			 if(info[6]=="Statique" || info[6]=="Speed 100" || info[6]=="DWF" || info[6]=="DNF" || info[6]=="16*50"){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 7 : 
			match, _ := regexp.MatchString("([[:digit:]]{1,5})", info[7] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 8 : 
			if(info[6]=="Statique" || info[6]=="Speed 100" || info[6]=="DWF" || info[6]=="DNF" || info[6]=="16*50"){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			case 9 : 
			match, _ := regexp.MatchString("([[:digit:]]{1,5})", info[9] )
			 if(match){
            
			}else{
			fmt.Println("Erreur sur " + info[n])
			}
			}
			if(info[n]==""){
			fmt.Println("Erreur valeur vide " + info[n])
			}
	        }
		}
	
	}
		