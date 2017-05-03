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
	}
	
	func (boardE ClassementEquipe) display(){
		fmt.Println(strconv.Itoa(boardE.id) + "; " +
		boardE.equipe + "; " +
		strconv.Itoa(boardE.point))
	}
	
	func newClassementE(id int, equipe string, point int)(*ClassementEquipe){
		boardE := new(Classement)
		boardE.id = id
		boardE.equipe = equipe
		boardE.point = point
		
		return boardE
	}
	