package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		queryParams map[string]string
		expected    string
		expectError bool
	}{
		{
			name:        "Empty base URL",
			baseURL:     "",
			queryParams: map[string]string{"key": "value"},
			expected:    "?key=value",
			expectError: false,
		},
		{
			name:        "Invalid base URL",
			baseURL:     "://invalid-url",
			queryParams: map[string]string{"key": "value"},
			expectError: true,
		},
		{
			name:        "No query params",
			baseURL:     "https://example.com",
			queryParams: nil,
			expected:    "https://example.com",
			expectError: false,
		},
		{
			name:    "With query params",
			baseURL: "https://example.com",
			queryParams: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected:    "https://example.com?key1=value1&key2=value2",
			expectError: false,
		},
		{
			name:    "Base URL with existing query",
			baseURL: "https://example.com?existing=param",
			queryParams: map[string]string{
				"key": "value",
			},
			expected:    "https://example.com?existing=param&key=value",
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := BuildURL(test.baseURL, test.queryParams)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("Expected URL %q, got %q", test.expected, result)
				}
			}
		})
	}
}

func TestIsSuccessStatusCode(t *testing.T) {
	tests := []struct {
		statusCode int
		expected   bool
	}{
		{http.StatusOK, true},
		{http.StatusCreated, true},
		{http.StatusAccepted, true},
		{http.StatusNoContent, true},
		{http.StatusBadRequest, false},
		{http.StatusUnauthorized, false},
		{http.StatusInternalServerError, false},
		{0, false},
		{99, false},
		{600, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("StatusCode_%d", test.statusCode), func(t *testing.T) {
			result := IsSuccessStatusCode(test.statusCode)
			if result != test.expected {
				t.Errorf("IsSuccessStatusCode(%d) = %v, expected %v", test.statusCode, result, test.expected)
			}
		})
	}
}

func TestParseQueryParams(t *testing.T) {
	tests := []struct {
		name        string
		queryString string
		expected    map[string]string
		expectError bool
	}{
		{
			name:        "Empty query string",
			queryString: "",
			expected:    map[string]string{},
			expectError: false,
		},
		{
			name:        "Invalid query string",
			queryString: "%invalid",
			expectError: true,
		},
		{
			name:        "Single parameter",
			queryString: "key=value",
			expected:    map[string]string{"key": "value"},
			expectError: false,
		},
		{
			name:        "Multiple parameters",
			queryString: "key1=value1&key2=value2",
			expected:    map[string]string{"key1": "value1", "key2": "value2"},
			expectError: false,
		},
		{
			name:        "Parameter with multiple values (should take first)",
			queryString: "key=value1&key=value2",
			expected:    map[string]string{"key": "value1"},
			expectError: false,
		},
		{
			name:        "Parameter without value",
			queryString: "key=",
			expected:    map[string]string{"key": ""},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseQueryParams(test.queryString)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Expected %v, got %v", test.expected, result)
				}
			}
		})
	}
}

func TestCreateHTTPClient(t *testing.T) {
	tests := []struct {
		name                string
		timeout             time.Duration
		maxIdleConns        int
		maxIdleConnsPerHost int
		maxConnsPerHost     int
	}{
		{
			name:                "Default values",
			timeout:             30 * time.Second,
			maxIdleConns:        10,
			maxIdleConnsPerHost: 5,
			maxConnsPerHost:     100,
		},
		{
			name:                "Custom values",
			timeout:             5 * time.Second,
			maxIdleConns:        20,
			maxIdleConnsPerHost: 10,
			maxConnsPerHost:     200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := CreateHTTPClient(test.timeout, test.maxIdleConns, test.maxIdleConnsPerHost, test.maxConnsPerHost)

			if client == nil {
				t.Fatal("Expected non-nil HTTP client")
			}

			if client.Timeout != test.timeout {
				t.Errorf("Expected timeout %v, got %v", test.timeout, client.Timeout)
			}

			transport, ok := client.Transport.(*http.Transport)
			if !ok {
				t.Fatal("Expected *http.Transport")
			}

			if transport.MaxIdleConns != test.maxIdleConns {
				t.Errorf("Expected MaxIdleConns %d, got %d", test.maxIdleConns, transport.MaxIdleConns)
			}

			if transport.MaxIdleConnsPerHost != test.maxIdleConnsPerHost {
				t.Errorf("Expected MaxIdleConnsPerHost %d, got %d", test.maxIdleConnsPerHost, transport.MaxIdleConnsPerHost)
			}

			if transport.MaxConnsPerHost != test.maxConnsPerHost {
				t.Errorf("Expected MaxConnsPerHost %d, got %d", test.maxConnsPerHost, transport.MaxConnsPerHost)
			}
		})
	}
}

