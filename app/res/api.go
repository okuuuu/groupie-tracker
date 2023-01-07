package res

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func makePetition(options map[string]interface{}, response interface{}) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	endpoint := ""

	for k, v := range options {
		switch v.(type) {
		case string:
			if k == "endpoint" {
				endpoint = v.(string)
			}

		case map[string]int:
			if k == "params" {
				// Modify the URL, i.e. the endpoint
				integer := strconv.FormatInt(int64(v.(map[string]int)["integer"]), 10)
				endpoint = endpoint + "/" + integer
				// Delete the parameters
				delete(options, "params")
			}
		}
	}

	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
