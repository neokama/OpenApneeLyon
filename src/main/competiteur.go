package main

import (
	"fmt"
	"strconv"
	)
	
	type Competiteur struct
	{
		id string
		prenom string
		nom string
		sexe string
		num_license string
		equipe string
		epreuve1 string
		temps1 int
		epreuve2 string
		temps2 int
	}
	
	func (pers Competiteur) display(){
		fmt.Println(pers.id + "; " +
		pers.prenom + "; " +
		pers.nom + "; " +
		pers.sexe + "; " +
		pers.num_license + "; " +
		pers.equipe + "; " +
		pers.epreuve1 + "; " +
		strconv.Itoa(pers.temps1) + "; " +
		pers.epreuve2 + "; " +
		strconv.Itoa(pers.temps2))
	}
	
	func newCompetiteur(id string, prenom string, nom string, sexe string, num_license string, equipe string, epreuve1 string, temps1 int, epreuve2 string, temps2 int)(*Competiteur){
		pers := new(Competiteur)
		pers.id = id
		pers.nom = nom
		pers.prenom = prenom
		pers.sexe = sexe
		pers.num_license = num_license
		pers.equipe = equipe
		pers.epreuve1 = epreuve1
		pers.temps1 = temps1
		pers.epreuve2 = epreuve2
		pers.temps2 = temps2
		
		return pers
	}
	