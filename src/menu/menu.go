package main

import "fmt"


func main() {
	
	var choix int
	choix=0
	
    
 
	for choix!= -1 {
	
	switch(choix){
	case(0): choix=menu()
	break
	case(1): choix=ajouterC()
	break
	case(2): choix=supprimerC()
	break
	case(3):choix=ajouterE()
	break
	case(4):choix=supprimerE()
	break
	case(9): choix= quitter()
	break
	default:fmt.Printf("Votre saisie n'est pas correcte \n")
	choix=0
	break
	}
	}
	
	
	
}
func menu()(int){
 fmt.Printf("\n==========MENU==========\n")
 fmt.Printf("1- Ajouter un compétiteur\n")
 fmt.Printf("2- Supprimer un compétiteur\n")
 fmt.Printf("3- Ajouter une équipe\n")
 fmt.Printf("4- Supprimer une équipe\n")
 fmt.Printf("9- Quitter\n")
 fmt.Printf("========FIN MENU========\n\n")
 fmt.Printf("Saisir votre Choix : ")
 var choix int
	fmt.Scan(&choix)
	return choix
}

func ajouterC()(int){
fmt.Printf("Ajouter un compétiteur\n")

fmt.Printf("Retour avec 0 :")
var retour int
	fmt.Scan(&retour)
	return retour

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
	 return -1} else {return 0}
	
}