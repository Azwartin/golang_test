package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//ftp commands
const (
	CD  = "CD"
	DIR = "DIR"
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:2121"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	exitIfError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	exitIfError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		handleCommand(string(buf[0:n]), conn)
	}
}

func handleCommand(command string, conn net.Conn) {
	if len(command) < 3 {
		conn.Write([]byte("Unknown command \n"))
		return
	}

	cmd := strings.ToUpper(command[0:3])
	if cmd[0:2] == CD {
		chdir(strings.TrimSuffix(command[3:], "\r\n"), conn)
	} else if cmd == DIR {
		dirlist(conn)
	} else if cmd == PWD {
		pwd(conn)
	} else {
		conn.Write([]byte("Unknown command \n"))
	}
}

//print name current directory
func pwd(conn net.Conn) {
	defer conn.Write([]byte("\n"))
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	conn.Write([]byte(s + "\n"))
}

//change directory to dir
func chdir(dir string, conn net.Conn) {
	defer conn.Write([]byte("\n"))
	err := os.Chdir(dir)
	if err == nil {
		conn.Write([]byte("Changed to " + dir))
	} else {
		conn.Write([]byte(err.Error()))
	}
}

//list information about the files in current directory
func dirlist(conn net.Conn) {
	defer conn.Write([]byte("\n"))
	dir, err := os.Open(".")
	if err != nil {
		conn.Write([]byte(err.Error()))
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		conn.Write([]byte(err.Error()))
	}

	for _, name := range names {
		conn.Write([]byte(name + "\n"))
	}
}

func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
