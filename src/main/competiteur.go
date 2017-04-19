package main

import (
	"fmt"
	"strconv"
	)
	
	type Competiteur struct
	{
		id int
		prenom string
		nom string
		sexe string
		num_license string
		equipe string
		epreuve1 string
		annonce1 int
		epreuve2 string
		annonce2 int
	}
	
	func (pers Competiteur) display(){
		fmt.Println(strconv.Itoa(pers.id) + "; " +
		pers.prenom + "; " +
		pers.nom + "; " +
		pers.sexe + "; " +
		pers.num_license + "; " +
		pers.equipe + "; " +
		pers.epreuve1 + "; " +
		strconv.Itoa(pers.annonce1) + "; " +
		pers.epreuve2 + "; " +
		strconv.Itoa(pers.annonce2))
	}
	
	func newCompetiteur(id int, prenom string, nom string, sexe string, num_license string, equipe string, epreuve1 string, annonce1 int, epreuve2 string, annonce2 int)(*Competiteur){
		pers := new(Competiteur)
		pers.id = id
		pers.nom = nom
		pers.prenom = prenom
		pers.sexe = sexe
		pers.num_license = num_license
		pers.equipe = equipe
		pers.epreuve1 = epreuve1
		pers.annonce1 = annonce1
		pers.epreuve2 = epreuve2
		pers.annonce2 = annonce2
		
		return pers
	}
	