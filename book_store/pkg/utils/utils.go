package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, io_err := ioutil.ReadAll(r.Body); io_err == nil {
		if unmrshl_err := json.Unmarshal(body, x); unmrshl_err != nil {
			return
		}
	}
}
