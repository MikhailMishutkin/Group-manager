package adapters

import (
	"encoding/json"
	"net/http"
)

// обработка ошибки для вывода пользователю
func JSONError(httpcode int, code, msg string, w http.ResponseWriter) {
	type Error struct {
		Code    *string `json:"code,omitempty"`
		Message *string `json:"message,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpcode)
	json.NewEncoder(w).Encode(
		Error{
			Code:    &code,
			Message: &msg,
		},
	)
}
