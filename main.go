package main

import (
	"fmt"
	"github.com/rajivnathan/kdo-poc/common"
	odocommon "github.com/rajivnathan/odo-poc/common"
	"github.com/spf13/cobra"
	"strings"
)

func main() {

	var echoTimes int

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "odo-poc"}
	rootCmd.AddCommand(common.CmdPrintKdo, cmdEcho)
	rootCmd.AddCommand(odocommon.CmdPrintOdo)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}
