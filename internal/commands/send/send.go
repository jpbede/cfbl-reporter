package send

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.bnck.me/cfbl-reporter/internal/commands"
	"go.bnck.me/cfbl-reporter/pkg/cfbl"
	"io"
	"os"
)

func init() {
	commands.RegisterCommand(&cli.Command{
		Name:    "send",
		Aliases: []string{"s"},
		Usage:   "Sends a CFBL report for given mail",
		Action:  runSend,
	})
}

func runSend(c *cli.Context) error {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("no pipe :(")
	} else {
		reader := bufio.NewReader(os.Stdin)
		all, _ := io.ReadAll(reader)

		report := cfbl.NewReport(&all)

		arfReport, _ := report.ComposeARFReport(
			cfbl.WithFrom("FBL Sender", "fbl@exmaple.com"),
		)
		fmt.Print(string(arfReport))
	}

	return nil
}
