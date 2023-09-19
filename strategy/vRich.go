package strategy

import "project/utils"

var RichStrategy ViewerStrategy = *NewViewerStrategy("Rich", rich)

func rich(budget float32, themes []utils.Theme, catalogs []utils.Catalog) map[string]utils.Subscription {
	res := make(map[string]utils.Subscription)

	// Parcours de chaque plateforme (1 catalogue = 1 plateforme)
	for _, plateform := range catalogs {
		//Recherche de l'abonnement le plus cher
		subToAdd := plateform.Subscriptions[0]
		for _, sub := range plateform.Subscriptions {
			if sub.Price > subToAdd.Price {
				subToAdd = sub
			}
		}

		//Ajout à la liste
		res[plateform.PlateformName] = subToAdd
	}

	return res
}
