package cmd

import (
	"binhong/kwcli/common"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ... (Your other imports)
var definedFlags = make(map[string]bool)

var outputFile string

var GlobalKWClient *common.KWClient // Declare the global variable

// ... your other commands ...

// The getKWClientInstance function
func getKWClientInstance() *common.KWClient {
	return GlobalKWClient
}

var rootCmd = &cobra.Command{
	Use:   "kwcli",
	Short: "Interact with Klocwork servers using the Web API",
	Long: `kwcli provides a command-line interface for leveraging
          the Klocwork Web API. Manage projects, reports, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action for the root command (could display help)
		fmt.Println("Welcome to kwcli! Use 'kwcli help' for usage.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "result.json", "The output file for results")
}
