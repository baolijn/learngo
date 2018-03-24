package real

import (
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic("err")
	}
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic("err")
	}
	return string(result)
}

