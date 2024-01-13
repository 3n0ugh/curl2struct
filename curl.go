package curl2struct

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/mattn/go-shellwords"
	"net/http"
	"strings"
)

type CURL struct {
	URL     string
	Method  string
	Headers http.Header
	Body    []byte
}

func NewCurl(curl string) (*CURL, error) {
	if !strings.HasPrefix(curl, "curl ") {
		return nil, errors.New("curl string must starts with curl")
	}

	args, err := shellwords.Parse(curl)
	if err != nil {
		return nil, err
	}

	c := &CURL{Headers: http.Header{}}
	c.parse(NewArgs(args))
	return c, nil
}

func (c *CURL) parse(as Args) {
	as.format()

	var state State

	for _, a := range as {
		if s, exists := Arg2State[a]; exists {
			state = s
			continue
		}

		if a.IsURL() {
			c.URL = a.String()
			continue
		}

		switch state {
		case METHOD:
			c.Method = a.String()
			state = ""
		case HEADER:
			key, value := a.ParseHeader()
			c.Headers.Add(key, value)
			state = ""
		case USER_AGENT:
			c.Headers.Set("User-Agent", a.String())
			state = ""
		case BODY:
			c.Body = []byte(a.String())
			state = ""
		case USER:
			c.Headers.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(a.String())))
			state = ""
		case COOKIE:
			c.Headers.Add("Cookie", a.String())
			state = ""
		}
	}

	c.checkMethod()
}

func (c *CURL) String() string {
	return fmt.Sprintf("URL: %s\nMethod: %s\nHeaders: %v\nBody: %s\n", c.URL, c.Method, c.Headers, string(c.Body))
}

func (c *CURL) checkMethod() {
	if c.Method != "" {
		return
	}

	c.Method = http.MethodGet
	if len(c.Body) > 0 {
		c.Method = http.MethodPost
	}
}
