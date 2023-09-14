package strategy

/*
import (
	"project/plateform"
	"project/utils"
	"sort"
)

// Get price of subscription
func SubsPrice(s utils.Subscription) float32 {
	return s.Price
}

func SetPrice(s utils.Subscription, newPrice float32) {
	s.Price = newPrice
}

// Le pourcentage de changement négatif pour la réduction
// Positive pour l'augmentation du prix
func ChangePrice(s utils.Subscription, percentage float32) {
	s.Price *= float32(percentage)
}



func BestAndWorstSubscription(p *plateform.Plateform, count []utils.Subscription) []utils.Subscription {
	type Count struct {
		sub utils.Subscription
		c   int
	}
	var occu []Count
	for key, value := range count {
		occu = append(occu, Count{key, value})
	}

	sort.Slice(occu, func(i, j int) bool {
		return occu[i].c > occu[j].c
	})

	return []utils.Subscription{occu[0].sub, occu[len(occu)-1].sub}
}

// On fait des promo pour la subscription le moins souscrit
func ProposePromotion(p *plateform.Plateform) {
	worstSub := p.BestAndWorstSubscription(p.SubscriptionCount())[0]
	NewPromotion(worstSub.Name, 0.8)
}

// On augmente le prix de la meilleure subscription
func (p plateform.Plateform) PriceIncrease() {
	bestSub := p.BestAndWorstSubscription(p.SubscriptionCount())[1]
	ChangePrice(bestSub, 1.2)
}

*/
