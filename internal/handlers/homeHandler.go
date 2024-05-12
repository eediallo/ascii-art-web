package handlers

import "net/http"

// MainPageHandler handles requests to the main page
func MainPageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return nil
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	return MainPageGetHandler(w, r)
}
