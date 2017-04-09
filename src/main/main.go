package main

import (
	"fmt"
	)
	
	
func main() {
	fmt.Println("Début: \n")

	// %%%%%%%%% Compétiteurs %%%%%%%%%
	Moi := newCompetiteur("ARI1","Arnaud","RICAUD","M","23111995N1","TeamNono","Stat",150,"16x50",1250)
	Moi.display()
	
	
	
	//%%%%%%%%% Bdd %%%%%%%%%
	// base := newBdd("../src/database/OpenApneeLyon")
	// base.reset()
	// fmt.Println("\n")
	// base.addCompetiteur(Moi)
	// fmt.Println("\n")
	// base.displayCompetiteur()
	// base.exportCompetiteur("","pourquoipas")
	// base.importCompetiteur("../bin/import.csv")
	// fmt.Println("\n")
	// base.displayCompetiteur()
	// fmt.Println("\n")
	// base.searchCompetiteur(3, "RICAUD")
	// base.modifCompetiteur("ARI1", 2, "nouveau-prenom")
	// fmt.Println("\n")
	// base.displayCompetiteur()
	// fmt.Println("\n")
	
	// base.deleteCompetiteur(3, "RICAUD")
	// fmt.Println("\n")
	
	
	//%%%%%%%%% Planning %%%%%%%%%
	base := newBdd("../src/database/OpenApneeLyon")
	base.reset()
	fmt.Println("\n")
	base.importCompetiteur("../bin/import.csv")
	base.exportCompetiteur("","pourquoipas")
	fmt.Println("\n")
	p := newPlanning("../src/database/OpenApneeLyon")
	p.getCompetiteur()
	p.displayCompetiteur()
	fmt.Println("\nConfiguration des épreuves")
	p.getConfigurationEpreuve()
	p.displayConfigurationEpreuve()
	fmt.Println("\n")
	p.generationPlanning()

}