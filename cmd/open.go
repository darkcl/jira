package cmd

import (
	"fmt"
	"os"

	helpers "github.com/darkcl/jira/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		var ticketNumber string
		if len(args) > 0 {
			ticketNumber = args[0]
		} else {
			ticketNumber = helpers.GetTicketNumberFromGit(dir)
		}

		if len(ticketNumber) == 0 {
			fmt.Printf("Jira ticket not provided.\n")
			return
		}

		ticketURL := fmt.Sprintf("%s/browse/%s", viper.GetString("host"), ticketNumber)
		fmt.Println(ticketURL)
		helpers.OpenBrowser(ticketURL)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
