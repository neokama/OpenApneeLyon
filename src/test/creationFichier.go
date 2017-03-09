package main

import "fmt"
import "log"
import "os"
import "io/ioutil"


func main() {

	
	//FILE I/O
	
	//
	// Ecriture dans un fichier
	//
	
	//
	//Création d'un fichier CSV
	file, err := os.Create("fichierTest.csv")
	
	//Vérification erreur
	if err != nil {
	log.Fatal(err)
	}
	
	//Ecriture dans le fichier CSV
	file.WriteString("Test;Leo;est;stupide\net; aurelien;est;intelligent")
	//file.close()
	
	
	//
	//Création d'un fichier TXT
	file2, err := os.Create("fichierTest.txt")
	
	//Vérification erreur
	if err != nil {
	log.Fatal(err)
	}
	
	//Ecriture dans le fichier TXT
	file2.WriteString("J'écris dans mon fichier Txt. Et là j'écris une deuxième phrase ! :) ")
	//file.close()
	
	//
	// Lecture dans un fichier TXT
	// 
	stream, err := ioutil.ReadFile("fichierTest.txt")
	if err != nil {
	log.Fatal(err)
	}
	
	readString := string(stream)
	fmt.Println(readString)
	
	//
	// Lecture dans un fichier CSV
	// 
	stream2, err := ioutil.ReadFile("fichierTest.csv")
	if err != nil {
	log.Fatal(err)
	}
	
	readString2 := string(stream2)
	fmt.Println(readString2)
	
	
	
}
