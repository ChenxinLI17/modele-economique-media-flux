package viewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/utils"
)

func (v *Viewer) decodePayment(r *http.Response) utils.Payment {
	// Décoder le payment json reçues par post request sur l'url
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp utils.Payment
	json.Unmarshal(buf.Bytes(), &resp)

	return resp
}

func (v *Viewer) SendPayment(port, id string) (err error) {
	//Préparation du payment
	req := utils.Payment{
		Id:       id,
		Amount:   float32(v.budget),
	}
	data, _ := json.Marshal(req)

	// Envoi de la requête de pay
	url := "http://localhost" + port + "/pay"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}

	// Si erreur, on quitte
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	// Sinon, on décode le payment
	res := v.decodePayment(resp)
	v.budget = float64(res.Amount)
	return nil
}