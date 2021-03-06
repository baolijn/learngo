package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<- rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bufferReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufferReader)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}