package handler

import (
	"fmt"
	"net/http"

	"github.com/faizalom/go-api/internal/service"
)

type ExampleHandler struct {
	serviceA *service.ServiceA
	serviceB *service.ServiceB
}

func NewExampleHandler(sA *service.ServiceA, sB *service.ServiceB) *ExampleHandler {
	return &ExampleHandler{
		serviceA: sA,
		serviceB: sB,
	}
}

func (h *ExampleHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	dataA, errA := h.serviceA.DoWorkA()
	if errA != nil {
		http.Error(w, errA.Error(), http.StatusInternalServerError)
		return
	}

	dataB, errB := h.serviceB.DoWorkB()
	if errB != nil {
		http.Error(w, errB.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Combined Data: [%s] and [%s]", dataA, dataB)
}
