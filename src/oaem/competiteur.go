package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Competiteur struct
type Competiteur struct {
	id         int
	prenom     string
	nom        string
	sexe       string
	numLicense string
	equipe     string
	epreuve1   string
	annonce1   int
	epreuve2   string
	annonce2   int
}

func (person Competiteur) String() string {
	return fmt.Sprintf("[%d, %s, %s, %s, %s, %s, %s, %d, %s, %d]", person.id,
		person.prenom,
		person.nom,
		person.sexe,
		person.numLicense,
		person.equipe,
		person.epreuve1,
		person.annonce1,
		person.epreuve2,
		person.annonce2)
}

func (person Competiteur) display() {
	fmt.Println(strconv.Itoa(person.id) + ";" +
		person.prenom + ";" +
		person.nom + ";" +
		person.sexe + ";" +
		person.numLicense + ";" +
		person.equipe + ";" +
		person.epreuve1 + ";" +
		strconv.Itoa(person.annonce1) + ";" +
		person.epreuve2 + ";" +
		strconv.Itoa(person.annonce2))
}

func (person Competiteur) check() bool {
	var verif = true
	verif = true
	match, _ := regexp.MatchString("^[\\p{L}- ]*$", person.prenom)
	if !match {
		verif = false
		fmt.Println("Erreur! Format du prénom.")
	}
	match, _ = regexp.MatchString("^[\\p{L}- ]*$", person.nom)
	if !match {
		verif = false
		fmt.Println("Erreur! Format du nom.")
	}
	match, _ = regexp.MatchString("([F|H])", person.sexe)
	if !match || len(person.sexe) > 1 {
		verif = false
		fmt.Println("Erreur! Format du sexe.")
	}
	match, _ = regexp.MatchString("^[A-Za-z0-9-]*$", person.numLicense)
	if !match {
		verif = false
		fmt.Println("Erreur! Format du numéro de license.")
	}
	match, _ = regexp.MatchString("^[\\p{L}0-9- _]*$", person.equipe)
	if !match {
		verif = false
		fmt.Println("Erreur! Format du nom d'équipe.")
	}
	if person.epreuve1 != "sta" && person.epreuve1 != "spd" && person.epreuve1 != "dwf" && person.epreuve1 != "dnf" && person.epreuve1 != "850" {
		verif = false
		fmt.Println("Erreur! Format du epreuve1 (Rappel des valeurs possibles: sta, spd, dwf, dnf, 850).")
	}
	match, _ = regexp.MatchString("(^[0-9]*$)", strconv.Itoa(person.annonce1))
	if !match {
		verif = false
		fmt.Println("Erreur! Format du annonce1.")
	}
	if person.epreuve2 != "sta" && person.epreuve2 != "spd" && person.epreuve2 != "dwf" && person.epreuve2 != "dnf" && person.epreuve2 != "850" {
		verif = false
		fmt.Println("Erreur! Format du epreuve2 (Rappel des valeurs possibles: sta, spd, dwf, dnf, 850).")
	}
	match, _ = regexp.MatchString("(^[0-9]*$)", strconv.Itoa(person.annonce2))
	if !match {
		verif = false
		fmt.Println("Erreur! Format du annonce2.")
	}

	if verif == true {
		return true
	}
	return false
}

func newCompetiteur(id int, prenom string, nom string, sexe string, numLicense string, equipe string, epreuve1 string, annonce1 int, epreuve2 string, annonce2 int) *Competiteur {
	person := new(Competiteur)
	person.id = id
	person.nom = nom
	person.prenom = prenom
	person.sexe = sexe
	person.numLicense = numLicense
	person.equipe = equipe
	person.epreuve1 = epreuve1
	person.annonce1 = annonce1
	person.epreuve2 = epreuve2
	person.annonce2 = annonce2

	return person
}
