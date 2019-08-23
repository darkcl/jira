package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/drgrib/alfred"

	jira "github.com/andygrunwald/go-jira"
	helpers "github.com/darkcl/jira/helpers"
	models "github.com/darkcl/jira/models"
)

var shouldOpen bool
var shouldOutputAlfred bool

func consoleLog(format string, a ...interface{}) {
	if shouldOutputAlfred {
		return
	}

	// Hack around
	fmt.Printf(format, a...)
}

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
			consoleLog("Jira ticket not provided.\n")
			return
		}

		browserURL := fmt.Sprintf("%s/browse/%s", viper.GetString("host"), ticketNumber)

		consoleLog("Issue URL: %s\n", browserURL)
		if shouldOpen {
			helpers.OpenBrowser(browserURL)
		}

		if shouldOutputAlfred {
			alfred.Add(alfred.Item{
				Title:    ticketNumber,
				Subtitle: "Issue Key",
				Arg:      browserURL,
				UID:      "issue.Key",
			})
		}

		client, _ := jira.NewClient(tp.Client(), viper.GetString("host"))
		issue, _, err := client.Issue.Get(ticketNumber, nil)

		if err != nil {
			consoleLog("\nerror: %v\n", err)
			return
		}
		consoleLog("[%s](%s) %s\n", issue.Key, issue.ID, issue.Fields.Summary)
		consoleLog("Status: %s \n", issue.Fields.Status.Name)
		consoleLog("Assignee: %s \n", issue.Fields.Assignee.DisplayName)

		prEndpoint := fmt.Sprintf("rest/dev-status/1.0/issue/detail?issueId=%s&applicationType=bitbucket&dataType=pullrequest", issue.ID)

		req, _ := client.NewRequest("GET", prEndpoint, nil)

		devStatus := new(models.DevelopStatus)
		_, err = client.Do(req, devStatus)
		if err != nil {
			panic(err)
		}

		consoleLog("\n")

		for _, detail := range devStatus.Detail {
			for _, branch := range detail.Branches {
				consoleLog("git checkout %s\t [%s]\n", branch.Name, branch.Repository.Name)
			}

			consoleLog("\n")

			for _, pr := range detail.PullRequests {
				consoleLog("Pull Request: %s [%s]\n", pr.Name, pr.Status)
				consoleLog("Pull Request URL: %s \n\n", pr.URL)

				if shouldOpen && pr.Status == "OPEN" {
					helpers.OpenBrowser(pr.URL)
				}

				if shouldOutputAlfred && pr.Status == "OPEN" {
					alfred.Add(alfred.Item{
						Title:    issue.Key,
						Subtitle: "Pull Request",
						Arg:      pr.URL,
						UID:      "pull.request",
					})
				}
			}
		}

		consoleLog("================================\n\n")
		epic := helpers.GetEpicFromIssue(issue.ID, client)
		if epic != nil {
			epicURL := fmt.Sprintf("%s/browse/%s", viper.GetString("host"), epic.Key)
			consoleLog("Epic URL: %s\n", epicURL)
			if shouldOpen {
				helpers.OpenBrowser(epicURL)
			}
			if shouldOutputAlfred {
				alfred.Add(alfred.Item{
					Title:    issue.Key,
					Subtitle: "Epic Key",
					Arg:      epicURL,
					UID:      "epic.Key",
				})
			}
		}
		if shouldOutputAlfred {
			alfred.Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
	rootCmd.PersistentFlags().BoolVarP(&shouldOpen, "open", "o", false, "Open in browser")
	rootCmd.PersistentFlags().BoolVarP(&shouldOutputAlfred, "alfred", "a", false, "Output in alfred")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
