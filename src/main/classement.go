package main
	
	import (
	"strconv"
	"fmt"
	)
	
	type Classement struct 
	{
	id int
	nom string 
	prenom string
	sexe string
	equipe string
	epreuve string
	annonce int
	resultat int 
	place int
	}
	
	func (board Classement) display(){
		fmt.Println(strconv.Itoa(board.id) + "; " +
		board.prenom + "; " +
		board.nom + "; " +
		board.sexe + "; " +
		board.equipe + "; " +
		board.epreuve + "; " +
		strconv.Itoa(board.annonce)+ "; " +
		strconv.Itoa(board.resultat)+ "; " +
		strconv.Itoa(board.place))
	}
	
	func newClassement(id int, prenom string, nom string, sexe string, equipe string, epreuve string, annonce int, resultat int, place int)(*Classement){
		board := new(Classement)
		board.id = id
		board.nom = nom
		board.prenom = prenom
		board.sexe = sexe
		board.equipe = equipe
		board.epreuve = epreuve
		board.annonce = annonce
		board.resultat = resultat
		board.place = place
		
		return board
	}
	