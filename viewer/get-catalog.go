package viewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/utils"
)

func (v *Viewer) readCatalog(r *http.Response) utils.Catalog {
	// Décoder les news json reçues par get request sur l'url ":8080/catalog"
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp utils.Catalog
	json.Unmarshal(buf.Bytes(), &resp)

	return resp
}

func (v *Viewer) GetCatalogs() (catalogs []utils.Catalog, err error) {
	for _, port := range v.serversUrl {
		var response *http.Response
		response, err = http.Get("http://localhost" + port + "/catalog")

		if err != nil {
			return
		}

		if response.StatusCode != http.StatusOK {
			err = fmt.Errorf("[%d] %s", response.StatusCode, response.Status)
			return
		}

		catalogs = append(catalogs, v.readCatalog(response))
	}

	return
}
