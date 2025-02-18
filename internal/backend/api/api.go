package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchData[T any](url string, data *T) error {
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
