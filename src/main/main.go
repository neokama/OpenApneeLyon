package main

import (
	"fmt"
	)
	
	
func main() {
	fmt.Println("Début: \n")
	
	//%%%%%%%%% Compétiteurs %%%%%%%%%
	Moi := newcomp("RICAUD","Arnaud")
	Moi.id = "ARI1"
	Moi.num_license = "23111995N1"
	Moi.equipe = "TeamNono"
	Moi.epreuve1 ="Stat"
	Moi.temps1 = 150
	Moi.epreuve2 ="16x50"
	Moi.temps2 = 1250
	Moi.disp()
	
	
	/*
	%%%%%%%%% Bdd %%%%%%%%%
	base := newBdd("../src/database/OpenApneeLyon")
	base.reset()
	fmt.Println("\n")
	base.addComp(Moi)
	fmt.Println("\n")
	base.disp_comp()
	base.export_comp("","pourquoipas")
	base.import_comp("../bin/import.csv")
	fmt.Println("\n")
	base.disp_comp()
	fmt.Println("\n")
	base.search_comp(3, "RICAUD")
	base.modif_comp("ARI1", 2, "nouveau-prenom")
	fmt.Println("\n")
	base.disp_comp()
	fmt.Println("\n")
	
	base.delComp(3, "RICAUD")
	fmt.Println("\n")
	*/
	
	//%%%%%%%%% Bdd %%%%%%%%%
	base := newBdd("../src/database/OpenApneeLyon")
	base.reset()
	fmt.Println("\n")
	base.addComp(Moi)
	base.addComp(Moi)
	base.addComp(Moi)
	base.addComp(Moi)
	fmt.Println("\n")
	p := newPlanning("../src/database/OpenApneeLyon")
	p.get_comp()
	p.disp_comp()
	
}