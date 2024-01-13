package curl2struct

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var curl = `curl 'https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1' \
	-X 'OPTIONS' \
	-H 'authority: test.com' \
	-H 'accept: */*' \
	-H 'accept-language: en-US,en;q=0.9' \
	-H 'access-control-request-headers: authorization,authorization-type,x-executor-user' \
	-H 'access-control-request-method: GET' \
	-H 'origin: https://abcd.com' \
	-H 'sec-fetch-dest: empty' \
	-H 'sec-fetch-mode: cors' \
	-H 'sec-fetch-site: same-site' \
    -u 'test:test'
	--user-agent 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' \
	-d '{
		"a": "b"
	}'
	--compressed
`

var curl2 = `curl 'https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1' \
	--request 'OPTIONS' \
    --cookie 'VisitCount=2; WebAbTesting=A_64-B_44-C_73-D_50-E_49-F_30-G_85-H_12-I_16-J_52-K_93-L_86-M_87-N_43-O_65-P_78-Q_40-R_14-S_86-T_98-U_87-V_93-W_13-X_38-Y_98-Z_50; hvtb=1'\
	-H 'authority: test.com' \
	-H 'accept: */*' \
	-H 'accept-language: en-US,en;q=0.9' \
	-H 'access-control-request-headers: authorization,authorization-type,x-executor-user' \
	-H 'access-control-request-method: GET' \
	-H 'origin: https://abcd.com' \
	-H 'sec-fetch-dest: empty' \
	-H 'sec-fetch-mode: cors' \
	-H 'sec-fetch-site: same-site' \
	-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' \
	--compressed
`

var curlWithoutMethodAndBody = `curl -i -s -k  \
	-H $'authority: test.com' \
	-H $'accept: */*' \
	-H $'accept-language: en-US,en;q=0.9' \
	-H $'access-control-request-headers: authorization,authorization-type,x-executor-user' \
	-H $'access-control-request-method: GET' \
	-H $'origin: https://abcd.com' \
	-H $'sec-fetch-dest: empty' \
	-H $'sec-fetch-mode: cors' \
	-H $'sec-fetch-site: same-site' \
	-H $'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' \
	$'https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1' 
`

var curlWithoutMethodWithBody = `curl --location 'https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1' \
	--header 'authority: test.com' \
	--header 'accept: */*' \
	--header 'accept-language: en-US,en;q=0.9' \
	--header 'access-control-request-headers: authorization,authorization-type,x-executor-user' \
	--header 'access-control-request-method: GET' \
	--header 'origin: https://abcd.com' \
	--header 'sec-fetch-dest: empty' \
	--header 'sec-fetch-mode: cors' \
	--header 'sec-fetch-site: same-site' \
	--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36' \
   	--data '{
    	"a": "b"
	}'
`

