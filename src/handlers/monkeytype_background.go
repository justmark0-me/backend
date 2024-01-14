package handlers

import (
	"net/http"
)

// GetMonkeytypeBackground returns static image that I use for monkeytype background
func (i *Implementation) GetMonkeytypeBackground(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	picturePath := "/app/files/monkeytype_background.jpg"
	http.ServeFile(w, req, picturePath)
}
