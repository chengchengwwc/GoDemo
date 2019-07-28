package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}  //任何类型
}


//空函数
func NilParser([]byte)ParseResult{
	return ParseResult{}
}