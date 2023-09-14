package viewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/utils"
)

func (v *Viewer) decodeContract(r *http.Response) utils.Contract {
	// Décoder les subs json reçues par post request sur l'url ":8080/subscriptions"
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp utils.Contract
	json.Unmarshal(buf.Bytes(), &resp)

	return resp
}

func (v *Viewer) SendSubscribe(port, subname string) (err error) {
	//Préparation du contrat
	req := utils.Contract{
		ClientName:       v.Name,
		SubscriptionName: subname,
		StrategyDemo:     v.Strategy.String(),
	}
	data, _ := json.Marshal(req)

	// Envoi de la requête de Subscription
	url := "http://localhost" + port + "/subscriptions"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}

	// Si erreur, on quitte
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	// Sinon, on décode le contrat avec notre ID (= validé)
	res := v.decodeContract(resp)
	v.contracts[res.PlatformName] = res
	return nil
}

func (v *Viewer) SendUnsubscribe(platform, port string) (err error) {
	contract, exists := v.contracts[platform]
	if !exists {
		log.Print("The platform " + platform + " does not exist whithin the Client " + v.Name)
		return
	}

	// Préparation du contrat
	req := utils.Contract{
		Id: contract.Id, // Juste l'id au cas où confusion POST
	}
	data, _ := json.Marshal(req)

	// Envoi de l'url
	url := "http://localhost/subscriptions" + port
	request, err := http.NewRequest("DELETE", url, bytes.NewReader(data))
	if err != nil {
		log.Print("Error in the send unsubscription", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	// Récupération de la réponse
	if resp.StatusCode != http.StatusOK {
		log.Print(v.Name+" failed to unsubscribe to "+platform, err)
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	return nil
}
