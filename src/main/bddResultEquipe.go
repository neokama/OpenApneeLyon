package main
	
	import (
	"strconv"
	//"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	//"bufio"
	//"strings"
	"time"
	)
		
	/*
	*
	*
	*
	*/
	func (base Bdd) exportClassementEquipe(){ 
	
	base.importEquipe()
	base.importPoint()
	
	
		t := time.Now()
			date := fmt.Sprint(t.Year(),"_",int(t.Month()),"_", t.Day(),"_",t.Hour(),"_", t.Minute(),"_", t.Second())
		
		file, err := os.Create(fmt.Sprint("export/archives/",date,"-equipe.csv"))
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export/archives\" dans le dossier de l'application?")
				log.Fatal(err)
			}
		file2, err := os.Create(fmt.Sprint("export/Classement-Equipe.csv"))
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
			log.Fatal(err)
		}
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classementequipe ORDER BY point ASC"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		
		defer base.resultat.Close()
	var info [4]string
	var numPlace int =0
	var tabPlace []int
	var tab []string
	
	file.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Equipe; Point; Place\r\n"))
	file2.WriteString(fmt.Sprint("\xEF\xBB\xBFId; Equipe; Point; Place\r\n"))	
		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)}				
			
				numPlace=numPlace+1
				tabPlace=append(tabPlace,numPlace)
				tab=append(tab,info[1])

		file.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", numPlace,"\r\n"))
		file2.WriteString(fmt.Sprint(info[0],";",info[1],";", info[2],";", numPlace,"\r\n"))
		}
		for n:=0;n<len(tabPlace);n++{
		   	base.modifPlace(tab[n],strconv.Itoa(tabPlace[n]))
			}	
		
	}
	
	
	func col_id2name3(col_num int, value string)(string, string){
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
			case 4:
				col_idr = "place"
				value = fmt.Sprint("'",value,"'")
			default:
				log.Fatal("Numéro invalide")
			}
		return col_idr, value
	}
	
	func (base Bdd) importEquipe(){
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM competiteurs GROUP BY equipe"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		var info [10]string
		var tab []string
		var equipe string
		for base.resultat.Next() {
		base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)}
			equipe = info[5]
		tab=append(tab,equipe)
		
		}
		for n:=0;n<len(tab);n++{
		classemt := newClassementE(0, tab[n],0 ,0)
		base.addEquipe(classemt)
		}
	}
	
	func (base Bdd) importPoint(){
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classementequipe GROUP BY equipe"))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		var info [4]string
		var tab []string
		var equipe string
		for base.resultat.Next() {
		base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)}
			equipe = info[1]
		tab=append(tab,equipe)
		
		}
		
		for n:=0;n<len(tab);n++{
		
		var id_col string
		var equipe2 string
		id_col, equipe2 = col_id2name2(5, tab[n])
		base.resultat, base.err = base.db.Query(fmt.Sprint("SELECT * FROM classement WHERE " + id_col + " = " + equipe2))
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête 1")
		}
		defer base.resultat.Close()
		var info [13]string
		var tabpoint []string
		var point string
		var nbpoint int
		var nb int
		for base.resultat.Next() {
		
		base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3],&info[4], &info[5], &info[6], &info[7],&info[8], &info[9], &info[10], &info[11], &info[12])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)}
				
			point = info[10]
			
		tabpoint=append(tabpoint,point)
		}
			for n:=0;n<len(tabpoint);n++{
			nb,_=strconv.Atoi(tabpoint[n])
		    nbpoint=nbpoint+nb
			}	
			base.modifPoint(tab[n],strconv.Itoa(nbpoint))
		}
		
		
	}
	
	func (base Bdd) modifPoint(equipe string, newvalue string){
		col_id2, equipe := col_id2name3(2, equipe)
		col_id, value := col_id2name3(3, newvalue)
		if(col_id2=="0"){
		//on doit utiliser col_id2
		}
		_, base.err = base.db.Exec("UPDATE classementequipe SET "  + col_id + " = " + value +  " WHERE equipe = " + equipe)
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : ", base.err)
			} else {
			//fmt.Println("Modification du competiteur " + strconv.Itoa(id_comp) + " avec " + col_id + " = " + value)
		}
	}
	
	func (base Bdd) modifPlace(equipe string, newvalue string){
		col_id2, equipe := col_id2name3(2, equipe)
		col_id, value := col_id2name3(4, newvalue)
		if(col_id2=="0"){
		//on doit utiliser col_id2
		}
		_, base.err = base.db.Exec("UPDATE classementequipe SET "  + col_id + " = " + value +  " WHERE equipe = " + equipe)
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : ", base.err)
			} else {
			//fmt.Println("Modification du competiteur " + strconv.Itoa(id_comp) + " avec " + col_id + " = " + value)
		}
	}
	
	func (base Bdd) addEquipe(board *ClassementEquipe){
		
		_, base.err = base.db.Exec("INSERT INTO classementequipe (equipe, point, place) VALUES('" +
		board.equipe + "'," +
		strconv.Itoa(board.point) + "," +
		strconv.Itoa(board.place)+")")
	
		if base.err != nil {
			fmt.Println("Echec lors de l'ajout : "+ board.equipe, base.err)
			} else {
			fmt.Println("Ajout validé de l'equipe " + board.equipe)
		}
	}
	
	func (base Bdd) displayEquipe(){
		base.resultat, base.err = base.db.Query("SELECT * FROM classementequipe")
		if base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(base.err)
		}
		defer base.resultat.Close()
		
		var info [4]string

		for base.resultat.Next() {
			base.err = base.resultat.Scan(&info[0], &info[1], &info[2], &info[3])
			if base.err != nil {
				fmt.Println("Erreur lors de la récupération des résultats: \n")
				log.Fatal(base.err)
			}
		fmt.Println(info[0] + "|" + info[1]+ "|" + info[2]+ "|" + info[3])
		}
	}