package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebResult struct {
}

func GetResults(query string) {

	url := "https://api.search.brave.com/res/v1/web/search?q=brave%2Bsearch"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	// req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("X-Subscription-Token", "")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]interface{}{}

	err = json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}
	parseMap(m)
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			// fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			// fmt.Println(key)
			if key != "videos" {
				parseArray(val.([]interface{}))
			}
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			// fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			// fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
