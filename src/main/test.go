package main

import (
	"fmt"
	)

	func testSuite(){
		var errPrenom, errNom, errSexe bool
		errPrenom = testCompetiteurPrenom()	
		errNom = testCompetiteurNom()
		errSexe = testCompetiteurSexe()
		
		fmt.Println("\n\nSYNTHESE:")
		if errPrenom {
			fmt.Println("PRENOM: ERROR")
		} else{
			fmt.Println("PRENOM: OK")
		}
		if errNom {
			fmt.Println("NOM: ERROR")
		} else{
			fmt.Println("NOM: OK")
		}
		if errSexe {
			fmt.Println("SEXE: ERROR")
		} else{
			fmt.Println("SEXE: OK")
		}
		
		
	}
	
	
	//TEST SUR LE PRENOM
	func testCompetiteurPrenom()(bool){
		var test bool
		var err bool
		comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe" , "sta", 123, "spd",456)
		test = true
		err = false
		
		
		fmt.Println("CHECK SUR LE PRENOM: \n")
		fmt.Println("TEST 1: Nom standard")
		test = comp.check()
		if (!test){
			fmt.Println("TEST PRENOM 1:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST PRENOM 1:	OK\n")
		}
		
		fmt.Println("TEST 2: PRENOM AVEC ACCENT")
		comp.prenom = "Frédèrîqûê"
		test = comp.check()
		if (!test){
			fmt.Println("TEST PRENOM 2:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST PRENOM 2:	OK\n")
		}
		
		fmt.Println("TEST 3: PRENOM AVEC TIRET")
		comp.prenom = "Jean-Eude-Saint-Jean"
		test = comp.check()
		if (!test){
			fmt.Println("TEST PRENOM 3:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST PRENOM 3:	OK\n")
		}
		
		fmt.Println("TEST 4: PRENOM AVEC SYMBOLE")
		comp.prenom = "Jeanª╝É¾àzaea"
		test = comp.check()
		if (!test){
			fmt.Println("TEST PRENOM 4:	OK\n")
			test = true
		} else {
			fmt.Println("TEST PRENOM 4:	ERROR\n")
			err = true
		}
		
		fmt.Println("TEST 5: PRENOM AVEC CHIFFRES")
		comp.prenom = "Bon123jour14"
		test = comp.check()
		if (!test){
			fmt.Println("TEST PRENOM 5:	OK\n")
			test = true
		} else {
			fmt.Println("TEST PRENOM 5:	ERROR\n")
			err = true
		}
		return err
	}
	
	//TEST SUR LE NOM
	func testCompetiteurNom()(bool){
		
		var test bool
		var err bool
		comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe" , "sta", 123, "spd",456)
		test = true
		err = false
		
		fmt.Println("CHECK SUR LE NOM: \n")
		fmt.Println("TEST 1: Nom standard")
		test = comp.check()
		if (!test){
			fmt.Println("TEST NOM 1:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST NOM 1:	OK\n")
		}
		
		fmt.Println("TEST 2: NOM AVEC ACCENT")
		comp.prenom = "Frédèrîqûê"
		test = comp.check()
		if (!test){
			fmt.Println("TEST NOM 2:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST NOM 2:	OK\n")
		}
		
		fmt.Println("TEST 3: NOM AVEC TIRET")
		comp.prenom = "Jean-Eude-Saint-Jean"
		test = comp.check()
		if (!test){
			fmt.Println("TEST NOM 3:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST NOM 3:	OK\n")
		}
		
		fmt.Println("TEST 4: NOM AVEC SYMBOLE")
		comp.prenom = "Jeanª╝É¾àzaea"
		test = comp.check()
		if (!test){
			fmt.Println("TEST NOM 4:	OK\n")
			test = true
		} else {
			fmt.Println("TEST NOM 4:	ERROR\n")
			err = true
		}
		
		fmt.Println("TEST 5: NOM AVEC CHIFFRES")
		comp.prenom = "Bon123jour14"
		test = comp.check()
		if (!test){
			fmt.Println("TEST NOM 5:	OK\n")
			test = true
		} else {
			fmt.Println("TEST NOM 5:	ERROR\n")
			err = true
		}
		
		return err
	}
	
	
	//TEST SUR LE SEXE
	func testCompetiteurSexe()(bool){
		
		var test bool
		var err bool
		comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe" , "sta", 123, "spd",456)
		test = true
		err = false
		
		fmt.Println("CHECK SUR LE SEXE: \n")
		fmt.Println("TEST 1: SEXE HOMME")
		test = comp.check()
		if (!test){
			fmt.Println("TEST SEXE 1:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST SEXE 1:	OK\n")
		}
		
		fmt.Println("TEST 2: SEXE FEMME")
		comp.sexe = "F"
		test = comp.check()
		if (!test){
			fmt.Println("TEST SEXE 2:	ERROR\n")
			test = true
			err = true
		} else {
			fmt.Println("TEST SEXE 2:	OK\n")
		}
		
		fmt.Println("TEST 3: SEXE PLUSIEURS CARACTERES")
		comp.sexe = "FFF"
		test = comp.check()
		if (!test){
			fmt.Println("TEST SEXE 3:	OK\n")
			test = true

		} else {
			fmt.Println("TEST SEXE 3:	ERROR\n")
			err = true
		}
		
		fmt.Println("TEST 4: SEXE NON H/F")
		comp.sexe = "I"
		test = comp.check()
		if (!test){
			fmt.Println("TEST SEXE 4:	OK\n")
			test = true
		} else {
			fmt.Println("TEST SEXE 4:	ERROR\n")
			err = true
		}
		
		fmt.Println("TEST 5: SEXE CHIFFRES")
		comp.sexe = "4"
		test = comp.check()
		if (!test){
			fmt.Println("TEST SEXE 5:	OK\n")
			test = true
		} else {
			fmt.Println("TEST SEXE 5:	ERROR\n")
			err = true
		}
		return err
	}

	