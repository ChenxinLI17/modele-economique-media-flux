package website

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project/demo"
	"project/strategy"
	"strconv"
)

// Lancement du serveur de démo
func Start() {
	http.HandleFunc("/", displayForm)
	http.HandleFunc("/result", handleForm)

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func displayForm(w http.ResponseWriter, r *http.Request) {
	//log.Print("Ouverture du fichier form.html")
	tmpl, err := template.ParseFiles("../website/form.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	availableStrategies := struct {
		Error              string
		ViewerStrategies   []strategy.ViewerStrategy
		PlatformStrategies []strategy.PlatformStrategy
	}{
		Error:              "",
		ViewerStrategies:   strategy.ViewerStrategies,
		PlatformStrategies: strategy.PlatformStrategies,
	}

	//log.Print("variable :", availableStrategies)
	err = tmpl.Execute(w, availableStrategies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	log.Print("Handle Form Request")
	if r.Method != "POST" {
		displayFormWithError(w, "Vous devez d'abord paramétrer la simulation avant d'accéder aux résultats.")
		return
	}

	log.Print("Getting the parameters")
	// Récupération des champs du formulaire
	n, err := strconv.Atoi(r.FormValue("n"))
	if err != nil {
		log.Print("n est invalide")
		displayFormWithError(w, "Le nombre de viewers doit être un nombre entier.")
		return
	}

	m, err := strconv.Atoi(r.FormValue("m"))
	if err != nil {
		log.Print("m est invalide")
		displayFormWithError(w, "Le nombre de plateforme doit être un nombre entier.")
		return
	}

	t, err := strconv.Atoi(r.FormValue("t"))
	if err != nil {
		log.Print("t est invalide")
		displayFormWithError(w, "La période de temps doit être un nombre entier.")
		return
	}

	viewerStrategies := r.Form["viewerStrategy[]"]
	platformStrategies := r.Form["platformStrategy[]"]

	log.Print("Convert tab to int")

	// Convertit les valeurs de viewerStrategies en entiers
	var vQuota = make([]int, len(viewerStrategies))
	for i, strToConvert := range viewerStrategies {
		vQuota[i], _ = strconv.Atoi(strToConvert)
	}

	// Convertit les valeurs de platformStrategies en entiers
	var pQuota = make([]int, len(platformStrategies))
	for i, strToConvert := range platformStrategies {
		pQuota[i], _ = strconv.Atoi(strToConvert)
	}

	//fmt.Println(viewerStrategies, platformStrategies)

	log.Print("Check args")
	// Vérifications des arguments
	if n < 1 {
		log.Print("n ne respecte pas les conditions")
		displayFormWithError(w, "Le nombre de viewers doit être supérieur à 1")
		return
	}

	if m < 1 || m > 5 {
		log.Print("m ne respecte pas les conditions")
		displayFormWithError(w, "Le nombre de plateformes doit être compris entre 1 et 5.")
		return
	}

	if t < 1 {
		log.Print("t ne respecte pas les conditions")
		displayFormWithError(w, "La période de temps doit être supérieure à 0.")
		return
	}

	log.Print("Launch Demonstration")
	// On lance la démonstration
	result, err := demo.LaunchDemo(n, m, t, vQuota, pQuota)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print(result, err)

	// Redirection vers la page de visualisation des résultats
	resultPage, err := template.ParseFiles("../website/result.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = resultPage.Execute(w, map[string]interface{}{
		"Plateforms": result.Plateforms,
		"Viewers":    result.Viewers,
	})

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func displayFormWithError(w http.ResponseWriter, errors string) {
	// Affichez le formulaire HTML avec l'alerte d'erreur
	tmpl, err := template.ParseFiles("./website/form.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	availableStrategies := struct {
		Error              string
		ViewerStrategies   []strategy.ViewerStrategy
		PlatformStrategies []strategy.PlatformStrategy
	}{
		Error:              errors,
		ViewerStrategies:   strategy.ViewerStrategies,
		PlatformStrategies: strategy.PlatformStrategies,
	}

	err = tmpl.Execute(w, availableStrategies)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
