package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string, v interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	return json.Unmarshal([]byte(sb), v)
}
