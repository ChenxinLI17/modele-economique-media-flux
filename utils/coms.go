package utils

type Catalog struct {
	PlateformName string         `json:"name"`
	Videos        []Video        `json:"catalog"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (c Catalog) String() string {
	str := "Catalogue de " + c.PlateformName

	for _, s := range c.Subscriptions {
		str += s.String()
	}

	for _, v := range c.Videos {
		str += v.String()
	}

	return str
}

// **********************************************************************
type Contract struct {
	ClientName       string `json:"client,omitempty"`
	SubscriptionName string `json:"subscription,omitempty"`
	Id               string `json:"id,omitempty"`
	PlatformName     string `json:"platform,omitempty"`
	StrategyDemo     string `json:"strategy,omitempty"`
}

func (c Contract) String() string {
	str := "Contrat de " + c.ClientName + " pour " + c.SubscriptionName + " (id: " + c.Id + ")"
	return str
}

// **********************************************************************

type Payment struct {
	Id     string  `json:"id,omitempty"`
	Amount float32 `json:"amount,omitempty"`
}

func (p Payment) String() string {
	str := "Paiement de " + p.Id
	return str
}

// **********************************************************************

type Watch struct {
	Id         string `json:"id"`
	VideoTitle string `json:"video"`
}

func (w Watch) String() string {
	str := "Watch de " + w.Id + "pour le video" + w.VideoTitle
	return str
}
