package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BuildURL constructs a URL from a base URL and query parameters.
//
// Parameters:
//   - baseURL: The base URL to which query parameters will be appended
//   - queryParams: A map of query parameter keys and values to be added to the URL
//
// Returns:
//   - string: The constructed URL with encoded query parameters
//   - error: An error if the base URL cannot be parsed
//
// Example:
//
//	url, err := net.BuildURL("https://api.example.com/users", map[string]string{
//		"page": "1",
//		"limit": "10",
//		"sort": "name",
//	})
//	// url will be "https://api.example.com/users?limit=10&page=1&sort=name"
func BuildURL(baseURL string, queryParams map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// GetJSON performs a GET request and unmarshals the JSON response into the provided interface.
//
// Parameters:
//   - url: The URL to send the GET request to
//   - target: A pointer to the struct or interface where the JSON response will be unmarshaled
//   - headers: A map of custom HTTP headers to include in the request (can be nil)
//
// Returns:
//   - error: An error if the request fails, the response status is not 2xx, or JSON unmarshaling fails
//
// Example:
//
//	type User struct {
//		ID   int    `json:"id"`
//		Name string `json:"name"`
//	}
//
//	var user User
//	err := net.GetJSON("https://api.example.com/users/1", &user, map[string]string{
//		"Authorization": "Bearer token123",
//	})
//	// user will contain the unmarshaled JSON response
func GetJSON(urlStr string, target any, headers map[string]string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, urlStr, http.NoBody)
	if err != nil {
		return err
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Always parse to get body content
	err = json.NewDecoder(resp.Body).Decode(target)

	if !IsSuccessStatusCode(resp.StatusCode) {
		return fmt.Errorf("error response from server: %d %s", resp.StatusCode, resp.Status)
	}

	return err
}

// PostJSON performs a POST request with JSON body and unmarshals the response into the provided interface.
//
// Parameters:
//   - url: The URL to send the POST request to
//   - body: The data to be marshaled to JSON and sent as the request body
//   - target: A pointer to the struct or interface where the JSON response will be unmarshaled
//   - headers: A map of custom HTTP headers to include in the request (can be nil)
//
// Returns:
//   - error: An error if JSON marshaling fails, the request fails, the response status is not 2xx, or JSON unmarshaling fails
//
// Example:
//
//	type LoginRequest struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//
//	type LoginResponse struct {
//		Token string `json:"token"`
//		User  struct {
//			ID   int    `json:"id"`
//			Name string `json:"name"`
//		} `json:"user"`
//	}
//
//	req := LoginRequest{Username: "user1", Password: "pass123"}
//	var resp LoginResponse
//	err := net.PostJSON("https://api.example.com/login", req, &resp, nil)
//	// resp will contain the unmarshaled JSON response with the token and user info
func PostJSON(urlStr string, body, target any, headers map[string]string) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(string(jsonBody)))
	if err != nil {
		return err
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Always parse to get body content
	err = json.NewDecoder(resp.Body).Decode(target)

	if !IsSuccessStatusCode(resp.StatusCode) {
		return fmt.Errorf("error response from server: %d %s", resp.StatusCode, resp.Status)
	}

	return err
}

