package main

	import (
	"strconv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"bufio"
	"strings"
	"time"
	)

	func (base Bdd) displayClassement(){
		base.resultat, base.err = base.db.Query("SELECT * FROM classement")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()

		var info [13]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9], &info[10], &info[11], &info[12])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6]+ "|" + info[7]+ "|" + info[8]+ "|" + info[9]+ "|" + info[10]+ "|" + info[11]+ "|" + info[12])
		}
	}

	/*
	* 		Bdd.searchCompetiteur:
	* Paramètres:
	*	- colNum: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description:
	*		Méthode permettant de rechercher un compétiteur en
	* 		competiteurs de la base de données
	*/

	func (base Bdd) searchCompetiteurClassement(colNum int, value string){

		var idCol string
		var searchValue string

		searchValue = fmt.Sprint("'%",value,"%'")
		idCol, value = idCol2name(colNum, value)

		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE ", idCol, " LIKE ", searchValue))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()

		var info [13]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9], &info[10], &info[11], &info[12])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6]+ "|" + info[7]+ "|" + info[8]+ "|" + info[9]+ "|" + info[10]+ "|" + info[11]+ "|" + info[12])
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

	func (base Bdd) addCompetiteurClassement(board *Classement){

		_, base.err = base.db.Exec("INSERT INTO classement ( id, prenom, nom, sexe, equipe, epreuve, annonce, resultat, place, rslt, plc, disq, description) VALUES(" +
		strconv.Itoa(board.id) + ",'" +
		board.prenom + "','" +
		board.nom + "','" +
		board.sexe + "','" +
		board.equipe + "','" +
		board.epreuve + "'," +
		strconv.Itoa(board.annonce) + "," +
		strconv.FormatFloat(board.resultat, 'f', -1, 64) + "," +
		strconv.Itoa(board.place) + "," +
		strconv.FormatFloat(board.rslt, 'f', -1, 64) + "," +
		strconv.Itoa(board.plc) + ",'" +
		strconv.FormatBool(board.disq) + "','" +
		board.description + "')")

		if base.err != nil {
			fmt.Println("Echec lors de l'ajout1 : "+ board.nom +" "+ board.prenom, base.err)
			} else {
			fmt.Println("Ajout validé du resulat compétiteur " + board.nom +" "+ board.prenom)
		}
	}

	/*
	* 		Bdd.deleteCompetiteur:
	* Paramètres:
	*	- colNum: 	numéro de la colonne sur laquelle effectuer la recherche (ex: 2 => prénom).
	*	- value:	valeur à rechercher dans la colonne choisie.
	*
	* Description:
	*		Méthode permettant de supprimer les compétiteurs en fonction des critères
	*		en entrée.
	*/

	func (base Bdd) deleteCompetiteurClassement(colNum int, value string){
		var idCol string
		value = fmt.Sprint("'",value,"'")

		if colNum==1 {
			idCol = "id"
		} else if colNum==2{
			idCol = "equipe"
		}

		if !(colNum < 1 && colNum > 2){
			_, base.err = base.db.Exec("DELETE FROM classement WHERE " + idCol + " = " + value)
			if base.err != nil {
				fmt.Println("Echec lors de la suppression: ", base.err)
				} else {
				fmt.Println("Suppression des competiteurs avec " + idCol + " = " + value)
			}
		} else {
			err := "Le numéro entré est invalide!"
			fmt.Println(err);
		}
	}

	func (base Bdd) resetClassement(){
		_, base.err = base.db.Exec("DELETE FROM classement")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la table classement: ", base.err)
		} else{
				fmt.Println("Remise à zéro des tables classement")
		}
	}

	func (base Bdd) resetClassementEquipe(){
		_, base.err = base.db.Exec("DELETE FROM classementequipe")
		if base.err != nil {
			fmt.Println("Echec lors de la remise à 0 de la table classementequipe: ", base.err)
		} else {
			_, base.err = base.db.Exec("DELETE FROM sqlite_sequence WHERE name='classementequipe'")
			if base.err != nil {
				fmt.Println("Echec lors de la remise à 0 de la table classementequipe: ", base.err)
			}
		}
	}

	/*
	* 		Bdd.importResultats:
	* Paramètres:
	*	- chemin: 	Chemin du fichier à importer avec le nom du fichier et l'extension.
	*
	* Description:
	*		Méthode permettant d'importer les compétiteurs contenu dans un fichier CSV
	*/

	func (base Bdd) importResultat(){
		file, err := os.Open("../var/import/classement.csv")
		if err != nil {
			println("Impossible d'ouvrir le fichier \"classement.csv\" dans le dossier import")
			log.Fatal(err)
		}
		defer file.Close()

		var firstCall bool
		firstCall = true
		var res float64
		var place int
		var plc int
		var equipe string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ",")

			if !firstCall{
				temps,er := strconv.ParseFloat(info[6],64)
				idd,errr := strconv.Atoi(info[0])
				annonce := base.recupAnnonce(info[0],info[5])
				disq,_ := strconv.ParseBool(info[7])
				sexe := info[3]
				equipe = info[4]
				if er != nil {
					log.Fatal(er)
				}
				if errr != nil {
					log.Fatal(errr)
				}
				switch(info[5]){
				case "spd": 
					res=calculResultat(sexe,equipe,"spd",annonce,info[6],info[7])
				break
				case "850":
					res=calculResultat(sexe,equipe,"850",annonce,info[6],info[7])
				break
				case "dnf":
					res=calculResultat(sexe,equipe,"dnf",annonce,info[6],info[7])
				break
				case "dwf":
					res=calculResultat(sexe,equipe,"dwf",annonce,info[6],info[7])
				break
				case "sta":
					res=calculResultat(sexe,equipe,"sta",annonce,info[6],info[7])
				break
				}

				classemt := newClassement(idd, info[1], info[2], info[3], info[4], info[5],annonce, temps, place, res, plc, disq, info[8])
				base.addCompetiteurClassement(classemt)
			}
			firstCall = false
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}


	/*
	*
	*
	*
	*/
	func (base Bdd)recupAnnonce(id string, epreuve string)(int){
		var idCol string
		idCol, id = idCol2name2(1, id)
		base.resultat, base.err = base.db.Query("SELECT * FROM competiteurs WHERE " + idCol + " = " + id)
			if base.err != nil {
				fmt.Println("Erreur lors de l'execution de la requête")
				log.Fatal(base.err)
			}
		defer base.resultat.Close()

		var info [10]string
		var resultat int
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
			if (epreuve==info[6]){
				resultat,_ = strconv.Atoi(info[7])
			}else if (epreuve==info[8]){
				resultat,_ = strconv.Atoi(info[9])
			} else{
				resultat = 0
			}
		}
		return resultat
	}

	/*
	* 		idCol2name:
	* Paramètres:
	*	- colNum:  Numéro de la colonne sur laquelle effectuer la modification (ex: 2 => prénom).
	*	- value:	Nouvelle valeur à entrée pour la colonne choisie.
	*
	* Description:
	*		Méthode permettant à partir d'un numéro de colonne, de retourner le nom de la colonne.
	*		De plus, la valeur entrée ("value") est retournée au format adéquat pour une requête SQL
	*		(Ex: "VariableString" => "'VariableString'")
	*/


	func idCol2name2(colNum int, value string)(string, string){
		var idrCol string

		switch colNum{
		    case 1:
				idrCol = "id"
				value = fmt.Sprint("'",value,"'")
			case 2:
				idrCol = "prenom"
				value = fmt.Sprint("'",value,"'")
			case 3:
				idrCol = "nom"
				value = fmt.Sprint("'",value,"'")
			case 4:
				idrCol = "sexe"
				value = fmt.Sprint("'",value,"'")
			case 5:
				idrCol = "equipe"
				value = fmt.Sprint("'",value,"'")
			case 6:
				idrCol = "epreuve"
				value = fmt.Sprint("'",value,"'")
			case 7:
				idrCol = "annonce"
				value = fmt.Sprint("'",value,"'")
			case 8:
				idrCol = "resultat"
				value = fmt.Sprint("'",value,"'")
			case 9:
				idrCol = "place"
				value = fmt.Sprint("'",value,"'")
			case 10:
				idrCol = "rslt"
				value = fmt.Sprint("'",value,"'")
			case 11:
				idrCol = "plc"
				value = fmt.Sprint("'",value,"'")
			case 12:
				idrCol = "disq"
				value = fmt.Sprint("'",value,"'")
			case 13:
				idrCol = "description"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return idrCol, value
	}


	/*
	*
	*
	*
	*/
	func (base Bdd) exportClassement(){
		t := time.Now()
			date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())

		file, err := os.Create(fmt.Sprint("../var/export/archives/",date,"-Classement.csv"))
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
				log.Fatal(err)
			}
		file2, err := os.Create(fmt.Sprint("../var/export/Classement.csv"))
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
			log.Fatal(err)
		}
		file.WriteString(fmt.Sprint("\xEF\xBB\xBFId,Prenom,Nom,Sexe,Equipe,Epreuve,Annonce,Resultat,Place,Resultat pris en compte equipe,Place Equipe,Disqualification,Description\r\n"))
		file2.WriteString(fmt.Sprint("\xEF\xBB\xBFId,Prenom,Nom,Sexe,Equipe,Epreuve,Annonce,Resultat,Place,Resultat pris en compte equipe,Place Equipe,Disqualification,Description\r\n"))

		var value string
		var epreuve string
		for i := 1; i <= 5; i++{
				//calcul de la place equipe
			switch(i){
			case 1:
				base.calculPlace("spd")
				value = "'spd'"
				epreuve = "spd"
			break
			case 2:
				base.calculPlace("850")
				value = "'850'"
				epreuve = "850"
			break
			case 3:
				base.calculPlace("dnf")
				value = "'dnf'"
				epreuve = "dnf"
			break
			case 4:
				base.calculPlace("dwf")
				value = "'dwf'"
				epreuve = "dwf"
			break
			case 5:
				base.calculPlace("sta")
				value = "'sta'"
				epreuve = "sta"
			break
			default:
				log.Fatal("Epreuve invalide")
			break
			}


		if(epreuve=="spd" || epreuve == "850"){
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE epreuve = ", value," ORDER BY sexe ASC, resultat ASC"))
			if base.err != nil {
				fmt.Println("Erreur lors de l'execution de la requête 1")
			}
		}else{
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE epreuve = ", value," ORDER BY sexe ASC, resultat DESC"))
			if base.err != nil {
				fmt.Println("Erreur lors de l'execution de la requête 1")
			}
		}
		defer base.resultat.Close()
	var info [13]string
	var numPlaceF int =1
	var numPlaceH int =1
	var nbF int =0
	var nbH int =0
	var sexe string ="F"
	var tabPlace []*Classement
	var nextResult *Classement
	var lastResult float64=-1
	var tabDisqH []*Classement
	var tabDisqF []*Classement
	var tabHomme []*Classement
	var egalF int =0
	var egalH int =0



		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9], &info[10], &info[11], &info[12])

			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)}



			if info[7]!="0" || info[11]=="false"{
				if(info[3]==sexe){
					info[8]=strconv.Itoa(numPlaceF)
					numPlaceF=numPlaceF+1

					Resultnow,_:=strconv.ParseFloat(info[7],64)
					if lastResult==Resultnow{
						info[8]=strconv.Itoa(numPlaceF-2-egalF)
						egalF =egalF + 1
					}else{
					egalF=0
					}
					lastResult,_=strconv.ParseFloat(info[7],64)
				}else{
					info[8]=strconv.Itoa(numPlaceH)
					numPlaceH=numPlaceH+1

					Resultnow,_:=strconv.ParseFloat(info[7],64)
					if lastResult==Resultnow{
						info[8]=strconv.Itoa(numPlaceH-2-egalH)
						egalH =egalH + 1
					}else{
					egalH=0
					}
					lastResult,_= strconv.ParseFloat(info[7],64)
				}
					id,_:=strconv.Atoi(info[0])
					annonce,_ := strconv.Atoi(info[6])
					resultat,_ := strconv.ParseFloat(info[7], 64)
					place,_ := strconv.Atoi(info[8])
					rslt,_ := strconv.ParseFloat(info[9], 64)
					plc,_ := strconv.Atoi(info[10])
					disq,_ := strconv.ParseBool(info[11])
					nextResult = newClassement(id, info[1], info[2], info[3], info[4], info[5], annonce, resultat, place, rslt, plc, disq, info[12])
					tabPlace=append(tabPlace,nextResult)

					if info[3]=="F" {
						file.WriteString(fmt.Sprint(info[0],",",info[1],",", info[2],",", info[3],",", info[4],",", info[5],",", info[6],",", info[7],",", info[8],",", info[9],",", info[10],",", info[11],",", info[12],"\r\n"))
						file2.WriteString(fmt.Sprint(info[0],",",info[1],",", info[2],",", info[3],",", info[4],",", info[5],",", info[6],",", info[7],",", info[8],",", info[9],",", info[10],",", info[11],",", info[12],"\r\n"))
					}else{
						tabHomme=append(tabHomme,nextResult)
					}
				}else{
					if(info[3]=="F"){
						nbF= nbF + 1
					}else{
						nbH= nbH + 1
					}
					info[8]=strconv.Itoa(100)
					id,_:=strconv.Atoi(info[0])
					annonce,_ := strconv.Atoi(info[6])
					resultat,_ := strconv.ParseFloat(info[7], 64)
					place,_ := strconv.Atoi(info[8])
					rslt,_ := strconv.ParseFloat(info[9], 64)
					plc,_ := strconv.Atoi(info[10])
					disq,_ := strconv.ParseBool(info[11])
					nextResult = newClassement(id, info[1], info[2], info[3], info[4], info[5], annonce, resultat, place, rslt, plc, disq, info[12])

					if(info[3]=="F"){
						tabDisqF=append(tabDisqF,nextResult)
					}else{
						tabDisqH=append(tabDisqH,nextResult)
					}
				}
		}

		for i := 0; i < len(tabDisqF); i++{
		var place string
				place=strconv.Itoa(numPlaceF+nbF-1)
				file.WriteString(fmt.Sprint(tabDisqF[i].id,",",tabDisqF[i].prenom,",", tabDisqF[i].nom,",", tabDisqF[i].sexe,",", tabDisqF[i].equipe,",", tabDisqF[i].epreuve,",", tabDisqF[i].annonce,",", tabDisqF[i].resultat,",", place,",", tabDisqF[i].rslt,",", tabDisqF[i].plc,",", tabDisqF[i].disq,",", tabDisqF[i].description,"\r\n"))
				file2.WriteString(fmt.Sprint(tabDisqF[i].id,",",tabDisqF[i].prenom,",", tabDisqF[i].nom,",", tabDisqF[i].sexe,",", tabDisqF[i].equipe,",", tabDisqF[i].epreuve,",", tabDisqF[i].annonce,",", tabDisqF[i].resultat,",", place,",", tabDisqF[i].rslt,",", tabDisqF[i].plc,",", tabDisqF[i].disq,",", tabDisqF[i].description,"\r\n"))

				pla,_:=strconv.Atoi(place)
				nextResult = newClassement(tabDisqF[i].id,tabDisqF[i].prenom,tabDisqF[i].nom,tabDisqF[i].sexe,tabDisqF[i].equipe,tabDisqF[i].epreuve,tabDisqF[i].annonce, tabDisqF[i].resultat,pla,tabDisqF[i].rslt,tabDisqF[i].plc, tabDisqF[i].disq, tabDisqF[i].description)
				tabPlace=append(tabPlace,nextResult)
		}
		for i := 0; i < len(tabHomme); i++{

				file.WriteString(fmt.Sprint(tabHomme[i].id,",",tabHomme[i].prenom,",", tabHomme[i].nom,",", tabHomme[i].sexe,",", tabHomme[i].equipe,",", tabHomme[i].epreuve,",", tabHomme[i].annonce,",", tabHomme[i].resultat,",", tabHomme[i].place,",", tabHomme[i].rslt,",", tabHomme[i].plc,",", tabHomme[i].disq,",", tabHomme[i].description,"\r\n"))
				file2.WriteString(fmt.Sprint(tabHomme[i].id,",",tabHomme[i].prenom,",", tabHomme[i].nom,",", tabHomme[i].sexe,",", tabHomme[i].equipe,",", tabHomme[i].epreuve,",", tabHomme[i].annonce,",", tabHomme[i].resultat,",", tabHomme[i].place,",", tabHomme[i].rslt,",", tabHomme[i].plc,",", tabHomme[i].disq,",", tabHomme[i].description,"\r\n"))
				nextResult = newClassement(tabHomme[i].id,tabHomme[i].prenom,tabHomme[i].nom,tabHomme[i].sexe,tabHomme[i].equipe,tabHomme[i].epreuve,tabHomme[i].annonce, tabHomme[i].resultat,tabHomme[i].place,tabHomme[i].rslt,tabHomme[i].plc, tabHomme[i].disq, tabHomme[i].description)
				tabPlace=append(tabPlace,nextResult)
		}

		for i := 0; i < len(tabDisqH); i++{
		var place string
				place=strconv.Itoa(numPlaceH+nbH-1)

				file.WriteString(fmt.Sprint(tabDisqH[i].id,",",tabDisqH[i].prenom,",", tabDisqH[i].nom,",", tabDisqH[i].sexe,",", tabDisqH[i].equipe,",", tabDisqH[i].epreuve,",", tabDisqH[i].annonce,",", tabDisqH[i].resultat,",", place,",", tabDisqH[i].rslt,",", tabDisqH[i].plc,",", tabDisqH[i].disq,",", tabDisqH[i].description,"\r\n"))
				file2.WriteString(fmt.Sprint(tabDisqH[i].id,",",tabDisqH[i].prenom,",", tabDisqH[i].nom,",", tabDisqH[i].sexe,",", tabDisqH[i].equipe,",", tabDisqH[i].epreuve,",", tabDisqH[i].annonce,",", tabDisqH[i].resultat,",", place,",", tabDisqH[i].rslt,",", tabDisqH[i].plc,",", tabDisqH[i].disq,",", tabDisqH[i].description,"\r\n"))

				pla,_:=strconv.Atoi(place)
				nextResult = newClassement(tabDisqH[i].id,tabDisqH[i].prenom,tabDisqH[i].nom,tabDisqH[i].sexe,tabDisqH[i].equipe,tabDisqH[i].epreuve,tabDisqH[i].annonce, tabDisqH[i].resultat,pla,tabDisqH[i].rslt,tabDisqH[i].plc, tabDisqH[i].disq, tabDisqH[i].description)
				tabPlace=append(tabPlace,nextResult)
		}

		for i := 0; i < len(tabPlace); i++{
				base.modifResult(tabPlace[i].id ,9, tabPlace[i].epreuve, strconv.Itoa(tabPlace[i].place))
			}
		 }
	}


	/*
	*
	*
	*/
	func (base Bdd) modifResult(idComp int, colNum int,epreuve string, newvalue string){
		idCol, value := idCol2name2(colNum, newvalue)
		id := strconv.Itoa(idComp)
		_, base.err = base.db.Exec("UPDATE classement SET "  + idCol + " = " + value +  " WHERE id = " + id + " AND epreuve = '" + epreuve +"'" )

		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : ", base.err)
			} else {
			//fmt.Println("Modification du competiteur " + strconv.Itoa(idComp) + " avec " + idCol + " = " + value)
		}
	}

	/*
	*
	*/
	func (base Bdd) calculPlace(epreuve string){
	var egalF int
	var egalH int
	var lastResult float64=-1
	var idCol string
		idCol, epreuve = idCol2name2(6, epreuve)

		if( epreuve == "'spd'" || epreuve == "'850'"){
	base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE ", idCol, " = ", epreuve," ORDER BY sexe ASC, rslt ASC"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		}else{
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE ", idCol, " = ", epreuve," ORDER BY sexe ASC, rslt DESC"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 2")
		}
		}
	var info [13]string
	var numPlaceF int =1
	var numPlaceH int =1
	var sexe string ="F"
	var tabClassement []*Classement
	var nextResult *Classement
	var tabH []*Classement
	var tabF []*Classement
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9], &info[10], &info[11], &info[12])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)}

			if (info[9]=="0" && info[11]=="true") || (info[9]=="0" && info[4]=="SOLO"){
				if info[3]=="H"{
			    id,_:=strconv.Atoi(info[0])
				annonce,_ := strconv.Atoi(info[6])
				resultat,_ := strconv.ParseFloat(info[7], 64)
				place,_ := strconv.Atoi(info[8])
				rslt,_ := strconv.ParseFloat(info[9], 64)
				plc,_ := strconv.Atoi(info[10])
				disq,_ := strconv.ParseBool(info[11])
			nextResult = newClassement(id, info[1], info[2], info[3], info[4], info[5], annonce, resultat, place, rslt, plc, disq, info[12])

			tabH=append(tabH,nextResult)
			}else{
			 id,_:=strconv.Atoi(info[0])
				annonce,_ := strconv.Atoi(info[6])
				resultat,_ := strconv.ParseFloat(info[7], 64)
				place,_ := strconv.Atoi(info[8])
				rslt,_ := strconv.ParseFloat(info[9], 64)
				plc,_ := strconv.Atoi(info[10])
				disq,_ := strconv.ParseBool(info[11])
			nextResult = newClassement(id, info[1], info[2], info[3], info[4], info[5], annonce, resultat, place, rslt, plc, disq, info[12])

			tabF=append(tabF,nextResult)}
			}else{
				if(info[3]==sexe){
					info[10]=strconv.Itoa(numPlaceF)
					numPlaceF=numPlaceF+1

					Resultnow,_:=strconv.ParseFloat(info[9], 64)
					if lastResult==Resultnow{
						info[10]=strconv.Itoa(numPlaceF-2-egalF)
						egalF =egalF + 1
					}else{
					egalF=0
					}
					lastResult,_=strconv.ParseFloat(info[9], 64)
				}else{
					info[10]=strconv.Itoa(numPlaceH)
					numPlaceH=numPlaceH+1

					Resultnow,_:=strconv.ParseFloat(info[9], 64)
					if lastResult==Resultnow{
						info[10]=strconv.Itoa(numPlaceH-2-egalH)
						egalH =egalH + 1
					}else{
					egalH=0
					}
					lastResult,_=strconv.ParseFloat(info[9], 64)
					}
			}

				id,_:=strconv.Atoi(info[0])
				annonce,_ := strconv.Atoi(info[6])
				resultat,_ := strconv.ParseFloat(info[7], 64)
				place,_ := strconv.Atoi(info[8])
				rslt,_ := strconv.ParseFloat(info[9], 64)
				plc,_ := strconv.Atoi(info[10])
				disq,_ := strconv.ParseBool(info[11])

				nextResult = newClassement(id, info[1], info[2], info[3], info[4], info[5], annonce, resultat, place, rslt, plc, disq, info[12])
				tabClassement = append(tabClassement, nextResult)
		}
		base.resultat.Close()

			for i := 0; i < len(tabClassement); i++{
				base.modifResult(tabClassement[i].id ,11,tabClassement[i].epreuve, strconv.Itoa(tabClassement[i].plc))
			}

			for i := 0; i < len(tabH); i++{
			placeH :=numPlaceH +len(tabH)-1
				base.modifResult(tabH[i].id ,11,tabH[i].epreuve, strconv.Itoa(placeH))
			}
			for i := 0; i < len(tabF); i++{
			placeF :=numPlaceF +len(tabF)-1
				base.modifResult(tabF[i].id ,11,tabF[i].epreuve,strconv.Itoa(placeF))
			}

	}


	func calculResultat(sexe string, equipe string, epreuve string, annonce int, resultat string, disq string)(float64){
		var sMin int
		var sMax int
		var res float64
		var result float64
		var tot float64
		var tot2 float64
		var tab[] *ConfigurationEpreuve
		var annoncef float64 = float64(annonce)
		var eventMaxTimeCounted = float64(600)

		result,_ = strconv.ParseFloat(resultat, 64)
		tab=getConfigurationEpreuve1()

		for i := 0; i < 5; i++{
			if (tab[i].id==epreuve){
				sMin=tab[i].seuilMin
				sMax=tab[i].seuilMax
			}

			max:=annonce+sMax
			min:=annonce+sMin

			if (equipe == "SOLO" || equipe == "Solo" || equipe == "solo"){
				res=0
			}else{
				if (result == 0 && disq == "true"){
					res=0
				}else{
					if(result>float64(max)){
						switch(epreuve){
						case "spd":
							tot =result + (result-(annoncef)+20)*3
						break
						case "850":
							if(sexe == "F"){
								eventMaxTimeCounted = float64(720)
							}
							if(result > eventMaxTimeCounted){
								tot = eventMaxTimeCounted
							} else {
								tot = result + (result-(annoncef)+40)*3
							}
						break
						case "dnf":
							tot = (annoncef+25)
						break
						case "dwf":
							tot = (annoncef+25)
						break
						case "sta":
							tot = (annoncef+60)
						break
						}
						res=tot
					}else if(result<float64(min)){
						switch(epreuve){
						case "spd":
							tot2=annoncef-10
						break
						case "850":
							tot2=annoncef-20
						break
						case "dnf":
							tot2=result-((annoncef-25)-result)*3
						break
						case "dwf":
							tot2=result-((annoncef-25)-result)*3
						break
						case "sta":
							tot2=result-((annoncef-60)-result)*3
						break
						}
					res=tot2
					}else if (result >= float64(min) && result <= float64(max)){
						res= result
					}
				}
			}
		}
		return res
	}

	/*
	*
	*
	*/
	func getConfigurationEpreuve1()([]*ConfigurationEpreuve){
	file, err := os.Open("../config/ConfigurationEpreuve.csv")
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
	var tabConfig[] *ConfigurationEpreuve
	tabConfig=tabConfig[:0]

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ",")
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
		tabConfig=append(tabConfig,nextconfig)
		}
		firstCall = false
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tabConfig
}
