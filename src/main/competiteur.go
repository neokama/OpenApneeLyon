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
		temps1 float32
		epreuve2 string
		temps2 float32
	}
	
	func (pers Competiteur) disp(){
		fmt.Println(pers.prenom + "; " +
		pers.id + "; " +
		pers.nom + "; " +
		pers.sexe + "; " +
		pers.num_license + "; " +
		pers.equipe + "; " +
		pers.epreuve1 + "; " +
		strconv.FormatFloat(float64(pers.temps1),'f', -1,  32) + "; " +
		pers.epreuve2 + "; " +
		strconv.FormatFloat(float64(pers.temps2),'f', -1,  32) + "\n")
	}
	
	func newcomp(nom string, prenom string)(*Competiteur){
		pers := new(Competiteur)
		pers.nom = nom
		pers.prenom = prenom
		pers.sexe = "M"
		pers.num_license = ""
		pers.epreuve1 =""
		pers.temps1 = 0
		pers.epreuve2 =""
		pers.temps2 = 0
		
		return pers
	}
	
	func newcomp2(id string, prenom string, nom string, sexe string, num_license string, equipe string, epreuve1 string, temps1 float32, epreuve2 string, temps2 float32)(*Competiteur){
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
	