package curl2struct_test

import (
	. "github.com/3n0ugh/curl2struct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Arg2State(t *testing.T) {
	assert.Equal(t, Arg2State["-A"], USER_AGENT)
	assert.Equal(t, Arg2State["--user-agent"], USER_AGENT)
	assert.Equal(t, Arg2State["-H"], HEADER)
	assert.Equal(t, Arg2State["--header"], HEADER)
	assert.Equal(t, Arg2State["-d"], BODY)
	assert.Equal(t, Arg2State["--data"], BODY)
	assert.Equal(t, Arg2State["--data-ascii"], BODY)
	assert.Equal(t, Arg2State["--data-raw"], BODY)
	assert.Equal(t, Arg2State["-u"], USER)
	assert.Equal(t, Arg2State["--user"], USER)
	assert.Equal(t, Arg2State["-X"], METHOD)
	assert.Equal(t, Arg2State["--request"], METHOD)
	assert.Equal(t, Arg2State["-b"], COOKIE)
	assert.Equal(t, Arg2State["--cookie"], COOKIE)
	assert.Equal(t, Arg2State[""], EMPTY)
	assert.Equal(t, Arg2State["\n"], EMPTY)
}
