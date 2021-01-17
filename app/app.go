package app

import (
	"encoding/json"
	"github.com/Javlopez/opiapi/infrastructure"
	"net/http"
	//"github.com/gorilla/context"
)

//AppContext struct
type OpiAppContext struct {
	Container infrastructure.Container
}

//AppHandler struct
type AppHandler struct {
	*OpiAppContext
	H func(*OpiAppContext, http.ResponseWriter, *http.Request) (int, interface{})
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode int
		data       interface{}
	)
	statusCode, data = ah.H(ah.OpiAppContext, w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	return
}
