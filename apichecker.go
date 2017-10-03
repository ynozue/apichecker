package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"net/url"
	"strings"
	"io/ioutil"
	"time"
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
		result = fmt.Sprintf("NG\n%s", err)
	} else {
		defer resp.Body.Close()
		expire := "-"
		if len(resp.TLS.PeerCertificates) > 0 {
			expireUTCTime := resp.TLS.PeerCertificates[0].NotAfter
			expireJSTTime := expireUTCTime.In(time.FixedZone("Asia/Tokyo", 9 * 60 * 60))
			expire = expireJSTTime.Format("06/01/02 15:04")
		}
		result = fmt.Sprintf("OK[expire=%s]\n%s", expire, endpoint)
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
