package main 

import "flag"
import "fmt"
import "bufio"
import "os"



/*
* 		Menu.readString:
* Description: 
* 		Méthode permettant de vérifier si un mot comporte un caratère accentué
*/

func checkCaracteres() (motFinal string) {
	Ascii := convertToAscii()
	motFinal = ""
	i := 0
	for (i < len(Ascii)) {
		if (i < len(Ascii)-2 && Ascii[i] == 83 && Ascii[i+1] == 204 && Ascii[i+2] == 140){
			motFinal += "è"
			i += 2
		} else if (i < len(Ascii)-2 && Ascii[i] == 97 && Ascii[i+1] == 204 && Ascii[i+2] == 136) {
			motFinal += "õ"
			i += 2
		} else if (i < len(Ascii)-2 && Ascii[i] == 99 && Ascii[i+1] == 204 && Ascii[i+2] == 167) {
			motFinal += "þ"
			i += 2
		} else if (i < len(Ascii)-2 && Ascii[i] == 97 && Ascii[i+1] == 204 && Ascii[i+2] == 129) {
			motFinal += "ß"
			i += 2
		} else if (i < len(Ascii)-2 && Ascii[i] == 105 && Ascii[i+1] == 204 && Ascii[i+2] == 128) {
			motFinal += "ý"
			i += 2
		} else if (Ascii[i] == 48) {
			motFinal += "0"
		} else if (Ascii[i] == 49) {
			motFinal += "1"
		} else if (Ascii[i] == 50) {
			motFinal += "2"
		} else if (Ascii[i] == 51) {
			motFinal += "3"
		} else if (Ascii[i] == 52) {
			motFinal += "4"
		} else if (Ascii[i] == 53) {
			motFinal += "5"
		} else if (Ascii[i] == 54) {
			motFinal += "6"
		} else if (Ascii[i] == 55) {
			motFinal += "7"
		} else if (Ascii[i] == 56) {
			motFinal += "8"
		} else if (Ascii[i] == 57) {
			motFinal += "9"
		} else if ((Ascii[i] < 65) || (Ascii[i] > 90 && Ascii[i] < 97) || (Ascii[i] > 122)) {
			lettre, j := traduireCaractere(i, Ascii)
			i = j
			motFinal += lettre
		} else {
			motFinal += string(Ascii[i])
		}
		i += 1
	}
	return motFinal
}

/*
* 		Menu.traduireCaractere:
* Description: 
* 		Méthode permettant de reconnaitre un caractère accentué à partir de son code Ascii
*/

func traduireCaractere(i int, Ascii []byte) (lettre string, j int) {
	lettre = ""
	if (Ascii[i] == 226 && Ascii[i+1]  == 128 && Ascii[i+2] == 154) {
		lettre = "é"
		i += 2
	} else if (Ascii[i] == 194 && Ascii[i+1] == 160) {
		lettre = "á"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 166){
		lettre = "à"
		i += 2
	} else if (Ascii[i] == 198 && Ascii[i+1] == 146) {
		lettre = "â"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 158) {
		lettre = "ä"
		i += 2
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 160) {
		lettre = "å"
		i += 2
	} else if (Ascii[i] == 195 && Ascii[i+1] == 134) {
		lettre = "ã"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 152) {
		lettre = "æ"
		i += 2
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 161) {
		lettre = "ç"
		i += 2
	} else if (Ascii[i] == 203 && Ascii[i+1] == 134) {
		lettre = "ê"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 176) {
		lettre = "ë"
		i += 2
	} else if (Ascii[i] == 194 && Ascii[i+1] == 141) {
		lettre = "ì"
		i += 1
	} else if (Ascii[i] == 197 && Ascii[i+1] == 146) {
		lettre = "î"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 185) {
		lettre = "ï"
		i += 2
	} else if (Ascii[i] == 194 && Ascii[i+1] == 164) {
		lettre += "ñ"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 156) {
		lettre = "ô"
		i += 2
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 157) {
		lettre = "ö"
		i += 2
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 186) {
		lettre = "ø"
		i += 2
	} else if (Ascii[i] == 195 && Ascii[i+1] == 144) {
		lettre = "ð"
		i += 1
	} else if (Ascii[i] == 203 && Ascii[i+1] == 156) {
		lettre = "ÿ"
		i += 1
	} else if (Ascii[i] == 194 && Ascii[i+1] == 129) {
		lettre = "ü"
		i += 1
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 148) {
		lettre = "ù"
		i += 1
	} else if (Ascii[i] == 194 && Ascii[i+1] == 163) {
		lettre = "ú"
		i += 1
	} else if (Ascii[i] == 39) {
		lettre = "'"
	} else if (Ascii[i] == 226 && Ascii[i+1] == 128 && Ascii[i+2] == 147) {
		lettre = "û"
		i += 2
	} else if (Ascii[i] == 45) {
		lettre = "-"
	} else if (Ascii[i] == 95) {
		lettre = "_"
	} else if (Ascii[i] == 32) {
		lettre = " "
	}
	j = i
	return lettre, j
}

