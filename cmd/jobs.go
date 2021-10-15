package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Get jobs by service id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		jsonString, err := http.Request(fmt.Sprintf("services/%s/jobs", serviceId))
		if err != nil {
			panic(err)
		}

		var cursorJobs []CursorJob
		if err := json.Unmarshal(jsonString, &cursorJobs); err != nil {
			panic(err)
		}

		var jobs []Job
		for _, cursorJob := range cursorJobs {
			jobs = append(jobs, cursorJob.Job)
		}

		if err := table.Print([]string{"Id", "ServiceId", "StartCommand", "CreatedAt", "FinishedAt", "Status"}, jobs); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(jobsCmd)
}
