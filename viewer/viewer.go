package viewer

import (
	"fmt"
	"log"
	"math/rand"
	"project/strategy"
	"project/utils"
	"sync"

	"github.com/go-faker/faker/v4"
)

type Viewer struct {
	Name string

	budget        float64
	initialBudget float64
	themePrefs    []utils.Theme

	serversUrl map[string]string
	contracts  map[string]utils.Contract

	Strategy strategy.ViewerStrategy
}

func NewViewer(budget float64, preferences []utils.Theme) *Viewer {
	return &Viewer{
		Name:          faker.FirstName() + " " + faker.LastName(),
		budget:        budget,
		initialBudget: budget,
		themePrefs:    preferences,
		Strategy:      strategy.PassionStrategy,
		contracts:     make(map[string]utils.Contract)}
}

func (v *Viewer) String() string {
	return fmt.Sprintf("Viewer: budget %.2f, preferences %v", v.budget, v.themePrefs)
}

func NewRandomViewers(n int) []*Viewer {
	viewers := make([]*Viewer, n)

	for i := 0; i < n; i++ {
		// Budget random
		budget := 10 + rand.Float64()*90

		preferences := []utils.Theme{}
		for j := 0; j < rand.Intn(4)+1; j++ {
			randtheme := utils.Themes[rand.Intn(4)]

			for utils.Contains(preferences, randtheme) {
				randtheme = utils.Themes[rand.Intn(4)]
			}

			preferences = append(preferences, randtheme)
		}

		// Add the viewer to the list
		viewers[i] = NewViewer(budget, preferences)
	}

	return viewers
}

func (v *Viewer) SetUrls(urls map[string]string) {
	v.serversUrl = urls
}

// ********************************************************* LANCEMENT

func (v *Viewer) StartMonth(wg *sync.WaitGroup) {
	defer wg.Done()

	// STEP 1 : Get News
	catalogs, err := v.GetCatalogs()
	if err != nil {
		log.Print(v.Name+" a rencontré une erreur : ", err)
		return
	}

	// STEP 2 : Lancer l'analyse et la traiter
	subs := v.Strategy.Apply(float32(v.budget), v.themePrefs, catalogs)
	v.sortSubs(subs)

	// STEP 3 : Visionner

	// loop for all the platform subscribed
	for _, contract := range v.contracts {

		// find the catalog of this platform
		var catalogOfContract utils.Catalog
		for _, catalog := range catalogs {
			if catalog.PlateformName == contract.PlatformName {
				catalogOfContract = catalog
			}
		}

		// find all videos in this catalog that has the theme preferred
		for _, video := range catalogOfContract.Videos {
			for _, theme := range video.Themes {
				if utils.Contains(v.themePrefs, theme) {
					v.SendWatch(v.serversUrl[contract.PlatformName], contract.Id, video.Title)
				}
			}
		}
	}

	// STEP 4 : Pay
	for _, contract := range v.contracts {
		v.SendPayment(v.serversUrl[contract.PlatformName], contract.Id)
	}

	// If the viewer can't afford it, the subscription is cancelled
	v.budget += 0.5 * v.initialBudget
}

func (v *Viewer) sortSubs(subs map[string]utils.Subscription) {
	// D'abord, on se désabonne de tout ce qui ne nous intéresse pas/plus
	for _, contract := range v.contracts {
		// Est-ce qu'elle est dans les subs conservés ?
		_, exists := subs[contract.PlatformName]
		if !exists { // Non
			v.SendUnsubscribe(contract.PlatformName, v.serversUrl[contract.PlatformName])
		}
	}

	// Ensuite, on s'abonne à qui nous intéresse + upgrade
	for platform, sub := range subs {
		// Est-ce que j'ai déjà un contract avec cette plateforme ?
		contract, exists := v.contracts[platform]
		if !exists { // Non
			v.SendSubscribe(v.serversUrl[platform], sub.Name)
		} else if contract.SubscriptionName != sub.Name { // Oui mais c'est un abonnement différent
			v.SendUnsubscribe(platform, v.serversUrl[platform])
			v.SendSubscribe(v.serversUrl[platform], sub.Name)
		}
	}
}
