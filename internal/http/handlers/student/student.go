package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/singhJasvinder101/basic-api/internal/types"
	"github.com/singhJasvinder101/basic-api/internal/utils/response"
	"github.com/singhJasvinder101/basic-api/storage/sqlite"
)

func New(storage *sqlite.Sqlite) http.HandlerFunc {
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


		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		
		if err != nil {
			slog.Error("error while creating a student", slog.String("error", err.Error()))
			response.JsonWrite(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		student.ID = lastId
		
		slog.Info("creating a student")
		
		response.JsonWrite(w, http.StatusCreated, response.Success(map[string]int64{
			"student_id": lastId,
		}))
	}
}

func GetStudentById(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		idStr := req.PathValue("id")
		slog.Info("getting student by id", slog.String("id", idStr))

		idInt, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			slog.Error("error while parsing id", slog.String("error", err.Error()))
			response.JsonWrite(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("id must be a number")))
			return
		}

		student, err := storage.GetStudentById(idInt)
		if err != nil {
			slog.Error("error while getting student by id", slog.String("error", err.Error()))
			response.JsonWrite(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if student.ID == 0 {
			slog.Error("student not found", slog.String("id", idStr))
			response.JsonWrite(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("student not found")))
			return
		}

		response.JsonWrite(w, http.StatusOK, response.Success(student))
		slog.Info("student found", slog.String("id", idStr), slog.String("name", student.Name))
	}
}



