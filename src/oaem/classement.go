package main
	
	import (
	"strconv"
	"fmt"
	)
	
	// Classement struct 
	type Classement struct 
	{
	id int
	nom string 
	prenom string
	sexe string
	equipe string
	epreuve string
	annonce int
	resultat float64 
	place int
	rsltEquipe float64
	plc int
	disq bool
	description string
	rsltIndiv float64
	}
	
	func (board Classement) display(){
		fmt.Println(strconv.Itoa(board.id) + ";" +
		board.prenom + ";" +
		board.nom + ";" +
		board.sexe + ";" +
		board.equipe + ";" +
		board.epreuve + ";" +
		strconv.Itoa(board.annonce)+ ";" +
		strconv.FormatFloat(board.resultat, 'f', -1, 64)+ ";" +
		strconv.Itoa(board.place)+ ";" +
		strconv.FormatFloat(board.rsltEquipe, 'f', -1, 64)+ ";" +
		strconv.Itoa(board.plc) + ";" +
		strconv.FormatBool(board.disq) + ";" +
		board.description + ";" +
		strconv.FormatFloat(board.rsltIndiv, 'f', -1, 64))
	}
	
	func newClassement(id int, prenom string, nom string, sexe string, equipe string, epreuve string, annonce int, resultat float64, place int, rsltEquipe float64, plc int, disq bool, description string, rsltIndiv float64)(*Classement){
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
		board.rsltEquipe = rsltEquipe
		board.plc = plc
		board.disq = disq
		board.description = description
		board.rsltIndiv = rsltIndiv
		
		return board
	}
	