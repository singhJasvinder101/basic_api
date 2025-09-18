package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/singhJasvinder101/basic-api/internal/types"
	"github.com/singhJasvinder101/basic-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var student types.Student

		err := json.NewDecoder(req.Body).Decode(&student)

		// body empty case
		if errors.Is(err, io.EOF) {
			// response.JsonWrite(w, http.StatusBadRequest, err.Error())
			// response.JsonWrite(w, http.StatusBadRequest, response.GeneralError(err))
			response.JsonWrite(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("body cannot be empty")))
			return
		}


		if err != nil {
			slog.Error("error while decoding the student", slog.String("error", err.Error()))
			response.JsonWrite(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			err := err.(validator.ValidationErrors)

			slog.Error("error while validating the student", slog.String("error", err.Error()))
			response.JsonWrite(w, http.StatusBadRequest, response.ValidationError(err))
			return
		}

		slog.Info("creating a student")
		response.JsonWrite(w, http.StatusCreated, response.Success(student))
	}
}
