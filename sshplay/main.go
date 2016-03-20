package main

import (
	"bufio"
	"fmt"

	"github.com/PieterD/crap/sshplay/term"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type myHandler struct{}

func (_ myHandler) Handle(reader *bufio.Reader, t *term.Full, c <-chan WindowSize) error {
	fmt.Printf("HANDLING!\n")
	go func() {
		for ws := range c {
			t.SetDimensions(ws.Width, ws.Height)
		}
	}()
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

	return nil
}

func main() {
	err := Run("./id_rsa", "127.0.0.1:12345", myHandler{})
	fmt.Printf("Run failed: %v\n", err)
}
