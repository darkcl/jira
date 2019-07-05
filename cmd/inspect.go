package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jira "github.com/andygrunwald/go-jira"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := viper.GetString("jira_user")
		accessKey := viper.GetString("jira_access_key")
		tp := jira.BasicAuthTransport{
			Username: user,
			Password: accessKey,
		}

		client, _ := jira.NewClient(tp.Client(), viper.GetString("host"))
		issue, _, err := client.Issue.Get(args[0], nil)

		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
		fmt.Printf("[%s] %s\n", issue.Key, issue.Fields.Summary)
		fmt.Printf("Open: %s/browse/%s", viper.GetString("host"), issue.Key)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
