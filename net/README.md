# net - HTTP Utility Functions for Go

The `net` package provides a comprehensive set of utility functions for working with HTTP requests, URLs, and file transfers in Go. It simplifies common HTTP operations like making API requests, handling JSON data, uploading and downloading files.

## Installation

```bash
go get github.com/gflydev/utils/net
```

## Usage

```go
import "github.com/gflydev/utils/net"
```

## Functions

### BuildURL

Constructs a URL from a base URL and query parameters.

```go
// Build a URL with query parameters
url, err := net.BuildURL("https://api.example.com/users", map[string]string{
    "page": "1",
    "limit": "10",
    "sort": "name",
})
// url will be "https://api.example.com/users?limit=10&page=1&sort=name"

// Base URL with existing query parameters
url, err := net.BuildURL("https://example.com?existing=param", map[string]string{
    "key": "value",
})
// url will be "https://example.com?existing=param&key=value"

// Empty query parameters
url, err := net.BuildURL("https://example.com", nil)
// url will be "https://example.com"
```

### GetJSON

Performs a GET request and unmarshals the JSON response into the provided interface.

```go
// Define a struct to hold the response
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Make a GET request with custom headers
var user User
err := net.GetJSON("https://api.example.com/users/1", &user, map[string]string{
    "Authorization": "Bearer token123",
    "X-API-Key": "abc123",
})
// user will contain the unmarshaled JSON response

// Make a GET request without custom headers
var users []User
err := net.GetJSON("https://api.example.com/users", &users, nil)
// users will contain the unmarshaled JSON response
```

### PostJSON

Performs a POST request with JSON body and unmarshals the response into the provided interface.

```go
// Define request and response structs
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
    User  struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    } `json:"user"`
}

// Make a POST request
req := LoginRequest{Username: "user1", Password: "pass123"}
var resp LoginResponse
err := net.PostJSON("https://api.example.com/login", req, &resp, map[string]string{
    "X-API-Key": "abc123",
})
// resp will contain the unmarshaled JSON response with the token and user info
```

### PutJSON

Performs a PUT request with JSON body and unmarshals the response into the provided interface.
This is typically used for updating existing resources on a server.

Parameters:
- `url`: The URL to send the PUT request to
- `body`: The data to be marshaled to JSON and sent as the request body
- `target`: A pointer to the struct or interface where the JSON response will be unmarshaled
- `headers`: A map of custom HTTP headers to include in the request (can be nil)

```go
// Define request and response structs
type UserUpdate struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

type UpdateResponse struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

// Make a PUT request
update := UserUpdate{Name: "test", Value: 42}
var response UpdateResponse
err := net.PutJSON("https://api.example.com/users/123", update, &response, map[string]string{
    "X-Test-Header": "test-value",
})
// response will contain the unmarshaled JSON response
```

### DeleteJSON

Performs a DELETE request and unmarshals the response into the provided interface.
This is typically used for deleting resources on a server.

Parameters:
- `url`: The URL to send the DELETE request to
- `target`: A pointer to the struct or interface where the JSON response will be unmarshaled (can be nil if no response body is expected)
- `headers`: A map of custom HTTP headers to include in the request (can be nil)

```go
// Define a response struct
type DeleteResponse struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

// Make a DELETE request
var response DeleteResponse
err := net.DeleteJSON("https://api.example.com/users/123", &response, map[string]string{
    "X-Test-Header": "test-value",
})
// response will contain the unmarshaled JSON response confirming the deletion
```

### DownloadFile

Downloads a file from the specified URL and returns its contents as a byte slice. The function takes a URL and a timeout value in seconds for the HTTP request.

Parameters:
- `url`: The URL of the file to download
- `timeout`: The timeout for the HTTP request in seconds

```go
// Download a file with a 10-second timeout
data, err := net.DownloadFile("https://example.com/files/document.pdf", 10)
if err != nil {
    log.Fatalf("Failed to download file: %v", err)
}

// Save the downloaded data to a file
err = os.WriteFile("document.pdf", data, 0644)
```

