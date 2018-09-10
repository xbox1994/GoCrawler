package fetcher

import (
	"bufio"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) (body []byte, err error) {
	<-rateLimiter

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, _ := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	newBufferReader := bufio.NewReader(resp.Body)
	reader := transform.NewReader(newBufferReader, getEncoding(newBufferReader).NewDecoder())
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(reader)
}

func getEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, e := reader.Peek(1024)
	if e != nil {
		log.Error("Fetcher error: %v", e)
		return unicode.UTF8
	}
	charsetString, _, _ := charset.DetermineEncoding(bytes, "")
	return charsetString
}
