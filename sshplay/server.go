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

type PassAuth func(conn ssh.ConnMetadata, password []byte) error
type PubkeyAuth func(conn ssh.ConnMetadata, key ssh.PublicKey) error
type InteractiveAuth func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) error

type HandlerFactory interface {
	Create() Handler
}

type Handler interface {
	Auth() (noauth bool, pass PassAuth, pubkey PubkeyAuth, inter InteractiveAuth)
	Handle(reader *bufio.Reader, t *term.Full, c <-chan WindowSize) error
}

func wrapAuth(handler Handler) (
	noAuth bool,
	passAuth func(conn ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error),
	pubkeyAuth func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error),
	interAuth func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error),
) {
	iNoAuth, iPassAuth, iPubkeyAuth, iInterAuth := handler.Auth()
	noAuth = iNoAuth
	if iPassAuth != nil {
		passAuth = func(conn ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			return nil, iPassAuth(conn, pass)
		}
	}
	if iPubkeyAuth != nil {
		pubkeyAuth = func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			return nil, iPubkeyAuth(conn, key)
		}
	}
	if iInterAuth != nil {
		interAuth = func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
			return nil, iInterAuth(conn, client)
		}
	}
	return
}

func Run(keyfile string, addr string, handlerfactory HandlerFactory) error {
	keybytes, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return err
	}

	key, err := ssh.ParsePrivateKey(keybytes)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		handler := handlerfactory.Create()
		noAuth, passAuth, pubkeyAuth, interAuth := wrapAuth(handler)
		config := &ssh.ServerConfig{
			NoClientAuth:                noAuth,
			PasswordCallback:            passAuth,
			PublicKeyCallback:           pubkeyAuth,
			KeyboardInteractiveCallback: interAuth,
		}
		config = config
		config.AddHostKey(key)

		go handleConn(conn, handler, config)
	}
}

func handleConn(conn net.Conn, handler Handler, config *ssh.ServerConfig) {
	sconn, channels, requests, err := ssh.NewServerConn(conn, config)
	if err != nil {
		conn.Close()
		// TODO: log
		fmt.Printf("Error accepting serverconn: %v\n", err)
		return
	}
	defer sconn.Close()

	go func() {
		for request := range requests {
			request.Reply(false, nil)
		}
	}()
	channel, requests, err := getSession(channels)
	if err != nil {
		fmt.Printf("Error getting session channel: %v\n", err)
		return
	}
	go ignoreChannels(channels)
	err = prepHandle(channel, requests, handler)
	if err != nil {
		fmt.Printf("Error in prephandle: %v\n", err)
		return
	}
}

func getSession(channels <-chan ssh.NewChannel) (c ssh.Channel, r <-chan *ssh.Request, err error) {
	for channel := range channels {
		if channel.ChannelType() == "session" {
			c, r, err = channel.Accept()
			return
		}
		channel.Reject(ssh.UnknownChannelType, "unknown channel type")
	}
	return
}

func ignoreChannels(channels <-chan ssh.NewChannel) {
	for channel := range channels {
		channel.Reject(ssh.UnknownChannelType, "unknown channel type")
	}
}

type WindowSize struct {
	Width  uint32
	Height uint32
}

func prepHandle(channel ssh.Channel, requests <-chan *ssh.Request, handler Handler) error {
	if channel == nil {
		return fmt.Errorf("No session channel received.\n")
	}
	defer channel.Close()

	var termtype string
	var width, height uint32
	var shell = false
	var pty = false
	after := time.After(time.Second * 3)
	for !(shell && pty) {
		select {
		case request, ok := <-requests:
			if !ok {
				return nil
			}
			fmt.Printf("request: %s\n", request.Type)
			switch request.Type {
			case "shell":
				// TODO: Maybe only allow zero-paylength shells (no command)
				// if len(request.Payload) == 0
				shell = true
				request.Reply(true, nil)
			case "pty-req":
				_, termtype, width, height = requestPtyReq(request)
				pty = true
				request.Reply(true, nil)
			default:
				request.Reply(false, nil)
			}
		case <-after:
			return fmt.Errorf("No shell and pty request after 3 seconds")
		}
	}

	t := term.New(term.Resolve(termtype), channel)
	if t == nil {
		return fmt.Errorf("Unknown terminal type '%s'", termtype)
	}
	t.SetDimensions(width, height)

	windowsize := make(chan WindowSize)

	go func() {
		for request := range requests {
			switch request.Type {
			case "window-change":
				_, width, height := requestWindowChange(request)
				windowsize <- WindowSize{Width: width, Height: height}
				// No reply to window change
				//request.Reply(true, nil)
			default:
				request.Reply(false, nil)
			}
		}
	}()

	reader := bufio.NewReader(channel)
	return handler.Handle(reader, t, windowsize)
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
