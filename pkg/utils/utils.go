package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if Body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(Body), x); err != nil {
			return
		}
	}

}
