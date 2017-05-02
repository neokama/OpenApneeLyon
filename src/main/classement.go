package main
	
	import (
	"strconv"
	//"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
	)
	
	type Classement struct 
	{
	id int
	nom string 
	prenom string
	sexe string
	equipe string
	epreuve string
	resultat int 
	}
	
	func (board Classement) display(){
		fmt.Println(strconv.Itoa(board.id) + "; " +
		board.prenom + "; " +
		board.nom + "; " +
		board.sexe + "; " +
		board.equipe + "; " +
		board.epreuve + "; " +
		strconv.Itoa(board.resultat))
	}
	
	func newClassement(id int, prenom string, nom string, sexe string, equipe string, epreuve string, resultat int)(*Classement){
		board := new(Classement)
		board.id = id
		board.nom = nom
		board.prenom = prenom
		board.sexe = sexe
		board.equipe = equipe
		board.epreuve = epreuve
		board.resultat = resultat
		
		return board
	}
	