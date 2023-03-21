package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"log"
	"net/http"
)

func checkPrescrption(w http.ResponseWriter, r *http.Request) {
	var newPresc model.Presc
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &newPresc)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("checking prescription")
}