// Mock HTTP server for testing HTTP request functions
func setupMockServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	server := httptest.NewServer(handler)
	t.Cleanup(func() {
		server.Close()
	})
	return server
}

func TestGetJSON(t *testing.T) {
	type testResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	// Test successful request
	t.Run("Successful request", func(t *testing.T) {
		expected := testResponse{
			Message: "Hello, World!",
			Status:  "success",
		}

		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("Expected GET request, got %s", r.Method)
			}

			// Check headers
			if r.Header.Get("X-Test-Header") != "test-value" {
				t.Errorf("Expected header X-Test-Header: test-value, got %s", r.Header.Get("X-Test-Header"))
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(expected)
		})

		var result testResponse
		err := GetJSON(server.URL, &result, map[string]string{"X-Test-Header": "test-value"})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	// Test non-success status code
	t.Run("Non-success status code", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"status":  "failed", "message": "Bad request"}`))
		})

		var result testResponse
		err := GetJSON(server.URL, &result, nil)

		if err == nil {
			t.Error("Expected error for non-success status code, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		var result testResponse
		err := GetJSON("invalid-url", &result, nil)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})

	// Test invalid JSON response
	t.Run("Invalid JSON response", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`invalid json`))
		})

		var result testResponse
		err := GetJSON(server.URL, &result, nil)

		if err == nil {
			t.Error("Expected error for invalid JSON, got nil")
		}
	})
}

func TestPostJSON(t *testing.T) {
	type testRequest struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	type testResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	// Test successful request
	t.Run("Successful request", func(t *testing.T) {
		request := testRequest{
			Name:  "test",
			Value: 42,
		}

		expected := testResponse{
			Message: "Created",
			Status:  "success",
		}

		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST request, got %s", r.Method)
			}

			// Check headers
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
			}

			if r.Header.Get("X-Test-Header") != "test-value" {
				t.Errorf("Expected header X-Test-Header: test-value, got %s", r.Header.Get("X-Test-Header"))
			}

			// Check request body
			var receivedRequest testRequest
			if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
				t.Fatalf("Failed to decode request body: %v", err)
			}

			if !reflect.DeepEqual(receivedRequest, request) {
				t.Errorf("Expected request body %+v, got %+v", request, receivedRequest)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(expected)
		})

		var result testResponse
		err := PostJSON(server.URL, request, &result, map[string]string{"X-Test-Header": "test-value"})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	// Test non-success status code
	t.Run("Non-success status code", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Bad request"}`))
		})

		var result testResponse
		err := PostJSON(server.URL, testRequest{}, &result, nil)

		if err == nil {
			t.Error("Expected error for non-success status code, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		var result testResponse
		err := PostJSON("invalid-url", testRequest{}, &result, nil)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})
}

func TestPutJSON(t *testing.T) {
	type testRequest struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	type testResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	// Test successful request
	t.Run("Successful request", func(t *testing.T) {
		request := testRequest{
			Name:  "test",
			Value: 42,
		}

		expected := testResponse{
			Message: "Updated",
			Status:  "success",
		}

		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("Expected PUT request, got %s", r.Method)
			}

			// Check headers
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
			}

			if r.Header.Get("X-Test-Header") != "test-value" {
				t.Errorf("Expected header X-Test-Header: test-value, got %s", r.Header.Get("X-Test-Header"))
			}

			// Check request body
			var receivedRequest testRequest
			if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
				t.Fatalf("Failed to decode request body: %v", err)
			}

			if !reflect.DeepEqual(receivedRequest, request) {
				t.Errorf("Expected request body %+v, got %+v", request, receivedRequest)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(expected)
		})

		var result testResponse
		err := PutJSON(server.URL, request, &result, map[string]string{"X-Test-Header": "test-value"})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	// Test non-success status code
	t.Run("Non-success status code", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Bad request"}`))
		})

		var result testResponse
		err := PutJSON(server.URL, testRequest{}, &result, nil)

		if err == nil {
			t.Error("Expected error for non-success status code, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		var result testResponse
		err := PutJSON("invalid-url", testRequest{}, &result, nil)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})
}

