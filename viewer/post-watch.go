package viewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/utils"
)

// func (v *Viewer) decodeWatch(r *http.Response) utils.Watch {
// 	// Décoder le watch json reçues par post request sur l'url
// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(r.Body)

// 	var resp utils.Watch
// 	json.Unmarshal(buf.Bytes(), &resp)

// 	return resp
// }

func (v *Viewer) SendWatch(port, id, video string) (err error) {
	// Préparation de Watch
	req := utils.Watch{
		Id:         id,
		VideoTitle: video,
	}
	data, _ := json.Marshal(req)

	// Envoi de la requête de Watch
	url := "http://localhost" + port + "/watch"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}

	// Si erreur, on quitte
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	// Sinon, watch succes
	// res := v.decodePayment(resp)

	return nil
}