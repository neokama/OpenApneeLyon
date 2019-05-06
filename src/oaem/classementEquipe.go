package main
	
	import (
	"strconv"
	"fmt"
	)
	
	/*
	* 		ClassementEquipe
	* Description:
	* La structure permet de gérer un classement equipe.
	*
	* Paramètres: 
	*	- id : numéro d'identification de l'équipe
	*	- equipe : nom de l'équipe
	*	- point : nombre de point de l'équipe
	*	- place : position de l'équipe
	**/
	type ClassementEquipe struct 
	{
	id int
	equipe string
	point int 
	place int
	etat bool
	}
	
	/*
	* 		ClassementEquipe.displayEquipe1:
	* Description:
	* 			La méthode permet d'afficher une structure équipe.
	
	*/
	func (boardE ClassementEquipe) displayEquipe1(){
		fmt.Println(strconv.Itoa(boardE.id) + "," +
		boardE.equipe + "," +
		strconv.Itoa(boardE.point)+ "," +
		strconv.Itoa(boardE.place)+ "," +
		strconv.FormatBool(boardE.etat))
	}
	
	
	/*
	* 		newClassementE:
	* Description:
	* 			La méthode permet de retrourner un ClassementEquipe
	* Paramètres:
	*	- id : numéro d'identification de l'équipe
	*	- equipe : nom de l'équipe
	*	- point : nombre de point de l'équipe
	*	- place : position de l'équipe
	*/
	func newClassementE(id int, equipe string, point int, place int, etat bool)(*ClassementEquipe){
		boardE := new(ClassementEquipe)
		boardE.id = id
		boardE.equipe = equipe
		boardE.point = point
		boardE.place = place
		boardE.etat = etat
		
		return boardE
	}
	