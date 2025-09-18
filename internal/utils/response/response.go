package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/singhJasvinder101/basic-api/internal/types"
)

type APIResponse = types.APIResponse


const (
	StatusOK    = "ok"
	StatusError = "error"
)



func Success(data interface{}) APIResponse {
	return APIResponse{
		Status: StatusOK,
		Data:   data,
	}
}

func JsonWrite(w http.ResponseWriter, status int, res APIResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(res)
}

func GeneralError(err error) APIResponse {
	return APIResponse{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.ValidationErrors) APIResponse {
	var msgs []string

	for _, e := range err {
		switch e.ActualTag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("%s is required", e.Field()))
		default:
			msgs = append(msgs, fmt.Sprintf("%s is not valid", e.Field()))
		}
	}

	return APIResponse{
		Status: StatusError,
		Error:  strings.Join(msgs, ", "),
	}
}