### IsSuccessStatusCode

Checks if the HTTP status code is in the 2xx range (200-299), which indicates a successful response according to HTTP standards.

Parameters:
- `statusCode`: The HTTP status code to check

```go
// Check if a status code indicates success
if net.IsSuccessStatusCode(http.StatusOK) {
    // Handle successful response (200 OK)
}

if net.IsSuccessStatusCode(http.StatusCreated) {
    // Handle successful response (201 Created)
}

if !net.IsSuccessStatusCode(http.StatusBadRequest) {
    // Handle error response (400 Bad Request)
}

if !net.IsSuccessStatusCode(http.StatusInternalServerError) {
    // Handle error response (500 Internal Server Error)
}
```

### UploadFile

Uploads a file to the specified URL using multipart/form-data encoding.
This is commonly used for file uploads to web servers.

Parameters:
- `url`: The URL to upload the file to
- `fieldName`: The name of the form field that will contain the file data
- `filePath`: The path to the file on the local filesystem
- `additionalFields`: A map of additional form fields to include in the request (can be nil)
- `headers`: A map of custom HTTP headers to include in the request (can be nil)

```go
// Upload a file with additional form fields and headers
resp, err := net.UploadFile(
    "https://api.example.com/upload",
    "file",                // form field name for the file
    "/path/to/image.jpg",  // path to the file
    map[string]string{     // additional form fields
        "field1": "value1",
    },
    map[string]string{     // custom headers
        "X-Test-Header": "test-value",
    },
)
if err != nil {
    log.Fatalf("Upload failed: %v", err)
}
defer resp.Body.Close()

// Check if upload was successful
if !net.IsSuccessStatusCode(resp.StatusCode) {
    log.Fatalf("Upload failed with status: %d %s", resp.StatusCode, resp.Status)
}
```

### CreateHTTPClient

Creates an HTTP client with custom timeout and transport options.
This allows for fine-tuning connection pooling and timeout settings for HTTP requests.

Parameters:
- `timeout`: The maximum time to wait for a complete response
- `maxIdleConns`: The maximum number of idle (keep-alive) connections across all hosts
- `maxIdleConnsPerHost`: The maximum number of idle (keep-alive) connections per host
- `maxConnsPerHost`: The maximum number of connections per host

```go
// Create a client with a 30-second timeout and custom connection pool settings
client := net.CreateHTTPClient(
    30*time.Second,  // timeout
    100,             // maxIdleConns
    10,              // maxIdleConnsPerHost
    100,             // maxConnsPerHost
)

// Use the client for requests
resp, err := client.Get("https://api.example.com/data")
```

### ParseQueryParams

Parses URL query parameters into a map of key-value pairs.
This function converts a URL query string (e.g., "name=John&age=30") into a map.
If a parameter appears multiple times, only the first value is kept.

Parameters:
- `queryString`: The URL query string to parse (without the leading '?')

Returns:
- `map[string]string`: A map containing the parsed query parameters
- `error`: An error if the query string cannot be parsed

```go
// Parse query parameters from a URL
url := "https://example.com/search?q=golang&page=1&sort=desc"
parsedURL, _ := url.Parse(url)
params, err := net.ParseQueryParams(parsedURL.RawQuery)
if err != nil {
    log.Fatalf("Failed to parse query params: %v", err)
}
// params will be: {"q": "golang", "page": "1", "sort": "desc"}

// Or parse a query string directly
params, err := net.ParseQueryParams("q=golang&page=1&sort=desc")
// params will be: {"q": "golang", "page": "1", "sort": "desc"}

// Parameter with multiple values (takes first value)
params, err := net.ParseQueryParams("key=value1&key=value2")
// params will be: {"key": "value1"}
```

## License

This package is licensed under the MIT License - see the LICENSE file for details.
