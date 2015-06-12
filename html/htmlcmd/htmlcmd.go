package htmlcmd

import (
	"flag"

	"github.com/PieterD/agoge/internal/flags"
)

func init() {
	flags.SetCommand("html", new(HtmlCmd))
}

type HtmlCmd struct {
	flags.Env
	Out string
}

func (c *HtmlCmd) Flags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "out", "", "Output .go file")
}

func (c *HtmlCmd) Check() string {
	if c.Out == "" {
		return "Expected -out"
	}
	return ""
}

func (c *HtmlCmd) Main() {
}
