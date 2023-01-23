package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "localhost:12345"

const proto = "tcp4"

const (
	share          = "Don't communicate by sharing memory, share memory by communicating"
	conc           = "Concurrency is not parallelism."
	channels       = "Channels orchestrate; mutexes serialize."
	zero           = "Make the zero value useful"
	emptyInterface = "interface{} says nothing"
)

var phrases []string = []string{share, conc, channels, zero, emptyInterface}

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleFunc(conn)
	}
}

func handleFunc(c net.Conn) {
	for {
		c.Write([]byte(phrases[rand.Intn(3)] + "\n"))
		// fmt.Println(phrases[rand.Intn(3)])
		time.Sleep(3 * time.Second)
	}
}

//         conn.Write([]byte(time.Now().String() + "\n"))
