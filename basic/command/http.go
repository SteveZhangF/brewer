package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SteveZhangF/brewer/basic/errors"
	"github.com/sirupsen/logrus"
)

func Error(w http.ResponseWriter, err *errors.HTTPError, status ...int) {
	encoder := json.NewEncoder(w)
	if err.Log == true {
		logrus.Errorf("%+v", err)
	}
	w.WriteHeader(err.HTTPStatus)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	er := encoder.Encode(err)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func Success(w http.ResponseWriter, value interface{}) {

	var e error
	if value != nil {
		if value == "ok" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			encoder := json.NewEncoder(w)
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			e = encoder.Encode(value)
		}

	} else {
		w.WriteHeader(http.StatusOK)
		// w.Write(nil)
	}
	if e != nil {
		fmt.Println(e)
	}
}
