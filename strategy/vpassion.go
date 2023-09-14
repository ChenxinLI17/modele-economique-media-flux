package strategy

import (
	"project/utils"
)

var PassionStrategy ViewerStrategy = *NewViewerStrategy("Passion", passion)

func passion(budget float32, themes []utils.Theme, catalogs []utils.Catalog) map[string]utils.Subscription {
	res := make(map[string]utils.Subscription)
	var FormatsPrefs = []utils.Format{"4K", "HD", "480p"}

	AnalyseMap := make(map[string]int)
	for i := 0; i < len(catalogs); i++ {
		score := 0
		themesvideos := []utils.Theme{}
		for j := 0; j < len(catalogs[i].Videos); j++ {
			themesunvideo := []utils.Theme{}
			for k := 0; k < len(catalogs[i].Videos[j].Themes); k++ {
				themesunvideo = append(themesunvideo, catalogs[i].Videos[j].Themes[k])
			}
			themesvideos = append(themesvideos, themesunvideo...)
		}
		for l := 0; l < len(themes); l++ {
			if utils.Contains(themesvideos, themes[l]) {
				score++
			}
		}
		AnalyseMap[catalogs[i].PlateformName] = score
	}
	maxscore := 0
	s := []string{}
	for _, v := range AnalyseMap {
		if maxscore < v {
			maxscore = v
		}
	}
	for k, v := range AnalyseMap {
		if v == maxscore {
			s = append(s, k)
		}
	}

	for i := 0; i < len(catalogs); i++ {
		if utils.Contains(s, catalogs[i].PlateformName) {
			for j := 0; j < len(catalogs[i].Subscriptions); j++ {
				//fmt.Println(catalogs[i].Subscriptions[j].Formats)
				if utils.CompareFormats(catalogs[i].Subscriptions[j].Formats, FormatsPrefs) {
					if catalogs[i].Subscriptions[j].Price < budget {
						res[catalogs[i].PlateformName] = catalogs[i].Subscriptions[j]
						budget = budget - catalogs[i].Subscriptions[j].Price
					}
				}
			}
		}
	}
	return res
}
