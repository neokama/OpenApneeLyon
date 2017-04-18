package main 

import "flag"
import "fmt"


/*
* 		Menu.readString:
* Description: 
* 		Méthode permettant de lire un String entré par l'utilisateur à l'entrée standard
*/

func readString(n int) ([]string, error) {
  in := make([]string, n)
  for i := range in {
    _, err := fmt.Scan(&in[i])
    if err != nil {
       return in[:i], err
    }
  }
  return in, nil
}

/*
* 		Menu.readfloat32:
* Description: 
* 		Méthode permettant de lire un float entré par l'utilisateur à l'entrée standard
*/


func readFloat(n int) ([]float32, error) {
  in := make([]float32, n)
  for i := range in {
    _, err := fmt.Scan(&in[i])
    if err != nil {
       return in[:i], err
    }
  }
  return in, nil
}

/*
* 		Menu.readfloat32:
* Description: 
* 		Méthode permettant de lire un entier entré par l'utilisateur à l'entrée standard
*/


func readInt(n int) ([]int, error) {
  in := make([]int, n)
  for i := range in {
    _, err := fmt.Scan(&in[i])
    if err != nil {
       return in[:i], err
    }
  }
  return in, nil
}

/*
* 		Menu.Parsage:
* Description: 
* 		Méthode permettant d'intéragir avec l'utilisateur pour que ce dernier puisse choisir la fonctionnalité qu'il souhaite effectuer
*/

