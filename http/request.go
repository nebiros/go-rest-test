package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nebiros/go-rest-test/config"
)

func MakeRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(config.Data.Ws.Auth.Username, config.Data.Ws.Auth.Password)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
