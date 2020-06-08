package spa_test

import (
	"net/http/httptest"
	"testing"

	"github.com/tj/assert"
	"github.com/tj/spa"
)

// Test server.
func TestServer(t *testing.T) {
	h := spa.Server{
		Dir: "example",
	}

	t.Run("with index", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "<h1>Hello</h1>")
	})

	t.Run("with missing file", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/some/random/stuff", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "<h1>Hello</h1>")
	})

	t.Run("with file outside of dir", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/../go.mod", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "invalid URL path\n")
	})

	t.Run("with existing file", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/style.css", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "text/css; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "16px Helvetica")
	})

	t.Run("with nested file", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/nested/some.txt", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "hello")
	})

	t.Run("with directory", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/nested", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "<h1>Hello</h1>")
	})

	t.Run("with extension", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/favicon.ico", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "404 page not found\n")
	})
}
