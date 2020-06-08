// Package spa provides a Single Page Application server.
package spa

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Server is an http.Handler which serves static files from Dir,
// and reverts to serving index.html for any missing files.
//
// Requests which contain an extension — such as /favicon.ico — are greated as
// requests for a file, yielding 404 instead of serving /index.html.
type Server struct {
	Dir string
}

// ServeHTTP implementation.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := filepath.Join(s.Dir, path.Clean(r.URL.Path))

	info, err := os.Stat(name)

	if (os.IsNotExist(err) || !info.Mode().IsRegular()) && path.Ext(name) == "" {
		name = filepath.Join(s.Dir, "/index.html")
	}

	http.ServeFile(w, r, name)
}
