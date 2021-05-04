package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sadmansakib/bdix_servers/internal"
	"github.com/spf13/cobra"
)

var (
	threads uint8
	timeout uint8
	cmd     = &cobra.Command{
		Use:     "bdix-tester",
		Short:   "CLI for finding compatible servers",
		Long:    "A small CLI for compatible BDIX servers from a predefined list of servers",
		Version: "v0.1.0",
		Run: func(_ *cobra.Command, _ []string) {
			if IsOnline() {
				internal.GetAccessibleServers(threads, timeout)
			} else {
				fmt.Println("unable to connect to the internet")
				os.Exit(1)
			}
		},
	}
)

func IsOnline() bool {
	_, err := http.Get("https://icanhazip.com/")
	//err = nil means online
	if err == nil {
		return true
	}
	//if the "return statement" in the if didn't executed,
	//this one will execute surely
	return false
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cmd.Flags().Uint8VarP(&threads,
		"workers",
		"w",
		2,
		"Define number of threads you want to assign to the program")
	cmd.Flags().Uint8VarP(&timeout,
		"timeout",
		"t",
		30,
		"Define network timeout (in seconds) if network response is unresponsive")
}
