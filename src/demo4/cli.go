package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/gitioapi"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		c string
		url string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&c, "c", "", "'code' parameter.")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Parse argument
	url = ""
	flags.Parse(args[1:])
	for flags.NArg() > 0 {
		if len(url) == 0 {
			url = flags.Args()[0]
		} else {
			fmt.Fprintln(cli.errStream, os.ErrInvalid, flags.Args()[0])
			return ExitCodeError
		}
		flags.Parse(flags.Args()[1:])
	}

	if len(url) == 0 {
		// Get URL from STDIN
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			url = strings.Trim(scanner.Text(), " \n\r")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(cli.errStream, os.ErrInvalid, err)
			return ExitCodeError
		}
	}

	// shortening URL
	shortUrl, err := gitioapi.Encode(&gitioapi.Param{Url: url, Code: c})
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	fmt.Fprint(cli.outStream, shortUrl)

	return ExitCodeOK
}
