package goirccommand

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	input    []string
	expected string
}

type ChannelCase struct {
	input    [][]string
	expected string
}

func TestPass(t *testing.T) {
	buf := new(bytes.Buffer)
	err := Pass(buf, "hogefuga")

	assert.Nil(t, err)
	assert.Equal(t, "PASS hogefuga\r\n", buf.String())
}

func TestNick(t *testing.T) {
	buf := new(bytes.Buffer)
	err := Nick(buf, "hogefuga")

	assert.Nil(t, err)
	assert.Equal(t, "NICK hogefuga\r\n", buf.String())
}

func TestUser(t *testing.T) {
	buf := new(bytes.Buffer)
	err := User(buf, "hogefuga", 0, "Hoge Fuga")

	assert.Nil(t, err)
	assert.Equal(t, "USER hogefuga 0 * :Hoge Fuga\r\n", buf.String())
}

func TestQuit(t *testing.T) {
	cases := []Case{
		{
			input:    []string{""},
			expected: "QUIT\r\n",
		},
		{
			input:    []string{"bye bye"},
			expected: "QUIT :bye bye\r\n",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := Quit(buf, c.input[0])

		assert.Nil(t, err)
		assert.Equal(t, c.expected, buf.String())
	}
}

func TestJoin(t *testing.T) {
	cases := []ChannelCase{
		{
			input: [][]string{
				{},
				{},
			},
			expected: "JOIN \r\n",
		},
		{
			input: [][]string{
				{"#hoge", "#fuga"},
				{},
			},
			expected: "JOIN #hoge,#fuga\r\n",
		},
		{
			input: [][]string{
				{"#hoge", "#fuga"},
				{"foo", "bar"},
			},
			expected: "JOIN #hoge,#fuga foo,bar\r\n",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := Join(buf, c.input[0], c.input[1])

		assert.Nil(t, err)
		assert.Equal(t, c.expected, buf.String())
	}
}

func TestPart(t *testing.T) {
	type PartCase struct {
		inputChan []string
		inputMsg  string
		expected  string
	}

	cases := []PartCase{
		{
			inputChan: []string{"#hoge", "#fuga"},
			inputMsg:  "",
			expected:  "PART #hoge,#fuga\r\n",
		},
		{
			inputChan: []string{"#hoge", "#fuga"},
			inputMsg:  "Leaving",
			expected:  "PART #hoge,#fuga :Leaving\r\n",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := Part(buf, c.inputChan, c.inputMsg)

		assert.Nil(t, err)
		assert.Equal(t, c.expected, buf.String())
	}
}

func TestNames(t *testing.T) {
	buf := new(bytes.Buffer)
	err := Names(buf, []string{"#hoge", "#fuga", "#foo", "#bar"})

	assert.Nil(t, err)
	assert.Equal(t, "NAMES :#hoge,#fuga,#foo,#bar\r\n", buf.String())
}

func TestPing(t *testing.T) {
	cases := []Case{
		{
			input:    []string{"", "", ""},
			expected: "PING\r\n",
		},
		{
			input:    []string{"", "", "hoge"},
			expected: "PING :hoge\r\n",
		},
		{
			input:    []string{"foo", "bar", ""},
			expected: "PING foo bar\r\n",
		},
		{
			input:    []string{"foo", "bar", "hoge"},
			expected: "PING foo bar :hoge\r\n",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := Ping(buf, c.input[0], c.input[1], c.input[2])

		assert.Nil(t, err)
		assert.Equal(t, c.expected, buf.String())
	}
}

func TestPong(t *testing.T) {
	cases := []Case{
		{
			input:    []string{"", "", ""},
			expected: "PONG\r\n",
		},
		{
			input:    []string{"foo", "", ""},
			expected: "PONG foo\r\n",
		},
		{
			input:    []string{"foo", "bar", ""},
			expected: "PONG foo bar\r\n",
		},
		{
			input:    []string{"foo", "bar", "hoge"},
			expected: "PONG foo bar :hoge\r\n",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := Pong(buf, c.input[0], c.input[1], c.input[2])

		assert.Nil(t, err)
		assert.Equal(t, c.expected, buf.String())
	}
}
