package myhttp

import (
	"strings"
)

type Request struct {
	Line   RequestLine
	Header map[string]string
	Body   string
}

type RequestLine struct {
	Method  string
	Url     string
	Version string
}

func NewRequest(str string) *Request {
	var mp map[string]string = make(map[string]string)
	arr := strings.Split(str, "\r\n")
	requestArr := arr[0]
	requestLine := strings.Split(requestArr, " ")

	for _, ele := range arr[1:] {
		if !strings.Contains(ele, ":") {
			break
		} else {
			keyVal := strings.Split(ele, ":")
			key := keyVal[0]
			val := strings.TrimSpace(keyVal[1])
			mp[key] = val

		}
	}

	return &Request{
		Line: RequestLine{
			Method:  requestLine[0],
			Url:     requestLine[1],
			Version: requestLine[2],
		},
		Header: mp,
		Body:   arr[len(arr)-1],
	}
}
