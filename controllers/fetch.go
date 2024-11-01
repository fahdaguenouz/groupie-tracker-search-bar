package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchData is a generic function to fetch data from a given URL.
func FetchData(url string, result any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
