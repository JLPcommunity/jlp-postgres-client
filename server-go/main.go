package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://pg_demo.virkea.com/get_data")
		if err != nil {
			errMsg := Msg{"ERROR", err.Error()}
			eByte, err := json.Marshal(errMsg)
			if err != nil {
				log.Fatal(err)
			}
			ResponseWithJSON(w, eByte, http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, body, http.StatusOK)
	}
	router.HandleFunc("/get", handler)

	http.ListenAndServe(":8383", router)
}

type Msg struct {
	Type  string `json:"type"`
	Cause string `json:"cause"`
}

// ResponseWithJSON compose to JSON before write thru http
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(json)
}
