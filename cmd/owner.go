package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

// ownerCmd represents the owner command
var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Get owner by id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ownerId := args[0]
		jsonString, err := http.Request(fmt.Sprintf("owners/%s", ownerId))
		if err != nil {
			panic(err)
		}

		var owner Owner
		if err := json.Unmarshal(jsonString, &owner); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "Name", "Email", "Type"}, []Owner{owner}); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ownerCmd)
}
