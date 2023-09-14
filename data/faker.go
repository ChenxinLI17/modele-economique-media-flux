package data

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"project/utils"
	"time"
)

type video struct {
	Title  string        `json:"titre"`
	Themes []utils.Theme `json:"themes"`
}

type content struct {
	Films         []video `json:"films"`
	Documentaires []video `json:"documentaires"`
	Series        []video `json:"series"`
	Emissions     []video `json:"emissions"`
}

func GenerateFakeVideo() (list []utils.Video) {
	// Read the JSON to get random title and themes
	var contenu content

	jsonData, err := os.ReadFile("../data/videos.json")
	log.Printf("osssss %s", jsonData)
	err = json.Unmarshal(jsonData, &contenu)
	if err != nil {
		log.Println(err)
		return nil
	}

	//fmt.Println(contenu)

	// Création de film
	for _, base := range contenu.Films {
		var v utils.Video

		v.Title = base.Title
		v.Themes = base.Themes
		v.Duration = duration()
		v.Cost = cost(v.Themes, v.Duration)
		v.Type = utils.Film

		list = append(list, v)
	}

	// Création de documentaires
	for _, base := range contenu.Documentaires {
		var v utils.Video

		v.Title = base.Title
		v.Themes = base.Themes
		v.Duration = duration()
		v.Cost = cost(v.Themes, v.Duration)
		v.Type = utils.Documentaire

		//fmt.Println("title :", base.Title, " / video : ", v)

		list = append(list, v)
	}

	// Création de séries
	// TODO

	// Créations d'émissions
	// TODO

	return list

}

func duration() int {
	rand.Seed(time.Now().UnixNano())

	// Intervalle entre 10 min et 3 heures
	minDuration := 10
	maxDuration := 3 * 60

	// Génération random d'une durée dans l'intervalle
	return rand.Intn(maxDuration-minDuration) + minDuration
}

func cost(themes []utils.Theme, duration int) float32 {
	cost := float32(0.0)

	// Augmentation du cout en fonction de la durée de la vidéo
	cost += float32(duration) * 0.2

	// Augmentation du cout en fonction des thèmes abordés
	for _, theme := range themes {
		switch theme {
		case utils.ScienceFiction, utils.Historique, utils.SuperHeros:
			cost += 400

		case utils.Archeologie, utils.Aventure, utils.Technologie:
			cost += 250

		case utils.Science, utils.Nature, utils.Guerre, utils.Action:
			cost += 175

		default:
			cost += 100
		}
	}

	return cost
}
