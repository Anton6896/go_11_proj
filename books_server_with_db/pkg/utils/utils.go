package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ParseBody(r *http.Request, x interface{}) {
	log.Info("unpacking json data")
	// will read request and put it to some interface
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
