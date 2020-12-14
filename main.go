package main

import (
	"os"
	"strings"

	"github.com/Srijan-Sengupta/Backup-Master/cli"
	"github.com/Srijan-Sengupta/Backup-Master/ui"
)

func main() {
	if len(os.Args) > 1 {
		if strings.Compare(os.Args[1], "cli") == 0 {
			cli.Cli()
		} else {
			ui.Gui()
		}
	} else {
		ui.Gui()
	}

}
