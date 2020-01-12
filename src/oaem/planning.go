package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Planning structure
type Planning struct {
	base         *Bdd
	comp         []*Competiteur
	cfgEpreuves  []*ConfigurationEpreuve
	planEpreuves []*PlanningEpreuve
}

func (p *Planning) getCompetiteur() {
	p.base.resultat, p.base.err = p.base.db.Query("SELECT * FROM competiteurs")
	if p.base.err != nil {
		fmt.Println("Erreur lors de l'execution de la requête")
		log.Fatal(p.base.err)
	}
	defer p.base.resultat.Close()

	var info [10]string
	var annonce1 int
	var annonce2 int
	var id int

	var nextcomp *Competiteur

	//On "clear" l'ancien tableau:
	p.comp = p.comp[:0]

	for p.base.resultat.Next() {
		p.base.err = p.base.resultat.Scan(&info[0], &info[1], &info[2], &info[3], &info[4], &info[5], &info[6], &info[7], &info[8], &info[9])
		if p.base.err != nil {
			fmt.Println("Erreur lors de la récupération des résultats: ")
			log.Fatal(p.base.err)
		}
		annonce1, _ = strconv.Atoi(info[7])
		annonce2, _ = strconv.Atoi(info[9])
		id, _ = strconv.Atoi(info[0])

		nextcomp = newCompetiteur(id, info[1], info[2], info[3], info[4], info[5], info[6], annonce1, info[8], annonce2)
		p.comp = append(p.comp, nextcomp)
	}
}

// Enregistrer les épreuves dans le tableau
func (p *Planning) getConfigurationEpreuve() {
	file, err := os.Open("../config/ConfigurationEpreuve.csv")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier \"ConfigurationEpreuve\": ")
		log.Fatal(err)
	}
	defer file.Close()

	var firstCall bool
	firstCall = true
	var nextconfig *ConfigurationEpreuve

	scanner := bufio.NewScanner(file)
	//On clear l'ancien tableau:
	p.cfgEpreuves = p.cfgEpreuves[:0]

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ";")
		if !firstCall {
			ordre, _ := strconv.Atoi(info[0])
			seuilMin, _ := strconv.Atoi(info[2])
			seuilMax, _ := strconv.Atoi(info[3])
			nbParPassage, _ := strconv.Atoi(info[4])
			dureeEchauffement, _ := strconv.Atoi(info[5])
			dureePassage, _ := strconv.Atoi(info[6])
			dureeAppel, _ := strconv.Atoi(info[7])
			surveillance, _ := strconv.Atoi(info[8])
			battementSerie, _ := strconv.Atoi(info[9])
			battementEpreuve, _ := strconv.Atoi(info[10])

			nextconfig = newConfigurationEpreuve(ordre, info[1], seuilMin, seuilMax, nbParPassage, dureeEchauffement, dureePassage, dureeAppel, surveillance,
				battementSerie, battementEpreuve, info[11])
			p.cfgEpreuves = append(p.cfgEpreuves, nextconfig)
		}
		firstCall = false
	}

	//Nombre de participants par épreuve
	for i := 0; i < len(p.cfgEpreuves); i++ {
		for j := 0; j < len(p.comp); j++ {
			if (p.comp[j].epreuve1 == p.cfgEpreuves[i].id) || (p.comp[j].epreuve2 == p.cfgEpreuves[i].id) {
				p.cfgEpreuves[i].nbParticipants = p.cfgEpreuves[i].nbParticipants + 1
			}
		}
	}

	sort.Sort(triEpreuves(p.cfgEpreuves))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// EpGeneration Epreuve generation function
func (p *Planning) EpGeneration(numEp int) {
	if numEp < len(p.cfgEpreuves) {
		p.planEpreuves = p.planEpreuves[:0]

		for j := 0; j < len(p.comp); j++ {
			if p.comp[j].epreuve1 == p.cfgEpreuves[numEp].id {
				plannEp := newPlanningEpreuve(p.cfgEpreuves[numEp].id, p.comp[j].id, p.comp[j].prenom, p.comp[j].nom, p.comp[j].sexe, p.comp[j].equipe, p.comp[j].annonce1)
				p.planEpreuves = append(p.planEpreuves, plannEp)

			} else if p.comp[j].epreuve2 == p.cfgEpreuves[numEp].id {
				plannEp := newPlanningEpreuve(p.cfgEpreuves[numEp].id, p.comp[j].id, p.comp[j].prenom, p.comp[j].nom, p.comp[j].sexe, p.comp[j].equipe, p.comp[j].annonce2)
				p.planEpreuves = append(p.planEpreuves, plannEp)
			}
		}
		if p.cfgEpreuves[numEp].id == "spd" || p.cfgEpreuves[numEp].id == "850" {
			sort.Sort(sort.Reverse(triAnnonces(p.planEpreuves)))
		} else {
			sort.Sort(triAnnonces(p.planEpreuves))
		}

	} else {
		fmt.Println("Numéro d'épreuve invalide")
	}
}

