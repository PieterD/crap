package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/PieterD/bites"
	"github.com/PieterD/crap/sshplay/term"
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

	// TODO: Authenticate client keys
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
		sconn, channels, requests, err := ssh.NewServerConn(conn, config)
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
			go prepHandle(sconn, schannel, srequests)
		}
	}
}

type ptyReq struct {
	width  uint32
	height uint32
	term   string
	pty    bool
}

func prepHandle(conn ssh.Conn, channel ssh.Channel, requests <-chan *ssh.Request) {
	defer conn.Close()
	defer channel.Close()
	err := handle(channel, requests)
	fmt.Printf("Error handling connection: %v\n", err)
}

func handle(channel ssh.Channel, requests <-chan *ssh.Request) error {
	ptychan := make(chan ptyReq)
	go func() {
		defer close(ptychan)
		for request := range requests {
			respond := true
			fmt.Printf("request: %#v\n", request)
			ok := false
			switch request.Type {
			case "shell":
				if len(request.Payload) == 0 {
					ok = true
				}
			case "pty-req":
				ok = true
				_, term, width, height := requestPtyReq(request)
				ptychan <- ptyReq{width: width, height: height, term: term, pty: true}
			case "window-change":
				respond = false
				_, width, height := requestWindowChange(request)
				ptychan <- ptyReq{width: width, height: height}
			}
			if respond {
				request.Reply(ok, nil)
			}
		}
	}()

	var t *term.Full = nil
	after := time.After(time.Second * 3)
	for t == nil {
		select {
		case pr, ok := <-ptychan:
			if !ok {
				return nil
			}
			if pr.pty {
				t = term.New(term.Resolve(pr.term), channel)
				if t == nil {
					return fmt.Errorf("Unknown terminal type '%s'", pr.term)
				}
				t.SetDimensions(pr.width, pr.height)
				fmt.Printf("set!\n")
			}
		case <-after:
			return fmt.Errorf("No pty request after 3 seconds")
		}
	}

	go func() {
		for pr := range ptychan {
			t.SetDimensions(pr.width, pr.height)
		}
	}()

	reader := bufio.NewReader(channel)
	for {
		r, _, err := reader.ReadRune()
		Panic(err)
		fmt.Printf("'%c'\n", r)

		t.Clear()
		t.Pos(10, 5)
		t.Printf("'%c' ", r)
		t.Attr().Fore(term.Red).Done()
		t.Printf("red ")
		t.Attr().Fore(term.Default).Done()
		t.Printf("default ")
		t.Attr().Fore(term.Red).Bright().Done()
		t.Printf("brightred ")
		t.Attr().Reset().Done()
		t.Printf("reset ")
		t.Printf("\r\n")
		if t.Error() != nil {
			fmt.Printf("%v\n", t.Error())
		}
	}
}

func requestWindowChange(request *ssh.Request) (ok bool, width, height uint32) {
	ok = !bites.Get(request.Payload).GetUint32(&width).GetUint32(&height).Error()
	return
}

func requestPtyReq(request *ssh.Request) (ok bool, term string, width, height uint32) {
	var length uint32
	var slice []byte
	ok = !bites.Get(request.Payload).GetUint32(&length).GetSlice(&slice, int(length)).GetUint32(&width).GetUint32(&height).Error()
	term = string(slice)
	return
}
