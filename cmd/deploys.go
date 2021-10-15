package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

type Commit struct {
	Id string `json:"id"`
	Message string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

type Deploy struct {
	Id string `json:"id"`
	Commit Commit `json:"commit"`
	Status string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	FinishedAt string `json:"finishedAt"`
}

type CursorDeploy struct {
	Cursor string `json:"cursor"`
	Deploy Deploy `json:"deploy"`
}

// deploysCmd represents the deploys command
var deploysCmd = &cobra.Command{
	Use:   "deploys",
	Short: "Get deploys by service id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		jsonString, err := http.Request(fmt.Sprintf("services/%s/deploys", serviceId))
		if err != nil {
			panic(err)
		}

		var cursorDeploys []CursorDeploy
		if err := json.Unmarshal(jsonString, &cursorDeploys); err != nil {
			panic(err)
		}

		var deploys []Deploy
		for _, cursorDeploy := range cursorDeploys {
			deploys = append(deploys, cursorDeploy.Deploy)
		}

		if err := table.Print([]string{"Id", "Status", "CreatedAt", "FinishedAt"}, deploys); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deploysCmd)
}
