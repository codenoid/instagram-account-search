package igaccountsearch

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

// GetHTML ...
func GetHTML(urlstr, ua string) (string, error) {
	//log.Print("getHTML")

	//NOTE: http://stackoverflow.com/questions/13130341/reading-gzipped-http-response-in-go

	if ua == "" {
		ua = "Googlebot"
	}
	req, _ := http.NewRequest("GET", urlstr, nil)
	// Set User-Agent to Googlebot
	req.Header.Set("User-Agent", ua)

	//gzip
	// req.Header.Add("Accept-Encoding", "gzip")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")

	// New Client
	tr := &http.Transport{
		//MaxIdleConns:       10,
		//IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	cl := &http.Client{
		Transport: tr,
	}
	// Send request
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Encoding") == "gzip" {

		zipread, e := gzip.NewReader(resp.Body)
		if e != nil {
			log.Print(e)
			return "", e
		}
		defer zipread.Close()

		reader := bufio.NewReader(zipread)
		var part []byte
		ret := ""

		for {
			if part, _, err = reader.ReadLine(); err != nil {
				break
			}
			ret += string(part)
		}
		return ret, nil
	}

	//log.Print("b, err := ioutil.ReadAll(resp.Body)")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
