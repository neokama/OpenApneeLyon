package main


import "log"
import "os"

func main() {

	
	//FILE I/O
	
	//Création d'un fichier CSV
	file, err := os.Create("fichierTest.csv")
	
	//Vérification erreur
	if err != nil {
	log.Fatal(err)
	}
	
	//Ecriture dans le fichier CSV
	file.WriteString("Test;Leo;est;stupide\net; aurelien;est;intelligent")
	//file.close()
	
	//Création d'un fichier TXT
	file2, err := os.Create("fichierTest.txt")
	
	//Vérification erreur
	if err != nil {
	log.Fatal(err)
	}
	
	//Ecriture dans le fichier TXT
	file2.WriteString("J'écris dans mon fichier Txt. Et là j'écris une deuxième phrase ! :) ")
	//file.close()
}
