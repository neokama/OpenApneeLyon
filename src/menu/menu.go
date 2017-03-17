package main

import "fmt"
import "os"
import"io"
import "bufio"
import "log"
import "io/ioutil"

func main() {
	var choix int
 choix=0
 
	menu(choix)
	}
	
func menu(choix int){
 
	
    //initialisationFichier()
 
	
	for choix!= -1 {
	fmt.Printf("\n==========MENU==========\n")
 fmt.Printf("1- Gestion des compétiteurs\n")
 fmt.Printf("2- Gestion des équipes\n")
 fmt.Printf("9- Quitter\n")
 fmt.Printf("========FIN MENU========\n\n")
 fmt.Printf("Saisir votre Choix : ")
 
	fmt.Scan(&choix)
	
	switch(choix){
	
	case(1): choix=ajouterC()
	break
	case(2): choix=supprimerC()
	break
	case(3):choix=ajouterE()
	break
	case(4):choix=supprimerE()
	break
	case(9): choix=quitter()
	break
	default:fmt.Printf("Votre saisie n'est pas correcte ! \n")	
	break
	}
}
}

func sousMenuAjouterC()(int){
fmt.Printf("\n===GESTION DES COMPETITEURS===\n")
 fmt.Printf("1- Lire Fichier compétiteur\n")
 fmt.Printf("2- ajouter un compétiteur\n")
 fmt.Printf("3- Rechercher un compétiteur\n")
 fmt.Printf("4- Supprimer un compétiteur\n")
 fmt.Printf("0- Retour\n")
 fmt.Printf("=========FIN AJOUTER=========\n\n")
 fmt.Printf("Saisir votre Choix : ")

var retour int
	fmt.Scan(&retour)
	return retour
	}
	
func ajouterC()(int){
	var choix int
	choix=9
	var rep string
	rep="a"
	
	for choix!= 0 {
	
	switch(choix){
	case(0): 
	break
	case(1): lireFichier()
	for rep!="o"{
	fmt.Printf("Pour retourner au sous-menu ? (o) ")
		fmt.Scan(&rep)
		} 
		choix=9
	break
	case(2): fmt.Printf("2")
	break
	case(3):fmt.Printf("3")
	break
	case(9):choix=sousMenuAjouterC()
	break
	default:fmt.Printf("Votre saisie n'est pas correcte \n")
	choix=9
	break
	}
}
return choix
}
		
func supprimerC()(int){
fmt.Printf("Supprimer un compétiteur\n")
fmt.Printf("Retour avec 0 :")
var retour int
	fmt.Scan(&retour)
	return retour

}

func ajouterE()(int){
fmt.Printf("Ajouter une équipe\n")
fmt.Printf("Retour avec 0 :")
var retour int
	fmt.Scan(&retour)
	return retour
	
	}
	func supprimerE()(int){
fmt.Printf("Supprimer une équipe\n")
fmt.Printf("Retour avec 0 :")
var retour int
	fmt.Scan(&retour)
	return retour
	
	}

func quitter()(int){

var res int 
res =0
for res==0 {

fmt.Printf("Voulez-vous vraiment quitter l'application ? (o/n) \n")
var choixQ string
	fmt.Scan(&choixQ)
	 
	if choixQ=="o"{
	fmt.Printf("Vous quittez l'appli ! Au revoir \n")
	res=1
	} else if choixQ=="n"{
		//fmt.Printf("Vous revenez au menu\n")
		res=2
		} else  {
	fmt.Printf("Votre saise n'est pas correcte \n")
	res=0
	}
	}
	
	 if res ==1 {
	 return -1 } else if res==2 {return 0} else {return 0}
	
}


func initialisationFichier(){
creationCompetiteur() 

}

func creationCompetiteur(){

	//Création d'un fichier CSV
	file, err := os.Create("Compétiteur.csv")
	
	//Vérification erreur
	if err != nil {
	log.Fatal(err)
	}
	
	//Ecriture dans le fichier CSV
	file.WriteString("Nom;Prenom;Sexe;No Licence\n")
	
	

}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func lireFichier(){
//Lecture du fichier CSV
dat, err := ioutil.ReadFile("Compétiteur.csv")
    check(err)
    fmt.Print(string(dat))

    // You'll often want more control over how and what
    // parts of a file are read. For these tasks, start
    // by `Open`ing a file to obtain an `os.File` value.
    f, err := os.Open("Compétiteur.csv")
    check(err)

    // Read some bytes from the beginning of the file.
    // Allow up to 5 to be read but also note how many
    // actually were read.
    b1 := make([]byte, 10)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // You can also `Seek` to a known location in the file
    // and `Read` from there.
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // The `io` package provides some functions that may
    // be helpful for file reading. For example, reads
    // like the ones above can be more robustly
    // implemented with `ReadAtLeast`.
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // There is no built-in rewind, but `Seek(0, 0)`
    // accomplishes this.
    _, err = f.Seek(0, 0)
    check(err)

    // The `bufio` package implements a buffered
    // reader that may be useful both for its efficiency
    // with many small reads and because of the additional
    // reading methods it provides.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    
    f.Close()

}