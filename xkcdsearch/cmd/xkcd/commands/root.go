package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	// IndexName is the Elasticsearch index name.
	IndexName string

	tWidth int
)

var rootCmd = &cobra.Command{
	Use:   "xkcd",
	Short: "xkcd allows you to index and search xkcd.com",
	// Long:  "TODO",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&IndexName, "index", "i", "xkcd", "Index name")
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
}

// Execute launches the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
