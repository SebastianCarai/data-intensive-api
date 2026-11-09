package utils

import (
	structs "data-intensive-api/cmd/internals/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Add required JSON validation
// https://dev.to/sanjaysinghrajpoot/adding-json-validation-to-structs-in-go-lang-4mg8

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, t interface{}) error{

	if ct := r.Header.Get("Content-Type"); ct != ""{
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			return structs.BadRequest{
				Status: http.StatusUnsupportedMediaType, 
				Msg: "Content-Type header is not application/json",
			}
        }
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&t)

    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
        var maxBytesError *http.MaxBytesError

        switch {
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return structs.BadRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := "Request body contains badly-formed JSON"
			return structs.BadRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return structs.BadRequest{Status: http.StatusBadRequest, Msg: msg}

        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return structs.BadRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
			return structs.BadRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.As(err, &maxBytesError):
            msg := fmt.Sprintf("Request body must not be larger than %d bytes", maxBytesError.Limit)
			return structs.BadRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

        default:
            return structs.BadRequest{Status: 500, Msg: "Internal server error"}
        }
    }

	return nil
}