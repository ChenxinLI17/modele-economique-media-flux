package plateform

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"

	"project/strategy"
	"project/utils"
	"sync"
	"time"
)

type Plateform struct {
	Name          string
	Catalog       utils.SynchronizedCatalog
	Subscriptions []utils.Subscription
	Subscribers   map[string][]*Subscriber
	WatchHistory  map[string]int
	Budget        float32 //Initial + Income - Expenses
	BudgetHistory []float32
	Strategy      strategy.PlatformStrategy
	sync.Mutex
}

type Subscriber struct {
	Id       string
	Name     string
	Strategy string //juste pour l'affichage pdt la demo
}

func (s Subscriber) String() string {
	return "<Subscriber n°" + s.Id + " (" + s.Name + ")>"
}

func NewSubscriber(i string, n string, strat string) *Subscriber {
	return &Subscriber{
		Id:       i,
		Name:     n,
		Strategy: strat,
	}
}

// ********************************************************* FONCTIONS DE BASE
func NewPlateform(name string, budget float32) *Plateform {
	// Set le nom + budget initial
	plateform := &Plateform{
		Name:     name,
		Budget:   budget,
		Strategy: strategy.BuyVideoStrategy,
	}

	plateform.Subscribers = make(map[string][]*Subscriber)

	// Génération random de subscriptions
	subs := generateSubscriptions()
	for _, sub := range subs {
		plateform.Subscriptions = append(plateform.Subscriptions, sub)
		plateform.Subscribers[sub.Name] = []*Subscriber{}
	}

	// Aucun subscribers au départ
	return plateform
}

func generateSubscriptions() (subs []utils.Subscription) {

	names := []string{"Basic", "Regular", "Premium"}

	// Nb d'abonnement à créer entre 1 et 3
	n := rand.Intn(3) + 1

	// Génération des prices
	var prices []float32
	for i := 0; i < n; i++ {
		prices = append(prices, 4.99+rand.Float32()*(19.99-4.99))
	}
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })

	for i := 0; i < n; i++ {
		var formats []utils.Format
		if n == 1 {
			formats = utils.Formats
		} else if n == 3 {
			formats = append(formats, utils.Formats[(n-1)-i])
		} else {
			formats = append(formats, utils.Formats[(n-1)-i])
		}
		subs = append(subs, utils.Subscription{Price: prices[i], Name: names[i], Formats: formats})
	}

	return
}

func (p *Plateform) String() string {
	str := p.Name + "(budget: " + utils.FormatFloat(p.Budget) + ")"

	str += "\n*********** CATALOGUE:\n"

	for _, video := range p.Catalog.Videos {
		str += video.String()
	}

	str += "\n*********** AVAILABLE SUBS :\n"
	for _, sub := range p.Subscriptions {
		str += sub.String()

		str += "\n--- SUBSCRIBERS :\n"
		for _, v := range p.Subscribers[sub.Name] {
			str += v.String()
		}
	}

	str += "\n\n"

	return str
}

// ********************************************************* GENERATION

func NewSemiRandomPlateforms(m int) []*Plateform {
	plateforms := make([]*Plateform, m)
	names := []string{"Netflix", "Disney+", "Amazon Prime Video", "Viki", "Apple TV"}

	for i := 0; i < m; i++ {
		plateforms[i] = NewPlateform(names[i], 1000) // No inequalities : elles commencent toutes avec 1000
	}

	return plateforms
}

func (p *Plateform) RandomBuy(store *utils.SynchronizedCatalog, n int) {
	// Ajout de n videos
	for i := 0; i < n; i++ {
		// Choose a random video in the list
		log.Printf(">>>>> %d", len(store.Videos))
		indice := rand.Intn(len(store.Videos))
		v := store.Videos[indice]

		// If we can afford it, we buy it = add to the catalog + remove from the list
		if p.Budget-v.Cost > 0 {
			store.Lock()
			p.Catalog.Lock()

			p.Catalog.Videos = append(p.Catalog.Videos, v)
			store.Videos = utils.RemoveVideo(store.Videos, indice)
			p.Budget -= v.Cost

			p.Catalog.Unlock()
			store.Unlock()
		}

		// else we pass
	}
}

// ********************************************************* UTILITAIRE

func (p *Plateform) SubscriptionExists(subName string) bool {
	for _, sub := range p.Subscriptions {
		if sub.Name == subName {
			return true
		}
	}

	return false
}

func (p *Plateform) ClientIdExists(id string) bool {
	for _, subscribers := range p.Subscribers {
		for _, sub := range subscribers {
			if sub.Id == id {
				return true
			}
		}
	}

	return false
}

func (p *Plateform) VideoExists(videoName string) bool {

	for _, video := range p.getCatalog().Videos {
		if videoName == video.Title {
			return true
		}
	}
	return false

}

func (p *Plateform) addSubscriber(subname, clientname, strategyForDemo string) string {
	// Génération d'un id unique - exemple : Netflix-EssentielPub-1
	newid := utils.GenerateUniqueID()
	log.Print("[Server] creating new subscriber, id=", newid)

	// Ajout d'un nouveau client à la liste
	p.Subscribers[subname] = append(p.Subscribers[subname], NewSubscriber(newid, clientname, strategyForDemo))

	return newid
}

func (p *Plateform) removeSubscriber(noclient string) {

	log.Println("[Server] deleting subscriber, id=", noclient)

	// Searching the subscriber
	for _, subscribers := range p.Subscribers {
		for index, subscriber := range subscribers {

			// Removing the subscriber
			if subscriber.Id == noclient {
				subscribers = append(subscribers[:index], subscribers[index+1:]...)
				log.Println(p.Subscribers)
			}
		}
	}
}

// ********************************************************* LANCEMENT

func (p *Plateform) StartRESTServer(port string) {
	// création du multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It's working !\n")
	})
	mux.HandleFunc("/catalog", p.sendCatalog)
	mux.HandleFunc("/subscriptions", p.handleSubscriptions)
	mux.HandleFunc("/pay", p.handlePay)
	mux.HandleFunc("/watch", p.handleWatch)

	// création du serveur http
	s := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// lancement du serveur
	log.Println("Listening on ", port)
	go log.Fatal(s.ListenAndServe())
}

func (p *Plateform) StartMonth(wg *sync.WaitGroup, store *utils.SynchronizedCatalog) {
	defer wg.Done()
	rand.Seed(int64(time.Now().UnixNano()))
	//buyVideo(watchedHistory map[*utils.Video]int, budget float32, catalog *utils.SynchronizedCatalog, store *utils.SynchronizedCatalog)
	store.Lock()
	p.Strategy.Apply(p.WatchHistory, p.Budget, &p.Catalog, store)
	store.Unlock()

}
