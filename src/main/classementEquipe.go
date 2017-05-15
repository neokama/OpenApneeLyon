package main
	
	import (
	"strconv"
	"fmt"
	)
	
	type ClassementEquipe struct 
	{
	id int
	equipe string
	point int 
	place int
	}
	
	func (boardE ClassementEquipe) displayEquipe(){
		fmt.Println(strconv.Itoa(boardE.id) + "; " +
		boardE.equipe + "; " +
		strconv.Itoa(boardE.point)+ "; " +
		strconv.Itoa(boardE.place))
	}
	
	func newClassementE(id int, equipe string, point int, place int)(*ClassementEquipe){
		boardE := new(ClassementEquipe)
		boardE.id = id
		boardE.equipe = equipe
		boardE.point = point
		boardE.place = place
		
		return boardE
	}
	