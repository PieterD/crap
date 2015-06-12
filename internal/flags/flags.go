package flags

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var flagmap = make(map[string]mainPair)

type mainPair struct {
	fs *flag.FlagSet
	c  Command
}

type Command interface {
	Flags(*flag.FlagSet)
	Check(args []string) string
	Main()
	setEnv()
}

type Env struct {
	Arch string
	OS   string
	File string
	Pkg  string
}

func (e *Env) setEnv() {
	e.Arch = os.Getenv("GOARCH")
	e.OS = os.Getenv("GOOS")
	e.File = os.Getenv("GOFILE")
	e.Pkg = os.Getenv("GOPACKAGE")
}

func SetCommand(name string, c Command) {
	mp, ok := flagmap[name]
	fs := mp.fs
	if !ok {
		fs = flag.NewFlagSet(name, flag.ExitOnError)
		fs.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of command \"%s\":\n", name)
			fs.PrintDefaults()
			fmt.Fprintf(os.Stderr, "\n")
		}
		c.Flags(fs)
		flagmap[name] = mainPair{
			fs: fs,
			c:  c,
		}
	}
}

func help() {
	var cmdlist []string
	for cmd := range flagmap {
		cmdlist = append(cmdlist, cmd)
	}
	sort.Strings(cmdlist)
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "	agoge command [arguments]\n\n")
	fmt.Fprintf(os.Stderr, "The commands are:\n")
	for _, cmd := range cmdlist {
		fmt.Fprintf(os.Stderr, "	%s\n", cmd)
	}
	fmt.Fprintf(os.Stderr, "\nUse \"agoge help [command]\" for more information about a command.\n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func helpCmd(cmd string) {
	mp, ok := flagmap[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "\nUnknown command \"%s\".\n\n", cmd)
		help()
	}
	mp.fs.Usage()
	os.Exit(1)
}

func Run() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "\nNo program name. This should never happen.\n\n")
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		help()
	}
	if os.Args[1] == "help" {
		if len(os.Args) == 2 {
			help()
		}
		if len(os.Args) == 3 {
			helpCmd(os.Args[2])
		}
		fmt.Fprintf(os.Stderr, "\nUnknown usage of \"help\" command.\n\n")
		help()
	}
	cmd := os.Args[1]
	arg := os.Args[2:]
	mp, ok := flagmap[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "\nUnknown command: \"%s\"\n\n", cmd)
		help()
	}
	if mp.fs.Parse(arg) == flag.ErrHelp {
		helpCmd(cmd)
	}
	str := mp.c.Check(mp.fs.Args())
	if str != "" {
		fmt.Fprintf(os.Stderr, "\n%s\n\n", str)
		helpCmd(cmd)
	}
	mp.c.Main()
}
