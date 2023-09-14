package strategy

import (
	"project/utils"
)

type Strategy struct {
	Name string
}

var ViewerStrategies = []ViewerStrategy{RichStrategy, EcoStrategy, PassionStrategy}
var PlatformStrategies = []PlatformStrategy{BuyVideoStrategy}

//************************************************ VIEWER STRATEGY

type ViewerStrategy struct {
	Strategy
	Apply func(float32, []utils.Theme, []utils.Catalog) map[string]utils.Subscription
}

func (v ViewerStrategy) String() string {
	return v.Strategy.Name
}

func NewViewerStrategy(name string, fct func(float32, []utils.Theme, []utils.Catalog) map[string]utils.Subscription) *ViewerStrategy {
	v := &ViewerStrategy{Apply: fct}
	v.Strategy.Name = name
	return v
}

//************************************************ PLATEFORM STRATEGY

type PlatformStrategy struct {
	Strategy
	Apply func(map[string]int, float32, *utils.SynchronizedCatalog, *utils.SynchronizedCatalog)
}

func NewPlatformStrategy(name string, fct func(map[string]int, float32, *utils.SynchronizedCatalog, *utils.SynchronizedCatalog)) *PlatformStrategy {
	v := &PlatformStrategy{Apply: fct}
	v.Strategy.Name = name
	return v
}

func (s PlatformStrategy) String() string {
	return s.Strategy.Name
}