func Test_NewCurl(t *testing.T) {
	type Expected struct {
		CURL *CURL
		Err  error
	}

	testCases := map[string]struct {
		Given    string
		Expected Expected
	}{
		"should return curl must starts with curl error when curl is not starts with curl": {
			Given: "abc",
			Expected: Expected{
				CURL: nil,
				Err:  errors.New("curl string must starts with curl"),
			},
		},
		"should return shell parse error when given string has invalid shell syntax": {
			Given: "curl $($",
			Expected: Expected{
				CURL: nil,
				Err:  errors.New("invalid command line string"),
			},
		},
		"should return CURL with all fields are parsed when flag are short": {
			Given: curl,
			Expected: Expected{
				CURL: &CURL{
					URL:    "https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1",
					Method: "OPTIONS",
					Headers: http.Header{
						"Authority":                      {"test.com"},
						"Authorization":                  {"Basic dGVzdDp0ZXN0"},
						"Accept":                         {"*/*"},
						"Accept-Language":                {"en-US,en;q=0.9"},
						"Access-Control-Request-Headers": {"authorization,authorization-type,x-executor-user"},
						"Access-Control-Request-Method":  {"GET"},
						"Origin":                         {"https://abcd.com"},
						"Sec-Fetch-Dest":                 {"empty"},
						"Sec-Fetch-Mode":                 {"cors"},
						"Sec-Fetch-Site":                 {"same-site"},
						"User-Agent":                     {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
					},
					Body: []byte(`{		"a": "b"	}`),
				},
				Err: nil,
			},
		},
		"should return CURL with all fields are parsed when flag are long": {
			Given: curl2,
			Expected: Expected{
				CURL: &CURL{
					URL:    "https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1",
					Method: "OPTIONS",
					Headers: http.Header{
						"Authority":                      {"test.com"},
						"Accept":                         {"*/*"},
						"Accept-Language":                {"en-US,en;q=0.9"},
						"Access-Control-Request-Headers": {"authorization,authorization-type,x-executor-user"},
						"Access-Control-Request-Method":  {"GET"},
						"Origin":                         {"https://abcd.com"},
						"Sec-Fetch-Dest":                 {"empty"},
						"Sec-Fetch-Mode":                 {"cors"},
						"Sec-Fetch-Site":                 {"same-site"},
						"User-Agent":                     {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
						"Cookie":                         {"VisitCount=2; WebAbTesting=A_64-B_44-C_73-D_50-E_49-F_30-G_85-H_12-I_16-J_52-K_93-L_86-M_87-N_43-O_65-P_78-Q_40-R_14-S_86-T_98-U_87-V_93-W_13-X_38-Y_98-Z_50; hvtb=1"},
					},
					Body: nil,
				},
				Err: nil,
			},
		},
		"should set GET method when curl has not method flag and body": {
			Given: curlWithoutMethodAndBody,
			Expected: Expected{
				CURL: &CURL{
					URL:    "https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1",
					Method: "GET",
					Headers: http.Header{
						"Authority":                      {"test.com"},
						"Accept":                         {"*/*"},
						"Accept-Language":                {"en-US,en;q=0.9"},
						"Access-Control-Request-Headers": {"authorization,authorization-type,x-executor-user"},
						"Access-Control-Request-Method":  {"GET"},
						"Origin":                         {"https://abcd.com"},
						"Sec-Fetch-Dest":                 {"empty"},
						"Sec-Fetch-Mode":                 {"cors"},
						"Sec-Fetch-Site":                 {"same-site"},
						"User-Agent":                     {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
					},
					Body: nil,
				},
				Err: nil,
			},
		},
		"should set POST method when curl has not method flag and has body": {
			Given: curlWithoutMethodWithBody,
			Expected: Expected{
				CURL: &CURL{
					URL:    "https://test.com/attendees?size=10&sort=createdDate&order=desc&page=1",
					Method: "POST",
					Headers: http.Header{
						"Authority":                      {"test.com"},
						"Accept":                         {"*/*"},
						"Accept-Language":                {"en-US,en;q=0.9"},
						"Access-Control-Request-Headers": {"authorization,authorization-type,x-executor-user"},
						"Access-Control-Request-Method":  {"GET"},
						"Origin":                         {"https://abcd.com"},
						"Sec-Fetch-Dest":                 {"empty"},
						"Sec-Fetch-Mode":                 {"cors"},
						"Sec-Fetch-Site":                 {"same-site"},
						"User-Agent":                     {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
					},
					Body: []byte(`{    	"a": "b"	}`),
				},
				Err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actualCURL, actualErr := NewCurl(tc.Given)
			assert.Equal(t, tc.Expected.Err, actualErr)
			assert.Equal(t, tc.Expected.CURL, actualCURL)
		})
	}
}

func TestCURL_String(t *testing.T) {
	given := &CURL{
		URL:     "https://abcd.com",
		Method:  "GET",
		Headers: http.Header{"a": []string{"b"}},
		Body:    []byte(`{"a": "b"}`),
	}
	expected := "URL: https://abcd.com\nMethod: GET\nHeaders: map[a:[b]]\nBody: {\"a\": \"b\"}\n"
	assert.Equal(t, expected, given.String())
}
