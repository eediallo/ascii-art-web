package servers

import (
	//"fmt"
	"net/http"

	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// Run starts the HTTP server.
func (s *server) Run() {
	//Call serveStaticFiles function to get the static files handler
	// errFileServer := s.serveStaticFiles(staticFilesAddress)

	// if errFileServer != nil {
	// 	fmt.Printf("file directory does not exit %s", errFileServer)
	// 	return
	// }
	
	mux := http.NewServeMux()
	// Register the static files handler with the server
	// mux.Handle("/static/", staticFilesHandler)

	mux.HandleFunc(homePagePath, makeHTTPHandlerFunc(handlers.MainPageHandler))
	mux.HandleFunc(asciiArtPage, makeHTTPHandlerFunc(handlers.AsciiArtHandler))
	err := http.ListenAndServe(s.listener, mux)
	if err != nil {
		return
	}
}
