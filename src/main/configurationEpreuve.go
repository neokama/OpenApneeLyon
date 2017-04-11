package main

import (
	"fmt"
	"strconv"
	)
	
	type triEpreuves []*ConfigurationEpreuve

	func (triEp triEpreuves) Len() int           { return len(triEp) }
	func (triEp triEpreuves) Swap(i, j int)      { triEp[i], triEp[j] = triEp[j], triEp[i] }
	func (triEp triEpreuves) Less(i, j int) bool { return triEp[i].ordre < triEp[j].ordre }
	
	type ConfigurationEpreuve struct
	{
		ordre int
		id string
		seuilMin int
		seuilMax int
		nbPassages int
		marge int
		heureDebut string
	}
	
	func (ep ConfigurationEpreuve) display(){
		fmt.Println(strconv.Itoa(ep.ordre) + "; " +
		ep.id + "; " +
		strconv.Itoa(ep.seuilMin) + "; " +
		strconv.Itoa(ep.seuilMax) + "; " +
		strconv.Itoa(ep.nbPassages) + "; " +
		strconv.Itoa(ep.marge) + "; " +
		ep.heureDebut)
	}
	
	func newConfigurationEpreuve(ordre int, id string, seuilMin int, seuilMax int, nbPassages int, marge int, heureDebut string)(*ConfigurationEpreuve){
		ep := new(ConfigurationEpreuve)
		ep.ordre = ordre
		ep.id = id
		ep.seuilMin = seuilMin
		ep.seuilMax = seuilMax
		ep.nbPassages = nbPassages
		ep.marge = marge
		ep.heureDebut = heureDebut
		
		return ep
	}
	