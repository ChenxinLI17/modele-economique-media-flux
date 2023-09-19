package strategy

import (
	"project/utils"
	"sort"
)

func calculateRanking(watchHistory map[string]int, catalog *utils.SynchronizedCatalog) (map[utils.Theme]int, map[utils.Type]int) {
	themeCount := make(map[utils.Theme]int)
	typeCount := make(map[utils.Type]int)

	for video, times := range watchHistory {
		var videostruct utils.Video
		for _, vid := range catalog.Videos {
			if vid.Title == video {
				videostruct = vid
			}
		}

		// 计算主题排行
		for _, theme := range videostruct.Themes {
			themeCount[theme] += times
		}

		// 计算类型排行
		typeCount[videostruct.Type] += times
	}

	return themeCount, typeCount
}

func sortRankingTheme(inputMap map[utils.Theme]int) []utils.Theme {

	pairs := make([]struct {
		Key   utils.Theme
		Value int
	}, len(inputMap))

	i := 0
	for key, value := range inputMap {
		pairs[i] = struct {
			Key   utils.Theme
			Value int
		}{key, value}
		i++
	}

	// 降序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	ranking := make([]utils.Theme, len(pairs))
	for i, pair := range pairs {
		ranking[i] = pair.Key
	}

	return ranking
}

// Analyse des vidéos visionnées de histoire
func SWFWatchHistory(watchHistoryLatest map[string]int, watchHistoryLatest2 map[string]int, watchHistoryLatest3 map[string]int, catalog *utils.SynchronizedCatalog) (profileTheme [][]utils.Theme, err error) {
	themeCount1, _ := calculateRanking(watchHistoryLatest, catalog)
	themeCount2, _ := calculateRanking(watchHistoryLatest2, catalog)
	themeCount3, _ := calculateRanking(watchHistoryLatest3, catalog)
	// 根据播放量排序主题排行榜
	rankingTheme1 := sortRankingTheme(themeCount1)
	rankingTheme2 := sortRankingTheme(themeCount2)
	rankingTheme3 := sortRankingTheme(themeCount3)

	profileTheme = make([][]utils.Theme, 0)

	// 添加 ranking1 到 profileTheme（5次）
	for i := 0; i < 5; i++ {
		profileTheme = append(profileTheme, rankingTheme1)
	}

	// 添加 ranking2 到 profileTheme（4次）
	for i := 0; i < 4; i++ {
		profileTheme = append(profileTheme, rankingTheme2)
	}

	// 添加 ranking3 到 profileTheme（3次）
	for i := 0; i < 3; i++ {
		profileTheme = append(profileTheme, rankingTheme3)
	}
	return profileTheme, nil
}
