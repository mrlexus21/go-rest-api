package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest_api/internal/apperror"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, request *http.Request) error {
	/*w.WriteHeader(200)
	w.Write([]byte("this is list of user"))

	return nil*/
	return apperror.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, request *http.Request) error {
	/*w.WriteHeader(201)
	w.Write([]byte("this is create of user"))

	return nil*/
	return fmt.Errorf("this is API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, request *http.Request) error {
	/*w.WriteHeader(200)
	w.Write([]byte("this is user by uuid"))

	return nil*/
	return apperror.NewAppError(nil, "test", "test", "t123")
}

func (h *handler) UpdateUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is update of user"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update of user"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete of user"))

	return nil
}
