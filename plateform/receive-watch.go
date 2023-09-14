package plateform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/utils"
)

func (p *Plateform) handleWatch(w http.ResponseWriter, r *http.Request) {
	log.Print("Plateform " + p.Name + " just received a Watch Request")

	p.Lock()
	// Check METHOD
	if r.Method == "POST" {
		log.Print("Action : Watch")
		p.watch(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		errmsg := "Method not allowed, the allowed method is POST"
		fmt.Fprint(w, errmsg)
		log.Print(errmsg)
		return
	}
	p.Unlock()
}

func (p *Plateform) watch(w http.ResponseWriter, r *http.Request) {
	watch, err := p.decodeWatch(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		log.Print("DECODE Error:", err.Error())
		return
	}

	// Checking the parameter Id
	if watch.Id == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field Id is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	// Checking the parameter Id
	if watch.VideoTitle == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		errmsg := "Field VideoTitle is mandatory but is empty"
		log.Print(errmsg)
		return
	}

	// Check the validity of the Id
	if !p.VideoExists(watch.VideoTitle) {
		w.WriteHeader(http.StatusNotFound) // 404
		errmsg := fmt.Sprintf("Video '%v' does not exist in plateform", watch.VideoTitle)
		log.Print(errmsg)
		return
	}

	// Saving the watch history
	if p.WatchHistory == nil {
		p.WatchHistory = make(map[string]int)
	}
	_, exist := p.WatchHistory[watch.VideoTitle]
	if exist {
		p.WatchHistory[watch.VideoTitle] += 1
	} else {
		p.WatchHistory[watch.VideoTitle] = 1
	}
	// Debug : pour tester si le nombre de watch a bien augmenté
	// println(p.WatchHistory[watch.VideoTitle])

	// Sending back ok header
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal("Watch successful")
	w.Write(serial)
}

func (*Plateform) decodeWatch(r *http.Request) (watch utils.Watch, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	err = json.Unmarshal(buf.Bytes(), &watch)

	return
}
