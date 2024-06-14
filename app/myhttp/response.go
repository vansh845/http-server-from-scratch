package myhttp

import (
	"fmt"
)

type Response struct {
	Version string
	Code    string
	Reason  string

	Headers ResponseHeader

	Body string
}

type ResponseHeader struct {
	ContentType   string
	ContentLength string
}

func NewResponse(body, contentType string) *Response {
	return &Response{
		Version: "HTTP/1.1",
		Code:    "200",
		Reason:  "OK",

		Headers: ResponseHeader{
			ContentType:   contentType,
			ContentLength: fmt.Sprintf("%d", len(body)),
		},

		Body: body,
	}
}

func (res *Response) ToString() string {

	return fmt.Sprintf("%s %s %s\r\nContent-Type: %s\r\nContent-Length: %s\r\n\r\n%s", res.Version, res.Code, res.Reason, res.Headers.ContentType, res.Headers.ContentLength, res.Body)

}
