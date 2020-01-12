package main

import (
	"fmt"
)

func testSuite() {
	var errPrenom, errNom, errSexe, errNumLicence, errEquipe, errEpreuve, errAnnonce bool
	errPrenom = testCompetiteurPrenom()
	errNom = testCompetiteurNom()
	errSexe = testCompetiteurSexe()
	errNumLicence = testCompetiteurNumLicence()
	errEquipe = testCompetiteurEquipe()
	errEpreuve = testCompetiteurEpreuve()
	errAnnonce = testCompetiteurAnnonce()

	fmt.Println("\n\nSYNTHESE:")
	if errPrenom {
		fmt.Println("PRENOM:		ERROR")
	} else {
		fmt.Println("PRENOM:		OK")
	}
	if errNom {
		fmt.Println("NOM:		ERROR")
	} else {
		fmt.Println("NOM:		OK")
	}
	if errSexe {
		fmt.Println("SEXE:		ERROR")
	} else {
		fmt.Println("SEXE:		OK")
	}
	if errNumLicence {
		fmt.Println("LICENCE:	ERROR")
	} else {
		fmt.Println("LICENCE:	OK")
	}
	if errEquipe {
		fmt.Println("EQUIPE:		ERROR")
	} else {
		fmt.Println("EQUIPE:		OK")
	}
	if errEpreuve {
		fmt.Println("EPREUVE:	ERROR")
	} else {
		fmt.Println("EPREUVE:	OK")
	}
	if errAnnonce {
		fmt.Println("ANNONCE:	ERROR")
	} else {
		fmt.Println("ANNONCE:	OK")
	}
}

//TEST SUR LE PRENOM
func testCompetiteurPrenom() bool {
	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LE PRENOM: ")
	fmt.Println("TEST 1: Nom standard")
	test = comp.check()
	if !test {
		fmt.Println("TEST PRENOM 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST PRENOM 1:	OK")
	}

	fmt.Println("TEST 2: PRENOM AVEC ACCENT")
	comp.prenom = "Frédèrîqûê"
	test = comp.check()
	if !test {
		fmt.Println("TEST PRENOM 2:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST PRENOM 2:	OK")
	}

	fmt.Println("TEST 3: PRENOM AVEC TIRET")
	comp.prenom = "Jean-Eude-Saint-Jean"
	test = comp.check()
	if !test {
		fmt.Println("TEST PRENOM 3:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST PRENOM 3:	OK")
	}

	fmt.Println("TEST 4: PRENOM AVEC SYMBOLE")
	comp.prenom = "Jeanª╝É¾àzaea"
	test = comp.check()
	if !test {
		fmt.Println("TEST PRENOM 4:	OK")
		test = true
	} else {
		fmt.Println("TEST PRENOM 4:	ERROR")
		err = true
	}

	fmt.Println("TEST 5: PRENOM AVEC CHIFFRES")
	comp.prenom = "Bon123jour14"
	test = comp.check()
	if !test {
		fmt.Println("TEST PRENOM 5:	OK")
		test = true
	} else {
		fmt.Println("TEST PRENOM 5:	ERROR")
		err = true
	}
	return err
}

//TEST SUR LE NOM
func testCompetiteurNom() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LE NOM: ")
	fmt.Println("TEST 1: Nom standard")
	test = comp.check()
	if !test {
		fmt.Println("TEST NOM 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST NOM 1:	OK")
	}

	fmt.Println("TEST 2: NOM AVEC ACCENT")
	comp.prenom = "Frédèrîqûê"
	test = comp.check()
	if !test {
		fmt.Println("TEST NOM 2:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST NOM 2:	OK")
	}

	fmt.Println("TEST 3: NOM AVEC TIRET")
	comp.prenom = "Jean-Eude-Saint-Jean"
	test = comp.check()
	if !test {
		fmt.Println("TEST NOM 3:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST NOM 3:	OK")
	}

	fmt.Println("TEST 4: NOM AVEC SYMBOLE")
	comp.prenom = "Jeanª╝É¾àzaea"
	test = comp.check()
	if !test {
		fmt.Println("TEST NOM 4:	OK")
		test = true
	} else {
		fmt.Println("TEST NOM 4:	ERROR")
		err = true
	}

	fmt.Println("TEST 5: NOM AVEC CHIFFRES")
	comp.prenom = "Bon123jour14"
	test = comp.check()
	if !test {
		fmt.Println("TEST NOM 5:	OK")
		test = true
	} else {
		fmt.Println("TEST NOM 5:	ERROR")
		err = true
	}

	return err
}

//TEST SUR LE SEXE
func testCompetiteurSexe() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "123456", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LE SEXE: ")
	fmt.Println("TEST 1: SEXE HOMME")
	test = comp.check()
	if !test {
		fmt.Println("TEST SEXE 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST SEXE 1:	OK")
	}

	fmt.Println("TEST 2: SEXE FEMME")
	comp.sexe = "F"
	test = comp.check()
	if !test {
		fmt.Println("TEST SEXE 2:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST SEXE 2:	OK")
	}

	fmt.Println("TEST 3: SEXE PLUSIEURS CARACTERES")
	comp.sexe = "FFF"
	test = comp.check()
	if !test {
		fmt.Println("TEST SEXE 3:	OK")
		test = true

	} else {
		fmt.Println("TEST SEXE 3:	ERROR")
		err = true
	}

	fmt.Println("TEST 4: SEXE NON H/F")
	comp.sexe = "I"
	test = comp.check()
	if !test {
		fmt.Println("TEST SEXE 4:	OK")
		test = true
	} else {
		fmt.Println("TEST SEXE 4:	ERROR")
		err = true
	}

	fmt.Println("TEST 5: SEXE CHIFFRES")
	comp.sexe = "4"
	test = comp.check()
	if !test {
		fmt.Println("TEST SEXE 5:	OK")
		test = true
	} else {
		fmt.Println("TEST SEXE 5:	ERROR")
		err = true
	}
	return err
}

