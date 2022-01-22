package main

import (
	"flag"

	"git.sr.ht/~anecdotal/ranj/server"
)

func main() {
	hostPtr := flag.String("host", "localhost", "server address")
	portPtr := flag.String("port", "8080", "server listening port")

	host := *hostPtr
	port := *portPtr

	flag.Parse()

	server.StartServer(host, port)
}
