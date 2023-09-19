package test

import (
	"fmt"
	"math/rand"
	"project/utils"
	"testing"
	"time"
)

func generateRandomTitle() string {
	// 生成随机标题的逻辑
	// 这里使用示例标题 "Video X"，其中 X 是一个随机数
	return fmt.Sprintf("Video %d", rand.Intn(100))
}

func generateRandomDuration() int {
	// 生成随机持续时间的逻辑
	return rand.Intn(300) + 60 // 假设视频时长在 60 到 360 秒之间
}

func generateRandomType() utils.Type {
	// 从 Types 中随机选择一个 Type
	return utils.Types[rand.Intn(len(utils.Types))]
}

func generateRandomThemes() []utils.Theme {
	numThemes := 3
	themes := make([]utils.Theme, numThemes)
	for i := 0; i < numThemes; i++ {
		themes[i] = utils.Themes[rand.Intn(len(utils.Themes))]
	}
	return themes
}

func generateRandomCost() float32 {
	// 生成随机成本的逻辑
	return rand.Float32() * 20 // 假设成本在 0 到 20 之间
}

func TestAdd(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	NumVideos := 30

	var videos []utils.Video

	for i := 0; i < NumVideos; i++ {
		video := utils.Video{
			Title:    generateRandomTitle(),
			Duration: generateRandomDuration(),
			Type:     generateRandomType(),
			Themes:   generateRandomThemes(),
			Cost:     generateRandomCost(),
		}
		videos = append(videos, video)
	}
	for i, v := range videos {
		fmt.Printf("Video %d:\n", i+1)
		fmt.Printf("Title: %s\n", v.Title)
		fmt.Printf("Duration: %d\n", v.Duration)
		fmt.Printf("Type: %s\n", v.Type)
		fmt.Printf("Themes: %v\n", v.Themes)
		fmt.Printf("Cost: %.2f\n", v.Cost)
		fmt.Println()
	}
	//countTheme := make(map[utils.Theme]int)
	//countTheme, _ = strategy.SWFTheme(videos)
	//// 将 countTheme 转换为字符串
	//countThemeStr := ""
	//for theme, count := range countTheme {
	//	countThemeStr += fmt.Sprintf("%s: %d, ", theme, count)
	//}
	//// 打印 countThemeStr
	//fmt.Println("SWF: " + countThemeStr)
	//bestThemes := strategy.MaxCountTheme(themeCount)

}
