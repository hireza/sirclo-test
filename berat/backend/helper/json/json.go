package json

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	Meta Meta
	Data interface{} `json:",omitempty"`
}

type Meta struct {
	StatusCode int
	Message    interface{}
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

func Response(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}, data interface{}) {
	meta := Meta{
		StatusCode: statusCode,
		Message:    message,
	}

	res := Res{
		Meta: meta,
		Data: data,
	}

	WriteJSON(w, statusCode, res)
}
