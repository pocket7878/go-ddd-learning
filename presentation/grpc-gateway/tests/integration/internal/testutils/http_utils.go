package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostJSON(url string, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	httpClient := &http.Client{}
	response, error := httpClient.Do(request)
	if error != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}
