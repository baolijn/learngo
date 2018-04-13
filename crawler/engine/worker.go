package engine

import "learngo/crawler/fetcher"

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching url %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//log.Printf("Fetcher: error fetching url %s %v",
		//	r.Url, err)
		//return ParserFunc{body, r.Url}, err
	}
	return r.ParserFunc(body, r.Url), nil
}