// DownloadFile downloads a file from the specified URL and returns its contents as a byte slice.
//
// Parameters:
//   - url: The URL of the file to download
//   - timeout: The timeout for the HTTP request in seconds
//
// Returns:
//   - []byte: The contents of the downloaded file as a byte slice
//   - error: An error if the request fails, the response status is not 2xx, or reading the response body fails
//
// Example:
//
//	data, err := net.DownloadFile("https://example.com/files/document.pdf", 10)
//	if err != nil {
//		log.Fatalf("Failed to download file: %v", err)
//	}
//	// Save the downloaded data to a file
//	err = os.WriteFile("document.pdf", data, 0644)
func DownloadFile(urlStr string, timeout int) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if !IsSuccessStatusCode(resp.StatusCode) {
		return nil, fmt.Errorf("error response from server: %d %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// IsSuccessStatusCode checks if the HTTP status code is in the 2xx range (200-299),
// which indicates a successful response according to HTTP standards.
//
// Parameters:
//   - statusCode: The HTTP status code to check
//
// Returns:
//   - bool: true if the status code is between 200 and 299 (inclusive), false otherwise
//
// Example:
//
//	resp, err := http.Get("https://api.example.com/users")
//	if err != nil {
//		log.Fatalf("Request failed: %v", err)
//	}
//	if !net.IsSuccessStatusCode(resp.StatusCode) {
//		log.Fatalf("Request failed with status: %d %s", resp.StatusCode, resp.Status)
//	}
//	// Process successful response
func IsSuccessStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

// PutJSON performs a PUT request with JSON body and unmarshals the response into the provided interface.
// This is typically used for updating existing resources on a server.
//
// Parameters:
//   - url: The URL to send the PUT request to
//   - body: The data to be marshaled to JSON and sent as the request body
//   - target: A pointer to the struct or interface where the JSON response will be unmarshaled
//   - headers: A map of custom HTTP headers to include in the request (can be nil)
//
// Returns:
//   - error: An error if JSON marshaling fails, the request fails, the response status is not 2xx, or JSON unmarshaling fails
//
// Example:
//
//	type UserUpdate struct {
//		Name  string `json:"name"`
//		Email string `json:"email"`
//	}
//
//	type UserResponse struct {
//		ID        int    `json:"id"`
//		Name      string `json:"name"`
//		Email     string `json:"email"`
//		UpdatedAt string `json:"updated_at"`
//	}
//
//	update := UserUpdate{Name: "John Smith", Email: "john.smith@example.com"}
//	var updatedUser UserResponse
//	err := net.PutJSON("https://api.example.com/users/123", update, &updatedUser, nil)
//	// updatedUser will contain the updated user information returned from the server
func PutJSON(urlStr string, body, target any, headers map[string]string) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodPut, urlStr, strings.NewReader(string(jsonBody)))
	if err != nil {
		return err
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Always parse to get body content
	err = json.NewDecoder(resp.Body).Decode(target)

	if !IsSuccessStatusCode(resp.StatusCode) {
		return fmt.Errorf("error response from server: %d %s", resp.StatusCode, resp.Status)
	}

	return err
}

// DeleteJSON performs a DELETE request and unmarshals the response into the provided interface.
// This is typically used for deleting resources on a server.
//
// Parameters:
//   - url: The URL to send the DELETE request to
//   - target: A pointer to the struct or interface where the JSON response will be unmarshaled (can be nil if no response body is expected)
//   - headers: A map of custom HTTP headers to include in the request (can be nil)
//
// Returns:
//   - error: An error if the request fails, the response status is not 2xx, or JSON unmarshaling fails
//
// Example:
//
//	type DeleteResponse struct {
//		Success bool   `json:"success"`
//		Message string `json:"message"`
//	}
//
//	var resp DeleteResponse
//	err := net.DeleteJSON("https://api.example.com/users/123", &resp, map[string]string{
//		"Authorization": "Bearer token123",
//	})
//	// resp will contain the response from the server confirming the deletion
func DeleteJSON(urlStr string, target any, headers map[string]string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, urlStr, http.NoBody)
	if err != nil {
		return err
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Always parse to get body content
	err = json.NewDecoder(resp.Body).Decode(target)

	if !IsSuccessStatusCode(resp.StatusCode) {
		return fmt.Errorf("error response from server: %d %s", resp.StatusCode, resp.Status)
	}

	return err
}

// UploadFile uploads a file to the specified URL using multipart/form-data encoding.
// This is commonly used for file uploads to web servers.
//
// Parameters:
//   - url: The URL to upload the file to
//   - fieldName: The name of the form field that will contain the file data
//   - filePath: The path to the file on the local filesystem
//   - additionalFields: A map of additional form fields to include in the request (can be nil)
//   - headers: A map of custom HTTP headers to include in the request (can be nil)
//
// Returns:
//   - *http.Response: The HTTP response from the server
//   - error: An error if opening the file fails, creating the request fails, or the request fails
//
// Example:
//
//	// Upload a profile picture with additional user information
//	resp, err := net.UploadFile(
//		"https://api.example.com/upload",
//		"profile_picture",
//		"/path/to/image.jpg",
//		map[string]string{
//			"user_id": "123",
//			"description": "Profile picture",
//		},
//		map[string]string{
//			"Authorization": "Bearer token123",
//		},
//	)
//	if err != nil {
//		log.Fatalf("Upload failed: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// Check if upload was successful
//	if !net.IsSuccessStatusCode(resp.StatusCode) {
//		log.Fatalf("Upload failed with status: %d %s", resp.StatusCode, resp.Status)
//	}
func UploadFile(urlStr, fieldName, filePath string, additionalFields, headers map[string]string) (*http.Response, error) {
	filePath = filepath.Clean(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Add additional form fields
	for key, value := range additionalFields {
		err = writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, err
	}

	// Set content type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30, // Longer timeout for file uploads
	}

	return client.Do(req)
}

// CreateHTTPClient creates an HTTP client with custom timeout and transport options.
// This allows for fine-tuning connection pooling and timeout settings for HTTP requests.
//
// Parameters:
//   - timeout: The maximum time to wait for a complete response
//   - maxIdleConns: The maximum number of idle (keep-alive) connections across all hosts
//   - maxIdleConnsPerHost: The maximum number of idle (keep-alive) connections per host
//   - maxConnsPerHost: The maximum number of connections per host
//
// Returns:
//   - *http.Client: A configured HTTP client with the specified settings
//
// Example:
//
//	// Create a client with a 30-second timeout and custom connection pool settings
//	client := net.CreateHTTPClient(
//		30*time.Second,
//		100,  // maxIdleConns
//		10,   // maxIdleConnsPerHost
//		100,  // maxConnsPerHost
//	)
//
//	// Use the client for requests
//	resp, err := client.Get("https://api.example.com/data")
func CreateHTTPClient(timeout time.Duration, maxIdleConns, maxIdleConnsPerHost, maxConnsPerHost int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        maxIdleConns,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		MaxConnsPerHost:     maxConnsPerHost,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

// ParseQueryParams parses URL query parameters into a map of key-value pairs.
// This function converts a URL query string (e.g., "name=John&age=30") into a map.
// If a parameter appears multiple times, only the first value is kept.
//
// Parameters:
//   - queryString: The URL query string to parse (without the leading '?')
//
// Returns:
//   - map[string]string: A map containing the parsed query parameters
//   - error: An error if the query string cannot be parsed
//
// Example:
//
//	// Parse query parameters from a URL
//	url := "https://example.com/search?q=golang&page=1&sort=desc"
//	parsedURL, _ := url.Parse(url)
//	params, err := net.ParseQueryParams(parsedURL.RawQuery)
//	if err != nil {
//		log.Fatalf("Failed to parse query params: %v", err)
//	}
//	// params will be: {"q": "golang", "page": "1", "sort": "desc"}
//
//	// Or parse a query string directly
//	params, err := net.ParseQueryParams("q=golang&page=1&sort=desc")
//	// params will be: {"q": "golang", "page": "1", "sort": "desc"}
func ParseQueryParams(queryString string) (map[string]string, error) {
	values, err := url.ParseQuery(queryString)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for key, vals := range values {
		if len(vals) > 0 {
			result[key] = vals[0]
		}
	}

	return result, nil
}
