package main

import (
	"fmt"
	)
	
	
	type personne struct
	{
		age int
		name string

	}
	
	func (pers personne) afficher(){
		fmt.Println("Bonjour, je suis ",pers.name," et j'ai ",pers.age, " ans\n")
	}
	
	func newpers(nom string, age int)(*personne){
		pers := new(personne)
		pers.name = nom
		pers.age = age
		
		return pers
	}
	