package plateform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/utils"
)

func (p *Plateform) handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	log.Print("Plateform " + p.Name + " just received a Subscription Request")

	p.Lock()
	// Check METHOD
	if r.Method == "POST" {
		log.Print("Action : Subscribe")
		p.subscribe(w, r)
	} else if r.Method == "DELETE" {
		log.Print("Action : Unsubscribe")
		p.unsubscribe(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		errmsg := "Method not allowed, the allowed methods are POST and DELETE"
		fmt.Fprint(w, errmsg)
		log.Print(errmsg)
		return
	}
	p.Unlock()
}

func (p *Plateform) subscribe(w http.ResponseWriter, r *http.Request) {
	contract, err := p.decodeContract(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		log.Print("DECODE Error:", err.Error())
		return
	}

	// Checking the parameters ClientName and SubscriptionName
	if contract.ClientName == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field ClientName is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	if contract.SubscriptionName == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field SubscriptionName is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	// Checking if Subscription actually exists
	if !p.SubscriptionExists(contract.SubscriptionName) {
		w.WriteHeader(http.StatusNotFound) // 404
		errmsg := fmt.Sprintf("Subscription '%v' does not exist'", contract.SubscriptionName)
		log.Print(errmsg)
		return
	}

	// Display Videos (pour savoir le nom de video pendant le test de /watch)
	fmt.Println(p.getCatalog().Videos[0].Title)

	// Saving the new Client
	contract.Id = p.addSubscriber(contract.SubscriptionName, contract.ClientName, contract.StrategyDemo)
	contract.PlatformName = p.Name

	// Sending back the filled contract
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(contract)
	w.Write(serial)
}

func (p *Plateform) unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Decode le contract
	contract, err := p.decodeContract(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		log.Print("DECODE Error:", err.Error())
		return
	}

	// Check the parameters
	if contract.Id == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field Id is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	// Check the validity of the Id
	if !p.ClientIdExists(contract.Id) {
		w.WriteHeader(http.StatusNotFound) // 404
		errmsg := fmt.Sprintf("Client '%v' does not exist in plateform", contract.Id)
		log.Print(errmsg)
		return
	}

	// Delete the client
	p.removeSubscriber(contract.Id)

	// Send Response
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal("Subscription delete successful")
	w.Write(serial)
}

func (*Plateform) decodeContract(r *http.Request) (contract utils.Contract, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	err = json.Unmarshal(buf.Bytes(), &contract)

	return
}
