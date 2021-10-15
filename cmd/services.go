package cmd

import (
	"encoding/json"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

type Service struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Repo string `json:"repo"`
	Name string `json:"name"`
	AutoDeploy bool `json:"autoDeploy"`
	Branch string `json:"branch"`
	CreatedAt string `json:"createdAt"`
	NotifyOnFail string `json:"notifyOnFail"`
	OwnerId string `json:"ownerId"`
	Slug string `json:"slug"`
	State string `json:"state"`
	UpdatedAt string `json:"updatedAt"`
}

type CursorService struct {
	Cursor string `json:"cursor"`
	Service Service `json:"service"`
}

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List services",
	Long: `It's hackday!'`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonString, err := http.Request("services")
		if err != nil {
			panic(err)
		}

		var cursorServices []CursorService
		if err := json.Unmarshal(jsonString, &cursorServices); err != nil {
			panic(err)
		}

		var services []Service
		for _, cursorService := range cursorServices {
			services = append(services, cursorService.Service)
		}

		if err := table.Print([]string{"Id", "Name", "Type", "State"}, services); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
