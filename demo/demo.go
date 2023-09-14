package demo

import (
	"log"
	"project/data"
	"project/plateform"
	"project/strategy"
	"project/utils"
	"project/viewer"
	"sync"
)

type Data struct {
	Plateforms []*plateform.Plateform
	Viewers    []*viewer.Viewer
}

var once sync.Once

func LaunchDemo(n, m, t int, vQuotas, pQuotas []int) (Data, error) {
	// Initialisation du "monde" : viewers, vidéos dispo, plateforms, etc.
	var store utils.SynchronizedCatalog
	store.Videos = data.GenerateFakeVideo()

	viewers := createViewers(n, vQuotas)
	platforms := createPlatforms(m, pQuotas)

	for _, p := range platforms {
		p.RandomBuy(&store, 10)
	}

	// Lancement des serveurs REST en Amont + Url à donnés aux Viewers
	// Once car on peut relancer une démo mais pas les rest sur les memes ports
	//(et on ne veut pas non plus en ouvrir 50)
	once.Do(func() { setupPorts(platforms, viewers) })

	// Lancement du mois
	StartMonths(n, m, t, viewers, platforms, &store)

	// Return Data to the front
	return Data{Plateforms: platforms, Viewers: viewers}, nil
}

func createPlatforms(m int, pQuota []int) []*plateform.Plateform {
	platforms := plateform.NewSemiRandomPlateforms(m)

	// Assignation des stratégies
	var s = len(strategy.PlatformStrategies) // total stratégies

	// Pour chaque stratégie
	for i := 0; i < s; i++ {
		// Pour quota de stratégie plateformes
		if i < len(pQuota) {
			for j := 0; j < pQuota[i]; j++ {
				platforms[j].Strategy = strategy.PlatformStrategies[i]
			}
		}
	}

	return platforms
}

func createViewers(n int, vQuota []int) []*viewer.Viewer {
	viewers := viewer.NewRandomViewers(n)
	// Assignation des stratégies
	var s = len(strategy.ViewerStrategies) // total stratégies

	next := 0
	// Pour chaque stratégie
	for i := 0; i < s; i++ {
		// Pour quota de stratégie viewers
		if i < len(vQuota) {
			for j := 0; j < vQuota[i]; j++ {
				viewers[next].Strategy = strategy.ViewerStrategies[i]
				next++
			}
			log.Print("Setting ", vQuota[i], " with ", strategy.ViewerStrategies[i])
		}

	}

	return viewers
}

// Une seule fois grâce à sync.Once
func setupPorts(plateforms []*plateform.Plateform, viewers []*viewer.Viewer) {
	log.Println("************* SETUP THE PORTS")
	next := 0
	availablePorts := []string{":9000", ":9001", ":9002", ":9003", ":9004", ":9005", ":9006"}
	var assigned = make(map[string]string)
	for _, p := range plateforms {
		assigned[p.Name] = availablePorts[next]
		go p.StartRESTServer(availablePorts[next])
		next++
	}

	for _, v := range viewers {
		v.SetUrls(assigned) //Slice des ports attribués
	}
}

// func StartMonths(n, m, t int, viewers []*viewer.Viewer, platforms []*plateform.Plateform, store *utils.SynchronizedCatalog) {
// 	// Start Month of all agents for t months
// 	for month := 0; month < t; month++ {
// 		fmt.Println("************************************************************** LE MOIS ", month, " COMMENCE")
// 		// Préparer la sync par mois
// 		var wg sync.WaitGroup

// 		// On lance le mois pour les Viewers (les Plateforms sont déjà lancé via le serveur REST)
// 		wg.Add(n + m)

// 		for _, v := range viewers {
// 			go v.StartMonth(&wg)
// 		}

// 		for _, p := range platforms {
// 			go p.StartMonth(&wg, store)
// 		}

// 		wg.Wait()
// 	}
// }

func StartMonths(n, m, t int, viewers []*viewer.Viewer, platforms []*plateform.Plateform, store *utils.SynchronizedCatalog) {
	
	// Préparer la sync par mois
	var wg sync.WaitGroup
		// On lance le mois pour les Viewers (les Plateforms sont déjà lancé via le serveur REST)
	wg.Add(n + m)

	for _, v := range viewers {
		go v.StartMonth(&wg)
	}

	for _, p := range platforms {
		go p.StartMonth(&wg, store)
	}
	wg.Wait()

}
