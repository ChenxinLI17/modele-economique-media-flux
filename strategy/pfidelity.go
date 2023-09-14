package strategy

import (
    // "errors"
    "fmt"
    "project/utils"
    "sort"
)

func maxCountTheme(countTheme map[utils.Theme]int) (favoriteTheme []utils.Theme) {
    type themeCount struct {
        theme utils.Theme
        c     int
    }
    var favoriteThemeCount []themeCount
    for key, value := range countTheme {
        favoriteThemeCount = append(favoriteThemeCount, themeCount{key, value})
    }

    sort.Slice(favoriteThemeCount, func(i, j int) bool {
        return favoriteThemeCount[i].c > favoriteThemeCount[j].c
    })

    for t := range favoriteThemeCount {
        if favoriteTheme == nil {
            favoriteTheme = append(favoriteTheme, favoriteThemeCount[t].theme)
        }
        if t != 0 && favoriteThemeCount[t].c >= favoriteThemeCount[0].c {
            favoriteTheme = append(favoriteTheme, favoriteThemeCount[t].theme)
        } else {
            break
        }
    }
    return favoriteTheme
}

func maxCountType(countType map[utils.Type]int) (favoriteType []utils.Type) {
    type typeCount struct {
        _type utils.Type
        c     int
    }
    var favoriteTypeCount []typeCount
    for key, value := range countType {
        favoriteTypeCount = append(favoriteTypeCount, typeCount{key, value})
    }

    sort.Slice(favoriteTypeCount, func(i, j int) bool {
        return favoriteTypeCount[i].c > favoriteTypeCount[j].c
    })

    for f := range favoriteTypeCount {
        if favoriteType == nil {
            favoriteType = append(favoriteType, favoriteTypeCount[f]._type)
        }
        if f != 0 && favoriteTypeCount[f].c >= favoriteTypeCount[0].c {
            favoriteType = append(favoriteType, favoriteTypeCount[f]._type)
        } else {
            break
        }
    }
    return favoriteType
}

// // Analyse des vidéos visionnées => trouver le theme aimé => produit ce type
// func SWFTheme(watched []utils.Video) (countTheme map[utils.Theme]int, err error) {
//     if len(watched) == 0 {
//         return nil, errors.New("Erreur : la liste visionnée est vide")
//     }
//     countTheme = make(map[utils.Theme]int)
//     for _, v := range watched {
//         for _, t := range v.Themes {
//             if _, ok := countTheme[t]; ok {
//                 countTheme[t] += 1
//             } else {
//                 countTheme[t] = 1
//             }
//         }
//     }
//     return countTheme, nil
// }

// func SCFTheme(watched []utils.Video) (favoriteTheme []utils.Theme, err error) {
//     countTheme := make(map[utils.Theme]int)
//     countTheme, err = SWFTheme(watched)
//     return maxCountTheme(countTheme), err
// }

// // Analyser les types aimés par les subscribers
// func SWFType(watched []utils.Video) (countFormat map[utils.Type]int, err error) {
//     if len(watched) == 0 {
//         return nil, errors.New("Erreur : la liste visionnée est vide")
//     }
//     countFormat = make(map[utils.Type]int)
//     for _, v := range watched {
//         if _, ok := countFormat[v.Type]; ok {
//             countFormat[v.Type] += 1
//         } else {
//             countFormat[v.Type] = 1
//         }
//     }
//     return countFormat, nil
// }

// func SCFType(watched []utils.Video) (favoriteType []utils.Type, err error) {
//     countType := make(map[utils.Type]int)
//     countType, err = SWFType(watched)
//     return maxCountType(countType), err
// }

var BuyVideoStrategy PlatformStrategy = *NewPlatformStrategy("BuyVideo", buyVideo)

func buyVideo(watchHistory map[string]int, budget float32, catalog *utils.SynchronizedCatalog, store *utils.SynchronizedCatalog) {
    // Décider le type et le theme des vidéos à acheter
    fmt.Println("Trying to buy video")
    themeCount := make(map[utils.Theme]int)
    if watchHistory == nil {
        return
    }

	// Find the most popular Theme
    for video, times := range watchHistory {
		// Find the Video structure correspond from catalog
        var videostruct utils.Video
        for _, vid := range catalog.Videos {
            if vid.Title == video {
				fmt.Println("Found Correspond")
                videostruct = vid
            }
        }
        themes := videostruct.Themes
		// Count theme appreance
        for _, theme := range themes {
            if _, ok := themeCount[theme]; ok {
                themeCount[theme] += times
            } else {
                themeCount[theme] = times
            }
        }
    }
    fmt.Println("Theme Count : ", themeCount)
    bestThemes := maxCountTheme(themeCount)
	fmt.Println("Best Themes :", bestThemes)

	// Find the most popular Type
    typeCount := make(map[utils.Type]int)
    for video, times := range watchHistory {
        var videostruct utils.Video
        for _, vid := range catalog.Videos {
            if vid.Title == video {
                videostruct = vid
            }
        }
        if _, ok := typeCount[videostruct.Type]; ok {
            typeCount[videostruct.Type] += times
        } else {
            typeCount[videostruct.Type] = times
        }
    }
	fmt.Println("Type Count : ", typeCount)
    bestTypes := maxCountType(typeCount)
	fmt.Println("Best Types :", bestTypes)

    if store.Videos == nil || (len(bestTypes)==0 && len(bestThemes)==0) {
        return
    }
    for indice, v := range store.Videos {
        if v.Type == bestTypes[0] {
            if budget-v.Cost > 0 {
				fmt.Println("Buying video corresponding to my favorite type, video title = ", v.Title, ", indice = ", indice)
                // store.Lock()
                catalog.Lock()
				fmt.Println(len(catalog.Videos))
                catalog.Videos = append(catalog.Videos, v)
				fmt.Println(len(catalog.Videos))
                store.Videos = utils.RemoveVideo(store.Videos, indice)
                budget -= v.Cost
                catalog.Unlock()
                // store.Unlock()
				continue
            }
        }
        for _, t := range v.Themes {
            if t == bestThemes[0] {
                if budget-v.Cost > 0 {
					fmt.Println("Buying video corresponding to my favorite theme, video title = ", v.Title, ", indice = ", indice)
                    // store.Lock()
                    catalog.Lock()
					fmt.Println(len(catalog.Videos))
                    catalog.Videos = append(catalog.Videos, v)
					fmt.Println(len(catalog.Videos))
                    store.Videos = utils.RemoveVideo(store.Videos, indice)
                    budget -= v.Cost

                    catalog.Unlock()
                    // store.Unlock()
					continue
                }
            }
        }
    }
	fmt.Println("End of buying video")
}