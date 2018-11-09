package main

import (
	"log"
	"net"
)




func main() {	
	target, err := net.Dial("unix", "/target.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer target.Close()

    spoof, err := net.Listen("unix", "/var/run/docker.sock")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		fd, err := spoof.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go serve("> ", fd, target)
		go serve("< ", target, fd)
	}
}


func serve(direction string, from net.Conn, to net.Conn) {
	for {	
		buf := make([]byte, 512)
		nr, err := from.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println(direction, string(data))			

		_, err = to.Write(data)
		if err != nil {
			log.Fatal("Writing error: ", err)
		}
	}
}

