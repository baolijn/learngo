package engine

type ParserFunc func(contents []byte) ParseResult

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url 	string
	Type 	string
	Id 		string
	Payload interface{}
}

type Request struct {
	Url string
	ParserFunc func([] byte) ParseResult
}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}