package main
	

func main() {
	base := newBdd("../src/database/test","../src/database/test.csv")
	base.connection()
	base.disp_requete("SELECT * FROM competiteurs")
	base.addComp("Ninja","LeRigolo")
	base.disp_requete("SELECT * FROM competiteurs")
	base.delComp("nom","'Ninja'")
	base.disp_requete("SELECT * FROM competiteurs")
	base.requete_export("SELECT * FROM competiteurs")
	
	
}