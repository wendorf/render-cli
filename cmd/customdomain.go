package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
)

// customDomainCmd represents the customDomains command
var customDomainCmd = &cobra.Command{
	Use:   "custom-domain",
	Short: "Get custom domains by service id and custom domain id or name",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		customDomainIdOrName := args[1]
		jsonString, err := http.Request(fmt.Sprintf("services/%s/custom-domains/%s", serviceId, customDomainIdOrName))
		if err != nil {
			panic(err)
		}

		var customDomain CustomDomain
		if err := json.Unmarshal(jsonString, &customDomain); err != nil {
			panic(err)
		}

		if err := table.Print([]string{"Id", "Name", "CreatedAt"}, []CustomDomain{customDomain}); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(customDomainCmd)
}
