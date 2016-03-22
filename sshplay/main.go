package main

import (
	"bufio"
	"fmt"

	"golang.org/x/crypto/ssh"

	"github.com/PieterD/crap/sshplay/term"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type myHandlerFactory struct{}

func (_ myHandlerFactory) Create() Handler {
	return myHandler{}
}

type myHandler struct{}

func (h myHandler) Auth() (noauth bool, pass PassAuth, pubkey PubkeyAuth, inter InteractiveAuth) {
	return false, nil, h.PublicKeyCallback, nil
}

func (_ myHandler) PublicKeyCallback(conn ssh.ConnMetadata, key ssh.PublicKey) error {
	fmt.Printf("key: %v\n", key)
	return nil
}

func (_ myHandler) Handle(reader *bufio.Reader, t *term.Full, c <-chan WindowSize) error {
	go func() {
		for ws := range c {
			t.SetDimensions(ws.Width, ws.Height)
		}
	}()

	t.Clear()
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

		if t.Error() != nil {
			fmt.Printf("%v\n", t.Error())
		}
	}

	return nil
}

func main() {
	err := Run("./id_rsa", "127.0.0.1:12345", myHandlerFactory{})
	fmt.Printf("Run failed: %v\n", err)
}
