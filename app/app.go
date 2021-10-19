package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	_, err := fmt.Fprint(writer, "the time")
	if err != nil {
		return
	}
}