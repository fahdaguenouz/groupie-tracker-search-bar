package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	// Check if the URL path begins with "/assets/css/" or "/assets/imgs/"
    if strings.HasPrefix(r.URL.Path, "/assets/css/") {
        handleCSSAssets(w, r)
    } else if strings.HasPrefix(r.URL.Path, "/assets/imgs/") {
        handleImageAssets(w, r)
    } else {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
}

func handleCSSAssets(w http.ResponseWriter, r *http.Request) {
    requestedFile := strings.TrimPrefix(r.URL.Path, "/assets/css/")

    // Check for path traversal and ensure it's a single file request
    if strings.Contains(requestedFile, "/") || requestedFile == "" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

    // Construct the file path
    filePath := filepath.Join("assets", "css", requestedFile)

    // Verify the file exists and is not a directory
    fileInfo, err := os.Stat(filePath)
    if os.IsNotExist(err) || fileInfo.IsDir() {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

    // Serve the CSS file if it exists
    http.ServeFile(w, r, filePath)
}

func handleImageAssets(w http.ResponseWriter, r *http.Request) {
    requestedFile := strings.TrimPrefix(r.URL.Path, "/assets/imgs/")

    // Check for path traversal and ensure it's a single file request
    if strings.Contains(requestedFile, "/") || requestedFile == "" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

    // Construct the file path
    filePath := filepath.Join("assets", "imgs", requestedFile)

    // Verify the file exists and is not a directory
    fileInfo, err := os.Stat(filePath)
    if os.IsNotExist(err) || fileInfo.IsDir() {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }

    // Serve the image file if it exists
    http.ServeFile(w, r, filePath)
}