package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/renderinc/cli/pkg/http"
	"github.com/renderinc/cli/pkg/table"
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open service by id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]
		cds := customDomains(serviceId)
		if len(cds) != 0 {
			if err := open.Run(fmt.Sprintf("https://%s", cds[0].Name)); err != nil {
				panic(err)
			}
			return
		}

		svc := service(serviceId)
		if err := open.Run(fmt.Sprintf("https://%s.onrender.com", svc.Slug)); err != nil {
			panic(err)
		}
		jsonString, err := http.Request(fmt.Sprintf("services/%s", serviceId))
		if err != nil {
			panic(err)
		}

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
	rootCmd.AddCommand(openCmd)
}
