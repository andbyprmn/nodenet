package handler

import (
	"encoding/json"
	"net/http"
	"nodenet/compute"
	"nodenet/log"
	"nodenet/model"

	"github.com/julienschmidt/httprouter"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	value, err := model.GetKeyValue(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.LogError(err)
		return
	}

	json.NewEncoder(w).Encode(KeyValue{Key: key, Value: value})
}

func SetKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if kv.Key == "compute" {
		err := compute.ExecuteProgram(kv.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.LogError(err)
			return
		}
	}

	err = model.SetKeyValue(kv.Key, kv.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.LogError(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.LogCommand("Set key: " + kv.Key)
}
