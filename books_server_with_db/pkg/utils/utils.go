package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ParseBody(r *http.Request, x interface{}) {
	log.Info("unpacking json data")

	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Error(err)
	}

	log.Infof("%+v", x)
}
