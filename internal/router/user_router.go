package router

import (
	"net/http"
)

func UserRouter(h *Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", protected(h.ListUsers))
	mux.Handle("POST /{$}", protected(h.CreateUser))
	mux.Handle("GET /{id}", protected(h.GetUserByID))
	mux.Handle("PUT /{id}", protected(h.UpdateUser))
	mux.Handle("DELETE /{id}", protected(h.DeleteUser))
	return mux
}
