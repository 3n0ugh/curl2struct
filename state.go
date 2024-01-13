package curl2struct

type State string

const (
	EMPTY      State = ""
	METHOD     State = "METHOD"
	HEADER     State = "HEADER"
	USER_AGENT State = "USER_AGENT"
	BODY       State = "BODY"
	USER       State = "USER"
	COOKIE     State = "COOKIE"
)

var Arg2State = map[Arg]State{
	"-A":           USER_AGENT,
	"--user-agent": USER_AGENT,
	"-H":           HEADER,
	"--header":     HEADER,
	"-d":           BODY,
	"--data":       BODY,
	"--data-ascii": BODY,
	"--data-raw":   BODY,
	"-u":           USER,
	"--user":       USER,
	"-X":           METHOD,
	"--request":    METHOD,
	"-b":           COOKIE,
	"--cookie":     COOKIE,
	"":             EMPTY,
	"\n":           EMPTY,
}
