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
func User(w io.Writer, user string, mode int, realname string) error {
	_, err := fmt.Fprintf(w, "USER %s %d * :%s\r\n", user, mode, realname)
	return err
}

// Quit writes a QUIT command to writer
// 3.1.7 in RFC 2812
func Quit(w io.Writer, message string) error {
	_, err := fmt.Fprint(w, "QUIT")
	if err != nil {
		return err
	}
	if len(message) > 0 {
		_, err = fmt.Fprintf(w, " :%s", message)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\r\n")
	return err
}

// Join writes a JOIN command to writer
// 3.2.1 in RFC 2812
func Join(w io.Writer, channels, keys []string) error {
	_, err := fmt.Fprintf(w, "JOIN %s", strings.Join(channels, ","))
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		_, err := fmt.Fprintf(w, " %s", strings.Join(keys, ","))
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\r\n")
	return err
}

// Part writes a PART command to writer
// 3.2.2 in RFC 2812
func Part(w io.Writer, channels []string, message string) error {
	//_, err := fmt.Fprintf(w, "PART %s\r\n", strings.Join(channels, ","))
	_, err := fmt.Fprintf(w, "PART %s", strings.Join(channels, ","))
	if err != nil {
		return err
	}
	if len(message) > 0 {
		_, err = fmt.Fprintf(w, " :%s", message)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\r\n")
	return err
}

// Names writes a NAMES command to writer
// 3.2.5 in RFC 2812
func Names(w io.Writer, channels []string) error {
	_, err := fmt.Fprintf(w, "NAMES :%s\r\n", strings.Join(channels, ","))
	return err
}

// Ping writes a PING command to writer
// 3.7.2 in RFC 2813
func Ping(w io.Writer, from, to, by string) error {
	_, err := fmt.Fprint(w, "PING")
	if err != nil {
		return err
	}
	if len(from) > 0 {
		_, err = fmt.Fprintf(w, " %s", from)
		if err != nil {
			return err
		}
	}
	if len(to) > 0 {
		_, err = fmt.Fprintf(w, " %s", to)
		if err != nil {
			return err
		}
	}
	if len(by) > 0 {
		_, err = fmt.Fprintf(w, " :%s", by)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\r\n")
	return err
}

// Pong writes a PONG command to writer
// 3.7.3 in RFC 2813
func Pong(w io.Writer, from, to, by string) error {
	_, err := fmt.Fprint(w, "PONG")
	if err != nil {
		return err
	}
	if len(from) > 0 {
		_, err = fmt.Fprintf(w, " %s", from)
		if err != nil {
			return err
		}
	}
	if len(to) > 0 {
		_, err = fmt.Fprintf(w, " %s", to)
		if err != nil {
			return err
		}
	}
	if len(by) > 0 {
		_, err = fmt.Fprintf(w, " :%s", by)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\r\n")
	return err
}
