package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/PieterD/bites"
	"golang.org/x/crypto/ssh"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// TODO: configurable key file
	keybytes, err := ioutil.ReadFile("./id_rsa")
	Panic(err)

	key, err := ssh.ParsePrivateKey(keybytes)
	Panic(err)

	config := &ssh.ServerConfig{
		NoClientAuth: true,
	}
	config = config
	config.AddHostKey(key)

	// TODO: configurable address
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	Panic(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		Panic(err)
		_, channels, requests, err := ssh.NewServerConn(conn, config)
		Panic(err)
		go func() {
			for request := range requests {
				fmt.Printf("request: %#v\n", request)
				request.Reply(false, nil)
			}
		}()
		for channel := range channels {
			fmt.Printf("%s\n", channel.ChannelType())
			if channel.ChannelType() != "session" {
				channel.Reject(ssh.UnknownChannelType, "unknown channel type")
				continue
			}
			schannel, srequests, err := channel.Accept()
			Panic(err)
			go handle(schannel, srequests)
		}
	}
}

func handle(channel ssh.Channel, requests <-chan *ssh.Request) {
	go func() {
		for request := range requests {
			// fmt.Printf("request: %#v\n", request)
			ok := false
			switch request.Type {
			case "pty-req":
				ok = true
			case "shell":
				ok = true
			case "window-change":
				b := bites.Bites(request.Payload)
				if len(b) >= 8 {
					var width, height uint32
					b.GetUint32(&width).GetUint32(&height)
					fmt.Printf("window %dx%d\n", width, height)
				}
			}
			request.Reply(ok, nil)
		}
	}()
	reader := bufio.NewReader(channel)
	for {
		r, _, err := reader.ReadRune()
		Panic(err)
		fmt.Printf("'%c'\n", r)
		fmt.Fprintf(channel, "'%c'\r\n", r)
	}
}
