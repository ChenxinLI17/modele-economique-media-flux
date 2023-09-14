package plateform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/utils"
)

func (p *Plateform) sendCatalog(w http.ResponseWriter, r *http.Request) {
	// Check GET
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		errmsg := "Method not allowed, the allowed method is GET"
		fmt.Fprint(w, errmsg)
		return
	}

	// Gestion de la concurrence : pas d'écriture en parallèle
	p.Catalog.RLock()
	defer p.Catalog.RUnlock()

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(p.getCatalog())
	w.Write(serial)
}

func (p *Plateform) getCatalog() utils.Catalog {
	var catalog utils.Catalog

	catalog.PlateformName = p.Name
	catalog.Subscriptions = p.Subscriptions
	catalog.Videos = p.Catalog.Videos //On ne leur envoie pas le verrou et tout, c'est juste une récup pour traitement

	return catalog
}
