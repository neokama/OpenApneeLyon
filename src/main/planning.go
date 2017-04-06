package main

import (
	"strconv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	)


type Planning struct
{
	base *Bdd
	comp []*Competiteur
}

func (p *Planning) get_comp(){
	p.base.resultat, p.base.err = p.base.db.Query("SELECT * FROM competiteurs")
		if p.base.err != nil {
			fmt.Println("Erreur lors de l'execution de la requête")
			log.Fatal(p.base.err)
		}
		defer p.base.resultat.Close()
		
		var info [10]string
		var temps1 float64
		var temps2 float64
		
		var nextcomp *Competiteur 
		
		for p.base.resultat.Next() {
			p.base.err = p.base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
			if p.base.err != nil {
			fmt.Println("Erreur lors de la récupération des résultats: \n")
			log.Fatal(p.base.err)
			}
		temps1,_ = strconv.ParseFloat(info[7], 64)
		temps2,_ = strconv.ParseFloat(info[9], 64)
		
		nextcomp = newcomp2(info[0], info[1], info[2], info[3], info[4], info[5], info[6], float32(temps1), info[8],float32(temps2))
		p.comp = append(p.comp, nextcomp)
		}
}



func (p Planning) disp_comp(){
	for j := 0; j < len(p.comp); j++ {
		p.comp[j].disp()
    }
}



func newPlanning(cheminBdd string)(*Planning){
	p := new(Planning)
	p.base = newBdd(cheminBdd)	
	
	return p
}