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
	/*
		get param from url
		localhost:8080/api/time?zone=Asia/Istanbul
		request.URL.Query()["zone"] --> return the "zone=" section
		request.URL.Query()["zone"] is return "Asia/Istanbul" in this case
	*/
	keys, _ := request.URL.Query()["zone"]
	location, _ := time.LoadLocation(string(keys[0]))
	_, err := fmt.Fprint(writer, "the time ", time.Now().In(location))
	if err != nil {
		return
	}
}