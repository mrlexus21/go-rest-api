package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest_api/internal/handlers"
	"rest_api/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, к *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of user"))
}

func (h *handler) CreateUser(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create of user"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is user by uuid"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update of user"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update of user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete of user"))
}
