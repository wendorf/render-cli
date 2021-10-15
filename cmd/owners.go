package cmd

import (
	"encoding/json"

	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"

	"github.com/renderinc/cli/pkg/http"
)

type Owner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Type string `json:"type"`
}

var ownersCmd = &cobra.Command{
	Use:   "owners",
	Short: "List owners",
	Long: `It's hack day!'`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonString, err := http.Request("owners")
		if err != nil {
			panic(err)
		}

		var owners []Owner
		if err := json.Unmarshal(jsonString, &owners); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "Name", "Email", "Type"}, owners); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ownersCmd)
}
