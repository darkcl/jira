package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jira "github.com/andygrunwald/go-jira"
	helpers "github.com/darkcl/jira/helpers"
	models "github.com/darkcl/jira/models"
)

var shouldOpen bool

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

		client, _ := jira.NewClient(tp.Client(), viper.GetString("host"))
		issue, _, err := client.Issue.Get(ticketNumber, nil)

		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
		fmt.Printf("[%s](%s) %s\n", issue.Key, issue.ID, issue.Fields.Summary)
		fmt.Printf("Status: %s \n", issue.Fields.Status.Name)
		fmt.Printf("Assignee: %s \n", issue.Fields.Assignee.DisplayName)

		prEndpoint := fmt.Sprintf("rest/dev-status/1.0/issue/detail?issueId=%s&applicationType=bitbucket&dataType=pullrequest", issue.ID)

		req, _ := client.NewRequest("GET", prEndpoint, nil)

		devStatus := new(models.DevelopStatus)
		_, err = client.Do(req, devStatus)
		if err != nil {
			panic(err)
		}

		fmt.Println()

		for _, detail := range devStatus.Detail {
			for _, branch := range detail.Branches {
				fmt.Printf("git checkout %s\t [%s]\n", branch.Name, branch.Repository.Name)
			}

			fmt.Println()

			for _, pr := range detail.PullRequests {
				fmt.Printf("Pull Request: %s [%s]\n", pr.Name, pr.Status)
				fmt.Printf("Pull Request URL: %s \n\n", pr.URL)

				if shouldOpen && pr.Status == "OPEN" {
					helpers.OpenBrowser(pr.URL)
				}
			}
		}

		browserURL := fmt.Sprintf("%s/browse/%s", viper.GetString("host"), issue.Key)

		fmt.Printf("Issue URL: %s\n", browserURL)
		if shouldOpen {
			helpers.OpenBrowser(browserURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
	rootCmd.PersistentFlags().BoolVarP(&shouldOpen, "open", "o", false, "Open in browser")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
