package htmlcmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/PieterD/crap/agoge/html"
	"github.com/PieterD/crap/agoge/internal/flags"
)

func init() {
	flags.SetCommand("html", new(HtmlCmd))
}

type HtmlCmd struct {
	flags.Env
	Out string
	In  []string
}

func (c *HtmlCmd) Flags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "out", "", "Output .go file")
}

func (c *HtmlCmd) Check(args []string) string {
	if c.Out == "" {
		return "Expected -out"
	}
	c.In = args
	return ""
}

func (c *HtmlCmd) Main() {
	err := html.Incorporate(c.Pkg, c.Out, c.In...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error incorporating html: %v", err)
		os.Exit(1)
	}
}
