package cmd

import (
	"errors"
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var dashboardTypeMap = map[string]string{
	"static_site": "static",
	"web_service": "web",
	"private_service": "pserv",
	"background_worker": "worker",
	"cron_job": "cron",
}

// dashboardCmd represents the open command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open service on the dashboard by id",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceId := args[0]

		svc := service(serviceId)
		if err := open.Run(fmt.Sprintf("https://dashboard.render.com/%s/%s", dashboardTypeMap[svc.Type], svc.Id)); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
