package main

import (
	"fmt"
	"strconv"
	)
	
	type triEpreuves []*ConfigurationEpreuve

	func (triEp triEpreuves) Len() int           { return len(triEp) }
	func (triEp triEpreuves) Swap(i, j int)      { triEp[i], triEp[j] = triEp[j], triEp[i] }
	func (triEp triEpreuves) Less(i, j int) bool { return triEp[i].ordre < triEp[j].ordre }

	// ConfigurationEpreuve structure
	type ConfigurationEpreuve struct
	{
		ordre int
		id string
		seuilMin int
		seuilMax int
		nbParPassage int
		nbParticipants int
		dureeEchauffement int
		dureePassage int
		dureeAppel int
		surveillance int
		battementSerie int
		battementEpreuve int
		heureDebut string
	}
	
	func (ep ConfigurationEpreuve) display(){
		fmt.Println(strconv.Itoa(ep.ordre) + ";" +
		ep.id + ";" +
		strconv.Itoa(ep.seuilMin) + ";" +
		strconv.Itoa(ep.seuilMax) + ";" +
		strconv.Itoa(ep.nbParPassage) + ";" +
		strconv.Itoa(ep.nbParticipants) + ";" +
		strconv.Itoa(ep.dureeEchauffement) + ";" +
		strconv.Itoa(ep.dureeAppel) + ";" +
		strconv.Itoa(ep.dureePassage) + ";" +
		strconv.Itoa(ep.surveillance) + ";" +
		strconv.Itoa(ep.battementSerie) + ";" +
		strconv.Itoa(ep.battementEpreuve) + ";" +
		ep.heureDebut)
	}
	
	func newConfigurationEpreuve(ordre int, id string, seuilMin int, seuilMax int, nbParPassage int,dureeEchauffement int,
		dureePassage int ,dureeAppel int, surveillance int, battementSerie int, battementEpreuve int, heureDebut string)(*ConfigurationEpreuve){
		ep := new(ConfigurationEpreuve)
		ep.ordre = ordre
		ep.id = id
		ep.seuilMin = seuilMin
		ep.seuilMax = seuilMax
		ep.nbParPassage = nbParPassage
		ep.dureePassage = dureePassage
		ep.dureeEchauffement = dureeEchauffement
		ep.dureeAppel = dureeAppel
		ep.surveillance = surveillance
		ep.battementSerie = battementSerie
		ep.battementEpreuve = battementEpreuve
		ep.heureDebut = heureDebut
		
		return ep
	}
	