func TestDeleteJSON(t *testing.T) {
	type testResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	// Test successful request
	t.Run("Successful request", func(t *testing.T) {
		expected := testResponse{
			Message: "Deleted",
			Status:  "success",
		}

		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("Expected DELETE request, got %s", r.Method)
			}

			// Check headers
			if r.Header.Get("X-Test-Header") != "test-value" {
				t.Errorf("Expected header X-Test-Header: test-value, got %s", r.Header.Get("X-Test-Header"))
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(expected)
		})

		var result testResponse
		err := DeleteJSON(server.URL, &result, map[string]string{"X-Test-Header": "test-value"})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	// Test non-success status code
	t.Run("Non-success status code", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Bad request"}`))
		})

		var result testResponse
		err := DeleteJSON(server.URL, &result, nil)

		if err == nil {
			t.Error("Expected error for non-success status code, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		var result testResponse
		err := DeleteJSON("invalid-url", &result, nil)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})
}

func TestDownloadFile(t *testing.T) {
	// Test successful download
	t.Run("Successful download", func(t *testing.T) {
		expected := []byte("file content")

		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("Expected GET request, got %s", r.Method)
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(expected)
		})

		result, err := DownloadFile(server.URL, 10)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test non-success status code
	t.Run("Non-success status code", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		_, err := DownloadFile(server.URL, 10)

		if err == nil {
			t.Error("Expected error for non-success status code, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		_, err := DownloadFile("invalid-url", 10)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})
}

func TestUploadFile(t *testing.T) {
	// Create a temporary file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	fileContent := []byte("test file content")

	err := os.WriteFile(tempFile, fileContent, 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test successful upload
	t.Run("Successful upload", func(t *testing.T) {
		server := setupMockServer(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST request, got %s", r.Method)
			}

			// Check that it's a multipart form
			if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
				t.Errorf("Expected multipart/form-data Content-Type, got %s", r.Header.Get("Content-Type"))
			}

			// Check headers
			if r.Header.Get("X-Test-Header") != "test-value" {
				t.Errorf("Expected header X-Test-Header: test-value, got %s", r.Header.Get("X-Test-Header"))
			}

			// Parse the multipart form
			err := r.ParseMultipartForm(10 << 20) // 10 MB
			if err != nil {
				t.Fatalf("Failed to parse multipart form: %v", err)
			}

			// Check form fields
			if r.FormValue("field1") != "value1" {
				t.Errorf("Expected field1=value1, got field1=%s", r.FormValue("field1"))
			}

			// Check file
			file, header, err := r.FormFile("file")
			if err != nil {
				t.Fatalf("Failed to get uploaded file: %v", err)
			}
			defer file.Close()

			if header.Filename != filepath.Base(tempFile) {
				t.Errorf("Expected filename %s, got %s", filepath.Base(tempFile), header.Filename)
			}

			uploadedContent := make([]byte, header.Size)
			_, err = file.Read(uploadedContent)
			if err != nil {
				t.Fatalf("Failed to read uploaded file: %v", err)
			}

			if !reflect.DeepEqual(uploadedContent, fileContent) {
				t.Errorf("Expected file content %v, got %v", fileContent, uploadedContent)
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Upload successful"))
		})

		additionalFields := map[string]string{
			"field1": "value1",
		}

		headers := map[string]string{
			"X-Test-Header": "test-value",
		}

		resp, err := UploadFile(server.URL, "file", tempFile, additionalFields, headers)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
	})

	// Test non-existent file
	t.Run("Non-existent file", func(t *testing.T) {
		_, err := UploadFile("http://example.com", "file", "non-existent-file.txt", nil, nil)

		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}
	})

	// Test invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		_, err := UploadFile("invalid-url", "file", tempFile, nil, nil)

		if err == nil {
			t.Error("Expected error for invalid URL, got nil")
		}
	})
}
