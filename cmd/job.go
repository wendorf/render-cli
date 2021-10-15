package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Get job by service id and job id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("provide exactly two arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		jobId := args[1]
		jsonString, err := http.Request(fmt.Sprintf("services/%s/jobs/%s", serviceId, jobId))
		if err != nil {
			panic(err)
		}

		var job Job
		if err := json.Unmarshal(jsonString, &job); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "Status", "CreatedAt", "FinishedAt"}, []Job{job}); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
}
