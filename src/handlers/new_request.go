package handlers

import (
	"log"
	"net/http"
)

type SaveRequestUseCase interface {
	SaveRequest(req *http.Request) error
}

// SaveRequest saves request to database for statistics
func (i *Implementation) SaveRequest(w http.ResponseWriter, req *http.Request) {
	err := i.saveRequestUseCase.SaveRequest(req)
	if err != nil {
		log.Printf("save request %v\n", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
