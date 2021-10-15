package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

type Job struct {
	Id           string `json:"id"`
	ServiceId    string `json:"serviceId"`
	StartCommand string `json:"startCommand"`
	PlanId       string `json:"planId"`
	CreatedAt    string `json:"createdAt"`
	StartedAt    string `json:"startedAt"`
	FinishedAt   string `json:"finishedAt"`
	Status       string `json:"status"`
}

type CursorJob struct {
	Cursor       string       `json:"cursor"`
	Job Job `json:"job"`
}

type CreateJobRequest struct {
	StartCommand string `json:"startCommand"`
}

// createJobCmd represents the createJob command
var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Create job by service id and start command",
	Long:  `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		startCommand := args[1]
		body := CreateJobRequest{StartCommand: startCommand}
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		jsonString, err := http.Create(fmt.Sprintf("services/%s/jobs", serviceId), bodyBytes)
		if err != nil {
			panic(err)
		}

		var job Job
		if err := json.Unmarshal(jsonString, &job); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "ServiceId", "StartCommand", "CreatedAt", "FinishedAt", "Status"}, []Job{job}); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createJobCmd)
}
