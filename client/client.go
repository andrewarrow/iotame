package client

import "fmt"
import "time"
import "io/ioutil"
import "net"
import "net/http"
import "bytes"

func Connect(node map[string]string) {
	fmt.Println(node)
	url := fmt.Sprintf("http://%s", node["id"])
	b := "{\"command\": \"getTrytes\", \"hashes\": [\"UYUKWGQZXIOMVNAZQXPCBBYFPJSDHZFBYMTPANPCZACZMBTHXJVEYYJTGWFGMAZQJEONSNYNLIEPZ9999\"]}"
	rd := bytes.NewReader([]byte(b))
	req, err := http.NewRequest("POST", url, rd)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IOTA-API-Version", "1")

	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		},
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
	resp.Body.Close()
}