func Parsage(){

	c := flag.String("c","deff", "a String") 	
	e := flag.String("e", "deff", "a String")
	bdd := flag.String("bdd", "deff", "a String")
	r := flag.String("r", "deff","a String")
	
	flag.Parse()  // permet d'initialiser le pointeur vers le flag.
	
	if *c != "deff" {
	
		if *c == "add"{
			base := newBdd("../src/database/OpenApneeLyon")
			var id int
			var prenom string
			var nom string
			var sexe string
			var num_license string
			var equipe string
			var epreuve1 string
			var temps1 int
			var epreuve2 string
			var temps2 int
			i := 1 
			for i!= 11 {
				if i == 1{
					fmt.Println("Quel est l'id du competiteur que vous souhaitez ajouter ? \n")
					in , err := readInt(1)
					fmt.Println(err)
					id = in[0]
					i++ 
				} else if i ==2 {
					fmt.Println("Quel est le prénom du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					fmt.Println(err)
					prenom = in[0]
					i++ 
				} else if i == 3 {
					fmt.Println("Quel est le nom du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					fmt.Println(err)
					nom = in[0]
					i++ 
				} else if i == 4 {
					fmt.Println("Quel est le sexe du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					fmt.Println(err)
					sexe = in[0]
					i++ 
				} else if i == 5 {
					fmt.Println("Quel est le numéro de license du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					fmt.Println(err)
					num_license = in[0]
					i++ 
				} else if i == 6 {
					fmt.Println("Quel est l'équipe du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					fmt.Println(err)
					equipe = in[0]
					i++ 
				} else if i == 7 {
					fmt.Println("Quel est la première épreuve à laquelle le competiteur que vous souhaitez ajouter va participer ? \n")
					in , err := readString(1)
					fmt.Println(err)
					epreuve1 = in[0]
					i++ 
				} else if i == 8 {
					fmt.Println("Quel est la première annonce du competiteur que vous souhaitez ajouter ? (en mètre ou en seconde) \n")
					in , err := readInt(1)
					fmt.Println(err)
					temps1 = in[0]
					i++ 
				} else if i == 9 {
					fmt.Println("Quel est la seconde épreuve à laquelle le competiteur que vous souhaitez ajouter va participer ? \n")
					in , err := readString(1)
					fmt.Println(err)
					epreuve2 = in[0]
					i++ 
				} else if i == 10 {
					fmt.Println("Quel est la seconde annonce du competiteur que vous souhaitez ajouter ? (en mètre ou en seconde) \n")
					in , err := readInt(1)
					fmt.Println(err)
					temps2 = in[0]
					i++ 
				} else {
					fmt.Println("Les informations saisies ne sont pas conformes, veuillez recommencer \n")
				}
			}
			competiteur := newCompetiteur(id, prenom, nom, sexe, num_license, equipe, epreuve1, temps1, epreuve2, temps2)
			base.addCompetiteur(competiteur)
			// fmt.Println(id, " | ", prenom, " | ", nom, " | ", sexe, " | ", num_license, " | ", equipe, " | ", epreuve1, " | ", temps1, " | ", epreuve2, " | ", temps2)
			
			
		} else if *c == "remove" {
			base := newBdd("../src/database/OpenApneeLyon")
			var col_num int
			var value string
			fmt.Println("Sur quel critère faire la recherche ?")
			fmt.Println("1- Id \n 2- Equipe")
			in1, err1 := readInt(1)
			fmt.Println(err1)
			col_num = in1[0]
			fmt.Println("Saisissez l'objet de votre recherche")
			in2, err2 := readString(1)
			fmt.Println(err2)
			value = in2[0]
			base.deleteCompetiteur(col_num, value)
		
		} else if *c == "modify" {
			base := newBdd("../src/database/OpenApneeLyon")
			var id_comp string
			var col_num int 
			var newvalue string
			fmt.Println("Saisissez l'id du participant que vous souhaitez modifier")
			in1, err1 := readString(1)
			fmt.Println(err1)
			id_comp = in1[0]
			fmt.Println("Quel est le critère que vous souhaitez modifier ?")
			fmt.Println("1- Id \n 2- Prénom \n 3- Nom \n 4- Sexe \n 5- Numéro de license \n 6- Equipe \n 7- Première épreuve du participant")
			fmt.Println("8- Première annonce \n 9- Deuxième épreuve du participant \n 10- Deuxième annonce")
			in2, err2 := readInt(1)
			fmt.Println(err2)
			col_num = in2[0]
			fmt.Println("Quelle la nouvelle valeur de ce critère ?")
			in3, err3 := readString(1)
			fmt.Println(err3)
			newvalue = in3[0]
			base.modifCompetiteur (id_comp, col_num, newvalue)
			
			
		} else if *c == "search" {
			base := newBdd("../src/database/OpenApneeLyon")
			var col_num int
			var value string
			fmt.Println("Sur quel critère faire la recherche ?")
			fmt.Println(" 1- Id \n2- Prénom \n3- Nom \n4- Sexe \n5- Numéro de license \n6- Equipe \n7- Première épreuve du participant")
			fmt.Println(" 8- Première annonce \n9- Deuxième épreuve du participant \n10- Deuxième annonce")
			in1, err1 := readInt(1)
			fmt.Println(err1)
			col_num = in1[0]
			fmt.Println("Saisissez l'objet de votre recherche")
			in2, err2 := readString(1)
			fmt.Println(err2)
			value = in2[0]
			base.searchCompetiteur(col_num, value)
			
		} else if *c == "display" {
			base := newBdd("../src/database/OpenApneeLyon")
			base.displayCompetiteur()
			
		} else if *c == "import" {
			base := newBdd("../src/database/OpenApneeLyon")
			base.importCompetiteur()
			
		} else if *c == "export" {
			base := newBdd("../src/database/OpenApneeLyon")
			base.exportCompetiteur()
			fmt.Println("Fichier \"competiteurs\" exporté dans le dossier \"export\".")		
		} else {
			fmt.Println("Vous pouvez ajouter un competiteur en tapant -c=add apres votre derniere commande\n")
			fmt.Println("Vous pouvez supprimer un competiteur en tapant -c=remove apres votre derniere commande\n")
			fmt.Println("Vous pouvez modifier un competiteur en tapant -c=modify apres votre derniere commande\n")
			fmt.Println("Vous pouvez rechercher un compétiteur en tapant -c=search apres votre derniere commande\n")
			fmt.Println("Vous pouvez afficher l'intégralité des compétiteurs en tapant -c=display apres votre derniere commande\n")
			fmt.Println("Vous pouvez importer un fichier csv contenant des compétiteurs en tapant -c=import apres votre derniere commande\n")
			fmt.Println("Vous pouvez exporter un fichier csv contenant des compétiteurs en tapant -c=export apres votre derniere commande\n")
		}
		
	} else if *e != "deff" {
	
		if *e == "check" {
			base := newBdd("../src/database/OpenApneeLyon")
			base.check_team()
			
		} else if *e == "planning" {
			plan := newPlanning("../src/database/OpenApneeLyon")
			fmt.Println("Saisissez le chemin et le nom du fichier de configuration des épreuves")
			in, err := readString(1)
			fmt.Println(err)
			chemin1 := in[0]
			plan.getCompetiteur()
			plan.getConfigurationEpreuve(chemin1)
			fmt.Println("Saisissez le chemin vers lequel vous souhaitez enregistrer votre planning")
			in, err = readString(1)
			chemin2 := in[0]
			fmt.Println("Saisissez le nom du fichier")
			in, err = readString(1)
			nom := in[0]
			cheminComplet := fmt.Sprint(chemin2, "/", nom)
			plan.generationPlanning(cheminComplet)
		
		} else {
			fmt.Println("Vous pouvez verifier la validite dune equipe en tapant -e=check apres votre derniere commande\n")
			fmt.Println("Vous pouvez générer le planning de la journée en tapant -e=planning apres votre derniere commande\n")
		}
		
	} else if *bdd != "deff" {
	
		if *bdd == "reset" {
				base := newBdd("../src/database/OpenApneeLyon")
				fmt.Println("Confirmez ? (o/n)")
				in, err := readString(1)
				fmt.Println(err)
				reponse := in[0]
				if reponse == "o" {
					base.reset()
				} else {
					Parsage()
				}
				
			} else if *bdd == "save" {
				fmt.Println("Copie les fichiers csv contenant les compétiteurs, les scores ainsi que le planning avec timestamp, les place dans un dossier save afin d’avoir des sauvegardes.")
			} else {
				fmt.Println("Vous pouvez remettre la bdd a zero en tapant -bdd=reset apres votre derniere commande\n")
				fmt.Println("Vous pouvez creer des sauvegardes en tapant -bdd=save apres votre derniere commande\n")
			}
			
	} else if *r != "deff" {
		if *r == "add" {
			fmt.Println("Demande le nom et la performance du compétiteur ainsi que le nom de l’épreuve sur laquelle on veut ajouter des résultats afin de l’ajouter au tableau des scores")
		} else if *r == "import" {
			fmt.Println("Importe un fichier .csv de résultats d’une épreuve")
		} else if *r == "modify" {
			fmt.Println("Modifie le score d’un compétiteur")
		} else if *r == "print" {
			fmt.Println("Génère une feuille de résultats d’une épreuve lorsqu’elle est terminée")
		} else if *r == "generateClass" {
			fmt.Println("Lance la génération du classement final une fois que les épreuves sont terminées")
		} else if *r == "remove" {
			fmt.Println("Supprime le score d’un participant à une épreuve")
		} else {
			fmt.Println("Vous pouvez inscrire la performance dun competiteur en tapant -r=add apres votre derniere commande\n")
			fmt.Println("Vous pouvez importer un fichier csv de resultats en tapant -r=import apres votre derniere commande\n")
			fmt.Println("Vous pouvez modifier le score dun competiteur en tapant -r=modify apres votre derniere commande\n")
			fmt.Println("Vous pouvez generer une feuille de resultat pour une epreuve donnee en tapant -r=print apres votre derniere commande\n")
			fmt.Println("Vous pouvez generer le classement final en tapant -r=generateClass apres votre derniere commande\n")
			fmt.Println("Vous pouvez supprimer le score dun participant a une epreuve en tapant -r=remove apres votre derniere commande\n")
		}
		
	} else {
		fmt.Println("Bienvenu dans l'aide ! \n")
		fmt.Println("Vous pouvez acceder aux options de gestion dun participant en tapant -c=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options de gestion dune equipe en tapant -e=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options dinteraction avec la bdd en tapant -bdd=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options de gestion des resultats en tapant -r=help apres votre derniere commande\n")
	}	
}

/* func main(){
	Parsage()
} */