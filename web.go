package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Check if the status code passed by the user is valid
func IsValidHTTPStatusCode(code int) bool {
	return code >= 100 && code <= 599
}

// Respond to k8s health probes
// Additional functionality of allowing user to pass arbitrary status codes for testing
func HTTPProbe(w http.ResponseWriter, r *http.Request) {
	// Error if the request isn't Get or Post
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Only GET and POST requests are allowed")
		return
	}

	// Send 200 for /
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" || path == "healthz" {
		// Default response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello! Pass a status code in the URL path (e.g., /200 or /404)")
		return
	}

	statusCode, _ := strconv.Atoi(path)
	_, err := strconv.Atoi(path)

	// Error if string can't be converted to integer
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "That is not a valid status code or integer.")
		return
	}

	// Error if integer is out of validity range
	err = nil
	if IsValidHTTPStatusCode(statusCode) == false {
		err = errors.New("That integer is not a valid status code")
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Errorf("Error processing data: %w", err))
		return
	}

	// Write value back to header and body
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%d", statusCode)
}
