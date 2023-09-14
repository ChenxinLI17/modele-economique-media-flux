package strategy

import (
	"math/rand"
	"project/utils"
)

var EcoStrategy ViewerStrategy = *NewViewerStrategy("Economique", economique)

func economique(budget float32, themes []utils.Theme, catalogs []utils.Catalog) map[string]utils.Subscription {
	res := make(map[string]utils.Subscription)
	minprice := 4.99 + rand.Float32()*float32(rand.Intn(4))
	// Parcours de chaque plateforme (1 catalogue = 1 plateforme)
	//Recherche de l'abonnement le moins cher dans toutes les plateforme
	for _, plateform := range catalogs {
		for _, sub := range plateform.Subscriptions {
			if sub.Price < minprice {
				minprice = sub.Price
			}
		}
	}
	if minprice > budget{
		return res
	}else {
		for _, plateform := range catalogs {
			for _, sub := range plateform.Subscriptions {
				if sub.Price == minprice {
					res[plateform.PlateformName] = sub
				}
			}
		}
	}
	

	return res
}
