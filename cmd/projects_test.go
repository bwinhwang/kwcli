package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjects(t *testing.T) {
	mockResponse := `[{"id": "test-project", "name": "Example"}]` // Sample JSON

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer ts.Close()

	// Temporarily modify your KWClient or command to point to the test server's URL

	// ... rest of your test setup ...

	// Run your projects command
	// ...

	// Assertions (check if the projects array is parsed correctly)
}
