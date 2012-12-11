package sockjs

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type infoData struct {
	Websocket    bool     `json:"websocket"`
	CookieNeeded bool     `json:"cookie_needed"`
	Origins      []string `json:"origins"`
	Entropy      uint32   `json:"entropy"`
}

func newInfoData(ws bool) infoData {
	return infoData{
		Websocket:    ws,
		CookieNeeded: true,
		Origins:      []string{"*:*"},
		Entropy:      rand.Uint32(),
	}
}

func handleInfo(h *Handler, w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	addCors(header, r)
	addContentTypeWithoutCache(header, "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json, err := json.Marshal(newInfoData(h.config.Websocket))
	if err != nil {
		panic(err)
	}
	w.Write(json)
}

func handleInfoOptions(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	addCorsAllowMethods(header, r, "OPTIONS, GET")
	addExpires(header)
	w.WriteHeader(http.StatusNoContent)
}
