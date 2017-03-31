/* package main 

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

} */


package main 

import "flag"
import "fmt"


func Parsage(){
	c := flag.String("c","deff", "a String") // On crée ici un flag => lorsqu'on lance le programme on peut ajouter une option 
	// pour choisir quelle fonctionnalité utiliser. Dans ce cas, le choix par défaut pour la valeur du flag est help et "a String" n'est 
	// qu'une description du flag.
	
	e := flag.String("e", "deff", "a String")
	bdd := flag.String("bdd", "deff", "a String")
	r := flag.String("r", "deff","a String")
	
	flag.Parse()  // permet d'initialiser le pointeur vers le flag.
	
	if *c == "add"{
		fmt.Println("le programme demande alors les nom, prénom, numéro, équipe etc… du compétiteur à ajouter.")
	} else if *c == "rmv" {
		fmt.Println("Le programme demande alors le nom et prénom du compétiteur à supprimer et demande confirmation après avoir affiché les données du compétiteur.")
	} else if *c == "modify" {
		fmt.Println("Le programme demande nom et prénom du compétiteur, puis la/les colonne(s) à modifier.")
	} else if *c == "search" {
		fmt.Println(" Le programme demande par quel élément la recherche se fait (nom, prénom, numéro…) puis le/les mot(s)-clé(s) de la recherche.")
	} else if *c == "genbad" {
		fmt.Println("Sélectionne un fichier csv (compétiteurs, bénévoles…) et édite les badges correspondants sous format pdf.")
	} else if *c == "help" {
		fmt.Println("Vous pouvez ajouter un competiteur en tapant -c=add apres votre derniere commande\n")
		fmt.Println("Vous pouvez supprimer un competiteur en tapant -c=rmv apres votre derniere commande\n")
		fmt.Println("Vous pouvez modifier un competiteur en tapant -c=modify apres votre derniere commande\n")
		fmt.Println("Vous pouvez rechercher un compétiteur en tapant -c=search apres votre derniere commande\n")
		fmt.Println("Vous pouvez generer des badges en tapant -c=genbad apres votre derniere commande\n")
	} else {
	
		if *e == "add" {
			fmt.Println("Ajout d'une equipe a la bdd")
		} else if *e == "rmv" {
			fmt.Println("suppression d'une equipe de la bdd")
		} else if *e == "modify" {
			fmt.Println("Modification d'une equipe de la bdd")
		} else if *e == "search" {
			fmt.Println("Recherche d'une equipe dans la bdd")
		} else if *e == "check" {
			fmt.Println("Permet de demander la verification de la validite d’une equipe")
		} else if *e == "help" {
			fmt.Println("Vous pouvez ajouter une equipe en tapant -e=add apres votre derniere commande\n")
			fmt.Println("Vous pouvez supprimer une equipe en tapant -e=rmv apres votre derniere commande\n")
			fmt.Println("Vous pouvez modifier une equipe en tapant -e=modify apres votre derniere commande\n")
			fmt.Println("Vous pouvez rechercher une equipe en tapant -e=search apres votre derniere commande\n")
			fmt.Println("Vous pouvez verifier la validite dune equipe en tapant -e=check apres votre derniere commande\n")
		} else {
		
			if *bdd == "reset" {
				fmt.Println("Remise a zero de la bdd")
			} else if *bdd == "save" {
				fmt.Println("Copie les fichiers csv contenant les compétiteurs, les scores ainsi que le planning avec timestamp, les place dans un dossier save afin d’avoir des sauvegardes.")
			} else if *bdd == "import" {
				fmt.Println("Demande le chemin d’accès au fichier csv contenant les compétiteurs")
			} else if *bdd == "genpla" {
				fmt.Println("Lance la vérification puis la génération du planning, ainsi que des fiches épreuves")
			} else if *bdd == "help" {
				fmt.Println("Vous pouvez remettre la bdd a zero en tapant -bdd=reset apres votre derniere commande\n")
				fmt.Println("Vous pouvez creer des sauvegardes en tapant -bdd=save apres votre derniere commande\n")
				fmt.Println("Vous pouvez importer les donnees dun fichier csv en tapant -bdd=import apres votre derniere commande\n")
				fmt.Println("Vous pouvez generer des plannings en tapant -bdd=genpla apres votre derniere commande\n")
			} else {
			
				if *r == "add" {
					fmt.Println("Demande le nom et la performance du compétiteur ainsi que le nom de l’épreuve sur laquelle on veut ajouter des résultats afin de l’ajouter au tableau des scores")
				} else if *r == "import" {
					fmt.Println("Importe un fichier .csv de résultats d’une épreuve")
				} else if *r == "modify" {
					fmt.Println("Modifie le score d’un compétiteur")
				} else if *r == "print" {
					fmt.Println("Génère une feuille de résultats d’une épreuve lorsqu’elle est terminée")
				} else if *r == "genclass" {
					fmt.Println("Lance la génération du classement final une fois que les épreuves sont terminées")
				} else if *r == "rmv" {
					fmt.Println("Supprime le score d’un participant à une épreuve")
				} else if *r == "help" {
					fmt.Println("Vous pouvez inscrire la performance dun competiteur en tapant -r=add apres votre derniere commande\n")
					fmt.Println("Vous pouvez importer un fichier csv de resultats en tapant -r=import apres votre derniere commande\n")
					fmt.Println("Vous pouvez modifier le score dun competiteur en tapant -r=modify apres votre derniere commande\n")
					fmt.Println("Vous pouvez generer une feuille de resultat pour une epreuve donnee en tapant -r=print apres votre derniere commande\n")
					fmt.Println("Vous pouvez generer le classement final en tapant -r=genclass apres votre derniere commande\n")
					fmt.Println("Vous pouvez supprimer le score dun participant a une epreuve en tapant -r=rmv apres votre derniere commande\n")
				} else {
					fmt.Println("Bienvenu dans l'aide ! \n")
					fmt.Println("Vous pouvez acceder aux options de gestion dun participant en tapant -c=help apres votre derniere commande\n")
					fmt.Println("Vous pouvez acceder aux options de gestion dune equipe en tapant -e=help apres votre derniere commande\n")
					fmt.Println("Vous pouvez acceder aux options dinteraction avec la bdd en tapant -bdd=help apres votre derniere commande\n")
					fmt.Println("Vous pouvez acceder aux options de gestion des resultats en tapant -r=help apres votre derniere commande\n")
				}
			}
		}
	}
	
	
	
}

func main() {
	Parsage()
}

// L'idée c'est donc de créer un flag qui selon sa valeur exécute la fonctionnalité recherchée :
// exemple : si le flag "choix" prend la valeur reg alors le programme exécute la méthode registration qui permet d'importer 
// les participants dans la base de données.