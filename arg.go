package curl2struct

import (
	"strings"
)

type Args []Arg

func NewArgs(args []string) Args {
	as := make(Args, len(args))
	for i, a := range args {
		as[i] = Arg(a)
	}
	return as
}

func (as Args) format() {
	args := Args{}
	for _, a := range as {
		if a == "\n" {
			continue
		}
		a.format()
		args = append(args, a)
	}
	copy(as, args)
}

type Arg string

func (a *Arg) format() {
	*a = Arg(strings.TrimSpace(a.String()))
	*a = Arg(strings.ReplaceAll(a.String(), "\n", ""))
	if !a.IsURL() {
		*a = Arg(strings.TrimLeft(a.String(), "$"))
	}
}

func (a *Arg) IsURL() bool {
	return strings.HasPrefix(a.String(), "http://") || strings.HasPrefix(a.String(), "https://")
}

func (a *Arg) ParseHeader() (key, value string) {
	key, value, _ = strings.Cut(a.String(), ":")
	return
}

func (a *Arg) String() string {
	if a != nil {
		return string(*a)
	}
	return ""
}
