package plateform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/utils"
)

func (p *Plateform) handlePay(w http.ResponseWriter, r *http.Request) {
	log.Print("Plateform " + p.Name + " just received a Payment Request")

	p.Lock()
	// Check METHOD
	if r.Method == "POST" {
		log.Print("Action : Pay")
		p.pay(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		errmsg := "Method not allowed, the allowed method is POST"
		fmt.Fprint(w, errmsg)
		log.Print(errmsg)
		return
	}
	p.Unlock()
}

func (p *Plateform) pay(w http.ResponseWriter, r *http.Request) {
	payment, err := p.decodePayment(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		log.Print("DECODE Error:", err.Error())
		return
	}

	// Checking the parameter Id
	if payment.Id == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field Id is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	// Check the validity of the Id
	if !p.ClientIdExists(payment.Id) {
		w.WriteHeader(http.StatusNotFound) // 404
		errmsg := fmt.Sprintf("ClientId '%v' does not exist in plateform", payment.Id)
		log.Print(errmsg)
		return
	}

	// Performing a payment
	newamount, ok := p.prelevement(payment.Id, payment.Amount)

	// Sending back the filled contract
	if ok {
		payment.Amount = newamount
		w.WriteHeader(http.StatusOK)
		serial, _ := json.Marshal(payment)
		w.Write(serial)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		serial, _ := json.Marshal(payment)
		w.Write(serial)
	}

}

func (*Plateform) decodePayment(r *http.Request) (payment utils.Payment, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &payment)
	return
}

func (p *Plateform) prelevement(id string, amount float32) (float32, bool) {
	// Find client's subscription
	for subscriptionName, subs := range p.Subscribers {
		for _, sub := range subs {
			if sub.Id == id {
				// Find the price of subscription
				for _, subscription := range p.Subscriptions {
					if subscription.Name == subscriptionName {
						if amount < subscription.Price {
							return amount, false
						} else {
							p.Budget += subscription.Price
							return amount - subscription.Price, true
						}
					}
				}
			}
		}
	}
	return amount, false
}