//TEST SUR LE NUMERO DE LICENCE
func testCompetiteurNumLicence() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "12-aBC-123", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LA LICENCE: ")
	fmt.Println("TEST 1: LICENCE CHIFFRE + LETTRES + \"-\"")
	test = comp.check()
	if !test {
		fmt.Println("TEST LICENCE 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST LICENCE 1:	OK")
	}

	fmt.Println("TEST 2: LICENCE ACCENTS")
	comp.numLicense = "ASD-123-1ÉÉéê"
	test = comp.check()
	if !test {
		fmt.Println("TEST LICENCE 2:	OK")
		test = true
	} else {
		fmt.Println("TEST LICENCE 2:	ERROR")
		err = true
	}

	fmt.Println("TEST 3: LICENCE SYMBOLES")
	comp.numLicense = "ABC-123-∟>¹1S"
	test = comp.check()
	if !test {
		fmt.Println("TEST LICENCE 3:	OK")
		test = true

	} else {
		fmt.Println("TEST LICENCE 3:	ERROR")
		err = true
	}
	return err
}

//TEST SUR LES NOM D'EQUIPE
func testCompetiteurEquipe() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "12-aBC-123", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LE NOM D'EQUIPE: ")
	fmt.Println("TEST 1: NOM STANDARD")
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 1:	OK")
	}

	fmt.Println("TEST 2: EQUIPE ACCENTS")
	comp.equipe = "équîpeNûmèrô"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 2:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 2:	OK")

	}

	fmt.Println("TEST 3: EQUIPE CHIFFRES")
	comp.equipe = "Loire42"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 3:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 3:	OK")

	}

	fmt.Println("TEST 4: EQUIPE ESPACES")
	comp.equipe = "Equipe du 42"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 4:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 4:	OK")

	}

	fmt.Println("TEST 5: EQUIPE TIRETS")
	comp.equipe = "Equipe-du-42"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 5:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 5:	OK")
	}

	fmt.Println("TEST 6: EQUIPE SYMBOLES")
	comp.equipe = "Equipe◙ãÍ╚┼"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 6:	OK")
		test = true
	} else {
		fmt.Println("TEST EQUIPE 6:	ERROR")
		err = true
	}
	return err
}

//TEST SUR LES EPREUVES
func testCompetiteurEpreuve() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "12-aBC-123", "equipe", "sta", 123, "sta", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR LES EPREUVES: ")
	fmt.Println("TEST 1: EPREUVE STA")
	test = comp.check()
	if !test {
		fmt.Println("TEST EPREUVE 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EPREUVE 1:	OK")
	}

	fmt.Println("TEST 2: EPREUVE SPD")
	comp.epreuve1 = "spd"
	comp.epreuve2 = "spd"
	test = comp.check()
	if !test {
		fmt.Println("TEST EPREUVE 2:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EPREUVE 2:	OK")

	}

	fmt.Println("TEST 3: EPREUVE DWF")
	comp.epreuve1 = "dwf"
	comp.epreuve2 = "dwf"
	test = comp.check()
	if !test {
		fmt.Println("TEST EPREUVE 3:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EPREUVE 3:	OK")

	}

	fmt.Println("TEST 4: EPREUVE DNF")
	comp.epreuve1 = "dnf"
	comp.epreuve2 = "dnf"
	test = comp.check()
	if !test {
		fmt.Println("TEST EPREUVE 4:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EPREUVE 4:	OK")
	}

	fmt.Println("TEST 5: EPREUVE 850")
	comp.epreuve1 = "850"
	comp.epreuve2 = "850"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 5:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST EQUIPE 5:	OK")
	}

	fmt.Println("TEST 6: EPREUVE AUTRE")
	comp.epreuve1 = "spd456"
	comp.epreuve2 = "123"
	test = comp.check()
	if !test {
		fmt.Println("TEST EQUIPE 6:	OK")
		test = true
	} else {
		fmt.Println("TEST EQUIPE 6:	ERROR")
		err = true
	}
	return err
}

//TEST SUR L'ANONCE
func testCompetiteurAnnonce() bool {

	var test bool
	var err bool
	comp := newCompetiteur(1, "Prenom", "Nom", "H", "12-aBC-123", "equipe", "sta", 123, "spd", 456)
	test = true
	err = false

	fmt.Println("CHECK SUR L'ANNONCE: ")
	fmt.Println("TEST 1: ANNONCE STANDARD")
	test = comp.check()
	if !test {
		fmt.Println("TEST LICENCE 1:	ERROR")
		test = true
		err = true
	} else {
		fmt.Println("TEST LICENCE 1:	OK")
	}

	fmt.Println("TEST 2: ANNONCE NEGATIVE")
	comp.annonce1 = -4561
	comp.annonce2 = -126
	test = comp.check()
	if !test {
		fmt.Println("TEST LICENCE 2:	OK")
		test = true
	} else {
		fmt.Println("TEST LICENCE 2:	ERROR")
		err = true
	}
	return err
}
