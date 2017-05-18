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
	"time"
	"regexp"
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
	* 		données située au chemin contenu dans l'attribut "cheminBdd"
	*/
	
	
	
	func (base Bdd) connection() (){
		base.db, base.err = sql.Open("sqlite3", base.cheminbdd)
		if base.err != nil {
			log.Fatal("Erreur de connection à la base de données:\n", base.err)
		}
	}
	
	
	/*
	* 		Bdd.displayCompetiteur:
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
		var test bool 
		test = comp.check()
	
		if (test) {
		_, base.err = base.db.Exec("INSERT INTO competiteurs (prenom, nom, sexe, num_license, equipe, epreuve1, annonce1, epreuve2, annonce2) VALUES('" +
		comp.prenom + "','" +
		comp.nom + "','" +
		comp.sexe + "','" +
		comp.num_license + "','" +
		comp.equipe + "','" +
		comp.epreuve1 + "'," +
		strconv.Itoa(comp.annonce1) + ",'" +
		comp.epreuve2 + "'," +
		strconv.Itoa(comp.annonce2) + ")")
		} else {
			log.Fatal(fmt.Sprint("Erreur lors de l'ajout du compétiteur ",comp.prenom," ",comp.nom,". données entrées eronnées."))
		}
		
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout: \n", base.err)
			} else {
			fmt.Println("Ajout validé du compétiteur " + comp.nom +" "+ comp.prenom)
		}
	}
	
	/*
	* 		Bdd.deleteCompetiteur:
	* Paramètres:
	*	- col_num: 	numéro de la colonne sur laquelle effectuer la recherche (1 => id/ 2 => équipe).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant de supprimer les compétiteurs en fonction des critères
	*		en entrée.
	*/

	func (base Bdd) deleteCompetiteur(col_num int, value string){
		var id_col string
		value = fmt.Sprint("'",value,"'")
		
		if col_num==1 {
			id_col = "id"		
		} else if col_num==2{
			id_col = "equipe"		
		}
		
		if !(col_num < 1 && col_num > 2){
			_, base.err = base.db.Exec("DELETE FROM competiteurs WHERE " + id_col + " = " + value)
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
	* 		Bdd.resetCompetiteurs:
	* Description: 		
	*		Méthode permettant de supprimer tous les compétiteurs contenus dans la base de
	*		données.
	*/
	
	func (base Bdd) resetCompetiteurs(){
		_, base.err = base.db.Exec("DELETE FROM competiteurs")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la table competiteurs. \n", base.err)
		} else {
			_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='competiteurs'")
			if base.err != nil {
				fmt.Println("Echec lors de la remise à 0 de la table competiteurs: \n", base.err)
				} 
			}
		}
	
		
	/*
	* 		Bdd.exportCompetiteur:
	* Description: 		
	*		Méthode permettant d'exporter un fichier CSV contenant tous les
	*		compétiteurs de la base de données.
	*/
	func (base Bdd) exportCompetiteur(){
	
		base.resultat, base.err = base.db.Query("SELECT * FROM competiteurs")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
		}
		defer base.resultat.Close()
		
			t := time.Now()
			date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())
		
		file, err := os.Create(fmt.Sprint("export/archives/",date,"-competiteurs.csv"))
		file2, err := os.Create(fmt.Sprint("export/competiteurs.csv"))
		
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
				log.Fatal(err)
			}
		
			var info [10]string
			
			file2.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Prenom; Nom; Sexe; Num_License; Equipe; Epreuve1; annonce1; Epreuve2; annonce2\r\n"))
			file.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Prenom; Nom; Sexe; Num_License; Equipe; Epreuve1; annonce1; Epreuve2; annonce2\r\n"))
			
			for base.resultat.Next() {
				base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
			}
		file.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", info[3],";", info[4],";", info[5],";", info[6],";", info[7],";", info[8],";", info[9],"\r\n"))
		file2.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", info[3],";", info[4],";", info[5],";", info[6],";", info[7],";", info[8],";", info[9],"\r\n"))
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
	
	func (base Bdd) importCompetiteur(){
		file, err := os.Open("import/competiteurs.csv")
		if err != nil {
			println("Impossible d'ouvrir le fichier \"competiteurs.csv\" dans le dossier import")
			log.Fatal(err)
		}
		defer file.Close()
		var firstCall bool
		
		firstCall = true
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ";")
			if !firstCall{
				temps1,errr := strconv.Atoi(info[6])
				temps2,er := strconv.Atoi(info[8])
				if er != nil {
				log.Fatal(er)
				}
				if errr != nil {
				log.Fatal(errr)
				}
				comp := newCompetiteur(0, info[0], info[1], info[2], info[3], info[4], info[5], temps1, info[7],temps2)
				base.addCompetiteur(comp)
			}
			firstCall = false
		}
			

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// Verification de l'unicité
		base.uniqueness()	
		//Importation equipe dans la Bdd 
		base.importEquipe()

		
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
	
	func (base Bdd) modifCompetiteur (id_comp int, col_num int, newvalue string){
		var test = true
		col_id, value := col_id2name(col_num, newvalue)
		
		test = verifValue(col_num, newvalue)
		if (test){
			_, base.err = base.db.Exec("UPDATE competiteurs SET " + col_id + " = " + value + " WHERE id = " + strconv.Itoa(id_comp))
			
			if base.err != nil {
				fmt.Println("Echec lors de la modification: \n", base.err)
			} else {
				fmt.Println("Modification du competiteur " + strconv.Itoa(id_comp) + " avec " + col_id + " = " + value)
			}
		} else {
			fmt.Println("Erreur lors de la modifications du compétiteur!")
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
	func verifValue(col_num int, value string)(bool){
		var verif = true
		verif = true
		switch col_num{
		    case 2, 3:
				match, _ := regexp.MatchString("^[\\p{L}- ]*$", value )
				if(!match){
					verif =false
					fmt.Println("Erreur! Format du prénom.")
				}
			case 4:
				match, _ := regexp.MatchString("([F|H])", value )
				if(!match || len(value) > 1){
					verif =false
					fmt.Println("Erreur! Format du sexe.")
				}
			case 5:
				match, _ := regexp.MatchString("^[A-Za-z0-9-]*$", value )
				if(!match){
					verif =false
					fmt.Println("Erreur! Format du numéro de license.")
				}
			case 6:
				match, _ := regexp.MatchString("^[\\p{L}0-9- _]*$", value )
				if(!match){
					verif =false
					fmt.Println("Erreur! Format du nom d'équipe.")
				}
			case 7,9:
				if(value!="sta" && value!="spd" && value!="dwf" && value!="dnf" && value!="1650"){
					verif =false
					fmt.Println("Erreur! Format du epreuve (Rappel des valeurs possibles: sta, spd, dwf, dnf, 1650).")
				}
			case 8,10:
				match, _ := regexp.MatchString("(^[0-9]*$)", value)
				if(!match){
					verif = false
					fmt.Println("Erreur! Format du format de l'annonce.")
				}
			default:
				log.Fatal("Numéro de colone invalide")
			}
		return verif
	}
	
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
				col_id = "annonce1"
				value = fmt.Sprint("'",value,"'")
			case 9:
				col_id = "epreuve2"
				value = fmt.Sprint("'",value,"'")
			case 10:
				col_id = "annonce2"
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
			case "sta": 
			nbSTA= nbSTA+1
			if (info[8]=="dwf" && info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "dwf":
			nbDWF= nbDWF+1
			if (info[8]=="spd" || info[8]=="sta" && info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "spd":
			nbSPE= nbSPE+1
			if (info[8]=="dnf" || info[8]=="dwf"&& info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "dnf":
			nbDNF= nbDNF+1
			if (info[8]=="1650" || info[8]=="spd"&& info[8]!=info[6]){
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "1650":
			nbSFC= nbSFC+1
			if (info[8]=="dnf"&& info[8]!=info[6]){
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
			case "sta": 
			nbSTA= nbSTA+1
			case "dwf":
			nbDWF= nbDWF+1
			case "spd":
			nbSPE= nbSPE+1
			case "dnf":
			nbDNF= nbDNF+1
			case "1650":
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
		t := time.Now()
		date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())
		
		file, err := os.Create(fmt.Sprint("export/archives/",date,"-FichierVerification.txt"))
		file2, err := os.Create(fmt.Sprint("export/FichierVerification.txt"))
		if err != nil {
				fmt.Println("Erreur lors de la création du fichier verification\n")
				log.Fatal(err)
			}
		file.WriteString("FICHIER VERIFICATION : il permet de visulaliser les erreurs liées à la composition des équipes !\r\n")
		file.WriteString("\r\n")
		file2.WriteString("FICHIER VERIFICATION : il permet de visulaliser les erreurs liées à la composition des équipes !\r\n\r\n")
		
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
			file2.WriteString(info[0] + "|" + nb_comp + "|" + "Homme : "+ nb_sexeH + "|" + "Femme : "+ nb_sexeF+ "|" + res +"|"+ res2 +"\r\n" )
			
			if (nb_comp!="5"){
				file.WriteString("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nb_comp + " compétiteurs !\r\n")
				file2.WriteString("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nb_comp + " compétiteurs !\r\n")
			}
			
			if (nb_sexeH != "3"){
				file.WriteString("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !\r\n")
				file2.WriteString("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !\r\n")
			}
			
			if (nb_sexeF != "2"){
				file.WriteString("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !\r\n")
				file2.WriteString("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !\r\n")
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
	