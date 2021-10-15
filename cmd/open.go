package cmd

import (
	"errors"
	"fmt"

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
		var url string
		if len(cds) != 0 {
			url = fmt.Sprintf("https://%s", cds[0].Name)
		} else {
			svc := service(serviceId)
			url = fmt.Sprintf("https://%s.onrender.com", svc.Slug)
		}
		if err := open.Run(url); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