/*
* 		Menu.convertToAscii:
* Description: 
* 		Méthode permettant de convertir un caractère en son code Ascii
*/

func convertToAscii() (Ascii []byte) {
	equipe := read()
	Ascii = []byte(equipe)
	//fmt.Println(Ascii)
	return Ascii
}

/*
* 		Menu.read:
* Description: 
* 		Méthode permettant de lire un String entré par l'utilisateur à l'entrée standard même ci celui-ci contient des accents
*/

func read() (string) {
	scanner := bufio.NewScanner(os.Stdin)
	equipe := ""
	for scanner.Scan(){
		equipe = scanner.Text()
		if (equipe != ""){
			break
		}
	}
	return equipe
}


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
			base := newBdd("database/OpenApneeLyon")
			var id int
			var prenom string
			var nom string
			var sexe string
			var num_license string
			equipe := ""
			var epreuve1 string
			var annonce1 int
			var epreuve2 string
			var annonce2 int
			i := 2 
			for i!= 11 {
				if i ==2 {
					fmt.Println("Quel est le prénom du competiteur que vous souhaitez ajouter ? \n")
					prenom = checkCaracteres()
					i++ 
				} else if i == 3 {
					fmt.Println("Quel est le nom du competiteur que vous souhaitez ajouter ? \n")
					nom = checkCaracteres()
					i++ 
				} else if i == 4 {
					fmt.Println("Quel est le sexe du competiteur que vous souhaitez ajouter (H/F) ? \n")
					in , err := readString(1)
					for (err != nil || (in[0] != "H" && in[0] != "F")) {
						fmt.Println("Veuillez saisir H ou F svp \n")
						in, err = readString(1)
						continue
						fmt.Println(err)
					}
					sexe = in[0]
					i++ 
				} else if i == 5 {
					fmt.Println("Quel est le numéro de license du competiteur que vous souhaitez ajouter ? \n")
					in , err := readString(1)
					if (err != nil) {
						fmt.Println(err)
					}
					num_license = in[0]
					i++ 
				} else if i == 6 {
					fmt.Println("Quel est l'équipe du competiteur que vous souhaitez ajouter ? \n")
					equipe = checkCaracteres()
					i++ 
				} else if i == 7 {
					fmt.Println("Quel est la première épreuve à laquelle le competiteur que vous souhaitez ajouter va participer ? \n")
					fmt.Println(" sta ?\n spd ?\n dwf ?\n dnf ?\n 850 ?\n") 
					in , err := readString(1)
					for (err != nil || (in[0] != "sta" && in[0] != "spd" && in[0] != "dwf" && in[0] != "dnf" && in[0] != "850")) {
						fmt.Println("Veuillez saisir une des options suivantes svp : sta ; spd ; dwf ; dnf ; 850")
						in, err = readString(1)
						continue
						fmt.Println(err)
					}
					epreuve1 = in[0]
					i++ 
				} else if i == 8 {
					fmt.Println("Quel est la première annonce du competiteur que vous souhaitez ajouter ? (en mètre ou en seconde) \n")
					in , err := readInt(1)
					for (err != nil) {
						fmt.Println("Veuillez saisir un entier positif svp \n")
						in, err = readInt(1)
						continue
						fmt.Println(err)
					}
					annonce1 = in[0]
					i++ 
				} else if i == 9 {
					fmt.Println("Quel est la seconde épreuve à laquelle le competiteur que vous souhaitez ajouter va participer ? \n")
					fmt.Println(" sta ?\n spd ?\n dwf ?\n dnf ?\n 850 ?\n")
					in , err := readString(1)
					for (err != nil || (in[0] != "sta" && in[0] != "spd" && in[0] != "dwf" && in[0] != "dnf" && in[0] != "850")) {
						fmt.Println("Veuillez saisir une des options suivantes svp : sta ; spd ; dwf ; dnf ; 850")
						in, err = readString(1)
						continue
						fmt.Println(err)
					}
					epreuve2 = in[0]
					i++ 
				} else if i == 10 {
					fmt.Println("Quel est la seconde annonce du competiteur que vous souhaitez ajouter ? (en mètre ou en seconde) \n")
					in , err := readInt(1)
					for (err != nil) {
						fmt.Println("Veuillez saisir un entier positif svp \n")
						in, err= readInt(1)
						continue
						fmt.Println(err)
					}
					annonce2 = in[0]
					i++ 
				} else {
					fmt.Println("Les informations saisies ne sont pas conformes, veuillez recommencer \n")
				}
			}
			competiteur := newCompetiteur(id, prenom, nom, sexe, num_license, equipe, epreuve1, annonce1, epreuve2, annonce2)
			base.addCompetiteur(competiteur)
			
			
		} else if *c == "remove" {
			base := newBdd("database/OpenApneeLyon")
			var col_num int
			var value string
			fmt.Println("Sur quel critère faire la recherche ?")
			fmt.Println(" 1- Id \n 2- Equipe")
			in1, err1 := readInt(1)
			for (err1 != nil || in1[0] > 2) {
				fmt.Println("Veuillez saisir un entier positif inférieur ou égal à 2 svp \n")
				in1, err1 = readInt(1)
				continue
				fmt.Println(err1)
			}
			col_num = in1[0]
			fmt.Println("Saisissez l'objet de votre recherche")
			if (col_num == 1) {
				in2, err2 := readString(1)
				if (err2 != nil) {
					fmt.Println(err2)
				}
				value = in2[0]
			} else if (col_num == 2) {
				value = checkCaracteres()
			}
			base.deleteCompetiteur(col_num, value)
		
		} else if *c == "modify" {
			base := newBdd("database/OpenApneeLyon")
			var id_comp int
			var col_num int 
			fmt.Println("Saisissez l'id du participant que vous souhaitez modifier")
			in1, err1 := readInt(1)
			if (err1 != nil) {
				fmt.Println(err1)
			}
			id_comp = in1[0]
			fmt.Println("Quel est le critère que vous souhaitez modifier ?")
			fmt.Println(" 1- Prénom \n 2- Nom \n 3- Sexe \n 4- Numéro de license \n 5- Equipe \n 6- Première épreuve du participant")
			fmt.Println(" 7- Première annonce \n 8- Deuxième épreuve du participant \n 9- Deuxième annonce")
			in2, err2 := readInt(1)
			for (err2 != nil || in2[0] > 9) {
				fmt.Println("Veuillez saisir un entier positif inférieur ou égal à 9 svp \n")
				in2, err2 = readInt(1)
				continue
				fmt.Println(err2)
			}
			col_num = in2[0]+1
			fmt.Println("Quelle la nouvelle valeur de ce critère ?")

			if (col_num == 2) {
				prenom := checkCaracteres()
				base.modifCompetiteur (id_comp, col_num, prenom)
			} else if (col_num == 3) {
				nom := checkCaracteres()
				base.modifCompetiteur (id_comp, col_num, nom)
			} else if (col_num == 4) {
				in , err := readString(1)
				for (err != nil || (in[0] != "H" && in[0] != "F")) {
					fmt.Println("Veuillez saisir H ou F svp \n")
					in, err = readString(1)
					continue
					fmt.Println(err)
				}
				sexe := in[0]
				base.modifCompetiteur (id_comp, col_num, sexe)
			} else if (col_num == 5) {
				in , err := readString(1)
				if (err != nil) {
					fmt.Println(err)
				}
				num_license := in[0]
				base.modifCompetiteur (id_comp, col_num, num_license)
			} else if (col_num == 6) {
				equipe := checkCaracteres()
				base.modifCompetiteur (id_comp, col_num, equipe)
			} else if (col_num == 7) {
				in , err := readString(1)
				for (err != nil || (in[0] != "sta" && in[0] != "spd" && in[0] != "dwf" && in[0] != "dnf" && in[0] != "850")) {
					fmt.Println("Veuillez saisir une des options suivantes svp : sta ; spd ; dwf ; dnf ; 850")
					in, err = readString(1)
					continue
					fmt.Println(err)
				}
				epreuve1 := in[0]
				base.modifCompetiteur (id_comp, col_num, epreuve1)
			} else if (col_num == 8) {
				annonce1 := checkCaracteres()
				base.modifCompetiteur (id_comp, col_num, annonce1)
			} else if (col_num == 9) {
				in , err := readString(1)
				for (err != nil || (in[0] != "sta" && in[0] != "spd" && in[0] != "dwf" && in[0] != "dnf" && in[0] != "850")) {
					fmt.Println("Veuillez saisir une des options suivantes svp : sta ; spd ; dwf ; dnf ; 850")
					in, err = readString(1)
					continue
					fmt.Println(err)
				}
				epreuve2 := in[0]
				base.modifCompetiteur (id_comp, col_num, epreuve2)
			} else if (col_num == 10) {
				annonce2 := checkCaracteres()
				base.modifCompetiteur (id_comp, col_num, annonce2)
			}			
			
		} else if *c == "search" {
			base := newBdd("database/OpenApneeLyon")
			var col_num int
			var value string
			fmt.Println("Sur quel critère faire la recherche ?")
			fmt.Println(" 1- Id \n2- Prénom \n3- Nom \n4- Sexe \n5- Numéro de license \n6- Equipe \n7- Première épreuve du participant")
			fmt.Println(" 8- Première annonce \n9- Deuxième épreuve du participant \n10- Deuxième annonce")
			in1, err1 := readInt(1)
			for (err1 != nil) {
				fmt.Println("Veuillez saisir un entier positif svp \n")
				in1, err1 = readInt(1)
				continue
				fmt.Println(err1)
			}
			col_num = in1[0]
			fmt.Println("Saisissez l'objet de votre recherche")
			in2, err2 := readString(1)
			if (err2 != nil) {
				fmt.Println(err2)
			}
			value = in2[0]
			base.searchCompetiteur(col_num, value)
			
		} else if *c == "display" {
			base := newBdd("database/OpenApneeLyon")
			fmt.Println("Contenu actuel de la base de donnee : \n")
			base.displayCompetiteur()
			
		} else if *c == "import" {
			base := newBdd("database/OpenApneeLyon")
			base.importCompetiteur()
			fmt.Println("\n importation dans la base de donnee effectuee ! \n")
			
		} else if *c == "export" {
			base := newBdd("database/OpenApneeLyon")
			base.exportCompetiteur()
			fmt.Println("Fichier \"competiteurs\" exporté dans le dossier \"export\".")		
		} else if *c == "reset" {
			base := newBdd("database/OpenApneeLyon")
			base.resetCompetiteurs()
			fmt.Println("Tous les compétiteurs ont été effacés de la base de donnée")
		} else {
			fmt.Println("Vous pouvez ajouter un competiteur en tapant -c=add apres votre derniere commande\n")
			fmt.Println("Vous pouvez supprimer un competiteur en tapant -c=remove apres votre derniere commande\n")
			fmt.Println("Vous pouvez modifier un competiteur en tapant -c=modify apres votre derniere commande\n")
			fmt.Println("Vous pouvez rechercher un compétiteur en tapant -c=search apres votre derniere commande\n")
			fmt.Println("Vous pouvez afficher l'intégralité des compétiteurs en tapant -c=display apres votre derniere commande\n")
			fmt.Println("Vous pouvez importer un fichier csv contenant des compétiteurs en tapant -c=import apres votre derniere commande\n")
			fmt.Println("Vous pouvez exporter un fichier csv contenant des compétiteurs en tapant -c=export apres votre derniere commande\n")
			fmt.Println("Vous pouvez supprimer tous les compétiteurs de la base de donnée en tapant -c=reset après votre derniere commande\n")
		}
		
	} else if *e != "deff" {
	
		if *e == "check" {
			base := newBdd("database/OpenApneeLyon")
			base.check_team()
			fmt.Println("\n Verrification effectuee !\n")
			
		} else if *e == "planning" {
			plan := newPlanning("database/OpenApneeLyon")
			plan.getCompetiteur()
			plan.getConfigurationEpreuve()
			plan.generationPlanning()
			fmt.Println("Generation des planning effectuee !")
			
		
		} else {
			fmt.Println("Vous pouvez verifier la validite dune equipe en tapant -e=check apres votre derniere commande\n")
			fmt.Println("Vous pouvez générer le planning de la journée en tapant -e=planning apres votre derniere commande\n")
		}
		
	} else if *bdd != "deff" {
	
		if *bdd == "reset" {
				base := newBdd("database/OpenApneeLyon")
				fmt.Println("Confirmez ? (o/n)")
				in, err := readString(1)
				if (err != nil) {
					fmt.Println(err)
				}
				reponse := in[0]
				if reponse == "o" {
					base.reset()
				} else {
					Parsage()
				}
			} else if *bdd == "results" {
				base := newBdd("database/OpenApneeLyon")
				fmt.Println("Affichage de tous les résultats : \n")
				base.displayClassement()
			} else if *bdd == "display" {
				fmt.Println("Affichge du classement par equipe : ")
				base := newBdd("database/OpenApneeLyon")
				base.displayEquipe()
			} else {
				fmt.Println("Vous pouvez remettre la bdd a zero en tapant -bdd=reset apres votre derniere commande\n")
				fmt.Println("Vous pouvez afficher l'intégralité des résultats enregistrés en tapant -bdd=results apres votre derniere commande\n")
				fmt.Println("Vous pouvez afficher le classement par équipe en tapant -bdd=display apres votre derniere commande\n")
			}
			
	} else if *r != "deff" {
		if *r == "import" {
			base := newBdd("database/OpenApneeLyon")
			base.importResultat()
			fmt.Println("importation des compétiteurs contenus dans le fichier \"classement.csv\" dans le dossier import \n")
		} else if *r == "export" {
			base := newBdd("database/OpenApneeLyon")
			base.exportClassement()
			fmt.Println("Export du classement réalisé")
		} else if *r == "team" {
			base := newBdd("database/OpenApneeLyon")
			base.exportClassementEquipe()
			fmt.Println("Generation du classement final par equipe effectuee !")
		} else if *r == "reset" {
			base := newBdd("database/OpenApneeLyon")
			base.resetClassement()
		} else if *r == "display" {
				base := newBdd("database/OpenApneeLyon")
				base.displayEquipe()
		}else {
			fmt.Println("Vous pouvez importer les résultats des épreuves en tapant -r=import apres votre derniere commande\n")
			fmt.Println("Vous pouvez generer un classement individuel -r=export apres votre derniere commande\n")
			fmt.Println("Vous pouvez generer le classement final par équipe en tapant -r=team apres votre derniere commande\n")
			fmt.Println("Vous pouvez reinitialiser les classements en tapant -r=resetclass apres votre derniere commande\n")
		}
		
	} else {
		fmt.Println("Bienvenu dans l'aide ! \n")
		fmt.Println("Vous pouvez acceder aux options de gestion dun participant en tapant -c=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options de gestion dune equipe en tapant -e=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options dinteraction avec la bdd en tapant -bdd=help apres votre derniere commande\n")
		fmt.Println("Vous pouvez acceder aux options de gestion des resultats en tapant -r=help apres votre derniere commande\n")
	}	
}