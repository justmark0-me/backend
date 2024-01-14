package handlers

import (
	"fmt"
	"net/http"
)

// GetRootPage index page handler
func (i *Implementation) GetRootPage(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "There is no documentation of API :'(")
	if err != nil {
		return
	}
}
