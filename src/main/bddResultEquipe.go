package main
	
	import (
	"strconv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
	)
		
	/*
	* 		Bdd.exportClassementEquipe:
	* Description: 		
	*		Méthode permettant d'exporter le classement équipe
	*/
	func (base Bdd) exportClassementEquipe(){
		//---variables---
		var info [5]string
		var numPlace int =0
		var tabPlace []int
		var tab []string
		//---------------
		
		//Importation du nom des équipes et du nombre de points dans la bdd classementEquipe
		//base.importEquipe()
		base.importPoint()
		
		//Génération du fichier export du classement équipe dans "export" et "export/archives"
		t := time.Now()
			date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())
		
		file, err := os.Create(fmt.Sprint("export/archives/",date,"-Equipe.csv"))
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export/archives\" dans le dossier de l'application?")
				log.Fatal(err)
			}
		file2, err := os.Create(fmt.Sprint("export/Classement-Equipe.csv"))
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
			log.Fatal(err)
		}
		
		//Requête sql pour ordonner de manière croissante les équipes en fonction de leur points
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classementequipe ORDER BY point ASC"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		
		//Ecriture dans les fichiers de la première ligne d'en-tête
		file.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Equipe; Point; Place; Etat\r\n"))
		file2.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Equipe; Point; Place; Etat\r\n"))	
		
		// Parcours du resultat de la requête
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4])
			if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
				}				
			numPlace=numPlace+1
			tabPlace=append(tabPlace,numPlace)
			tab=append(tab,info[1])
			
			//Ecriture dans les fichiers
			file.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", numPlace,";", info[4],"\r\n"))
			file2.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", numPlace,";", info[4],"\r\n"))
		}
		
		for n:=0;n<len(tabPlace);n++{
		   	base.modifPlace(tab[n],strconv.Itoa(tabPlace[n]))
		}		
	}
	
	/*
	* 		col_id2name3:
	* Paramètres:
	*	- col_num:  Numéro de la colonne sur laquelle effectuer la modification (ex: 2 => équipe).
	*	- value:	Nouvelle valeur à entrée pour la colonne choisie.
	*
	* Description: 		
	*		Méthode permettant à partir d'un numéro de colonne, de retourner le nom de la colonne.
	*		De plus, la valeur entrée ("value") est retournée au format adéquat pour une requête SQL
	*		(Ex: "VariableString" => "'VariableString'")
	*/
	func col_id2name3(col_num int, value string)(string, string){
		//---variables---
		var col_idr string
		//---------------
		
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
			case 4:
				col_idr = "place"
				value = fmt.Sprint("'",value,"'")
			case 5:
				col_idr = "etat"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return col_idr, value
	}
	
	/*
	* 		Bdd.importEquipe:
	* Description: 		
	*		Méthode permettant d'importer le nom des équipes dans la base de données adéquate
	*/
	func (base Bdd) importEquipe(){
		//---variables---
		var info [10]string
		var tab []string
		var equipe string
		//---------------
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs GROUP BY equipe"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
					}
				equipe = info[5]
				if equipe != "SOLO"{
					tab=append(tab,equipe)
				}
		}
		
		for n:=0;n<len(tab);n++{
			classemt := newClassementE(0, tab[n],0 ,0,false)
			base.addEquipe(classemt)
		}
	}
	
	/*
	* 		Bdd.importPoint:
	* Description: 		
	*		Méthode permettant d'importer le nombre de points de chaque équipe dans la base de donnée classement équipe
	*/
	func (base Bdd) importPoint(){
		//---variables---
		var info [5]string
		var tab []string
		var equipe string
		//--------------
		
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classementequipe GROUP BY equipe"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
					}
				equipe = info[1]
			tab=append(tab,equipe)
		}
		
		for n:=0;n<len(tab);n++{
			var id_col string
			var equipe2 string
			var info [13]string
			var tabpoint []string
			var point string
			var nbpoint int
			var nb int
			var nbCompetiteur int =0
			
			//On "clear" l'ancien tableau:
			tabpoint = tabpoint[:0]
			
			id_col, equipe2 = col_id2name2(5, tab[n])
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE " + id_col + " = " + equipe2))
			if base.err != nil {
				fmt.Println("Erreur lors de l'execution de la requête 1")
			}
			defer base.resultat.Close()
		
			for base.resultat.Next() {
				base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3],&info[4], &info[5], &info[6], &info[7],&info[8], &info[9], &info[10], &info[11], &info[12])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: \n")
					log.Fatal(base.err)
				}
				point = info[10]	
				nbCompetiteur = nbCompetiteur + 1
				tabpoint=append(tabpoint,point)
			}
			
			//Vérification prise en comte de 10 résultats correspond à 2 épreuves par compétiteur
			if nbCompetiteur != 10{
				fmt.Println("Classement équipe non complet !")
				fmt.Println("Equipe : ", equipe2," nb compétiteur pris en compte : ",nbCompetiteur)
			}
			
			
			for n:=0;n<len(tabpoint);n++{
				nb,_=strconv.Atoi(tabpoint[n])
				nbpoint=nbpoint+nb
			}	
			base.modifPoint(tab[n],strconv.Itoa(nbpoint))
		}
	}
	
	/*
	* 		Bdd.modifPoint:
	* Description: 		
	*		Méthode permettant de modifier les points d'une équipe dans la Bdd
	* Paramètres:
	*	- equipe : nom de l'équipe dont on veut modifier le nombre de point
	* 	- newvalue : nouvelle valeur de point à insérer dans la Bdd
	*/
	func (base Bdd) modifPoint(equipe string, newvalue string){
		col_id2, equipe := col_id2name3(2, equipe)
		col_id, value := col_id2name3(3, newvalue)
		
		//on doit utiliser col_id2
		if(col_id2=="0"){
		}
		
		_, base.err = base.db.Exec("UPDATE classementequipe SET "  + col_id + " = " + value +  " WHERE equipe = " + equipe)
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : ", base.err)
		}else{
			//fmt.Println("Modification du competiteur " + strconv.Itoa(id_comp) + " avec " + col_id + " = " + value)
		}
	}
	
	/*
	* 		Bdd.modifPlace:
	* Description: 		
	*		Méthode permettant de modifier la place d'une équipe dans la Bdd
	* Paramètres:
	*	- equipe : nom de l'équipe dont on veut modifier la place 
	* 	- newvalue : nouvelle valeur de point à insérer dans la Bdd
	*/
	func (base Bdd) modifPlace(equipe string, newvalue string){
		col_id2, equipe := col_id2name3(2, equipe)
		col_id, value := col_id2name3(4, newvalue)
		if(col_id2=="0"){
		//on doit utiliser col_id2
		}
		
		_, base.err = base.db.Exec("UPDATE classementequipe SET "  + col_id + " = " + value +  " WHERE equipe = " + equipe)
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : ", base.err)
		}else {
			//fmt.Println("Modification du competiteur " + strconv.Itoa(id_comp) + " avec " + col_id + " = " + value)
		}
	}
	/*
	* 		Bdd.addEquipe:
	* Description: 		
	*		Méthode permettant d'ajouter une équipe dans la Bdd
	* Paramètres:
	*	- ClassementEquipe : structure ClassementEquipe 
	*/
	func (base Bdd) addEquipe(board *ClassementEquipe){
		_, base.err = base.db.Exec("INSERT INTO classementequipe (equipe, point, place, etat) VALUES('" +
		board.equipe + "'," +
		strconv.Itoa(board.point) + "," +
		strconv.Itoa(board.place) + ",'" +
		strconv.FormatBool(board.etat) +"')")
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : "+ board.equipe, base.err)
			} else {
			fmt.Println("Ajout validé de l'equipe " + board.equipe)
		}
	}
	/*
	* 		Bdd.addEquipe:
	* Description: 		
	*		Méthode permettant d'afficher la Bdd classementequipe
	*/
	func (base Bdd) displayEquipe(){
		
		var info [5]string

		base.resultat, base.err = base.db.Query("SELECT * FROM classementequipe")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3]+ "|" + info[4])
		}
	}