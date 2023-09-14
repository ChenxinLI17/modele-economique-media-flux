package utils

import (
	"fmt"
	"strconv"
	"sync"
)

type Subscription struct {
	Name    string   `json:"name"`
	Price   float32  `json:"price"`
	Formats []Format `json:"formats"`
}

func (s Subscription) String() string {
	return "<div class='fw-bold'>" + s.Name + "</div>Price : " + fmt.Sprintf("%.2f", s.Price) + " €\n"
}

//****************************************************** FORMAT

type Format string

const ( //High, Mid and Low Qualities
	HQ Format = "4K"
	MQ Format = "HD"
	LQ Format = "480p"
)

var Formats = []Format{HQ, MQ, LQ}

func (f Format) String() string {
	return string(f)
}

//****************************************************** VIDEO

type Video struct {
	Title    string  `json:"title"`
	Duration int     `json:"duration"`
	Type     Type    `json:"type"`
	Themes   []Theme `json:"themes"`
	Cost     float32 `json:"cost"`
}

type SynchronizedCatalog struct {
	sync.RWMutex
	Videos []Video `json:"videos"`
}

func formatDuration(duration int) string {
	// Conversion en hueures et récup du reste en minutes
	hours := int(duration) / 60
	minutes := int(duration) % 60

	// Concat "h" s'il y a des heures
	durationString := ""
	if hours > 0 {
		durationString += strconv.Itoa(hours) + "h"
	}

	// Concat minutes s'il y a des minutes
	if minutes > 0 {
		durationString += strconv.Itoa(minutes) + "min"
	}

	return durationString
}

func (v Video) String() string {
	themes := ""
	for i, theme := range v.Themes {
		themes += theme.String()
		if i <= len(themes)-1 {
			themes += ", "
		}
	}
	return fmt.Sprintf(`<b>%s</b><br>Type : %s (%s) \n<i>Thèmes: %s </i>`, v.Title, v.Type, formatDuration(v.Duration), themes)
}

//****************************************************** THEME ENUM

type Type string

const (
	Film         Type = "Film"
	Serie        Type = "Serie"
	Documentaire Type = "Documentaire"
	Emission     Type = "Emission"
)

var Types = []Type{Film, Serie, Documentaire, Emission}

func (t Type) String() string {
	return string(t)
}

//****************************************************** THEME ENUM
type Theme string

const (
	Aventure       Theme = "Aventure"
	Action         Theme = "Action"
	Historique     Theme = "Historique"
	Guerre         Theme = "Guerre"
	SuperHeros     Theme = "Super-héros"
	Comedie        Theme = "Comédie"
	Education      Theme = "Éducation"
	Drame          Theme = "Drame"
	Medecine       Theme = "Médecine"
	ScienceFiction Theme = "Science-fiction"
	Mystere        Theme = "Mystère"
	Art            Theme = "Art"
	Nature         Theme = "Nature"
	Societe        Theme = "Société"
	Religion       Theme = "Religion"
	Technologie    Theme = "Technologie"
	Astronomie     Theme = "Astronomie"
	Science        Theme = "Science"
	Anthropologie  Theme = "Anthropologie"
	Archeologie    Theme = "Archéologie"
	Biologie       Theme = "Biologie"
	Politique      Theme = "Politique"
	Architecture   Theme = "Architecture"
)

var Themes = []Theme{
	Aventure,
	Action,
	Historique,
	Guerre,
	SuperHeros,
	Comedie,
	Education,
	Drame,
	Medecine,
	ScienceFiction,
	Mystere,
	Art,
	Nature,
	Societe,
	Religion,
	Technologie,
	Astronomie,
	Science,
	Anthropologie,
	Archeologie,
	Biologie,
	Politique,
	Architecture,
}

func (t Theme) String() string {
	return string(t)
}
