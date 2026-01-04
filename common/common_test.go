package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ... other imports ...

func Test_KWResponse_Validate(t *testing.T) {
	testCases := []struct {
		name          string
		response      KWResponse
		expectedError bool
	}{
		{"missing status", KWResponse{Status: 0, Message: "some message"}, true},
		{"missing message", KWResponse{Status: 1, Message: ""}, true},
		{"valid response", KWResponse{Status: 1, Message: "success"}, false},
		// ... add more cases ...
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.response.Validate()
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func Test_Execute_Success(t *testing.T) {
	// Setup a mock HTTP server
	// Klocwork API 成功时返回数据行（不是 JSON），失败时返回 JSON 格式的 KWResponse
	// 所以成功响应应该是非 JSON 格式的数据行
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"id": "123", "name": "test_project"}`) // 数据行，不是 KWResponse
	}))
	defer ts.Close()

	// Create a KWClient pointing to the mock server
	client := NewKWClient(ts.URL, "user", "token")

	// Execute with some data
	lines, err := client.Execute(map[string]interface{}{"action": "test"})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(lines) == 0 {
		t.Errorf("Expected data lines, got empty")
	}
}

func Test_Execute_APIError(t *testing.T) {
	// Klocwork API 返回 JSON 格式的 KWResponse 表示错误
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status": 1, "message": "project not found"}`)
	}))
	defer ts.Close()

	client := NewKWClient(ts.URL, "user", "token")

	_, err := client.Execute(map[string]interface{}{"action": "test"})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// ... Add more tests for Execute with different responses, errors, etc.
// ... Tests for constructFields ...
