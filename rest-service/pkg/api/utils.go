package api

import (
	"encoding/json"
	"net/http"
)

// writeJSON wrapping user-friendly message with HTTP status as Json format and write it to the HTTP writer
func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {

		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON write user=friendly error json into the HTTP writer
func (app *Application) errorJSON(w http.ResponseWriter, httpStatusCode int, err error) {
	type jsonErr struct {
		Message string `json:"message"`
	}

	theError := jsonErr{
		Message: err.Error(),
	}

	err = app.writeJSON(w, httpStatusCode, theError, "error")
	if err != nil {

		return
	}
}
