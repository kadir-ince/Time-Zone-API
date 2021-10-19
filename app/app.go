package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

func getTime(writer http.ResponseWriter, request *http.Request) {
	keys, _ := request.URL.Query()["zone"] // get params in url
	location, _ := time.LoadLocation(string(keys[0]))
	_, err := fmt.Fprint(writer, "the time ", time.Now().In(location))
	if err != nil {
		return
	}
}