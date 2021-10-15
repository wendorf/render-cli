package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

type CreateServiceRequest struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Env string `json:"env"`
	Repo string `json:"repo"`
	OwnerId string `json:"ownerId"`
}

var svcType string
var name string
var env string
var repo string
var ownerId string

// createServiceCmd represents the createService command
var createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Create service",
	Long:  `It's hack day!'`,
	Run: func(cmd *cobra.Command, args []string) {
		body := CreateServiceRequest{
			Type: svcType,
			Name: name,
			Env: env,
			Repo: repo,
			OwnerId: ownerId,
		}
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bodyBytes))
		jsonString, err := http.Create("services", bodyBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonString))

		var service Service
		if err := json.Unmarshal(jsonString, &service); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "Name", "Type", "State"}, []Service{service}); err != nil {
			panic(err)
		}
	},
}

func init() {
	createServiceCmd.Flags().StringVarP(&svcType, "type", "t", "", "Type (required)")
	_ = createServiceCmd.MarkFlagRequired("type")
	createServiceCmd.Flags().StringVarP(&name, "name", "n", "", "Name (required)")
	_ = createServiceCmd.MarkFlagRequired("name")
	createServiceCmd.Flags().StringVarP(&env, "env", "e", "", "Env (required)")
	_ = createServiceCmd.MarkFlagRequired("env")
	createServiceCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo (required)")
	_ = createServiceCmd.MarkFlagRequired("repo")
	createServiceCmd.Flags().StringVarP(&ownerId, "owner-id", "o", "", "Owner Id (required)")
	_ = createServiceCmd.MarkFlagRequired("owner-id")
	rootCmd.AddCommand(createServiceCmd)
}
