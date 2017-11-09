package goirccommand

import (
	"fmt"
	"io"
	"strings"
)

// Pass writes a PASS command to writer
// 3.1.1 in RFC 2812
func Pass(w io.Writer, pass string) error {
	_, err := fmt.Fprintf(w, "PASS %s\r\n", pass)
	return err
}

// Nick writes a NICK command to writer
// 3.1.2 in RFC 2812
func Nick(w io.Writer, nick string) error {
	_, err := fmt.Fprintf(w, "NICK %s\r\n", nick)
	return err
}

// User writes a USER command to writer
// 3.1.3 in RFC 2812
func User(w io.Writer, username string, mode int, realname string) error {
	_, err := fmt.Fprintf(w, "USER %s %d * :%s\r\n", username, mode, realname)
	return err
}

// Join writes a JOIN command to writer
// 3.2.1 in RFC 2812
func Join(w io.Writer, channels []string) error {
	_, err := fmt.Fprintf(w, "JOIN %s\r\n", strings.Join(channels, ","))
	return err
}

// Part writes a PART command to writer
// 3.2.2 in RFC 2812
func Part(w io.Writer, channels []string) error {
	_, err := fmt.Fprintf(w, "PART %s\r\n", strings.Join(channels, ","))
	return err
}

// Names writes a NAMES command to writer
// 3.2.5 in RFC 2812
func Names(w io.Writer, channels []string) error {
	_, err := fmt.Fprintf(w, "NAMES :%s\r\n", strings.Join(channels, ","))
	return err
}

// Pong writes a PONG command to writer
// 3.7.3 in RFC 2813
func Pong(w io.Writer, server string) error {
	_, err := fmt.Fprintf(w, "PONG :%s\r\n", server)
	return err
}
