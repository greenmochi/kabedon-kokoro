package kokoro

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/greenmochi/kabedon-kokoro/logger"
)

// Run TODO
func Run(port int, gatewayPort int, shutdown chan<- bool) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/shutdown", shutdownHandler(shutdown))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	reply, err := json.Marshal(struct{ Message string }{"pong"})
	if err != nil {
		logger.Error("unable to marshal reply for pong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reply)
}

func shutdownHandler(shutdown chan<- bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reply, err := json.Marshal(struct{ Message string }{"shutdown request received"})
		if err != nil {
			logger.Error("unable to marshal reply for shutdown")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(reply)

		logger.Infof("shutdown request received. shutting down.")
		switch r.Method {
		case http.MethodGet:
			shutdown <- true
		}
	}
}