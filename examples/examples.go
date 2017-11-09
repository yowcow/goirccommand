package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"

	command "github.com/yowcow/goirccommand"
)

var server = "localhost:6667"

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	conn, err := net.Dial("tcp", server)
	if err != nil {
		logger.Println(err)
		return
	}
	defer conn.Close()

	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)
	done := make(chan bool)

	go writeCommand(w, logger, done)
	go readCommand(r, logger)

	for {
		select {
		case <-done:
			logger.Println("got a done signal. waiting 3 secs to finish up.")
			time.Sleep(3 * time.Second)
			return
		}
	}
}

type commandFunc func(w io.Writer)

func writeCommand(w *bufio.Writer, logger *log.Logger, done chan<- bool) {
	defer func() {
		done <- true
	}()

	cmds := []commandFunc{
		func(w io.Writer) {
			command.Nick(w, "mynick")
		},
		func(w io.Writer) {
			command.User(w, "myuser", 0, "my name")
		},
		func(w io.Writer) {
			command.Ping(w, "", "", "mynick")
		},
		func(w io.Writer) {
			command.Pong(w, "", "", "mynick")
		},
		func(w io.Writer) {
			command.Join(w, []string{"#test", "#test1", "#test2"}, []string{})
		},
		func(w io.Writer) {
			command.Names(w, []string{"#test", "#test1", "#test2"})
		},
		func(w io.Writer) {
			command.Part(w, []string{"#test1", "#test2"}, "")
		},
		func(w io.Writer) {
			command.Part(w, []string{"#test"}, "Bye!")
		},
		func(w io.Writer) {
			command.Quit(w, "Closed?")
		},
	}

	for _, cmd := range cmds {
		cmd(w)
		w.Flush()
		time.Sleep(1 * time.Second)
	}
}

func readCommand(r *bufio.Reader, logger *log.Logger) {
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			logger.Println(err)
			return
		}
		logger.Println(string(line))
	}
}
