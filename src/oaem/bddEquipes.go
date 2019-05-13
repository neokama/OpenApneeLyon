package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
	)

	/*
	* 		EtatEquipe
	* Description:
	* La structure permet de connaitre l'etat d'une equipe. On calcule le classement equipe avec uniquement des equipes valides
	*
	* Paramètres:
	*	- equipe : nom de l'équipe
	*	- etat : validiter de l'equipe
	**/
	// EtatEquipe struct
	type EtatEquipe struct
	{
		equipe string
		etat bool
	}

	/*
	* 		Bdd.orderByComp:
	*
	* Description:
	*		Méthode permettant de trier les compétiteurs
	*		par équipe.
	*/

	func (base Bdd) orderByComp(){


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
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3] + "|" + info[4]+ "|" + info[5]+ "|" + info[6]+ "|" + info[7]+ "|" + info[8]+ "|" + info[9])
		}
	}

	/*
	* 		Bdd.countComp:
	*
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe
	*
	*/
	func (base Bdd) countComp(colNum int, value string)(string){
	var idCol string
		idCol, value = idCol2name(colNum, value)


		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", idCol, " = ", value))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()

		var info [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}

		}
		return info[0]
	}

	/*
	* 		Bdd.countEpreuveComp:
	*
	* Description: Méthode permettent de vérifier le nombre de compétiteur par équipe
	*
	* Edition 2019 : Order should be as : STA, DWF, SPE100, DNF, 8x50m
	*/
	func (base Bdd) countEpreuveComp(colNum int, value string)(bool,bool,bool,bool,bool,int,int,int,int,int){
	var idCol string
		idCol, value = idCol2name(colNum, value)


		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs WHERE ", idCol, " = ", value))
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
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
			switch (info[6]){
			case "Statique", "sta":
			nbSTA= nbSTA+1
			if (info[8]=="DWF" || info[8]=="dwf" && info[8]!=info[6]) {
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "DWF", "dwf":
			nbDWF= nbDWF+1
			if (info[8]=="Speed 100" || info[8]=="Statique" || info[8]=="spd" || info[8]=="sta"  && info[8]!=info[6]) {
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "Speed 100", "spd":
			nbSPE= nbSPE+1
			if (info[8]=="DNF" || info[8]=="DWF" || info[8]=="dnf" || info[8]=="dwf" && info[8]!=info[6]) {
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "DNF", "dnf":
			nbDNF= nbDNF+1
			if (info[8]=="8*50" || info[8]=="Speed 100" || info[8]=="850" || info[8]=="Spd" && info[8]!=info[6]) {
			switch(n){
			case 1: comp1=false
			case 2: comp2=false
			case 3: comp3=false
			case 4: comp4=false
			case 5: comp5=false
			}
			}
			case "8*50", "850":
			nbSFC= nbSFC+1
			if (info[8]=="DNF" || info[8]=="dnf" && info[8]!=info[6]){
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
			case "8*50":
			nbSFC= nbSFC+1
			case "sta":
			nbSTA= nbSTA+1
			case "dwf":
			nbDWF= nbDWF+1
			case "spd":
			nbSPE= nbSPE+1
			case "dnf":
			nbDNF= nbDNF+1
			case "850":
			nbSFC= nbSFC+1
			}
		}
		return comp1, comp2, comp3, comp4, comp5,nbSTA,nbDWF,nbSPE,nbDNF,nbSFC
	}

	/*
	* 		Bdd.countSexeComp:
	* Description:
	*		Méthode permettent de vérifier le nombre de compétiteur par équipe.
	*
	*/
	func (base Bdd) countSexeComp(colNum int, value string)(string,string){
	var idCol string
	var idCol2 string
	var colNum2 int = 4
	var valueH string = "H"
	var valueF string = "F"
		idCol, value = idCol2name(colNum, value)
		idCol2, valueH = idCol2name(colNum2, valueH)
		idCol2, valueF = idCol2name(colNum2, valueF)


		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", idCol, " = ", value," AND ", idCol2, " = ", valueH))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()

		var infoH [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&infoH[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
		}
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", idCol, " = ", value," AND ", idCol2, " = ", valueF))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()

		var infoF [1]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&infoF[0])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: ")
				log.Fatal(base.err)
			}
		}
		return infoH[0],infoF[0]
	}


	/*
	* 		Bdd.checkTeam:
	* Description:
	*		Méthode permettant de vérifier la validité des equipes
	*/

	func (base Bdd) checkTeam(){
		// CREATION FICHIER
		t := time.Now()
		date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())

		file, err := os.Create(fmt.Sprint("../var/export/archives/",date,"-FichierVerification.txt"))
		file2, err := os.Create(fmt.Sprint("../var/export/FichierVerification.txt"))
		if err != nil {
				fmt.Println("Erreur lors de la création du fichier verification")
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
		var nbSolo int =0

		var tabEquipe []*EtatEquipe
		//On "clear" l'ancien tableau:
			tabEquipe = tabEquipe[:0]


		for base.resultat.Next() {
			if info[0]=="SOLO"{
				nbSolo = nbSolo + 1
			//Faire vérification sur temps de repos
			}else{
				var ep1 bool
				var ep2 bool
				var ep3 bool
				var ep4 bool
				var ep5 bool
				var etat1 bool =false
				var etat2 bool =false
				var etat3 bool =false
				var etat4 bool =false
				var etat5 bool =false
				var etat bool =false
				EquipeEnCours := newEtatEquipe("vide",false)
				var nbSTA int =0
				var nbDWF int =0
				var nbSPE int =0
				var nbDNF int =0
				var nbSFC int =0
				var res string
				var res2 string
				var nb_sexeH string ="0"
				var nb_sexeF string ="0"
				base.err = base.resultat.Scan(&info[0])
				if base.err != nil {
					fmt.Println("Erreur lors de la récupération des résultats: ")
					log.Fatal(base.err)
				}
				var nbComp string = base.countComp(6,info[0])
				nb_sexeH,nb_sexeF=base.countSexeComp(6,info[0])
				ep1,ep2,ep3,ep4,ep5,nbSTA,nbDWF,nbSPE,nbDNF,nbSFC=base.countEpreuveComp(6,info[0])

				if(ep1 && ep2 && ep3 && ep4 && ep5){
				res = "Temps repos : OK"
				etat1 =true
				}else{
				res = "Erreur temps de repos"
				etat1 =false
				}

				if(nbSTA==2 && nbDNF==2 && nbDWF==2 && nbSFC==2 && nbSPE==2){
				res2="Nombres épreuves corrects"
				etat2=true
				}else{
				res2="Erreur sur nombres épreuves"
				etat2=false
				}

				file.WriteString(info[0] + "|" + nbComp + "|" + "Homme : "+ nb_sexeH + "|" + "Femme : "+ nb_sexeF+ "|" + res +"|"+ res2 +"\r\n" )
				file2.WriteString(info[0] + "|" + nbComp + "|" + "Homme : "+ nb_sexeH + "|" + "Femme : "+ nb_sexeF+ "|" + res +"|"+ res2 +"\r\n" )

				if (nbComp!="5"){
					file.WriteString("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nbComp + " compétiteurs !\r\n")
					file2.WriteString("Erreur nombre de compétiteur dans l'equipe "+ info[0] +" où il y a "+ nbComp + " compétiteurs !\r\n")
					etat3=false
				}else{
					etat3=true
				}

				if (nb_sexeH != "3"){
					file.WriteString("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !\r\n")
					file2.WriteString("Erreur nombre d'homme dans l'equipe " + info[0] + " où il y a "+ nb_sexeH + " hommes !\r\n")
					etat4 =false
				}else{
				etat4 =true
				}

				if (nb_sexeF != "2"){
					file.WriteString("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !\r\n")
					file2.WriteString("Erreur nombre de femme dans l'equipe " + info[0] + " où il y a " + nb_sexeF + " femmes !\r\n")
					etat5 =false
				}else{
					etat5=true

				}

				if (etat1==true && etat2==true && etat3==true && etat4==true && etat5==true){
					etat=true
					fmt.Println("Equipe : ",info[0],", Etat : OK")
				}else{
					etat=false
					fmt.Println("Equipe : ",info[0],", Etat : Erreur")

				}

				EquipeEnCours.equipe=info[0]
				EquipeEnCours.etat=etat
				tabEquipe=append(tabEquipe,EquipeEnCours)
			}
		}

		for n:=0;n<len(tabEquipe);n++{

		  base.modifEtat(tabEquipe[n].equipe,tabEquipe[n].etat)
		}
		if nbSolo != 0{
			fmt.Println("Présence de ",nbSolo," compétiteurs dans l'équipe SOLO")
		}
	}

		/*
	* 		Bdd.verif:
	* Description:
	*		Méthode permettant d'effectuer des vérifications sur la validité des champs de la bdd
	*/
	func (base Bdd) verif(val string, num int ){
		var idCol string
		var value string = val
		if num == 1{
		idCol, value = idCol2name(1, value)
		}else{idCol, value = idCol2name(5, value)}
			base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT COUNT(*) FROM competiteurs WHERE ", idCol, " = ", value))
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
	/*
	* 		newEtatEquipe:
	* Description:
	* 			La méthode permet de retourner un EtatEquipe
	* Paramètres:
	*	- equipe : nom de l'équipe
	*	- etat : validiter de l'equipe
	* Sortie:
	*	- une structure EtatEquipe
	*/
	func newEtatEquipe(equipe string, etat bool)(*EtatEquipe){
		board := new(EtatEquipe)
		board.equipe = equipe
		board.etat = etat

		return board
	}
