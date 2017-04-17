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
	
	type Fichier struct{
		cheminFichier string
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
		strconv.Itoa(comp.id) + "','" +
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
		
		file, err := os.Create(fmt.Sprint(cheminFichier,".csv"))
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
			var verif0 bool = false
			var verif1 bool = false
			var verif2 bool = false
			var verif3 bool = false
			var verif4 bool = false
			var verif5 bool = false
			var verif6 bool = false
			var verif7 bool = false
			var verif8 bool = false
			var verif9 bool = false
			info := strings.Split(scanner.Text(), ";")
			if !firstCall{
			num_comp = num_comp + 1
			temps1,errr := strconv.Atoi(info[7])
			temps2,er := strconv.Atoi(info[9])
			if er != nil {
			log.Fatal(er)
			}
			if errr != nil {
			log.Fatal(errr)
			}
			for n := 0; n < 10; n++{
			
				switch(n){
				case 0 : match, _ := regexp.MatchString("([:digit:]*)", info[0] )
					if(match){
					verif0 =true
					}
				case 1 : match, _ := regexp.MatchString("([:alpha:]*)([:digit:]{0})", info[1] )
					if(match){
					verif1 =true
					}
				case 2:  
				match, _ := regexp.MatchString("([:alpha:]*)([:digit:]{0})", info[2] )
					if(match){
					verif2 =true
					}
				case 3 : 
				match, _ := regexp.MatchString("([F|H]{1})", info[3] )
				 if(match){
				 verif3 = true
				}
				case 4 : 
				match, _ := regexp.MatchString("([:digit:]*)+([:alpha:]*)", info[4] )
					if(match){
					verif4 =true
					}
				case 5 : 
				match, _ := regexp.MatchString("([:alpha:]*)", info[5] )
					 if(match){
					verif5 =true
					}
				case 6 : 
					if(info[6]=="Statique" || info[6]=="Speed 100" || info[6]=="DWF" || info[6]=="DNF" || info[6]=="16*50"){
					verif6 =true
					}
				case 7 : 
				match, _ := regexp.MatchString("([[:digit:]]{1,5})", info[7] )
					if(match){
					verif7 =true
					}
				case 8 : 
					if(info[8]=="Statique" || info[8]=="Speed 100" || info[8]=="DWF" || info[8]=="DNF" || info[8]=="16*50"){
					verif8 =true
					}
				case 9 : 
				match, _ := regexp.MatchString("([[:digit:]]{1,4})", info[9] )
					if(match){
					verif9 =true
					}
				}
			}
			if (verif0 && verif1 && verif2 && verif3 && verif4 && verif5 && verif6 && verif7 && verif8 && verif9){
			comp := newCompetiteur(num_comp, info[1], info[2], info[3], info[4], info[5], info[6], temps1, info[8],temps2)
			base.addCompetiteur(comp)
			}else{
			fmt.Println("Echec lors de l'ajout de "+ info[1] + " " + info[2])
			}
			}
			firstCall = false
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// Verification de l'unicité
		fmt.Println("Unicité:")
		base.uniqueness()
		fmt.Println("\r\n")
		
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
	func (base Bdd) count_epreuve_comp(col_num int, value string)(bool,bool,bool,bool,bool,int,int,int,int,int){
	var id_col string
		id_col, value = col_id2name(col_num, value)
		
	
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs WHERE ", id_col, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [10]string
		var n int = 0
		var comp1 bool = true
		var comp2 bool = true
		var comp3 bool = true
		var comp4 bool = true
		var comp5 bool = true
		var nbSTA int =0
		var nbDWF int =0
		var nbSPE int =0
		var nbDNF int =0
		var nbSFC int =0
		for base.resultat.Next() {
			n=n+1
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
			switch (info[6]){
			case "Statique": 
			nbSTA= nbSTA+1
			if (info[8]=="DWF" && info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "DWF":
			nbDWF= nbDWF+1
			if (info[8]=="Speed 100" || info[8]=="Statique" && info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "Speed 100":
			nbSPE= nbSPE+1
			if (info[8]=="DNF" || info[8]=="DWF"&& info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "DNF":
			nbDNF= nbDNF+1
			if (info[8]=="16*50" || info[8]=="Speed 100"&& info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "16*50":
			nbSFC= nbSFC+1
			if (info[8]=="DNF"&& info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}		
			}
			switch(info[8]){
			case "Statique": 
			nbSTA= nbSTA+1
			case "DWF":
			nbDWF= nbDWF+1
			case "Speed 100":
			nbSPE= nbSPE+1
			case "DNF":
			nbDNF= nbDNF+1
			case "16*50":
			nbSFC= nbSFC+1
			}
		}
		return comp1, comp2, comp3, comp4, comp5,nbSTA,nbDWF,nbSPE,nbDNF,nbSFC
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
		
		// CREATION FICHIER
		file, err := os.Create("fichierVerification.txt")
		if err != nil {
				fmt.Println("Erreur lors de la création du fichier planning:\n")
				log.Fatal(err)
			}
		file.WriteString("FICHIER VERIFICATION : il permet de visulaliser les erreurs liées à lacomposition des équipes !\r\n")
				
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT DISTINCT equipe FROM competiteurs "))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()	
	
		var info [1]string
		var nb_sexeH string ="0"
		var nb_sexeF string ="0"
		
		for base.resultat.Next() {
		var ep1 bool
		var ep2 bool
		var ep3 bool
		var ep4 bool
		var ep5 bool
		var nbSTA int =0
		var nbDWF int =0
		var nbSPE int =0
		var nbDNF int =0
		var nbSFC int =0
		var res string
		var res2 string
			base.err = base.resultat.Scan(&info[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
			var nb_comp string = base.count_comp(6,info[0])
			nb_sexeH,nb_sexeF=base.count_sexe_comp(6,info[0])
			ep1,ep2,ep3,ep4,ep5,nbSTA,nbDWF,nbSPE,nbDNF,nbSFC=base.count_epreuve_comp(6,info[0])
			
			if(ep1 && ep2 && ep3 && ep4 && ep5){
			res = "Temps repos : OK"
			}else{
			res = "Erreur temps de repos"}
			
			if(nbSTA==2 && nbDNF==2 && nbDWF==2 && nbSFC==2 && nbSPE==2){
			res2="Nombres épreuves corrects"
			}else{
			res2="Erreur sur nombres épreuves"}
			
			file.WriteString(info[0] + "|" + nb_comp + "|" + "Homme : "+ nb_sexeH + "|" + "Femme : "+ nb_sexeF+ "|" + res +"|"+ res2 +"\r\n" )
			
			if (nb_comp!="5"){
				file.WriteString("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nb_comp + " compétiteurs !\r\n")
			}
			
			if (nb_sexeH != "3"){
				file.WriteString("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !\r\n")
			}
			
			if (nb_sexeF != "2"){
				file.WriteString("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !\r\n")
			}
			
		
		}
		
	}
	
		
	/*
	* 		Bdd.uniqueness:
	* 
	* Description: Méthode permettent de vérifier l'unicité des champs id et licence censé être unique	
	*		
	*/
	func (base Bdd) uniqueness(){	
	
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		//defer base.resultat.Close()
		
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
		if num == 1{
		id_col, value = col_id2name(1, value)
		}else{id_col, value = col_id2name(5, value)}
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", id_col, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête (2)")
			log.Fatal(base.err)}
			
			var inf [1]string
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&inf[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats (2): ")
				log.Fatal(base.err)
			}
			if inf[0]!="1" && num == 1{
			fmt.Println("Erreur doublons sur "+value )
			} else if inf[0]!="1" && num == 2 {
				fmt.Println("Erreur doublons sur "+value )
			}
		}	
	}
	