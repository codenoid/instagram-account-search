package igaccountsearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// UserSearch  ....
func UserSearch(query string) (IGSearchResult, error) {
	//url := "https://www.instagram.com/web/search/topsearch/?context=user&query=" + string + "&rank_token=0.6885969395144884"
	query = strings.Trim(query, "　 .\n\r\t&^%$$#@!*()_+-=/,[]{}<>?")
	result := IGSearchResult{}
	q := url.QueryEscape(query)
	if q == "" {
		return result, fmt.Errorf("need search query")
	}
	urlstr := fmt.Sprintf("https://www.instagram.com/web/search/topsearch/?context=user&query=%s", url.QueryEscape(query))

	// Create a Resty Client
	client := resty.New()
	client.SetProxy(getProxy())

	resp, err := client.R().
		Get(urlstr)
	//log.Print(body)
	if err != nil {
		log.Print(err)
		return result, err
	}
	//parse JSON
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Print(err)
		log.Print(string(resp.Body()))
		return result, err
	}
	if result.Status != "ok" {
		log.Print("status is not ok")
		return result, fmt.Errorf("status is not ok")
	}
	return result, nil

}

func getProxy() string {
	b, err := ioutil.ReadFile(os.Getenv("PROXY_FILE_PATH"))
	if err != nil {
		return ""
	}
	proxies := strings.Split(string(b), "\n")
	rand.Seed(time.Now().Unix())
	return proxies[rand.Intn(len(proxies))]
}
