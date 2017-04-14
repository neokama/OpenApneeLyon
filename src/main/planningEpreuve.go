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
		sexe string
		equipe string
		annonce int
		seuilMin int
		seuilMax int
		numSerie int
		numPassage int
		heurePassage string
	}
	
	func (planep PlanningEpreuve) display(){
		fmt.Println(planep.idEpreuve + "; " +
		planep.idComp + "; " +
		planep.prenom + "; " +
		planep.nom + "; " +
		planep.sexe + "; " +
		planep.equipe + "; " +
		strconv.Itoa(planep.annonce) + "; " +
		strconv.Itoa(planep.numSerie) + "; " +
		strconv.Itoa(planep.numPassage) + "; " +
		planep.heurePassage)
	}
	
	func newPlanningEpreuve(idEpreuve string, idComp string, prenom string, nom string, sexe string, equipe string, annonce int)(*PlanningEpreuve){
		planep := new(PlanningEpreuve)
		planep.idEpreuve = idEpreuve
		planep.idComp = idComp
		planep.prenom = prenom
		planep.nom = nom
		planep.annonce = annonce
		planep.seuilMin = 0
		planep.seuilMax = 0
		planep.sexe = sexe
		planep.equipe = equipe
		planep.numSerie = 0
		planep.numPassage = 0
		planep.heurePassage = "00:00"
		
		return planep
	}
	