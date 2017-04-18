package main

import (
	"fmt"
	//"log"
	//"bufio"
	//"io/ioutil"
	//"strings"
	//"os"
	"time"
	
	)
	func main(){
	/*base := newBdd("../src/database/OpenApneeLyon")
	base.reset()
	fmt.Println("\n")
	base.importCompetiteur()
	fmt.Println("\n ")
	base.disp_comp()
	fmt.Println("\n unicité ->")
	base.uniqueness()
	fmt.Println("\n Trier ->")
	base.orderby_comp()
	fmt.Println("\n Verification équipe ->")
	base.check_team()
	*/
	
	
	
	 //p := fmt.Println

    // Here's a basic example of formatting a time
    // according to RFC3339, using the corresponding layout
    // constant.
    t := time.Now()

  var date string 
	date=t.Format("2006-01-02_15:04:05")
	fmt.Printf("date ="+date)
	}

	
	
	
	
	
	
	
	
	
	
	/*func Compare(a, b string) int

Compare returns an integer comparing two strings 
lexicographically. The result will be 0 if a==b, -1 if a < b, 
and +1 if a > b. 
	*/
	
	// vérification pour l'identifiant 
	// fmt.Println(strings.HasPrefix("Gopher", "Go"))  true ou false 
	
	/*func main(){
	
	stream, err := ioutil.ReadFile("testbdd.csv")
	if err != nil {
	log.Fatal(err)
	
	}
	
	readString := string(stream)
	fmt.Println(readString)*/
	
	
	/*file, err := os.Open("testbdd.csv")
		scanner := bufio.NewScanner(file)
		var team string = "" 
		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ";")
			
			team = info[5]
			
			fmt.Println("Equipe :",team)
			if team !=""{
			break}
			
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		
	
	
	for scanner.Scan() {
		
			inf := strings.Split(scanner.Text(), ";")
			
		if inf[5]==team{
			fmt.Println(inf[0],inf[1],inf[2],inf[3],inf[4],inf[5],inf[6],inf[7],inf[8],inf[9])
			}
			
			} */
	//	teamSplit(3)

	//}
	/*
	* 		teamSplit:
	* Description: 	
	* 		Méthode permettant d'afficher l'integralité des
	* 		des compétiteur en fonction de leur équipe
	*/
	
	/*func teamSplit(nb int)() {
	var teamName [20] string 
	var i int = 0
	var j int
	var k int
	for i <= nb { //=  !!!
	teamName[i]=""
	i=i+1
	}
	fmt.Println("Debut teamSplit :")
	file, _ := os.Open("testbdd.csv")
		scanner := bufio.NewScanner(file)
	for j = 0; j < 20; j++{
	//fmt.Println("j:",j)
	if j==nb{break}
	if teamName[j]==""{
		
		
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		//fmt.Println(teamName[0],"et",teamName[1])
		for scanner.Scan() {
			info := strings.Split(scanner.Text(), ";")
			//fmt.Println(info[0],info[5])
			for k = 0; k <= j; k++{
			//fmt.Println(k)
			//fmt.Println(strings.EqualFold(info[5],teamName[k]))
			if (strings.EqualFold(info[5],teamName[k])==false){
			
			teamName[j] = info[5]
			fmt.Println("Equipe :",teamName[j])
			}else{//fmt.Println("casse toi")
			break
			}
			}
			if teamName[j] !=""{
			//fmt.Println("casse")
			break}
			
		}
		
		file2, _ := os.Open("testbdd.csv")
		scanner2 := bufio.NewScanner(file2)
		if err := scanner2.Err(); err != nil {
			log.Fatal(err)
		}
		defer file2.Close()
	
	for scanner2.Scan() {
		
			inf := strings.Split(scanner2.Text(), ";")
			
		if inf[5]==teamName[j]{
			fmt.Println(inf[0],inf[1],inf[2],inf[3],inf[4],inf[5],inf[6],inf[7],inf[8],inf[9])
			}
			
			} 	
	
	}
	
	}
	
	}	*/
	
		/*	type Fichier struct
	{
		chemincsv string
		err error
	}
	func newFile(chemincsv string)(*Fichier){
		file := new(Bdd)
		
		file.chemincsv = chemincsv
		
		file.db, file.err = os.Open(base.chemincsv)
		if file.err != nil {
		log.Fatal("Erreur de connection à la base de données:\n", file.err)
		}
		
		
		return file
	}*/
	