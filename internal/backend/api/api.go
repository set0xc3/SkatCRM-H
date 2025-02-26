package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostData[T any](url string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080"+url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add: %s", resp.Status)
	}

	return nil
}

func FetchData[T any](url string, data T) error {
	resp, err := http.Get("http://localhost:8080" + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

func FetchOne[T any](params string) (data T, err error) {
	url := params
	err = FetchData(url, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func Fetch[T any](params string) (data []T, err error) {
	url := params
	err = FetchData(url, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
