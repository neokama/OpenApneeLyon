package main

import (
	"fmt"
	"strconv"
	"regexp"
	)
	
	// Competiteur struct
	type Competiteur struct
	{
		id int
		prenom string
		nom string
		sexe string
		numLicense string
		equipe string
		epreuve1 string
		annonce1 int
		epreuve2 string
		annonce2 int
	}
	
	func (pers Competiteur) display(){
		fmt.Println(strconv.Itoa(pers.id) + ";" +
		pers.prenom + ";" +
		pers.nom + ";" +
		pers.sexe + ";" +
		pers.numLicense + ";" +
		pers.equipe + ";" +
		pers.epreuve1 + ";" +
		strconv.Itoa(pers.annonce1) + ";" +
		pers.epreuve2 + ";" +
		strconv.Itoa(pers.annonce2))
	}
	
	func (pers Competiteur) check()(bool){
		var verif = true
		verif = true
		match, _ := regexp.MatchString("^[\\p{L}- ]*$", pers.prenom )
			if(!match){
				verif =false
				fmt.Println("Erreur! Format du prénom.")
			}
		match, _ = regexp.MatchString("^[\\p{L}- ]*$", pers.nom )
			if(!match){
				verif =false
				fmt.Println("Erreur! Format du nom.")
			}
		match, _ = regexp.MatchString("([F|H])", pers.sexe )
			if(!match || len(pers.sexe) > 1){
				verif =false
				fmt.Println("Erreur! Format du sexe.")
			}
		match, _ = regexp.MatchString("^[A-Za-z0-9-]*$", pers.numLicense )
			if(!match){
				verif =false
				fmt.Println("Erreur! Format du numéro de license.")
			}
		match, _ = regexp.MatchString("^[\\p{L}0-9- _]*$", pers.equipe )
			if(!match){
				verif =false
				fmt.Println("Erreur! Format du nom d'équipe.")
			}
			if(pers.epreuve1!="sta" && pers.epreuve1 !="spd" && pers.epreuve1!="dwf" && pers.epreuve1!="dnf" && pers.epreuve1!="850"){
				verif =false
				fmt.Println("Erreur! Format du epreuve1 (Rappel des valeurs possibles: sta, spd, dwf, dnf, 850).")
			}
		match, _ = regexp.MatchString("(^[0-9]*$)", strconv.Itoa(pers.annonce1))
			if(!match){
				verif = false
				fmt.Println("Erreur! Format du annonce1.")
			}
			if(pers.epreuve2 !="sta" && pers.epreuve2 !="spd" && pers.epreuve2 !="dwf" && pers.epreuve2 !="dnf" && pers.epreuve2 !="850"){
				verif = false
				fmt.Println("Erreur! Format du epreuve2 (Rappel des valeurs possibles: sta, spd, dwf, dnf, 850).")
			}
		match, _ = regexp.MatchString("(^[0-9]*$)", strconv.Itoa(pers.annonce2))
			if(!match){
				verif = false
				fmt.Println("Erreur! Format du annonce2.")
			}
			
			if (verif == true) {
				return true
			}
			return false
	}
 
	func newCompetiteur(id int, prenom string, nom string, sexe string, numLicense string, equipe string, epreuve1 string, annonce1 int, epreuve2 string, annonce2 int)(*Competiteur){
		pers := new(Competiteur)
		pers.id = id
		pers.nom = nom
		pers.prenom = prenom
		pers.sexe = sexe
		pers.numLicense = numLicense
		pers.equipe = equipe
		pers.epreuve1 = epreuve1
		pers.annonce1 = annonce1
		pers.epreuve2 = epreuve2
		pers.annonce2 = annonce2
		
		return pers
	}
	