func (p *Planning) generationHoraires(fichier string) {
	var nbCompPassage int
	nbCompPassage = 0
	var numSerie int
	numSerie = 1

	//Calcul de l'heure de début
	if len(p.cfgEpreuves) < 1 {
		log.Fatal("Aucune épreuve enregistrée.")
	}

	heureSTR := strings.Split(p.cfgEpreuves[0].heureDebut, ":")
	var heure []int
	h, _ := strconv.Atoi(heureSTR[0])
	m, _ := strconv.Atoi(heureSTR[1])
	heure = append(heure, h, m)

	for i := 0; i < len(p.cfgEpreuves); i++ {
		//Initialisation des séries/passages:
		nbCompPassage = 0
		numSerie = 1

		//Recherche des compétiteurs pour cette épreuve
		p.EpGeneration(i)

		//Pour chaque competiteurs:
		for j := 0; j < len(p.planEpreuves); j++ {
			if p.planEpreuves[j].idEpreuve == p.cfgEpreuves[i].id {

				//Si le nombre de compétiteur est plein pour ce passage:
				// On modifie l'heure du prochain passage (le temps de l'épreuve)

				if nbCompPassage == p.cfgEpreuves[i].nbParPassage {
					nbCompPassage = 0
					numSerie = numSerie + 1
					heure[1] = heure[1] + p.cfgEpreuves[i].dureePassage + p.cfgEpreuves[i].surveillance + p.cfgEpreuves[i].battementSerie
					for heure[1] >= 60 {
						heure[1] = heure[1] - 60
						heure[0] = heure[0] + 1
					}
				}
				//On formate l'heure
				if heure[1] < 10 {
					p.planEpreuves[j].heurePassage = fmt.Sprint(heure[0], ":0", heure[1])
				} else {
					p.planEpreuves[j].heurePassage = fmt.Sprint(heure[0], ":", heure[1])
				}

				//Configuration des seuils de pénalité:
				p.planEpreuves[j].seuilMin = p.planEpreuves[j].annonce + p.cfgEpreuves[i].seuilMin
				p.planEpreuves[j].seuilMax = p.planEpreuves[j].annonce + p.cfgEpreuves[i].seuilMax

				//Configuration des numéro de passages
				nbCompPassage = nbCompPassage + 1
				p.planEpreuves[j].numSerie = numSerie
				p.planEpreuves[j].numPassage = nbCompPassage
			}
		}
		//Ouverture de l'épreuve suivante:
		heure[1] = heure[1] + p.cfgEpreuves[i].dureePassage + p.cfgEpreuves[i].surveillance + p.cfgEpreuves[i].battementSerie + p.cfgEpreuves[i].battementEpreuve
		for heure[1] >= 60 {
			heure[1] = heure[1] - 60
			heure[0] = heure[0] + 1
		}
		//Export de l'épreuve dans le fichier
		p.exportPlanEpreuve(fichier)
	}
}

func (p *Planning) exportPlanCompetition() {
	t := time.Now()
	date := fmt.Sprint(t.Year(), "_", int(t.Month()), "_", t.Day(), "_", t.Hour(), "_", t.Minute(), "_", t.Second())
	file, err := os.Create(fmt.Sprint("../var/export/archives/", date, "-PlanningCompetition.csv"))
	file2, err := os.Create(fmt.Sprint("../var/export/PlanningCompetition.csv"))
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier. Avez vous créé un dossier \"export\" dans le dossier de l'application?")
		log.Fatal(err)
	}

	file.WriteString(fmt.Sprint("\xEF\xBB\xBFId Epreuve;HeureOuverture(1ereEpreuve);Echauffement;Annonce;Temps/Series;nbSeries;Battement Epreuve\r\n"))
	file2.WriteString(fmt.Sprint("\xEF\xBB\xBFId Epreuve;HeureOuverture(1ereEpreuve);Echauffement;Annonce;Temps/Series;nbSeries;Battement Epreuve\r\n"))
	for i := 0; i < len(p.cfgEpreuves); i++ {
		var nbSeries int

		idEpreuve := p.cfgEpreuves[i].id
		heureOuverture := p.cfgEpreuves[i].heureDebut
		echauffement := p.cfgEpreuves[i].dureeEchauffement
		appel := p.cfgEpreuves[i].dureeAppel
		tempsSerie := p.cfgEpreuves[i].dureePassage + p.cfgEpreuves[i].surveillance + p.cfgEpreuves[i].battementSerie
		nbSeries = p.cfgEpreuves[i].nbParticipants / p.cfgEpreuves[i].nbParPassage

		if p.cfgEpreuves[i].nbParticipants%p.cfgEpreuves[i].nbParPassage != 0 {
			nbSeries = nbSeries + 1
		}
		battementEp := p.cfgEpreuves[i].battementEpreuve
		file.WriteString(fmt.Sprint(idEpreuve, ";", heureOuverture, ";", echauffement, ";", appel, ";", tempsSerie, ";", nbSeries, ";", battementEp, "\r\n"))
		file2.WriteString(fmt.Sprint(idEpreuve, ";", heureOuverture, ";", echauffement, ";", appel, ";", tempsSerie, ";", nbSeries, ";", battementEp, "\r\n"))
	}
}

