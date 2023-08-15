package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func GetResponse200(w http.ResponseWriter, r *http.Request) {
	intStatus := http.StatusOK
	resp := response{
		Status:  http.StatusText(intStatus),
		Message: http.StatusText(intStatus),
		Code:    intStatus,
	}

	render.Status(r, intStatus)
	render.JSON(w, r, resp)
}

func GetResponse400(w http.ResponseWriter, r *http.Request) {
	intStatus := http.StatusBadRequest
	resp := response{
		Status:  http.StatusText(intStatus),
		Message: http.StatusText(intStatus),
		Code:    intStatus,
	}

	render.Status(r, intStatus)
	render.JSON(w, r, resp)
}

func GetResponse500(w http.ResponseWriter, r *http.Request) {
	intStatus := http.StatusInternalServerError
	resp := response{
		Status:  http.StatusText(intStatus),
		Message: http.StatusText(intStatus),
		Code:    intStatus,
	}

	render.Status(r, intStatus)
	render.JSON(w, r, resp)
}
