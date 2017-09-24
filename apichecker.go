package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"net/url"
	"strings"
	"io/ioutil"
)

func main() {
	var endpoint = flag.String("endpoint", "", "check target Endpoint URL")
	var lineToken = flag.String("token", "", "LINE notify token")
	flag.Parse()

	var apiResult = getAPI(*endpoint)
	var result = postLINE(*lineToken, apiResult)

	fmt.Printf("LINE Post result [%t]\n", result)
}

func getAPI(endpoint string) string {
	if endpoint == "" {
		log.Fatal("not endpoint")
		return "not endpoint"
	}

	var result = ""
	resp, err := http.Get(endpoint)
	if err != nil {
		result = fmt.Sprintf("NG [%s]", err)
	} else {
		defer resp.Body.Close()
		result = fmt.Sprintf("OK [%s]", endpoint)
	}

	return result
}

func postLINE(token string, message string) bool {
	if token == "" {
		log.Fatal("not token")
		return false
	} else if message == "" {
		log.Fatal("not text")
		return false
	}

	data := url.Values{"message": {message}}
	r, _ := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(data.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
