package main

import (
	goflag "flag"
	"log"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var (
	rootCmd = &cobra.Command{
		Use:   "invtools",
		Short: "Invoice Tools",
		Long:  "CLI containing a number of invoice-related tools",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			goflag.Parse()
		},
	}
)

func init() {

	// Add Go flags to cobra
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

// Main function to be exeuted
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Root command execution failed, error: %v", err)
	}
}