func (p *Planning) exportPlanEpreuve(fichier string) {

	file, err := os.OpenFile(fichier, os.O_APPEND|os.O_WRONLY, 0777)
	file2, err := os.OpenFile("../var/export/PlanningEpreuve.csv", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier")
		log.Fatal(err)
	}

	for j := 0; j < len(p.planEpreuves); j++ {
		file.WriteString(fmt.Sprint(p.planEpreuves[j].idEpreuve, ";", p.planEpreuves[j].numSerie, ";", p.planEpreuves[j].numPassage, ";", p.planEpreuves[j].idComp, ";",
			p.planEpreuves[j].prenom, ";", p.planEpreuves[j].nom, ";", p.planEpreuves[j].sexe, ";", p.planEpreuves[j].equipe, ";",
			strconv.Itoa(p.planEpreuves[j].annonce), ";", strconv.Itoa(p.planEpreuves[j].seuilMin), ";", strconv.Itoa(p.planEpreuves[j].seuilMax), ";", p.planEpreuves[j].heurePassage, "\r\n"))

		file2.WriteString(fmt.Sprint(p.planEpreuves[j].idEpreuve, ";", p.planEpreuves[j].numSerie, ";", p.planEpreuves[j].numPassage, ";", p.planEpreuves[j].idComp, ";",
			p.planEpreuves[j].prenom, ";", p.planEpreuves[j].nom, ";", p.planEpreuves[j].sexe, ";", p.planEpreuves[j].equipe, ";",
			strconv.Itoa(p.planEpreuves[j].annonce), ";", strconv.Itoa(p.planEpreuves[j].seuilMin), ";", strconv.Itoa(p.planEpreuves[j].seuilMax), ";", p.planEpreuves[j].heurePassage, "\r\n"))
	}

}

func (p *Planning) generationPlanning() {
	t := time.Now()
	date := fmt.Sprint(t.Year(), "_", int(t.Month()), "_", t.Day(), "_", t.Hour(), "_", t.Minute(), "_", t.Second())
	fichier := fmt.Sprint("../var/export/archives/", date, "-PlanningEpreuve.csv")

	file, err := os.Create(fichier)
	file2, err := os.Create("../var/export/PlanningEpreuve.csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier planning:")
		log.Fatal(err)
	}
	file.WriteString(fmt.Sprint("\xEF\xBB\xBFEpreuve;Num Serie;Num Passage;Id Competiteur;Prenom;Nom;Sexe;Equipe;Annonce(s/m);Seuil Min;Seuil Max;Heure de passage\r\n"))
	file2.WriteString(fmt.Sprint("\xEF\xBB\xBFEpreuve;Num Serie;Num Passage;Id Competiteur;Prenom;Nom;Sexe;Equipe;Annonce(s/m);Seuil Min;Seuil Max;Heure de passage\r\n"))
	p.getCompetiteur()
	p.getConfigurationEpreuve()

	p.planEpreuves = p.planEpreuves[:0]
	p.generationHoraires(fichier)
	p.exportPlanCompetition()
}

func (p Planning) displayCompetiteur() {
	for j := 0; j < len(p.comp); j++ {
		p.comp[j].display()
	}
}

func (p Planning) displayConfigurationEpreuve() {
	for j := 0; j < len(p.cfgEpreuves); j++ {
		p.cfgEpreuves[j].display()
	}
}

func (p Planning) displayPlanningEpreuve() {
	for j := 0; j < len(p.planEpreuves); j++ {
		p.planEpreuves[j].display()
	}
}

func newPlanning(cheminBdd string) *Planning {
	p := new(Planning)
	p.base = newBdd(cheminBdd)

	return p
}
