package main

import (
	"fmt"
	"strconv"

	)
	
	type triAnnonces []*PlanningEpreuve

	func (trieur triAnnonces) Len() int           { return len(trieur) }
	func (trieur triAnnonces) Swap(i, j int)      { trieur[i], trieur[j] = trieur[j], trieur[i] }
	func (trieur triAnnonces) Less(i, j int) bool { return trieur[i].annonce > trieur[j].annonce }

	type PlanningEpreuve struct
	{
		idEpreuve string
		idComp string
		prenom string
		nom string
		annonce int
		numSerie int
		numPassage int
		heurePassage string
	}
	
	func (planep PlanningEpreuve) display(){
		fmt.Println(planep.idEpreuve + "; " +
		planep.idComp + "; " +
		planep.prenom + "; " +
		planep.nom + "; " +
		strconv.Itoa(planep.annonce) + "; " +
		planep.heurePassage)
	}
	
	func newPlanningEpreuve(idEpreuve string, idComp string, prenom string, nom string, annonce int, heurePassage string)(*PlanningEpreuve){
		planep := new(PlanningEpreuve)
		planep.idEpreuve = idEpreuve
		planep.idComp = idComp
		planep.prenom = prenom
		planep.nom = nom
		planep.annonce = annonce
		planep.heurePassage = heurePassage
		
		return planep
	}
	